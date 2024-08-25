HCLCONFIG {
    go_output_dir = "./generated"
    package_name = "models"
    file_name = "main"
}
# comment
User {
    authentication {
        id = "string"
        mfa_enabled = "boolean"
        repeated provider {
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

