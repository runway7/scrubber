package scrubber

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

func init() {
	http.HandleFunc("/", infoHandler)
	http.HandleFunc("/strict", strictHandler)
	http.HandleFunc("/ugc", ugcHandler)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func strictHandler(w http.ResponseWriter, r *http.Request) {
	handleWithPolicy(w, r, bluemonday.StrictPolicy())
}

func ugcHandler(w http.ResponseWriter, r *http.Request) {
	handleWithPolicy(w, r, bluemonday.UGCPolicy())
}

func handleWithPolicy(w http.ResponseWriter, r *http.Request, policy *bluemonday.Policy) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Fprint(w, string(policy.SanitizeBytes(data)))
}
