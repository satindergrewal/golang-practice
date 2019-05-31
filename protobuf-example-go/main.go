package main

import (
	"fmt"
	addressbookpb "golang-practice/protobuf-example-go/src/AddressBook"
	complexpb "golang-practice/protobuf-example-go/src/complex"
	enumpb "golang-practice/protobuf-example-go/src/enum_example"
	simplepb "golang-practice/protobuf-example-go/src/simple"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()

	doAddressBook()
}

func doAddressBook() {
	p1 := addressbookpb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*addressbookpb.Person_PhoneNumber{
			{Number: "555-4321", Type: addressbookpb.Person_HOME},
		},
	}

	book := addressbookpb.AddressBook{People: []*addressbookpb.Person{
		{
			Name:  "John Doe",
			Id:    101,
			Email: "john@example.com",
		},
		{
			Name: "Jane Doe",
			Id:   102,
		},
		{
			Name:  "Jack Doe",
			Id:    201,
			Email: "jack@example.com",
			Phones: []*addressbookpb.Person_PhoneNumber{
				{Number: "555-555-5555", Type: addressbookpb.Person_WORK},
			},
		},
		{
			Name:  "Jack Buck",
			Id:    301,
			Email: "buck@example.com",
			Phones: []*addressbookpb.Person_PhoneNumber{
				{Number: "555-555-0000", Type: addressbookpb.Person_HOME},
				{Number: "555-555-0001", Type: addressbookpb.Person_MOBILE},
				{Number: "555-555-0002", Type: addressbookpb.Person_WORK},
			},
		},
		{
			Name:  "Janet Doe",
			Id:    1001,
			Email: "janet@example.com",
			Phones: []*addressbookpb.Person_PhoneNumber{
				{Number: "555-777-0000"},
				{Number: "555-777-0001", Type: addressbookpb.Person_HOME},
			},
		},
	}}

	fmt.Println(p1)
	fmt.Println(book)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY

	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {

	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct: ", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{Indent: "   "}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)

	fmt.Println("Read the content: ", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to a file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffer's struct", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)

	// Always prefer to use helper funcs for such Id info instead of directly using sm.Id
	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
