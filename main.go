package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/erickmx/compiller/lexical"
	"github.com/erickmx/compiller/utils"
)

func processFile(fileData string) {
	withOutSpaces := strings.Split(fileData, " ")
	var trimmed string
	for _, word := range withOutSpaces {
		trimmed = strings.TrimSpace(word)
		if len(trimmed) != 0 {
			fmt.Println(trimmed, ": ", lexical.Classify(trimmed))
		}
	}
}

func processConsole(stdin io.Reader) {
	reader := bufio.NewReader(stdin)
	var expression string
	var splitted []string
	for {
		fmt.Println("Ingrese una expression: ")
		fmt.Println("Presione enter para salir")
		expression, _ = reader.ReadString('\n')
		expression = strings.Replace(expression, "\n", "", -1)
		if len(expression) == 0 {
			break
		}
		splitted = strings.Split(expression, " ")
		for _, word := range splitted {
			fmt.Println(word, ": ", lexical.Classify(word))
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		fileData, err := utils.OpenFile(args[0])
		if err != nil {
			fmt.Printf("An error occurred while opening the file. %v", err)
		}
		processFile(fileData)
	} else {
		processConsole(os.Stdin)
	}
}
