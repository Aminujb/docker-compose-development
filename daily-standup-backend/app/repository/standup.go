package repository

import (
	"daily-standup/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostRepository ...
type ReportRepository struct {
	Collection *mongo.Collection
}

// NewReportRepository will create an object that represent the domain.ReportRepository interface
func NewReportRepository(collection *mongo.Collection) domain.ReportRepository {
	return &ReportRepository{collection}
}

// Fetch all report entries ...
func (c *ReportRepository) FetchReports(ctx context.Context) (res *[]primitive.M, err error) {
	var reports []bson.M

	var filter bson.M = bson.M{}

	cur, err := c.Collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	cur.All(context.Background(), &reports)

	return &reports, nil
}

// Create a report ...
func (c *ReportRepository) CreateReport(ctx context.Context, report domain.Report) (*mongo.InsertOneResult, error) {

	report.DateCreated = time.Now()
	report.DateUpdated = report.DateCreated	

	response, err := c.Collection.InsertOne(context.Background(), report)

	if err != nil {
		return response, err
	}

	return response, nil
}

// Update a report ...
func (c *ReportRepository) UpdateReport(ctx context.Context, id string, report domain.Report) error {

	report.DateUpdated = time.Now()

	update := bson.M{
		"$set": report,
	}

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	response, err := c.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		return err
	}

	if response.ModifiedCount == 0 {
		return errors.New("update failed")
	}

	return nil
}

// Delete a report ...
func (c *ReportRepository) DeleteReport(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := c.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		return err
	}

	return nil
}
