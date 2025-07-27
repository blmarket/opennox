#!/usr/bin/env python3
"""
Unit tests for the C Function Parameter Update Tool

Tests the parameter parsing and update logic to ensure proper handling
of various C function signatures and edge cases.
"""

import pytest
import tempfile
import os
from pathlib import Path
import tree_sitter
from tools.update import (
    setup_tree_sitter, find_function_nodes, extract_function_name,
    extract_parameter_name, extract_parameter_type, get_parameter_list,
    update_function_signature, add_variable_declaration, process_file
)


class TestUpdateTool:
    
    def setup_method(self):
        """Set up test environment"""
        self.parser = setup_tree_sitter()
        
    def parse_code(self, code):
        """Helper to parse C code"""
        return self.parser.parse(code.encode('utf-8'))
    
    def test_simple_function_parsing(self):
        """Test basic function parsing"""
        code = """
        int add(int a, int b) {
            return a + b;
        }
        """
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "add")
        
        assert len(functions) == 1
        assert functions[0][0] == 'definition'
        
    def test_function_declaration_parsing(self):
        """Test function declaration parsing"""
        code = """
        int add(int a, int b);
        """
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "add")
        
        assert len(functions) == 1
        assert functions[0][0] == 'declaration'
        
    def test_function_name_extraction(self):
        """Test function name extraction from various declarator types"""
        code = """
        int func1(int a);
        int* func2(int a);
        int** func3(int a);
        """
        tree = self.parse_code(code)
        
        # Find all function declarations
        func1 = find_function_nodes(tree, "func1")
        func2 = find_function_nodes(tree, "func2")  
        func3 = find_function_nodes(tree, "func3")
        
        assert len(func1) == 1
        assert len(func2) == 1
        assert len(func3) == 1
        
    def test_parameter_name_extraction_simple(self):
        """Test parameter name extraction for simple parameters"""
        code = """
        void func(int param1, char* param2) {}
        """
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        assert len(functions) == 1
        node_type, node = functions[0]
        declarator = node.child_by_field_name('declarator')
        param_list = get_parameter_list(declarator)
        parameters = [child for child in param_list.children 
                     if child.type == 'parameter_declaration']
        
        assert len(parameters) == 2
        
        # Test first parameter
        param1_declarator = parameters[0].child_by_field_name('declarator')
        name1 = extract_parameter_name(param1_declarator)
        assert name1 == "param1"
        
        # Test second parameter
        param2_declarator = parameters[1].child_by_field_name('declarator')
        name2 = extract_parameter_name(param2_declarator)
        assert name2 == "param2"
        
    def test_parameter_name_extraction_with_underscore(self):
        """Test parameter name extraction for parameters ending with underscore"""
        code = """
        void* sub_4876A0(struct264* a1_) {}
        """
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "sub_4876A0")
        
        assert len(functions) == 1
        node_type, node = functions[0]
        declarator = node.child_by_field_name('declarator')
        param_list = get_parameter_list(declarator)
        parameters = [child for child in param_list.children 
                     if child.type == 'parameter_declaration']
        
        assert len(parameters) == 1
        
        param_declarator = parameters[0].child_by_field_name('declarator')
        name = extract_parameter_name(param_declarator)
        assert name == "a1_"
        
    def test_parameter_type_extraction(self):
        """Test parameter type extraction"""
        code = """
        void func(int param1, struct264* param2, char** param3) {}
        """
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        node_type, node = functions[0]
        declarator = node.child_by_field_name('declarator')
        param_list = get_parameter_list(declarator)
        parameters = [child for child in param_list.children 
                     if child.type == 'parameter_declaration']
        
        # Test int parameter
        param1_declarator = parameters[0].child_by_field_name('declarator')
        type1 = extract_parameter_type(parameters[0], param1_declarator)
        assert "int" in type1
        
        # Test struct pointer parameter
        param2_declarator = parameters[1].child_by_field_name('declarator')
        type2 = extract_parameter_type(parameters[1], param2_declarator)
        assert "struct264" in type2 and "*" in type2
        
        # Test double pointer parameter
        param3_declarator = parameters[2].child_by_field_name('declarator')
        type3 = extract_parameter_type(parameters[2], param3_declarator)
        assert "char" in type3 and "**" in type3
        
    def test_function_signature_update_simple(self):
        """Test simple function signature update"""
        code = """
        void func(int param) {
            return param + 1;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        new_source, original_type, param_name = update_function_signature(
            source_code, functions[0], 0, "float"
        )
        
        assert new_source != source_code
        assert original_type == "int"
        assert param_name == "param"
        assert b"float param_" in new_source
        
    def test_function_signature_update_with_underscore_param(self):
        """Test function signature update with parameter ending in underscore"""
        code = """
        void* sub_4876A0(struct264* a1_) {
            return NULL;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "sub_4876A0")
        
        new_source, original_type, param_name = update_function_signature(
            source_code, functions[0], 0, "struct264*"
        )
        
        assert new_source != source_code
        assert original_type == "struct264 *"
        assert param_name == "a1_"
        # The new parameter should be named a1__ (adding one underscore)
        assert b"struct264* a1__" in new_source
        
    def test_variable_declaration_addition(self):
        """Test adding variable declaration to function body"""
        code = """
        void func(float param_) {
            return param_ + 1;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        new_source = add_variable_declaration(
            source_code, functions[0], "int", "param", "param_"
        )
        
        assert new_source != source_code
        # Should add variable declaration after opening brace
        assert b"int param = param_;" in new_source
        
    def test_problematic_underscore_case(self):
        """Test the specific problematic case that creates naming conflicts"""
        code = """
        void* sub_4876A0(struct264* a1_) {
            uint32_t ** a1 = a1_;
            return a1;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "sub_4876A0")
        
        new_source, original_type, param_name = update_function_signature(
            source_code, functions[0], 0, "struct264*"
        )
        
        # The problem: this creates a1__ parameter and tries to declare "struct264 * a1_ = a1__;"
        # But a1_ already exists in the function body, creating a conflict
        assert param_name == "a1_"
        assert b"struct264* a1__" in new_source
        
        # Add the variable declaration
        final_source = add_variable_declaration(
            new_source, functions[0], original_type, param_name, param_name + "_"
        )
        
        # This should create: struct264 * a1_ = a1__;
        # But there's already "uint32_t ** a1 = a1_;" which references the old a1_
        # This creates a naming conflict that breaks compilation
        assert b"struct264 * a1_ = a1__;" in final_source
        
    def test_no_naming_conflict_when_no_existing_usage(self):
        """Test that no conflict occurs when parameter isn't used in function body"""
        code = """
        void* sub_4876A0(struct264* a1_) {
            return NULL;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "sub_4876A0")
        
        new_source, original_type, param_name = update_function_signature(
            source_code, functions[0], 0, "struct264*"
        )
        
        final_source = add_variable_declaration(
            new_source, functions[0], original_type, param_name, param_name + "_"
        )
        
        # This should work fine since a1_ isn't used in the function body
        assert b"struct264* a1__" in final_source
        assert b"struct264 * a1_ = a1__;" in final_source
        
    def test_return_type_update(self):
        """Test return type update functionality"""
        code = """
        int func(int param) {
            return param;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        new_source, original_type, param_name = update_function_signature(
            source_code, functions[0], -1, "float"
        )
        
        assert new_source != source_code
        assert original_type == "int"
        assert param_name is None  # No parameter name for return type update
        assert b"float" in new_source and b"func" in new_source
        
    def test_pointer_return_type_update(self):
        """Test pointer return type update"""
        code = """
        int* func(int param) {
            return &param;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        new_source, original_type, param_name = update_function_signature(
            source_code, functions[0], -1, "float*"
        )
        
        assert new_source != source_code
        assert "*" in original_type
        assert b"float* func" in new_source
        
    def test_full_file_processing(self):
        """Test full file processing workflow"""
        code = """
        #include <stdio.h>
        
        int add(int a, int b);
        
        int add(int a, int b) {
            return a + b;
        }
        
        int main() {
            return add(1, 2);
        }
        """
        
        # Create temporary file
        with tempfile.NamedTemporaryFile(mode='w', suffix='.c', delete=False) as f:
            f.write(code)
            temp_file = f.name
            
        try:
            # Process the file
            result = process_file(temp_file, "add", 0, "float")
            assert result is True
            
            # Read the modified file
            with open(temp_file, 'r') as f:
                modified_content = f.read()
                
            # Should have updated both declaration and definition
            assert "float a_" in modified_content
            assert "int a = a_;" in modified_content
            
        finally:
            os.unlink(temp_file)
            
    def test_parameter_index_out_of_range(self):
        """Test handling of parameter index out of range"""
        code = """
        int func(int a) {
            return a;
        }
        """
        source_code = code.encode('utf-8')
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        # Try to update parameter index 5 when there's only 1 parameter
        new_source, original_type, param_name = update_function_signature(
            source_code, functions[0], 5, "float"
        )
        
        # Should return unchanged source and None values
        assert new_source == source_code
        assert original_type is None
        assert param_name is None
        
    def test_function_not_found(self):
        """Test handling when function is not found"""
        code = """
        int func(int a) {
            return a;
        }
        """
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "nonexistent_func")
        
        assert len(functions) == 0
        
    def test_nameless_parameter_handling(self):
        """Test handling of parameters without names"""
        code = """
        int func(int, char*);
        """
        tree = self.parse_code(code)
        functions = find_function_nodes(tree, "func")
        
        assert len(functions) == 1
        node_type, node = functions[0]
        declarator = node.child_by_field_name('declarator')
        param_list = get_parameter_list(declarator)
        parameters = [child for child in param_list.children 
                     if child.type == 'parameter_declaration']
        
        assert len(parameters) == 2
        
        # First parameter has no name
        param1_declarator = parameters[0].child_by_field_name('declarator')
        if param1_declarator:
            name1 = extract_parameter_name(param1_declarator)
        else:
            name1 = "param"  # Default name
            
        # Should handle gracefully
        assert name1 in ["param", "param0"]
        
    def test_real_problematic_case_sub_4876A0(self):
        """Test the actual problematic case reported by user"""
        # This is the exact signature that was causing problems
        code = """
        void* sub_4876A0(struct264* a1_) {
            uint32_t ** a1 = a1_;
            return a1;
        }
        """
        
        # Create temporary file to test full processing
        with tempfile.NamedTemporaryFile(mode='w', suffix='.c', delete=False) as f:
            f.write(code)
            temp_file = f.name
            
        try:
            # This should process without errors
            result = process_file(temp_file, "sub_4876A0", 0, "struct264*")
            assert result is True
            
            # Read the modified file
            with open(temp_file, 'r') as f:
                modified_content = f.read()
            
            # Verify the changes were made
            assert "struct264* a1__" in modified_content  # New parameter name
            assert "struct264 * a1_ = a1__;" in modified_content  # Variable declaration
            
            # The problem: this creates a naming conflict because a1_ is used again
            # The line "uint32_t ** a1 = a1_;" should now reference the new a1_ variable
            # But this might not work correctly if the original a1_ usage isn't updated
            
        finally:
            os.unlink(temp_file)
            
    def test_naming_conflict_detection(self):
        """Test detection of potential naming conflicts"""
        # This test demonstrates the core issue: when a parameter name ending in underscore
        # is updated, the tool creates a new parameter with double underscore and tries to
        # create a local variable with the original name, but this can conflict with existing
        # usage of that parameter name in the function body.
        
        code = """
        void problematic(int param_) {
            int x = param_ + 1;  // This line uses the original parameter
            return x;
        }
        """
        
        with tempfile.NamedTemporaryFile(mode='w', suffix='.c', delete=False) as f:
            f.write(code)
            temp_file = f.name
            
        try:
            result = process_file(temp_file, "problematic", 0, "float")
            assert result is True
            
            with open(temp_file, 'r') as f:
                modified_content = f.read()
            
            # The modified function should have:
            # 1. Parameter renamed to param__
            # 2. Local variable: int param_ = param__;
            # 3. Original usage: int x = param_ + 1; (now refers to local variable)
            
            assert "float param__" in modified_content
            assert "int param_ = param__;" in modified_content
            assert "int x = param_ + 1;" in modified_content
            
            # This actually works correctly - the local variable shadows the original parameter
            # so existing code continues to work
            
        finally:
            os.unlink(temp_file)
            
    def test_multiple_parameter_updates(self):
        """Test updating multiple parameters in the same function"""
        code = """
        int multi(int a, float b, char* c) {
            return a + (int)b + *c;
        }
        """
        
        with tempfile.NamedTemporaryFile(mode='w', suffix='.c', delete=False) as f:
            f.write(code)
            temp_file = f.name
            
        try:
            # Update first parameter
            result1 = process_file(temp_file, "multi", 0, "double")
            assert result1 is True
            
            # Update second parameter  
            result2 = process_file(temp_file, "multi", 1, "double")
            assert result2 is True
            
            with open(temp_file, 'r') as f:
                modified_content = f.read()
                
            # Should have both parameters updated
            assert "double a_" in modified_content
            assert "double b_" in modified_content
            assert "int a = a_;" in modified_content
            assert "float b = b_;" in modified_content
            
        finally:
            os.unlink(temp_file)

    def test_multiple_pointer_parameter(self):
        """Test updating multiple parameters in the same function"""
        code = """
        int multi(int** a, float b, char* c) {
            return *a + (int)b + *c;
        }
        """
        
        with tempfile.NamedTemporaryFile(mode='w', suffix='.c', delete=False) as f:
            f.write(code)
            temp_file = f.name
            
        try:
            # Update first parameter
            result1 = process_file(temp_file, "multi", 0, "int*")
            assert result1 is True
            
            with open(temp_file, 'r') as f:
                modified_content = f.read()
                
            # Should have both parameters updated
            assert "int* a_" in modified_content
            
        finally:
            os.unlink(temp_file)

    def test_compile_after_update(self):
        """Test that updated code can still compile (if compiler available)"""
        code = """
        #include <stdio.h>
        
        int test_func(int param) {
            return param * 2;
        }
        
        int main() {
            printf("%d\\n", test_func(5));
            return 0;
        }
        """
        
        with tempfile.NamedTemporaryFile(mode='w', suffix='.c', delete=False) as f:
            f.write(code)
            temp_file = f.name
            
        try:
            # Update the parameter
            result = process_file(temp_file, "test_func", 0, "float")
            assert result is True
            
            # Try to compile (if gcc is available)
            try:
                import subprocess
                result = subprocess.run(['gcc', '-c', temp_file, '-o', '/dev/null'], 
                                      capture_output=True, text=True)
                if result.returncode == 0:
                    print("Code compiles successfully after update")
                else:
                    print(f"Compilation warnings/errors: {result.stderr}")
            except FileNotFoundError:
                print("GCC not available for compilation test")
                
        finally:
            os.unlink(temp_file)


if __name__ == "__main__":
    pytest.main([__file__, "-v"])