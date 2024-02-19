package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"linkingsign/utils"
	"log"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/joho/godotenv"
)

func UploadLocal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	var tail = handler.Filename[len(handler.Filename)-4:]
	var nameFile = "upload-*" + tail
	tempFile, err := ioutil.TempFile("temp-files", nameFile)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func UploadCloud(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Max upload size

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Initialize Cloudinary configuration
	// load .env file
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	// Create Cloudinary uploader
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_API_KEY"), os.Getenv("CLOUD_API_SECRET"))
	if err != nil {
		fmt.Println("Error Initializing Cloudinary Configuration")
		fmt.Println(err)
		return
	}
	uploadApi := cld.Upload

	// get file name
	var publicID = handler.Filename[:len(handler.Filename)-4]

	// Upload file to Cloudinary
	result, err := uploadApi.Upload(context.Background(), file, uploader.UploadParams{PublicID: publicID})
	if err != nil {
		fmt.Println("Error Uploading File to Cloudinary")
		fmt.Println(err)
		return
	}

	fmt.Println("File Successfully Uploaded to Cloudinary")
	fmt.Println("Public ID:", result.PublicID)
	fmt.Println("URL:", result.SecureURL)

	// Return response to client
	w.WriteHeader(http.StatusOK)
	// format a response object
	res := utils.UploadResponse{
		Message: "File Successfully Uploaded to Cloudinary",
		URL:     result.SecureURL,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
