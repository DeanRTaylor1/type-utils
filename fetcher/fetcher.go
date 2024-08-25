package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/deanrtaylor1/type-utils/listener"
	"github.com/deanrtaylor1/type-utils/parser"
)

func ParseSchemasFromGitRepo(repo, path string) ([]*listener.SchemaListener, error) {
	// Assuming the repo is in the format "username/repo"
	baseURL := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path)

	// Fetch the list of files in the specified directory
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schema files: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch schema files, status: %d", resp.StatusCode)
	}

	var files []struct {
		Name string `json:"name"`
		URL  string `json:"download_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	schemas := make([]*listener.SchemaListener, 0)

	for _, file := range files {
		if !strings.HasSuffix(file.Name, ".hcl") {
			continue
		}

		// Fetch the content of each .hcl file
		content, err := fetchFileContent(file.URL)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch content of %s: %v", file.Name, err)
		}

		// Parse the schema
		fileSchemas, err := parseSchema(content)
		if err != nil {
			return nil, fmt.Errorf("failed to parse schema in %s: %v", file.Name, err)
		}

		schemas = append(schemas, fileSchemas)

	}

	return schemas, nil
}

func fetchFileContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch file, status: %d", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func parseSchema(input string) (*listener.SchemaListener, error) {
	lexer := parser.NewHCLLikeDSLLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewHCLLikeDSLParser(stream)
	schemaListener := listener.NewSchemaListener()
	antlr.ParseTreeWalkerDefault.Walk(schemaListener, p.File())

	return schemaListener, nil
}
