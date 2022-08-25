package execises17

import (
	"fmt"
	"time"
)

func Hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func MyHello(chantuto chan bool) {

	fmt.Println("hello go routine is going to sleep")

	time.Sleep(6 * time.Second)

	fmt.Println("now go routine has awoken up and is about pass value to the channel chantuto")

	time.Sleep(6 * time.Second)
	chantuto <- true
}

func CalcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func CalcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func Mycount(numberschan chan int) {

	for i := 0; i < 10; i++ {

		numberschan <- i
	}

	close(numberschan)
}

func Write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
