package controller

type Controllers struct {
	user UserController
}

func NewControllers(user UserController) Controllers {
	return Controllers{
		user: user,
	}
}

func (c Controllers) User() UserController {
	return c.user
}
