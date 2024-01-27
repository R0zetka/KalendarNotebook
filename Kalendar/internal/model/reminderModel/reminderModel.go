package reminderModel

import (
	"Kalendar/internal/model/userModel"
	"time"
)

type ReminderModel struct {
	ID                string         `json:"id"`
	NameReminder      string         `json:"namereminder"`
	Description       string         `json:"description"`
	StatusReminder    bool           `json:"statusreminder"`
	DateReminderStart time.Time      `json:"datereminderstart"`
	DateReminderFnish time.Time      `json:"datereminderfinish"`
	User              userModel.User `json:"user"`
}
