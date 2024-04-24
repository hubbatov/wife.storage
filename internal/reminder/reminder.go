package reminder

type Reminder struct {
	ID     int    `gorm:"primary_key" json:"id"`
	UserID int    `gorm:"index" json:"-"`
	Title  string `gorm:"type:varchar(100)" json:"title"`
	Text   string `gorm:"type:varchar(100)" json:"text"`
}
