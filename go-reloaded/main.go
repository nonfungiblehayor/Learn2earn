package main

import (
	"fmt"
	"os"
	"github.com/nonfungiblehayor/Learn2earn.git/utils"
)

func main() {
	content, err := os.ReadFile("test-files/sample.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(content)
	fmt.Println(utils.LoopText(text))
}