package handlers

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api

	// used to read the environment variable
	// package used to covert string into int type

	"linkingsign/models"       // models package where the models are defined
	"linkingsign/repositories" // package where all the db operations are defined
	"linkingsign/utils"        // utility package to format the response

	"github.com/gorilla/mux" // used to get the params from the route
)

// CreateTopic create a topic in the postgres db
func CreateTopic(w http.ResponseWriter, r *http.Request) {
	// create an empty topic of type models.topic
	var topic models.Topic

	// decode the json request to topic
	err := json.NewDecoder(r.Body).Decode(&topic)

	fmt.Println(topic)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert topic function and pass the topic
	insertTopic := repositories.InsertTopic(topic)

	// format a response object
	res := utils.Response{
		Message: "Topic created successfully",
		Data:    insertTopic,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// Get topic will return a single topic by its id
func GetTopic(w http.ResponseWriter, r *http.Request) {
	// get the topicid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Unable to convert the string into int.  %v", err)
	// }

	id := params["id"]

	// call the gettopic function with topic id to retrieve a single topic
	topic, err := repositories.GetTopic(id)

	if err != nil {
		log.Fatalf("Unable to get topic. %v", err)
	}

	// format a response object
	res := utils.Response{
		Message: "Get single topic successfully",
		Data:    topic,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetAlltopic will return all the topics
func GetAllTopic(w http.ResponseWriter, r *http.Request) {

	// get all the topics in the db
	topics, err := repositories.GetAllTopics()

	if err != nil {
		log.Fatalf("Unable to get all topic. %v", err)
	}

	fmt.Println("Get all topics")

	// format a response object
	res := utils.Response{
		Message: "Get all topics successfully",
		Data:    topics,
	}

	// send all the topics as response
	json.NewEncoder(w).Encode(res)
}

// Updatetopic update topic's detail in the postgres db
func UpdateTopic(w http.ResponseWriter, r *http.Request) {

	// get the topicid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Unable to convert the string into int.  %v", err)
	// }

	id := params["id"]

	// create an empty topic of type models.topic
	var topic models.Topic

	// decode the json request to topic
	err := json.NewDecoder(r.Body).Decode(&topic)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update topic to update the topic
	updatedRows := repositories.UpdateTopic(id, topic)

	// format the message string
	msg := fmt.Sprintf("topic updated successfully. Total rows/record affected %v", updatedRows)

	topic.ID = id

	// format the response message
	res := utils.Response{
		Message: msg,
		Data:    topic,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// Deletetopic delete topic's detail in the postgres db
func DeleteTopic(w http.ResponseWriter, r *http.Request) {

	// get the topicid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Unable to convert the string into int.  %v", err)
	// }

	id := params["id"]

	// call the deletetopic, convert the int to int64
	deletedRows := repositories.DeleteTopic(id)

	// format the message string
	msg := fmt.Sprintf("topic delete successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := utils.Response{
		Message: msg,
		Data:    deletedRows,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
