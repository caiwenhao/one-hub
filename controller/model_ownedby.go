package controller

import (
	"errors"
	"net/http"
	"one-api/common"
	"one-api/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllModelOwnedBy(c *gin.Context) {
	modelOwnedBies, err := model.GetAllModelOwnedBy()
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    modelOwnedBies,
	})
}

func GetModelOwnedBy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	modelOwnedBy, err := model.GetModelOwnedBy(id)
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    modelOwnedBy,
	})
}

func CreateModelOwnedBy(c *gin.Context) {
	var modelOwnedBy model.ModelOwnedBy
	if err := c.ShouldBindJSON(&modelOwnedBy); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	if checkModelOwnedByReserveID(modelOwnedBy.Id) {
		common.APIRespondWithError(c, http.StatusOK, errors.New("invalid id"))
		return
	}

	if err := model.CreateModelOwnedBy(&modelOwnedBy); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}

func UpdateModelOwnedBy(c *gin.Context) {
	var modelOwnedBy model.ModelOwnedBy
	if err := c.ShouldBindJSON(&modelOwnedBy); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	if err := model.UpdateModelOwnedBy(&modelOwnedBy); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}

func DeleteModelOwnedBy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	if checkModelOwnedByReserveID(id) {
		common.APIRespondWithError(c, http.StatusOK, errors.New("invalid id"))
		return
	}

	if err := model.DeleteModelOwnedBy(id); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}

func GetModelBrandOverview(c *gin.Context) {
	query := &model.ModelBrandOverviewQuery{}

	query.Keyword = strings.TrimSpace(c.Query("keyword"))
	query.Group = strings.TrimSpace(c.Query("group"))

	if owned := strings.TrimSpace(c.Query("owned_by_type")); owned != "" {
		value, err := strconv.Atoi(owned)
		if err != nil {
			common.APIRespondWithError(c, http.StatusOK, errors.New("invalid owned_by_type"))
			return
		}
		query.OwnedByType = &value
	}

	if channel := strings.TrimSpace(c.Query("channel_type")); channel != "" {
		value, err := strconv.Atoi(channel)
		if err != nil {
			common.APIRespondWithError(c, http.StatusOK, errors.New("invalid channel_type"))
			return
		}
		query.ChannelType = &value
	}

	if issue := strings.TrimSpace(c.Query("issue")); issue != "" {
		rawIssues := strings.Split(issue, ",")
		for _, item := range rawIssues {
			if trimmed := strings.TrimSpace(item); trimmed != "" {
				query.Issues = append(query.Issues, trimmed)
			}
		}
	}

	if pageStr := strings.TrimSpace(c.Query("page")); pageStr != "" {
		if pageValue, err := strconv.Atoi(pageStr); err == nil {
			query.Page = pageValue
		}
	}

	if pageSizeStr := strings.TrimSpace(c.Query("page_size")); pageSizeStr != "" {
		if sizeValue, err := strconv.Atoi(pageSizeStr); err == nil {
			query.PageSize = sizeValue
		}
	}

	overview, err := model.QueryModelBrandOverview(query)
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    overview,
	})
}

func checkModelOwnedByReserveID(id int) bool {
	return id <= model.ModelOwnedByReserveID
}

type ModelBrandBatchUpdateRequest struct {
	Models      []string `json:"models" binding:"required"`
	OwnedByType *int     `json:"owned_by_type"`
}

func BatchUpdateModelBrand(c *gin.Context) {
	req := &ModelBrandBatchUpdateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	if len(req.Models) == 0 {
		common.APIRespondWithError(c, http.StatusOK, errors.New("models is required"))
		return
	}

	if req.OwnedByType == nil {
		common.APIRespondWithError(c, http.StatusOK, errors.New("owned_by_type is required"))
		return
	}

	if err := model.BatchUpdatePriceOwnedByType(req.Models, *req.OwnedByType); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}
