package domain

import (
	"context"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Standup schema struct
type Standup struct {
	Yesterday	string	`json:"yesterday" bson:"yesterday"`
	Today    	string	`json:"today" bson:"today"`
	Blockers    string	`json:"blockers" bson:"blockers"`
}

// Report schema struct
type Report struct {
	ID           	string `json:"_id,omitempty" bson:"_id,omitempty"`
	Username		string		`json:"username" bson:"username"`
	DateCreated    	time.Time	`json:"date_created" bson:"date_created"`
	DateUpdated    	time.Time	`json:"date_updated" bson:"date_updated"`
	Report			Standup		`json:"report" bson:"report"`
}

// ReportEntity interface
type ReportEntity interface {
	CreateReport(ctx context.Context, report Report) (*mongo.InsertOneResult, error)
	FetchReports(ctx context.Context) (*[]primitive.M, error)
	UpdateReport(ctx context.Context, id string, report Report) error
	DeleteReport(ctx context.Context, id string) error
}

// ReportRepository interface
type ReportRepository interface {
	CreateReport(ctx context.Context, report Report) (*mongo.InsertOneResult, error)
	FetchReports(ctx context.Context) (*[]primitive.M, error)
	UpdateReport(ctx context.Context, id string, report Report) error
	DeleteReport(ctx context.Context, id string) error
}
