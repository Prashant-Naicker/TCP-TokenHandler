package server

import (
    "fmt"
    "net"
    "net/http"
)

// Entry.
func Start() error {
    l, err := net.Listen("tcp", ":8081")
    if err != nil { fmt.Printf("%v", err); return err }

    for {
        conn, err := l.Accept()
        if err != nil { fmt.Printf("%v", err); continue }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Printf("works")
}

// Headers.
func SetGeneralHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
}
