package repository

import (
	"wife/internal/reminder"
	"wife/internal/user"
)

func (d *Repository) Reminders() []reminder.Reminder {
	var table []reminder.Reminder
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *Repository) RemindersFromUser(login string) []reminder.Reminder {
	var user user.User
	var reminders []reminder.Reminder
	result := d.DataBase.Where("login = ?", login).First(&user)
	if result.Error != nil {
		return reminders
	}

	return user.Reminders
}
