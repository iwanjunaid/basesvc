package controller

type AppController struct {
	// User   interface{ UserController }
	Author interface{ AuthorController }
}
