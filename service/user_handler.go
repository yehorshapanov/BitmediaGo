package service

import (
	"encoding/json"
	"github.com/yehorshapanov/BitmediaGo/db"
	"github.com/yehorshapanov/BitmediaGo/logger"
	"net/http"
)

func getAllUsersHandle(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		users, err := deps.DB.ListAllUsers(req.Context())
		if err != nil {
			logger.Get().Info("Error fetching data", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(users)
		if err != nil {
			logger.Get().Info("Error marshaling data", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func createUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)

		var u db.User
		err := decoder.Decode(&u)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = deps.DB.Create(req.Context(), u)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	})
}