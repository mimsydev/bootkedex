package main

import "strings"

func cleanInput(text string) []string {
	return strings.Split(strings.Trim(strings.ToLower(text), " \t\n\r"), " ")
	// return make([]string, 0)
}
