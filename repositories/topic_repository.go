package repositories

import (
	"database/sql"
	"fmt"
	"linkingsign/database"
	"linkingsign/models"
	"log"
)

// ------------------------- handler functions ----------------
// insert one topic in the DB
func InsertTopic(topic models.Topic) models.Topic {
	db := database.CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO topics (topic_name, topic_image, number_learned_lesson, total_lesson, state) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var id string

	err := db.QueryRow(sqlStatement, topic.TopicName, topic.TopicImage, topic.NumberLearnedLesson, topic.TotalLesson, topic.State).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	topic.ID = id

	fmt.Printf("Inserted a single record %v", id)
	return topic
}

// get one topic from the DB by its topicid
func GetTopic(id string) (models.Topic, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create a topic of models.Topic type
	var topic models.Topic

	// create the select sql query
	sqlStatement := `SELECT * FROM topics WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to topic
	err := row.Scan(&topic.ID, &topic.TopicName, &topic.TopicImage, &topic.NumberLearnedLesson, &topic.TotalLesson, &topic.State)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return topic, nil
	case nil:
		return topic, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty topic on error
	return topic, err
}

// get one topic from the DB by its topicid
func GetAllTopics() ([]models.Topic, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	var topics []models.Topic

	// create the select sql query
	sqlStatement := `SELECT * FROM topics`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var topic models.Topic

		// unmarshal the row object to topic
		err = rows.Scan(&topic.ID, &topic.TopicName, &topic.TopicImage, &topic.NumberLearnedLesson, &topic.TotalLesson, &topic.State)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the topic in the topics slice
		topics = append(topics, topic)

	}

	// return empty topic on error
	return topics, err
}

// update topic in the DB
func UpdateTopic(id string, topic models.Topic) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE topics SET topic_name=$2, topic_image=$3, number_learned_lesson=$4, total_lesson=$5, state=$6 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, topic.TopicName, topic.TopicImage, topic.NumberLearnedLesson, topic.TotalLesson, topic.State)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete topic in the DB
func DeleteTopic(id string) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM topics WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
