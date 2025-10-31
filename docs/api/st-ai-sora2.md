# sora2创建视频

## OpenAPI Specification

```yaml
openapi: 3.0.1
info:
  title: ''
  description: ''
  version: 1.0.0
paths:
  /v1/videos:
    post:
      summary: sora2创建视频
      deprecated: false
      description: ''
      tags: []
      parameters:
        - name: Authorization
          in: header
          description: ''
          required: true
          example: Bearer {{api_token}}
          schema:
            type: string
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
              properties:
                model:
                  type: string
                  enum:
                    - sora_video2
                    - sora_video2-portrait
                    - sora_video2-landscape
                    - sora_video2-portrait-15s
                    - sora_video2-landscape-15s
                    - sora_video2-portrait-hd
                    - sora_video2-landscape-hd
                    - sora_video2-portrait-hd-15s
                    - sora_video2-landscape-hd-15s
                    - sora_video2-portrait-hd-25s
                    - sora_video2-landscape-hd-25s
                  x-apifox-enum:
                    - value: sora_video2
                      name: ''
                      description: sora2视频模型，默认生成竖屏标清视频
                    - value: sora_video2-portrait
                      name: ''
                      description: sora2视频模型，竖屏标清视频
                    - value: sora_video2-landscape
                      name: ''
                      description: sora2视频模型，横屏标清视频
                    - value: sora_video2-portrait-15s
                      name: ''
                      description: sora2视频模型，竖屏标清视频，15s
                    - value: sora_video2-landscape-15s
                      name: ''
                      description: sora2视频模型，横屏标清视频，15s
                    - value: sora_video2-portrait-hd
                      name: ''
                      description: sora2视频模型，竖屏高清视频
                    - value: sora_video2-landscape-hd
                      name: ''
                      description: sora2视频模型，横屏高清视频
                    - value: sora_video2-portrait-hd-15s
                      name: ''
                      description: sora2视频模型，竖屏高清视频，15s
                    - value: sora_video2-landscape-hd-15s
                      name: ''
                      description: sora2视频模型，横屏高清视频，15s
                    - value: sora_video2-portrait-hd-25s
                      name: ''
                      description: sora2视频模型，竖屏高清视频，25s
                    - value: sora_video2-landscape-hd-25s
                      name: ''
                      description: sora2视频模型，横屏高清视频，25s
                  default: sora_video2
                  description: 模型类型
                  example: sora_video2
                prompt:
                  description: 提示词
                  example: 制作成动画
                  type: string
                input_reference:
                  type: string
                  format: binary
                  default: 本地图片
                  description: 本地图片，不能图片链接。
                  example: >-
                    cmMtdXBsb2FkLTE3NjEyMzIxODkwNTEtNQ==/0c8ae9162add47ffad0577f9e84c3478.png
              required:
                - model
                - prompt
            examples: {}
      responses:
        '200':
          description: ''
          content:
            application/json:
              schema:
                type: object
                properties:
                  id:
                    type: string
                  object:
                    type: string
                  model:
                    type: string
                  status:
                    type: string
                  progress:
                    type: 'null'
                  created_at:
                    type: integer
                  seconds:
                    type: integer
                  size:
                    type: string
                required:
                  - id
                  - object
                  - model
                  - status
                  - progress
                  - created_at
                  - seconds
                  - size
              example:
                id: video_c24cfb41-7dc1-4072-9a7e-a179d41e0e0d
                object: video
                model: sora_video2
                status: queued
                progress: null
                created_at: 1761227541
                seconds: 10
                size: 720x720
          headers: {}
          x-apifox-name: 成功
      security: []
      x-apifox-folder: ''
      x-apifox-status: released
      x-run-in-apifox: https://app.apifox.com/web/project/7291708/apis/api-365919918-run
components:
  schemas: {}
  securitySchemes:
    令牌:
      type: bearer
      scheme: bearer
servers: []
security: []

```

# sora2视频查询

## OpenAPI Specification

```yaml
openapi: 3.0.1
info:
  title: ''
  description: ''
  version: 1.0.0
paths:
  /v1/videos/video_id:
    get:
      summary: sora2视频查询
      deprecated: false
      description: 视频ID通过创建接口获取，对应里面的ID
      tags: []
      parameters:
        - name: Authorization
          in: header
          description: ''
          required: true
          example: Bearer {{api_token}}
          schema:
            type: string
      responses:
        '200':
          description: ''
          content:
            application/json:
              schema:
                type: object
                properties:
                  completed_at:
                    type: integer
                  created_at:
                    type: integer
                  error:
                    type: 'null'
                  expires_at:
                    type: 'null'
                  id:
                    type: string
                  model:
                    type: string
                  object:
                    type: string
                  progress:
                    type: integer
                  remixed_from_video_id:
                    type: 'null'
                  seconds:
                    type: integer
                  size:
                    type: string
                  status:
                    type: string
                  video_url:
                    type: string
                required:
                  - completed_at
                  - created_at
                  - error
                  - expires_at
                  - id
                  - model
                  - object
                  - progress
                  - remixed_from_video_id
                  - seconds
                  - size
                  - status
                  - video_url
              example:
                completed_at: 1761227671
                created_at: 1761227541
                error: null
                expires_at: null
                id: video_c24cfb41-7dc1-4072-9a7e-a179d41e0e0d
                model: sora_video2
                object: video
                progress: 100
                remixed_from_video_id: null
                seconds: 10
                size: 720x720
                status: completed
                video_url: >-
                  https://zeakai-oss.oss-cn-hongkong.aliyuncs.com/sora/f3c3a8f8-bf28-4973-a4a3-3e1397a91a11.mp4
          headers: {}
          x-apifox-name: 成功
      security: []
      x-apifox-folder: ''
      x-apifox-status: released
      x-run-in-apifox: https://app.apifox.com/web/project/7291708/apis/api-365920651-run
components:
  schemas: {}
  securitySchemes:
    令牌:
      type: bearer
      scheme: bearer
servers: []
security: []

```