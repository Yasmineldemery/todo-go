package models

import "gopkg.in/mgo.v2/bson"

type Todo struct{
	Id bson.ObjectId `json:"id" bson:"_id"`
	Task string  `json:"task" bson:"task"`
	Gone bool  `json:"done" bson:"done"`
	Priority int `json:"priority" bson:"priority"`
}