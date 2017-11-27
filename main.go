// https://gobyexample.com/string-formatting
// http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
// https://stackoverflow.com/questions/40159137/golang-reading-from-stdin-how-to-detect-special-keys-enter-backspace-etc
// https://github.com/mattn/go-tty/blob/master/tty_unix.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input, err := os.Open("/dev/tty")
	if err != nil {
		fmt.Println("Error opening the file")
	}

	for {
		bin := bufio.NewReader(input)
		fmt.Printf("\033[32mIn:\033[0m ")

		input, _ := bin.ReadByte()

		fmt.Printf("\033[31mOut:\033[0m %d\n", input)
	}

	// Reset
	// \033[0m

	// Basic ANSI colors
	// fmt.Println("\033[30mA")
	// fmt.Println("\033[31mB")
	// fmt.Println("\033[32mC")
	// fmt.Println("\033[33mD")
	// fmt.Println("\033[34mE")
	// fmt.Println("\033[35mF")
	// fmt.Println("\033[36mG")
	// fmt.Println("\033[37mH")

	// 16 Colors
	// fmt.Println("\033[30;1mA")
	// fmt.Println("\033[31;1mB")
	// fmt.Println("\033[32;1mC")
	// fmt.Println("\033[33;1mD")
	// fmt.Println("\033[34;1mE")
	// fmt.Println("\033[35;1mF")
	// fmt.Println("\033[36;1mG")
	// fmt.Println("\033[37;1mH")

	//for i := 0; i < 16; i++ {
	//	for j := 0; j < 16; j++ {
	//		code := i*16 + j
	// 		fmt.Printf("\033[38;5;%dm%4d", code, code)
	// 	}
	// 	fmt.Printf("\n")
	//}

	// First time
	//for {
	//	reader := bufio.NewReaderSize(os.Stdin, 1)
	//	fmt.Printf("\033[32mIn:\033[0m ")
	//	input, _ := reader.ReadByte()

	//	ascii := input

	//	fmt.Printf("\033[31mOut:\033[0m %d\n", ascii)
	//}

}
