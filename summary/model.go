package summary

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Summary struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	Link           string             `json:"link,omitempty" validate:"required"`
	Prompt         string             `json:"prompt,omitempty" validate:"required"`
	GptSummary     string             `json:"gpt_summary,omitempty" validate:"required"`
	HindiTranslate string             `json:"hindi_translate,omitempty" validate:"required"`
	DownloadLink   string             `json:"download_link,omitempty"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
