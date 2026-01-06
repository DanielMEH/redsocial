package database

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/infrastructure/utils"
	"github.com/redsocial/internal/module/users/adapters/services/database/gorm/models"
	"github.com/redsocial/internal/module/users/domains/command"
	"github.com/redsocial/internal/module/users/domains/entities/login"
	"github.com/redsocial/internal/module/users/domains/entities/profile"
	"github.com/redsocial/internal/module/users/domains/entities/register"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServicesDatabaseAdapter struct {
	dbgorm *gorm.DB
}

func NewServicesDatabase() *ServicesDatabaseAdapter {
	dsn := os.Getenv("URL_POSTGRESS_DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
	}

	err = db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Post{}, &models.Like{})
	if err != nil {
		slog.Error("Error en la migración", "error", err)
	}
	slog.Info("Conectado a la base de datos")

	return &ServicesDatabaseAdapter{
		dbgorm: db,
	}
}

func (sd *ServicesDatabaseAdapter) RegisterAccount(data command.EntityRegisterAccountCommand) (register.EntityRegisterAccountResponse, error) {

	/* uuid para el la session */
	uuidValue := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return register.EntityRegisterAccountResponse{}, config.NewInternalServerError(err)
	}

	/* onjeto para creal el usuario */
	newUser := models.User{
		Email:    data.Email,
		Password: string(hashedPassword),
		Profile: models.Profile{
			Alias:     data.DisplayName,
			BirthDate: time.Now(),
		},
	}

	/* ejecutar consulta*/
	result := sd.dbgorm.Create(&newUser)

	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" {
				return register.EntityRegisterAccountResponse{},
					config.NewBadRequestError(errors.New("el correo electrónico  ya está registrado"))
			}
		}
		return register.EntityRegisterAccountResponse{}, config.NewInternalServerError(result.Error)
	}

	/* jwt*/
	value, errJwt := utils.GenerateJWT(string(newUser.ID.String()), uuidValue)

	if errJwt != nil {
		println("error")
		return register.EntityRegisterAccountResponse{}, config.NewInternalServerError(errJwt)
	}

	return register.EntityRegisterAccountResponse{
		Message: "ok",
		Details: struct {
			AccountToken     string "json:\"account_token\""
			AccountSessionId string "json:\"account_session_id\""
		}{
			AccountToken:     value,
			AccountSessionId: uuidValue,
		},
	}, nil
}

func (sd *ServicesDatabaseAdapter) LoginAccount(data command.EntityLoginAccountCommand) (login.EntityLoginAccountResponse, error) {

	var TableUsers models.User

	resposneRow := sd.dbgorm.Where("email = ?", data.Email).First(&TableUsers)

	if resposneRow.Error != nil {
		if errors.Is(resposneRow.Error, gorm.ErrRecordNotFound) {
			slog.Error(resposneRow.Error.Error())
			return login.EntityLoginAccountResponse{}, config.NewAuthInvalidCredentials(errors.New("credenciales incorrectas"))
		}
		slog.Error(resposneRow.Error.Error())
		return login.EntityLoginAccountResponse{}, config.NewInternalServerError(resposneRow.Error)
	}

	err := bcrypt.CompareHashAndPassword([]byte(TableUsers.Password), []byte(data.Password))
	if err != nil {
		return login.EntityLoginAccountResponse{}, config.NewBadRequestError(errors.New("credenciales incorrectas"))
	}
	sessionID := uuid.New().String()
	token, errJwt := utils.GenerateJWT(TableUsers.ID.String(), sessionID)
	if errJwt != nil {
		return login.EntityLoginAccountResponse{}, config.NewInternalServerError(errJwt)
	}

	return login.EntityLoginAccountResponse{
		Message: "ok",
		Details: struct {
			AccountToken     string "json:\"account_token\""
			AccountSessionId string "json:\"account_session_id\""
		}{
			AccountToken:     token,
			AccountSessionId: sessionID,
		},
	}, nil
}

func (sd *ServicesDatabaseAdapter) GetProfile(accountId string) (profile.EntityGetProfileAccountResponse, error) {
	var user models.User
	result := sd.dbgorm.
		Select("id", "email").
		Preload("Profile", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id", "alias")
		}).
		Where("id = ?", accountId).
		First(&user)

	if result.Error != nil {
		return profile.EntityGetProfileAccountResponse{}, config.NewInternalServerError(result.Error)

	}
	return profile.EntityGetProfileAccountResponse{
		Message: "Ok",
		Details: struct {
			Email     string    "json:\"email\""
			Alias     string    "json:\"alias\""
			BirthDate time.Time "json:\"bith_date\""
		}{
			Email:     user.Email,
			Alias:     user.Profile.Alias,
			BirthDate: user.Profile.BirthDate,
		},
	}, nil
}
