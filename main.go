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
	file, err := os.Open("./names.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		cName := capitalise(name)
		print(cName)
	}

	fmt.Println(time.Since(start))
}

func capitalise(name string) string {
	delay(2e3)
	return strings.ToUpper(name)
}

func print(name string) {
	delay(1e3)
	fmt.Println("hello this is ----->>>>  ", name)
}
