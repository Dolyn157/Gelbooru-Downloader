package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func OnlyContainsDigits(str string) bool {

	regPattern := "\\d+"
	result, _ := regexp.MatchString(regPattern, str)
	return result

}

//"https://img3.gelbooru.com/images/53/af/53afdc173f70721211c61baa00159ddf.jpg"

func GetNameFromURI(rawUrl string) string {

	reversed := Reverse(rawUrl)
	cutPoint := strings.Index(reversed, "/")
	prefix := reversed[0:cutPoint]

	revFileName := strings.TrimSuffix(prefix, "/")
	FileName := Reverse(revFileName)
	fmt.Println(FileName)

	return FileName
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
