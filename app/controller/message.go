package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yuhando/simpleapi/app/lib"
	"github.com/yuhando/simpleapi/app/model"
)

// PostMessage will save the received request(Message) to a file
func PostMessage(w http.ResponseWriter, r *http.Request) {

	pathFile := "app/assets/message.json"
	readMessageFile := lib.ReadFile(w, r, pathFile)
	messageStored := []model.Message{}
	json.Unmarshal(readMessageFile, &messageStored)

	messageModel := []model.Message{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&messageModel); err != nil {
		lib.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	currentTime := time.Now()
	formatedTime := currentTime.Format("2006-01-02 15:04:05")
	for i := 0; i < len(messageModel); i++ {
		messageModel[i].CreatedAt = formatedTime
		err := messageModel[i].Validate()
		if err != nil {
			lib.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		messageStored = append(messageStored, messageModel[i])
	}

	messageJSON, err := json.Marshal(messageStored)
	if err != nil {
		lib.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	lib.RespondJSON(w, http.StatusCreated, messageModel)
	_ = ioutil.WriteFile(pathFile, messageJSON, 0644)

}

// GetMessage get and return all previously message from the file that was created by PostMessage
func GetMessage(w http.ResponseWriter, r *http.Request) {
	pathFile := "app/assets/message.json"
	readMessageFile := lib.ReadFile(w, r, pathFile)
	messageStored := []model.Message{}
	json.Unmarshal(readMessageFile, &messageStored)
	if len(messageStored) == 0 {
		lib.RespondJSON(w, http.StatusOK, "There's no data to display")
		return
	}

	lib.RespondJSON(w, http.StatusOK, messageStored)
}
