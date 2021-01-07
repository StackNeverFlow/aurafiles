package data

import "go.mongodb.org/mongo-driver/bson/primitive"

// File represents a file in the mongo database
type File struct {

	// Id is a random generated sequence of string
	Id string `json:"id,onitempty" bson:id,onitempty"`

	// OldName is the old name of the file before randomizing
	OldName string `json:"oldname,onitempty" bson:oldname,onitempty"`

	// NewName is the new name of the file after randomizing
	NewName string `json:"newname,onitempty" bson:newname,onitempty"`

	// Downloads is the amount of downloads
	Downloads int64 `json:"downloads,onitempty" bson:downloads,onitempty"`

	// Upload is the date of upload
	Upload primitive.DateTime `json:"upload,onitempty" bson:upload,onitempty"`

	// Type is the type of content
	Type string `json:"type,onitempty" bson:type,onitempty"`

	// Size is the size in byte
	Size int64 `json:"size,onitempty" bson:size,onitempty"`
}
