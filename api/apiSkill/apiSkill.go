package aboutcontroller

import (
	skillmodel "be-portfolio/models/skillModel"
	"encoding/json"
	"net/http"
)

func All(w http.ResponseWriter, r *http.Request) {
	skills := skillmodel.GetAll()

	skillsMap := make(map[string]map[string]string)
	for _, skill := range skills {
		skillsMap[skill.Name] = map[string]string{
			"Level":    skill.Level,
			"Category": skill.Category,
		}
	}

	data := map[string]interface{}{
		"skill": skillsMap,
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