package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error

	length := 4
	generatedNumber := generateNumber(length)
	//fmt.Println("Nomor yang tergenerate : " + generatedNumber)
	fmt.Println("Nomor telah tergenerate yuk ditebak :D ")
	input := ""
	for input != generatedNumber {
		input, err = readInput(length)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		validateNumber(generatedNumber, input)
	}

	fmt.Println("tebakan anda benar")
}

// generate 6 different 6 number
func generateNumber(length int) string {

	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})

	randomNumber := ""
	for i := 0; i < length; i++ {
		randomNumber = randomNumber + strconv.Itoa(digits[i])
	}
	return randomNumber
}

// receive inputed number from user
func readInput(length int) (string, error) {

	text := ""
	var err error

	for len(text) != length {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter " + strconv.Itoa(length) + " digit text: ")
		text, err = reader.ReadString('\n')
		text = strings.Trim(text, " \n\t")
		if err != nil {
			return "", err
		}
	}

	return text, nil
}

// validate submitted number
func validateNumber(generatedNumber string, submittedNumber string) {
	//check benar berapa
	submittedSlice := []byte(submittedNumber)
	generatedSlice := []byte(generatedNumber)
	valid := 0
	position := 0
	for i := 0; i < len(submittedNumber); i++ {
		for j := 0; j < len(generatedNumber); j++ {
			if submittedSlice[i] == generatedSlice[j] {
				valid++
				if i == j {
					position++
				}
			}
		}
	}

	fmt.Println("Angka Benar : " + strconv.Itoa(valid) + " Posisi Benar : " + strconv.Itoa(position))
}
