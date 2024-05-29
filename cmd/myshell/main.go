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
        text, err := bufio.NewReader(os.Stdin).ReadString('\n')
        if err != nil {
            fmt.Fprintf(os.Stderr, "error while processing stdin: %s", err)
            os.Exit(1)
        }
        fmt.Fprintf(os.Stdout, "%s: command not found\n", strings.TrimSpace(text))
    }
}
