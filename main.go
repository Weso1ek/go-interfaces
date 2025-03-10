package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	//eb := englishBot{}
	//sb := spanishBot{}
	//
	//printGreeting(eb)
	//printGreeting(sb)

	//===============================================================================

	// HTTP APPLICATION
	//resp, err := http.Get("http://www.google.com")
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	os.Exit(1)
	//}
	//
	//// number of empty elements in element
	//bs := make([]byte, 99999)
	//
	//resp.Body.Read(bs)
	//
	//fmt.Println(string(bs))
	//
	//lw := logWriter{}
	//
	////io.Copy(os.Stdout, resp.Body)
	//io.Copy(lw, resp.Body)

	//==============================================================================

	t := triangle{
		height: 10.0,
		base:   30.0,
	}

	s := square{
		sideLength: 10.0,
	}

	printAres(t)
	printAres(s)

	// FILE EXAMPLE =============================================================================

	fileName := os.Args[1]

	file, err := os.Open(fileName) // For read access.
	if err != nil {
		fmt.Println("Error opening file:", fileName)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

	fmt.Println(file)
}

func (lw logWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("just read that many bytes", len(p))
	return len(p), nil
}
