package main

import (
	"flag"
	"log"
	"os"

	"github.com/kalambet/iglo"
)

var outputFilePath string
var inputFilePath string

func main() {

	flag.StringVar(&outputFilePath, "out", "api.html", "Filename of the HTML output")
	flag.StringVar(&inputFilePath, "in", "api.apib", "Filename of the API Blueprint input")

	flag.Parse()

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	err = iglo.MarkdownToHTML(outputFile, inputFile)
	if err != nil {
		log.Fatal(err)
	}
}
