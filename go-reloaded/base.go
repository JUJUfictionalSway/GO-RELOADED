package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HandleConv(str string) string {
	str = strings.ReplaceAll(str, "( ", "(")
	lines := strings.Fields(str)
	for i := 0; i < len(lines); i++ {
		lineo := lines[i]
		if lineo == "(hex)" || strings.HasPrefix(lineo, "(hex,") {
			if lineo == "(hex)" {
				value, err := strconv.ParseInt(lines[i-1], 16, 64)
				if err != nil {
					fmt.Println("err:::ln18")
				}
				lines[i-1] = strconv.FormatInt(value, 10)
			} else if i+1 < len(lines) {
				numstr := strings.TrimSuffix(lines[i+1], ")")
				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						value, err := strconv.ParseInt(lines[i-j], 16, 64)
						if err != nil {
							fmt.Println("err:::ln28")
						}
						lines[i-j] = strconv.FormatInt(value, 10)
					}
				}
			}
			lines = append(lines[:i], lines[i+1:]...)
			i--
		}
		if lineo == "(bin)" || strings.HasPrefix(lineo, "(bin,") {
			if lineo == "(bin)" {
				value, err := strconv.ParseInt(lines[i-1], 2, 64)
				if err != nil {
					fmt.Println("err:::ln41")
				}
				lines[i-1] = strconv.FormatInt(value, 10)
			} else if i+1 < len(lines) {
				numstr := strings.TrimSuffix(lines[i+1], ")")
				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						value, err := strconv.ParseInt(lines[i-j], 2, 64)
						if err != nil {
							fmt.Println("err:::ln51")
						}
						lines[i-j] = strconv.FormatInt(value, 10)
					}
				}
			}
			lines = append(lines[:i], lines[i+1:]...)
			i--
		}
	}
	return strings.Join(lines, " ")
}
