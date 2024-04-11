package ticketing_system

import (
	"flag"
	"fmt"
	"lld/ticketing_system/inputreader"
	"lld/ticketing_system/readerbuilder"
)

func TicketingSystem() {

	reader := &inputreader.InputReader{}
	filePath := flag.String("filePath", "", "File from which to read instructions. Leave empty to read from stdin")
	flag.Parse()
	if *filePath == "" {
		fmt.Println("Enter the instructions. Type `exit`(case-sensitive) to exit:")
		reader.Read(readerbuilder.GetStdinReader(), "$ ", "exit")
	} else {
		inFile := readerbuilder.GetFileReader(*filePath)
		defer inFile.Close()
		reader.Read(inFile, "", "")
	}
}
