#%%
import os
import sys
import tree_sitter_c as tsc
from tree_sitter import Language, Parser, Node

os.chdir(os.path.join(os.path.dirname(__file__), ".."))

from tools.inspect import dbg, print_tree
from tools.ast_utils import find_function_definition, find_struct_definition, find_variable_references
from tools.type_analysis import find_variable_type, infer_from_func
from tools.struct_utils import field_by_offset
from tools.translator import SourceFile, TranslationContext


#%%
FILE_PATH = "audio.c"
# FUNC_NAME = "sub_451920"
# FUNC_NAME = "sub_452190"
# FUNC_NAME = "sub_451CA0"
# FUNC_NAME = "sub_4523D0"
FUNC_NAME = "sub_452410"
FUNC_NAME = "sub_452490"
FUNC_NAME = "sub_452510"
FUNC_NAME = "sub_452690"
FUNC_NAME = "nox_xxx_draw_452300"
FUNC_NAME = "sub_452E90"
FUNC_NAME = "sub_452F10"
FUNC_NAME = "sub_451F30"

FUNC_NAME = sys.argv[-1]

#%%
source_file = SourceFile(FILE_PATH)

#%%
# node = source_file.root_node.children[62]
# print(dbg(node, source_file.source_code))
# print_tree(node, source_file.source_code)

#%%
func_node = source_file.find_function_definition(FUNC_NAME)
if func_node:
    print(f"Found function '{FUNC_NAME}' at line {func_node.start_point[0] + 1}")
else:
    raise RuntimeError(f"Function '{FUNC_NAME}' not found")

#%%
SRC, TGT = infer_from_func(func_node, source_file.source_code)

#%%
src_type = find_variable_type(func_node, SRC, source_file.source_code)
tgt_type = find_variable_type(func_node, TGT, source_file.source_code)

print(f"Variable '{SRC}' type: {src_type}")
print(f"Variable '{TGT}' type: {tgt_type}")

# #%%
# # Debug: Print parse tree structure of function node

# print("\nFunction parse tree structure:")
# print_tree(func_node, source_file.source_code, max_depth=4)

#%%
src_references = find_variable_references(func_node, SRC, source_file.source_code)
print(f"\nFound {len(src_references)} references to '{SRC}':")

for i, ref in enumerate(src_references):
    line_num = ref.start_point[0] + 1
    
    # Get the full statement context by walking up to expression_statement or declaration
    parent = ref.parent
    while parent and parent.type not in ["expression_statement", "declaration", "return_statement"]:
        parent = parent.parent
    
    if parent:
        context = source_file.source_code[parent.start_byte:parent.end_byte]
        # Remove semicolon and clean up whitespace
        context = context.rstrip(';').strip()
    else:
        context = source_file.source_code[ref.start_byte:ref.end_byte]
    
    print(f"  {i+1}. Line {line_num}: {context}")

# %%
# Parse defs.h to find struct200 definition
defs_file = SourceFile("defs.h")

#%%
struct200_node = find_struct_definition(defs_file.root_node, "struct200", defs_file.source_code)
struct200_field_map = field_by_offset(struct200_node, defs_file.source_code)
struct200_field_map

# %%
# struct200 field mapping validated - see test_struct_utils.py for assertions

#%%
struct576_field_map = field_by_offset(find_struct_definition(defs_file.root_node, "struct576", defs_file.source_code), defs_file.source_code)
struct576_field_map[188] = "timerGroup_46.field_0.field_1"
struct576_field_map

#%%
# struct576 field mapping validated - see test_struct_utils.py for assertions

#%%
structs_map = {
    "struct200*": struct200_field_map,
    "struct576*": struct576_field_map,
}

#%%
SIZE_MAP = {
    "uint32_t*": 4,
    "int": 1,
}

#%%
# Create translation context and update the translation call
context = source_file.create_translation_context(
    src=SRC,
    tgt=TGT,
    src_type=src_type,
    tgt_type=tgt_type,
    structs_map=structs_map,
    size_map=SIZE_MAP
)

context.replace_translation(
    src_references[1:][-1]
)

# %%
# ref = src_references[1].parent
# print_tree(ref, source_file.source_code)

# %%
