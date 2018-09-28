package main

import (
	"fmt"
	"github.com/carousell/ProtobufGrpc/complex"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	someStuff()
}

func someStuff() {

	//sm := testpackage.SimpleMessage {
	//
	//	FirstName:"Vipul",
	//	LastName:"Sodha",
	//	Age:24,
	//}
	//
	////writeToFile("test", &sm)
	////
	////
	//sm2 := &testpackage.SimpleMessage{}
	////
	////readFromFile("test", sm2)
	////
	////fmt.Println(sm2.GetFirstName())
	//
	//json, _ := toJson(&sm)
	//
	//fromJson(json, sm2)
	//
	//fmt.Println(sm2)

	//en := simple_enum.SomeMessage{
	//	Id:1,
	//	E:simple_enum.SomeEnum_ONE,
	//}
	//
	//
	//fmt.Println(en)


	dummy_one := complexpb.DummyMessage{
		Id:1,
		Name: "Vipul",
	}

	dummy_two := complexpb.DummyMessage{
		Id:1,
		Name: "Vipul",
	}

	dummy_three := complexpb.DummyMessage{
		Id:1,
		Name: "Vipul",
	}


	comp := complexpb.ComplexMessage{
		OneDummy: &dummy_one,
		MultipleDummy:[]*complexpb.DummyMessage {
			&dummy_two,
			&dummy_three,
		},
	}


	fmt.Println(comp)

}

func fromJson(in string, pb proto.Message)  {

	jsonpb.UnmarshalString(in,pb)

}

func toJson(pb proto.Message) (string, error) {

	marshaler := jsonpb.Marshaler{}

	out, err := marshaler.MarshalToString(pb)

	if err != nil {
		return "", err
	}

	return out, nil

}


func readFromFile(fname string, pb proto.Message)  error {

	in, err := ioutil.ReadFile(fname)

	if err != nil {
		return err
	}

	proto.Unmarshal(in, pb)

	fmt.Println("Done reading")
	return nil
}

func writeToFile(fname string, pb proto.Message)  error{

	out, err := proto.Marshal(pb)

	if err !=nil {
		log.Fatal(err)
		return err
	}

	ioutil.WriteFile(fname, out, 0644)

	fmt.Print("Done")

	return nil

}
