package video

import (
	"strings"
	"testing"
	"time"
)

func TestEncryptDecryptPayload(t *testing.T) {
	key := "test-secret-key-12345678"
	
	payload := VideoURLPayload{
		URL:       "https://example.com/video.mp4",
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		VideoID:   "test-video-123",
	}

	// 测试加密
	encrypted, err := encryptPayload(payload, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	if encrypted == "" {
		t.Fatal("Encrypted string is empty")
	}

	t.Logf("Encrypted token: %s", encrypted)

	// 测试解密
	decrypted, err := DecryptPayload(encrypted, key)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	// 验证解密结果
	if decrypted.URL != payload.URL {
		t.Errorf("URL mismatch: got %s, want %s", decrypted.URL, payload.URL)
	}

	if decrypted.VideoID != payload.VideoID {
		t.Errorf("VideoID mismatch: got %s, want %s", decrypted.VideoID, payload.VideoID)
	}

	if decrypted.ExpiresAt != payload.ExpiresAt {
		t.Errorf("ExpiresAt mismatch: got %d, want %d", decrypted.ExpiresAt, payload.ExpiresAt)
	}
}

func TestDecryptWithWrongKey(t *testing.T) {
	key1 := "correct-key"
	key2 := "wrong-key"

	payload := VideoURLPayload{
		URL:       "https://example.com/video.mp4",
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		VideoID:   "test-video-123",
	}

	// 使用 key1 加密
	encrypted, err := encryptPayload(payload, key1)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	// 使用 key2 解密（应该失败）
	_, err = DecryptPayload(encrypted, key2)
	if err == nil {
		t.Fatal("Expected decrypt to fail with wrong key, but it succeeded")
	}

	t.Logf("Correctly failed with wrong key: %v", err)
}

func TestExpiredToken(t *testing.T) {
	key := "test-key"

	// 创建已过期的payload
	payload := VideoURLPayload{
		URL:       "https://example.com/video.mp4",
		ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(), // 1小时前过期
		VideoID:   "test-video-123",
	}

	encrypted, err := encryptPayload(payload, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	// 解密应该失败（token已过期）
	_, err = DecryptPayload(encrypted, key)
	if err == nil {
		t.Fatal("Expected decrypt to fail with expired token, but it succeeded")
	}

	if !strings.Contains(err.Error(), "expired") {
		t.Errorf("Expected 'expired' error, got: %v", err)
	}

	t.Logf("Correctly detected expired token: %v", err)
}

func TestProxyVideoURL(t *testing.T) {
	// 测试未配置CF Workers的情况
	originalURL := "https://upstream.com/video.mp4"
	result := ProxyVideoURL(originalURL, "test-123")
	
	if result != originalURL {
		t.Errorf("Expected original URL when CF not configured, got: %s", result)
	}

	// 测试空URL
	result = ProxyVideoURL("", "test-123")
	if result != "" {
		t.Errorf("Expected empty string for empty URL, got: %s", result)
	}
}

func TestProxyVideoURLWithConfig(t *testing.T) {
	// 模拟配置（注意：这需要在实际环境中设置config包的变量）
	// 这里只是演示测试结构
	
	originalURL := "https://upstream.com/video.mp4"
	videoID := "test-video-123"
	
	// 在实际测试中，你需要：
	// 1. 设置 config.CFWorkerVideoUrl = "https://proxy.workers.dev"
	// 2. 设置 config.CFWorkerVideoKey = "test-key"
	// 3. 调用 ProxyVideoURL
	// 4. 验证返回的URL格式
	
	result := ProxyVideoURL(originalURL, videoID)
	
	// 如果配置了CF Workers，URL应该被替换
	// if !strings.Contains(result, "proxy.workers.dev") {
	//     t.Errorf("Expected proxy URL, got: %s", result)
	// }
	
	t.Logf("Proxy result: %s", result)
}

func BenchmarkEncryptPayload(b *testing.B) {
	key := "benchmark-key-12345678"
	payload := VideoURLPayload{
		URL:       "https://example.com/video.mp4",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		VideoID:   "bench-video-123",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encryptPayload(payload, key)
	}
}

func BenchmarkDecryptPayload(b *testing.B) {
	key := "benchmark-key-12345678"
	payload := VideoURLPayload{
		URL:       "https://example.com/video.mp4",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		VideoID:   "bench-video-123",
	}

	encrypted, _ := encryptPayload(payload, key)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptPayload(encrypted, key)
	}
}
