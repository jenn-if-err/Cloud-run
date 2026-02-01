package function

import (
	"fmt"
	"net/http"
	"os"
)

// HTTP entry point for Google Cloud Functions
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = os.Getenv("NAME")
	}
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
