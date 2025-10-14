package main

import (
	"log"

	"github.com/pkoukk/tiktoken-go"
)

// 预热预定义模型的 tiktoken 词表缓存，避免运行时拉取
func main() {
	models := []string{
		"gpt-3.5-turbo",
		"gpt-4",
		"gpt-4o",
	}

	for _, model := range models {
		if _, err := tiktoken.EncodingForModel(model); err != nil {
			log.Fatalf("prefetch %s failed: %v", model, err)
		}
		log.Printf("prefetch %s success", model)
	}
}

