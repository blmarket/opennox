from tree_sitter import Node

def dbg(node: Node, source_code: str) -> str:
    """Extract source code text for a tree-sitter node"""
    return source_code[node.start_byte:node.end_byte]

def print_tree(node: Node, source_code: str, depth=0, max_depth=3):
    """Print tree structure for debugging"""
    if depth > max_depth:
        return
    
    indent = "  " * depth
    node_text = source_code[node.start_byte:node.end_byte]
    # Truncate long text and replace newlines
    if len(node_text) > 50:
        node_text = node_text[:50] + "..."
    node_text = node_text.replace('\n', '\\n').replace('\t', '\\t')
    
    print(f"{indent}{node.type}: '{node_text}'")
    
    for child in node.children:
        print_tree(child, source_code, depth + 1, max_depth)