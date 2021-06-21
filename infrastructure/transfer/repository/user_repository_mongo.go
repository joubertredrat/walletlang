package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/dao"
	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USERS_COLLECTION = "users"
)

type UserRepositoryMongo struct {
	db *mongo.Database
}

func NewUserRepositoryMongo(db *mongo.Database) UserRepositoryMongo {
	return UserRepositoryMongo{
		db: db,
	}
}

func (r *UserRepositoryMongo) GetPayerByID(ID string) (*entity.Payer, error) {
	return nil, nil
}

func (r *UserRepositoryMongo) GetPayeeByID(ID string) (*entity.Payee, error) {
	return nil, nil
}

func (r *UserRepositoryMongo) CreatePayer(payer *entity.Payer) (*entity.Payer, error) {
	document := dao.User{
		AppID:  uuid.NewString(),
		Amount: payer.Amount,
	}

	if _, err := r.db.Collection(USERS_COLLECTION).InsertOne(context.Background(), document); err != nil {
		return nil, err
	}

	payer.ID = document.AppID

	return payer, nil
}

func (r *UserRepositoryMongo) CreatePayee(payee *entity.Payee) (*entity.Payee, error) {
	document := dao.User{
		AppID:  uuid.NewString(),
		Amount: payee.Amount,
	}

	if _, err := r.db.Collection(USERS_COLLECTION).InsertOne(context.Background(), document); err != nil {
		return nil, err
	}

	payee.ID = document.AppID

	return payee, nil
}

func (r *UserRepositoryMongo) UpdatePayer(payer *entity.Payer) (*entity.Payer, error) {
	return nil, nil
}

func (r *UserRepositoryMongo) UpdatePayee(payee *entity.Payee) (*entity.Payee, error) {
	return nil, nil
}
