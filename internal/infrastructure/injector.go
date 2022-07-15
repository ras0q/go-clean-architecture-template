package infrastructure

import (
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/controller"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
)

func InjectControllers(ec *ent.Client) controller.Controllers {
	return controller.NewControllers(
		controller.NewUserController(
			repository.NewUserRepository(ec.User),
		),
	)
}
