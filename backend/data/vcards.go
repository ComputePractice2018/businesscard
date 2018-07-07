package data

//Vcard структура для хранения визитных карточек
type Vcard struct {
	Name      string `json:"name"`
	Position  string `json:"posion"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Workplace string `json:"workplace"`
	Address   string `json:"address"`
}

//VcardList хранит список визитных карточек
var VcardList = []Vcard{Vcard{
	Name:      "ФИО",
	Position:  "Должность",
	Phone:     "+7-999-999-99-99",
	Email:     "user@domain.ru",
	Workplace: "Место работы",
	Address:   "Адрес"}}
