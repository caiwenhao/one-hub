图生视频
POST https://api.vidu.cn/ent/v2/img2video
请求头
字段	值	描述
Content-Type	application/json	数据交换格式
Authorization	Token {your api key}	将 {your api key} 替换为您的 token
请求体
参数名称	类型	必填	参数描述
model	String	是	模型名称
可选值：viduq2-pro、viduq2-turbo、viduq1 、viduq1-classic 、vidu2.0、vidu1.5
- viduq2-pro：新模型，效果好，细节丰富
- viduq2-turbo：新模型，效果好，生成快
- viduq1：画面清晰，平滑转场，运镜稳定
- viduq1-classic：画面清晰，转场、运镜更丰富
- vidu2.0：生成速度快
- vidu1.5：动态幅度大
images	Array[String]	是	首帧图像
模型将以此参数中传入的图片为首帧画面来生成视频。
注1：支持传入图片 Base64 编码或图片URL（确保可访问）；
注2：只支持输入 1 张图；
注3：图片支持 png、jpeg、jpg、webp格式；
注4：图片比例需要小于 1:4 或者 4:1 ；
注5：图片大小不超过 50 MB；
注6：请注意，base64 decode之后的字节长度需要小于10M，且编码必须包含适当的内容类型字符串，例如：
data:image/png;base64,{base64_encode}
prompt	String	可选	文本提示词
生成视频的文本描述。
注：字符长度不能超过 2000 个字符
duration	Int	可选	视频时长
viduq2-pro 默认为 5，可选：2、3、4、5、6、7、8
viduq2-turbo 默认为 5，可选：2、3、4、5、6、7、8
viduq1 默认为 5，可选：5
viduq1-classic 默认为 5，可选：5
vidu2.0 默认为 4，可选：4、8
vidu1.5 默认为 4，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
- viduq2-pro 1-8秒：默认 720p，可选：720p、1080p
- viduq2-turbo 1-8秒：默认 720p，可选：720p、1080p
- viduq1 5秒：默认 1080p，可选：1080p
- viduq1-classic 5秒：默认 1080p，可选：1080p
- vidu2.0 4秒：默认 360p，可选：360p、720p、1080p
- vidu2.0 8秒：默认 720p，可选：720p
- vidu1.5 4秒：默认 360p，可选：360p、720p、1080p
- vidu1.5 8秒：默认 720p，可选：720p
movement_amplitude	String	可选	运动幅度
默认 auto，可选值：auto、small、medium、large
bgm	Bool	可选	是否为生成的视频添加背景音乐。
默认：false，可选值 true 、false
- 传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。
- BGM不限制时长，系统根据视频时长自动适配
payload	String	可选	透传参数
不做任何处理，仅数据传输
注：最多 1048576个字符
off_peak	Bool	可选	错峰模式，默认为：false，可选值：
- true：错峰生成视频；
- false：即时生成视频；
注1：错峰模式消耗的积分更低，具体请查看产品定价
注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分；
注3：您也可以手动取消错峰任务
watermark	Bool	可选	是否添加水印
- true：添加水印；
- false：不添加水印；
注1：目前水印内容为固定，内容由AI生成，默认不加
注2：您可以通过watermarked_url参数查询获取带水印的视频内容，详情见查询任务接口
callback_url	String	可选	Callback 协议
需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务有状态变化时，Vidu 将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致
回调返回的"status"包括以下状态：
- processing 任务处理中
- success 任务完成（如发送失败，回调三次）
- failed 任务失败（如发送失败，回调三次）
Vidu采用回调签名算法进行认证，详情见：回调签名算法
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "model": "viduq2-pro",
    "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/image2video.png"],
    "prompt": "The astronaut waved and the camera moved up.",
    "duration": 5,
    "seed": 0,
    "resolution": "1080p",
    "movement_amplitude": "auto",
    "off_peak": false
}' https://api.vidu.cn/ent/v2/img2video
响应体
字段	类型	描述
task_id	String	Vidu 生成的任务ID
state	String	处理状态
可选值：
created 创建成功
queueing 任务排队中
processing 任务处理中
success 任务成功
failed 任务失败
model	String	本次调用的模型名称
prompt	String	本次调用的提示词参数
images	Array[String]	本次调用的图像参数
duration	Int	本次调用的视频时长参数
seed	Int	本次调用的随机种子参数
resolution	String	本次调用的分辨率参数
bgm	Bool	本次调用的背景音乐参数
movement_amplitude	String	本次调用的镜头动态幅度参数
payload	String	本次调用时传入的透传参数
off_peak	Bool	本次调用时是否使用错峰模式
credits	String	本次调用使用的积分数
watermark	Bool	本次提交任务是否使用水印
created_at	String	任务创建时间
{
  "task_id": "your_task_id_here",
  "state": "created",
  "model": "viduq2-pro",
  "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/image2video.png"],
  "prompt": "The astronaut waved and the camera moved up.",
  "duration": 5,
  "seed": random_number,
  "resolution": "1080p",
  "movement_amplitude": "auto",
  "payload":"",
  "off_peak": false,
  "credits":credits_number,
  "created_at": "2025-01-01T15:41:31.968916Z"
}


参考生视频
POST https://api.vidu.cn/ent/v2/reference2video
请求头
字段	值	描述
Content-Type	application/json	数据交换格式
Authorization	Token {your api key}	将 {your api key} 替换为您的 token
请求体
参数名称	类型	必填	参数描述
model	String	是	模型名称
可选值：viduq1、vidu2.0、vidu1.5
- viduq1：画面清晰，平滑转场，运镜稳定
- vidu2.0：生成速度快
- vidu1.5：动态幅度大
images	Array[String]	是	图像参考
支持上传1～7张图片，模型将以此参数中传入的图片中的主题为参考生成具备主体一致的视频。
注1：支持传入图片 Base64 编码或图片URL（确保可访问）；
注2：图片支持 png、jpeg、jpg、webp格式
注3：图片像素不能小于 128*128，且比例需要小于1:4或者4:1
注4：且大小不超过50M。
注5：请注意，base64 decode之后的字节长度需要小于10M，且编码必须包含适当的内容类型字符串，例如：
data:image/png;base64,{base64_encode}
prompt	String	是	文本提示词
生成视频的文本描述。
注：字符长度不能超过 2000 个字符
duration	Int	可选	视频时长参数，默认值依据模型而定：
viduq1：默认5秒，可选：5
vidu2.0：默认4秒，可选：4
vidu1.5: 默认4秒，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
aspect_ratio	String	可选	比例
默认 16:9，可选值如下：16:9、9:16、1:1
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
viduq1 （5秒）：默认 1080p, 可选：1080p
vidu2.0 （4秒）：默认 360p, 可选：360p、720p
vidu1.5（4秒）：默认 360p，可选：360p、720p、1080p
vidu1.5（8秒）：默认 720p，可选：720p
movement_amplitude	String	可选	运动幅度
默认 auto，可选值：auto、small、medium、large
bgm	Bool	可选	是否为生成的视频添加背景音乐。
默认：false，可选值 true 、false
- 传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。
- BGM不限制时长，系统根据视频时长自动适配
payload	String	可选	透传参数
不做任何处理，仅数据传输
注：最多 1048576个字符
off_peak	Bool	可选	错峰模式，默认为：false，可选值：
- true：错峰生成视频；
- false：即时生成视频；
注1：错峰模式消耗的积分更低，具体请查看产品定价
注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分；
注3：您也可以手动取消错峰任务
watermark	Bool	可选	是否添加水印
- true：添加水印；
- false：不添加水印；
注1：目前水印内容为固定，内容由AI生成，默认不加
注2：您可以通过watermarked_url参数查询获取带水印的视频内容，详情见查询任务接口
callback_url	String	可选	Callback 协议
需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务有状态变化时，Vidu 将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致
回调返回的"status"包括以下状态：
- processing 任务处理中
- success 任务完成（如发送失败，回调三次）
- failed 任务失败（如发送失败，回调三次）
Vidu采用回调签名算法进行认证，详情见：回调签名算法
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "model": "vidu2.0",
    "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-1.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-2.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-3.png"],
    "prompt": "Santa Claus and the bear hug by the lakeside.",
    "duration": 4,
    "seed": 0,
    "aspect_ratio": "16:9",
    "resolution": "720p",
    "movement_amplitude": "auto",
    "off_peak": false
}' https://api.vidu.cn/ent/v2/reference2video
响应体
字段	类型	描述
task_id	String	Vidu 生成的任务ID
state	String	处理状态
可选值：
created 创建成功
queueing 任务排队中
processing 任务处理中
success 任务成功
failed 任务失败
model	String	本次调用的模型名称
prompt	String	本次调用的提示词参数
images	Array[String]	本次调用的图像参数
duration	Int	本次调用的视频时长参数
seed	Int	本次调用的随机种子参数
aspect_ratio	String	本次调用的 比例 参数
resolution	String	本次调用的分辨率参数
bgm	Bool	本次调用的背景音乐参数
movement_amplitude	String	本次调用的镜头动态幅度参数
payload	String	本次调用时传入的透传参数
off_peak	Bool	本次调用时是否使用错峰模式
credits	String	本次调用使用的积分数
watermark	Bool	本次提交任务是否使用水印
created_at	String	任务创建时间
{
  "task_id": "your_task_id_here",
  "state": "created",
  "model": "vidu2.0",
  "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-1.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-2.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-3.png"],
  "prompt": "Santa Claus and the bear hug by the lakeside.",
  "duration": 4,
  "seed": random_number,
  "aspect_ratio": "16:9",
  "resolution": "720p",
  "movement_amplitude": "auto",
  "payload":"",
  "off_peak": false,
  "credits":credits_number,
  "created_at": "2025-01-01T15:41:31.968916Z"
}


首尾帧
POST https://api.vidu.cn/ent/v2/start-end2video
请求头
字段	值	描述
Content-Type	application/json	数据交换格式
Authorization	Token {your api key}	将 {your api key} 替换为您的 token
请求体
参数名称	类型	必填	参数描述
model	String	是	模型名称
可选值：viduq2-pro、viduq2-turbo、viduq1 、viduq1-classic、vidu2.0、vidu1.5
- viduq2-pro：新模型，效果好，细节丰富
- viduq2-turbo：新模型，效果好，生成快
- viduq1：画面清晰，平滑转场，运镜稳定
- viduq1-classic：画面清晰，转场、运镜更丰富
- vidu2.0：生成速度快
- vidu1.5：动态幅度大
images	Array[String]	是	图像
支持输入两张图，上传的第一张图片视作首帧图，第二张图片视作尾帧图，模型将以此参数中传入的图片来生成视频

注1: 首尾帧两张输入图的分辨率需相近，首帧图的分辨率/尾帧图的分辨率要在0.8～1.25之间。且图片比例需要小于1:4或者4:1；
注2: 支持传入图片 Base64 编码或图片URL（确保可访问）；
注3: 图片支持 png、jpeg、jpg、webp格式；
注4: 图片大小不超过50M；
注5: 请注意，base64 decode之后的字节长度需要小于10M，且编码必须包含适当的内容类型字符串，例如：
data:image/png;base64,{base64_encode}
prompt	String	可选	文本提示词
生成视频的文本描述。
注：字符长度不能超过 2000 个字符
duration	Int	可选	视频时长参数，默认值依据模型而定：
- viduq2-pro 默认为 5，可选：2、3、4、5、6、7、8
- viduq2-turbo 默认为 5，可选：2、3、4、5、6、7、8
- viduq1 和 viduq1-classic 默认 5 秒，可选：5
- vidu2.0 默认 4 秒，可选：4、8
- vidu1.5 默认 4 秒，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
- viduq2-pro 1-8秒：默认 720p，可选：720p、1080p
- viduq2-turbo 1-8秒：默认 720p，可选：720p、1080p
- viduq1 和 viduq1-classic 5秒：默认 1080p，可选：1080p
- vidu2.0 4秒：默认 360p，可选：360p、720p、1080p
- vidu2.0 8秒：默认 720p，可选：720p
- vidu1.5 4秒：默认 360p，可选：360p、720p、1080p
- vidu1.5 8秒：默认 720p，可选：720p
movement_amplitude	String	可选	运动幅度
默认 auto，可选值：auto、small、medium、large
bgm	bool	可选	是否为生成的视频添加背景音乐。
默认：false，可选值 true 、false
- 传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。
- BGM不限制时长，系统根据视频时长自动适配
payload	String	可选	透传参数
不做任何处理，仅数据传输
注：最多 1048576个字符
off_peak	Bool	可选	错峰模式，默认为：false，可选值：
- true：错峰生成视频；
- false：即时生成视频；
注1：错峰模式消耗的积分更低，具体请查看产品定价
注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分；
注3：您也可以手动取消错峰任务
watermark	Bool	可选	是否添加水印
- true：添加水印；
- false：不添加水印；
注1：目前水印内容为固定，内容由AI生成，默认不加
注2：您可以通过watermarked_url参数查询获取带水印的视频内容，详情见查询任务接口
callback_url	String	可选	Callback 协议
需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务有状态变化时，Vidu 将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致
回调返回的"status"包括以下状态：
- processing 任务处理中
- success 任务完成（如发送失败，回调三次）
- failed 任务失败（如发送失败，回调三次）
Vidu采用回调签名算法进行认证，详情见：回调签名算法
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "model": "viduq2-pro",
    "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/startend2video-1.jpeg","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/startend2video-2.jpeg"],
    "prompt": "The camera zooms in on the bird, which then flies to the right. With its flight being smooth and natural, the bird soars in the sky. with a red light effect following and surrounding it from behind.",
    "duration": 5,
    "seed": 0,
    "resolution": "1080p",
    "movement_amplitude": "auto",
    "off_peak": false
}' https://api.vidu.cn/ent/v2/start-end2video
响应体
字段	类型	描述
task_id	String	Vidu 生成的任务ID
state	String	处理状态
可选值：
created 创建成功
queueing 任务排队中
processing 任务处理中
success 任务成功
failed 任务失败
model	String	本次调用的模型名称
prompt	String	本次调用的提示词参数
images	Array[String]	本次调用的图像参数
duration	Int	本次调用的视频时长参数
seed	Int	本次调用的随机种子参数
resolution	String	本次调用的分辨率参数
bgm	bool	本次调用的背景音乐参数
movement_amplitude	String	本次调用的镜头动态幅度参数
payload	String	本次调用时传入的透传参数
off_peak	Bool	本次调用时是否使用错峰模式
credits	String	本次调用使用的积分数
watermark	Bool	本次提交任务是否使用水印
created_at	String	任务创建时间
{
  "task_id": "your_task_id_here",
  "state": "created",
  "model": "viduq2-turbo",
  "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/startend2video-1.jpeg","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/startend2video-2.jpeg"],
  "prompt": "The camera zooms in on the bird, which then flies to the right. The bird's flight is smooth and natural, with a red light effect following and surrounding it from behind.",
  "duration": 5,
  "seed": random_number,
  "resolution": "1080p",
  "movement_amplitude": "auto",
  "payload":"",
  "off_peak": false,
  "credits":credits_number,
  "created_at": "2025-01-01T15:41:31.968916Z"
}


文生视频
POST https://api.vidu.cn/ent/v2/text2video
请求头
字段	值	描述
Content-Type	application/json	数据交换格式
Authorization	Token {your api key}	将 {your api key} 替换为您的 token
请求体
参数名称	类型	必填	参数描述
model	String	是	模型名称
可选值：viduq1 、vidu1.5
- viduq1：画面清晰，平滑转场，运镜稳定
- vidu1.5：动态幅度大
style	String	可选	风格
默认 general，可选值：general、anime
general：通用风格，可以通过提示词来控制风格
anime：动漫风格，仅在动漫风格表现突出，可以通过不同的动漫风格提示词来控制
prompt	String	是	文本提示词
生成视频的文本描述。
注：字符长度不能超过 2000 个字符
duration	Int	可选	视频时长参数，默认值依据模型而定：
- viduq1 : 默认5秒，可选：5
- vidu1.5 : 默认4秒，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
aspect_ratio	String	可选	比例
默认 16:9，可选值：16:9、9:16、1:1
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
- viduq1 5秒：默认 1080p，可选：1080p
- vidu1.5 4秒：默认 360p，可选：360p、720p、1080p
- vidu1.5 8秒：默认 720p，可选：720p
movement_amplitude	String	可选	运动幅度
默认 auto，可选值：auto、small、medium、large
bgm	Bool	可选	是否为生成的视频添加背景音乐。
默认：false，可选值 true 、false
传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。
- BGM不限制时长，系统根据视频时长自动适配
payload	String	可选	透传参数
不做任何处理，仅数据传输
注：最多 1048576个字符
off_peak	Bool	可选	错峰模式，默认为：false，可选值：
- true：错峰生成视频；
- false：即时生成视频；
注1：错峰模式消耗的积分更低，具体请查看产品定价
注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分；
注3：您也可以手动取消错峰任务
watermark	Bool	可选	是否添加水印
- true：添加水印；
- false：不添加水印；
注1：目前水印内容为固定，内容由AI生成，默认不加
注2：您可以通过watermarked_url参数查询获取带水印的视频内容，详情见查询任务接口
callback_url	String	可选	Callback 协议
需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务有状态变化时，Vidu 将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致
回调返回的"status"包括以下状态：
- processing 任务处理中
- success 任务完成（如发送失败，回调三次）
- failed 任务失败（如发送失败，回调三次）
Vidu采用回调签名算法进行认证，详情见：回调签名算法
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "model": "viduq1",
    "style": "general",
    "prompt": "In an ultra-realistic fashion photography style featuring light blue and pale amber tones, an astronaut in a spacesuit walks through the fog. The background consists of enchanting white and golden lights, creating a minimalist still life and an impressive panoramic scene.",
    "duration": 5,
    "seed": 0,
    "aspect_ratio": "16:9",
    "resolution": "1080p",
    "movement_amplitude": "auto",
    "off_peak": false
}' https://api.vidu.cn/ent/v2/text2video
响应体
字段	类型	描述
task_id	String	Vidu 生成的任务ID
state	String	处理状态
可选值：
created 创建成功
queueing 任务排队中
processing 任务处理中
success 任务成功
failed 任务失败
model	String	本次调用的模型名称
prompt	String	本次调用的提示词参数
duration	Int	本次调用的视频时长参数
seed	Int	本次调用的随机种子参数
aspect_ratio	String	本次调用的 比例 参数
resolution	String	本次调用的分辨率参数
bgm	Bool	本次调用的背景音乐参数
movement_amplitude	String	本次调用的镜头动态幅度参数
payload	String	本次调用时传入的透传参数
off_peak	Bool	本次调用时是否使用错峰模式
credits	String	本次调用使用的积分数
watermark	Bool	本次提交任务是否使用水印
created_at	String	任务创建时间
{
  "task_id": "your_task_id_here",
  "state": "created",
  "model": "viduq1",
  "style": "general",
  "prompt": "In an ultra-realistic fashion photography style featuring light blue and pale amber tones, an astronaut in a spacesuit walks through the fog. The background consists of enchanting white and golden lights, creating a minimalist still life and an impressive panoramic scene.",
  "duration": 5,
  "seed": random_number,
  "aspect_ratio": "16:9",
  "resolution": "1080p",
  "movement_amplitude": "auto",
  "payload":"",
  "off_peak": false,
  "credits": credits_number,
  "created_at": "2025-01-01T15:41:31.968916Z"
}

参考生图
POST https://api.vidu.cn/ent/v2/reference2image
请求头
字段	值	描述
Content-Type	Application/json	数据交换格式
Authorization	Token{your api key}	将 {your api key} 替换为您的 token
请求体
参数名称	类型	必填	参数描述
model	String	是	模型名称
可选值：viduq1
images	Array[String]	是	图像参考
支持输入 1～7 张图片，模型将以此参数中传入的图片中的主题为参考生成具备主体一致的图片。
注1：支持传入图片 Base64 编码或图片URL（确保可访问)
注2：图片支持 png、jpeg、jpg、webp格式
注3：图片像素不能小于 128*128，且比例需要小于1:4或者4:1
注4: 且大小不超过50M
注5：请注意，base64 decode之后的字节长度需要小于50M，且编码必须包含适当的内容类型字符串，例如：
data:image/png;base64,{base64_encode}
prompt	String	是	文本提示词
视频生成的文本描述，长度不能超过 2000 个字符
seed	Int	可选	随机种子参数
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
aspect_ratio	String	可选	比例参数
默认 16:9，可选值如下：16:9、9:16、1:1、auto
- auto 生成图片的比例与第一张输入图片一致
payload	String	可选	透传参数
不做任何处理，仅数据传输
注：最多 1048576个字符
callback_url	String	可选	Callback 协议
需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务有状态变化时，Vidu 将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致
回调返回的“status”包括以下状态：
- processing 任务处理中
- success 任务完成（如发送失败，回调三次）
- failed 任务失败（如发送失败，回调三次）
Vidu采用回调签名算法进行认证，详情见：回调签名算法
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "model": "viduq1",
    "images":["your_image_url1","your_image_url2","your_image_url3"],
    "prompt": "your_prompt",
    "seed": 0,
    "aspect_ratio": "16:9",
    "payload":""
}' https://api.vidu.cn/ent/v2/reference2image
响应体
字段	类型	描述
task_id	String	Vidu 生成的任务ID
state	String	处理状态
可选值：
created 创建成功
queueing 任务排队中
processing 任务处理中
success 任务成功
failed 任务失败
model	String	本次调用的模型名称
prompt	String	本次调用的提示词参数
images	Array[String]	本次调用的图像参数
seed	Int	本次调用的随机种子参数
aspect_ratio	String	本次调用的比例参数
callback_url	String	本次调用时传入的回调参数
payload	String	本次调用时传入的透传参数
credits	Int	本次调用的积分数
created_at	String	任务创建时间
{
  "task_id": "your_task_id_here",
  "state": "created",
  "model": "viduq1",
  "images": ["your_image_url1","your_image_url2","your_image_url3"],
  "prompt": "your_prompt",
  "seed": 0,
  "aspect_ratio": "16:9",
  "payload": "",
  "credits": 2,
  "created_at": "2025-09-08T09:53:22.083033428Z"
}


查询任务接口
GET https://api.vidu.cn/ent/v2/tasks/{id}/creations
请求头
字段	值	描述
Content-Type	application/json	数据交换格式
Authorization	Token {your api key}	将 {your api key} 替换为您的 token
请求体
参数名称	类型	必填	参数描述
id	String	是	任务id，由创建任务接口创建成功返回
curl -X GET -H "Authorization: Token {your_api_key}" https://api.vidu.cn/ent/v2/tasks/{your_id}/creations
响应体
字段	子字段	类型	描述
id		String	任务ID
state		String	处理状态
可选值：
created 创建成功
queueing 任务排队中
processing 任务处理中
success 任务成功
failed 任务失败
err_code		String	错误码，具体见错误码表
credits		Int	该任务消耗的积分数量，单位：积分
payload		String	本次任务调用时传入的透传参数
bgm		Bool	本次任务调用是否使用bgm
off_peak		Bool	本次任务调用是否使用错峰模式
creations		Array	生成物结果
id	String	生成物id，用来标识不同的生成物
url	String	生成物URL， 一个小时有效期
cover_url	String	生成物封面，一个小时有效期
watermarked_url	String	带水印的生成物url，一小时有效期
{
  "id":"your_task_id",
  "state": "success",
  "err_code": "",
  "credits": 4,
  "payload": "",
  "creations": [
    {
      "id": "your_creations_id",
      "url": "your_generated_results_url",
      "cover_url": "your_generated_results_cover_url"
      "watermarked_url": "your_generated_results_watermarked_url"
    }
  ]
}


取消任务接口
POST https://api.vidu.cn/ent/v2/tasks/{id}/cancel
请求头
字段	值	描述
Content-Type	application/json	数据交换格式
Authorization	Token {your api key}	将 {your api key} 替换为您的 token
请求体
参数名称	类型	必填	参数描述
id	String	是	任务id，由创建任务接口创建成功返回
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "id": "your_task_id_here"
}' https://api.vidu.cn/ent/v2/tasks/{your_id}/cancel
响应体
字段	类型	描述
取消任务成功返回空值取消任务失败返回错误码，详情见：错误码
成功示例
{}
失败示例
{
    "code": 400,
    "reason": "BadRequest",
    "message": "task state is scheduled, can not cancel",
    "metadata": {
        "trace_id": "04e5c2fe159ff7c574acd0424e78c35f"
    }
}


产品定价
Vidu 开放平台提供丰富的视频生成 API 能力，以满足企业和开发者的多样需求。目前，API 调用采用积分计价方式。
积分单价
积分	价格（RMB）
1	0.3125
可选充值套餐价格
商务定制，大量采购，请联系 platform@vidu.cn 我们将竭诚为您服务
金额（RMB）	购得积分数	可用并发数	有效期
500	1600 积分	≤ 5	1年
1000	3200 积分	≤ 5	1年
2000	6400 积分	≤ 5	1年
5000	16000 积分	≤ 5	1年
20000	64000 积分	≤ 5	1年
注意事项：
生成失败的视频不计积分消耗（包含内容审核拒绝导致的失败）；
可用并发数≤ 5，指最多可以使用5个并发任务。在极少数情况下，系统负载可能导致您的并发量低于最大值。
积分消耗明细
视频生成能力
ViduQ2
ViduQ2图生与首尾帧积分消耗明细：
模型	清晰度	时长	积分消耗	错峰积分消耗
viduq2-turbo	720p	2S	1	1
viduq2-turbo	720p	3S	2	1
viduq2-turbo	720p	4S	3	2
viduq2-turbo	720p	5S	4	2
viduq2-turbo	720p	6S	5	3
viduq2-turbo	720p	7S	6	3
viduq2-turbo	720p	8S	7	4
viduq2-turbo	1080p	2S	5	3
viduq2-turbo	1080p	3S	6	3
viduq2-turbo	1080p	4S	7	4
viduq2-turbo	1080p	5S	8	4
viduq2-turbo	1080p	6S	9	5
viduq2-turbo	1080p	7S	10	5
viduq2-turbo	1080p	8S	11	6
viduq2-pro	720p	2S	3	2
viduq2-pro	720p	3S	4	2
viduq2-pro	720p	4S	5	3
viduq2-pro	720p	5S	6	3
viduq2-pro	720p	6S	7	4
viduq2-pro	720p	7S	8	4
viduq2-pro	720p	8S	9	5
viduq2-pro	1080p	2S	8	4
viduq2-pro	1080p	3S	10	5
viduq2-pro	1080p	4S	12	6
viduq2-pro	1080p	5S	14	7
viduq2-pro	1080p	6S	16	8
viduq2-pro	1080p	7S	18	9
viduq2-pro	1080p	8S	20	10
注意：积分计算目前暂不支持小数点计数，Q2模型在错峰模式下按向上取整计费，其他模型保持不变。
ViduQ1
能力	模型版本	时长	清晰度	风格	积分消耗	错峰积分消耗
参考-视频	viduq1	5S	1080p	通用	8	4
图-视频	viduq1	5S	1080p	通用	8	4
图-视频	viduq1-classic	5S	1080p	通用	8	4
首尾帧	viduq1	5S	1080p	通用	8	4
首尾帧	viduq1-classic	5S	1080p	通用	8	4
文-视频	viduq1	5S	1080p	通用	8	4
文-视频	viduq1	5S	1080p	动漫	8	4
Vidu2.0
能力	模型版本	时长	清晰度	风格	积分消耗	错峰积分消耗
图-视频	vidu2.0	4S	360p	通用	2	1
图-视频	vidu2.0	4S	720p	通用	4	2
图-视频	vidu2.0	4S	1080p	通用	10	5
图-视频	vidu2.0	8S	720p	通用	10	5
首尾帧	vidu2.0	4S	360p	通用	2	1
首尾帧	vidu2.0	4S	720p	通用	4	2
首尾帧	vidu2.0	4S	1080p	通用	10	5
首尾帧	vidu2.0	8S	720p	通用	10	5
参考-视频	vidu2.0	4S	360p	通用	8	4
参考-视频	vidu2.0	4S	720p	通用	8	4
Vidu1.5
能力	模型版本	时长	清晰度	风格	积分消耗	错峰积分消耗
图-视频	vidu1.5	4S	360P	通用	4	2
图-视频	vidu1.5	4S	720P	通用	10	5
图-视频	vidu1.5	4S	1080P	通用	20	10
图-视频	vidu1.5	8S	720P	通用	20	10
首尾帧	vidu1.5	4S	360P	通用	4	2
首尾帧	vidu1.5	4S	720P	通用	10	5
首尾帧	vidu1.5	4S	1080P	通用	20	10
首尾帧	vidu1.5	8S	720P	通用	20	10
参考-视频	vidu1.5	4S	360P	通用	8	4
参考-视频	vidu1.5	4S	720P	通用	20	10
参考-视频	vidu1.5	4S	1080P	通用	40	20
参考-视频	vidu1.5	8S	720P	通用	40	20
文-视频	vidu1.5	4S	360P	通用	4	2
文-视频	vidu1.5	4S	720P	通用	10	5
文-视频	vidu1.5	4S	1080P	通用	20	10
文-视频	vidu1.5	8S	720P	通用	20	10
文-视频	vidu1.5	4S	360P	动漫	4	2
文-视频	vidu1.5	4S	720P	动漫	10	5
文-视频	vidu1.5	4S	1080P	动漫	20	10
文-视频	vidu1.5	8S	720P	动漫	20	10
智能超清-尊享	-	1S	1080P	-	1	-
智能超清-尊享	-	1S	2K	-	2	-
智能超清-尊享	-	1S	4K	-	4	-
智能超清-尊享	-	1S	8K	-	16	-
音频生成能力
能力	模型版本	时长	积分消耗
文生音频	audio1.0	小于5秒	1
文生音频	audio1.0	小于10秒	2
可控文生音效	audio1.0	小于5秒	1
可控文生音效	audio1.0	小于10秒	2
图像生成能力
能力	积分消耗
参考生图	2
场景模板
我们提供了“拥抱”、“亲吻”、“万物生花”等多种热度高、效果稳定的场景模板，并持续更新中； 不同的场景模板，积分消耗规则不同； 详情请访问：场景示例中心
微调
微调模型的价格和您选择的 Batch Size 以及 Step 有关，消耗明细如下
Batch Size	步数（Step）	所需积分	预估耗时/小时
4	1000	20000	8
4	2000	40000	14
4	3000	60000	18
4	4000	80000	24
8	1000	40000	14
8	2000	80000	24
8	3000	120000	36
8	4000	160000	50
训练好的微调模型，每次推理调用消耗 4 积分。
当前微调功能处于内测阶段。如需体验，请发送邮件至 platform@vidu.cn 申请权限。
推荐提示词
基于输出提示词的个数来计费。每输出 5 个提示词，扣 1 积分。
对口型
基于对口型视频时长来计费。每生成 5 秒视频，扣 2 积分。
视频延长
目前支持延长4s，最多60s，原视频分辨率不同，延长时的积分消耗不同，具体见下表：
时长	清晰度	积分消耗
4S	360p	2
4S	720p	4
4S	1080p	10
