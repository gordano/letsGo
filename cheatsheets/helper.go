package main

import "fmt"

var GlobalVariable = "OH Go, dear I need beer."

func helpApp(helpMenu string, helpVersion int) (bool, bool, string) {

	fmt.Printf("This is app: %v\n with version app: %v\n", helpMenu, helpVersion)

	return true, false, "yes"
}
