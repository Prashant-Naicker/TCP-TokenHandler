package server

import (
    "fmt"
    "net"
    "net/http"
)

// Entry.
func Start() error {
    fmt.Printf("Server is running..")
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
    _, err := conn.Write([]byte("message"))
    if err != nil { fmt.Printf("%v", err); return }
    return
}

func awaitData(conn net.Conn, totalSize int) ([]byte, error) {
    buffer := make([]byte, totalSize)
    readSize := 0

    for (readSize < totalSize) {
        length, err := conn.Read(buffer[readSize:]) //Read method stores bytes being read into the buffer and returns the length of bytes read.
        if err != nil { return nil, err }

        readSize += length
    }
    fmt.Println(readSize)
    return buffer, nil
}

// Headers.
func SetGeneralHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
}
