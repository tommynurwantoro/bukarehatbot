package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bot/bukarehatbot/entity"
	"github.com/bot/bukarehatbot/utility/mysql"
)

// CreateMicrobreakHistoryHandler _
func CreateMicrobreakHistoryHandler(w http.ResponseWriter, r *http.Request) {
	var microbreak entity.MicrobreakRequest
	_ = json.NewDecoder(r.Body).Decode(&microbreak)

	log.Printf("%v", microbreak)

	if microbreak.Username == "" || microbreak.GroupID == 0 {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	user := mysql.FindUserByUsernameAndGroupID(microbreak.Username, microbreak.GroupID)

	mysql.InsertHistory(user.ID, user.GroupID)

	response := entity.NewJSONResponse("Selamat kamu mendapatkan 100 poin karena telah melakukan microbreak tepat waktu", 200)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
