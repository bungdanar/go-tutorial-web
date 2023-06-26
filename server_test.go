package tutorialweb

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:5500",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
