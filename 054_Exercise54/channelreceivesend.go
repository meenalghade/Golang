package main

func main() {
	/*cs := make(chan<- int)

	go func() {
		cs <- 42           PART 1 : recevier to send-only type chan<- int
	}()
	fmt.Println(<-cs)


	fmt.Print("------------------\n")
	fmt.Printf("cs\t%T\n", cs)*/

	/*cr := make(<-chan int)
	go func() {
		cr <- 42               PART 2 : send to receiver-only type <-chan int
	}()

	fmt.Println(<-cr)
	fmt.Print("------------------\n")
	fmt.Printf("cs\t%T\n", cr)*/

}
