package controllers

import (
	"backend/db"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CryptoSorter []models.Crypto

func (c CryptoSorter) Len() int           { return len(c) }
func (c CryptoSorter) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CryptoSorter) Less(i, j int) bool { return c[i].Likes > c[j].Likes }

func responseWithError(w http.ResponseWriter, code int, msg string) {
	responseWithJSON(w, code, map[string]string{"error": msg})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer database.Session.Close()

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	var cryptos []models.Crypto
	cryptos, err = cryptoRepository.FindAll()
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if cryptos != nil {
		sort.Sort(CryptoSorter(cryptos))
		responseWithJSON(w, http.StatusOK, cryptos)
	} else {
		responseWithJSON(w, http.StatusOK, map[string]string{"result": "There are no registered cryptos"})
	}
}

func FindById(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error
	var crypto models.Crypto

	params := mux.Vars(r)
	id := params["id"]

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer database.Session.Close()

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	crypto, err = cryptoRepository.FindById(id)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"result": "Crypto not found. Please enter an existing id"})
		return
	}

	responseWithJSON(w, http.StatusOK, crypto)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error
	var crypto models.Crypto

	crypto.Id = bson.NewObjectId()
	err = json.NewDecoder(r.Body).Decode(&crypto)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer database.Session.Close()

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	err = cryptoRepository.Create(&crypto)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJSON(w, http.StatusCreated, map[string]string{"result": "Crypto successfully created"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error
	var newCrypto models.Crypto
	var crypto models.Crypto

	err = json.NewDecoder(r.Body).Decode(&newCrypto)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer database.Session.Close()

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	crypto, err = cryptoRepository.FindById(newCrypto.Id.Hex())
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"result": "Crypto not found. Please enter an existing id"})
		return
	}

	newCrypto.Id = crypto.Id
	newCrypto.Likes = crypto.Likes
	newCrypto.Dislikes = crypto.Dislikes

	err = cryptoRepository.Update(&newCrypto)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "Crypto successfully updated"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error

	params := mux.Vars(r)
	id := params["id"]

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer database.Session.Close()

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	err = cryptoRepository.Delete(id)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "Crypto successfully deleted"})
}
