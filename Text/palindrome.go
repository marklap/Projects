/*
**Check if Palindrome** - Checks if the string entered by the user is a palindrome. That is that it reads the same forwards as backwards like “racecar”
 */
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
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

    re := regexp.MustCompile(`\PL`) // we're looking for anything that's not a unicode Letter

    text := re.ReplaceAllString(strings.ToLower(stdin.Text()), "")
    rtext := reverse(text)
    if text == rtext {
        fmt.Println("It IS a palindrome!")
    } else {
        fmt.Println("Bummer, no palindrome.")
    }

}
