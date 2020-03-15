package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	start := time.Now()
	quit := make(chan int)
	names := names("./names.txt")

	capitalisedNames := cCapitalise(names)
	cPrint(capitalisedNames, quit)

	<-quit

	fmt.Println(time.Since(start))
}

func names(fileName string) <-chan string {
	file, err := os.Open(fileName)
	checkErr(err)
	names := make(chan string)
	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			names <- scanner.Text()
		}
		close(names)
	}()
	return names
}

func cCapitalise(names <-chan string) <-chan string {
	capitalisedNames := make(chan string)
	go func() {
		for name := range names {
			cname := capitalise(name)
			fmt.Println("capitalised -> ", cname)
			capitalisedNames <- cname
		}
		close(capitalisedNames)
	}()
	return capitalisedNames
}

func cPrint(cNames <-chan string, quit chan<- int) {
	go func() {
		for name := range cNames {
			print(name)
		}
		quit <- 1
	}()
}

func capitalise(name string) string {
	delay(2e3)
	return strings.ToUpper(name)
}

func print(name string) {
	delay(3e3)
	fmt.Println("hello this is ----->>>>  ", name)
}
