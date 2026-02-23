package utils

import (
	"strings"
	"regexp"
	"fmt"
	"strconv"
)

func capitalizeFirstLetter(word string) string{
	if len(word) == 0 {
		return ""
	}
	capitalizeLetter := strings.ToUpper(word[:1]) + word[1:]
	return capitalizeLetter
}

func capitalizeAllLettters(word string) string {
	capitalizeWord := strings.ToUpper(word)
	return capitalizeWord
}

func binToDecimal(word string) string {
	output, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return fmt.Sprint(output)
}

func hexToDecimal(word string) string {
	output, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return fmt.Sprint(output)
}

func changeLetterToLowerCase(word string) string {
	lowerCaseLetter := strings.ToLower(word)
	return lowerCaseLetter
}

func checkForVowels(word string) bool {
	vowels := "aeiou"
	arrFormat := strings.Split(word, "")
	if len(word) > 2 {
	if strings.Contains(vowels, string(arrFormat[0])) {
		return true
	}
	}
	return false
}

func fixExtraSpace(sentence string) string {
	re := regexp.MustCompile(`\s+([.,!?:;])`)
    sentence = re.ReplaceAllString(sentence, "$1")
	reAfter := regexp.MustCompile(`([.,!?:;])([^\s.,!?:;])`)
    sentence = reAfter.ReplaceAllString(sentence, "$1 $2")
    reSpaces := regexp.MustCompile(`\s+`)
    sentence = reSpaces.ReplaceAllString(sentence, " ")
    return strings.TrimSpace(sentence)
}

func LoopText(file_content string) string {
	var capWords []string
	regexVersion := regexp.MustCompile(`\([^)]+\)|[a-zA-Z0-9']+|[.,!?:;]`)
	finalArray := regexVersion.FindAllString(file_content, -1)
	for _, word:= range finalArray {
		capWords = append(capWords, word)
	}
	for index, comp:= range capWords {
		switch comp {
		case "(cap)": 
			if index > 0 {
				capWords[index] = ""
				capWords[index-1] = capitalizeFirstLetter(capWords[index-1])				
			}
		case "(up)": 
			if index > 0 {
				capWords[index] = ""
				capWords[index-1] = capitalizeAllLettters(capWords[index-1])				
			}
		case "(low)":
			if index > 0 {
				capWords[index] = ""
				capWords[index-1] = changeLetterToLowerCase(capWords[index-1])
			}
		case "(bin)": 
		if index > 0 {
			capWords[index] = ""
			capWords[index -1] = binToDecimal(capWords[index-1])
		}
		case "(hex)":
		if index > 0 {
			capWords[index] = ""
			capWords[index -1] = hexToDecimal(capWords[index-1])
		}
		}
		if checkForVowels(comp) {
			if capWords[index-1] == "a" {
				fmt.Println(capWords[index-1])
				capWords[index-1] = "an"
			}
		}
		if strings.HasPrefix(comp, "(cap,") {
			numStr := strings.TrimSuffix(strings.TrimPrefix(comp, "(cap, "), ")")
			n, _ := strconv.Atoi(numStr)
			capWords[index] = ""
			count := 0
				for i := index - 1; i >= 0 && count < n; i-- {
					capWords[i] = capitalizeAllLettters(capWords[i])
					count++
			}
		}
		if strings.HasPrefix(comp, "(low,") {
			numStr := strings.TrimSuffix(strings.TrimPrefix(comp, "(low, "), ")")
			n, _ := strconv.Atoi(numStr)
			capWords[index] = ""
			count := 0
				for i := index - 1; i >= 0 && count < n; i-- {
					capWords[i] = changeLetterToLowerCase(capWords[i])
					count++
			}
		}
		
	}
	tempString := strings.Join(capWords, " ")
	cleanedSlice := strings.Fields(tempString)
	finalString := strings.Join(cleanedSlice, " ")
	return fixExtraSpace(finalString)
}