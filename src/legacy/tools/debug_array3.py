import tree_sitter_c as tsc
from tree_sitter import Language, Parser, Node

C_LANGUAGE = Language(tsc.language(), 'c')
parser = Parser()
parser.set_language(C_LANGUAGE)

with open('defs.h', 'r') as f:
    defs_source = f.read()

defs_tree = parser.parse(bytes(defs_source, 'utf8'))

exec(open('alias.py').read().split('struct200_field_map')[0])

struct200_node = find_struct_definition(defs_tree.root_node, 'struct200')
if struct200_node:
    print("Testing calculate_field_offsets with full processing:")
    fields = calculate_field_offsets(struct200_node)
    
    print(f"\nAll calculated fields ({len(fields)} total):")
    for i, field in enumerate(fields):
        print(f"  {i:2d}: +{field['offset']:3d}: {field['type']:<16} {field['name']}")
        if field['name'] == 'field_32':
            print(f"       *** Found field_32 at offset {field['offset']} with size {field['size']} ***")

    # Check the map
    offset_map = field_by_offset(struct200_node)
    print(f"\nfield_by_offset map ({len(offset_map)} entries):")
    for k in sorted(offset_map.keys()):
        marker = " *** ARRAY ***" if offset_map[k] == 'field_32' else ""
        print(f"  {k:3d}: \"{offset_map[k]}\"{marker}")