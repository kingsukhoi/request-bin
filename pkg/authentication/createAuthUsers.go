package authentication

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/kingsukhoi/request-bin/pkg/db"
	"github.com/kingsukhoi/request-bin/pkg/sqlc"
)

// InitUsers will create the admin user if it doesn't exist
// and hash any plain text passwords
func InitUsers(ctx context.Context) error {
	pool := db.MustGetDatabase()
	queries := sqlc.New(pool)

	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		return err
	}

	adminFound := false
	for _, user := range users {
		if user.Username == "admin" {
			adminFound = true
		}
		if !strings.HasPrefix(user.Password, "$argon2id") {
			hashedPassword, errL := createPassword(user.Password)
			if errL != nil {
				return errL
			}
			errL = queries.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
				Password: hashedPassword,
				Username: user.Username,
			})
			if errL != nil {
				return errL
			}
		}
	}

	if !adminFound {
		adminPassword := genPassword()
		slog.Info("Admin user created, you can change the password in the database and restart the server to hash it",
			"password", adminPassword)

		var hashedPassword string
		hashedPassword, err = createPassword(adminPassword)
		if err != nil {
			return err
		}

		err = queries.CreateUser(ctx, sqlc.CreateUserParams{
			Username: "admin",
			Password: hashedPassword,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func createPassword(password string) (string, error) {
	argonParams := argon2id.DefaultParams
	argonParams.Iterations = 3
	hashedPassword, err := argon2id.CreateHash(password, argonParams)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}

func VerifyPassword(ctx context.Context, username string, password string) (bool, error) {
	pool := db.MustGetDatabase()
	queries := sqlc.New(pool)

	user, err := queries.GetUser(ctx, username)
	if err != nil {
		return false, err
	}

	passwordMatch, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return false, err
	}
	return passwordMatch, nil
}

func genPassword() string {
	randomBytes := make([]byte, 32)
	_, _ = rand.Read(randomBytes)

	return base64.URLEncoding.EncodeToString(randomBytes)[:30]
}
