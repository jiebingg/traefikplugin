package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	caCertPEM, err := ioutil.ReadFile("ca.crt")
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")

	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				RootCAs:            roots,
				InsecureSkipVerify: true,
			},
		},
	}

	// Make a request
	r, err := client.Get("https://127.0.0.1:8000/hello")
	if err != nil {
		log.Fatalf("Error with GET request: %s", err)
	}

	log.Print("RESPONSE HEADER: ", r.Header)

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Print("RESPONSE BODY: ", bodyString)

	log.Print("client: exiting")
}
