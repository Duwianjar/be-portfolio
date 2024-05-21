package profilecontroller

import (
	"be-portfolio/entities"
	profilemodel "be-portfolio/models/profileModel"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func Index(w http.ResponseWriter, r *http.Request) {
	profiles := profilemodel.GetAll()
	CV := profilemodel.CV()
	PP := profilemodel.PP()

	data := map[string]interface{}{
		"profiles": profiles,
		"cvAddress":CV.Address,
		"ppAddress":PP.Address,
	}

	// If isset message success
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

	temp, err := template.ParseFiles("views/profile/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/profile/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var profile entities.Profile

		profile.Name = r.FormValue("name")
		profile.Value = r.FormValue("value")
		profile.CreatedAt = time.Now()
		profile.UpdatedAt = time.Now()

		var errorMessage string
		if profile.Name == "" || profile.Value == "" {
			// Jika ada kolom yang kosong, set pesan kesalahan
			errorMessage = "Name and value fields are required."
			temp, err := template.ParseFiles("views/profile/create.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Menampilkan pesan kesalahan ke template
			temp.Execute(w, errorMessage)
			return
		}

		if ok := profilemodel.Create(profile); !ok {
			temp, _ := template.ParseFiles("views/profile/create/html")
			temp.Execute(w, nil)
		}

		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["success"] = "Successfully added data"
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)

	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	
	temp, err := template.ParseFiles("views/profile/edit.html")
	if err != nil {
		panic(err)
	}

	// Conversi idString to integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	profile := profilemodel.Detail(id)
	
	data := map[string]interface{}{
		"profile": profile,
		"error":   "", 
		"success":   "", 
	}

	// IF Method GET
	if r.Method == "GET" {
		temp.Execute(w, data)
	}

	// IF Method POSt
	if r.Method == "POST" {
		var updatedProfile entities.Profile

		updatedProfile.Name = r.FormValue("name")
		updatedProfile.Value = r.FormValue("value")
		updatedProfile.UpdatedAt = time.Now()

		if updatedProfile.Name == "" || updatedProfile.Value == "" {
			data["error"] = "Name and value fields are required." 
			temp.Execute(w, data)
			return
		}

		if ok := profilemodel.Update(id, updatedProfile); !ok {
			data["error"] = http.StatusSeeOther
			temp.Execute(w, data)
			return
		}

		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["success"] = "Successfully updated "+updatedProfile.Name
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
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
	profile := profilemodel.Detail(id)

	if err := profilemodel.Delete(id); err != nil{
		panic(err)
	}

	var store = sessions.NewCookieStore([]byte("secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["success"] = "Successfully Deleted "+profile.Name
	session.Save(r, w)

	http.Redirect(w, r, "/profile", http.StatusSeeOther)

}

func UpdateCV(w http.ResponseWriter, r *http.Request) {
	// check method post
	if r.Method != http.MethodPost {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Method not allowed"
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}

	// Menerima file dari form dengan nama "pdfFile"
	file, handler, err := r.FormFile("pdfFile")
	if err != nil {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Failed to retrieve file"
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
	defer file.Close()

	// Membuat nama file baru dengan format "cv_tanggalwaktu.pdf"
	filename := fmt.Sprintf("cv_%s.pdf", time.Now().Format("20060102150405"))

	// Mendefinisikan path lengkap untuk menyimpan file
	savePath := filepath.Join("assets/files/", filename)

	// Membuat file baru di server
	dst, err := os.Create(savePath)
	if err != nil {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Failed to create file : "+handler.Filename
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
	defer dst.Close()

	// Menyalin data file ke file yang baru dibuat di server
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// ambil nama file sebelum di update
	CV := profilemodel.CV()

	// mengupdate data di sql
	if ok := profilemodel.UpdateCV(filename); !ok {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Failed to update database CV "
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}

	// hapus file lama
    if _, err := os.Stat("assets/files/" + CV.Address); err == nil {
        err := os.Remove("assets/files/" + CV.Address)
        if err != nil {
			var store = sessions.NewCookieStore([]byte("secret"))
			session, _ := store.Get(r, "session-name")
			session.Values["error"] = "Failed to delete existing file CV "
			session.Save(r, w)

			http.Redirect(w, r, "/profile", http.StatusSeeOther)
        }
    }
	
	var store = sessions.NewCookieStore([]byte("secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["success"] = "Successfully to update CV "
	session.Save(r, w)

	http.Redirect(w, r, "/profile", http.StatusSeeOther)

}

func UpdatePP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Method not allowed"
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
	
	file, handler, err := r.FormFile("photoFile")
	if err != nil {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Failed to retrieve file"
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
	defer file.Close()

	extension := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("pp_%s"+extension, time.Now().Format("20060102150405"))

	savePath := filepath.Join("assets/img/", filename)

	dst, err := os.Create(savePath)
	if err != nil {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Failed to create file : "+handler.Filename
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
	
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	PP := profilemodel.PP()

	if ok := profilemodel.UpdatePP(filename); !ok {
		var store = sessions.NewCookieStore([]byte("secret"))
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Failed to update database CV "
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
	
    if _, err := os.Stat("assets/img/" + PP.Address); err == nil {
        err := os.Remove("assets/img/" + PP.Address)
        if err != nil {
			var store = sessions.NewCookieStore([]byte("secret"))
			session, _ := store.Get(r, "session-name")
			session.Values["error"] = "Failed to delete existing file CV "
			session.Save(r, w)

			http.Redirect(w, r, "/profile", http.StatusSeeOther)
        }
    }
	
	var store = sessions.NewCookieStore([]byte("secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["success"] = "Successfully to update Photo Profile "
	session.Save(r, w)

	http.Redirect(w, r, "/profile", http.StatusSeeOther)

}



