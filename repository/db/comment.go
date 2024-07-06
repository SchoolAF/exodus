package db

import (
	"errors"
	"fmt"
	"github.com/SchoolAF/exodus/config"
	"github.com/SchoolAF/exodus/model"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertComment(requestData model.Comments) error {
	db := mongo.MongoConnect(DBHLCYN)
	insertedID := mongo.InsertOneDoc(db, config.CommentsColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func GetCommentFilter(filter bson.M) ([]model.Comments, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var datas []model.Comments
	err := mongo.GetAllDocByFilter[model.Comments](db, config.CommentsColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneCommentFilter(filter bson.M) (model.Comments, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Comments
	err := mongo.GetOneDocByFilter[model.Comments](db, config.CommentsColl, filter, &data)
	if err != nil {
		return model.Comments{}, err
	}
	return data, nil
}

func DeleteComment(filter bson.M) (model.Comments, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Comments
	result := mongo.DeleteOneDoc(db, config.CommentsColl, filter)

	if result == nil || result.DeletedCount == 0 {
		return model.Comments{}, fmt.Errorf("failed to delete document: no documents matched the filter")
	}

	return data, nil
}

func DeleteAllComment(filter bson.M) (model.Comments, error) {
	db := mongo.MongoConnect(DBHLCYN)
	var data model.Comments
	result := mongo.DeleteDoc(db, config.CommentsColl, filter)

	if result == nil || result.DeletedCount == 0 {
		return model.Comments{}, fmt.Errorf("failed to delete document: no documents matched the filter")
	}

	return data, nil
}

func UpdateComment(filter bson.M, updateData model.Comments) error {
	db := mongo.MongoConnect(DBHLCYN)
	result := mongo.ReplaceOneDoc(db, config.CommentsColl, filter, updateData)
	if result == nil || result.MatchedCount == 0 {
		return fmt.Errorf("no matching document found for update")
	}
	return nil
}
