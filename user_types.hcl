HCLCONFIG {
    output_dir = "./generated"
    package_name = "users"
    file_name = "user_types"
}

# Custom type using imported type
UserProfile {
    user = User  # Assuming User is defined in user_types.hcl
    preferences = Preferences  # Assuming Preferences is defined in user_types.hcl
    metadata = Metadata  # Assuming Metadata is defined in common_types.hcl
}

Address {
    street = "string"
    city = "string"
    country = "string"
}

# Complex type with various field types
ComplexType {
    id = "string"
    name = "string"
    age = "number"
    is_verified = "boolean"
    optional repeated tags = "string"
    created_at = "time.Duration"  # Assuming Timestamp is defined in common_types.hcl
}

