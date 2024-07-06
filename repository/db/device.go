package db

import (
	"errors"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDevice(requestData model.Device) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.DeviceColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func CountDevice(codename string) error {
	db := mongo.MongoConnect(DBHLCYN)
	filter := bson.M{"codename": codename}
	checkData, err := mongo.CountDocuments(db, config.DeviceColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func GetDeviceFilter(filter bson.M) ([]model.Device, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.Device
	err := mongo.GetAllDocByFilter[model.Device](db, config.DeviceColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneDeviceFilter(filter bson.M) (model.Device, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Device
	err := mongo.GetOneDocByFilter[model.Device](db, config.DeviceColl, filter, &data)
	if err != nil {
		return model.Device{}, err
	}
	return data, nil
}
