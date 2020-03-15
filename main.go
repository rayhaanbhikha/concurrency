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

	names := names("./names.txt")
	capitalisedNames := capitalise(names)
	writeToFile("f-names.txt", capitalisedNames)

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

func capitalise(names <-chan string) <-chan string {
	capitalisedNames := make(chan string)
	go func() {
		for name := range names {
			delay(2e3)
			capitalisedNames <- strings.ToUpper(name)
		}
		close(capitalisedNames)
	}()
	return capitalisedNames
}

func writeToFile(fileName string, capitalisedNames <-chan string) {
	file, err := os.Create(fileName)
	checkErr(err)
	defer file.Close()
	buffWriter := bufio.NewWriter(file)

	for name := range capitalisedNames {
		delay(3e3)
		buffWriter.WriteString(fmt.Sprintf("%s\n", name))
	}
	buffWriter.Flush()
}
