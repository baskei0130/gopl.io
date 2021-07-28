package bank

type wreq struct {
	amount  int
	succeed chan bool
}

var deposits = make(chan int)    // 入金額を送信
var withdraws = make(chan *wreq) // 引き落とし額を送信
var balances = make(chan int)    // 残高を送信

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	succ := make(chan bool)
	withdraws <- &wreq{amount, succ}
	return <-succ
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdraw := <-withdraws:
			if balance < withdraw.amount {
				withdraw.succeed <- false
			} else {
				balance -= withdraw.amount
				withdraw.succeed <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start monitor goroutine
}
