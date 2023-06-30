package main

import "fmt"

type Package interface {
	GetSenderAddress() string
	GetRecipientAddress() string
	Send()
}

type Box struct {
	AddressSender    string
	AddressRecipient string
}

type Envelope struct {
	AddressSender    string
	AddressRecipient string
}

func (b Box) GetSenderAddress() string {
	return b.AddressSender
}

func (b Box) GetRecipientAddress() string {
	return b.AddressRecipient
}

func (b Box) Send() {
	fmt.Println("Sending box to:", b.GetRecipientAddress(), "from:", b.GetSenderAddress())
}

func (e Envelope) GetSenderAddress() string {
	return e.AddressSender
}

func (e Envelope) GetRecipientAddress() string {
	return e.AddressRecipient
}

func (e Envelope) Send() {
	fmt.Println("Sending envelope to:", e.GetRecipientAddress(), "from:", e.GetSenderAddress())
}

type SortingDepartment struct {
	deliverySave map[int]int
}

func (sd *SortingDepartment) SortAndSend(p Package) {

	var delivery string

	fmt.Println("Fast or regular delivery for the box?")
	fmt.Scan(&delivery)

	if delivery == "Fast" || delivery == "fast" {
		fmt.Println("Sorting box and sending it by fast delivery")
		sd.deliverySave[1]++
		p.Send()
	} else if delivery == "Regular" || delivery == "regular" {
		fmt.Println("Sorting box and sending it by regular delivery")
		sd.deliverySave[2]++
		p.Send()
	}

	fmt.Println("Fast delivery:", sd.deliverySave[1])
	fmt.Println("Regular delivery:", sd.deliverySave[2])
}
func main() {
	b := Box{
		AddressSender:    "12 High Street, London, SW1A 1AA",
		AddressRecipient: "27 Park Avenue, Manchester, M14 5PT",
	}

	e := Envelope{
		AddressSender:    "8 Queen's Road, Birmingham, B1 1RD",
		AddressRecipient: "45 Windsor Gardens, Edinburgh, EH1 2HU",
	}

	sd := SortingDepartment{
		deliverySave: make(map[int]int),
	}

	sd.deliverySave = map[int]int{
		1: 0,
		2: 0,
	}

	sd.SortAndSend(&b)
	sd.SortAndSend(&e)
}


