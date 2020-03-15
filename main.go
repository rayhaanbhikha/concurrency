package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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
	c1 := capitalise(names)
	c2 := capitalise(names)
	c3 := capitalise(names)
	c4 := capitalise(names)
	c5 := capitalise(names)

	writeToFile("f-names.txt", merge(c1, c2, c3, c4, c5))

	fmt.Println(time.Since(start))
}

func merge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
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
			fmt.Println("capitalise ", name)
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
	var wg sync.WaitGroup

	for name := range capitalisedNames {
		wg.Add(1)
		go func(name string, buffWriter *bufio.Writer) {
			defer wg.Done()
			fmt.Println("writing ", name)
			delay(3e3)
			buffWriter.WriteString(fmt.Sprintf("%s\n", name))
		}(name, buffWriter)
	}
	wg.Wait()
	buffWriter.Flush()
}
