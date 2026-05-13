package main

import "strings"

func handleAtotheAn(str string) string {
	lines := strings.Fields(str)
	for i, lineo := range lines {
		if lineo == "a" || lineo == "A" {
			j := i + 1
			for j < len(lines) {
				if lines[j] == "'" || lines[j] == "\"" {
					j++
				} else {
					break
				}
			}
			if j < len(lines) {
				f := strings.ToLower(string(lines[j][0]))
				if f == "a" || f == "e" || f == "i" ||
					f == "o" || f == "u" || f == "h" {
					if lineo == "a" {
						lines[i] = "an"
					} else {
						lines[i] = "An"
					}
				}
			}
		}
		if lineo == "an" || lineo == "An" {
			j := i + 1
			for j < len(lines) {
				if lines[j] == "'" || lines[j] == "\"" {
					j++
				} else {
					break
				}
			}
			if j < len(lines) {
				f := strings.ToLower(string(lines[j][0]))
				if f != "a" && f != "e" && f != "i" &&
					f != "o" && f != "u" {
					if lineo == "an" {
						lines[i] = "an"
					} else {
						lines[i] = "An"
					}
				}
			}
		}
	}
	return strings.Join(lines, " ")
}
