package db

import (
	"errors"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount(requestData model.Accounts) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.AccountsColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func UpdateAccount(requestData model.Accounts, filter bson.M) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.ReplaceOneDoc(db, config.AccountsColl, filter, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func CountTeleID(id string) error {
	db := mongo.MongoConnect(DBHLCYN)
	filter := bson.M{"id": id}
	checkData, err := mongo.CountDocuments(db, config.AccountsColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func CountUsername(username string) error {
	db := mongo.MongoConnect(DBHLCYN)
	filter := bson.M{"username": username}
	checkData, err := mongo.CountDocuments(db, config.AccountsColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func GetAccountFilter(filter bson.M) ([]model.Accounts, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.Accounts
	err := mongo.GetAllDocByFilter[model.Accounts](db, config.AccountsColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneAccountFilter(filter bson.M) (model.Accounts, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Accounts
	err := mongo.GetOneDocByFilter[model.Version](db, config.AccountsColl, filter, &data)
	if err != nil {
		return model.Accounts{}, err
	}
	return data, nil
}
