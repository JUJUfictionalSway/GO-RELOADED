package main

import "regexp"

func handlepunchnQ(str string) string {
	str = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(str, " '$1' ")
	return str

}
