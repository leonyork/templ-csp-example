package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/leonyork/templ-csp-example/components"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Hx-Request") == "true" {
		handleHtmx(w, r)
		return
	}

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
	csp := fmt.Sprintf("connect-src 'self'; script-src 'sha256-%s' 'sha256-%s' '%s' 'unsafe-hashes'; style-src '%s'", base64.StdEncoding.EncodeToString(funcSha), base64.StdEncoding.EncodeToString(callSha), components.HTMX_SCRIPT_SHA, components.HTMX_STYLE_SHA)
	// ... and without 'unsafe-hashes'
	//csp := fmt.Sprintf( "connect-src 'self'; script-src 'sha256-%s' 'sha256-%s' '%s'; style-src '%s'", base64.StdEncoding.EncodeToString(funcSha), base64.StdEncoding.EncodeToString(callSha), components.HTMX_SCRIPT_SHA, components.HTMX_STYLE_SHA)
	// If we set the connect-src to be none, then the `hx-get` will fail to load.
	//csp := fmt.Sprintf("connect-src 'none'; script-src 'sha256-%s' 'sha256-%s' '%s' 'unsafe-hashes'; style-src '%s'", base64.StdEncoding.EncodeToString(funcSha), base64.StdEncoding.EncodeToString(callSha), components.HTMX_SCRIPT_SHA, components.HTMX_STYLE_SHA)

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

func handleHtmx(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("<div id=\"htmx\">Loaded htmx!</div><script type=\"text/javascript\">console.log(\"Loaded javascript!\")</script>")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
