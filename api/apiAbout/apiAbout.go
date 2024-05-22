package aboutcontroller

import (
	aboutmodel "be-portfolio/models/aboutModel"
	"encoding/json"
	"net/http"
)

func All(w http.ResponseWriter, r *http.Request) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	abouts := aboutmodel.GetAll()
	FOTO := aboutmodel.FOTO()

	aboutsMap := make(map[string]string)
	for _, about := range abouts {
		aboutsMap[about.Name] = about.Value
	}
	aboutsMap["photoAbout"] = scheme + r.Host + "/photo/" + FOTO.Address

	data := map[string]interface{}{
		"about": aboutsMap,
	}
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyetel header HTTP untuk menunjukkan bahwa konten adalah JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Menulis data JSON ke http.ResponseWriter
	w.Write(jsonData)

}