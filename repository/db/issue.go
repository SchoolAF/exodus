package db

import (
	"errors"
	"fmt"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertIssue(requestData model.IssueData) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.IssuesColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func GetIssueFilter(filter bson.M) ([]model.IssueData, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.IssueData
	err := mongo.GetAllDocByFilter[model.IssueData](db, config.IssuesColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneIssueFilter(filter bson.M) (model.IssueData, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.IssueData
	err := mongo.GetOneDocByFilter[model.IssueData](db, config.IssuesColl, filter, &data)
	if err != nil {
		return model.IssueData{}, err
	}
	return data, nil
}

func DeleteIssue(filter bson.M) (model.IssueData, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.IssueData
	result := mongo.DeleteOneDoc(db, config.IssuesColl, filter)

	if result == nil || result.DeletedCount == 0 {
		return model.IssueData{}, fmt.Errorf("failed to delete document: no documents matched the filter")
	}

	return data, nil
}

func UpdateIssue(filter bson.M, updateData model.IssueData) error {
	db := mongo.MongoConnect(DBHLCYN)
	result := mongo.ReplaceOneDoc(db, config.IssuesColl, filter, updateData)
	if result == nil || result.MatchedCount == 0 {
		return fmt.Errorf("no matching document found for update")
	}
	return nil
}
