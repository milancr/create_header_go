package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"golang.design/x/clipboard"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter header text: ")
	headerText, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	headerText = strings.ToUpper(strings.TrimSuffix(headerText, "\n"))
	separator := "    ///////////////////////////////////////////////////////////////"
	output := fmt.Sprintf("%s\n%s%s%s\n%s", separator, "    ", strings.Repeat(" ", (64-len(headerText))/2), headerText, separator)
	fmt.Println(output)

	err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(output))
	fmt.Println("Copied to clipboard!")
}
