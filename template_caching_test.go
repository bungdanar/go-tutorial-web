package tutorialweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templatesForCache embed.FS

var myTemplates = template.Must(template.ParseFS(templatesForCache, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, h *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello Bro!")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5500", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
