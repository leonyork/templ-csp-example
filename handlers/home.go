package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/leonyork/templ-csp-example/components"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Need both the SHA of the function contents...
	funcSha, err := functionSha()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	
	
	//... And the sha on the `onload="..."` contents
	callSha, err := callSha()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// With 'unsafe-hashes'...
	csp := fmt.Sprintf( "default-src 'self'; script-src 'sha256-%s' 'sha256-%s' 'unsafe-hashes';", base64.StdEncoding.EncodeToString(funcSha), base64.StdEncoding.EncodeToString(callSha))
	// ... and without 'unsafe-hashes'
	//csp := fmt.Sprintf( "default-src 'self'; script-src 'sha256-%s' 'sha256-%s';", base64.StdEncoding.EncodeToString(funcSha), base64.StdEncoding.EncodeToString(callSha))

	w.Header().Add("Content-Security-Policy", csp)
	components.Page().Render(r.Context(), w)
}

func functionSha() ([]byte, error) {
	sha := sha256.New()

	if _, err := sha.Write(([]byte)(components.App().Function)); err != nil {
		return nil, err
	}

	return sha.Sum(nil), nil
}


func callSha() ([]byte, error) {
	sha := sha256.New()

	if _, err := sha.Write(([]byte)(components.App().Call)); err != nil {
		return nil, err
	}

	return sha.Sum(nil), nil
}