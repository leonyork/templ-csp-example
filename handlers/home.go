package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/leonyork/templ-csp-example/components"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Now only need the SHA for the combined function and it's call.
	funcSha, err := functionSha()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	

	csp := fmt.Sprintf( "default-src 'self'; script-src 'sha256-%s';", base64.StdEncoding.EncodeToString(funcSha))

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