package main

import (
	"log"

	"github.com/youruser/yourrepo/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "SomeName"
	log.Println(myVar.TypeName)
}

//To create a package firstly go in the root level of the directory and run command go mod init github.com/chayanbajaj/projectname
//Then in the root level directory create a folder with the name of package
//inside the created folder place the file with the same name as of the folder followed by .go extension
