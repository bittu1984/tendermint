package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type transaction struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var transactionList = []transaction{
	transaction{ID: 1, Title: "Transaction 1", Content: "Transaction 1 body"},
	transaction{ID: 2, Title: "Transaction 2", Content: "Transaction 2 body"},
}

func getAllTransactions() []transaction {
	// Was working here to implement interface for tendermint environment.
	resp, err := http.Get("http://127.0.0.1:26657/abci_query?data=\"tendermint\"")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body)) // response for "tendermint" logged in console.
	// :::End:::::
	return transactionList
}

func getTransactionByID(id int) (*transaction, error) {
	// get
	for _, a := range transactionList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Transaction not found")
}

// To create
func createNewTransaction(title, content string) (*transaction, error) {
	/*
		resp, err := http.Get("localhost:26657/broadcast_tx_commit?tx=\"tendermint=rocks\"")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	*/
	a := transaction{ID: len(transactionList) + 1, Title: title, Content: content}
	transactionList = append(transactionList, a)
	return &a, nil
}
