package profilecontroller

import (
	"be-portfolio/entities"
	aboutmodel "be-portfolio/models/aboutModel"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func Index(w http.ResponseWriter, r *http.Request) {
	abouts := aboutmodel.GetAll()
	photo := aboutmodel.FOTO()


	data := map[string]interface{}{
		"abouts": abouts,
		"photoAddress":photo.Address,
	}

	var store = sessions.NewCookieStore([]byte("secret"))
	session, _ := store.Get(r, "session-name")
	successMessage, ok := session.Values["success"].(string)
    if ok {
        delete(session.Values, "success")
        session.Save(r, w)
		data["success"] = successMessage
    }

	errorMessage, okErr := session.Values["error"].(string)
    if okErr {
        delete(session.Values, "error")
        session.Save(r, w)
		data["error"] = errorMessage
    }

	temp, err := template.ParseFiles("views/about/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/about/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var about entities.About

		about.Name = r.FormValue("name")
		about.Value = r.FormValue("value")
		about.CreatedAt = time.Now()
		about.UpdatedAt = time.Now()

		var errorMessage string
		if about.Name == "" || about.Value == "" {
			// Jika ada kolom yang kosong, set pesan kesalahan
			errorMessage = "Name and value fields are required."
			temp, err := template.ParseFiles("views/about/create.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Menampilkan pesan kesalahan ke template
			temp.Execute(w, errorMessage)
			return
		}

		if ok := aboutmodel.Create(about); !ok {
			temp, _ := template.ParseFiles("views/about/create/html")
			temp.Execute(w, nil)
		}

		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["success"] = "Successfully added data"
		session.Save(r, w)

		http.Redirect(w, r, "/about", http.StatusSeeOther)

	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	
	temp, err := template.ParseFiles("views/about/edit.html")
	if err != nil {
		panic(err)
	}

	// Conversi idString to integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	about := aboutmodel.Detail(id)
	
	data := map[string]interface{}{
		"profile": about,
		"error":   "", 
		"success":   "", 
	}

	// IF Method GET
	if r.Method == "GET" {
		temp.Execute(w, data)
	}

	// IF Method POSt
	if r.Method == "POST" {
		var updatedAbout entities.About

		updatedAbout.Name = r.FormValue("name")
		updatedAbout.Value = r.FormValue("value")
		updatedAbout.UpdatedAt = time.Now()

		if updatedAbout.Name == "" || updatedAbout.Value == "" {
			data["error"] = "Name and value fields are required." 
			temp.Execute(w, data)
			return
		}

		if ok := aboutmodel.Update(id, updatedAbout); !ok {
			data["error"] = http.StatusSeeOther
			temp.Execute(w, data)
			return
		}

		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["success"] = "Successfully updated "+updatedAbout.Name
		session.Save(r, w)

		http.Redirect(w, r, "/about", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	// Get name before delete
	about := aboutmodel.Detail(id)

	if err := aboutmodel.Delete(id); err != nil{
		panic(err)
	}

	var store = sessions.NewCookieStore([]byte("secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["success"] = "Successfully Deleted "+about.Name
	session.Save(r, w)

	http.Redirect(w, r, "/about", http.StatusSeeOther)

}

