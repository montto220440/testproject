package repo

import (
	"fmt"
	db "testproject/_config"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOneStatement(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	ctx, cancel := db.InitContext()
	option := options.FindOne()
	option.SetSort(filterOption)
	defer cancel()
	fmt.Println("filter", filter)
	err := resource.DB.Collection(collection).FindOne(ctx, filter).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func GetManyStatement(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	ctx, cancel := db.InitContext()
	option := options.Find()
	option.SetSort(filterOption)
	defer cancel()
	obj, err := resource.DB.Collection(collection).Find(ctx, filter, option)
	if err != nil {
		return err
	}
	err = obj.All(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(resource *db.Resource, collection string, data interface{}) error {
	ctx, cancel := db.InitContext()
	defer cancel()
	_, err := resource.DB.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(resource *db.Resource, collection string, filter interface{}, data interface{}) error {
	ctx, cancel := db.InitContext()
	defer cancel()
	_, err := resource.DB.Collection(collection).UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(resource *db.Resource, collection string, filter interface{}) error {
	ctx, cancel := db.InitContext()
	defer cancel()
	_, err := resource.DB.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
