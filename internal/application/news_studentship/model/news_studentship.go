package model

type NewsStudentship struct {
	ID      int    `gorm:"column:id_profilberitakesiswaan"`
	Title   string `gorm:"column:judul"`
	Content string `gorm:"column:isi"`
	Author  string `gorm:"column:sumber"`
	Image   string `gorm:"column:picture"`
	Date    string `gorm:"column:tgl"`
}

func (NewsStudentship) TableName() string {
	return "t_profil_berita_kesiswaan"
}
