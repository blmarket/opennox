#%%
import os
import sys
import tree_sitter_c as tsc
from tree_sitter import Language, Parser, Node

os.chdir(os.path.join(os.path.dirname(__file__), ".."))

from tools.inspect import dbg, print_tree
from tools.ast_utils import find_function_definition, find_struct_definition, find_variable_references, enumerate_func_definitions
from tools.type_analysis import find_variable_type, infer_from_func
from tools.struct_utils import field_by_offset
from tools.translator import SourceFile, TranslationContext


# %%
# Parse defs.h to find struct200 definition
defs_file = SourceFile("defs.h")

#%%
struct200_node = find_struct_definition(defs_file.root_node, "struct200", defs_file.source_code)
struct200_field_map = field_by_offset(struct200_node, defs_file.source_code)
struct200_field_map[92] = "field_22.field_1"
struct200_field_map[20] = "field_4.field_1"
struct200_field_map

#%%
struct576_field_map = field_by_offset(find_struct_definition(defs_file.root_node, "struct576", defs_file.source_code), defs_file.source_code)
struct576_field_map[188] = "timerGroup_46.field_0.field_1"
struct576_field_map[248] = "timerGroup_46.field_16"
struct576_field_map

#%%
structs_map = {
    "struct200*": struct200_field_map,
    "struct576*": struct576_field_map,
}

SIZE_MAP = {
    "uint32_t*": 4,
    "int": 1,
}

#%%
FILE_PATH = "audio.c"


#%%
while True:
    has_change = False
    source_file = SourceFile(FILE_PATH)
    nodes = enumerate_func_definitions(source_file.root_node)

    """Reverse order so that modification in later node won't affect earlier node"""
    for func_node in nodes[::-1]:
        SRC, TGT = infer_from_func(func_node, source_file.source_code)
        if SRC is None or TGT is None:
            continue
        src_type = find_variable_type(func_node, SRC, source_file.source_code)
        tgt_type = find_variable_type(func_node, TGT, source_file.source_code)
        print(f"Variable '{SRC}' type: {src_type}")
        print(f"Variable '{TGT}' type: {tgt_type}")

        if tgt_type not in structs_map:
            continue

        src_references = find_variable_references(func_node, SRC, source_file.source_code)

        context = source_file.create_translation_context(
            src=SRC,
            tgt=TGT,
            src_type=src_type,
            tgt_type=tgt_type,
            structs_map=structs_map,
            size_map=SIZE_MAP
        )

        if len(src_references) <= 1:
            continue

        if context.replace_translation( src_references[1:][-1]):
            has_change = True
    if not has_change:
        break


# %%
