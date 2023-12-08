package model

type NewsCurriculum struct {
	ID      int    `gorm:"column:id_profilberitakurikulum"`
	Title   string `gorm:"column:judul"`
	Content string `gorm:"column:isi"`
	Author  string `gorm:"column:sumber"`
	Image   string `gorm:"column:picture"`
	Date    string `gorm:"column:tgl"`
}

func (NewsCurriculum) TableName() string {
	return "t_profil_berita_kurikulum"
}
