package entity

import (
	"daily-standup/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ReportEntity ...
type ReportEntity struct {
	repo domain.ReportRepository
}

// NewPostEntity will create new an articleUsecase object representation of domain.ReportEntity interface
func NewReportEntity(a domain.ReportRepository) domain.ReportEntity {
	return &ReportEntity{
		repo: a,
	}
}

// FetchPost retrives post record(s)...
func (a *ReportEntity) FetchReports(c context.Context) (*[]primitive.M, error) {
	return a.repo.FetchReports(c)
}

//CreatePost creates a single post record...
func (a *ReportEntity) CreateReport(c context.Context, post domain.Report) (*mongo.InsertOneResult, error) {
	return a.repo.CreateReport(c, post)
}

//UpdatePost creates a single post record...
func (a *ReportEntity) UpdateReport(c context.Context, id string, post domain.Report) error {
	return a.repo.UpdateReport(c, id, post)
}

//DeletePost creates a single post record...
func (a *ReportEntity) DeleteReport(c context.Context, id string)  error {
	return a.repo.DeleteReport(c, id)
}
