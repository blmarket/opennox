from tree_sitter import Node, Parser, Language
import tree_sitter_c as tsc

from tools.inspect import dbg, print_tree
from tools.ast_utils import find_function_definition, find_struct_definition, find_variable_references
from tools.type_analysis import find_variable_type, infer_from_func
from tools.struct_utils import field_by_offset

class TranslationContext(object):
    def __init__(self, src: str, tgt: str, src_type: str, tgt_type: str, 
                 structs_map: dict, size_map: dict, source_file):
        self._src = src
        self._tgt = tgt
        self._src_type = src_type
        self._tgt_type = tgt_type
        self._structs_map = structs_map
        self._size_map = size_map
        self._source_file = source_file
    
    @property
    def src(self) -> str:
        return self._src
    
    @property
    def tgt(self) -> str:
        return self._tgt
    
    @property
    def src_type(self) -> str:
        return self._src_type
    
    @property
    def tgt_type(self) -> str:
        return self._tgt_type
    
    @property
    def structs_map(self) -> dict:
        return self._structs_map
    
    @property
    def size_map(self) -> dict:
        return self._size_map
    
    @property
    def source_file(self):
        return self._source_file

    def _extract_index_from_multiplication(self, array_node: Node) -> str:
        """
        Extract index from binary expression that starts with "4 *".
        Returns the right operand of the multiplication as the index.
        
        For example:
        - "4 * v2" -> "v2"
        - "4 * *(uint32_t*)(&a1_->field_42)" -> "*(uint32_t*)(&a1_->field_42)"
        """
        if array_node.type != "binary_expression":
            return None
        
        source_code = self.source_file.source_code
        
        # Check if this is a multiplication with 3 children: left * right
        if len(array_node.children) != 3:
            return None
        
        left_operand = array_node.children[0]
        operator = array_node.children[1]
        right_operand = array_node.children[2]
        
        # Check if operator is multiplication
        if source_code[operator.start_byte:operator.end_byte] != "*":
            return None
        
        # Check if left operand is "4"
        left_text = source_code[left_operand.start_byte:left_operand.end_byte]
        if left_text != "4":
            return None
        
        # Return the right operand as the index
        return source_code[right_operand.start_byte:right_operand.end_byte]

    def replace_translation(self, node: Node):
        """ 
        Given the node, perform in-place translation in the code. Update file in source_file.
        Returns the updated source_code.
        """
        if not self.source_file:
            raise ValueError("TranslationContext must have a source_file to perform replacement")

        node = node.parent
        if node.type == "binary_expression" and node.children[1].type == '+' and node.parent.type == "binary_expression" and node.parent.children[1].type == '+':
            # Special support for 3 entries addition case
            node = node.parent
        
        source_code = self.source_file.source_code
        file_path = self.source_file.file_path
        
        # Get the original text and its translation
        translated_text = self.translate(node)
        
        # Only proceed if there's actually a change to make
        original_text = source_code[node.start_byte:node.end_byte]
        if original_text == translated_text:
            return source_code
        
        # Replace the text in source_code
        new_source_code = (
            source_code[:node.start_byte] + 
            translated_text + 
            source_code[node.end_byte:]
        )
        
        # Update the source file's content
        self.source_file.source_code = new_source_code
        
        # Write the updated code back to the file
        with open(file_path, 'w') as f:
            f.write(new_source_code)
        
        print(f"Replaced '{original_text}' with '{translated_text}' in {file_path}")

    def translate(self, node: Node) -> str:
        """
        Translate C pointer/array access to struct field access.
        
        Converts patterns like:
        - *a2 -> a2_->field_0
        - a2[1] -> a2_->field_1 
        - a2[2] -> a2_->field_2
        - a2 + 4 -> &a2_->field_4
        etc.
        """
        assert self.tgt_type in self.structs_map, "unknown struct type"
        source_code = self.source_file.source_code
        struct = self.structs_map[self.tgt_type]

        if node.type == "pointer_expression":
            # Handle *a2 -> a2_->field_0 (offset 0)
            return f"{self.tgt}->field_0"
        
        if node.type == "argument_list":
            ret = ""
            for child in node.children:
                child_txt = source_code[child.start_byte:child.end_byte]
                if child.type == "identifier" and child_txt == self.src:
                    ret += self.tgt
                else:
                    ret += child_txt
            return ret
        
        elif node.type == "subscript_expression":
            # Handle a2[index] -> a2_->field_X
            # Find the index value
            for child in node.children:
                if child.type == "number_literal":
                    index = int(source_code[child.start_byte:child.end_byte])
                    # Calculate byte offset: index * 4 (since we're dealing with uint32_t*)
                    byte_offset = index * 4
                    
                    # Look up field name from struct200_field_map
                    assert byte_offset in struct, f"cannot resolve offset {byte_offset}"
                    field_name = struct[byte_offset]
                    return f"{self.tgt}->{field_name}"
        
        elif node.type == "binary_expression":
            # Handle pointer arithmetic like a2 + 4 -> &a2_->field_4
            # Check if this is an addition with a number literal
            operator = None
            identifier_node = None
            number_node = None

            if node.children[0].type == "binary_expression":
                # Handle 3-term addition: base + array + offset -> dst->field_array[idx]
                if self.src_type != 'int':
                    print(f"Warning: src_type is '{self.src_type}', not 'int'. Skipping translation.")
                    return source_code[node.start_byte:node.end_byte]
                
                # Parse the nested binary expression (first two terms)
                left_expr = node.children[0]
                operator2 = node.children[1] if len(node.children) > 1 else None
                right_term = node.children[2] if len(node.children) > 2 else None

                if operator2.type != '+':
                    print(f"Warning: unexpected operator")
                    return source_code[node.start_byte:node.end_byte]

                # Extract components from left expression (base + array)
                base_node = None
                array_node = None
                offset_node = right_term
                
                for child in left_expr.children:
                    if child.type == "identifier" and source_code[child.start_byte:child.end_byte] == self.src:
                        base_node = child
                    elif child.type == "binary_expression":
                        array_node = child

                offset = int(source_code[offset_node.start_byte:offset_node.end_byte])

                # Extract index from binary expression starting with "4 *"
                idx = self._extract_index_from_multiplication(array_node)
                if idx is None:
                    print(f"Warning: cannot resolve idx from {self.source_file.dbg(array_node)}")
                    return source_code[node.start_byte:node.end_byte]

                return f"&{self.tgt}->{struct[offset]}[{idx}]"
            
            for child in node.children:
                if child.type == "identifier" and source_code[child.start_byte:child.end_byte] == self.src:
                    identifier_node = child
                elif child.type == "number_literal":
                    number_node = child
                elif child.type == "+":
                    operator = "+"
            
            if operator == "+" and identifier_node and number_node:
                # Calculate the offset in elements (assuming uint32_t* arithmetic)
                offset = int(source_code[number_node.start_byte:number_node.end_byte])
                assert self.src_type in self.size_map, f"FIXME: Current code does not support {self.src_type} type"
                byte_offset = offset * self.size_map[self.src_type]
                
                assert byte_offset in struct, f"cannot resolve offset {byte_offset}"
                field_name = struct[byte_offset]
                return f"&{self.tgt}->{field_name}"

            if node.children[1].type in ["==", "!="]:
                ret = ""
                for child in node.children:
                    child_txt = source_code[child.start_byte:child.end_byte]
                    print(child.type, child_txt)
                    if child.type == "identifier" and child_txt == self.src:
                        ret += self.tgt
                        continue
                    ret += child_txt
                return ret

        elif node.type == "cast_expression" and node.parent.type == "argument_list":
            return self.tgt

        elif node.type == "assignment_expression":
            ret = ""
            for child in node.children:
                child_txt = source_code[child.start_byte:child.end_byte]
                if child.type == "identifier" and child_txt == self.src:
                    ret += self.tgt
                    continue
                ret += child_txt
            return ret

        elif node.type == "parenthesized_expression" and len(node.children) == 3:
            child = node.children[1]
            child_txt = source_code[child.start_byte:child.end_byte]
            if child_txt == self.src:
                return f"({self.tgt})"

                
        # Fallback for other cases
        return source_code[node.start_byte:node.end_byte]


class SourceFile(object):
    def __init__(self, file_path: str):
        self.file_path = file_path
        with open(file_path, 'r') as f:
            self.source_code = f.read()
        
        self.parser = Parser()
        self.parser.set_language(Language(tsc.language(), "c"))
        self.tree = self.parser.parse(bytes(self.source_code, "utf8"))
        self.root_node = self.tree.root_node

    def dbg(self, node: Node):
        return dbg(node, self.source_code)
    
    def find_function_definition(self, func_name: str) -> Node:
        """Find function definition by name in this source file"""
        return find_function_definition(self.root_node, func_name, self.source_code)
    
    def create_translation_context(self, src: str, tgt: str, src_type: str, tgt_type: str, 
                                 structs_map: dict, size_map: dict) -> TranslationContext:
        """Create a TranslationContext associated with this SourceFile"""
        return TranslationContext(src, tgt, src_type, tgt_type, structs_map, size_map, self)

