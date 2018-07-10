package data

import "testing"

var testVcards = []Vcard{
	{
		Name:      "Evgeny",
		Position:  "Zavedooushy Laboratory",
		Phone:     "+79998888899",
		Email:     "vyazilov@mail.ru",
		Workplace: "VNI GMI",
		Address:   "Central street",
	},
	{
		Name:      "Vasya",
		Position:  "Dvornik",
		Phone:     "+79997778899",
		Email:     "netu@mail.ru",
		Workplace: "Street",
		Address:   "no",
	},
}

func TestAddVcard(t *testing.T) {
	vl := NewVcardList()

	vl.AddVcard(testVcards[0])

	if vl.GetVcards()[0] != testVcards[0] {
		t.Errorf("AddVcard is no working")
	}
}

func TestEditVcard(t *testing.T) {
	vl := NewVcardList()
	vl.AddVcard(testVcards[0])

	err := vl.EditVcard(testVcards[1], 0)

	if vl.GetVcards()[0] != testVcards[1] {
		t.Errorf("EditVcard is no working")
	}
	if err != nil {
		t.Errorf("Unexpected EditVcard error")
	}

	err = vl.EditVcard(testVcards[1], -1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}

func TestRemoveVcard(t *testing.T) {
	vl := NewVcardList()
	vl.AddVcard(testVcards[0])
	vl.AddVcard(testVcards[1])

	err := vl.RemoveVcard(0)

	if vl.GetVcards()[0] != testVcards[1] {
		t.Errorf("RemoveVcard is no working")
	}
	if err != nil {
		t.Errorf("Unexpected RemoveVcard error")
	}

	err = vl.RemoveVcard(-1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}
