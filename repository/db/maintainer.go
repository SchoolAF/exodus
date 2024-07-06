package db

import (
	"errors"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertMaintainer(requestData model.Maintainer) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.MaintainerColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func CountMaintainer(username string) error {
	db := mongo.MongoConnect(DBHLCYN)
	filter := bson.M{"username": username}
	checkData, err := mongo.CountDocuments(db, config.MaintainerColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func GetMaintainerFilter(filter bson.M) ([]model.Maintainer, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.Maintainer
	err := mongo.GetAllDocByFilter[model.Maintainer](db, config.MaintainerColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneMaitainerFilter(filter bson.M) (model.Maintainer, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Maintainer
	err := mongo.GetOneDocByFilter[model.Maintainer](db, config.MaintainerColl, filter, &data)
	if err != nil {
		return model.Maintainer{}, err
	}
	return data, nil
}
