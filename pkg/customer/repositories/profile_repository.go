package user_repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/messages"
	"gituh.com/adi-kmt/ecommerce-ixl-go/pkg/entities"
)

func (repo *UserRepository) InsertUser(ctx *fiber.Ctx, name, email, address, password string, isAdmin bool) *messages.AppError {

	err0 := repo.q.InsertIntoUsersTable(ctx.Context(), db.InsertIntoUsersTableParams{
		Name:     name,
		Email:    email,
		Password: password,
		Address:  address,
		Isadmin:  isAdmin,
	})

	if err0 != nil {
		log.Debugf("Error Inserting User: %v", err0)
		return messages.InternalServerError("Error Inserting User")
	}
	return nil
}

func (repo *UserRepository) GetUserDetails(ctx *fiber.Ctx, email string) (*db.GetUserEmailAndPasswordByEmailRow, *messages.AppError) {
	user, err := repo.q.GetUserEmailAndPasswordByEmail(ctx.Context(), email)
	if err != nil {
		log.Debugf("Error Getting User: %v", err)
		return nil, messages.InternalServerError("Error Getting User")
	}
	return user, nil
}

func (repo *UserRepository) GetUserDetailsAndOrders(ctx *fiber.Ctx, userId int64) (*entities.UserDetailsAndOrdersDto, *messages.AppError) {
	userAndOrderDetails, err := repo.q.GetUserDetailsAndOrders(ctx.Context(), userId)
	if err != nil {
		log.Debugf("Error Getting Profile: %v", err)
		return nil, messages.InternalServerError("Error Getting Profile")
	}
	return entities.UserDetailsAndOrdersFromDb(userAndOrderDetails), nil
}
