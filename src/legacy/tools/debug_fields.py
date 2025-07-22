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
actual_map = field_by_offset(struct200_node)
expected_map = {
    0: 'field_0',
    4: 'field_1',  
    8: 'field_2',
    12: 'field_3',
    16: 'field_4',
    48: 'field_12',
    52: 'field_13',
    56: 'field_14',
    60: 'field_15',
    64: 'field_16',
    68: 'field_17',
    72: 'field_18',
    76: 'field_19',
    80: 'field_20',
    84: 'snd_name',
    88: 'field_22',
    100: 'field_25',
    104: 'field_26',
    108: 'field_27',
    112: 'field_28',
    124: 'field_31',
    128: 'field_32',
    192: 'field_48',
    196: 'field_49',
}

print('Actual map:')
for k in sorted(actual_map.keys()):
    print(f'  {k}: "{actual_map[k]}"')

print()
print('Expected map:')
for k in sorted(expected_map.keys()):
    print(f'  {k}: "{expected_map[k]}"')

print()
print('Differences:')
all_keys = set(actual_map.keys()) | set(expected_map.keys())
for k in sorted(all_keys):
    if k not in actual_map:
        print(f'  Missing in actual: {k}: "{expected_map[k]}"')
    elif k not in expected_map:
        print(f'  Extra in actual: {k}: "{actual_map[k]}"')
    elif actual_map[k] != expected_map[k]:
        print(f'  Different: {k}: actual="{actual_map[k]}" expected="{expected_map[k]}"')