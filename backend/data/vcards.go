package data

import (
	"fmt"
)

//Vcard структура для хранения визитных карточек
type Vcard struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Workplace string `json:"workplace"`
	Address   string `json:"address"`
}

// VcardList структура для списка записей визиток
type VcardList struct {
	vcards []Vcard
}

//Editable интерфейс для работы со списком записей
type Editable interface {
	GetVcards() []Vcard
	AddVcard(vcard Vcard) int
	EditVcard(vcard Vcard, id int) error
	RemoveVcard(id int) error
}

// NewVcardList конструктор списка визиток
func NewVcardList() *VcardList {
	return &VcardList{}
}

//GetVcards возвращает список визитных карточек
func (vl *VcardList) GetVcards() []Vcard {
	return vl.vcards
}

//AddVcard добавляет карточку vcard в конец списка и возвращает id
func (vl *VcardList) AddVcard(vcard Vcard) int {
	id := len(vl.vcards) + 1
	vcard.ID = id
	vl.vcards = append(vl.vcards, vcard)
	return id
}

//EditVcard изменяет карточку c id на vcard
func (vl *VcardList) EditVcard(vcard Vcard, id int) error {
	if id < 1 || id > len(vl.vcards) {
		return fmt.Errorf("incorrect ID")
	}
	vl.vcards[id-1] = vcard
	return nil
}

//RemoveVcard удаляет контакт по id
func (vl *VcardList) RemoveVcard(id int) error {
	if id < 1 || id > len(vl.vcards) {
		return fmt.Errorf("incorrect ID")
	}
	id--
	vl.vcards = append(vl.vcards[:id], vl.vcards[id+1:]...)
	return nil
}
