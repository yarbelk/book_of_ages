package parser

import (
	"encoding/xml"
	. "github.com/yarbelk/book_of_ages/db"
	"fmt"
)




func Parse(world *World, inputXML []byte) {
	err := xml.Unmarshal(inputXML, world)
	if err != nil {
		fmt.Printf("failed to unmarshal the xml", err)
		return
	}
}
