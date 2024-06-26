package model

type Headline struct {
	ID      int    `gorm:"column:id"`
	Title   string `gorm:"column:judul"`
	Content string `gorm:"column:isi"`
	Author  string `gorm:"column:sumber"`
	Image   string `gorm:"column:picture"`
	Date    string `gorm:"column:tgl"`
}

func (Headline) TableName() string {
	return "v_headline"
}
