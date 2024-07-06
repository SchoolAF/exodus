package db

import (
	"errors"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertVersion(requestData model.Version) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.VersionColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func CountVersion(branch string) error {
	db := mongo.MongoConnect(DBHLCYN)
	filter := bson.M{"branch": branch}
	checkData, err := mongo.CountDocuments(db, config.VersionColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func GetVersionFilter(filter bson.M) ([]model.Version, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.Version
	err := mongo.GetAllDocByFilter[model.Version](db, config.VersionColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneVersionFilter(filter bson.M) (model.Version, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Version
	err := mongo.GetOneDocByFilter[model.Version](db, config.VersionColl, filter, &data)
	if err != nil {
		return model.Version{}, err
	}
	return data, nil
}
