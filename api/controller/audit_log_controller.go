package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hms-api/domain"
	"net/http"
)

type AuditLogController struct {
	AuditLogUsecase domain.AuditLogUsecase
}

func NewAuditLogController(usecase domain.AuditLogUsecase) *AuditLogController {
	return &AuditLogController{
		AuditLogUsecase: usecase,
	}
}

func (alc *AuditLogController) Create(c *gin.Context) {
	var auditLog domain.AuditLog

	err := c.ShouldBind(&auditLog)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = alc.AuditLogUsecase.Create(c, &auditLog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, auditLog)
}

func (alc *AuditLogController) Fetch(c *gin.Context) {
	auditLogs, err := alc.AuditLogUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, auditLogs)
}

func (alc *AuditLogController) FetchByID(c *gin.Context) {
	auditLogID := c.Param("id")
	if auditLogID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "audit log id is required"})
		return
	}

	parsedID, err := uuid.Parse(auditLogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid audit log id format"})
		return
	}

	auditLog, err := alc.AuditLogUsecase.FetchByID(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if auditLog.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Audit log not found"})
		return
	}

	c.JSON(http.StatusOK, auditLog)
}

func (alc *AuditLogController) Update(c *gin.Context) {
	auditLogID := c.Param("id")

	if auditLogID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Audit log id is required"})
		return
	}

	parsedID, err := uuid.Parse(auditLogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid auditoLog id"})
		return
	}

	var auditLog domain.AuditLog
	err = c.ShouldBind(&auditLog)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	auditLog.ID = parsedID

	err = alc.AuditLogUsecase.Update(c, &auditLog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, auditLog)
}

func (alc *AuditLogController) Delete(c *gin.Context) {
	auditLogID := c.Param("id")
	if auditLogID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Audit log id is required"})
		return
	}

	parsedID, err := uuid.Parse(auditLogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid audit log id"})
		return
	}

	err = alc.AuditLogUsecase.Delete(c, parsedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
