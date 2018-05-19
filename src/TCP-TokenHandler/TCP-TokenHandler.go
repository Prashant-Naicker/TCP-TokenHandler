package main

import (
    "fmt"
    "net/http"
    "golang.org/x/crypto/acme/autocert"
)

// Entry.
func main() {
    errc := start()
    err := errc
	fmt.Printf("Fatal Error: %v\n", err)
}

func start() chan error {
	errc := make(chan error)
	httpsMux := http.NewServeMux()
	httpsMux.HandleFunc("/", requestHandler)

    acmeMgr := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("127.0.0.1"),
		Cache:      autocert.DirCache("autocert-cache"),
	}
	go func() {
		errc <- http.ListenAndServe(":http", acmeMgr.HTTPHandler(nil))
	}()

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP521, tls.CurveP384, tls.CurveP256},
		GetCertificate:           acmeMgr.GetCertificate,
		PreferServerCipherSuites: true,
	}

	httpsServer := &http.Server{
		Addr:         ":443",
		Handler:      httpsMux,
		TLSConfig:    tlsConfig,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     log.New(ioutil.Discard, "", 0),
		TLSNextProto: nil,
	}

	go func() {
		errc <- httpsServer.ListenAndServe("", "")
	}()

	return errc
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
    SetGeneralHeaders(w)
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
