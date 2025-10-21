package controller

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// collectStringArray 支持 name / name[] / name[0] 形式的数组参数
func collectStringArray(c *gin.Context, key string) []string {
	query := c.Request.URL.Query()
	result := make([]string, 0)
	seen := make(map[string]struct{})

	addValues := func(values []string) {
		for _, v := range values {
			if v == "" {
				continue
			}
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	if values, ok := query[key]; ok {
		addValues(values)
	}
	if values, ok := query[key+"[]"]; ok {
		addValues(values)
	}

	prefix := key + "["
	for k, values := range query {
		if strings.HasPrefix(k, prefix) {
			addValues(values)
		}
	}

	return result
}

func collectIntArray(c *gin.Context, key string) []int {
	stringVals := collectStringArray(c, key)
	result := make([]int, 0, len(stringVals))
	seen := make(map[int]struct{})
	for _, item := range stringVals {
		v, err := strconv.Atoi(item)
		if err != nil {
			continue
		}
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		result = append(result, v)
	}
	return result
}

func GetLogsList(c *gin.Context) {
	var params model.LogsListParams
	if err := c.ShouldBindQuery(&params); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	// 扩展：支持多选数组过滤
	modelNames := collectStringArray(c, "model_names")
	channelIDs := collectIntArray(c, "channel_ids")
	usernames := collectStringArray(c, "usernames")

	var (
		logs *model.DataResult[model.Log]
		err  error
	)
	if len(modelNames) > 0 || len(channelIDs) > 0 || len(usernames) > 0 {
		logs, err = model.GetLogsListWithArrays(&params, modelNames, channelIDs, usernames, nil)
	} else {
		logs, err = model.GetLogsList(&params)
	}
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    logs,
	})
}

func GetUserLogsList(c *gin.Context) {
	userId := c.GetInt("id")

	var params model.LogsListParams
	if err := c.ShouldBindQuery(&params); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
	// 扩展：支持多选数组过滤
	modelNames := collectStringArray(c, "model_names")
	channelIDs := collectIntArray(c, "channel_ids")
	// 用户列表不支持 usernames[]
	var logs *model.DataResult[model.Log]
	var err error
	if len(modelNames) > 0 || len(channelIDs) > 0 {
		logs, err = model.GetLogsListWithArrays(&params, modelNames, channelIDs, nil, &userId)
	} else {
		logs, err = model.GetUserLogsList(userId, &params)
	}
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    logs,
	})
}

// -------------------- 选项接口（供下拉多选搜索） --------------------

type optionItem struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

// GetLogOptions 管理员：模型名 / 渠道 / 用户名
func GetLogOptions(c *gin.Context) {
	field := strings.TrimSpace(c.Query("field"))
	q := strings.TrimSpace(c.Query("q"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit <= 0 || limit > 50 {
		limit = 20
	}

	var res []optionItem
	switch field {
	case "model_name":
		var rows []string
		tx := model.DB.Model(&model.Log{}).Distinct("model_name").Where("model_name <> ''")
		if q != "" {
			tx = tx.Where("model_name LIKE ?", "%"+q+"%")
		}
		if err := tx.Limit(limit).Order("model_name ASC").Pluck("model_name", &rows).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
		for _, m := range rows {
			res = append(res, optionItem{Label: m, Value: m})
		}
	case "channel_id":
		type row struct {
			Id   int
			Name string
		}
		var rows []row
		tx := model.DB.Model(&model.Channel{}).Select("id,name")
		if q != "" {
			tx = tx.Where("name LIKE ?", "%"+q+"%")
		}
		if err := tx.Limit(limit).Order("name ASC").Find(&rows).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
		for _, r := range rows {
			res = append(res, optionItem{Label: r.Name, Value: r.Id})
		}
	case "username":
		var rows []string
		tx := model.DB.Model(&model.User{}).Distinct("username")
		if q != "" {
			tx = tx.Where("username LIKE ?", "%"+q+"%")
		}
		if err := tx.Limit(limit).Order("username ASC").Pluck("username", &rows).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
		for _, u := range rows {
			res = append(res, optionItem{Label: u, Value: u})
		}
	default:
		c.JSON(http.StatusOK, gin.H{"success": true, "data": []optionItem{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": res})
}

// GetUserLogOptions 用户：模型名 / 渠道（限制为自身日志）
func GetUserLogOptions(c *gin.Context) {
	uid := c.GetInt("id")
	field := strings.TrimSpace(c.Query("field"))
	q := strings.TrimSpace(c.Query("q"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit <= 0 || limit > 50 {
		limit = 20
	}

	var res []optionItem
	switch field {
	case "model_name":
		var rows []string
		tx := model.DB.Model(&model.Log{}).Where("user_id = ?", uid).Distinct("model_name").Where("model_name <> ''")
		if q != "" {
			tx = tx.Where("model_name LIKE ?", "%"+q+"%")
		}
		if err := tx.Limit(limit).Order("model_name ASC").Pluck("model_name", &rows).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
		for _, m := range rows {
			res = append(res, optionItem{Label: m, Value: m})
		}
	case "channel_id":
		type row struct {
			Id   int
			Name string
		}
		var rows []row
		// 仅返回用户使用过的渠道
		sub := model.DB.Model(&model.Log{}).Select("DISTINCT channel_id").Where("user_id = ?", uid)
		if err := model.DB.Model(&model.Channel{}).Select("id,name").Where("id IN (?)", sub).Find(&rows).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
		for _, r := range rows {
			if q == "" || strings.Contains(strings.ToLower(r.Name), strings.ToLower(q)) {
				res = append(res, optionItem{Label: r.Name, Value: r.Id})
			}
		}
		if len(res) > limit {
			res = res[:limit]
		}
	default:
		c.JSON(http.StatusOK, gin.H{"success": true, "data": []optionItem{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": res})
}

func GetLogsStat(c *gin.Context) {
	// logType, _ := strconv.Atoi(c.Query("type"))
	startTimestamp, _ := strconv.ParseInt(c.Query("start_timestamp"), 10, 64)
	endTimestamp, _ := strconv.ParseInt(c.Query("end_timestamp"), 10, 64)
	tokenName := c.Query("token_name")
	username := c.Query("username")
	modelName := c.Query("model_name")
	channel, _ := strconv.Atoi(c.Query("channel"))
	quotaNum := model.SumUsedQuota(startTimestamp, endTimestamp, modelName, username, tokenName, channel)
	//tokenNum := model.SumUsedToken(logType, startTimestamp, endTimestamp, modelName, username, "")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"quota": quotaNum,
			//"token": tokenNum,
		},
	})
}

func GetLogsSelfStat(c *gin.Context) {
	username := c.GetString("username")
	// logType, _ := strconv.Atoi(c.Query("type"))
	startTimestamp, _ := strconv.ParseInt(c.Query("start_timestamp"), 10, 64)
	endTimestamp, _ := strconv.ParseInt(c.Query("end_timestamp"), 10, 64)
	tokenName := c.Query("token_name")
	modelName := c.Query("model_name")
	channel, _ := strconv.Atoi(c.Query("channel"))
	quotaNum := model.SumUsedQuota(startTimestamp, endTimestamp, modelName, username, tokenName, channel)
	//tokenNum := model.SumUsedToken(logType, startTimestamp, endTimestamp, modelName, username, tokenName)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"quota": quotaNum,
			//"token": tokenNum,
		},
	})
}

func DeleteHistoryLogs(c *gin.Context) {
	targetTimestamp, _ := strconv.ParseInt(c.Query("target_timestamp"), 10, 64)
	if targetTimestamp == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "target timestamp is required",
		})
		return
	}
	count, err := model.DeleteOldLog(targetTimestamp)
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
		"data":    count,
	})
}

// buildLogsExportQuery 按请求参数构建导出查询（支持单值与多值 IN 过滤）
func buildLogsExportQuery(c *gin.Context, base *gorm.DB, forUserID *int) *gorm.DB {
	// 起点
	tx := base

	// 仅当前用户
	if forUserID != nil {
		tx = tx.Where("user_id = ?", *forUserID)
	}

	// 类型
	if v := strings.TrimSpace(c.Query("log_type")); v != "" {
		if logType, err := strconv.Atoi(v); err == nil && logType != model.LogTypeUnknown {
			tx = tx.Where("type = ?", logType)
		}
	}

	// 月份参数优先：month=YYYY-MM
	if month := strings.TrimSpace(c.Query("month")); month != "" {
		// 解析月份
		t, err := time.ParseInLocation("2006-01", month, time.Local)
		if err == nil {
			start := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local).Unix()
			end := time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 0, time.Local).Unix()
			tx = tx.Where("created_at >= ?", start).Where("created_at <= ?", end)
		}
	} else {
		// start / end
		if v := strings.TrimSpace(c.Query("start_timestamp")); v != "" {
			if ts, err := strconv.ParseInt(v, 10, 64); err == nil && ts > 0 {
				tx = tx.Where("created_at >= ?", ts)
			}
		}
		if v := strings.TrimSpace(c.Query("end_timestamp")); v != "" {
			if ts, err := strconv.ParseInt(v, 10, 64); err == nil && ts > 0 {
				tx = tx.Where("created_at <= ?", ts)
			}
		}
	}

	// 等值与 IN 过滤
	// model_name / model_names[]
	modelNames := collectStringArray(c, "model_names")
	if len(modelNames) > 0 {
		tx = tx.Where("model_name IN ?", modelNames)
	} else if v := strings.TrimSpace(c.Query("model_name")); v != "" {
		tx = tx.Where("model_name = ?", v)
	}

	// token_name 单值
	if v := strings.TrimSpace(c.Query("token_name")); v != "" {
		tx = tx.Where("token_name = ?", v)
	}

	// channel_id / channel_ids[]
	channelIDs := collectIntArray(c, "channel_ids")
	if len(channelIDs) > 0 {
		tx = tx.Where("channel_id IN ?", channelIDs)
	} else if v := strings.TrimSpace(c.Query("channel_id")); v != "" {
		if id, err := strconv.Atoi(v); err == nil && id > 0 {
			tx = tx.Where("channel_id = ?", id)
		}
	}

	// username / usernames[]（仅管理员 route 会传 forUserID=nil 才允许生效）
	if forUserID == nil {
		usernames := collectStringArray(c, "usernames")
		if len(usernames) > 0 {
			tx = tx.Where("username IN ?", usernames)
		} else if v := strings.TrimSpace(c.Query("username")); v != "" {
			tx = tx.Where("username = ?", v)
		}
	}

	// source_ip
	if v := strings.TrimSpace(c.Query("source_ip")); v != "" {
		tx = tx.Where("source_ip = ?", v)
	}

	// 排序（简单兜底：created_at desc, id desc）
	orderStr := strings.TrimSpace(c.Query("order"))
	if orderStr == "" {
		tx = tx.Order("created_at DESC").Order("id DESC")
	} else {
		// 兼容前端传入 "-created_at" 或 "created_at" 格式
		for _, field := range strings.Split(orderStr, ",") {
			f := strings.TrimSpace(field)
			if f == "" {
				continue
			}
			desc := strings.HasPrefix(f, "-")
			if desc {
				f = f[1:]
			}
			// 白名单字段
			switch f {
			case "created_at", "channel_id", "user_id", "token_name", "model_name", "type", "source_ip":
				if desc {
					tx = tx.Order(f + " DESC")
				} else {
					tx = tx.Order(f)
				}
			}
		}
	}

	// 预加载渠道名称用于导出
	tx = tx.Preload("Channel", func(db *gorm.DB) *gorm.DB { return db.Select("id, name") })

	return tx
}

func quotaValuesForCSV(quota int) (raw string, currency string) {
	raw = strconv.Itoa(quota)
	if config.DisplayInCurrencyEnabled && config.QuotaPerUnit > 0 {
		currency = fmt.Sprintf("%.6f", float64(quota)/config.QuotaPerUnit)
	}
	return
}

// ExportLogsCSV 管理员导出日志 CSV（不包含输入/输出片段）
func ExportLogsCSV(c *gin.Context) {
	// 构建查询
	tx := buildLogsExportQuery(c, model.DB, nil)

	// 输出头
	filename := time.Now().In(time.Local).Format("logs_20060102_150405.csv")
	c.Header("Content-Type", "text/csv; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename="+filename)

	// 写入 UTF-8 BOM 以便 Excel 识别
	if _, err := c.Writer.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
		return
	}

	w := csv.NewWriter(c.Writer)
	header := []string{"时间", "用户", "渠道", "模型", "令牌"}
	if config.DisplayInCurrencyEnabled && config.QuotaPerUnit > 0 {
		header = append(header, "金额(货币)")
	}
	header = append(header, "请求耗时(s)", "来源IP", "流式")
	_ = w.Write(header)

	// 分批导出，避免占用过多内存
	const batchSize = 1000
	exported := 0
	for page := 0; ; page++ {
		var rows []*model.Log
		err := tx.Limit(batchSize).Offset(page * batchSize).Find(&rows).Error
		if err != nil {
			w.Flush()
			return
		}
		if len(rows) == 0 {
			break
		}
		for _, r := range rows {
			createdAt := time.Unix(r.CreatedAt, 0).In(time.Local).Format("2006-01-02 15:04:05")
			username := r.Username
			channel := ""
			if r.Channel != nil {
				channel = r.Channel.Name
			}
			isStream := "否"
			if r.IsStream {
				isStream = "是"
			}
			_, quotaCurrency := quotaValuesForCSV(r.Quota)

			rec := []string{
				createdAt,
				username,
				channel,
				r.ModelName,
				r.TokenName,
			}
			if quotaCurrency != "" {
				rec = append(rec, quotaCurrency)
			}
			rec = append(rec,
				strconv.Itoa(r.RequestTime),
				r.SourceIp,
				isStream,
			)
			_ = w.Write(rec)
			exported++
		}
		w.Flush()
		if err := w.Error(); err != nil {
			return
		}
		// 简单安全阈值，防止超大导出导致 OOM；可按需放宽或改为异步任务
		if exported > 1000000 { // 100 万行安全阈值
			break
		}
	}
}

// ExportUserLogsCSV 导出当前登录用户的日志 CSV
func ExportUserLogsCSV(c *gin.Context) {
	uid := c.GetInt("id")
	tx := buildLogsExportQuery(c, model.DB, &uid)

	filename := time.Now().In(time.Local).Format("logs_20060102_150405.csv")
	c.Header("Content-Type", "text/csv; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename="+filename)

	// UTF-8 BOM
	if _, err := c.Writer.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
		return
	}

	w := csv.NewWriter(c.Writer)
	header := []string{"时间", "用户", "渠道", "模型", "令牌"}
	if config.DisplayInCurrencyEnabled && config.QuotaPerUnit > 0 {
		header = append(header, "金额(货币)")
	}
	header = append(header, "请求耗时(s)", "来源IP", "流式")
	_ = w.Write(header)

	const batchSize = 1000
	for page := 0; ; page++ {
		var rows []*model.Log
		err := tx.Limit(batchSize).Offset(page * batchSize).Find(&rows).Error
		if err != nil {
			w.Flush()
			return
		}
		if len(rows) == 0 {
			break
		}
		for _, r := range rows {
			createdAt := time.Unix(r.CreatedAt, 0).In(time.Local).Format("2006-01-02 15:04:05")
			channel := ""
			if r.Channel != nil {
				channel = r.Channel.Name
			}
			isStream := "否"
			if r.IsStream {
				isStream = "是"
			}
			_, quotaCurrency := quotaValuesForCSV(r.Quota)
			rec := []string{
				createdAt,
				r.Username,
				channel,
				r.ModelName,
				r.TokenName,
			}
			if quotaCurrency != "" {
				rec = append(rec, quotaCurrency)
			}
			rec = append(rec,
				strconv.Itoa(r.RequestTime),
				r.SourceIp,
				isStream,
			)
			_ = w.Write(rec)
		}
		w.Flush()
		if err := w.Error(); err != nil {
			return
		}
	}
}
