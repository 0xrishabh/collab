
package service

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"
	"encoding/json"

	"github.com/TylerBrock/colorjson"
	"golang.org/x/crypto/acme/autocert"

)


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

func Http_run(domain string) {
	var obj map[string]interface{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello HTTP/2")
		fmt.Println(r.Header)
		json.Unmarshal(r.Header, &obj)
		s, _ := colorjson.Marshal(obj)
		fmt.Println(string(s))
	
	})

	fmt.Println("TLS domain", domain)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache("certs"),
	}

	tlsConfig := certManager.TLSConfig()
	tlsConfig.GetCertificate = getSelfSignedOrLetsEncryptCert(&certManager)
	server := http.Server{
		Addr:      ":443",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	fmt.Println("Server listening on", server.Addr)
	if err := server.ListenAndServeTLS("", ""); err != nil {
		fmt.Println(err)
	}
}
