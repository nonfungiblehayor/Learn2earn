package main

import (
	"fmt"
	"os"
	"github.com/nonfungiblehayor/Learn2earn.git/utils"
)

func main() {
	inputFile := os.Args[1]
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(content)
	contentResult := utils.LoopText(text)
	newContent := []byte(contentResult)
	errFile := os.WriteFile(inputFile, newContent, 0644)
	if err != nil {
		panic(errFile)
	}
	fmt.Println("File has been auto-corrected successfully!!!")
}