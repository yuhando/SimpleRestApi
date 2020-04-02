package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// RespondNoContent makes the response with payload as json format
func RespondNoContent(w http.ResponseWriter, statusCode int, r *http.Request) {
	allowedMethod := r.Header.Get("Access-Control-Request-Method")
	xForwardFor := r.Header.Get("X-Forwarded-For")
	currentTIme := time.Now().Round(0)
	dateUTC := currentTIme.UTC().Format(http.TimeFormat)

	w.Header().Set("Date", dateUTC)
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Server", "Kestrel")
	w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
	w.Header().Set("Access-Control-Allow-Methods", allowedMethod)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Processing-Time", "0")
	w.Header().Set("X-Forwarded-For", xForwardFor)
	w.WriteHeader(statusCode)
}

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, statusCode int, paylod interface{}) {
	response, err := json.Marshal(paylod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, statusCode int, message string) {
	RespondJSON(w, statusCode, map[string]string{"error": message})
}

// RespondBlank ...
func RespondBlank(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

// CheckFile if the file does not exist then it will create new one
func CheckFile(pathToFile string) error {
	_, err := os.Stat(pathToFile)
	if os.IsNotExist(err) {
		_, err := os.Create(pathToFile)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadFile ...
func ReadFile(w http.ResponseWriter, r *http.Request, pathFile string) []byte {

	err := CheckFile(pathFile)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	File, err := os.Open(pathFile)

	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return nil
	}

	defer File.Close()

	convertToByteVal, _ := ioutil.ReadAll(File)
	return convertToByteVal
}
