// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start_time := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("%v time: %v\n", s, time.Since(start_time).Nanoseconds())

	start_time_join := time.Now()
	fmt.Printf("%v time: %v\n", strings.Join(os.Args[0:], ""), time.Since(start_time_join).Nanoseconds())
}
