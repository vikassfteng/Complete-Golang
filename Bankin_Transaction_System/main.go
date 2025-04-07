package main

import (
	"fmt"
	"os"
	"sync"
)

type Account struct {
	ID      int
	Balance float64
	mu      sync.Mutex
}

type Transaction struct {
	Type      string
	Amount    float64
	AccountID int
}

var accounts = map[int]*Account{}
var fileMu sync.Mutex

func (a *Account) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Balance >= amount {
		fmt.Printf("Withdrawing %.2f from account %d\n", amount, a.ID)
		a.Balance -= amount
		fmt.Printf("New balance: %.2f\n", a.Balance)
		return true
	}
	fmt.Printf("Insufficient funds for account %d. Current balance: %.2f cannot withdraw %.2f\n", a.ID, a.Balance, amount)
	return false
}

func porcesTransacton(transactions <-chan Transaction, wg *sync.WaitGroup) {

	defer wg.Done()
	for tx := range transactions {
		acc, exists := accounts[tx.AccountID]
		if !exists {
			fmt.Printf("Account %d does not exist\n", tx.AccountID)
			continue
		}
		success := false
		if tx.Type == "deposit" {
			acc.Deposit(tx.Amount)
			success = true
		} else if tx.Type == "withdraw" {
			acc.Withdraw(tx.Amount)
			success = true
		}
		if success {
			logTransaction(acc)
		} else {
			fmt.Printf("Transaction failed for account %d\n", tx.AccountID)
		}
	}
}

func logTransaction(acc *Account) {
	fileMu.Lock()
	defer fileMu.Unlock()
	f, err := os.OpenFile("transactions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()
	logline := fmt.Sprintf("Account %d: New Balance: %.2f\n", acc.ID, acc.Balance)
	f.WriteString(logline)
}

func main() {
	accounts[1] = &Account{ID: 1, Balance: 1000}
	accounts[2] = &Account{ID: 2, Balance: 1000}
	transactions := make(chan Transaction, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go porcesTransacton(transactions, &wg)

	txList := []Transaction{
		{"deposit", 500, 1},
		{"withdraw", 300, 1},
		{"withdraw", 2000, 2},
		{"withdraw", 400, 2},
	}
	for _, tx := range txList {
		fmt.Printf("Sending transaction to channel: Type: %s, Amount: %.2f, AccountID: %d\n", tx.Type, tx.Amount, tx.AccountID)
		transactions <- tx
	}
	close(transactions)
	wg.Wait()
	fmt.Println("All transactions processed.")
	for _, acc := range accounts {
		fmt.Printf("Account %d: Final Balance: %.2f\n", acc.ID, acc.Balance)
	}
	f, err := os.OpenFile("transactions.log", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		fmt.Printf("Error getting file stats: %v\n", err)
		return
	}
	if stat.Size() == 0 {
		fmt.Println("No transactions logged.")
		return
	}
	content := make([]byte, stat.Size())
	_, err = f.Read(content)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("Transaction log:")
	fmt.Println(string(content))
	fmt.Println("End of transaction log.")
	fmt.Println("End of program.")
}
