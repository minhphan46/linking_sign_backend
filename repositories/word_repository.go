package repositories

import (
	"database/sql"
	"fmt"
	"linkingsign/database"
	"linkingsign/models"
	"log"

	"github.com/google/uuid"
)

// insert one word in the DB
func InsertWord(word models.Word) models.Word {
	db := database.CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO words (topic_id, word_name, example1, example2, video, is_learned) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id string

	err := db.QueryRow(sqlStatement, word.TopicID, word.WordName, word.Example1, word.Example2, word.Video, word.IsLearned).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	word.ID = id

	fmt.Printf("Inserted a single record %v", id)
	return word
}

// get one word from the DB by its wordid
func GetWord(id string) (models.Word, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create a word of models.word type
	var word models.Word

	// create the select sql query
	sqlStatement := `SELECT * FROM words WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to word
	err := row.Scan(&word.ID, &word.TopicID, &word.WordName, &word.Example1, &word.Example2, &word.Video, &word.IsLearned)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return word, nil
	case nil:
		return word, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty word on error
	return word, err
}

// get one word from the DB by its wordid
func GetAllWords() ([]models.Word, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	var words []models.Word

	// create the select sql query
	sqlStatement := `SELECT * FROM words`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var word models.Word

		// unmarshal the row object to word
		err = rows.Scan(&word.ID, &word.TopicID, &word.WordName, &word.Example1, &word.Example2, &word.Video, &word.IsLearned)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the word in the words slice
		words = append(words, word)

	}

	// return empty word on error
	return words, err
}

// get one word from the DB by its wordid
// get all words from the DB by topicId
func GetAllWordsByTopicId(topicId string) ([]models.Word, error) {
	// Check if topicId is a valid UUID
	_, err := uuid.Parse(topicId)
	if err != nil {
		log.Fatalf("TopicId is not a valid UUID. %v", err)
	}

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	var words []models.Word

	// create the select sql query
	sqlStatement := `SELECT * FROM words WHERE topic_id = $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, topicId)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var word models.Word

		// unmarshal the row object to word
		err = rows.Scan(&word.ID, &word.TopicID, &word.WordName, &word.Example1, &word.Example2, &word.Video, &word.IsLearned)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the word in the words slice
		words = append(words, word)
	}

	// return the words
	return words, nil
}

// update word in the DB
func UpdateWord(id string, word models.Word) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE words SET topic_id=$2, word_name=$3, example1=$4, example2=$5, video=$6, is_learned=$7 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, word.TopicID, word.WordName, word.Example1, word.Example2, word.Video, word.IsLearned)

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

// delete word in the DB
func DeleteWord(id string) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM words WHERE id=$1`

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
