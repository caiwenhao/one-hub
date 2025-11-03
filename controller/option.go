package controller

import (
    "encoding/json"
    "net/http"
    "one-api/common/config"
    "one-api/common/utils"
    "one-api/model"
    "one-api/safty"
    "strings"

    "github.com/gin-gonic/gin"
)

// 前端固化配置项：不对外暴露写入能力
var frontendControlledOptionKeys = map[string]bool{
    "SystemName":       true,
    "Logo":             true,
    "HomePageContent":  true,
    "About":            true,
    "Footer":           true,
}

func GetOptions(c *gin.Context) {
    var options []*model.Option
    for k, v := range config.GlobalOption.GetAll() {
        if strings.HasSuffix(k, "Token") || strings.HasSuffix(k, "Secret") {
            continue
        }
        // 从选项列表中过滤前端固化配置，避免在前端显示/误改
        if frontendControlledOptionKeys[k] {
            continue
        }
        options = append(options, &model.Option{
            Key:   k,
            Value: utils.Interface2String(v),
        })
    }
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    options,
	})
	return
}

func GetSafeTools(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    safty.GetAllSafeToolsName(),
	})
	return
}

func UpdateOption(c *gin.Context) {
    var option model.Option
    err := json.NewDecoder(c.Request.Body).Decode(&option)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "无效的参数",
        })
        return
    }
    // 拦截前端固化配置项写入
    if frontendControlledOptionKeys[option.Key] {
        c.JSON(http.StatusOK, gin.H{
            "success": false,
            "message": "该配置项已通过前端代码固定，禁止后端修改",
        })
        return
    }
    switch option.Key {
    case "GitHubOAuthEnabled":
        if option.Value == "true" && config.GitHubClientId == "" {
            c.JSON(http.StatusOK, gin.H{
                "success": false,
                "message": "无法启用 GitHub OAuth，请先填入 GitHub Client Id 以及 GitHub Client Secret！",
			})
			return
		}
	case "OIDCAuthEnabled":
		if option.Value == "true" && (config.OIDCClientId == "" || config.OIDCClientSecret == "" || config.OIDCIssuer == "" || config.OIDCScopes == "" || config.OIDCUsernameClaims == "") {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用 OIDC，请先填入OIDC信息！",
			})
			return
		}
	case "EmailDomainRestrictionEnabled":
		if option.Value == "true" && len(config.EmailDomainWhitelist) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用邮箱域名限制，请先填入限制的邮箱域名！",
			})
			return
		}
	case "WeChatAuthEnabled":
		if option.Value == "true" && config.WeChatServerAddress == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用微信登录，请先填入微信登录相关配置信息！",
			})
			return
		}
	case "TurnstileCheckEnabled":
		if option.Value == "true" && config.TurnstileSiteKey == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用 Turnstile 校验，请先填入 Turnstile 校验相关配置信息！",
			})
			return
		}
	}
	err = model.UpdateOption(option.Key, option.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}
