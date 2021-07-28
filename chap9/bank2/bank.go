package bank

var (
	sema    = make(chan struct{}, 1) // balance を保護するバイナリセマフォ
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
