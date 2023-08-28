/*
Serve is a very simple static file server in go
Usage:

	-p="8100": port to serve on
	-d=".":    the directory of static files to host

Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	"golang.org/x/crypto/acme/autocert"
)

//var (
//	domain string
//)

func getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if !ok {
			dirCache = "certs"
		}
		keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
		crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			fmt.Printf("%s\nFalling back to Letsencrypt\n", err)
			return certManager.GetCertificate(hello)
		}
		fmt.Println("Loaded selfsigned certificate.")
		return &certificate, err
	}
}

// const DefaultACMEDirectory = "https://acme-v02.api.letsencrypt.org/directory"
const DefaultACMEDirectory = "https://acme-staging-v02.api.letsencrypt.org/directory"

func main() {
	port := flag.String("port", "80", "port to serve on")
	https_port := flag.String("https_port", "443", "HTTPS port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	domain := flag.String("domain", "dilworth.uk", "the domain")
	//flag.StringVar(&domain, "domain", "", "domain name to request your certificate")
	flag.Parse()

	//	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	//	log.Fatal(http.ListenAndServe(":"+*port, nil))

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(*directory)))

	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello HTTP/2")
	})
	fmt.Println("TLS domain", *domain)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(*domain),
		Cache:      autocert.DirCache("certs"),
	}
	//certManager.Client.DirectoryURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
	tlsConfig := certManager.TLSConfig()
	tlsConfig.GetCertificate = getSelfSignedOrLetsEncryptCert(&certManager)
	server := http.Server{
		Addr:      ":" + *https_port,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}
	go http.ListenAndServe(":"+*port, certManager.HTTPHandler(nil))
	fmt.Println("Server listening on", server.Addr)
	if err := server.ListenAndServeTLS("", ""); err != nil {
		fmt.Println(err)
	}
}
