package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var convert_map = map[rune]rune{
	'A': 'Ⓐ',
	'B': 'Ⓑ',
	'C': 'Ⓒ',
	'D': 'Ⓓ',
	'E': 'Ⓔ',
	'F': 'Ⓕ',
	'G': 'Ⓖ',
	'H': 'Ⓗ',
	'I': 'Ⓘ',
	'J': 'Ⓙ',
	'K': 'Ⓚ',
	'L': 'Ⓛ',
	'M': 'Ⓜ',
	'N': 'Ⓝ',
	'O': 'Ⓞ',
	'P': 'Ⓟ',
	'Q': 'Ⓠ',
	'R': 'Ⓡ',
	'S': 'Ⓢ',
	'T': 'Ⓣ',
	'U': 'Ⓤ',
	'V': 'Ⓥ',
	'W': 'Ⓦ',
	'X': 'Ⓧ',
	'Y': 'Ⓨ',
	'Z': 'Ⓩ',
	'a': 'ⓐ',
	'b': 'ⓑ',
	'c': 'ⓒ',
	'd': 'ⓓ',
	'e': 'ⓔ',
	'f': 'ⓕ',
	'g': 'ⓖ',
	'h': 'ⓗ',
	'i': 'ⓘ',
	'j': 'ⓙ',
	'k': 'ⓚ',
	'l': 'ⓛ',
	'm': 'ⓜ',
	'n': 'ⓝ',
	'o': 'ⓞ',
	'p': 'ⓟ',
	'q': 'ⓠ',
	'r': 'ⓡ',
	's': 'ⓢ',
	't': 'ⓣ',
	'u': 'ⓤ',
	'v': 'ⓥ',
	'w': 'ⓦ',
	'x': 'ⓧ',
	'y': 'ⓨ',
	'z': 'ⓩ',
	'1': '①',
	'2': '②',
	'3': '③',
	'4': '④',
	'5': '⑤',
	'6': '⑥',
	'7': '⑦',
	'8': '⑧',
	'9': '⑨',
	'0': '⓪',
}

func bubblize(r rune) rune {
	if t, ok := convert_map[r]; ok {
		return t
	}
	return r
}

func main() {
	var input string
	info, _ := os.Stdin.Stat()
	if info.Size() > 0 {
		raw, _ := ioutil.ReadAll(os.Stdin)
		input = strings.TrimSpace(string(raw))
	} else {
		input = strings.Join(os.Args[1:], " ")
	}
	bubbles := strings.Map(bubblize, input)
	cmd := exec.Command("pbcopy")
	stdin, _ := cmd.StdinPipe()
	fmt.Fprint(stdin, bubbles)
	stdin.Close()
	cmd.Run()
	fmt.Println(bubbles)
}
