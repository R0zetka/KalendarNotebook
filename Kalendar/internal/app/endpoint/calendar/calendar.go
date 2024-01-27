package calendar

import (
	"Kalendar/internal/app/endpoint/user"
	"Kalendar/internal/app/service"
	"Kalendar/internal/model/reminderModel"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
)

var Calendar []reminderModel.ReminderModel
var wg sync.WaitGroup

func CreateReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reminder reminderModel.ReminderModel
	params := mux.Vars(r)
	personEmail := params["personEmail"]
	personPassword := params["personPassword"]

	wg.Add(1)
	person := service.SerchUser(&wg, personEmail, personPassword, user.People)
	wg.Wait()
	if person.ID != "" {

		err := json.NewDecoder(r.Body).Decode(&reminder)
		if err != nil {
			return
		}
	}

	if reminder.NameReminder != "" || reminder.ID != "" {
		reminder.User = person
		if reminder.DateReminderStart.Unix() < time.Now().Unix() || reminder.DateReminderFnish.Unix() < reminder.DateReminderStart.Unix() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		reminder.StatusReminder = false
		Calendar = append(Calendar, reminder)
		reminder.User.Password = ""
		err := json.NewEncoder(w).Encode(reminder)
		if err != nil {
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusNotFound)

}

func GetReminder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personEmail := params["personEmail"]
	personPassword := params["personPassword"]
	reminderID := params["reminderID"]
	wg.Add(1)
	person := service.SerchUser(&wg, personEmail, personPassword, user.People)
	wg.Wait()
	if person.ID != "" {
		reminder := service.SerchReminder(reminderID, Calendar)
		if reminder.ID != "" {
			reminder.User.Password = ""
			err := json.NewEncoder(w).Encode(reminder)
			if err != nil {
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateStatusReminder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personEmail := params["personEmail"]
	personPassword := params["personPassword"]
	reminderID := params["reminderID"]
	wg.Add(1)
	person := service.SerchUser(&wg, personEmail, personPassword, user.People)
	wg.Wait()
	if person.ID != "" {
		for i, reminder := range Calendar {
			if reminderID == reminder.ID {
				Calendar[i].StatusReminder = !reminder.StatusReminder
				reminder.StatusReminder = !reminder.StatusReminder
				reminder.User.Password = ""
				err := json.NewEncoder(w).Encode(reminder)
				if err != nil {
					return
				}
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateReminder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personEmail := params["personEmail"]
	personPassword := params["personPassword"]
	reminderID := params["reminderID"]
	wg.Add(1)
	person := service.SerchUser(&wg, personEmail, personPassword, user.People)
	wg.Wait()
	var newReminder reminderModel.ReminderModel
	if person.ID != "" {
		err := json.NewDecoder(r.Body).Decode(&newReminder)
		if err != nil {
			return
		}
		for i := range Calendar {
			if reminderID == Calendar[i].ID {
				if newReminder.DateReminderStart.Unix() < time.Now().Unix() || newReminder.DateReminderFnish.Unix() < newReminder.DateReminderStart.Unix() || newReminder.NameReminder == "" {
					w.WriteHeader(http.StatusBadRequest)
				} else {

					newReminder.User = Calendar[i].User
					newReminder.ID = Calendar[i].ID
					Calendar[i] = newReminder
					newReminder.User.Password = ""
					err := json.NewEncoder(w).Encode(newReminder)
					if err != nil {
						return
					}
					return

				}

			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteReminder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personEmail := params["personEmail"]
	personPassword := params["personPassword"]
	reminderID := params["reminderID"]
	wg.Add(1)
	person := service.SerchUser(&wg, personEmail, personPassword, user.People)
	wg.Wait()
	if person.ID != "" {
		for i, reminder := range Calendar {
			if reminder.ID == reminderID {
				Calendar = append(Calendar[:i], Calendar[i+1:]...)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
