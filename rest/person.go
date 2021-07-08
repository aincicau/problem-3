package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pb3/db"
	"pb3/entity"

	"github.com/sirupsen/logrus"
)

func PostVehicle(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var vehicle = entity.Vehicle{}
	err = json.Unmarshal(bodyBytes, &vehicle)
	if hasError(rw, err, "Internal issue") {
		return
	}

	db.GetDB().Create(&vehicle)
	rw.Write(bodyBytes)
}

func GetVehicle(rw http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("id")
	var vehicle entity.Vehicle
	result := db.GetDB().Find(&vehicle, "id=?", value)
	if result.RecordNotFound() {
		http.Error(rw, "No Record Found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
	}
	vehicleData, _ := json.Marshal(vehicle)
	rw.Write(vehicleData)
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)
	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}
	return false
}

func DeleteVehicle(rw http.ResponseWriter, r *http.Request) {
	idValue := r.URL.Query().Get("id")
	db.GetDB().Delete(&entity.Vehicle{}, "id=?", idValue)
}
