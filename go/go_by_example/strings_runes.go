package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	const s2 = "hmmm"

	// since a string is of type []byte, len(s) is the number of bytes
	// "สวัสดี" 4 characters, but is encoded in UTF-8 so its actually 18 bytes
	// strings are NOT null-terminated!
	fmt.Println("len('สวัสดี'):", len(s)) // 18 bytes

	// prints out the hex value of each byte -> NOT each character!
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	// number of 'runes' is number of unicode code-points
	// in Thai, som e characters are made of multiple unicode code points,
	// so this isn't nesicarily the number of 'characters' that appear in the final text
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	// ranges work special for strings: idx is the offset of the runeValue
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	// manually looping over characters
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}
	fmt.Println()

	fmt.Println("len('hmmm'):", len(s2)) // 4 bytes
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%x ", s2[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s2))
	for idx, runeValue := range s2 {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
