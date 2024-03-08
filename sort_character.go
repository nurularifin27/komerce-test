package main

import (
	"fmt"
  "strings"
)

func isVowel(c string) bool {
  return strings.Contains("aeiuo",c)
}

func sortChar(c string, arrStr []string) []string {
  for i, s := range arrStr {
    if strings.Contains(s,c){
      arrStr[i] = arrStr[i] + c
      return arrStr
    }
  }
  arrStr = append(arrStr,c)
  return arrStr
}

func main() {
	str := "Sample Case" // input string
  vowels := make([]string,0)
  consonants := make([]string,0)
  strWithoutWhiteSpace := strings.ReplaceAll(str," ","")
  for _, s := range strings.ToLower(strWithoutWhiteSpace){
    if isVowel(string(s)){
      vowels = sortChar(string(s),vowels)
    }else{
      consonants = sortChar(string(s),consonants)
    }
  }
  fmt.Printf("Vowels : %s \n", strings.Join(vowels,""))
  fmt.Printf("Consonants : %s \n", strings.Join(consonants,""))
}
