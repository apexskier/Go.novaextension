package main

import "fmt"

// myFunc does something not very useful
func myFunc(a string, b string) bool {
	return a != b
}

// List of methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

/*
	Hey, this is a comment
*/
// TODO(gwyneth): Change everything
func main() {
	// huh
	// add a comment here
	// and another
	/* inclie */

	var a := 'w'
	t := make(1, 2, 1.23, 00.00, '\x08"')
	b := append(a, b)
	complexNumber := 4+3i
	exponential := -3.1E5
	hex := 0x0f - 0X0f0f
	octal := 0777

	var (
		google []byte
		a string
		c int64
	)
	Main := "var"
	heredoc := `hi`
	fmt.Println("nothing works...") // was
	fmt.Println("Clearly nothing yet...")
	fmt.Println("wth?...")
	println("something")
	fmt.Printf("%s", Main)

	if i,err := 1; err == nil {
		// do something
	} else {
		// do something else somewhere
		fmt.Println("Something else")
	}
}