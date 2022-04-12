package main

import (
	"fmt"
	"module_001/auxiliary"
	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Writing in the file main")
	auxiliary.AbcPrint()

	err1 := checkmail.ValidateFormat("rprojetos@gmail.com")
	if err1 != nil {
        fmt.Println("err1", err1)
    }

	err2 := checkmail.ValidateFormat("rprojetos@@gmail.com")
	if err2 != nil {
        fmt.Println("err2", err2)
    }
}
