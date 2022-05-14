// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CheckEmailExist(ctx context.Context, email string) (bool, error)
	CheckIDExist(ctx context.Context, id uuid.UUID) (bool, error)
	CheckUsernameExist(ctx context.Context, username string) (bool, error)
	CountData(ctx context.Context) (int64, error)
	CountMessage(ctx context.Context) (int64, error)
	CountUser(ctx context.Context) (int64, error)
	CreateData(ctx context.Context, arg CreateDataParams) (Datum, error)
	CreateFile(ctx context.Context, arg CreateFileParams) (CreateFileRow, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) error
	DeleteFile(ctx context.Context, url sql.NullString) error
	DeleteMessageByID(ctx context.Context, id uuid.UUID) error
	DeleteOldRefreshToken(ctx context.Context) error
	DeleteRefreshToken(ctx context.Context, id uuid.UUID) error
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	ExistUserByEmailAndConfirmCode(ctx context.Context, arg ExistUserByEmailAndConfirmCodeParams) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	GetBoAllDatas(ctx context.Context, arg GetBoAllDatasParams) ([]Datum, error)
	GetBoAllUsers(ctx context.Context, arg GetBoAllUsersParams) ([]User, error)
	GetCodeByEmail(ctx context.Context, email string) (sql.NullString, error)
	GetDataByID(ctx context.Context, id uuid.UUID) (Datum, error)
	GetDataByUserID(ctx context.Context, arg GetDataByUserIDParams) ([]Datum, error)
	GetEmailByUserID(ctx context.Context, id uuid.UUID) (string, error)
	GetFileByURL(ctx context.Context, url sql.NullString) (File, error)
	GetMessageByID(ctx context.Context, id uuid.UUID) (Contact, error)
	GetMessagesFiltered(ctx context.Context, arg GetMessagesFilteredParams) ([]Contact, error)
	GetOldRefreshToken(ctx context.Context) (RefreshToken, error)
	GetRefreshToken(ctx context.Context, token string) (GetRefreshTokenRow, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserRandom(ctx context.Context) (User, error)
	InsertMessage(ctx context.Context, arg InsertMessageParams) error
	ListData(ctx context.Context, title string) ([]Datum, error)
	ListDataMostRecent(ctx context.Context, arg ListDataMostRecentParams) ([]Datum, error)
	ListRefreshTokenByUserID(ctx context.Context, arg ListRefreshTokenByUserIDParams) ([]RefreshToken, error)
	ListUsers(ctx context.Context, firstname sql.NullString) ([]User, error)
	ListUsersMostRecent(ctx context.Context, arg ListUsersMostRecentParams) ([]User, error)
	LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserRow, error)
	Signup(ctx context.Context, arg SignupParams) (User, error)
	UpdateAvatarUser(ctx context.Context, arg UpdateAvatarUserParams) error
	UpdateData(ctx context.Context, arg UpdateDataParams) error
	UpdateDescriptionUser(ctx context.Context, arg UpdateDescriptionUserParams) error
	UpdatePasswordUserWithconfirmCode(ctx context.Context, arg UpdatePasswordUserWithconfirmCodeParams) error
	UpdateUserBo(ctx context.Context, arg UpdateUserBoParams) error
	UpdateUserConfirmCode(ctx context.Context, arg UpdateUserConfirmCodeParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
