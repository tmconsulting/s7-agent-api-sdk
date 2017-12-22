package main

import (
	"gitlab.tm-consulting.ru/avia/services/s7-sdk"

	"encoding/xml"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
)

func main() {
	// name := "test/response_error.xml"
	// name := "test/response_3d_agreement.xml"
	// name := "test/response_transit_code_share.xml"
	// name := "test/response_transfer_spa.xml"
	name := "test/response_direct_s7_flight.xml"

	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}

	env := &sdk.Envelope{
		Body: &sdk.Body{
			AirDocDisplayRS: new(sdk.AirDocDisplayRS),
		},
	}

	err = xml.Unmarshal(data, env)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", spew.Sdump(env))
}