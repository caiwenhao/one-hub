package model

import (
	"encoding/json"
	"one-api/common/config"
	"one-api/common/logger"
	"strconv"
	"strings"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func removeKeyIndexMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202405152141",
		Migrate: func(tx *gorm.DB) error {
			dialect := tx.Dialector.Name()
			if dialect == "sqlite" {
				return nil
			}

			if !tx.Migrator().HasIndex(&Channel{}, "idx_channels_key") {
				return nil
			}

			err := tx.Migrator().DropIndex(&Channel{}, "idx_channels_key")
			if err != nil {
				logger.SysLog("remove idx_channels_key  Failure: " + err.Error())
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}

func changeTokenKeyColumnType() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202411300001",
		Migrate: func(tx *gorm.DB) error {
			// 如果表不存在，说明是新数据库，直接跳过
			if !tx.Migrator().HasTable("tokens") {
				return nil
			}

			dialect := tx.Dialector.Name()
			var err error

			switch dialect {
			case "mysql":
				err = tx.Exec("ALTER TABLE tokens MODIFY COLUMN `key` varchar(59)").Error
			case "postgres":
				err = tx.Exec("ALTER TABLE tokens ALTER COLUMN key TYPE varchar(59)").Error
			case "sqlite":
				return nil
			}

			if err != nil {
				logger.SysLog("修改 tokens.key 字段类型失败: " + err.Error())
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if !tx.Migrator().HasTable("tokens") {
				return nil
			}

			dialect := tx.Dialector.Name()
			var err error

			switch dialect {
			case "mysql":
				err = tx.Exec("ALTER TABLE tokens MODIFY COLUMN `key` char(48)").Error
			case "postgres":
				err = tx.Exec("ALTER TABLE tokens ALTER COLUMN key TYPE char(48)").Error
			}
			return err
		},
	}
}

func addOwnedByTypeToPrice() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202411300002",
		Migrate: func(tx *gorm.DB) error {
			if !tx.Migrator().HasTable(&Price{}) {
				return nil
			}

			if !tx.Migrator().HasColumn(&Price{}, "owned_by_type") {
				if err := tx.Migrator().AddColumn(&Price{}, "owned_by_type"); err != nil {
					logger.SysLog("新增 owned_by_type 字段失败: " + err.Error())
					return err
				}
			}

			if err := tx.Model(&Price{}).Where("owned_by_type = ?", 0).Update("owned_by_type", gorm.Expr("channel_type")).Error; err != nil {
				logger.SysLog("初始化价格品牌字段失败: " + err.Error())
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if !tx.Migrator().HasTable(&Price{}) || !tx.Migrator().HasColumn(&Price{}, "owned_by_type") {
				return nil
			}
			return tx.Migrator().DropColumn(&Price{}, "owned_by_type")
		},
	}
}

func migrationBefore(db *gorm.DB) error {
	// 从库不执行
	if !config.IsMasterNode {
		logger.SysLog("从库不执行迁移前操作")
		return nil
	}

	// 如果是第一次运行 直接跳过
	if !db.Migrator().HasTable("channels") {
		return nil
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		removeKeyIndexMigration(),
		changeTokenKeyColumnType(),
		addOwnedByTypeToPrice(),
	})
	return m.Migrate()
}

func addStatistics() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202408100001",
		Migrate: func(tx *gorm.DB) error {
			go UpdateStatistics(StatisticsUpdateTypeALL)
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}

func changeChannelApiVersion() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202408190001",
		Migrate: func(tx *gorm.DB) error {
			plugin := `{"customize": {"1": "{version}/chat/completions", "2": "{version}/completions", "3": "{version}/embeddings", "4": "{version}/moderations", "5": "{version}/images/generations", "6": "{version}/images/edits", "7": "{version}/images/variations", "9": "{version}/audio/speech", "10": "{version}/audio/transcriptions", "11": "{version}/audio/translations"}}`

			// 查询 channel 表中的type 为 8，且 other = disable 的数据,直接更新
			var jsonMap map[string]map[string]interface{}
			err := json.Unmarshal([]byte(strings.Replace(plugin, "{version}", "", -1)), &jsonMap)
			if err != nil {
				logger.SysLog("changeChannelApiVersion Failure: " + err.Error())
				return err
			}
			disableApi := map[string]interface{}{
				"other":  "",
				"plugin": datatypes.NewJSONType(jsonMap),
			}

			err = tx.Model(&Channel{}).Where("type = ? AND other = ?", 8, "disable").Updates(disableApi).Error
			if err != nil {
				logger.SysLog("changeChannelApiVersion Failure: " + err.Error())
				return err
			}

			// 查询 channel 表中的type 为 8，且 other != disable 并且不为空 的数据,直接更新
			var channels []*Channel
			err = tx.Model(&Channel{}).Where("type = ? AND other != ? AND other != ?", 8, "disable", "").Find(&channels).Error
			if err != nil {
				logger.SysLog("changeChannelApiVersion Failure: " + err.Error())
				return err
			}

			for _, channel := range channels {
				var jsonMap map[string]map[string]interface{}
				err := json.Unmarshal([]byte(strings.Replace(plugin, "{version}", "/"+channel.Other, -1)), &jsonMap)
				if err != nil {
					logger.SysLog("changeChannelApiVersion Failure: " + err.Error())
					return err
				}
				changeApi := map[string]interface{}{
					"other":  "",
					"plugin": datatypes.NewJSONType(jsonMap),
				}
				err = tx.Model(&Channel{}).Where("id = ?", channel.Id).Updates(changeApi).Error
				if err != nil {
					logger.SysLog("changeChannelApiVersion Failure: " + err.Error())
					return err
				}
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	}
}

func initUserGroup() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202410300001",
		Migrate: func(tx *gorm.DB) error {
			userGroups := map[string]*UserGroup{
				"default": {
					Symbol: "default",
					Name:   "默认分组",
					Ratio:  1,
					Public: true,
				},
				"vip": {
					Symbol: "vip",
					Name:   "vip分组",
					Ratio:  1,
					Public: false,
				},
				"svip": {
					Symbol: "svip",
					Name:   "svip分组",
					Ratio:  1,
					Public: false,
				},
			}
			option, err := GetOption("GroupRatio")
			if err == nil && option.Value != "" {
				oldGroup := make(map[string]float64)
				err = json.Unmarshal([]byte(option.Value), &oldGroup)
				if err != nil {
					return err
				}

				for k, v := range oldGroup {
					isPublic := false
					if k == "default" {
						isPublic = true
					}
					userGroups[k] = &UserGroup{
						Symbol: k,
						Name:   k,
						Ratio:  v,
						Public: isPublic,
					}
				}
			}

			for k, v := range userGroups {
				err := tx.Where("symbol = ?", k).FirstOrCreate(v).Error
				if err != nil {
					return err
				}
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	}
}

func addOldTokenMaxId() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202411300002",
		Migrate: func(tx *gorm.DB) error {
			var token Token
			tx.Last(&token)
			tokenMaxId := token.Id
			option := Option{
				Key: "OldTokenMaxId",
			}

			DB.FirstOrCreate(&option, Option{Key: "OldTokenMaxId"})
			option.Value = strconv.Itoa(tokenMaxId)
			return DB.Save(&option).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	}
}

func addExtraRatios() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202504300001",
		Migrate: func(tx *gorm.DB) error {
			extraTokenPriceJson := ""
			extraRatios := make(map[string]map[string]float64)
			// 先查询数据库中是否存在extra_ratios
			option, err := GetOption("ExtraTokenPriceJson")
			if err == nil {
				extraTokenPriceJson = option.Value

			} else {
				extraTokenPriceJson = GetDefaultExtraRatio()
			}

			err = json.Unmarshal([]byte(extraTokenPriceJson), &extraRatios)
			if err != nil {
				return err
			}

			if len(extraRatios) == 0 {
				return nil
			}

			models := make([]string, 0)
			for model := range extraRatios {
				models = append(models, model)
			}

			// 查询数据库中是否存在
			var prices []*Price
			err = tx.Where("model IN (?)", models).Find(&prices).Error
			if err != nil {
				return err
			}

			for _, price := range prices {
				extraRatios := extraRatios[price.Model]
				jsonData := datatypes.NewJSONType(extraRatios)
				price.ExtraRatios = &jsonData
				err = tx.Model(&Price{}).Where("model = ?", price.Model).Updates(map[string]interface{}{
					"extra_ratios": jsonData,
				}).Error
				if err != nil {
					return err
				}
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Rollback().Error
		},
	}
}
func migrationAfter(db *gorm.DB) error {
	// 从库不执行
	if !config.IsMasterNode {
		logger.SysLog("从库不执行迁移后操作")
		return nil
	}
    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        addStatistics(),
        changeChannelApiVersion(),
        initUserGroup(),
        addOldTokenMaxId(),
        addExtraRatios(),
        fixSutuiSoraTimesPricing(),
        fixAllTimesPricingIO(),
        fixSutuiVeo3TimesPricing(),
        fixVeo3OwnedByToGemini(),
    })
    return m.Migrate()
}

// 将 Sutui 上游的 Sora2 变体（sora_video2*）修正为按次计费（times），并将 output 与 input 对齐
func fixSutuiSoraTimesPricing() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "20251031_fix_sutui_sora_times",
        Migrate: func(tx *gorm.DB) error {
            if !tx.Migrator().HasTable(&Price{}) {
                return nil
            }
            // 需要修正的模型列表
            models := []string{
                "sora_video2",
                "sora_video2-portrait",
                "sora_video2-landscape",
                "sora_video2-portrait-15s",
                "sora_video2-landscape-15s",
                "sora_video2-portrait-hd",
                "sora_video2-landscape-hd",
                "sora_video2-portrait-hd-15s",
                "sora_video2-landscape-hd-15s",
                "sora_video2-portrait-hd-25s",
                "sora_video2-landscape-hd-25s",
            }

            // 仅对已存在且类型不为 times 的记录进行修正
            if err := tx.Model(&Price{}).
                Where("model IN (?)", models).
                Updates(map[string]any{"type": TimesPriceType, "output": gorm.Expr("CASE WHEN output=0 THEN input ELSE output END"), "input": 0} ).Error; err != nil {
                logger.SysLog("修正 Sutui Sora2 按次计费失败: " + err.Error())
                return err
            }
            return nil
        },
        Rollback: func(tx *gorm.DB) error { return nil },
    }
}

// 将所有 times 类型的价格统一为“输入免费（0）、输出收费（沿用原有输出；若输出为0则取原输入）”
func fixAllTimesPricingIO() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "20251031_fix_times_pricing_io",
        Migrate: func(tx *gorm.DB) error {
            if !tx.Migrator().HasTable(&Price{}) {
                return nil
            }
            if err := tx.Model(&Price{}).
                Where("type = ?", TimesPriceType).
                Updates(map[string]any{"output": gorm.Expr("CASE WHEN output=0 THEN input ELSE output END"), "input": 0}).Error; err != nil {
                logger.SysLog("修正 Times 类型输入/输出字段失败: " + err.Error())
                return err
            }
            return nil
        },
        Rollback: func(tx *gorm.DB) error { return nil },
    }
}

// 将 Sutui Veo3 系列（veo3、veo3.1、veo3-pro、veo3.1-pro、veo3.1-components）统一设置为 times 类型；
// 并按“输入0、输出收费（若输出为0则取输入）”修正历史数据。
func fixSutuiVeo3TimesPricing() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "20251031_fix_sutui_veo3_times",
        Migrate: func(tx *gorm.DB) error {
            if !tx.Migrator().HasTable(&Price{}) {
                return nil
            }
            models := []string{"veo3", "veo3.1", "veo3-pro", "veo3.1-pro", "veo3.1-components"}
            if err := tx.Model(&Price{}).
                Where("model IN (?)", models).
                Updates(map[string]any{"type": TimesPriceType, "output": gorm.Expr("CASE WHEN output=0 THEN input ELSE output END"), "input": 0}).Error; err != nil {
                logger.SysLog("修正 Sutui Veo3 按次计费失败: " + err.Error())
                return err
            }
            return nil
        },
        Rollback: func(tx *gorm.DB) error { return nil },
    }
}

// 将 Veo3 系列的品牌归属统一修正为 Google Gemini（owned_by_type = ChannelTypeGemini）
func fixVeo3OwnedByToGemini() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "20251031_fix_veo3_ownedby_gemini",
        Migrate: func(tx *gorm.DB) error {
            if !tx.Migrator().HasTable(&Price{}) {
                return nil
            }
            models := []string{"veo3", "veo3.1", "veo3-pro", "veo3.1-pro", "veo3.1-components"}
            if err := tx.Model(&Price{}).
                Where("model IN (?)", models).
                Update("owned_by_type", config.ChannelTypeGemini).Error; err != nil {
                logger.SysLog("修正 Veo3 品牌归属为 Gemini 失败: " + err.Error())
                return err
            }
            return nil
        },
        Rollback: func(tx *gorm.DB) error { return nil },
    }
}
