package controllers

import (
	"backend/db"
	"backend/models"
	"backend/repositories"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func AddLike(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error
	var crypto models.Crypto

	params := mux.Vars(r)
	id := params["id"]

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	crypto, err = cryptoRepository.FindById(id)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"result": "Crypto not found. Please enter an existing id"})
		return
	}

	crypto.Likes += 1

	err = cryptoRepository.Update(&crypto)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "Like added"})
}

func RemoveLike(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error
	var crypto models.Crypto

	params := mux.Vars(r)
	id := params["id"]

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	crypto, err = cryptoRepository.FindById(id)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"result": "Crypto not found. Please enter an existing id"})
		return
	}

	crypto.Likes -= 1
	if crypto.Likes < 0 {
		crypto.Likes = 0
	}

	err = cryptoRepository.Update(&crypto)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "Like removed"})
}

func AddDislike(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error
	var crypto models.Crypto

	params := mux.Vars(r)
	id := params["id"]

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	crypto, err = cryptoRepository.FindById(id)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"result": "Crypto not found. Please enter an existing id"})
		return
	}

	crypto.Dislikes += 1

	err = cryptoRepository.Update(&crypto)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "Dislike added"})
}

func RemoveDislike(w http.ResponseWriter, r *http.Request) {
	var database *mgo.Database
	var err error
	var crypto models.Crypto

	params := mux.Vars(r)
	id := params["id"]

	database, err = db.Connect()
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	cryptoRepository := repositories.CryptoDAO{
		Db:         database,
		Collection: "cryptos",
	}

	crypto, err = cryptoRepository.FindById(id)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"result": "Crypto not found. Please enter an existing id"})
		return
	}

	crypto.Dislikes -= 1
	if crypto.Dislikes < 0 {
		crypto.Dislikes = 0
	}

	err = cryptoRepository.Update(&crypto)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "Dislike removed"})
}
