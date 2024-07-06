package db

import (
	"errors"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertForm(requestData model.MaintainerShipForm) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.ApplyFormColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func GetApplyFormFilter(filter bson.M) ([]model.MaintainerShipForm, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.MaintainerShipForm
	err := mongo.GetAllDocByFilter[model.MaintainerShipForm](db, config.ApplyFormColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}
