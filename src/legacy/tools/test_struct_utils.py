import os
import pytest
import tree_sitter_c as tsc
from tree_sitter import Language, Parser

from .ast_utils import find_struct_definition
from .struct_utils import field_by_offset


@pytest.fixture
def parser():
    """Setup tree-sitter parser for C language."""
    C_LANGUAGE = Language(tsc.language(), "c")
    parser = Parser()
    parser.set_language(C_LANGUAGE)
    return parser


@pytest.fixture
def defs_source(parser):
    """Load and parse defs.h file."""
    DEFS_FILE = os.path.join(os.path.dirname(__file__), "../defs.h")
    with open(DEFS_FILE, 'r') as f:
        defs_source = f.read()
    
    defs_tree = parser.parse(bytes(defs_source, "utf8"))
    defs_root = defs_tree.root_node
    
    return defs_source, defs_root


class TestStruct200FieldMapping:
    """Test struct200 field offset mappings."""
    
    def test_struct200_field_offsets(self, defs_source, parser):
        """Test that struct200 field offsets match expected values."""
        defs_source_text, defs_root = defs_source
        
        struct200_node = find_struct_definition(defs_root, "struct200", defs_source_text)
        struct200_field_map = field_by_offset(struct200_node, defs_source_text)
        
        expected_struct200_field_map = {
            0: "field_0",        # uint32_t field_0
            4: "field_1",        # uint32_t field_1  
            8: "field_2",        # uint32_t field_2
            12: "field_3",       # uint32_t field_3
            16: "field_4",       # timer field_4 (32 bytes)
            48: "field_12",      # uint32_t field_12
            52: "field_13",      # uint32_t field_13
            56: "field_14",      # uint32_t field_14
            60: "field_15",      # uint32_t field_15
            64: "field_16",      # uint32_t field_16
            68: "field_17",      # uint32_t field_17
            72: "field_18",      # uint32_t field_18
            76: "field_19",      # uint32_t field_19
            80: "field_20",      # uint32_t field_20
            84: "snd_name",      # uint32_t snd_name (sound name field accessed in sub_451850)
            88: "field_22",      # nox_list_item_t field_22 (12 bytes)
            100: "field_25",     # uint32_t field_25
            104: "field_26",     # uint32_t field_26
            108: "field_27",     # uint32_t field_27
            112: "field_28",     # nox_list_item_t field_28 (12 bytes)
            124: "field_31",     # uint32_t field_31
            128: "field_32",     # uint16_t field_32[32] (64 bytes)
            192: "field_48",     # uint32_t field_48
            196: "field_49",     # uint32_t field_49
        }
        
        assert struct200_field_map == expected_struct200_field_map


class TestStruct576FieldMapping:
    """Test struct576 field offset mappings."""
    
    def test_struct576_field_offsets(self, defs_source, parser):
        """Test that struct576 field offsets match expected values."""
        defs_source_text, defs_root = defs_source
        
        struct576_node = find_struct_definition(defs_root, "struct576", defs_source_text)
        struct576_field_map = field_by_offset(struct576_node, defs_source_text)
        
        assert struct576_field_map[0] == 'next', "The first field should be `next`"
        assert struct576_field_map[40] == 'field_10', "field at offset 40 should be field_10"
        assert struct576_field_map[168] == 'field_42', "field at offset 168 should be field_42"
    
    def test_struct576_specific_field_mappings(self, defs_source, parser):
        """Test specific struct576 field mappings."""
        defs_source_text, defs_root = defs_source
        
        struct576_node = find_struct_definition(defs_root, "struct576", defs_source_text)
        struct576_field_map = field_by_offset(struct576_node, defs_source_text)
        
        # Test first field
        assert struct576_field_map[0] == 'next', "The first field should be `next`"
        
        # Test field at offset 40
        assert struct576_field_map[40] == 'field_10', "field at offset 40 should be field_10"
        
        # Test field at offset 168
        assert struct576_field_map[168] == 'field_42', "field at offset 168 should be field_42"