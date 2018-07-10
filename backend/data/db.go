package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DBVcardList struct {
	db *gorm.DB
}

func NewDBVcardList(connection string) (*DBVcardList, error) {
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Vcard{})
	return &DBVcardList{db: db}, db.Error
}

func (vl *DBVcardList) Close() {
	vl.db.Close()
}

func (vl *DBVcardList) GetVcards() []Vcard {
	var vcards []Vcard
	vl.db.Find(&vcards)
	return vcards
}

func (vl *DBVcardList) AddVcard(vcard Vcard) int {
	vl.db.Create(&vcard)
	return vcard.ID
}

func (vl *DBVcardList) EditVcard(vcard Vcard, id int) error {
	var vcards []Vcard
	vl.db.Where("id = ?", id).Find(&vcards)
	if len(vcards) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	vcard.ID = vcards[0].ID
	vl.db.Save(&vcard)
	return vl.db.Error
}

func (vl *DBVcardList) RemoveContact(id int) error {
	var vcards []Vcard
	vl.db.Where("id = ?", id).Find(&vcards)
	if len(vcards) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}
	vl.db.Delete(&vcards[0])
	return vl.db.Error
}
