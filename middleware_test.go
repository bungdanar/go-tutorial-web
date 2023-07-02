package tutorialweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After execute handler")
}

type ErrHandlerMiddleware struct {
	Handler http.Handler
}

func (middleware *ErrHandlerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("Recover: ", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
		}
	}()

	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inner hanlder executed")
		fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic executed")
		panic("Ups")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	errHandlerMiddleware := new(ErrHandlerMiddleware)
	errHandlerMiddleware.Handler = logMiddleware

	server := http.Server{
		Addr:    "localhost:5500",
		Handler: errHandlerMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
