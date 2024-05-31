package skillcontroller

import (
	"be-portfolio/entities"
	skillmodel "be-portfolio/models/skillModel"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func Index(w http.ResponseWriter, r *http.Request) {
	skills := skillmodel.GetAll()


	data := map[string]interface{}{
		"skills": skills,
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

	temp, err := template.ParseFiles("views/skill/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/skill/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var skill entities.Skill

		skill.Name = r.FormValue("name")
		skill.Level = r.FormValue("level")
		skill.Category = r.FormValue("category")
		skill.CreatedAt = time.Now()
		skill.UpdatedAt = time.Now()

		var errorMessage string
		if skill.Name == "" || skill.Level == "" || skill.Category == "" {
			// Jika ada kolom yang kosong, set pesan kesalahan
			errorMessage = "Name, Level and Category fields are required."
			temp, err := template.ParseFiles("views/skill/create.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Menampilkan pesan kesalahan ke template
			temp.Execute(w, errorMessage)
			return
		}

		if ok := skillmodel.Create(skill); !ok {
			temp, _ := template.ParseFiles("views/skill/create/html")
			temp.Execute(w, nil)
		}

		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["success"] = "Successfully added data"
		session.Save(r, w)

		http.Redirect(w, r, "/skill", http.StatusSeeOther)

	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	
	temp, err := template.ParseFiles("views/skill/edit.html")
	if err != nil {
		panic(err)
	}

	// Conversi idString to integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	skill := skillmodel.Detail(id)
	
	data := map[string]interface{}{
		"skill": skill,
		"error":   "", 
		"success":   "", 
	}

	// IF Method GET
	if r.Method == "GET" {
		temp.Execute(w, data)
	}

	// IF Method POSt
	if r.Method == "POST" {
		var updatedSkill entities.Skill

		updatedSkill.Name = r.FormValue("name")
		updatedSkill.Level = r.FormValue("level")
		updatedSkill.Category = r.FormValue("category")
		updatedSkill.UpdatedAt = time.Now()

		if updatedSkill.Name == "" || updatedSkill.Level == "" || updatedSkill.Category == "" {
			data["error"] = "Name, Level and Category fields are required." 
			temp.Execute(w, data)
			return
		}

		if ok := skillmodel.Update(id, updatedSkill); !ok {
			data["error"] = http.StatusSeeOther
			temp.Execute(w, data)
			return
		}

		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["success"] = "Successfully updated "+updatedSkill.Name
		session.Save(r, w)

		http.Redirect(w, r, "/skill", http.StatusSeeOther)
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
	skill := skillmodel.Detail(id)

	if err := skillmodel.Delete(id); err != nil{
		panic(err)
	}

	var store = sessions.NewCookieStore([]byte("secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["success"] = "Successfully Deleted " + skill.Name
	session.Save(r, w)

	http.Redirect(w, r, "/skill", http.StatusSeeOther)

}