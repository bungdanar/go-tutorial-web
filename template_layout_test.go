package tutorialweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/layout.gohtml", "./templates/header.gohtml", "./templates/footer.gohtml"))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Layout",
		"Name":  "Danar",
		// "Address": map[string]interface{}{
		// 	"Street": "Jalan XXX",
		// 	"City":   "Mexico City",
		// },
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5500", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
