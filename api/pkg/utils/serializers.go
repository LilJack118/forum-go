package utils

import (
	"go.mongodb.org/mongo-driver/bson"
)

func StructToBson(s interface{}) (*bson.M, error) {
	bytesS, err := bson.Marshal(s)
	if err != nil {
		return nil, err
	}

	var bsonS bson.M
	if err := bson.Unmarshal(bytesS, &bsonS); err != nil {
		return nil, err
	}

	return &bsonS, nil
}
