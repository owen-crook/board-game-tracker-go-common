package documents

import "time"

type ImageUploadCreate struct {
	ID               string    `firestore:"id" json:"id"`
	Bucket           string    `firestore:"bucket" json:"bucket"`
	Path             string    `firestore:"path" json:"path"`
	LlmParsedContent *string   `firestore:"parsed_content" json:"parsed_content,omitempty"`
	CreatedBy        *string   `firestore:"created_by" json:"created_by,omitempty"`
	CreatedAt        time.Time `firestore:"created_at" json:"created_at"`
}
