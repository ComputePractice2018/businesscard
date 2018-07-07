package data

import (
	"fmt"
)

//Vcard структура для хранения визитных карточек
type Vcard struct {
	Name      string `json:"name"`
	Position  string `json:"posion"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Workplace string `json:"workplace"`
	Address   string `json:"address"`
}

// vcards хранит список визитных карточек
var vcards []Vcard

//GetVcards возвращает список визитных карточек
func GetVcards() []Vcard {
	return vcards
}

//AddVcard добавляет карточку vcard в конец списка и возвращает id
func AddVcard(vcard Vcard) int {
	id := len(vcards)
	vcards = append(vcards, vcard)
	return id
}

//EditVcard изменяет карточку c id на vcard
func EditVcard(vcard Vcard, id int) error {
	if id < 0 || id >= len(vcards) {
		return fmt.Errorf("incorrect ID")
	}
	vcards[id] = vcard
	return nil
}

//RemoveVcard удаляет контакт по id
func RemoveVcard(id int) error {
	if id < 0 || id >= len(vcards) {
		return fmt.Errorf("incorrect ID")
	}
	vcards = append(vcards[:id], vcards[id+1:]...)
	return nil
}
