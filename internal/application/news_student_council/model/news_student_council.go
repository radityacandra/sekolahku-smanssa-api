package model

type NewsStudentCouncil struct {
	ID      int    `gorm:"column:id_profilberitaosis"`
	Title   string `gorm:"column:judul"`
	Content string `gorm:"column:isi"`
	Author  string `gorm:"column:sumber"`
	Image   string `gorm:"column:picture"`
	Date    string `gorm:"column:tgl"`
}

func (NewsStudentCouncil) TableName() string {
	return "t_profil_berita_osis"
}
