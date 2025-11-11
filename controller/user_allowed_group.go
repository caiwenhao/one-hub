package controller

import (
    "net/http"
    "one-api/common"
    "one-api/model"
    "strconv"

    "github.com/gin-gonic/gin"
)

type AllowedGroupsPayload struct {
    Groups []string `json:"groups"`
}

// GetSelfAllowedGroups 返回当前用户被授权的分组列表
func GetSelfAllowedGroups(c *gin.Context) {
    userId := c.GetInt("id")
    groups, err := model.GetAllowedGroupsByUser(userId)
    if err != nil {
        common.APIRespondWithError(c, http.StatusOK, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": groups})
}

// Admin: 获取某用户的授权分组列表
func GetUserAllowedGroups(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
        return
    }
    groups, err := model.GetAllowedGroupsByUser(id)
    if err != nil {
        common.APIRespondWithError(c, http.StatusOK, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": groups})
}

// Admin: 覆盖设置某用户授权分组
func SetUserAllowedGroups(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
        return
    }
    payload := AllowedGroupsPayload{}
    if err := c.ShouldBindJSON(&payload); err != nil {
        common.APIRespondWithError(c, http.StatusOK, err)
        return
    }
    if err := model.ReplaceAllowedGroups(id, payload.Groups); err != nil {
        common.APIRespondWithError(c, http.StatusOK, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

