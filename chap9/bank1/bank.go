package bank

var deposits = make(chan int) // 入金額を送信
var balances = make(chan int) // 残高を送信

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amout := <-deposits:
			balance += amout
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start monitor goroutine
}
