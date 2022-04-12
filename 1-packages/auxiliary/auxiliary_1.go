package auxiliary

import "fmt"

// Write a mesage on the shell
func AbcPrint(){
	fmt.Println("Writing on the package auxiliary")
	abcPrint()
}

// Notes:

// Function started with letter Maíuscula can be
// Used by other packages

// Function initiated with minuscule letter only
// can be used within the package itself

// It is good practice to comment on the function, 
// describing what it does.