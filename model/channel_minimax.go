package model

import (
	"encoding/json"
	"strings"

	"one-api/common/config"
)

const (
	miniMaxOfficialBaseURL = "https://api.minimaxi.com"
	miniMaxPPInfraBaseURL  = "https://api.ppinfra.com"
)

// PopulateMiniMaxUpstream 从自定义参数或 BaseURL 推断上游供应商，用于回显到前端。
func (channel *Channel) PopulateMiniMaxUpstream() {
	if channel == nil || channel.Type != config.ChannelTypeMiniMax {
		return
	}

	if channel.MiniMaxUpstream != "" {
		channel.MiniMaxUpstream = strings.TrimSpace(channel.MiniMaxUpstream)
	}

	if channel.MiniMaxUpstream == "" {
		custom := channel.GetCustomParameter()
		if custom != "" {
			var payload map[string]json.RawMessage
			if err := json.Unmarshal([]byte(custom), &payload); err == nil {
				if videoRaw, ok := payload["video"]; ok {
					var video map[string]interface{}
					if err := json.Unmarshal(videoRaw, &video); err == nil {
						if upstream, ok := video["upstream"].(string); ok && strings.TrimSpace(upstream) != "" {
							channel.MiniMaxUpstream = strings.TrimSpace(upstream)
						}
					}
				}
				if channel.MiniMaxUpstream == "" {
					if topRaw, ok := payload["upstream"]; ok {
						var upstream string
						if err := json.Unmarshal(topRaw, &upstream); err == nil && strings.TrimSpace(upstream) != "" {
							channel.MiniMaxUpstream = strings.TrimSpace(upstream)
						}
					}
				}
			}
		}
	}

	if channel.MiniMaxUpstream == "" {
		if base := strings.TrimSpace(channel.GetBaseURL()); base != "" && strings.Contains(strings.ToLower(base), "ppinfra") {
			channel.MiniMaxUpstream = "ppinfra"
		} else {
			channel.MiniMaxUpstream = "official"
		}
	}
}

// ApplyMiniMaxUpstreamConfig 根据上游选择生成自定义参数、默认 BaseURL 等。
func (channel *Channel) ApplyMiniMaxUpstreamConfig() {
	if channel == nil || channel.Type != config.ChannelTypeMiniMax {
		return
	}

	upstream := strings.TrimSpace(channel.MiniMaxUpstream)
	if upstream == "" {
		upstream = "official"
	}

	var payload map[string]interface{}
	custom := channel.GetCustomParameter()
	if custom != "" {
		_ = json.Unmarshal([]byte(custom), &payload)
	}
	if payload == nil {
		payload = make(map[string]interface{})
	}

	video := extractMiniMaxVideoConfig(payload)
	video["upstream"] = upstream
	delete(video, "api_key")

	switch upstream {
	case "ppinfra":
		if _, ok := video["submit_path_template"]; !ok {
			video["submit_path_template"] = "/v3/async/%s"
		}
		if _, ok := video["query_path_template"]; !ok {
			if _, ok2 := video["query_path"]; !ok2 {
				video["query_path"] = "/v3/async/task-result"
			}
		}
		base := strings.TrimSpace(channel.GetBaseURL())
		if base == "" || !strings.Contains(strings.ToLower(base), "ppinfra") {
			channel.SetBaseURL(miniMaxPPInfraBaseURL)
		}
	default:
		if hasDefaultValue(video, "submit_path_template", "/v3/async/%s") {
			delete(video, "submit_path_template")
		}
		if hasDefaultValue(video, "query_path", "/v3/async/task-result") {
			delete(video, "query_path")
		}
		base := strings.TrimSpace(channel.GetBaseURL())
		if base == "" || strings.Contains(strings.ToLower(base), "ppinfra") {
			channel.SetBaseURL(miniMaxOfficialBaseURL)
		}
	}

	if len(video) == 0 {
		delete(payload, "video")
	} else {
		payload["video"] = video
	}

	delete(payload, "upstream")
	delete(payload, "api_key")

	if len(payload) == 0 {
		channel.CustomParameter = nil
	} else {
		if bytes, err := json.Marshal(payload); err == nil {
			str := string(bytes)
			channel.CustomParameter = &str
		}
	}
}

func extractMiniMaxVideoConfig(root map[string]interface{}) map[string]interface{} {
	if root == nil {
		return make(map[string]interface{})
	}
	if raw, ok := root["video"]; ok {
		switch v := raw.(type) {
		case map[string]interface{}:
			return v
		default:
			if bytes, err := json.Marshal(v); err == nil {
				var decoded map[string]interface{}
				if json.Unmarshal(bytes, &decoded) == nil {
					return decoded
				}
			}
		}
	}
	return make(map[string]interface{})
}

func hasDefaultValue(video map[string]interface{}, key, expected string) bool {
	if video == nil {
		return false
	}
	if val, ok := video[key]; ok {
		if str, ok2 := val.(string); ok2 {
			return strings.EqualFold(strings.TrimSpace(str), expected)
		}
	}
	return false
}

// SetBaseURL 为渠道设置 BaseURL 工具函数。
func (channel *Channel) SetBaseURL(url string) {
	trimmed := strings.TrimSpace(url)
	if trimmed == "" {
		channel.BaseURL = nil
		return
	}
	value := trimmed
	channel.BaseURL = &value
}
