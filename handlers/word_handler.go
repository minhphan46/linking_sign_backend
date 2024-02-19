package handlers

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api

	// used to read the environment variable
	"strconv" // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	"linkingsign/database"     // package where the db connection is created
	"linkingsign/models"       // models package where the models are defined
	"linkingsign/repositories" // package where all the db operations are defined
)

// Createword create a word in the postgres db
func CreateWord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateWord")

	// create an empty word of type models.word
	var word models.Word

	// decode the json request to word
	err := json.NewDecoder(r.Body).Decode(&word)

	fmt.Println(word)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert word function and pass the word
	insertID := repositories.InsertWord(word)

	// format a response object
	res := database.Response{
		Message: "word created successfully",
		Data:    insertID,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// Get word will return a single word by its id
func GetWord(w http.ResponseWriter, r *http.Request) {
	// get the wordid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getword function with word id to retrieve a single word
	word, err := repositories.GetWord(int64(id))

	if err != nil {
		log.Fatalf("Unable to get word. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(word)
}

// GetAllword will return all the words
func GetAllWord(w http.ResponseWriter, r *http.Request) {

	// get all the words in the db
	words, err := repositories.GetAllWords()

	if err != nil {
		log.Fatalf("Unable to get all word. %v", err)
	}

	fmt.Println("Get all words")

	// send all the words as response
	json.NewEncoder(w).Encode(words)
}

// Updateword update word's detail in the postgres db
func UpdateWord(w http.ResponseWriter, r *http.Request) {

	// get the wordid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty word of type models.word
	var word models.Word

	// decode the json request to word
	err = json.NewDecoder(r.Body).Decode(&word)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update word to update the word
	updatedRows := repositories.UpdateWord(int64(id), word)

	// format the message string
	msg := fmt.Sprintf("word updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := database.Response{
		Message: msg,
		Data:    updatedRows,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// Deleteword delete word's detail in the postgres db
func DeleteWord(w http.ResponseWriter, r *http.Request) {

	// get the wordid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteword, convert the int to int64
	deletedRows := repositories.DeleteWord(int64(id))

	// format the message string
	msg := fmt.Sprintf("word updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := database.Response{
		Message: msg,
		Data:    deletedRows,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}