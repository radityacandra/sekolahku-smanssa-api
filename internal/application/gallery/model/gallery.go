package model

type Gallery struct {
	ID    int    `gorm:"column:id_galerifoto"`
	Title string `gorm:"column:judulgalerifoto"`
	Image string `gorm:"column:picture"`
	Date  string `gorm:"column:tgl"`
}

func (Gallery) TableName() string {
	return "t_galerifoto"
}
