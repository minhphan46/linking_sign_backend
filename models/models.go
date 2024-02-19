package models

type Topic struct {
	ID                  string `json:"id"`
	TopicName           string `json:"topic_name"`
	TopicImage          string `json:"topic_image"`
	NumberLearnedLesson int    `json:"number_learned_lesson"`
	TotalLesson         int    `json:"total_lesson"`
	State               string `json:"state"`
}

// state: "learning", "learned", "not learning"

type Word struct {
	ID        string `json:"id"`
	TopicID   int    `json:"topic_id"`
	WordName  string `json:"word_name"`
	Example1  string `json:"example1"`
	Example2  string `json:"example2"`
	Video     string `json:"video"`
	IsLearned bool   `json:"is_learned"`
}
