package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numbers := make(chan int, 100)
	squares := make(chan int, 100)
	errors := make(chan error, 100)

	for i := 1; i <= 100; i++ {
		numbers <- i
	}
	close(numbers)
	// =============================================
	var wg sync.WaitGroup

	for num := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			result, err := square(num)
			if err != nil {
				errors <- err
			}
			squares <- result
		}(num)
	}
	wg.Wait()

	// =============================================
	runSelect := true
	for runSelect {
		select {
		case res := <-squares:
			fmt.Println(res)
		case err := <-errors:
			fmt.Println(err)
		default:
			fmt.Println("finished")
			runSelect = false
		}
	}
	fmt.Println("hi")
}

func worker(numbers, results chan int) {
	for number := range numbers {
		results <- number
	}
}

func square(num int) (int, error) {
	time.Sleep(time.Millisecond * 500)
	if num%5 == 0 {
		return 0, fmt.Errorf("some error %d", num)
	}
	return num * num, nil
}
