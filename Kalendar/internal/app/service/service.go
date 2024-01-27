package service

import (
	"Kalendar/internal/model/reminderModel"
	"Kalendar/internal/model/userModel"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"sync"
)

func Hash(wg *sync.WaitGroup, value string) string {
	var s = append([]byte(value))
	// уменьшение счетчика на 1
	wg.Done()
	b := sha256.Sum256(s)
	return hex.EncodeToString(b[:])
}

func SerchUser(wg *sync.WaitGroup, personEmail string, personPassword string, people []userModel.User) userModel.User {

	personPassword = Hash(wg, personPassword)
	wg.Wait()
	for _, person := range people {
		if strings.ToLower(personEmail) == strings.ToLower(person.EmailUser) && person.Password == personPassword {
			return person
		}
	}
	return userModel.User{}
}

func SerchReminder(reminderID string, arrayCalendar []reminderModel.ReminderModel) reminderModel.ReminderModel {
	for _, reminder := range arrayCalendar {
		if reminderID == reminder.ID {
			return reminder
		}
	}
	return reminderModel.ReminderModel{}
}
