#!/bin/bash

# 可灵AI官方接口兼容性测试脚本
# 使用方法: ./test_kling_api.sh <your_domain> <your_token>

DOMAIN=${1:-"localhost:3000"}
TOKEN=${2:-"your_test_token"}

echo "=== 可灵AI官方接口兼容性测试 ==="
echo "测试域名: $DOMAIN"
echo "测试Token: $TOKEN"
echo ""

# 基础URL
BASE_URL="http://$DOMAIN/kling/v1"

# 测试1: 文生视频任务创建
echo "1. 测试文生视频任务创建..."
TEXT2VIDEO_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/text2video" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v1",
    "prompt": "一只可爱的小猫在花园里玩耍",
    "mode": "std",
    "duration": "5",
    "aspect_ratio": "16:9",
    "negative_prompt": "低质量，模糊",
    "cfg_scale": 0.5,
    "external_task_id": "test_text2video_001"
  }')

echo "响应: $TEXT2VIDEO_RESPONSE"
echo ""

# 提取任务ID
TASK_ID=$(echo $TEXT2VIDEO_RESPONSE | grep -o '"task_id":"[^"]*"' | cut -d'"' -f4)
echo "提取到的任务ID: $TASK_ID"
echo ""

# 测试2: 图生视频任务创建
echo "2. 测试图生视频任务创建..."
IMAGE2VIDEO_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/image2video" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v1",
    "image": "https://example.com/test.jpg",
    "prompt": "让图片中的内容动起来",
    "mode": "std",
    "duration": "5",
    "aspect_ratio": "16:9",
    "external_task_id": "test_image2video_001"
  }')

echo "响应: $IMAGE2VIDEO_RESPONSE"
echo ""

# 提取图生视频任务ID
IMAGE_TASK_ID=$(echo $IMAGE2VIDEO_RESPONSE | grep -o '"task_id":"[^"]*"' | cut -d'"' -f4)
echo "提取到的图生视频任务ID: $IMAGE_TASK_ID"
echo ""

# 测试3: 查询单个任务 (通过task_id)
if [ ! -z "$TASK_ID" ]; then
  echo "3. 测试查询单个任务 (通过task_id)..."
  TASK_QUERY_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/text2video/$TASK_ID" \
    -H "Authorization: Bearer $TOKEN")
  
  echo "响应: $TASK_QUERY_RESPONSE"
  echo ""
fi

# 测试4: 查询单个任务 (通过external_task_id)
echo "4. 测试查询单个任务 (通过external_task_id)..."
EXTERNAL_TASK_QUERY_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/text2video/test_text2video_001" \
  -H "Authorization: Bearer $TOKEN")

echo "响应: $EXTERNAL_TASK_QUERY_RESPONSE"
echo ""

# 测试5: 查询任务列表 (默认分页)
echo "5. 测试查询任务列表 (默认分页)..."
TASK_LIST_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/text2video" \
  -H "Authorization: Bearer $TOKEN")

echo "响应: $TASK_LIST_RESPONSE"
echo ""

# 测试6: 查询任务列表 (自定义分页)
echo "6. 测试查询任务列表 (自定义分页)..."
TASK_LIST_CUSTOM_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/text2video?pageNum=1&pageSize=10" \
  -H "Authorization: Bearer $TOKEN")

echo "响应: $TASK_LIST_CUSTOM_RESPONSE"
echo ""

# 测试7: 摄像机控制参数测试
echo "7. 测试摄像机控制参数..."
CAMERA_CONTROL_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/text2video" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v1",
    "prompt": "一个旋转的地球",
    "mode": "std",
    "duration": "5",
    "camera_control": {
      "type": "simple",
      "config": {
        "horizontal": 0,
        "vertical": 0,
        "pan": 5,
        "tilt": 0,
        "roll": 0,
        "zoom": 0
      }
    },
    "external_task_id": "test_camera_control_001"
  }')

echo "响应: $CAMERA_CONTROL_RESPONSE"
echo ""

# 测试8: 图生视频动态笔刷测试
echo "8. 测试图生视频动态笔刷..."
DYNAMIC_MASK_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/image2video" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v1",
    "image": "https://example.com/test.jpg",
    "prompt": "宇航员站起身走了",
    "mode": "pro",
    "duration": "5",
    "cfg_scale": 0.5,
    "dynamic_masks": [
      {
        "mask": "https://example.com/mask.png",
        "trajectories": [
          {"x": 279, "y": 219},
          {"x": 417, "y": 65}
        ]
      }
    ],
    "external_task_id": "test_dynamic_mask_001"
  }')

echo "响应: $DYNAMIC_MASK_RESPONSE"
echo ""

# 测试9: 图生视频查询测试
if [ ! -z "$IMAGE_TASK_ID" ]; then
  echo "9. 测试图生视频查询 (通过task_id)..."
  IMAGE_TASK_QUERY_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/image2video/$IMAGE_TASK_ID" \
    -H "Authorization: Bearer $TOKEN")
  
  echo "响应: $IMAGE_TASK_QUERY_RESPONSE"
  echo ""
fi

# 测试10: 图生视频任务列表查询
echo "10. 测试图生视频任务列表查询..."
IMAGE2VIDEO_LIST_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/image2video?pageNum=1&pageSize=5" \
  -H "Authorization: Bearer $TOKEN")

echo "响应: $IMAGE2VIDEO_LIST_RESPONSE"
echo ""

# 测试11: 错误处理测试
echo "11. 测试错误处理 (缺少必填参数)..."
ERROR_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/text2video" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{}')

echo "响应: $ERROR_RESPONSE"
echo ""

# 测试12: 多图参考生视频测试
echo "12. 测试多图参考生视频..."
MULTI_IMAGE_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/multi-image2video" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v1-6",
    "image_list": [
      {"image": "https://example.com/image1.jpg"},
      {"image": "https://example.com/image2.jpg"},
      {"image": "https://example.com/image3.jpg"}
    ],
    "prompt": "三张图片的内容串联成一个连贯的故事",
    "mode": "std",
    "duration": "5",
    "aspect_ratio": "16:9",
    "external_task_id": "test_multi_image_001"
  }')

echo "响应: $MULTI_IMAGE_RESPONSE"
echo ""

# 提取多图生视频任务ID
MULTI_IMAGE_TASK_ID=$(echo $MULTI_IMAGE_RESPONSE | grep -o '"task_id":"[^"]*"' | cut -d'"' -f4)
echo "提取到的多图生视频任务ID: $MULTI_IMAGE_TASK_ID"
echo ""

# 测试13: 多图参考生视频查询测试
if [ ! -z "$MULTI_IMAGE_TASK_ID" ]; then
  echo "13. 测试多图参考生视频查询..."
  MULTI_IMAGE_QUERY_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/multi-image2video/$MULTI_IMAGE_TASK_ID" \
    -H "Authorization: Bearer $TOKEN")
  
  echo "响应: $MULTI_IMAGE_QUERY_RESPONSE"
  echo ""
fi

# 测试14: 多图参考生视频任务列表查询
echo "14. 测试多图参考生视频任务列表查询..."
MULTI_IMAGE_LIST_RESPONSE=$(curl -s -X GET "$BASE_URL/videos/multi-image2video?pageNum=1&pageSize=5" \
  -H "Authorization: Bearer $TOKEN")

echo "响应: $MULTI_IMAGE_LIST_RESPONSE"
echo ""

# 测试15: 多模态视频编辑 - 初始化待编辑视频
echo "15. 测试多模态视频编辑 - 初始化待编辑视频..."
INIT_SELECTION_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/multi-elements/init-selection" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "video_url": "https://example.com/test_video.mp4"
  }')

echo "响应: $INIT_SELECTION_RESPONSE"
echo ""

# 提取session_id
SESSION_ID=$(echo $INIT_SELECTION_RESPONSE | grep -o '"session_id":"[^"]*"' | cut -d'"' -f4)
echo "提取到的session_id: $SESSION_ID"
echo ""

# 测试16: 多模态视频编辑 - 增加视频选区
if [ ! -z "$SESSION_ID" ]; then
  echo "16. 测试多模态视频编辑 - 增加视频选区..."
  ADD_SELECTION_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/multi-elements/add-selection" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"session_id\": \"$SESSION_ID\",
      \"frame_index\": 10,
      \"points\": [
        {\"x\": 0.3, \"y\": 0.4},
        {\"x\": 0.7, \"y\": 0.6}
      ]
    }")
  
  echo "响应: $ADD_SELECTION_RESPONSE"
  echo ""
fi

# 测试17: 多模态视频编辑 - 预览已选区视频
if [ ! -z "$SESSION_ID" ]; then
  echo "17. 测试多模态视频编辑 - 预览已选区视频..."
  PREVIEW_SELECTION_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/multi-elements/preview-selection" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"session_id\": \"$SESSION_ID\"
    }")
  
  echo "响应: $PREVIEW_SELECTION_RESPONSE"
  echo ""
fi

# 测试18: 多模态视频编辑 - 创建编辑任务
if [ ! -z "$SESSION_ID" ]; then
  echo "18. 测试多模态视频编辑 - 创建编辑任务..."
  MULTI_ELEMENTS_RESPONSE=$(curl -s -X POST "$BASE_URL/videos/multi-elements" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"model_name\": \"kling-v1-6\",
      \"session_id\": \"$SESSION_ID\",
      \"edit_mode\": \"addition\",
      \"image_list\": [
        {\"image\": \"https://example.com/add_image.jpg\"}
      ],
      \"prompt\": \"基于<<<video_1>>>中的原始内容，以自然生动的方式，将<<<image_1>>>中的元素，融入<<<video_1>>>的场景中\",
      \"mode\": \"std\",
      \"duration\": \"5\",
      \"external_task_id\": \"test_multi_elements_001\"
    }")
  
  echo "响应: $MULTI_ELEMENTS_RESPONSE"
  echo ""
fi

# 测试19: 图像生成任务创建
echo "19. 测试图像生成任务创建..."
IMAGE_GEN_RESPONSE=$(curl -s -X POST "$BASE_URL/images/generations" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v1",
    "prompt": "一只可爱的小猫在花园里玩耶",
    "negative_prompt": "低质量，模糊",
    "resolution": "1k",
    "n": 2,
    "aspect_ratio": "16:9",
    "external_task_id": "test_image_gen_001"
  }')

echo "响应: $IMAGE_GEN_RESPONSE"
echo ""

# 提取图像生成任务ID
IMAGE_GEN_TASK_ID=$(echo $IMAGE_GEN_RESPONSE | grep -o '"task_id":"[^"]*"' | cut -d'"' -f4)
echo "提取到的图像生成任务ID: $IMAGE_GEN_TASK_ID"
echo ""

# 测试20: 图生图任务创建 
echo "20. 测试图生图任务创建..."
IMAGE2IMAGE_RESPONSE=$(curl -s -X POST "$BASE_URL/images/generations" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v1-5",
    "prompt": "美丽的晨曦风景",
    "image": "https://example.com/reference.jpg",
    "image_reference": "subject",
    "image_fidelity": 0.7,
    "human_fidelity": 0.5,
    "resolution": "2k",
    "n": 1,
    "aspect_ratio": "1:1",
    "external_task_id": "test_image2image_001"
  }')

echo "响应: $IMAGE2IMAGE_RESPONSE"
echo ""

# 测试21: 多图参考生图任务创建
echo "21. 测试多图参考生图任务创建..."
MULTI_IMAGE2IMAGE_RESPONSE=$(curl -s -X POST "$BASE_URL/images/multi-image2image" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "model_name": "kling-v2",
    "prompt": "综合多张图片的特点生成新图像",
    "subject_image_list": [
      {"subject_image": "https://example.com/subject1.jpg"},
      {"subject_image": "https://example.com/subject2.jpg"},
      {"subject_image": "https://example.com/subject3.jpg"}
    ],
    "scene_image": "https://example.com/scene.jpg",
    "style_image": "https://example.com/style.jpg",
    "n": 3,
    "aspect_ratio": "4:3",
    "external_task_id": "test_multi_image2image_001"
  }')

echo "响应: $MULTI_IMAGE2IMAGE_RESPONSE"
echo ""

# 测试22: 图像任务查询
if [ ! -z "$IMAGE_GEN_TASK_ID" ]; then
  echo "22. 测试图像任务查询 (通过task_id)..."
  IMAGE_TASK_QUERY_RESPONSE=$(curl -s -X GET "$BASE_URL/images/generations/$IMAGE_GEN_TASK_ID" \
    -H "Authorization: Bearer $TOKEN")
  
  echo "响应: $IMAGE_TASK_QUERY_RESPONSE"
  echo ""
fi

# 测试23: 图像任务列表查询
echo "23. 测试图像任务列表查询..."
IMAGE_LIST_RESPONSE=$(curl -s -X GET "$BASE_URL/images/generations?pageNum=1&pageSize=5" \
  -H "Authorization: Bearer $TOKEN")

echo "响应: $IMAGE_LIST_RESPONSE"
echo ""

# 测试24: 多图参考生图任务列表查询
echo "24. 测试多图参考生图任务列表查询..."
MULTI_IMAGE2IMAGE_LIST_RESPONSE=$(curl -s -X GET "$BASE_URL/images/multi-image2image?pageNum=1&pageSize=5" \
  -H "Authorization: Bearer $TOKEN")

echo "响应: $MULTI_IMAGE2IMAGE_LIST_RESPONSE"
echo ""

echo "=== 测试完成 ==="
echo ""
echo "对比官方API格式:"
echo "1. 检查响应字段是否包含: code, message, request_id, data"
echo "2. 检查任务数据是否包含: task_id, task_status, created_at, updated_at, task_info"
echo "3. 检查错误响应格式是否符合官方规范"
echo ""
echo "注意事项:"
echo "- 确保已配置可灵AI渠道和模型"
echo "- 确保Token具有相应权限"
echo "- 实际视频生成需要时间，可通过查询接口检查进度"
echo "- 支持文生视频和图生视频两种模式"
echo "- 支持动态笔刷、静态笔刷和摄像机控制功能"