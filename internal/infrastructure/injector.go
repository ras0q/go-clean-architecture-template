package infrastructure

import (
	"github.com/ras0q/go-clean-architecture-template/internal/interfaces/controller"
	"github.com/ras0q/go-clean-architecture-template/internal/interfaces/repository"
	"github.com/ras0q/go-clean-architecture-template/internal/interfaces/repository/ent"
)

func InjectControllers(ec *ent.Client) controller.Controllers {
	return controller.NewControllers(
		controller.NewUserController(
			repository.NewUserRepository(ec.User),
		),
	)
}
