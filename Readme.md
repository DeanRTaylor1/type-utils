# Type Utils

Write types once, generate anywhere.

This application is a schema generator that uses a language based on the HashiCorp Configuration Language (HCL) to define types and configurations, which can then be used to generate schemas in various programming languages. The tool allows for flexible configuration, including specifying different output directories for various languages, enabling multi-repository support, and pulling schema definitions from remote GitHub repositories.

## Features

- **Multi-language Support**: Generate schemas in multiple languages including Go, TypeScript, JavaScript, Java, and Swift.
- **Configurable Outputs**: Specify output directories for different languages.
- **GitHub Integration**: Pull schema definitions directly from GitHub repositories.
- **Constructor & Getter Generation**: Automatically generate constructors and getter methods for defined types.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Defining a Schema File](#defining-a-schema-file)
  - [Repository Configuration](#repository-configuration)
  - [Supported Types](#supported-types)
  - [Supported Languages](#supported-languages)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install the HCL Schema Generator, clone the repository and build the application:

```bash
git clone https://github.com/your-repo/schema-generator.git
cd schema-generator
make build
```

## Usage

### Defining a Schema File

Each schema file begins with a `Type_Config` block, which specifies where the generated files should be saved for different programming languages. This is followed by the type definitions. Here’s an example:

```hcl
Type_Config {
    go_output_dir = "./generated"
    typescript_output_dir = "./types"
    javascript_output_dir = "./js"
    package_name = "models"
    file_name = "main"
}

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
```

#### Type Configuration

The `Type_Config` block is mandatory and allows you to define output directories for different languages, package names, and file names.

- `go_output_dir`: Directory for Go files.
- `typescript_output_dir`: Directory for TypeScript files.
- `javascript_output_dir`: Directory for JavaScript files.
- `package_name`: The package or namespace name for the generated files.
- `file_name`: The base name for the generated files.

### Repository Configuration

The repository-level configuration is done via a YAML file (e.g., `config.yaml`). This file specifies global settings such as the target language, schema directory, and optional GitHub repository for schema sources.

```yaml
version: "v1"
language: "go"
generate_constructors: true
generate_getters: true
schemas_dir_name: "schemas"
git_repo:
  url: "your-github-repo/schema-repo"
  path: "schemas"
```

#### YAML Config Options

- `version`: The version of the configuration.
- `language`: The language for which the schemas should be generated.
- `generate_constructors`: Whether to generate constructors.
- `generate_getters`: Whether to generate getter methods.
- `schemas_dir_name`: (Optional) The directory where the schemas are stored.
- `git_repo`: (Optional) Define a GitHub repository and path to pull schema definitions.

> [!note]-
> Users should define at least a git_repo option or a schemas_dir_name

### Supported Types

The following types are supported in HCL schema definitions:

- **Basic Types**:

  - `string`
  - `int`
  - `int32`
  - `int64`
  - `float`
  - `double`
  - `boolean`

- **Date and Time Types**:
  - `date`
  - `time`
  - `datetime`
  - `timestamp`

### Supported Languages

The tool currently supports the following languages for schema generation:

- **Go**
- **TypeScript**
- **JavaScript**
- **Java**
- **Swift**

The output directory for each language is defined in the `Type_Config` block of your schema files.

## Examples

### Example Schema File

- Please see below for example of what this code outputs

```hcl
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

```

### Example Repository Configuration

```yaml
version: "v1"
language: "go"
generate_constructors: true
generate_getters: true
schemas_dir_name: "schemas"
```

## Example outputs

### Go output

```go
package users

type User struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Username string `json:"username,omitempty"`
	Authentication Authentication `json:"authentication,omitempty"`
	Name []string `json:"name"`
	Age int32 `json:"age,omitempty"`
	Address []string `json:"address,omitempty"`
}

func NewUser(phone string, username string, authentication Authentication, name []string, age int32, address []string, email string) *User {
	return &User{
		Age: age,
		Address: address,
		Email: email,
		Phone: phone,
		Username: username,
		Authentication: authentication,
		Name: name,
	}
}

type Authentication struct {
	Provider []Provider `json:"provider,omitempty"`
	Id string `json:"id"`
	Mfa_enabled bool `json:"mfa_enabled"`
}

func NewAuthentication(id string, mfa_enabled bool, provider []Provider) *Authentication {
	return &Authentication{
		Id: id,
		Mfa_enabled: mfa_enabled,
		Provider: provider,
	}
}

type Provider struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func NewProvider(name string, url string) *Provider {
	return &Provider{
		Name: name,
		Url: url,
	}
}




// types generated by type-utils
```

### Javascript output

```javascript
// Generated types for users

/**
 * @typedef {Object} Authentication
 * @property {string} id
 * @property {boolean} mfa_enabled
 * @property {Provider[]} provider [optional]
 */
class Authentication {
  constructor(data) {
    this.id = data.id;
    this.mfa_enabled = data.mfa_enabled;
    this.provider = data.provider;
  }
}

/**
 * @typedef {Object} Provider
 * @property {string} name
 * @property {string} url
 */
class Provider {
  constructor(data) {
    this.name = data.name;
    this.url = data.url;
  }
}

/**
 * @typedef {Object} User
 * @property {string} phone [optional]
 * @property {string} username [optional]
 * @property {Authentication} authentication [optional]
 * @property {string[]} name
 * @property {number} age [optional]
 * @property {string[]} address [optional]
 * @property {string} email [optional]
 */
class User {
  constructor(data) {
    this.authentication = data.authentication;
    this.name = data.name;
    this.age = data.age;
    this.address = data.address;
    this.email = data.email;
    this.phone = data.phone;
    this.username = data.username;
  }
}

// Export all types
export { Authentication, Provider, User };

// Types generated by type-utils
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
