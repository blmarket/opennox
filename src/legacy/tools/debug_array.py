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

def print_tree(node, depth=0, max_depth=4):
    if depth > max_depth:
        return
    
    indent = '  ' * depth
    node_text = defs_source[node.start_byte:node.end_byte]
    if len(node_text) > 50:
        node_text = node_text[:50] + '...'
    node_text = node_text.replace('\n', '\\n').replace('\t', '\\t')
    
    print(f'{indent}{node.type}: "{node_text}"')
    
    for child in node.children:
        print_tree(child, depth + 1, max_depth)

struct200_node = find_struct_definition(defs_tree.root_node, 'struct200')
if struct200_node:
    field_list = None
    for child in struct200_node.children:
        if child.type == 'field_declaration_list':
            field_list = child
            break
    
    if field_list:
        count = 0
        for i, field_decl in enumerate(field_list.children):
            if field_decl.type == 'field_declaration':
                field_text = defs_source[field_decl.start_byte:field_decl.end_byte]
                if 'field_32' in field_text:
                    print(f'Found field_32 at field {count}:')
                    print_tree(field_decl)
                    print()
                    print('Children of field_declaration:')
                    for child in field_decl.children:
                        print(f'  {child.type}: "{defs_source[child.start_byte:child.end_byte]}"')
                    break
                count += 1