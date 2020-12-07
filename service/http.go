
package service

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"
	_ "encoding/json"

	_ "github.com/TylerBrock/colorjson"
	"github.com/logrusorgru/aurora"
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
			return certManager.GetCertificate(hello)
		}
		return &certificate, err
	}
}

func Http_run(domain string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello HTTP/2")
		for u,v := range r.Header{
			fmt.Printf("%s: %s\n", aurora.Green(u),aurora.Blue(v))
		}
	
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
