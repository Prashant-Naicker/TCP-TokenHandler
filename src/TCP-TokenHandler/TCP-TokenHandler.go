package main

import (
    "fmt"

    "server"
)

// Entry.
func main() {
    errc := server.Start()
    err := errc
	fmt.Printf("Fatal Error: %v\n", err)
}
