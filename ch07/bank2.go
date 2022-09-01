package main

var (
	sema    = make(chan struct{}, 1)
	balance int
)

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {

}
