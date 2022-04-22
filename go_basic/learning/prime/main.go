package main

import (
	"fmt"
)

func putNumber(numberChannel chan<- int) {
	for i := 2; i <= 80; i++ {
		numberChannel <- i
	}
	//close(numberChannel)
}

func judgePrime(numberChannel <-chan int, primeChannel chan<- int) {
	for {
		select {
		case number := <-numberChannel:
			isPrime := true
			for i := 2; i < number; i++ {
				if number%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primeChannel <- number
			}

		}
		//time.Sleep(time.Second)
	}
}

func main() {

	goroutines := 10

	numberChannel := make(chan int, 10)
	primeChannel := make(chan int, 10)
	//judgeExitChannel := make(chan bool, goroutines)

	go putNumber(numberChannel)

	for i := 0; i < goroutines; i++ {
		go judgePrime(numberChannel, primeChannel)
	}
label:
	for {
		select {
		case prime := <-primeChannel:
			fmt.Printf("%d\n", prime)
		default:
			break label
		}
		//time.Sleep(time.Second)
	}

}
