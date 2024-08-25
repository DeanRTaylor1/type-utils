HCLCONFIG {
    go_output_dir = "./generated"
    typescript_output_dir = "./types"
    javascript_output_dir = "./js"
    package_name = "models"
    file_name = "main"
}
# comment
User {
    optional authentication {
        id = "string"
        mfa_enabled = "boolean"
        optional repeated provider {
            name = "string"
            url = "string"
        }
    }
    optional UserProfile {
        email = "string"
        first_name = "string"
        last_name = "string"
        phone = "string"
        username = "string"
    }
}

Example {
    repeated name = "string"
    optional age = "int"
    optional repeated address = "string"
}
