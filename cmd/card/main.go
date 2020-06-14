package main

import (
	"fmt"
	"github.com/DABronskikh/go-lesson-3_task-2/pkg/card"
	"time"
)

func main() {
	master := &card.Card{
		Id:       1,
		Issuer:   "MasterCard",
		Balance:  50_000_00,
		Currency: "RUB",
		Number:   "0001",
		Icon:     "MasterCardIcon",
	}

	var transactionsCard []card.Transaction

	transactionsCard = append(transactionsCard, card.Transaction{
		Id:     1,
		Amount: -735_55,
		Date:   time.Now().Unix(),
		MCC:    "5411",
		Type:   "payment",
		Status: "in progress",
	})
	transactionsCard = append(transactionsCard, card.Transaction{
		Id:     2,
		Amount: 2000_00,
		Date:   time.Now().Unix(),
		MCC:    "",
		Type:   "top up",
		Status: "completed",
	})
	transactionsCard = append(transactionsCard, card.Transaction{
		Id:     3,
		Amount: -1203_91,
		Date:   time.Now().Unix(),
		MCC:    "5411",
		Type:   "payment",
		Status: "in progress",
	})

	master.Transactions = transactionsCard

	newTransaction := &card.Transaction{
		Id:     4,
		Amount: 4000_00,
		Date:   time.Now().Unix(),
		MCC:    "",
		Type:   "top up",
		Status: "completed",
	}

	card.AddTransaction(master, newTransaction)
	fmt.Println(master)

	mccArray := []string{"5411", "5812"}

	total := card.SumByMCC(master.Transactions, mccArray)
	fmt.Println(total)

	category := card.TranslateMCC(master.Transactions[0].MCC)
	fmt.Println(category)

}
