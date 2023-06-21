// BEGIN WITH
// go mod init booking-app

// packege name
package main

// imoports
import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	const thisIsConstantA uint = 30 // -30 error
	const thisIsConstantB int = -50

	var aVariableA string = "Hey text"
	aVariableB := "another text"
	var varGo string // string, integer, map, array, boolean
	varGo = "Go"

	fmt.Println("Are you", varGo, "?")
	fmt.Printf("Our constants:\n 1-%v\n 2-%v\nOur variables:\n 1-%v\n 2-%v\n",
		thisIsConstantA, thisIsConstantB, aVariableA, aVariableB)

	fmt.Println("What's your name?")
	var firstName string
	fmt.Scan(&firstName) // use Pointer to object variable
	fmt.Println("Niceto meet you", firstName, "!")

	fmt.Println("For operations pls use SAME TYPE, like", thisIsConstantB-10)

	ohArrrayA := [10]int{1, 2, 3, 4} // fix size array
	ohArrrayB := []int{}             // flex size array or Slice
	// len(ohArrrayA)
	// ohArrrayA[0] = 1
	ohArrrayB = append(ohArrrayB, 2)

	fmt.Println("Array fixed", ohArrrayA[0])
	fmt.Println("Array flex", ohArrrayB[0])

	for {
		randOmh := rand.Intn(10)

		if randOmh != 4 {
			fmt.Println("Continue", randOmh)
			continue
		} else {
			fmt.Println("Oh God, stop pls")
			break
		}
	}

	slogans := [3]string{"Hi fi", "My coffe", "Friend sugar"}

	for _, slogan := range slogans {
		firstWord := strings.Fields(slogan)[0]

		fmt.Println("Slogan from slogans:", firstWord)
	}

}
