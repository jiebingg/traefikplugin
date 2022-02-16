package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/jiebingg/traefikplugin"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("TEST... SERVER RUNNING")

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	cNames := []string{"localhost", "test"}
	handler, err := traefikplugin.New(ctx, next, cNames, "demo-plugin")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", handler.ServeHTTP)
	caCertPEM, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal("read cert: ", err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")

	server := &http.Server{
		Addr: ":8000",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ClientCAs:    roots,
		},
	}
	err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
