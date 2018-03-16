package main

import (
	"fmt"
	"time"
	"errors"
)

//create MyError,this is an error implemeation that includes a time and message.
type MyError struct {
	When time.Time
	Whate string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.Whate)
}

func oop() error {
	return MyError {
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oop(); err != nil {
		fmt.Println(err)
	}
	err := errors.New("This is a new error")
	if err != nil {
		fmt.Println(err)
	}

	const name, id = "ycl", 27
	err = fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Println(err)
	}
	err = fmt.Errorf("against");
	fmt.Println(err)
}