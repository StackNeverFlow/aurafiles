## AuraFiles | Anonymous file sharing

Aurafiles is the first real program I made with the Go programming language. Please don't judge me, but I'm still
learning Go.

### How to run

#### Using Git

1) Clone locally using `git clone github.com/Auralax/aurafiles.git`
2) Import into your IDE, I prefer GoLand

#### Using releases

1) Download the latest version from the releases tab
2) Extract the file to a folder on your server or computer

#### General setup

3) Copy `example.env`, rename to `.env` and change all data:

- `BACKEND_PORT` The port under which the server will be accessible (default is 8000)
- `FRONTEND_PORT` The port under which the frontend will be accessible (default is 8080)
- `MONGO_CONNECTION_URL` The url on which the mongo database is located
- `SECRET_USERNAME` The secret username used to authenticate requests (default is "Admin")
- `SECRET_PASSWORD` The secret password used to authenticate requests (default is "12345")
- `CUSTOM_NAME_LENGTH` The amount of letters custom names can have (default is 16 letters)

4) Run `startserver.sh` and `startfrontent.sh` to start the backend and frontend server

5) Have fun uploading

### Rest API

You have to add basic auth to the header of your requests. The credentials are specified in the `.env` file.

- `example.com:port/[id]` returns the file with a specific id
- `example.com:port/fileinfo/[id]` returns a json response contain information about a file with a specific id
    - `Id` The id of the upload
    - `OldName` The old name of the file before randomizing
    - `NewName` The name of the file after randomizing
    - `Downloads` The amount of downloads this upload has
    - `Date` The date on which the upload was made
    - `Type` The type of file
    - `Size` The size of the file in byte
- `example.com:port/upload` is used to upload a file to server. You have to add the file to upload in the body with the
  Key `upload`. Returns the same as `/fileinfo/`
- `example.com:port/download/[id]` is used to increase the download amount of an upload with a specific id, returns the
  same as `/fileinfo/`
    

