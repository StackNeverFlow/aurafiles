package data

import (
	"aurafiles/backend/database"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	db             = database.Connect("aurafiles")
	fileCollection = db.Collection("upload")
)

// DefaultRoute is used when someone is accessing a default route like "/"
func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	newRequest(r)

	http.Error(w, "no content", http.StatusNotFound)
}

// UploadFileRoute is used when someone is sending a post request to the upload route
func UploadFileRoute(w http.ResponseWriter, r *http.Request) {
	newRequest(r)
	if Auth(w, r) {
		// maximal size of file
		var maxSize int64 = 20

		err := r.ParseMultipartForm(maxSize)
		if err != nil {
			fmt.Println("File to big")
			fmt.Println(err)
			http.Error(w, "file is to big", http.StatusInternalServerError)
			return
		}

		file, handler, err := r.FormFile("upload")
		if err != nil {
			fmt.Println("Failed to get the file to upload")
			fmt.Println(err)
			http.Error(w, "error while retrieving the file", http.StatusInternalServerError)
			return
		}

		ending := strings.Split(handler.Filename, ".")
		fileType := ending[len(ending)-1]

		defer file.Close()

		num := randNum(16)
		tmpFile, err := ioutil.TempFile("./upload/", "*-"+num+"."+fileType)

		if err != nil {
			fmt.Println("Failed to create temporary directory")
			fmt.Println(err)
			http.Error(w, "error while uploading the file", http.StatusInternalServerError)
			return
		}

		defer tmpFile.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Failed to create temporary file")
			fmt.Println(err)
			http.Error(w, "error while uploading the file", http.StatusInternalServerError)
			return
		}

		tmpFile.Write(bytes)

		info := File{
			Id:        num,
			OldName:   handler.Filename,
			NewName:   tmpFile.Name(),
			Downloads: 0,
			Upload:    primitive.NewDateTimeFromTime(time.Now()),
			Type:      fileType,
		}

		fmt.Println(info)

		_, err = fileCollection.InsertOne(context.TODO(), &info)
		if err != nil {
			fmt.Println("Failed to insert file info into database")
			fmt.Println(err)
			http.Error(w, "error while uploading the file", http.StatusInternalServerError)
			return
		}

		fmt.Println("Successfully inserted file info into database")

		http.Error(w, "success", http.StatusOK)
	}
}

// GetFileInfoRoute is used when someone is requesting information about a uploaded file
func GetFileInfoRoute(w http.ResponseWriter, r *http.Request) {
	newRequest(r)
	if Auth(w, r) {
		w.Header().Set("connection-type", "application/json")

		var params = mux.Vars(r)
		var id = params["id"]

		var file File

		filter := bson.M{"id": id}
		err := fileCollection.FindOne(context.TODO(), filter).Decode(&file)
		if err != nil {
			fmt.Println("File with id ", id, " is not available!")
			fmt.Println(err)
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(file)
	}
}

// GetFileRoute is used when someone is requesting a file
func GetFileRoute(w http.ResponseWriter, r *http.Request) {
	newRequest(r)
	//if Auth(w, r) {
	w.Header().Set("connection-type", "application/json")

	var params = mux.Vars(r)
	var id = params["id"]

	var file File

	filter := bson.M{"id": id}
	err := fileCollection.FindOne(context.TODO(), filter).Decode(&file)
	if err != nil {
		fmt.Println("File with id ", id, " is not available!")
		fmt.Println(err)
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	b, err := ioutil.ReadFile(file.NewName)

	w.Write(b)
	//}
}

// AddDownloadRoute is used when someone is adding a download to a file
func AddDownloadRoute(w http.ResponseWriter, r *http.Request) {
	newRequest(r)
	if Auth(w, r) {
		params := mux.Vars(r)
		id := params["id"]

		filter := bson.M{"id": id}

		var result File

		err := fileCollection.FindOne(context.TODO(), bson.D{}).Decode(&result)
		if err != nil {
			fmt.Println("Failed to get info about file ", id)
			fmt.Println(err)
			http.Error(w, "unknown id", http.StatusNotFound)
			return
		}

		update := bson.D{
			{"$set", bson.D{
				{"id", id},
				{"oldname", result.OldName},
				{"newname", result.NewName},
				{"downloads", result.Downloads + 1},
				{"upload", result.Upload},
				{"type", result.Type},
			}},
		}

		res := fileCollection.FindOneAndUpdate(context.TODO(), filter, update)

		if res.Err() == mongo.ErrNoDocuments {
			fmt.Println("Unknown id ", id)
			fmt.Println(res.Err())
			http.Error(w, "unknown id", http.StatusNotFound)
			return
		}

		fmt.Println("Successfully increased the download amount of file "+id+" to ", result.Downloads+1)

		http.Error(w, "success", http.StatusOK)
	}
}

// randNum is used to generate a new string containing a random sequence of numbers
// The length of the sequence is specified by the length argument
func randNum(length int) string {
	num := make([]byte, length)
	rand.Read(num)
	return fmt.Sprintf("%x", num)
}

// newRequest is used to print information to the console when a new request is send to the server
// Example: "New Request from: 127.0.0.1:37050 type: POST at:  2021-01-07 15:41:01.08680352 +0100 CET m=+5.812450204"
func newRequest(r *http.Request) {
	fmt.Println("New Request from: "+r.RemoteAddr+" type: "+r.Method+" at: ", time.Now())
}
