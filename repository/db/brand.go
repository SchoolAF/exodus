package db

import (
	"errors"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertBrand(requestData model.Brand) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.BrandColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func CountBrand(brand string) error {
	db := mongo.MongoConnect(DBHLCYN)
	filter := bson.M{"brandlower": brand}
	checkData, err := mongo.CountDocuments(db, config.BrandColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func GetBrandFilter(filter bson.M) ([]model.Brand, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.Brand
	err := mongo.GetAllDocByFilter[model.Brand](db, config.BrandColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneBrandFilter(filter bson.M) (model.Brand, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Brand
	err := mongo.GetOneDocByFilter[model.Brand](db, config.BrandColl, filter, &data)
	if err != nil {
		return model.Brand{}, err
	}
	return data, nil
}
