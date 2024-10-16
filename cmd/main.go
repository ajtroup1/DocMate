package main

import (
	"fmt"
	"log"

	"github.com/ajtroup1/DocMate/internal/lexer"
	"github.com/ajtroup1/DocMate/internal/parser"
	"github.com/ajtroup1/DocMate/internal/utils"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Clear  = "\033[0m"
)

func main() {
	fmt.Print()
	settings, err := utils.GetOrRetrieveSettings()
	if err != nil {
		log.Fatalf(Red+"Error retrieving / creating settings: %v\n"+Clear, err)
	}

	lexer := lexer.New(settings.IncludeTests, settings.ProjectPath)
	comments, err := lexer.ExtractComments()

	parser := parser.New(comments, settings.CapitalizeItems)
	parser.ParseComments()

}
