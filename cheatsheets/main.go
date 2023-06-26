// BEGIN WITH
// go mod init booking-app

// packege name
package main

// imoports
import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

	// FOR, IF
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

	// STRING, FOR

	slogans := [3]string{"Hi fi", "My coffe", "Friend sugar"}

	for _, slogan := range slogans {
		firstWord := strings.Fields(slogan)[0]
		strings.Contains(firstWord, "Hi")
		fmt.Println("Slogan from slogans:", firstWord)
	}

	// TIME
	start_time_join := time.Now()
	fmt.Printf("result: %v, time spend: %v\n", strings.Join(slogans[0:], " "), time.Since(start_time_join).Nanoseconds())
	// time.Sleep(10 * time.Seconds)

	// SWITCH

	city := "London"

	switch city {
	case "London", "Sidney":
	// some code
	case "Denver":
	// some code
	default:
		// some code
	}

	// PACKEGES GLOBAL VAR
	helpMenu := "Help menu"
	helpVersion := 10
	helpApp(helpMenu, helpVersion)

	fmt.Printf("IT's a global thing - %v\n", GlobalVariable)

	// MAPS
	var listOfMaps = make([]map[string]string, 0)
	var notEmptyMap = make(map[string]string)
	notEmptyMap["one"] = "1"
	notEmptyMap["two"] = "2"
	// convert int 3 to string
	notEmptyMap[strconv.FormatUint(uint64(3), 10)] = "three"

	listOfMaps = append(listOfMaps, notEmptyMap)
	// CONCURENCY SYMPLE

	waitSeconds(10)
	workStruct()

}
func waitSeconds(seconds int) {
	fmt.Printf("PLease wait %v sec....\n", seconds)
	time.Sleep(10 * time.Second)
	fmt.Println("Thank you for your waiting!")

}
func workStruct() {
	//STUCT
	type UserData struct {
		firstName string
		lastName  string
		age       int
	}

	var listOfStructs = make([]UserData, 0)

	var userData_1 = UserData{
		firstName: "Bob",
		lastName:  "Korel",
		age:       19,
	}

	var UserData_2 = UserData{
		firstName: "Jhon",
		lastName:  "Rohl",
		age:       56,
	}

	listOfStructs = append(listOfStructs, userData_1)
	listOfStructs = append(listOfStructs, UserData_2)

	for index, listStruct := range listOfStructs {
		fmt.Printf("%v name is: %v", index, listStruct.firstName)
	}

}
