package main

import "fmt"

type Package interface {
	GetSenderAdress() string
	GetRecipientAdress() string
	Send()
}

type Box struct {
	AdressSender    string
	AdressRecipient string
}

type Envelope struct {
	AdressSender    string
	AdressRecipient string
}

func (b Box) GetSenderAdress() string {
	return b.AdressSender
}

func (b Box) GetRecipientAdress() string {
	return b.AdressRecipient
}

func (b Box) Send() {
	fmt.Println("Sending box to:", b.GetRecipientAdress(), "from:", b.GetSenderAdress())
}

func (e Envelope) GetSenderAdress() string {
	return e.AdressSender
}

func (e Envelope) GetRecipientAdress() string {
	return e.AdressRecipient
}

func (e Envelope) Send() {
	fmt.Println("Sending envelope to:", e.GetRecipientAdress(), "from:", e.GetSenderAdress())
}

type SortingDepartment struct{}

func (sd SortingDepartment) SortAndSend(p Package, pack string) {
	
	fmt.Scan(&pack)
	if pack == "Box" || pack == "box" {
		fmt.Println("Sorting box...")
		p.Send()
	}
	
	fmt.Scan(&pack)
	if pack == "Envelope" || pack == "envelope" {
		fmt.Println("Sorting envelope...")
		p.Send()
	}
}

func main() {
	b := Box{
		AdressSender:    "12 High Street, London, SW1A 1AA",
		AdressRecipient: "27 Park Avenue, Manchester, M14 5PT",
	}

	e := Envelope{
		AdressSender:    "8 Queen's Road, Birmingham, B1 1RD",
		AdressRecipient: "45 Windsor Gardens, Edinburgh, EH1 2HU",
	}
	var pack string

	sd := SortingDepartment{}
	sd.SortAndSend(b, pack)
	sd.SortAndSend(e, pack)
}

