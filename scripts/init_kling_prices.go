package main

import (
	"encoding/json"
	"fmt"
	"log"
	"one-api/model"
)

// KlingPrice 可灵AI价格配置
type KlingPrice struct {
	Model       string  `json:"model"`
	Type        string  `json:"type"`
	ChannelType int     `json:"channel_type"`
	Input       float64 `json:"input"`
	Output      float64 `json:"output"`
}

// initKlingPrices 初始化可灵AI模型价格
func initKlingPrices() error {
	// 视频生成模型价格表
	videoPrices := []KlingPrice{
		// V1模型
		{"kling-video_kling-v1_std_5", "times", 53, 1.0, 1.0},
		{"kling-video_kling-v1_std_10", "times", 53, 2.0, 2.0},
		{"kling-video_kling-v1_pro_5", "times", 53, 3.5, 3.5},
		{"kling-video_kling-v1_pro_10", "times", 53, 7.0, 7.0},
		
		// V1.5模型
		{"kling-video_kling-v1-5_std_5", "times", 53, 2.0, 2.0},
		{"kling-video_kling-v1-5_std_10", "times", 53, 4.0, 4.0},
		{"kling-video_kling-v1-5_pro_5", "times", 53, 3.5, 3.5},
		{"kling-video_kling-v1-5_pro_10", "times", 53, 7.0, 7.0},
		
		// V1.6模型
		{"kling-video_kling-v1-6_std_5", "times", 53, 2.0, 2.0},
		{"kling-video_kling-v1-6_std_10", "times", 53, 4.0, 4.0},
		{"kling-video_kling-v1-6_pro_5", "times", 53, 3.5, 3.5},
		{"kling-video_kling-v1-6_pro_10", "times", 53, 7.0, 7.0},
		
		// V2大师版模型
		{"kling-video_kling-v2-master_5", "times", 53, 10.0, 10.0},
		{"kling-video_kling-v2-master_10", "times", 53, 20.0, 20.0},
		
		// V2.1模型
		{"kling-video_kling-v2-1_std_5", "times", 53, 2.0, 2.0},
		{"kling-video_kling-v2-1_std_10", "times", 53, 4.0, 4.0},
		{"kling-video_kling-v2-1_pro_5", "times", 53, 3.5, 3.5},
		{"kling-video_kling-v2-1_pro_10", "times", 53, 7.0, 7.0},
		{"kling-video_kling-v2-1-master_5", "times", 53, 10.0, 10.0},
		{"kling-video_kling-v2-1-master_10", "times", 53, 20.0, 20.0},
	}

	// 图像生成模型价格表
	imagePrices := []KlingPrice{
		// V1.0模型
		{"kling-image_kling-v1_text2image", "times", 53, 1.0, 1.0},
		{"kling-image_kling-v1_image2image", "times", 53, 1.0, 1.0},
		
		// V1.5模型
		{"kling-image_kling-v1-5_text2image", "times", 53, 4.0, 4.0},
		{"kling-image_kling-v1-5_image2image", "times", 53, 8.0, 8.0},
		
		// V2.0模型
		{"kling-image_kling-v2_text2image", "times", 53, 4.0, 4.0},
		{"kling-image_kling-v2_image2image", "times", 53, 8.0, 8.0},
		{"kling-image_kling-v2-new_image2image", "times", 53, 8.0, 8.0},
		{"kling-image_kling-v2_multi_reference", "times", 53, 16.0, 16.0},
		
		// V2.1模型
		{"kling-image_kling-v2-1_text2image", "times", 53, 4.0, 4.0},
		
		// 图像编辑
		{"kling-image_expand", "times", 53, 8.0, 8.0},
	}

	// 虚拟试穿模型价格表
	virtualTryOnPrices := []KlingPrice{
		{"kling-virtual-try-on_v1", "times", 53, 1.0, 1.0},
		{"kling-virtual-try-on_v1-5", "times", 53, 1.0, 1.0},
	}

	// 合并所有价格
	allPrices := append(videoPrices, imagePrices...)
	allPrices = append(allPrices, virtualTryOnPrices...)

	// 转换为model.Price格式并插入数据库
	for _, klingPrice := range allPrices {
		price := &model.Price{
			Model:       klingPrice.Model,
			Type:        klingPrice.Type,
			ChannelType: klingPrice.ChannelType,
			Input:       klingPrice.Input,
			Output:      klingPrice.Output,
			Locked:      false, // 允许后续更新
		}

		// 检查是否已存在
		existingPrice := model.PricingInstance.GetPrice(price.Model)
		if existingPrice != nil {
			fmt.Printf("模型 %s 价格已存在，跳过\n", price.Model)
			continue
		}

		// 添加价格
		if err := model.PricingInstance.AddPrice(price); err != nil {
			log.Printf("添加模型 %s 价格失败: %v", price.Model, err)
		} else {
			fmt.Printf("成功添加模型 %s 价格: input=%.1f, output=%.1f\n", 
				price.Model, price.Input, price.Output)
		}
	}

	return nil
}

// exportKlingPricesJSON 导出价格配置为JSON文件
func exportKlingPricesJSON() error {
	prices := model.PricingInstance.GetAllPrices()
	var klingPrices []KlingPrice

	// 筛选可灵AI相关模型
	for modelName, price := range prices {
		if price.ChannelType == 53 { // 可灵AI渠道类型
			klingPrices = append(klingPrices, KlingPrice{
				Model:       modelName,
				Type:        price.Type,
				ChannelType: price.ChannelType,
				Input:       price.Input,
				Output:      price.Output,
			})
		}
	}

	// 输出JSON
	jsonData, err := json.MarshalIndent(klingPrices, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println("可灵AI模型价格配置:")
	fmt.Println(string(jsonData))
	return nil
}

func main() {
	fmt.Println("=== 可灵AI模型价格初始化工具 ===")
	
	// 这里应该初始化数据库连接
	// model.SetupDB()
	// defer model.CloseDB()
	
	fmt.Println("1. 初始化可灵AI模型价格...")
	if err := initKlingPrices(); err != nil {
		log.Fatal("初始化价格失败:", err)
	}
	
	fmt.Println("\n2. 导出价格配置...")
	if err := exportKlingPricesJSON(); err != nil {
		log.Fatal("导出配置失败:", err)
	}
	
	fmt.Println("\n✅ 可灵AI模型价格初始化完成！")
}