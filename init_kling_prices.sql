-- 可灵AI价格初始化脚本
-- 根据官方资源包定价规则配置One Hub中的可灵AI模型价格
-- 所有价格采用"times"计费类型，按次收费

-- 删除现有的可灵AI价格配置（如果存在）
DELETE FROM prices WHERE channel_type = 53;

-- 视频生成模型价格配置
-- V1 系列
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-video_kling-v1_std_5', 'times', 53, 5.0, 5.0, false),
('kling-video_kling-v1_std_10', 'times', 53, 10.0, 10.0, false),
('kling-video_kling-v1_pro_5', 'times', 53, 15.0, 15.0, false),
('kling-video_kling-v1_pro_10', 'times', 53, 30.0, 30.0, false);

-- V1.5 系列
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-video_kling-v1.5_std_5', 'times', 53, 5.0, 5.0, false),
('kling-video_kling-v1.5_std_10', 'times', 53, 10.0, 10.0, false),
('kling-video_kling-v1.5_pro_5', 'times', 53, 15.0, 15.0, false),
('kling-video_kling-v1.5_pro_10', 'times', 53, 30.0, 30.0, false);

-- V1.6 系列
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-video_kling-v1-6_std_5', 'times', 53, 10.0, 10.0, false),
('kling-video_kling-v1-6_std_10', 'times', 53, 20.0, 20.0, false),
('kling-video_kling-v1-6_pro_5', 'times', 53, 30.0, 30.0, false),
('kling-video_kling-v1-6_pro_10', 'times', 53, 60.0, 60.0, false);

-- V2-Master 系列
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-video_kling-v2-master_std_5', 'times', 53, 15.0, 15.0, false),
('kling-video_kling-v2-master_std_10', 'times', 53, 30.0, 30.0, false),
('kling-video_kling-v2-master_pro_5', 'times', 53, 45.0, 45.0, false),
('kling-video_kling-v2-master_pro_10', 'times', 53, 90.0, 90.0, false);

-- V2.1-Master 系列
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-video_kling-v2-1-master_std_5', 'times', 53, 15.0, 15.0, false),
('kling-video_kling-v2-1-master_std_10', 'times', 53, 30.0, 30.0, false),
('kling-video_kling-v2-1-master_pro_5', 'times', 53, 45.0, 45.0, false),
('kling-video_kling-v2-1-master_pro_10', 'times', 53, 90.0, 90.0, false);

-- 图像生成模型价格配置
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-image_kling-v1_std', 'times', 53, 5.0, 5.0, false),
('kling-image_kling-v1_pro', 'times', 53, 15.0, 15.0, false),
('kling-image_kling-v1.5_std', 'times', 53, 5.0, 5.0, false),
('kling-image_kling-v1.5_pro', 'times', 53, 15.0, 15.0, false),
('kling-image_kling-v1-6_std', 'times', 53, 10.0, 10.0, false),
('kling-image_kling-v1-6_pro', 'times', 53, 30.0, 30.0, false),
('kling-image_kling-v2-master_std', 'times', 53, 15.0, 15.0, false),
('kling-image_kling-v2-master_pro', 'times', 53, 45.0, 45.0, false),
('kling-image_kling-v2-1-master_std', 'times', 53, 15.0, 15.0, false),
('kling-image_kling-v2-1-master_pro', 'times', 53, 45.0, 45.0, false);

-- 虚拟试穿模型价格配置
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-try-on_kling-v1_std', 'times', 53, 5.0, 5.0, false),
('kling-try-on_kling-v1_pro', 'times', 53, 15.0, 15.0, false),
('kling-try-on_kling-v1.5_std', 'times', 53, 5.0, 5.0, false),
('kling-try-on_kling-v1.5_pro', 'times', 53, 15.0, 15.0, false),
('kling-try-on_kling-v1-6_std', 'times', 53, 10.0, 10.0, false),
('kling-try-on_kling-v1-6_pro', 'times', 53, 30.0, 30.0, false);

-- 兼容性配置（为了与现有实现兼容）
INSERT INTO prices (model, type, channel_type, input, output, locked) VALUES 
('kling-v1', 'times', 53, 5.0, 5.0, false),
('kling-v1.5', 'times', 53, 5.0, 5.0, false),
('kling-v1-6', 'times', 53, 10.0, 10.0, false),
('kling-v2-master', 'times', 53, 15.0, 15.0, false),
('kling-v2-1-master', 'times', 53, 15.0, 15.0, false);

-- 查询验证插入的价格配置
SELECT 
    model, 
    type, 
    channel_type, 
    input AS price_yuan, 
    CASE 
        WHEN model LIKE '%_std_%' THEN '标准模式'
        WHEN model LIKE '%_pro_%' THEN '专家模式'
        ELSE '兼容模式'
    END AS mode,
    CASE 
        WHEN model LIKE '%_5' THEN '5秒'
        WHEN model LIKE '%_10' THEN '10秒'
        WHEN model LIKE 'kling-video%' THEN '视频生成'
        WHEN model LIKE 'kling-image%' THEN '图像生成'
        WHEN model LIKE 'kling-try-on%' THEN '虚拟试穿'
        ELSE '通用'
    END AS duration_or_type
FROM prices 
WHERE channel_type = 53 
ORDER BY 
    CASE 
        WHEN model LIKE 'kling-video%' THEN 1
        WHEN model LIKE 'kling-image%' THEN 2
        WHEN model LIKE 'kling-try-on%' THEN 3
        ELSE 4
    END,
    model;