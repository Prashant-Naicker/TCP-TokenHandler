package main

import (
    "fmt"

    "server"
)

// Entry.
func main() {
    errc := server.Start()
	fmt.Printf("Fatal Error: %v\n", errc)
}
