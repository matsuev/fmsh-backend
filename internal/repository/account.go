package repository

import (
	"context"
	"fmsh-backend/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AccountRepository struct {
	db *pgx.Conn
}

// NewAccountRepository ...
func NewAccountRepository(db *pgx.Conn) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

// Create ..
func (r *AccountRepository) Create(username, password string) (*models.Account, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	sqlQuery := `
		INSERT INTO public.account
		(username, password)
		VALUES ($1, $2)
		RETURNING *
		;
	`

	account := &models.Account{}

	if err := r.db.QueryRow(context.Background(), sqlQuery, username, encryptedPassword).Scan(
		&account.ID,
		&account.UserName,
		&account.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return account, nil
}

// Read ...
func (r *AccountRepository) Read(id uuid.UUID) (*models.Account, error) {
	sqlQuery := `
		SELECT * FROM public.account
		WHERE id=$1
	`

	account := &models.Account{}

	if err := r.db.QueryRow(context.Background(), sqlQuery, id).Scan(
		&account.ID,
		&account.UserName,
		&account.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return account, nil
}

// FindByUserName ...
func (r *AccountRepository) FindByUserName(username string) (*models.Account, error) {
	sqlQuery := `
		SELECT * FROM public.account
		WHERE username=$1
	`

	account := &models.Account{}

	if err := r.db.QueryRow(context.Background(), sqlQuery, username).Scan(
		&account.ID,
		&account.UserName,
		&account.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return account, nil
}

// Update ...
func (r *AccountRepository) Update(account *models.Account) error {
	sqlQuery := `
		UPDATE public.account
		WHERE id=$1
		SET username=$2, password=$3
	`

	_, err := r.db.Exec(context.Background(), sqlQuery,
		account.ID,
		account.UserName,
		account.EncryptedPassword,
	)

	return err
}

// Delete ..
func (r *AccountRepository) Delete(id uuid.UUID) error {
	sqlQuery := `
		DELETE public.account
		WHERE id=$1
	`
	_, err := r.db.Exec(context.Background(), sqlQuery, id)

	return err
}
