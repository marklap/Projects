/*
**Reverse a String** - Enter a string and the program will reverse it and print it out.
 */
package main

import (
    "bufio"
    "fmt"
    "os"
)

func reverse(text string) string {
    t := []rune(text) // use a rune array to be explicit we're dealing with text
    n := len(t)
    h := n / 2
    for i := 0; i < h; i++ {
        t[i], t[n-i-1] = t[n-i-1], t[i]
    }
    return string(t)
}

func main() {

    fmt.Printf("Enter a text to reverse: ")

    stdin := bufio.NewScanner(os.Stdin)

    stdin.Scan() // get text entered
    if err := stdin.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input", err)
    }

    fmt.Println(reverse(stdin.Text()))

}
