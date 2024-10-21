package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func hex(hexString string) int64 {
	hexInt, err := strconv.ParseInt(hexString, 16, 32)
	if err != nil {
		fmt.Println("Error:", err)
		// return
	}
	return hexInt
}
func bin(binString string) int64 {
	decimalInt, err := strconv.ParseInt(binString, 2, 32)
	if err != nil {
		fmt.Println("Error:", err)
		// return
	}
	return decimalInt
}
func Notmain() {
	/*
		///////////////////////////////////
				get file data
		\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input filename> <output filename>")
		return
	}
	inputFile := os.Args[1]
	filePath := os.Args[2]
	txt, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Print(err)
		return
	}
	strtxt := string(txt)
	/*
	   /////////////////////////////////
	   		string to array
	   \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/
	words := strings.Fields(strtxt)
	/*
	   /////////////////////////////////
	   	  	  search for keyword
	   \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/
	for i, j := range words {
		if i > 0 {
			switch j {
			case "(hex)":
				new := strconv.Itoa(int(hex(words[i-1])))
				words[i-1] = new
				words = append(words[:i], words[i+1:]...)
			case "(bin)":
				new := strconv.Itoa(int(bin(words[i-1])))
				words[i-1] = new
				words = append(words[:i], words[i+1:]...)
			case "(up)":
				new := strings.ToUpper(words[i-1])
				words[i-1] = new
				words = append(words[:i], words[i+1:]...)
			case "(low)":
				new := strings.ToLower(words[i-1])
				words[i-1] = new
				words = append(words[:i], words[i+1:]...)
			case "(cap)":
				new := strings.Title(words[i-1])
				words[i-1] = new
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	/*
	   //////////////////////////////////////
	   keyword with num
	   \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/
	for i, j := range words {
		if j == "(up," {
			n := words[i+1][0]
			words[i+1] = string(n)
			num, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			for k := num; k >= 1; k-- {
				new := strings.ToUpper(words[i-k])
				words[i-k] = new
			}
			words = append(words[:i], words[i+2:]...)
		}
		if j == "(low," {
			n := words[i+1][0]
			words[i+1] = string(n)
			num, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			for k := num; k >= 1; k-- {
				new := strings.ToLower(words[i-k])
				words[i-k] = new
			}
			words = append(words[:i], words[i+2:]...)
		}
		if j == "(cap," {
			n := words[i+1][0]
			words[i+1] = string(n)
			num, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			for k := num; k >= 1; k-- {
				new := strings.Title(words[i-k])
				words[i-k] = new
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	// amora asked tala: "noor was here?????
	/*
	   //////////////////////////////////////
	   a ---->  an
	   \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/
	for i := range words {
		if len(words[i]) > 1 {
			switch words[i][0] {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'h', 'H':
				if i > 0 {
					switch words[i-1] {
					case "a":
						words[i-1] = "an"
					case "A":
						words[i-1] = "An"
					}
				}
			}
		}
	}
	// apple   a elephant   a kkajaj
	/*
	   //////////////////////////////////////
	   correct punctatuion
	   \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/
	for j, w := range words {
		for i, o := range w {
			if unicode.IsPunct(o) && i == 0 && o != '"' && o != 39 {
				if j > 0 {
					words[j-1] += string(o)
					words[j] = words[j][1:]
				}
			}
		}
	}
	/*
	   //////////////////////////////////////
	   fix the quote
	   \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/
	count := 0
	for i, r := range words {
		if r == "'" {
			count++
			if count == 1 {
				new := "'" + words[i+1]
				words[i+1] = new
				words = append(words[:i], words[i+1:]...)
			}
			if count == 2 {
				new := words[len(words)-2] + "'"
				words[len(words)-2] = new
				words = append(words[:len(words)-1])
			}
		}
	}
	/*
	   ///////////////////////////////////////
	   write the modified text in the new file
	   \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	*/
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := strings.Join(words, " ")
	modifiedStr := strings.ReplaceAll(result, "  ", " ")
	_, err = io.WriteString(file, modifiedStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}
