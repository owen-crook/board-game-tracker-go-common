package documents

import "time"

type ScorecardDocumentCreate struct {
	ID                    string            `firestore:"id" json:"id"`
	ImageUploadMetadataID string            `firestore:"image_upload_metadata_id" json:"image_upload_metadata_id"`
	Game                  string            `firestore:"game" json:"game"`
	Date                  time.Time         `firestore:"date" json:"date"`
	IsCompleted           bool              `firestore:"is_completed" json:"is_completed"`
	Location              *string           `firestore:"location" json:"location,omitempty"`
	PlayerScores          *[]map[string]any `firestore:"player_scores" json:"player_scores,omitempty"`
	CreatedBy             *string           `firestore:"created_by" json:"created_by,omitempty"`
	CreatedAt             time.Time         `firestore:"created_at" json:"created_at"`
}

type ScorecardDocumentUpdate struct {
	ID           string            `firestore:"id" json:"id"`
	Game         *string           `firestore:"game,omitempty" json:"game,omitempty"`
	Date         *time.Time        `firestore:"date,omitempty" json:"date,omitempty"`
	IsCompleted  *bool             `firestore:"is_completed,omitempty" json:"is_completed,omitempty"`
	Location     *string           `firestore:"location,omitempty" json:"location,omitempty"`
	PlayerScores *[]map[string]any `firestore:"player_scores,omitempty" json:"player_scores,omitempty"`
	UpdatedBy    string            `firestore:"updated_by,omitempty" json:"updated_by,omitempty"`
	UpdatedAt    time.Time         `firestore:"updated_at,omitempty" json:"updated_at,omitempty"`
}
