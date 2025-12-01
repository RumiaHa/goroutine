package main

import "fmt"


func inputStream(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func doubler(in chan int, out chan int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func printer(ch chan int) {
	for num := range ch {
		fmt.Println("Результат:", num)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)

	go inputStream(chan1)

	go doubler(chan1, chan2)


	printer(chan2)

}

