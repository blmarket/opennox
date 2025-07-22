from tree_sitter import Node

def calculate_field_offsets(struct_node: Node, source_code: str) -> list[dict]:
    """Calculate byte offsets for struct fields (assuming 32-bit/4-byte alignment)"""
    fields = []
    offset = 0
    
    # Find the field declaration list
    field_list = None
    for child in struct_node.children:
        if child.type == "field_declaration_list":
            field_list = child
            break
    
    if not field_list:
        return fields
    
    for field_decl in field_list.children:
        if field_decl.type == "field_declaration":
            # Extract type and field names
            field_type = None
            field_names = []
            size = None  # Initialize size as None
            
            for child in field_decl.children:
                if child.type in ["primitive_type", "type_identifier", "struct_specifier"]:
                    field_type = source_code[child.start_byte:child.end_byte]
                elif child.type == "field_identifier":
                    field_names.append(source_code[child.start_byte:child.end_byte])
                elif child.type == "field_declarator":
                    # Handle pointer or array declarators
                    for decl_child in child.children:
                        if decl_child.type == "field_identifier":
                            field_names.append(source_code[decl_child.start_byte:decl_child.end_byte])
                        elif decl_child.type == "pointer_declarator":
                            for ptr_child in decl_child.children:
                                if ptr_child.type == "field_identifier":
                                    field_names.append(source_code[ptr_child.start_byte:ptr_child.end_byte])
                                    field_type += "*"
                elif child.type == "pointer_declarator":
                    # Handle direct pointer declarators (like "struct struct576* next;")
                    if field_type:
                        field_type += "*"
                    for ptr_child in child.children:
                        if ptr_child.type == "field_identifier":
                            field_names.append(source_code[ptr_child.start_byte:ptr_child.end_byte])
                        elif ptr_child.type == "identifier":
                            field_names.append(source_code[ptr_child.start_byte:ptr_child.end_byte])
                        elif ptr_child.type == "array_declarator":
                            # Handle pointer to array like "void* field_10[32]"
                            array_size = 1
                            field_name = None
                            for array_child in ptr_child.children:
                                if array_child.type == "field_identifier":
                                    field_name = source_code[array_child.start_byte:array_child.end_byte]
                                elif array_child.type == "number_literal":
                                    array_size = int(source_code[array_child.start_byte:array_child.end_byte])
                            
                            field_names.append(field_name)
                            base_size = 4  # default pointer size
                            if field_type and "*" in field_type:
                                base_size = 4  # pointer size on 32-bit
                            size = base_size * array_size
                elif child.type == "array_declarator":
                    # Handle array declarators like field_32[32] as direct children of field_declaration
                    array_size = 1
                    field_name = None
                    for array_child in child.children:
                        if array_child.type == "field_identifier":
                            field_name = source_code[array_child.start_byte:array_child.end_byte]
                        elif array_child.type == "number_literal":
                            array_size = int(source_code[array_child.start_byte:array_child.end_byte])
                    
                    field_names.append(field_name)
                    # For arrays, calculate element size and total size
                    base_size = 4  # default
                    if field_type == "uint16_t":
                        base_size = 2
                    elif field_type == "uint8_t" or field_type == "char":
                        base_size = 1
                    elif field_type == "uint32_t" or field_type == "int":
                        base_size = 4
                    elif field_type == "uint64_t":
                        base_size = 8
                    elif "*" in field_type:
                        base_size = 4  # pointer size on 32-bit
                    size = base_size * array_size
            
            # Calculate size based on type (assuming 32-bit architecture)
            # Only calculate size if not already set (e.g., for arrays)
            if field_type and size is None:
                if "*" in field_type or field_type in ["void*", "char*", "int*"]:
                    size = 4  # pointer size on 32-bit
                elif field_type in ["char", "unsigned char", "int8_t", "uint8_t"]:
                    size = 1
                elif field_type in ["short", "unsigned short", "int16_t", "uint16_t"]:
                    size = 2
                elif field_type in ["int", "unsigned int", "long", "unsigned long", "int32_t", "uint32_t", "float"]:
                    size = 4
                elif field_type in ["long long", "unsigned long long", "int64_t", "uint64_t", "double"]:
                    size = 8
                elif field_type == "timer":
                    size = 32  # timer struct is 32 bytes
                elif field_type == "timerGroup":
                    size = 96
                elif field_type == "nox_list_item_t":
                    size = 12  # nox_list_item_t struct is 12 bytes
                else:
                    size = 4  # default assumption
            
            # Process fields if we have both type and size (either calculated or pre-set for arrays)
            if field_type and size is not None:
                # Add padding for alignment based on 32-bit architecture rules
                # For most types, align to 4-byte boundaries (GOARCH=386)
                alignment = 4  # default 4-byte alignment for 32-bit
                if field_type in ["char", "unsigned char", "int8_t", "uint8_t"]:
                    alignment = 1
                elif field_type in ["short", "unsigned short", "int16_t", "uint16_t"]:
                    alignment = 2
                elif field_type in ["long long", "unsigned long long", "int64_t", "uint64_t", "double"]:
                    alignment = 4  # Even 64-bit types use 4-byte alignment on 32-bit systems
                
                padding = offset % alignment
                if padding != 0:
                    offset += alignment - padding
                
                for field_name in field_names:
                    fields.append({
                        'name': field_name,
                        'type': field_type,
                        'offset': offset,
                        'size': size
                    })
                    offset += size
    
    return fields

def field_by_offset(node: Node, source_code: str) -> dict[int, str]:
    """Convert struct node to dictionary mapping byte offsets to field names"""
    if not node:
        return {}
    
    fields = calculate_field_offsets(node, source_code)
    offset_map = {}
    
    for field in fields:
        offset_map[field['offset']] = field['name']
    
    return offset_map