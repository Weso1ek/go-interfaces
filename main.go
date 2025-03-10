package main

import (
	"fmt"
	"net/http"
	"os"
)

//type bot interface {
//	getGreeting() string
//}
//
//type englishBot struct{}
//type spanishBot struct{}

func main() {
	//eb := englishBot{}
	//sb := spanishBot{}
	//
	//printGreeting(eb)
	//printGreeting(sb)

	// HTTP APPLICATION
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// number of empty elements in element
	bs := make([]byte, 99999)

	resp.Body.Read(bs)

	fmt.Println(string(bs))
}

//func printGreeting(b bot) {
//	fmt.Println(b.getGreeting())
//}
//
//func (eb englishBot) getGreeting() string {
//	return "Hello, World!"
//}
//
//func (sb spanishBot) getGreeting() string {
//	return "Holla!"
//}
