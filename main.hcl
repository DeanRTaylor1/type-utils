DCLCONFIG {
    output_dir = "./generated"
    package_name = "models"
}


# Basic types
BasicTypes {
    string_field = "string"
    number_field = "number"
    boolean_field = "boolean"
    repeated string_array = "string"
    repeated number_array = "number"
}

# Custom type using imported type
UserProfile {
    user = User  # Assuming User is defined in user_types.hcl
    preferences = Preferences  # Assuming Preferences is defined in user_types.hcl
    metadata = Metadata  # Assuming Metadata is defined in common_types.hcl
}

# Nested structure
NestedType {
    name = "string"
    repeated addresses = Address
}

Address {
    street = "string"
    city = "string"
    country = "string"
}

# Complex type with various field types
ComplexType {
    id = UUID  # Assuming UUID is defined in common_types.hcl
    name = "string"
    age = "number"
    is_verified = "boolean"
    repeated tags = "string"
    created_at = Timestamp  # Assuming Timestamp is defined in common_types.hcl
    union_type = "string | number"  # Assuming we support union types
}

NestedType {
    name = "string"
    repeated addresses {
        street = "string"
        city = "string"
        country = "string"
    }
}
