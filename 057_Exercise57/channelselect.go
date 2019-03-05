package main

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//recieve
	receive(eve, odd, quit)
}

func send(e, o, q chan<- int) {

}

func receive(e, o, q <-chan int) {

}
