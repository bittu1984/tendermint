package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	transactions := getAllTransactions()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": transactions}, "index.html")
}

func showTransactionCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Transaction"}, "create-transaction.html")
}

func getTransaction(c *gin.Context) {
	if transactionID, err := strconv.Atoi(c.Param("transaction_id")); err == nil {
		if transaction, err := getTransactionByID(transactionID); err == nil {
			render(c, gin.H{
				"title":   transaction.Title,
				"payload": transaction}, "transaction.html")

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createTransaction(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewTransaction(title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
