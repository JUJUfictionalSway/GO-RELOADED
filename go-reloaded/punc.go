package main

import "strings"

func punc(str string) string {
	punc := ".,/?;:!_-|"
	var output string
	alph := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' && i+1 < len(str) && strings.Contains(punc, string(str[i+1])) {
			continue
		}
		output += string(str[i])

		if strings.Contains(punc, string(str[i])) && i+1 < len(str) && !strings.Contains(punc, string(str[i+1])) && strings.Contains(alph, string(str[i+1])) && str[i+1] != ' ' {
			output += " "
		}
	}
	return output
}
