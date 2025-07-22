#%%
import os
import sys
import tree_sitter_c as tsc
from tree_sitter import Language, Parser, Node

os.chdir(os.path.join(os.path.dirname(__file__), ".."))

#%%
FILE_PATH = "GAME3_2.c"

#%%
# Initialize tree-sitter parser
C_LANGUAGE = Language(tsc.language(), "c")
parser = Parser()
parser.set_language(C_LANGUAGE)

# Read and parse the file
with open(FILE_PATH, 'r') as f:
    source_code = f.read()

tree = parser.parse(bytes(source_code, "utf8"))
root_node = tree.root_node

#%%
root_node.has_error
# node = root_node.children[62]
# print(dbg(node, source_code))
# print_tree(node, source_code)

#%%
def find_errors(node):
    if node.type == 'ERROR':
        print(f"Error found at line {node.start_point[0]+1}, column {node.start_point[1]+1}: {node.text.decode('utf8')}")
    for child in node.children:
        find_errors(child)

find_errors(root_node)
