package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/dao"
	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		AppID:     uuid.NewString(),
		PayerID:   transaction.Payer.ID,
		PayeeID:   transaction.Payee.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: time.Now(),
	}

	if _, err := r.db.Collection(TRANSACTIONS_COLLECTION).InsertOne(context.Background(), document); err != nil {
		return nil, err
	}

	transaction.ID = document.AppID

	return &transaction, nil
}

func (r *TransactionRepositoryMongo) Update(transaction entity.Transaction) error {
	ctx := context.Background()
	options := options.Update().SetUpsert(true)

	filterID := bson.D{{"app_id", transaction.ID}}
	documentData := bson.D{{"$set", bson.M{
		"payer_id":   transaction.Payer.ID,
		"payee_id":   transaction.Payee.ID,
		"amount":     transaction.Amount,
		"status":     transaction.Status,
		"updated_at": time.Now(),
	}}}

	if _, err := r.db.Collection(TRANSACTIONS_COLLECTION).UpdateOne(ctx, filterID, documentData, options); err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryMongo) GetByID(ID string) (*entity.Transaction, error) {
	filterID := bson.D{{"app_id", ID}}
	var transactionDao dao.Transaction

	if err := r.db.Collection(TRANSACTIONS_COLLECTION).FindOne(context.Background(), filterID).Decode(&transactionDao); err != nil {
		return nil, err
	}

	transaction := &entity.Transaction{
		ID:     transactionDao.AppID,
		Payer:  entity.NewPayer(transactionDao.PayerID),
		Payee:  entity.NewPayee(transactionDao.PayeeID),
		Amount: transactionDao.Amount,
		Status: transactionDao.Status,
	}

	return transaction, nil
}
