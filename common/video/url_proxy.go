package video

import (
    "context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"one-api/common/config"
    "one-api/common/requester"
    "one-api/common/utils"
    "one-api/common/logger"
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
        logger.SysDebug("video.proxy: CFWorkerVideoUrl empty, return realURL")
        return realURL
    }

    // 空URL直接返回
    if strings.TrimSpace(realURL) == "" {
        logger.SysDebug("video.proxy: realURL empty, return as-is")
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

	// 构造代理URL基础：去除尾斜杠、补 https://
	base := strings.TrimSpace(config.CFWorkerVideoUrl)
	proxyURL := ensureHTTPS(strings.TrimSuffix(base, "/"))

	// 默认回退方案（无 KV 短链）：/v/<shorthex>.mp4?token=...
	sum := sha256.Sum256([]byte(encrypted))
	shortHex := fmt.Sprintf("%x", sum)[:12]
	fallback := fmt.Sprintf("%s/v/%s.mp4?token=%s", proxyURL, shortHex, url.QueryEscape(encrypted))

	// 未启用 KV 短链则直接返回回退样式
    if !config.CFWorkerVideoKVEnabled {
        logger.SysDebug(fmt.Sprintf("video.proxy: KV disabled -> fallback shortHex=%s", shortHex))
        return fallback
    }

	// 计算 slug（base62(sha256(token)) 截断）
	slug := hashToBase62(sum[:], clampShortLen(config.CFWorkerVideoShortLength))

	// 计算 ttl：尽量与 payload 的过期保持一致（秒）
	ttl := 24 * 3600
	now := time.Now().Unix()
	if exp := now + int64(ttl); exp > now { // 默认 24h，若后续调整 payload.ExpiresAt 可取最小值
		// no-op，保留模版
	}

	// 调用 Worker /register 写入 KV 映射
    if tryRegisterShortLink(proxyURL, slug, encrypted, ttl) == nil {
        logger.SysDebug(fmt.Sprintf("video.proxy: register success slug=%s ttl=%d worker=%s", slug, ttl, proxyURL))
        return fmt.Sprintf("%s/v/%s.mp4", proxyURL, slug)
    }

    // 注册失败则回退
    logger.SysError(fmt.Sprintf("video.proxy: register failed -> fallback shortHex=%s slug=%s worker=%s", shortHex, slug, proxyURL))
    return fallback
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

// --- 短链注册与工具 ---

// 将 sha256 摘要编码为 base62 并截断到指定长度
func hashToBase62(hash []byte, length int) string {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// 简易 base62：逐步把 bytes 视为大整数编码
	// 为避免大整数库，这里用常规除法实现，足够生成固定长度 slug
	// 复制一份可变切片
	buf := make([]byte, len(hash))
	copy(buf, hash)
	// 将字节转为大整数形式在 base256 下进行除法到 base62
	// 为简单与稳定，生成足够长度后截断
	out := make([]byte, 0, 44)
	digits := make([]byte, len(buf))
	for len(buf) > 0 && len(out) < 64 {
		// 模拟大整数除法：buf / 62
		var rem int
		n := 0
		for i := 0; i < len(buf); i++ {
			n = rem*256 + int(buf[i])
			q := n / 62
			rem = n % 62
			digits[i] = byte(q)
		}
		out = append(out, alphabet[rem])
		// 去掉前导 0
		i := 0
		for i < len(digits) && digits[i] == 0 {
			i++
		}
		buf = buf[:copy(buf, digits[i:])]
	}
	// 反转（最高位在末尾）
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	if length <= 0 || length > len(out) {
		length = len(out)
	}
	return string(out[:length])
}

func clampShortLen(n int) int {
	if n <= 0 {
		return 12
	}
	if n > 32 {
		return 32
	}
	return n
}

// 调用 Worker 的 /register 将 slug->token 写入 KV
func tryRegisterShortLink(base, slug, token string, ttl int) error {
    if strings.TrimSpace(base) == "" || strings.TrimSpace(slug) == "" || strings.TrimSpace(token) == "" {
        return errors.New("invalid params")
    }
    url := ensureHTTPS(strings.TrimSuffix(base, "/")) + "/register"
    body := map[string]any{
        "slug":  slug,
        "token": token,
        "ttl":   ttl,
    }
    headers := map[string]string{
        "Content-Type": "application/json",
        "X-API-Key":    config.CFWorkerVideoKey,
    }
    // 避免泄露敏感 token：仅记录长度与前后若干字符
    tokenPreview := ""
    if n := len(token); n > 16 {
        tokenPreview = fmt.Sprintf("%s...%s(len=%d)", token[:8], token[n-8:], n)
    } else {
        tokenPreview = fmt.Sprintf("%s(len=%d)", token, len(token))
    }
    logger.SysDebug(fmt.Sprintf("video.proxy: register begin url=%s slug=%s ttl=%d token=%s", url, slug, ttl, tokenPreview))
    // 使用带超时的上下文，避免阻塞主链路
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    req, err := utils.RequestBuilder(ctx, http.MethodPost, url, body, nil)
    if err != nil {
        logger.SysError(fmt.Sprintf("video.proxy: build request error slug=%s err=%v", slug, err))
        return err
    }
    for k, v := range headers {
        req.Header.Set(k, v)
    }
    resp, err := requester.HTTPClient.Do(req)
    if err != nil {
        logger.SysError(fmt.Sprintf("video.proxy: http do error slug=%s err=%v", slug, err))
        return err
    }
    defer resp.Body.Close()
    // 读取少量响应体做诊断
    b, _ := io.ReadAll(resp.Body)
    msg := string(b)
    if len(msg) > 512 {
        msg = msg[:512]
    }
    logger.SysDebug(fmt.Sprintf("video.proxy: register resp slug=%s status=%d body=%q", slug, resp.StatusCode, msg))
    if resp.StatusCode/100 != 2 {
        return fmt.Errorf("register failed: %d %s", resp.StatusCode, msg)
    }
    return nil
}
