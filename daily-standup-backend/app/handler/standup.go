package handler

import (
	"daily-standup/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReportHandler ...
type ReportHandler struct {
	ReportEntity domain.ReportEntity
}

// NewReportHandler ...
func NewReportHandler(r *gin.RouterGroup, pe domain.ReportEntity) {
	handler := &ReportHandler{
		ReportEntity: pe,
	}

	r.POST("/daily-standup", handler.CreateReport)
	r.GET("/daily-standup", handler.FetchReports)
	r.PUT("/daily-standup/:id", handler.UpdateReport)
	r.DELETE("/daily-standup/:id", handler.DeleteReport)

}

// FetchReports ... 
func (a *ReportHandler) FetchReports(c *gin.Context) {

	post, err := a.ReportEntity.FetchReports(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// CreateReport ...
func (a *ReportHandler) CreateReport(c *gin.Context) {
	var report domain.Report 

	err := c.ShouldBind(&report)

	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err.Error()})
		return
	}

	response, err := a.ReportEntity.CreateReport(c.Request.Context(), report)
	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response.InsertedID})

}

// UpdateReport ...
func (a *ReportHandler) UpdateReport(c *gin.Context) {
	var report domain.Report
	err := c.ShouldBind(&report)

	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err})
		return
	}

	err = a.ReportEntity.UpdateReport(c.Request.Context(), c.Param("id"), report)
	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Ok"})
}

// DeleteReport ...
func (a *ReportHandler) DeleteReport(c *gin.Context) {

	err := a.ReportEntity.DeleteReport(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError,  gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Ok"})
}
