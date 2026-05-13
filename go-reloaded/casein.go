package main

import (
	"strconv"
	"strings"
)

func handleCases(str string) string {
	lines := strings.Fields(str)
	for i, lineo := range lines {
		if lineo == "(up)" || strings.HasPrefix(lineo, "(up,") {
			if lineo == "(up)" {
				lines[i-1] = strings.ToUpper(lines[i-1])
			} else if i+1 < len(lines) {
				numstr := strings.TrimSuffix(lines[i+1], ")")
				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						lines[i-j] = strings.ToUpper(lines[i-j])
					}
				}
			}
			lines = append(lines[:i], lines[i+2:]...)
			i--
		}
		if lineo == "(low)" || lineo == "(low," {
			if lineo == "(low)" {
				lines[i-1] = strings.ToLower(lines[i-1])
			} else if i+1 < len(lines) {
				numstr := strings.TrimSuffix(lines[i+1], " ")
				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						lines[i-j] = strings.ToLower(lines[i-j])
					}
				}
			}
			lines = append(lines[:i], lines[i+2:]...)
			i--
		}
		if lineo == "(cap)" || strings.HasPrefix(lineo, "(cap,") {
			if lineo == "(cap)" {
				lines[i-1] = strings.Title(lines[i-1])
			} else if i+1 < len(lines) {
				numstr := strings.TrimSuffix(lines[i+1], ")")
				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						lines[i-j] = strings.Title(lines[i-j])
					}
				}
			}
			lines = append(lines[:i], lines[i+2:]...)
			i--
		}
	}
	return strings.Join(lines, " ")
}
