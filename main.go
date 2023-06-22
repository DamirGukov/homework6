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
	fmt.Println("I send box to:", b.GetRecipientAdress(), ",from:", b.GetSenderAdress())
}

func (e Envelope) GetSenderAdress() string {
	return e.AdressSender
}

func (e Envelope) GetRecipientAdress() string {
	return e.AdressRecipient
}

func (e Envelope) Send() {
	fmt.Println("I send box to:", e.GetRecipientAdress(), ",from:", e.GetSenderAdress())
}

func SendPackage(p Package) {
	p.Send()

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

	SendPackage(b)
	SendPackage(e)
}
