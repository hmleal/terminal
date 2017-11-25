// https://gobyexample.com/string-formatting
// http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    // Reset
    // \033[0m

    // Basic ANSI colors
    fmt.Println("\033[30mA")
    fmt.Println("\033[31mB")
    fmt.Println("\033[32mC")
    fmt.Println("\033[33mD")
    fmt.Println("\033[34mE")
    fmt.Println("\033[35mF")
    fmt.Println("\033[36mG")
    fmt.Println("\033[37mH")

    // 16 Colors
    fmt.Println("\033[30;1mA")
    fmt.Println("\033[31;1mB")
    fmt.Println("\033[32;1mC")
    fmt.Println("\033[33;1mD")
    fmt.Println("\033[34;1mE")
    fmt.Println("\033[35;1mF")
    fmt.Println("\033[36;1mG")
    fmt.Println("\033[37;1mH")

    for i := 0; i < 16; i++ {
        for j := 0; j < 16; j++ {
            code := i * 16 + j
            fmt.Printf("\033[38;5;%dm%4d", code, code)
        }
        fmt.Printf("\n")
    }

    for {
        reader := bufio.NewReader(os.Stdin)
        char, _ := reader.ReadByte()

        if char == 10 {
            fmt.Printf(">>> %d", char)
        }
    }

}
