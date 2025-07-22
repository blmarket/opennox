import tree_sitter_c as tsc
from tree_sitter import Language, Parser, Node

C_LANGUAGE = Language(tsc.language(), 'c')
parser = Parser()
parser.set_language(C_LANGUAGE)

with open('defs.h', 'r') as f:
    defs_source = f.read()

defs_tree = parser.parse(bytes(defs_source, 'utf8'))

def find_struct_definition(node, struct_name):
    if node.type == 'struct_specifier':
        for child in node.children:
            if child.type == 'type_identifier' and defs_source[child.start_byte:child.end_byte] == struct_name:
                return node
    
    for child in node.children:
        result = find_struct_definition(child, struct_name)
        if result:
            return result
    return None

def calculate_field_offsets_debug(struct_node):
    """Debug version of calculate_field_offsets"""
    fields = []
    offset = 0
    
    # Find the field declaration list
    field_list = None
    for child in struct_node.children:
        if child.type == "field_declaration_list":
            field_list = child
            break
    
    if not field_list:
        return fields
    
    for field_decl in field_list.children:
        if field_decl.type == "field_declaration":
            # Extract type and field names
            field_type = None
            field_names = []
            size = None  # Initialize size as None
            
            field_text = defs_source[field_decl.start_byte:field_decl.end_byte]
            print(f"Processing field: {field_text.strip()}")
            
            for child in field_decl.children:
                print(f"  Child type: {child.type}, text: '{defs_source[child.start_byte:child.end_byte]}'")
                
                if child.type in ["primitive_type", "type_identifier", "struct_specifier"]:
                    field_type = defs_source[child.start_byte:child.end_byte]
                    print(f"    Set field_type to: {field_type}")
                elif child.type == "field_identifier":
                    field_names.append(defs_source[child.start_byte:child.end_byte])
                    print(f"    Added field_name: {defs_source[child.start_byte:child.end_byte]}")
                elif child.type == "array_declarator":
                    print(f"    Found array_declarator!")
                    # Handle array declarators like field_32[32] as direct children of field_declaration
                    array_size = 1
                    field_name = None
                    for array_child in child.children:
                        print(f"      Array child type: {array_child.type}, text: '{defs_source[array_child.start_byte:array_child.end_byte]}'")
                        if array_child.type == "field_identifier":
                            field_name = defs_source[array_child.start_byte:array_child.end_byte]
                            print(f"        Array field_name: {field_name}")
                        elif array_child.type == "number_literal":
                            array_size = int(defs_source[array_child.start_byte:array_child.end_byte])
                            print(f"        Array size: {array_size}")
                    
                    if field_name:
                        field_names.append(field_name)
                        print(f"        Added array field_name: {field_name}")
                        # For arrays, multiply the base type size by array size
                        base_size = 4  # default
                        if field_type == "uint16_t":
                            base_size = 2
                        elif field_type == "uint8_t" or field_type == "char":
                            base_size = 1
                        elif field_type == "uint32_t" or field_type == "int":
                            base_size = 4
                        elif field_type == "uint64_t":
                            base_size = 8
                        
                        # Override the size calculation for arrays
                        size = base_size * array_size
                        print(f"        Calculated array size: {size} (base_size={base_size} * array_size={array_size})")
            
            print(f"  Final: field_type={field_type}, field_names={field_names}, size={size}")
            
            if 'field_32' in str(field_names):
                print(f"  *** This is field_32! Breaking here for debug ***")
                break
    
    return fields

struct200_node = find_struct_definition(defs_tree.root_node, 'struct200')
if struct200_node:
    print("Found struct200, debugging field processing:")
    calculate_field_offsets_debug(struct200_node)