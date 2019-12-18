package editor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(displayText string) string {

	displayText = fmt.Sprintf("%s:", displayText)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(displayText)
	inputText, _ := reader.ReadString('\n')

	return inputText
}

func PromptBool(displayText string) bool {

	displayText = fmt.Sprintf("%s (y/n):", displayText)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(displayText)
	inputText, _ := reader.ReadString('\n')

	inputText = strings.ToLower(inputText)

	if inputText == "yes" || inputText == "y" {
		return true
	}

	return false
}
