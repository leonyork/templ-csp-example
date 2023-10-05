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

	// Now only need the SHA for the combined function and it's call.
	funcSha, err := functionSha()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	csp := fmt.Sprintf("connect-src 'self'; script-src 'sha256-%s' '%s'; style-src '%s'", base64.StdEncoding.EncodeToString(funcSha), components.HTMX_SCRIPT_SHA, components.HTMX_STYLE_SHA)

	w.Header().Add("Content-Security-Policy", csp)
	components.Page().Render(r.Context(), w)
}

func functionSha() ([]byte, error) {
	sha := sha256.New()

	script := components.App()

	if _, err := sha.Write(([]byte)(script.Function + ";" + script.Call)); err != nil {
		return nil, err
	}

	return sha.Sum(nil), nil
}

func handleHtmx(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("<div id=\"htmx\">Loaded htmx!</div>")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
