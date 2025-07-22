from tree_sitter import Node

def extract_type_and_name(node: Node, source_code: str) -> tuple[str, str]:
    """Extract type and variable name from a declaration node"""
    decl_type = None
    var_name = None
    is_pointer = False
    
    for child in node.children:
        if child.type in ["primitive_type", "type_identifier", "struct_specifier"]:
            decl_type = source_code[child.start_byte:child.end_byte]
        elif child.type == "identifier":
            var_name = source_code[child.start_byte:child.end_byte]
        elif child.type == "pointer_declarator":
            is_pointer = True
            for ptr_child in child.children:
                if ptr_child.type == "identifier":
                    var_name = source_code[ptr_child.start_byte:ptr_child.end_byte]
        elif child.type == "init_declarator":
            # Handle local variable declarations
            for init_child in child.children:
                if init_child.type == "pointer_declarator":
                    is_pointer = True
                    for ptr_child in init_child.children:
                        if ptr_child.type == "identifier":
                            var_name = source_code[ptr_child.start_byte:ptr_child.end_byte]
                elif init_child.type == "identifier":
                    var_name = source_code[init_child.start_byte:init_child.end_byte]
    
    final_type = decl_type + "*" if is_pointer and decl_type else decl_type
    return final_type, var_name

def find_variable_type(func_node: Node, var_name: str, source_code: str) -> str:
    """Find variable/parameter types in function"""
    # Check function parameters
    for child in func_node.children:
        if child.type == "pointer_declarator":
            child = child.children[1]
        if child.type == "function_declarator":
            for param_child in child.children:
                if param_child.type == "parameter_list":
                    for param in param_child.children:
                        if param.type == "parameter_declaration":
                            param_type, param_name = extract_type_and_name(param, source_code)
                            if param_name == var_name:
                                return param_type
    
    # Check local variable declarations
    def search_declarations(node):
        if node.type == "declaration":
            # Get the base type first
            base_type = None
            for child in node.children:
                if child.type in ["primitive_type", "type_identifier", "struct_specifier"]:
                    base_type = source_code[child.start_byte:child.end_byte]
                    break
            
            # Look for the variable in init_declarators
            for child in node.children:
                if child.type == "init_declarator":
                    # Check for pointer declarator
                    for init_child in child.children:
                        if init_child.type == "pointer_declarator":
                            for ptr_child in init_child.children:
                                if ptr_child.type == "identifier" and source_code[ptr_child.start_byte:ptr_child.end_byte] == var_name:
                                    return base_type + "*" if base_type else None
                        elif init_child.type == "identifier" and source_code[init_child.start_byte:init_child.end_byte] == var_name:
                            return base_type
        
        for child in node.children:
            result = search_declarations(child)
            if result:
                return result
        return None
    
    return search_declarations(func_node)

def infer_from_func(node: Node, source_code: str) -> tuple[str, str]:
    """
    Check the first statement of the function compound statement. If it's declaration with init_declarator then infer SRC, TGT from the assignment and return as a tuple.
    """
    # Find the compound statement (function body)
    compound_stmt = None
    for child in node.children:
        if child.type == "compound_statement":
            compound_stmt = child
            break
    
    if not compound_stmt:
        return None, None
    
    # Get the first statement in the compound statement
    first_stmt = None
    for child in compound_stmt.children:
        if child.type == "declaration":
            first_stmt = child
            break
    
    if not first_stmt:
        return None, None
    
    # Look for init_declarator in the declaration
    for child in first_stmt.children:
        if child.type == "init_declarator":
            # The structure should be: init_declarator -> identifier, "=", identifier
            identifiers = []
            for init_child in child.children:
                if init_child.type == "identifier":
                    identifiers.append(source_code[init_child.start_byte:init_child.end_byte])
                elif init_child.type == "pointer_declarator":
                    identifiers.append(source_code[init_child.children[1].start_byte:init_child.children[1].end_byte])
            
            # Should have exactly 2 identifiers: [declared_var, assigned_value]
            if len(identifiers) == 2:
                # Return (SRC, TGT) where SRC is the new variable and TGT is what it's assigned from
                return identifiers[0], identifiers[1]
    
    return None, None