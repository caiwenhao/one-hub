# MiniMax 视频渠道集成说明

## 模型输入与映射（文本 / 语音 / 视频）

本节说明 One Hub 中“输入的模型名”与“上游调用/计费键”的关系，便于前后端统一与排查。

- 文本（OpenAI 兼容 Chat）
  - 输入模型：`MiniMax-M1`（推荐）、`MiniMax-Text-01`
  - 对外接口：`POST /v1/chat/completions`
  - 价格（人民币/1k tokens，已内置，可后台覆盖）：
    - MiniMax-M1（默认采用 0–32k 档）：输入 0.0008，输出 0.008
    - MiniMax-Text-01：输入 0.001，输出 0.008

- 语音（TTS）
  - 输入模型：
    - HD 档：`speech-2.5-hd-preview`、`speech-02-hd`、`speech-01-hd`
    - Turbo 档：`speech-2.5-turbo-preview`、`speech-02-turbo`、`speech-01-turbo`
  - 对外接口：`POST /v1/audio/speech`（Provider 内部映射 minimaxi `/v1/t2a_v2`；如需异步请改走 `/v1/t2a_async_v2`）
  - 价格（人民币/1k 字符，已内置，可后台覆盖）：HD=0.35；Turbo=0.2

- 视频（Video Generation）
  - 动作与自动识别：
    - `text2video`：仅提供 `prompt`
    - `image2video`：提供 `first_frame_image`（可带 `prompt`）
    - `start-end2video`：同时提供 `first_frame_image` 与 `last_frame_image`
  - 官方上游（默认）
    - 输入模型：`MiniMax-Hailuo-02`
    - 提交：`POST /v1/video_generation`；查询：`GET /v1/query/video_generation`
    - 计费键模式：`minimax-<action>-minimax-hailuo-02-<resolution>-<duration>s`
      - 例：`minimax-text2video-minimax-hailuo-02-768p-10s`
    - 建议分辨率/时长：512P/768P/1080P × 6s/10s（1080P 支持 10s 档）
  - PPInfra 上游（推荐：前端“上游供应商”选择 PPInfra；或在渠道自定义参数顶层设置 `{"upstream":"ppinfra"}`；如需仅视频覆盖，使用 `{"video":{"upstream":"ppinfra"}}`）
    - 输入模型与用途：
      - `T2V-01` / `T2V-01-Director`（文生视频）
      - `I2V-01` / `I2V-01-live`（图生视频）
      - `S2V-01`（首尾帧）
    - 提交路径模板：`/v3/async/%s`（其中 `%s` 为模型段，如 `t2v-01`）
    - 查询：`/v3/async/task-result`
    - 计费键模式：`minimax-<action>-<model-seg>-<resolution>-<duration>s`
      - 例：`minimax-text2video-t2v-01-720p-6s`、`minimax-start-end2video-s2v-01-720p-6s`

> 说明：计费键由系统运行时根据“动作 + 输入模型 + 分辨率 + 时长”自动生成，无需前端传入。

## 路由与接口

- 主路径：`/minimaxi/v1/videos/:action`
  - 支持 `text2video`、`image2video`、`start-end2video` 动作，通过 `task.RelayTaskSubmit` 统一处理。
- 查询接口：
  - `GET /minimaxi/v1/query/video_generation?task_id=xxx`
  - `GET /minimaxi/v1/tasks`、`GET /minimaxi/v1/tasks/:task_id`
- 官方兼容别名：
  - `POST /v1/video_generation`
  - `GET /v1/query/video_generation`
  - 别名路径与官方完全一致，内部会自动推断动作类型。

## 渠道配置（Custom Parameter → video）


> 当前前端会根据“上游供应商”选项自动生成以下 `video` 配置，并复用渠道通用密钥。通常无需手动编辑 JSON，除非要覆盖更多细粒度参数。

```json
{
  "video": {
    "upstream": "official",
    "base_url": "https://api.minimaxi.com"
  }
}
```

### 顶层 upstream 快捷配置

如不需要分能力细化，建议使用顶层 `upstream` 直接切换：

```json
{"upstream": "ppinfra"}
```

- 模型字段（后台→渠道→模型）建议包含：
  - 文本/语音：`MiniMax-M1, MiniMax-Text-01, speech-02-turbo`（或你的常用项）
  - 视频（强烈推荐）：`minimax-*`（一条通配即可覆盖所有 minimaxi 视频计费键），以及 `MiniMax-Hailuo-02, T2V-01, I2V-01, S2V-01` 等常用名，方便筛选与统计。

- 也可以直接在顶层设置：`{"upstream":"ppinfra","api_key":"Bearer ppinfra_xxx"}`（默认会使用渠道密钥，一般无需单独写 `api_key`）。

- `upstream=ppinfra` 时：
  - 默认 `submit_path_template=/v3/async/%s`（模型名会自动转小写并拼到路径尾）。
  - 默认 `query_path=/v3/async/task-result`，请求会自动携带 `task_id` 查询参数。
  - 如上游返回结构包含 `videos`/`images`/`audios` 等字段，会写入任务数据并自动映射状态与失败原因。
  - `enable_prompt_expansion` 未显式设置时，统一走渠道默认；入参仍可覆盖。
  - 默认计费表已按 PPInfra 最新报价（0.6/1/2/4/6 元）写入 `model/price.go`，如果官方调价，可在后台“模型价格”中覆盖或直接修改该映射表。

## 计费模型键

- 统一格式：`minimax-<action>-<model>-<resolution>-<duration>s`
- 默认 price（单位：人民币，详见 `model/price.go`）：
  - text2video（MiniMax-Hailuo-02）：512P/768P/1080P × 6s；512P/768P × 10s
  - image2video：同上，单价略低
  - start-end2video：同上，单价略高
- 若需自定义，请在后台 **模型价格** 模块新增或覆盖相应键值。

## 任务存储

- `task.Properties`：记录 `model`、`action`、`duration`、`resolution`
- `task.Data`：保存最新一次查询得到的 `MiniMaxVideoQueryResponse`
- 状态映射：
  - Success → `SUCCESS`
  - Fail / Failed → `FAILURE`
  - 其余 Running/Processing/Pending → `IN_PROGRESS`

## 测试建议

1. 文生视频：`POST /minimaxi/v1/videos/text2video`
2. 图生视频：`POST /minimaxi/v1/videos/image2video`
3. 首尾帧：`POST /minimaxi/v1/videos/start-end2video`
4. 查询：`GET /minimaxi/v1/query/video_generation?task_id=xxx`（若上游为 PPInfra，会调用 `/v3/async/task-result` 并解析 `videos[]/task.status` 等字段）
5. 别名：`POST /v1/video_generation`（带首尾帧字段，自动识别动作）
6. 若接入 PPInfra，上述每项需再验证一次，确认路径模板与鉴权无误。也可直接在前端下拉选择 PPInfra，系统会自动写入 `{"upstream":"ppinfra"}`。

## 运维提示

- 监控：建议在 Prometheus 中新增 MiniMax 标签（成功率、耗时、失败类型）。
- 配额：后台价格调整后需运行一次「模型价格更新」，确保计费使用最新配置。
- 回调：如启用 `callback_url`，务必在回调服务中校验签名或 Token，避免伪造通知。
