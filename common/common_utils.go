package common

import "strconv"

func IsNumeric(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func IsBoolean(s string) bool {
	if s == "true" || s == "false" {
		return true
	}

	return false
}

func UniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us
}

func ConvertStringArrToIntArr(strings []string) []int {
	var res []int
	for _, data := range strings {
		resInt, _ := strconv.Atoi(data)
		res = append(res, resInt)
	}

	return res
}

func GetLastNCharacters(string string, numberOfCharacter int) string {
	return string[len(string)-numberOfCharacter:]
}

func RemoveZeroInfrontOfInteger(value string) string {
	var temp string
	if value[:0] == "0" {
		temp = value[:1]
	} else {
		temp = value
	}

	return temp
}
