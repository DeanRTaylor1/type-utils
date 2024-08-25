package fetcher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/deanrtaylor1/type-utils/listener"
	"github.com/deanrtaylor1/type-utils/parser"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

func ParseSchemasFromGitRepo(repoURL, path string) ([]*listener.SchemaListener, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %v", err)
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	tmpDir := filepath.Join(homeDir, "tmp", "git-clone-"+uuid.String())
	err = os.MkdirAll(tmpDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	fullUrl := fmt.Sprintf("https://github.com/%s", repoURL)
	cmd := exec.Command("git", "clone", "--quiet", "--depth", "1", fullUrl, tmpDir)
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to clone repository: %v", err)
	}

	fullPath := filepath.Join(tmpDir, path)
	files, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	schemas := make([]*listener.SchemaListener, 0)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".hcl") {
			continue
		}
		content, err := os.ReadFile(filepath.Join(fullPath, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", file.Name(), err)
		}
		fileSchema, err := parseSchema(string(content))
		if err != nil {
			return nil, fmt.Errorf("failed to parse schema in %s: %v", file.Name(), err)
		}
		schemas = append(schemas, fileSchema)
	}

	return schemas, nil
}

func parseSchema(input string) (*listener.SchemaListener, error) {
	lexer := parser.NewHCLLikeDSLLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewHCLLikeDSLParser(stream)
	schemaListener := listener.NewSchemaListener()
	antlr.ParseTreeWalkerDefault.Walk(schemaListener, p.File())
	return schemaListener, nil
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
