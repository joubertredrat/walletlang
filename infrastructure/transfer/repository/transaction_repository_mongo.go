package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/dao"
	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TRANSACTIONS_COLLECTION = "transactions"
)

type TransactionRepositoryMongo struct {
	db *mongo.Database
}

func NewTransactionRepositoryMongo(db *mongo.Database) TransactionRepositoryMongo {
	return TransactionRepositoryMongo{
		db: db,
	}
}

func (r *TransactionRepositoryMongo) Create(transaction entity.Transaction) (*entity.Transaction, error) {
	document := dao.Transaction{
		AppID:   uuid.NewString(),
		PayerID: transaction.Payer.ID,
		PayeeID: transaction.Payer.ID,
		Amount:  transaction.Amount,
		Status:  transaction.Status,
	}

	if _, err := r.db.Collection(TRANSACTIONS_COLLECTION).InsertOne(context.Background(), document); err != nil {
		return nil, err
	}

	transaction.ID = document.AppID

	return &transaction, nil
}

func (r *TransactionRepositoryMongo) Update(transaction entity.Transaction) error {
	return nil
}

func (r *TransactionRepositoryMongo) GetByID(ID string) (*entity.Transaction, error) {
	return nil, nil
}
