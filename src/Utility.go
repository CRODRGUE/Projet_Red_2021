package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func isAlpha(s string) bool {
	for _, char := range s {
		if !(rune(char) >= 97 && rune(char) <= 122) && !(rune(char) >= 65 && rune(char) <= 90) && !(rune(char) >= 48 && rune(char) <= 57) {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, char := range s {
		if !(rune(char) >= 97 && rune(char) <= 122) {
			return false
		}
	}
	return true
}

func toLower(s string) string {
	res := ""
	for _, char := range s {
		if rune(char) >= 65 && rune(char) <= 90 {
			new_rune := rune(char) + 32
			res += string(new_rune)
		} else {
			res += string(char)
		}
	}
	return res
}

func toUpper(s string) string {
	res := ""
	for _, char := range s {
		if rune(char) >= 97 && rune(char) <= 122 {
			new_rune := rune(char) - 32
			res += string(new_rune)
		} else {
			res += string(char)
		}
	}
	return res
}

func isUpper(s string) bool {
	for _, char := range s {
		if !(rune(char) >= 65 && rune(char) <= 90) {
			return false
		}
	}
	return true
}

func Capitalize(s string) string {
	str := ""
	for index, char := range s {
		new_char := ""
		if index == 0 && isLower(string(char)) {
			new_char = toUpper(string(char))
		} else if isLower(string(char)) && !(isAlpha(string(s[index-1]))) {
			new_char = toUpper(string(char))
		} else {
			if index != 0 && isUpper(string(char)) && isAlpha(string(s[index-1])) {
				new_char = toLower(string(char))
			} else {
				new_char = string(char)
			}
		}
		str += new_char
	}
	return str
}

//Affiche les choix des diffrents menus
func PrintMenu(tab []string) {
	for nb, words := range tab {
		fmt.Print((nb + 1), "-", words, "\n")
	}
}

//Supprime le terminal
func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func QuiSontIls() {
	fmt.Println(Blue, "A", Purple, "B", Red, "B", Yellow, "A", Reset)
	fmt.Print(Blue, "Ste", Purple, "ven ", Red, "Spie", Yellow, "lberg", Reset, "\n")
}

// génère un nombre aléatoire entre min et max
func RandomNbr(min, max int) int {
	time.Sleep(3)
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}
