package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	Id        string             `json:"id,onitempty" bson:id,onitempty"`
	OldName   string             `json:"oldname,onitempty" bson:oldname,onitempty"`
	NewName   string             `json:"newname,onitempty" bson:newname,onitempty"`
	Downloads int64              `json:"downloads,onitempty" bson:downloads,onitempty"`
	Upload    primitive.DateTime `json:"upload,onitempty" bson:upload,onitempty"`
	Type      string             `json:"type,onitempty" bson:type,onitempty"`
	Size      int64              `json:"size,onitempty" bson:size,onitempty"`
}
