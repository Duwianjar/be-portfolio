package main

import (
	apiAbout "be-portfolio/api/apiAbout"
	apiProfile "be-portfolio/api/apiProfile"
	apiSkill "be-portfolio/api/apiSkill"
	"be-portfolio/config"
	aboutcontroller "be-portfolio/controller/aboutController"
	homecontroller "be-portfolio/controller/homeController"
	profilecontroller "be-portfolio/controller/profileController"
	skillcontroller "be-portfolio/controller/skillController"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	config.ConnectDB()
	router := mux.NewRouter()

	// For Handler Get /{id}
	http.Handle("/", router)

	// For Handler layout & File 
	router.HandleFunc("/navbar", homecontroller.Navbar)
	router.HandleFunc("/file/{name}", homecontroller.Files)
	router.HandleFunc("/photo/{name}", homecontroller.Photo)

	// 1. Home
	router.HandleFunc("/", homecontroller.Welcome)

	// 2. Profile
	http.HandleFunc("/profile", profilecontroller.Index)
	http.HandleFunc("/profile/add", profilecontroller.Add)
	http.HandleFunc("/profile/updateCV", profilecontroller.UpdateCV)
	http.HandleFunc("/profile/updatePP", profilecontroller.UpdatePP)
	router.HandleFunc("/profile/edit/{id}", profilecontroller.Edit)
	router.HandleFunc("/profile/delete/{id}", profilecontroller.Delete)

	// 3. About Me
	http.HandleFunc("/about", aboutcontroller.Index)
	http.HandleFunc("/about/add", aboutcontroller.Add)
	router.HandleFunc("/about/edit/{id}", aboutcontroller.Edit)
	router.HandleFunc("/about/delete/{id}", aboutcontroller.Delete)
	http.HandleFunc("/profile/updatePhotoAbout", aboutcontroller.UpdatePhoto)
	
	// 3. Skills 
	http.HandleFunc("/skill", skillcontroller.Index)
	http.HandleFunc("/skill/add", skillcontroller.Add)
	router.HandleFunc("/skill/edit/{id}", skillcontroller.Edit)
	router.HandleFunc("/skill/delete/{id}", skillcontroller.Delete)

	// API
	http.HandleFunc("/api/profile/all", apiProfile.All)
	http.HandleFunc("/api/about/all", apiAbout.All)
	http.HandleFunc("/api/skill/all", apiSkill.All)

	// START RUNNING SERVER
	port := 8080
	hyperlink := fmt.Sprintf("\033]8;;http://127.0.0.1:%d/\033\\%shttp://localhost:%d/\033[0m\033]8;;\033\\", port, "\033[1;32m", port)
	log.Printf("Server running on port %d: %s", port, hyperlink)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


}