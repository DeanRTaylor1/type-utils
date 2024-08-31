Type_Config {
    go_output_dir = "./users"
    typescript_output_dir = "./types/users"
    javascript_output_dir = "./src/users"
    package_name = "users"
    file_name = "users"
}

User {
    repeated name = "string"
    optional age = "int"
    optional repeated address = "string"
    optional email = "string"
    optional phone = "string"
    optional username = "string"
    optional authentication {
        id = "string"
        mfa_enabled = "boolean"
        optional repeated provider {
            name = "string"
            url = "string"
        }
    }
}
