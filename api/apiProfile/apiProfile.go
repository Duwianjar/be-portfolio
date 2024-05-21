package profilecontroller

import (
	profilemodel "be-portfolio/models/profileModel"
	"encoding/json"
	"net/http"
)

func All(w http.ResponseWriter, r *http.Request) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	profiles := profilemodel.GetAll()
	CV := profilemodel.CV()
	PP := profilemodel.PP()

	profilesMap := make(map[string]string)
	for _, profile := range profiles {
		profilesMap[profile.Name] = profile.Value
	}
	profilesMap["CV"] = scheme + r.Host + "/file/" + CV.Address
	profilesMap["photoProfile"] = scheme + r.Host + "/photo/" + PP.Address

	data := map[string]interface{}{
		"profile": profilesMap,
	}
	// Meng-encode data ke JSON
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