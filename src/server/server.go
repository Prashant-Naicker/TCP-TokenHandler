package server

import (
    "fmt"
    "net"
    "net/http"
    "crypto/hmac"
    "encoding/binary"
    "crypto/sha256"
    "encoding/base64"
)

// Entry.
func Start() error {
    fmt.Println("Server is running..")
    l, err := net.Listen("tcp", ":8081")
    if err != nil { fmt.Printf("%v", err); return err }

    for {
        conn, err := l.Accept()
        if err != nil { fmt.Printf("%v", err); continue }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    secretKey := []byte("prashIsCool")

    fmt.Println("get size bytes")
    sizeBytes, err := awaitData(conn, 2)
    if (err != nil) {
        conn.Close()
        fmt.Println("Disconnected")
        return
    }
    fmt.Println("convert to uint16")
    size := binary.BigEndian.Uint16(sizeBytes)
    fmt.Println(size)
    fmt.Println("await message data")
    data, err := awaitData(conn, int(size))
    if (err != nil) {
        conn.Close()
        fmt.Println("Disconnected")
        return
    }
    fmt.Println("write message back to user")
    h := hmac.New(sha256.New, secretKey)
    h.Write(data)
    fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))
    _, err = conn.Write([]byte(base64.StdEncoding.EncodeToString(h.Sum(nil))))
    if (err != nil) {
        conn.Close()
        fmt.Println("Disconnected")
        return
    }
    fmt.Println("finish")

    return
}

func awaitData(conn net.Conn, totalSize int) ([]byte, error) {
    fmt.Println("awaitData")
    buffer := make([]byte, totalSize)
    readSize := 0

    for (readSize < totalSize) {
        length, err := conn.Read(buffer[readSize:]) //Read method stores bytes being read into the buffer and returns the length of bytes read.
        if err != nil { return nil, err }

        readSize += length
    }
    fmt.Println("ReadSize Value: ", readSize)
    fmt.Println("Data buffer: ", buffer)
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
