#%%
import os
import sys
import tree_sitter_c as tsc
from tree_sitter import Language, Parser, Node

os.chdir(os.path.join(os.path.dirname(__file__), ".."))

from tools.type_analysis import find_variable_type, infer_from_func
from tools.translator import SourceFile

#%%
src = SourceFile("audio.c")

#%%
FUNC_NAME = 'sub_452050'
func_node = src.find_function_definition(FUNC_NAME)
if func_node:
    print(f"Found function '{FUNC_NAME}' at line {func_node.start_point[0] + 1}")
else:
    raise RuntimeError(f"Function '{FUNC_NAME}' not found")

#%%
print(infer_from_func(func_node, src.source_code))
assert ("v1", "v1p") == infer_from_func(func_node, src.source_code), "Failed to extract variables"

print(find_variable_type(func_node, "v1p", src.source_code))
# %%
