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
	cNames := make(chan string)

	go func() {
		defer close(cNames)
		for {
			name, ok := <-names
			if !ok {
				return
			}
			cname := capitalise(name)
			fmt.Println("capitalised -> ", cname)
			cNames <- cname

		}
	}()

	go func() {
		for {
			name, ok := <-cNames
			if !ok {
				quit <- 1
			}
			print(name)
		}
	}()

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

func capitalise(name string) string {
	delay(2e3)
	return strings.ToUpper(name)
}

func print(name string) {
	delay(3e3)
	fmt.Println("hello this is ----->>>>  ", name)
}
