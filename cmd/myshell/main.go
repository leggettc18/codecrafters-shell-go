package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
        text = strings.TrimSpace(text)
        commands := strings.Split(text, " ")
        switch commands[0] {
        case "exit":
            exitCode, err := strconv.Atoi(commands[1])
            if err != nil {
                fmt.Fprintf(os.Stderr, "error while parsing exit code: %s", err)
            }
            os.Exit(exitCode)
        default:
            fmt.Fprintf(os.Stdout, "%s: command not found\n", text)
        }
    }
}
