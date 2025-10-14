# MiniMax 视频渠道集成说明

## 路由与接口

- 主路径：`/minimax/v1/videos/:action`
  - 支持 `text2video`、`image2video`、`start-end2video` 动作，通过 `task.RelayTaskSubmit` 统一处理。
- 查询接口：
  - `GET /minimax/v1/query/video_generation?task_id=xxx`
  - `GET /minimax/v1/tasks`、`GET /minimax/v1/tasks/:task_id`
- 官方兼容别名：
  - `POST /v1/video_generation`
  - `GET /v1/query/video_generation`
  - 别名路径与官方完全一致，内部会自动推断动作类型。

## 渠道配置（Custom Parameter → video）

```json
{
  "video": {
    "upstream": "official",        // official | ppinfra
    "api_key": "",                // 可选，缺省使用渠道主 key
    "base_url": "https://api.minimaxi.com",
    "submit_path_template": "/v3/async/%s",  // ppinfra 时可覆盖
    "query_path_template": "/v3/async/%s/%s",
    "enable_prompt_expansion": false,
    "callback_url": "",
    "extra_headers": {}
  }
}
```

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

1. 文生视频：`POST /minimax/v1/videos/text2video`
2. 图生视频：`POST /minimax/v1/videos/image2video`
3. 首尾帧：`POST /minimax/v1/videos/start-end2video`
4. 查询：`GET /minimax/v1/query/video_generation?task_id=xxx`（PPInfra 会调用 `/v3/async/task-result` 并解析 `videos[]/task.status` 等字段）
5. 别名：`POST /v1/video_generation`（带首尾帧字段，自动识别动作）
6. 若接入 PPInfra，上述每项需再验证一次，确认路径模板与鉴权无误。

## 运维提示

- 监控：建议在 Prometheus 中新增 MiniMax 标签（成功率、耗时、失败类型）。
- 配额：后台价格调整后需运行一次「模型价格更新」，确保计费使用最新配置。
- 回调：如启用 `callback_url`，务必在回调服务中校验签名或 Token，避免伪造通知。
