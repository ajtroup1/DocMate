package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ajtroup1/DocMate/internal/types"
)

type Parser struct {
	comments        []types.CommentBlock
	Packages        []types.Package
	Errors          types.Error
	capitalizeItems bool
}

func New(comments []types.CommentBlock, capItems bool) *Parser {
	return &Parser{comments: comments, capitalizeItems: capItems}
}

func (p *Parser) ParseComments() {
	fmt.Print()
	// First, retrieve all package names to properly assign the root nodes for structured data
	pkgNames := p.retrievePackages()

	// Initialize all packages before evalutaion so you can assign nodes to a root
	for _, name := range pkgNames {
		p.createPackage(name)
	}

	// fmt.Printf("%v\n", p.Packages)

	for _, comment := range p.comments {
		p.parseIndividualCommentBlock(comment)
	}

	// fmt.Printf("%v\n", p.Packages)
}

func (p *Parser) retrievePackages() []string {
	var pkgNames []string
	uniquePkgs := make(map[string]bool)

	for _, comment := range p.comments {
		// Since names are being capitalized optionally, use non case-sensitive checking
		lowerPkgName := strings.ToLower(comment.Package)

		if !uniquePkgs[lowerPkgName] {
			// CapitalizeItems in the settings
			if p.capitalizeItems {
				comment.Package = strings.ToUpper(string(comment.Package[0])) + comment.Package[1:]
			}

			// Mark the lowercase package name as seen
			uniquePkgs[lowerPkgName] = true

			pkgNames = append(pkgNames, comment.Package)
		}
	}

	return pkgNames
}

func (p *Parser) parseIndividualCommentBlock(comment types.CommentBlock) {
	// pkgName := comment.Package
	// header := ""

	// Determine the header value
	if strings.HasPrefix(strings.TrimSpace(comment.Text[0]), "--") {
		// User is declaring a header value
		comment.Text[0] = strings.TrimSpace(strings.TrimPrefix(comment.Text[0], "--"))
		// header = comment.Text[0]
	}
	// fmt.Printf("%s\n", header)

	// Remove the header line before evaluation
	comment.Text = comment.Text[1:]

	// Deconstruct the tags
	for _, line := range comment.Text {
		// Determine the tag
		line = skipWhitespace(line)

		p.parseLine(line)
	}
}

func (p *Parser) createPackage(name string) {
	pkg := types.Package{
		Name: name,
	}

	p.Packages = append(p.Packages, pkg)
}

func (p *Parser) parseLine(line string) {
	// fmt.Printf("%s\n", line)
	if strings.HasPrefix(line, "@") {
		line = line[1:]
		// Extract the tag name
		tag, line := extractTagName(line)
		fmt.Printf("%s\n%s\n\n", tag, line)
	}
}

func extractTagName(line string) (string, string) {
	var tagName strings.Builder
	var i int

	for i = 0; i < len(line); i++ {
		if unicode.IsSpace(rune(line[i])) {
			break
		}
		tagName.WriteByte(line[i])
	}

	return tagName.String(), strings.TrimLeftFunc(line[i:], unicode.IsSpace)
}

// Skips any whitespace up until it reaches an actual char value
func skipWhitespace(line string) string {
	for i, ch := range line {
		if ch != ' ' && ch != '\t' && ch != '\n' && ch != '\r' {
			// Return the substring from the first non-whitespace character
			return line[i:]
		}
	}
	// If the entire string is whitespace, return an empty string
	return ""
}
