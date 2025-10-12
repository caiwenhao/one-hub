package model

import (
	"one-api/common/config"
	"sort"
	"strings"
)

const (
	IssueOwnedByMissing       = "OWNED_BY_MISSING"
	IssuePriceChannelMismatch = "PRICE_CHANNEL_MISMATCH"
	IssueChannelBrandConflict = "CHANNEL_BRAND_CONFLICT"
	IssuePriceChannelUnknown  = "PRICE_CHANNEL_UNKNOWN"
	IssueNoChannelBound       = "NO_CHANNEL_BOUND"
)

type ModelBrandChannel struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ChannelType     int    `json:"channel_type"`
	ChannelTypeName string `json:"channel_type_name"`
	Group           string `json:"group"`
	Priority        int    `json:"priority"`
	Weight          uint   `json:"weight"`
	Status          int    `json:"status"`
}

type ModelBrandOverview struct {
	Model           string              `json:"model"`
	OwnedByType     int                 `json:"owned_by_type"`
	OwnedByName     string              `json:"owned_by_name"`
	ChannelType     int                 `json:"channel_type"`
	ChannelTypeName string              `json:"channel_type_name"`
	Groups          []string            `json:"groups"`
	Channels        []ModelBrandChannel `json:"channels"`
	Issues          []string            `json:"issues"`
	Price           *Price              `json:"price"`
}

type ModelBrandOverviewQuery struct {
	Keyword     string
	Group       string
	OwnedByType *int
	ChannelType *int
	Issues      []string
	Page        int
	PageSize    int
}

type ModelBrandOverviewResult struct {
	List       []*ModelBrandOverview `json:"list"`
	Total      int                   `json:"total"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"page_size"`
	IssueStats map[string]int        `json:"issue_stats"`
	Groups     []string              `json:"groups"`
}

func QueryModelBrandOverview(query *ModelBrandOverviewQuery) (*ModelBrandOverviewResult, error) {
	if query == nil {
		query = &ModelBrandOverviewQuery{}
	}

	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	if query.PageSize > 200 {
		query.PageSize = 200
	}
	if query.Page <= 0 {
		query.Page = 1
	}

	allOverview, groups, err := buildModelBrandOverview()
	if err != nil {
		return nil, err
	}

	filtered := filterModelBrandOverview(allOverview, query)

	total := len(filtered)
	start := (query.Page - 1) * query.PageSize
	if start > total {
		start = total
	}

	end := start + query.PageSize
	if end > total {
		end = total
	}

	list := filtered[start:end]

	issueStats := make(map[string]int)
	for _, item := range filtered {
		for _, issue := range item.Issues {
			issueStats[issue]++
		}
	}

	result := &ModelBrandOverviewResult{
		List:       list,
		Total:      total,
		Page:       query.Page,
		PageSize:   query.PageSize,
		IssueStats: issueStats,
		Groups:     groups,
	}

	return result, nil
}

func buildModelBrandOverview() ([]*ModelBrandOverview, []string, error) {
    var result []*ModelBrandOverview

    ownedByNames := ModelOwnedBysInstance.GetAll()
    modelGroups := ChannelGroup.GetModelsGroups()
    modelChannels := buildModelChannels(ownedByNames)

    // 仅处理“已接入渠道”的模型：以已构建出的 modelChannels 为准
    // 不再包含未绑定任何渠道的模型
    modelSet := map[string]struct{}{}
    for name := range modelChannels {
        modelSet[name] = struct{}{}
    }

    allModels := make([]string, 0, len(modelSet))
    for modelName := range modelSet {
        allModels = append(allModels, modelName)
    }
	sort.Strings(allModels)

	for _, modelName := range allModels {
		price := PricingInstance.GetPrice(modelName)
		price.Normalize()

		ownedByType := price.GetOwnedByType()
		channelType := price.ChannelType

		ownedByName := lookupOwnedByName(ownedByType, ownedByNames)
		channelTypeName := lookupOwnedByName(channelType, ownedByNames)

		groups := sortedGroups(modelGroups[modelName])
		channels := modelChannels[modelName]
		issues := inspectModelIssues(price, ownedByType, channels)

		item := &ModelBrandOverview{
			Model:           modelName,
			OwnedByType:     ownedByType,
			OwnedByName:     ownedByName,
			ChannelType:     channelType,
			ChannelTypeName: channelTypeName,
			Groups:          groups,
			Channels:        channels,
			Issues:          issues,
			Price:           price,
		}

		result = append(result, item)
	}

	groupSet := make(map[string]struct{})
	for _, overview := range result {
		for _, group := range overview.Groups {
			groupSet[group] = struct{}{}
		}
	}

	groups := make([]string, 0, len(groupSet))
	for group := range groupSet {
		groups = append(groups, group)
	}
	sort.Strings(groups)

	return result, groups, nil
}

func buildModelChannels(ownedByNames map[int]*ModelOwnedBy) map[string][]ModelBrandChannel {
	modelChannels := make(map[string][]ModelBrandChannel)

	ChannelGroup.RLock()
	defer ChannelGroup.RUnlock()

	for group, models := range ChannelGroup.Rule {
		for modelName, priorityList := range models {
			for priorityIdx, channelIDs := range priorityList {
				for _, channelID := range channelIDs {
					choice, ok := ChannelGroup.Channels[channelID]
					if !ok || choice.Channel == nil {
						continue
					}
					channel := choice.Channel

					weight := config.DefaultChannelWeight
					if channel.Weight != nil {
						weight = *channel.Weight
					}

					channelTypeName := lookupOwnedByName(channel.Type, ownedByNames)

					item := ModelBrandChannel{
						ID:              channel.Id,
						Name:            channel.Name,
						ChannelType:     channel.Type,
						ChannelTypeName: channelTypeName,
						Group:           group,
						Priority:        priorityIdx,
						Weight:          weight,
						Status:          channel.Status,
					}

					modelChannels[modelName] = append(modelChannels[modelName], item)
				}
			}
		}
	}

	for modelName := range modelChannels {
		sort.Slice(modelChannels[modelName], func(i, j int) bool {
			left := modelChannels[modelName][i]
			right := modelChannels[modelName][j]

			if left.Priority == right.Priority {
				return left.ID < right.ID
			}
			return left.Priority < right.Priority
		})
	}

	return modelChannels
}

func inspectModelIssues(price *Price, ownedByType int, channels []ModelBrandChannel) []string {
	var issues []string

	if ownedByType == config.ChannelTypeUnknown {
		issues = append(issues, IssueOwnedByMissing)
	}

	if price.ChannelType == config.ChannelTypeUnknown {
		issues = append(issues, IssuePriceChannelUnknown)
	} else if ownedByType != config.ChannelTypeUnknown && price.ChannelType != ownedByType {
		issues = append(issues, IssuePriceChannelMismatch)
	}

	if len(channels) == 0 {
		issues = append(issues, IssueNoChannelBound)
	}

	if ownedByType != config.ChannelTypeUnknown {
		for _, ch := range channels {
			if ch.ChannelType != ownedByType && ch.ChannelType != config.ChannelTypeUnknown {
				issues = append(issues, IssueChannelBrandConflict)
				break
			}
		}
	}

	return uniqueStrings(issues)
}

func filterModelBrandOverview(items []*ModelBrandOverview, query *ModelBrandOverviewQuery) []*ModelBrandOverview {
	if query == nil {
		return items
	}

	keyword := strings.TrimSpace(strings.ToLower(query.Keyword))
	groupFilter := strings.TrimSpace(query.Group)

	issueFilter := make([]string, 0, len(query.Issues))
	for _, issue := range query.Issues {
		if trimmed := strings.TrimSpace(issue); trimmed != "" {
			issueFilter = append(issueFilter, trimmed)
		}
	}

	result := make([]*ModelBrandOverview, 0, len(items))
	for _, item := range items {
		if keyword != "" {
			hit := strings.Contains(strings.ToLower(item.Model), keyword) ||
				strings.Contains(strings.ToLower(item.OwnedByName), keyword) ||
				strings.Contains(strings.ToLower(item.ChannelTypeName), keyword)
			if !hit {
				continue
			}
		}

		if query.OwnedByType != nil && item.OwnedByType != *query.OwnedByType {
			continue
		}

		if query.ChannelType != nil && item.ChannelType != *query.ChannelType {
			continue
		}

		if groupFilter != "" && !containsString(item.Groups, groupFilter) {
			continue
		}

		if len(issueFilter) > 0 {
			if !containsAnyString(issueFilter, item.Issues) {
				continue
			}
		}

		result = append(result, item)
	}

	return result
}

func sortedGroups(groups map[string]bool) []string {
	if len(groups) == 0 {
		return []string{}
	}
	result := make([]string, 0, len(groups))
	for group := range groups {
		result = append(result, group)
	}
	sort.Strings(result)
	return result
}

func lookupOwnedByName(channelType int, ownedByNames map[int]*ModelOwnedBy) string {
	if channelType == config.ChannelTypeUnknown {
		return UnknownOwnedBy
	}
	if owned, ok := ownedByNames[channelType]; ok && owned != nil && strings.TrimSpace(owned.Name) != "" {
		return owned.Name
	}
	name := ModelOwnedBysInstance.GetName(channelType)
	if strings.TrimSpace(name) == "" {
		return UnknownOwnedBy
	}
	return name
}

func uniqueStrings(items []string) []string {
	if len(items) == 0 {
		return items
	}
	seen := make(map[string]struct{}, len(items))
	result := make([]string, 0, len(items))
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}

func containsString(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func containsAnyString(options []string, values []string) bool {
	for _, value := range values {
		for _, option := range options {
			if value == option {
				return true
			}
		}
	}
	return false
}
