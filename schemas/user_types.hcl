HCLCONFIG {
    go_output_dir = "./generated"
    typescript_output_dir = "./types"
    javascript_output_dir = "./js"
    package_name = "users"
    file_name = "user_types"
}

# Custom type using imported type
UserProfile {
    user = "models.User"  # Assuming User is defined in user_types.hcl
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

