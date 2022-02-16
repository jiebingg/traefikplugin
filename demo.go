// Package traefikplugin a demo plugin.
package traefikplugin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Demo a Demo plugin.
type CertValidator struct {
	next       http.Handler
	allowedCNs []string
	name       string
	template   *template.Template
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, CNames []string, name string) (http.Handler, error) {
	if len(CNames) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &CertValidator{
		allowedCNs: CNames,
		next:       next,
		name:       name,
	}, nil
}

func (a *CertValidator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Print("HANDLING")
	rw.Write([]byte("This is an example server.\n"))
	currCN := req.TLS.PeerCertificates[0].Subject.CommonName
	log.Print("CERTIFICATE CN: ", currCN)
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
