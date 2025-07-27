package documents

import "time"

type ImageUploadCreate struct {
	ID                    string    `firestore:"id" json:"id"`
	GoogleCloudStorageUrl string    `firestore:"google_cloud_storage_url" json:"google_cloud_storage_url"`
	LlmParsedContent      *string   `firestore:"parsed_content" json:"parsed_content,omitempty"`
	CreatedBy             *string   `firestore:"created_by" json:"created_by,omitempty"`
	CreatedAt             time.Time `firestore:"created_at" json:"created_at"`
}
