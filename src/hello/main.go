package main

import (
	"fmt"
	"os"
	"time"
	"errors"
	"hello/model"
)

func main() {

	jumperList := model.GetList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
	//declare variables with type inference
	// type inference <variable-name> := <some-value>
	// manual type declaration
	//var <variable-name> <data-type>
	//<variable name> = <some-value>
	// common data types: int, string, bool, []string

	args := os.Args
	var message string
	message = "Hello from go. This is Yi\n "
	if len(args) > 1 { //array with program arguments, starting with the name of the executable and followed by user-supplied arguments
		fmt.Println(args[1])
	} else {
		fmt.Println(message)
	}

	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(greeting)

	for i := 0; i < 5; i++ {
		fmt.Println(i);
	}

	j := 0
	isLessThanFive := true
	for isLessThanFive {
		if j >= 4 {
			isLessThanFive = false
		}
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		if k > 3 {
			break
		}
		fmt.Println(k)
		k++
	}

  //array
	// var langs [3]string
	// langs[0] = "Go"
	// langs[1] = "Ruby"
	// langs[2] = "JavaScript"
	// fmt.Println(langs)

	//slices
	// var langs []string
	// langs = append(langs, "Go")
	// langs = append(langs, "Ruby")
	// langs = append(langs, "JavaScript")
	// langs = append(langs, "blahblah")
	// fmt.Println(langs)

	//langs := []string{"Go", "Ruby", "JS"}
	//fmt.Println(langs[1])

	langs := getLangs()
	for _, element := range langs {
		fmt.Println(element)
	}

}

func getLangs() []string {
	return []string{"Go", "Ruby", "JS"}
}

//accordin to docs for time package it returns type int
// fucntion signature: for every function in Go it start with type of argument and type of whats in return
// in Go, we communicate error via a separate return value
func getGreeting(hour int) (string, error) {
	var message string
	if hour < 7 {
		err :=errors.New("Too early")
		return message, err
	}
	if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}
