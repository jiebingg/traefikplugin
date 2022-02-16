// Package traefikplugin a demo plugin.
package traefikplugin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

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
	log.Print("CERTIFICATE CN: ", req.TLS.PeerCertificates[0].Subject.CommonName)
	if req.TLS.PeerCertificates[0].Subject.CommonName != "localhost" {
		http.Error(rw, "Certificate provided is invalid.", http.StatusForbidden)
	}
	//for key, value := range a.headers {
	//	tmpl, err := a.template.Parse(value)
	//	if err != nil {
	//		http.Error(rw, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	writer := &bytes.Buffer{}
	//
	//	err = tmpl.Execute(writer, req)
	//	if err != nil {
	//		http.Error(rw, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	req.Header.Set(key, writer.String())
	//}

	a.next.ServeHTTP(rw, req)
}
