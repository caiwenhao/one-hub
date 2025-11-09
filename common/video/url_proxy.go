package video

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"one-api/common/config"
	"strings"
	"time"
)

// VideoURLPayload 视频URL加密载荷
type VideoURLPayload struct {
	URL       string `json:"url"`        // 真实视频URL
	ExpiresAt int64  `json:"expires_at"` // 过期时间戳
	VideoID   string `json:"video_id"`   // 视频ID（可选，用于日志追踪）
}

// ProxyVideoURL 将真实的视频URL转换为CF Workers代理URL
// 如果未配置CF Workers，则返回原始URL
func ProxyVideoURL(realURL, videoID string) string {
	// 未配置CF Workers，返回原URL
	if config.CFWorkerVideoUrl == "" {
		return realURL
	}

	// 空URL直接返回
	if strings.TrimSpace(realURL) == "" {
		return realURL
	}

	// 构造加密载荷
	payload := VideoURLPayload{
		URL:       realURL,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 24小时过期
		VideoID:   videoID,
	}

	// 加密载荷
	encrypted, err := encryptPayload(payload, config.CFWorkerVideoKey)
	if err != nil {
		// 加密失败，返回原URL（降级处理）
		return realURL
	}

    // 构造代理URL：形如 https://worker.example.com/v/<short>.mp4?token=<long>
    // - 文件名短一些，token 放在查询参数中
    // - 仅 Proxy 模式，确保看起来像视频直链
    // 规范化域名，去除空格与尾部斜杠，并确保有 https:// 前缀
    base := strings.TrimSpace(config.CFWorkerVideoUrl)
    proxyURL := ensureHTTPS(strings.TrimSuffix(base, "/"))
    // 生成短ID（基于 token 的 sha256 前缀）
    sum := sha256.Sum256([]byte(encrypted))
    short := fmt.Sprintf("%x", sum)[:12]
    return fmt.Sprintf("%s/v/%s.mp4?token=%s", proxyURL, short, url.QueryEscape(encrypted))
}

// encryptPayload 使用AES-256-GCM加密载荷
func encryptPayload(payload VideoURLPayload, key string) (string, error) {
	// 将载荷序列化为JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload failed: %w", err)
	}

	// 使用SHA256将密钥转换为32字节
	keyHash := sha256.Sum256([]byte(key))
	
	// 创建AES cipher
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return "", fmt.Errorf("create cipher failed: %w", err)
	}

	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("create gcm failed: %w", err)
	}

	// 生成随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("generate nonce failed: %w", err)
	}

	// 加密数据
	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)

    // Base64编码（URL安全、去除 padding 以便用于路径/查询）
    encoded := base64.RawURLEncoding.EncodeToString(ciphertext)
    return encoded, nil
}

// DecryptPayload 解密载荷（用于测试或服务端验证）
func DecryptPayload(encrypted, key string) (*VideoURLPayload, error) {
    // Base64解码：优先无 padding 的 RawURLEncoding，失败则回退标准 URLEncoding
    ciphertext, err := base64.RawURLEncoding.DecodeString(encrypted)
    if err != nil {
        if ct2, err2 := base64.URLEncoding.DecodeString(encrypted); err2 == nil {
            ciphertext = ct2
        } else {
            return nil, fmt.Errorf("base64 decode failed: %w", err)
        }
    }

	// 使用SHA256将密钥转换为32字节
	keyHash := sha256.Sum256([]byte(key))

	// 创建AES cipher
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return nil, fmt.Errorf("create cipher failed: %w", err)
	}

	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create gcm failed: %w", err)
	}

	// 提取nonce
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 解密
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decrypt failed: %w", err)
	}

	// 反序列化
	var payload VideoURLPayload
	if err := json.Unmarshal(plaintext, &payload); err != nil {
		return nil, fmt.Errorf("unmarshal payload failed: %w", err)
	}

	// 检查过期时间
	if payload.ExpiresAt > 0 && time.Now().Unix() > payload.ExpiresAt {
		return nil, fmt.Errorf("token expired")
	}

    return &payload, nil
}

// ensureHTTPS 若缺少协议则补齐 https://
func ensureHTTPS(u string) string {
    t := strings.TrimSpace(u)
    s := strings.ToLower(t)
    if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
        return t
    }
    return "https://" + t
}
