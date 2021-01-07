package data

import "go.mongodb.org/mongo-driver/bson/primitive"

// FileInfo represents info about a file stored in the database
type FileInfo struct {
	Id        string             `json:"id,onitempty" bson:"id,onitempty"`               // Id is a random generated sequence of string
	OldName   string             `json:"oldname,onitempty" bson:"oldname,onitempty"`     // OldName is the old name of the file before randomizing
	NewName   string             `json:"newname,onitempty" bson:"newname,onitempty"`     // NewName is the new name of the file after randomizing
	Downloads int64              `json:"downloads,onitempty" bson:"downloads,onitempty"` // Downloads is the amount of downloads
	Date      primitive.DateTime `json:"date,onitempty" bson:"date,onitempty"`           // Date is the date of upload
	Type      string             `json:"type,onitempty" bson:"type,onitempty"`           // Type is the type of content
	Size      int64              `json:"size,onitempty" bson:"size,onitempty"`           // Size is the size in byte
}
