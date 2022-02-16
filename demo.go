// Package traefikplugin a demo plugin.
package traefikplugin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

// Demo a Demo plugin.
type CertValidator struct {
	next       http.Handler
	allowedCNs []string
	name       string
	template   *template.Template
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &CertValidator{
		allowedCNs: strings.Split(config.Headers["CNames"], ","),
		next:       next,
		name:       name,
	}, nil
}

func (a *CertValidator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Print("HANDLING")
	rw.Write([]byte("This is an example server.\n"))
	currCN := req.TLS.PeerCertificates[0].Subject.CommonName
	log.Print("CERTIFICATE CN: ", currCN)
	log.Print("ALLOWED CNs: ", a.allowedCNs)
	log.Print("RESULT: ", contains(a.allowedCNs, currCN))
	if !contains(a.allowedCNs, currCN) {
		http.Error(rw, "Certificate provided is invalid.", http.StatusForbidden)
	}
	a.next.ServeHTTP(rw, req)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
