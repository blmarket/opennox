from tree_sitter import Node

def find_function_definition(node: Node, func_name: str, source_code: str) -> Node:
    """Find function definition by name"""
    if node.type == "function_definition":
        declarator = None
        for child in node.children:
            if child.type == "function_declarator":
                declarator = child
                break
            elif child.type == "pointer_declarator":
                # Handle pointer return types: look for function_declarator inside pointer_declarator
                for ptr_child in child.children:
                    if ptr_child.type == "function_declarator":
                        declarator = ptr_child
                        break
                if declarator:
                    break
        
        if declarator:
            for child in declarator.children:
                if child.type == "identifier" and source_code[child.start_byte:child.end_byte] == func_name:
                    return node
    
    for child in node.children:
        result = find_function_definition(child, func_name, source_code)
        if result:
            return result
    
    return None

def find_struct_definition(node: Node, struct_name: str, source_code: str) -> Node:
    """Find struct definition by name"""
    if node.type == "struct_specifier":
        for child in node.children:
            if child.type == "type_identifier" and source_code[child.start_byte:child.end_byte] == struct_name:
                return node
    
    for child in node.children:
        result = find_struct_definition(child, struct_name, source_code)
        if result:
            return result
    return None

def find_variable_references(node: Node, var_name: str, source_code: str) -> list[Node]:
    """Find all nodes where a variable is referenced"""
    references = []
    
    def search_references(node):
        if node.type == "identifier" and source_code[node.start_byte:node.end_byte] == var_name:
            references.append(node)
        
        for child in node.children:
            search_references(child)
    
    search_references(node)
    return references