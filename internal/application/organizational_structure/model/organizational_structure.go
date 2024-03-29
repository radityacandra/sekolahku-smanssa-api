package model

type OrganizationalStructure struct {
	ID    string `gorm:"column:kd"`
	Name  string `gorm:"column:nmstrukturorganisasidetail"`
	Image string `gorm:"column:picturestrukturorganisasidetail"`
	Title string `gorm:"column:keteranganstrukturorganisasidetail"`
}

func (OrganizationalStructure) TableName() string {
	return "v_profilstrukturorganisasi"
}
