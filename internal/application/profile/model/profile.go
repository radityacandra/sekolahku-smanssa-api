package model

type Profile struct {
	ID                        int    `gorm:"column:id_berandaprofilvisimisi" json:"id"`
	Code                      string `gorm:"column:kdvisimisi" json:"code"`
	Vision                    string `gorm:"column:nmvisi" json:"vision"`
	Mission                   string `gorm:"column:nmmisi" json:"mission"`
	Motto                     string `gorm:"column:motto" json:"motto"`
	MottoDetail               string `gorm:"column:deskripsimotto" json:"mottoDetail"`
	Service                   string `gorm:"column:jenispelayanan" json:"service"`
	Aim                       string `gorm:"column:tujuansekolah" json:"aim"`
	History                   string `gorm:"column:sejarahsingkat" json:"history"`
	HistoryYoutubeID          string `gorm:"column:fileyoutubesejarahsingkat" json:"historyYoutubeId"`
	Hymn                      string `gorm:"column:hymnesekolah" json:"hymn"`
	HymnYoutubeID             string `gorm:"column:filealamathymneyoutube" json:"hymnYoutubeId"`
	OrganizationalStructure   string `gorm:"column:filestrukturorganisasi" json:"organizationalStructure"`
	SchoolCommittee           string `gorm:"column:filekomitesekolah" json:"schoolCommittee"`
	AcademicCalendar          string `gorm:"column:filekalenderakademik" json:"academicCalendar"`
	PrincipalWelcome          string `gorm:"column:sambutankepalasekolah" json:"principalWelcome"`
	PrincipalWelcomeYoutubeID string `gorm:"column:fileyoutubesambutankepalasekolah" json:"principalWelcomeYoutubeId"`
	VisionMissionYoutubeID    string `gorm:"column:fileyoutubevisimisi" json:"visionMissionYoutubeId"`
}

func (Profile) TableName() string {
	return "t_berandaprofilvisimisi"
}
