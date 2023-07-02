package tutorialweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	http.ServeFile(w, r, fmt.Sprintf("./resources/%s", fileName))
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5500",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
