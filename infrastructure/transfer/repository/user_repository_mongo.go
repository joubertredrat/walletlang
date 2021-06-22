package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/dao"
	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"github.com/joubertredrat/walletlang/internal/transfer/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	user, err := r.getUserByID(ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.UserRepositoryPayerNotFoundError
		}

		return nil, err
	}

	payer := entity.Payer{
		ID:     user.AppID,
		Amount: user.Amount,
	}

	return &payer, nil
}

func (r *UserRepositoryMongo) GetPayeeByID(ID string) (*entity.Payee, error) {
	user, err := r.getUserByID(ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, repository.UserRepositoryPayeeNotFoundError
		}

		return nil, err
	}

	payee := entity.Payee{
		ID:     user.AppID,
		Amount: user.Amount,
	}

	return &payee, nil
}

func (r *UserRepositoryMongo) CreatePayer(payer *entity.Payer) (*entity.Payer, error) {
	document := dao.User{
		AppID:     uuid.NewString(),
		Amount:    payer.Amount,
		CreatedAt: time.Now(),
	}

	if err := r.createUser(document); err != nil {
		return nil, err
	}

	payer.ID = document.AppID

	return payer, nil
}

func (r *UserRepositoryMongo) CreatePayee(payee *entity.Payee) (*entity.Payee, error) {
	document := dao.User{
		AppID:     uuid.NewString(),
		Amount:    payee.Amount,
		CreatedAt: time.Now(),
	}

	if err := r.createUser(document); err != nil {
		return nil, err
	}

	payee.ID = document.AppID

	return payee, nil
}

func (r *UserRepositoryMongo) UpdatePayer(payer *entity.Payer) (*entity.Payer, error) {
	ctx := context.Background()
	options := options.Update().SetUpsert(true)

	filterID := bson.D{{"app_id", payer.ID}}
	documentData := bson.D{{"$set", bson.M{
		"amount":     payer.Amount,
		"updated_at": time.Now(),
	}}}

	if _, err := r.db.Collection(USERS_COLLECTION).UpdateOne(ctx, filterID, documentData, options); err != nil {
		return nil, err
	}

	return payer, nil
}

func (r *UserRepositoryMongo) UpdatePayee(payee *entity.Payee) (*entity.Payee, error) {
	ctx := context.Background()
	options := options.Update().SetUpsert(true)

	filterID := bson.D{{"app_id", payee.ID}}
	documentData := bson.D{{"$set", bson.M{
		"amount":     payee.Amount,
		"updated_at": time.Now(),
	}}}

	if _, err := r.db.Collection(USERS_COLLECTION).UpdateOne(ctx, filterID, documentData, options); err != nil {
		return nil, err
	}

	return payee, nil
}

func (r *UserRepositoryMongo) getUserByID(ID string) (*dao.User, error) {
	filterID := bson.D{{"app_id", ID}}
	var user dao.User

	if err := r.db.Collection(USERS_COLLECTION).FindOne(context.Background(), filterID).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryMongo) createUser(user dao.User) error {
	if _, err := r.db.Collection(USERS_COLLECTION).InsertOne(context.Background(), user); err != nil {
		return err
	}

	return nil
}
