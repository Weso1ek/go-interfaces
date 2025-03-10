package main

import "fmt"

//type bot interface {
//	getGreeting() string
//}
//
//type englishBot struct{}
//type spanishBot struct{}

//type logWriter struct{}

type shape interface {
	area() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

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

	// number of empty elements in element
	//bs := make([]byte, 99999)
	//
	//resp.Body.Read(bs)
	//
	//fmt.Println(string(bs))

	//lw := logWriter{}

	//io.Copy(os.Stdout, resp.Body)
	//io.Copy(lw, resp.Body)

	//==============================================================================

	t := triangle{
		height: 10.0,
		base:   30.0,
	}

	s := square{
		sideLength: 10.0,
	}

	fmt.Println(t.area())
	fmt.Println(s.area())
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

//func (lw logWriter) Write(p []byte) (n int, err error) {
//	fmt.Println(string(p))
//	fmt.Println("just read that many bytes", len(p))
//	return len(p), nil
//}

func (t triangle) area() float64 {
	return 0.5 * t.height * t.base
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
