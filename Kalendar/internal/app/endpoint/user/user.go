package user

import (
	"Kalendar/internal/app/service"
	"Kalendar/internal/model/userModel"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"sync"
)

var People []userModel.User
var wg sync.WaitGroup

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user userModel.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	for _, person := range People {
		if strings.ToLower(person.EmailUser) == strings.ToLower(user.EmailUser) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	user.EmailUser = strings.ToLower(user.EmailUser)

	// счетчик 1
	wg.Add(1)
	b := service.Hash(&wg, user.Password)
	//ожидание когда счетчик будет равен 0
	wg.Wait()
	// пароль кодировки sha256 сохраняется в string
	user.Password = b
	People = append(People, user)
	user.Password = ""
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	personEmail := params["personEmail"]
	personPassword := params["personPassword"]

	wg.Add(1)

	person := service.SerchUser(&wg, personEmail, personPassword, People)
	wg.Wait()
	if person.ID != "" {
		person.Password = ""
		err := json.NewEncoder(w).Encode(person)
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
