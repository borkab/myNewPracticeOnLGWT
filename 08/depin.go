package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s", name) the name is getting printed out, but it goes to stdout
	fmt.Fprintf(writer, "Hello, %s", name) //Fprintf() takes a Writer as its argument, but Printf() not
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

/*
Our first round of code was not easy to test
because it wrote data to somewhere we couldn't control.

we refactored the code so we could control where the data was written
by injecting a dependency which allowed us to:

Test our code: If you can't test a function easily, it's usually because of
dependencies hard-wired into a function or global state.
DI will motivate you to inject in a database dependency (via an interface) which
you can then mock out with something you can control in your tests.

Separate our concerns, decoupling where the data goes from how to generate it.

Allow our code to be re-used in different contexts:
 The first "new" context our code can be used in is inside tests.
 But further on if someone wants to try something new with your function
  they can inject their own dependencies.
*/
