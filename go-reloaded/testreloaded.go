package main

// func HandleConv(str string) string {
// 	str = strings.ReplaceAll(str, "( ", "(")
// 	lines := strings.Fields(str)
// 	for i := 0; i < len(lines); i++ {
// 		lineo := lines[i]
// 		if lineo == "(hex)" || strings.HasPrefix(lineo, "(hex,") {
// 			if lineo == "(hex)" {
// 				value, err := strconv.ParseInt(lines[i-1], 16, 64)
// 				if err != nil {
// 					fmt.Println("err:::ln18")
// 				}
// 				lines[i-1] = strconv.FormatInt(value, 10)
// 			} else if i+1 < len(lines) {
// 				numstr := strings.TrimSuffix(lines[i+1], ")")
// 				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
// 				for j := 1; j <= num; j++ {
// 					if i-j >= 0 {
// 						value, err := strconv.ParseInt(lines[i-j], 16, 64)
// 						if err != nil {
// 							fmt.Println("err:::ln28")
// 						}
// 						lines[i-j] = strconv.FormatInt(value, 10)
// 					}
// 				}
// 			}
// 			lines = append(lines[:i], lines[i+1:]...)
// 			i--
// 		}
// 		if lineo == "(bin)" || strings.HasPrefix(lineo, "(bin,") {
// 			if lineo == "(bin)" {
// 				value, err := strconv.ParseInt(lines[i-1], 2, 64)
// 				if err != nil {
// 					fmt.Println("err:::ln41")
// 				}
// 				lines[i-1] = strconv.FormatInt(value, 10)
// 			} else if i+1 < len(lines) {
// 				numstr := strings.TrimSuffix(lines[i+1], ")")
// 				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
// 				for j := 1; j <= num; j++ {
// 					if i-j >= 0 {
// 						value, err := strconv.ParseInt(lines[i-j], 2, 64)
// 						if err != nil {
// 							fmt.Println("err:::ln51")
// 						}
// 						lines[i-j] = strconv.FormatInt(value, 10)
// 					}
// 				}
// 			}
// 			lines = append(lines[:i], lines[i+1:]...)
// 			i--
// 		}
// 	}
// 	return strings.Join(lines, " ")
// }
// func handleCases(str string) string {
// 	lines := strings.Fields(str)
// 	for i, lineo := range lines {
// 		if lineo == "(up)" || strings.HasPrefix(lineo, "(up,") {
// 			if lineo == "(up)" {
// 				lines[i-1] = strings.ToUpper(lines[i-1])
// 			} else if i+1 < len(lines) {
// 				numstr := strings.TrimSuffix(lines[i+1], ")")
// 				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
// 				for j := 1; j <= num; j++ {
// 					if i-j >= 0 {
// 						lines[i-j] = strings.ToUpper(lines[i-j])
// 					}
// 				}
// 			}
// 			lines = append(lines[:i], lines[i+1:]...)
// 			i--
// 		}
// 		if lineo == "(low)" || lineo == "(low," {
// 			if lineo == "(low)" {
// 				lines[i-1] = strings.ToLower(lines[i-1])
// 			} else if i+1 < len(lines) {
// 				numstr := strings.TrimSuffix(lines[i+1], " ")
// 				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
// 				for j := 1; j <= num; j++ {
// 					if i-j >= 0 {
// 						lines[i-j] = strings.ToLower(lines[i-j])
// 					}
// 				}
// 			}
// 			lines = append(lines[:i], lines[i+1:]...)
// 			i--
// 		}
// 		if lineo == "(cap)" || strings.HasPrefix(lineo, "(cap,") {
// 			if lineo == "(cap)" {
// 				lines[i-1] = strings.Title(lines[i-1])
// 			} else if i+1 < len(lines) {
// 				numstr := strings.TrimSuffix(lines[i+1], ")")
// 				num, _ := strconv.Atoi(strings.TrimSpace(numstr))
// 				for j := 1; j <= num; j++ {
// 					if i-j >= 0 {
// 						lines[i-j] = strings.Title(lines[i-j])
// 					}
// 				}
// 			}
// 			lines = append(lines[:i], lines[i+1:]...)
// 			i--
// 		}
// 	}
// 	return strings.Join(lines, " ")
// }
// func handleAtotheAn(str string) string {
// 	lines := strings.Fields(str)
// 	for i, lineo := range lines {
// 		if lineo == "a" || lineo == "A" {
// 			j := i + 1
// 			for j < len(lines) {
// 				if lines[j] == "'" || lines[j] == "\"" {
// 					j++
// 				} else {
// 					break
// 				}
// 			}
// 			if j < len(lines) {
// 				f := strings.ToLower(string(lines[j][0]))
// 				if f == "a" || f == "e" || f == "i" ||
// 					f == "o" || f == "u" || f == "h" {
// 					if lineo == "a" {
// 						lines[i] = "an"
// 					} else {
// 						lines[i] = "An"
// 					}
// 				}
// 			}
// 		}
// 		if lineo == "an" || lineo == "An" {
// 			j := i + 1
// 			for j < len(lines) {
// 				if lines[j] == "'" || lines[j] == "\"" {
// 					j++
// 				} else {
// 					break
// 				}
// 			}
// 			if j < len(lines) {
// 				f := strings.ToLower(string(lines[j][0]))
// 				if f != "a" && f != "e" && f != "i" &&
// 					f != "o" && f != "u" {
// 					if lineo == "an" {
// 						lines[i] = "an"
// 					} else {
// 						lines[i] = "An"
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return strings.Join(lines, " ")
// }
// func handlepunchnQ(str string) string {
// 	str = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(str, " '$1' ")
// 	return str

// }
// func punc(str string) string {
// 	punc := ".,\"/?;:'!_-|"
// 	var output string
// 	alph := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == ' ' && i+1 < len(str) && strings.Contains(punc, string(str[i+1])) {
// 			continue
// 		}
// 		output += string(str[i])

// 		if strings.Contains(punc, string(str[i])) && i+1 < len(str) && !strings.Contains(punc, string(str[i+1])) && strings.Contains(alph, string(str[i+1])) && str[i+1] != ' ' {
// 			output += " "
// 		}
// 	}
// 	return output
// }
