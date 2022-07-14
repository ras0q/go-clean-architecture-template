package infrastructure

import (
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/controller"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository"
)

func injectControllers() controller.Controllers {
	return controller.NewControllers(
		controller.NewUserController(
			repository.NewUserRepository(),
		),
	)
}
