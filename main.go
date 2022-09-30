package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"deadly.surgery/t/esp32in/fizzbuzz"
)

func main() {

	gm := fizzbuzz.GameManager{
		Games: make(map[string]*fizzbuzz.Game),
	}

	http.HandleFunc("/fizzbuzz", gm.Game)

	http.HandleFunc("/update", handleUpdate)

	http.ListenAndServe(":9000", nil)
}

func handleUpdate(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		rw.WriteHeader(http.StatusBadRequest)
		Encode(rw, message{Error: "No ID Provided"})
		return
	}

	log.Printf("Request from %v (%v)", id, r.RemoteAddr)
	// fmt.Printf("Request from #%v (%v)\n", id, r.RemoteAddr)
}

// Encode writes the body as json to the writer.
func Encode(w io.Writer, body interface{}) {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	if err := e.Encode(body); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encode response: %v\n", err)
	}
}

type message struct {
	Error string `json:"error"`
}
