package lexer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ajtroup1/DocMate/internal/types"
)

type Lexer struct {
	src          string
	position     int
	readPosition int
	ch           byte
	includeTests bool
	projectPath  string
}

func New(include bool, path string) *Lexer {
	return &Lexer{includeTests: include, projectPath: path}
}

func (e *Lexer) ExtractComments() ([]types.CommentBlock, error) {
	fmt.Print()
	var comments []types.CommentBlock

	err := filepath.WalkDir(e.projectPath, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// if !entry.IsDir() && entry.Name() == "go.mod" {
		// 	projectName, err = e.extractGoMod(path)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			if e.includeTests || (!e.includeTests && !strings.HasSuffix(entry.Name(), "_test.go")) {
				fileComments, err := e.extractCommentsFromFile(path)
				if err != nil {
					return err
				}
				if len(fileComments) > 0 {
					fmt.Printf("%d comments found in `%s`\n", len(fileComments), path)
					comments = append(comments, fileComments...)
					e.resetState()
				}
			}
		}
		return nil

	})

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (e *Lexer) extractGoMod(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "module ") {
			modulePath := strings.TrimSpace(strings.TrimPrefix(line, "module "))

			parts := strings.Split(modulePath, "/")
			if len(parts) > 0 {
				return parts[len(parts)-1], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to scan go.mod file: %v", err)
	}

	return "", fmt.Errorf("module line not found in go.mod file")
}

func (e *Lexer) extractCommentsFromFile(filePath string) ([]types.CommentBlock, error) {
	var comments []types.CommentBlock
	fmt.Printf("\033[32mReading comments from %s\n\033[0m", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := e.readFileContent(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", filePath, err)
	}
	e.src = content

	pkgName, err := e.extractPkgName()
	if err != nil {
		return comments, fmt.Errorf("error receiving package name")
	}

	e.readPosition = 0
	e.readChar()

	for !e.isAtEnd() {
		if e.isGoDocComment() {
			comment, err := e.extractBlockComment(filePath, pkgName)
			if err != nil {
				return comments, err
			}
			if !isEmptyComment(comment) {
				comments = append(comments, comment)
			}
		} else {
			e.readChar()
		}
	}

	return comments, nil
}

func (e *Lexer) readFileContent(file *os.File) (string, error) {
	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return sb.String(), nil
}

func (e *Lexer) extractPkgName() (string, error) {
	e.readChar()
	e.skipWhitespace()

	var pkgName strings.Builder

	for e.ch != 0 {
		if e.ch == '/' && e.peekChar(0) == '/' {
			e.skipLineComment()
			continue
		}
		if e.ch == '/' && e.peekChar(0) == '*' {
			e.skipBlockComment()
			continue
		}

		// Look for the 'package' keyword
		if e.ch == 'p' && strings.HasPrefix(e.src[e.position:], "package ") {
			// Move the position past 'package '
			e.position += len("package ")
			e.readPosition = e.position
			e.readChar()

			e.skipWhitespace()

			// Extract the entire package name
			for e.ch != ' ' && e.ch != '\n' && e.ch != '\r' && e.ch != 0 {
				pkgName.WriteByte(e.ch)
				e.readChar()
			}
			break
		}

		e.readChar()
	}

	packageName := pkgName.String()

	// Default to "main" if no valid package name is found
	if strings.TrimSpace(packageName) == "" {
		packageName = "main"
	}

	return packageName, nil
}


func (e *Lexer) extractBlockComment(filePath, pkgName string) (types.CommentBlock, error) {
	var lines []string
	e.advanceBy(4)

	for !e.isAtEnd() {
		e.skipWhitespace()

		if e.ch == '*' && e.peekChar(0) == '/' {
			e.advanceBy(2)
			break
		}

		var sb strings.Builder
		for e.ch != '\n' && e.ch != 0 {
			sb.WriteByte(e.ch)
			e.readChar()
		}

		line := strings.TrimSpace(sb.String())
		if len(line) > 0 {
			lines = append(lines, line)
		}

		e.readChar()
	}

	if len(lines) == 0 {
		return types.CommentBlock{}, fmt.Errorf("empty comment found in %s", filePath)
	}

	return types.CommentBlock{
		Filepath: filePath,
		Package:  pkgName,
		Text:     lines,
	}, nil
}

func (e *Lexer) resetState() {
	e.position = 0
	e.readPosition = 0
}

func (e *Lexer) peekChar(ahead int) byte {
	if e.readPosition+ahead >= len(e.src) {
		return 0
	}
	return e.src[e.readPosition]
}

func (e *Lexer) readChar() {
	if e.readPosition >= len(e.src) {
		e.ch = 0
	} else {
		e.ch = e.src[e.readPosition]
	}
	e.position = e.readPosition
	e.readPosition += 1
}

func (e *Lexer) advanceBy(n int) {
	for i := 0; i < n; i++ {
		e.readChar()
	}
}

func (e *Lexer) skipWhitespace() {
	for e.ch == ' ' || e.ch == '\t' || e.ch == '\n' || e.ch == '\r' {
		e.readChar()
	}
}

func (e *Lexer) isAtEnd() bool {
	return e.readPosition >= len(e.src)
}

func (e *Lexer) skipLineComment() {
	for e.ch != '\n' && e.ch != 0 {
		e.readChar()
	}
}

func (e *Lexer) skipBlockComment() {
	e.readChar()
	e.readChar()
	for !(e.ch == '*' && e.peekChar(0) == '/') && e.ch != 0 {
		e.readChar()
	}
	e.readChar()
	e.readChar()
}

func (e *Lexer) isGoDocComment() bool {
	return e.ch == '/' && e.peekChar(0) == '*' && e.peekChar(1) == '*' && e.peekChar(2) == '*'
}

func isEmptyComment(comment types.CommentBlock) bool {
	return len(comment.Text) == 0
}
