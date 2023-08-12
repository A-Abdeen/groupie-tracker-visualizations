package gt
import (
	"fmt"
	gt "gt/webapp/API"
	"html/template"
	"net/http"
)
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BaseHandler is called.")
	// Verify Request Method
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Path
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	Artists := gt.LoadArtist()
	t, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	t.Execute(w, Artists)
}

