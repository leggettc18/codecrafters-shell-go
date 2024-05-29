package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

    for {
	    // Print the prompt
        fmt.Fprint(os.Stdout, "$ ")
	    // Wait for user input
        text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        fmt.Fprintf(os.Stdout, "%s: command not found\n", strings.TrimSpace(text))
    }
}
