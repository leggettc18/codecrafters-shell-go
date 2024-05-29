package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func findExecutable(name string, paths []string) (string, bool) {
    for _, path := range paths {
        fullpath := filepath.Join(path, name)
        if _, err := os.Stat(fullpath); err == nil {
            return fullpath, true
        }
    }
    return "", false
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

    for {
	    // Print the prompt
        fmt.Fprint(os.Stdout, "$ ")
	    // Wait for user input
        text, err := bufio.NewReader(os.Stdin).ReadString('\n')
        if err != nil {
            fmt.Fprintf(os.Stderr, "error while processing stdin: %s\n", err)
            os.Exit(1)
        }
        text = strings.TrimSpace(text)
        commands := strings.Split(text, " ")
        switch commands[0] {
        case "exit":
            exitCode, err := strconv.Atoi(commands[1])
            if err != nil {
                fmt.Fprintf(os.Stderr, "error while parsing exit code: %s\n", err)
                os.Exit(1)
            }
            os.Exit(exitCode)
        case "echo":
            for index, command := range commands[1:] {
                fmt.Fprint(os.Stdout, command)
                if (index < len(commands) - 2) {
                    fmt.Fprint(os.Stdout, " ")
                } else {
                    fmt.Fprint(os.Stdout, "\n")
                }
            }
        case "type":
            switch commands[1] {
            case "exit", "echo", "type":
                fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", commands[1])
            default:
                path, result := findExecutable(commands[1], strings.Split(os.Getenv("PATH"), ":"))
                if result {
                    fmt.Fprintf(os.Stdout, "%s is %s\n", commands[1], path)
                } else {
                    fmt.Fprintf(os.Stdout, "%s not found\n", commands[1])
                }
        }
        default:
            fmt.Fprintf(os.Stdout, "%s: command not found\n", text)
        }
    }
}
