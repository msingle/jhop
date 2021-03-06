package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/cooldrip/jhop"
)

func main() {
	f, err := os.Open("recipes.json")
	if err != nil {
		log.Fatalf("file opening failed: %s", err)
	}
	h, err := jhop.NewHandler(f)
	if err != nil {
		log.Fatalf("handler initialization failed: %s", err)
	}

	s := httptest.NewServer(h)
	defer s.Close()

	resp, err := http.Get(fmt.Sprintf("%s/recipes/1", s.URL))
	if err != nil {
		log.Fatalf("request to test server failed: %s", err)
	}

	io.Copy(os.Stdout, resp.Body) // {"difficulty":"hard","id":1,"prep_time":"1h"}
}
