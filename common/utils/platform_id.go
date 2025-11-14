package utils

import (
    mrand "math/rand"
    "strconv"
    "strings"
    "time"

    ulid "github.com/oklog/ulid/v2"
)

// Platform task id 编解码：task_<base36(lowercase)>
// 仅对内部自增 ID 做简单 base36 封装，满足短 ID 与可读性。

const PlatformTaskPrefix = "task_"

// EncodePlatformTaskID 将自增主键编码为平台任务ID（task_<base36>）
func EncodePlatformTaskID(id int64) string {
	if id <= 0 {
		return ""
	}
	s := strings.ToLower(strconv.FormatInt(id, 36))
	// 规范：最少 8 位，左侧 0 填充，避免早期自增很小导致过短（如 task_s）
	const minLen = 8
	if len(s) < minLen {
		s = strings.Repeat("0", minLen-len(s)) + s
	}
	return PlatformTaskPrefix + s
}

// NewPlatformULID 生成 ULID 字符串，长度 26，使用大写 Crockford Base32
func NewPlatformULID() string {
    t := time.Now()
    // math/rand 源用于 ulid 的熵，需种子
    ms := ulid.Timestamp(t)
    entropy := mrand.New(mrand.NewSource(t.UnixNano()))
    id := ulid.MustNew(ms, entropy)
    return strings.ToUpper(id.String())
}

// AddTaskPrefix 确保返回以 task_ 开头的展示ID
func AddTaskPrefix(id string) string {
    if id == "" {
        return ""
    }
    if strings.HasPrefix(strings.ToLower(id), PlatformTaskPrefix) {
        return id
    }
    return PlatformTaskPrefix + id
}

// StripTaskPrefix 去掉 task_ 前缀
func StripTaskPrefix(id string) string {
    if strings.HasPrefix(strings.ToLower(id), PlatformTaskPrefix) {
        return id[len(PlatformTaskPrefix):]
    }
    return id
}

// IsULID 校验是否为合法 ULID 字符串
func IsULID(s string) bool {
    if len(s) != 26 {
        return false
    }
    _, err := ulid.Parse(strings.ToLower(s))
    return err == nil
}

// DecodePlatformTaskID 解析平台任务ID，返回内部自增主键与是否成功
func DecodePlatformTaskID(platformID string) (int64, bool) {
	s := strings.TrimSpace(platformID)
	if s == "" {
		return 0, false
	}
	if !strings.HasPrefix(strings.ToLower(s), PlatformTaskPrefix) {
		return 0, false
	}
	raw := s[len(PlatformTaskPrefix):]
	if raw == "" {
		return 0, false
	}
	v, err := strconv.ParseInt(raw, 36, 64)
	if err != nil || v <= 0 {
		return 0, false
	}
	return v, true
}
