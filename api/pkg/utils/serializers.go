package utils

import (
	"strconv"

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

func StringToIntWithDefault(s string, def int) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return def, err
	}

	return i, nil
}
