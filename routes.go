package main

func initializeRoutes() {
	router.Use(setUserStatus())
	router.GET("/", showIndexPage)
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}
	transactionRoutes := router.Group("/transaction")
	{
		transactionRoutes.GET("/view/:transaction_id", getTransaction)
		transactionRoutes.GET("/create", ensureLoggedIn(), showTransactionCreationPage)
		transactionRoutes.POST("/create", ensureLoggedIn(), createTransaction)
	}
}
