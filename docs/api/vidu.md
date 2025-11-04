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
注1：字符长度不能超过 2000 个字符
注2：若使用is_rec推荐提示词参数，模型将不考虑此参数所输入的提示词
is_rec	Bool	可选	是否使用推荐提示词
- true：是，由系统自动推荐提示词，并使用提示词内容生成视频，推荐提示词数量=1
- false：否，根据输入的prompt生成视频
注意：启用推荐提示词后，每个任务多消耗10积分
duration	Int	可选	视频时长
viduq2-pro 默认为 5，可选：1、2、3、4、5、6、7、8
viduq2-turbo 默认为 5，可选：1、2、3、4、5、6、7、8
viduq1 默认为 5，可选：5
viduq1-classic 默认为 5，可选：5
vidu2.0 默认为 4，可选：4、8
vidu1.5 默认为 4，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
- viduq2-pro 1-8秒：默认 720p，可选：540p、720p、1080p
- viduq2-turbo 1-8秒：默认 720p，可选：540p、720p、1080p
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
wm_position	Int	可选	水印位置，表示水印出现在图片的位置，可选项为：
1：左上角
2：右上角
3：右下角
4：左下角
默认为：3
wm_url	String	可选	水印内容，此处为图片URL
不传时，使用默认水印：内容由AI生成
meta_data	String	可选	元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：
{
"Label": "your_label","ContentProducer": "yourcontentproducer","ContentPropagator": "your_content_propagator","ProduceID": "yourproductid", "PropagateID": "your_propagate_id","ReservedCode1": "yourreservedcode1", "ReservedCode2": "your_reserved_code2"
}
该参数为空时，默认使用vidu生成的元数据标识
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
credits	Int	本次调用使用的积分数
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
可选值：viduq2、viduq1、vidu2.0、vidu1.5
- viduq2：最新模型
- viduq1：画面清晰，平滑转场，运镜稳定
- vidu2.0：生成速度快
- vidu1.5：动态幅度大
images	Array[String]	是	图像参考
支持上传1～7张图片，模型将以此参数中传入的图片中的主题为参考生成具备主体一致的视频。
注1： viduq2 viduq1 支持上传1～7张图片
注2： vidu2.0 vidu1.5支持上传1～3张图片
注3：支持传入图片 Base64 编码或图片URL（确保可访问）
注4：图片支持 png、jpeg、jpg、webp格式
注5：图片像素不能小于 128*128，且比例需要小于1:4或者4:1，且大小不超过50M。
注6：请注意，base64 decode之后的字节长度需要小于10M，且编码必须包含适当的内容类型字符串，例如：
data:image/png;base64,{base64_encode}
prompt	String	是	文本提示词
生成视频的文本描述。
注：字符长度不能超过 2000 个字符
duration	Int	可选	视频时长参数，默认值依据模型而定：
viduq2：默认5秒，可选：1-8
viduq1：默认5秒，可选：5
vidu2.0：默认4秒，可选：4
vidu1.5: 默认4秒，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
aspect_ratio	String	可选	比例
默认 16:9，可选值如下：16:9、9:16、4:3、3:4、1:1
注：4:3、3:4仅支持q2模型
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
viduq2 （1-8秒）：默认 720p, 可选：540p、720p、1080p
viduq1 （5秒）：默认 1080p, 可选：1080p
vidu2.0 （4秒）：默认 360p, 可选：360p、720p
vidu1.5（4秒）：默认 360p，可选：360p、720p、1080p
vidu1.5（8秒）：默认 720p，可选：720p
movement_amplitude	String	可选	运动幅度
默认 auto，可选值：auto、small、medium、large
注：使用q2模型时该参数不生效
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
wm_position	Int	可选	水印位置，表示水印出现在图片的位置，可选项为：
1：左上角
2：右上角
3：右下角
4：左下角
默认为：3
wm_url	String	可选	水印内容，此处为图片URL
不传时，使用默认水印：内容由AI生成
meta_data	String	可选	元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：
{
"Label": "your_label","ContentProducer": "yourcontentproducer","ContentPropagator": "your_content_propagator","ProduceID": "yourproductid", "PropagateID": "your_propagate_id","ReservedCode1": "yourreservedcode1", "ReservedCode2": "your_reserved_code2"
}
该参数为空时，默认使用vidu生成的元数据标识
callback_url	String	可选	Callback 协议
需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务有状态变化时，Vidu 将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致
回调返回的"status"包括以下状态：
- processing 任务处理中
- success 任务完成（如发送失败，回调三次）
- failed 任务失败（如发送失败，回调三次）
Vidu采用回调签名算法进行认证，详情见：回调签名算法
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "model": "viduq2",
    "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-1.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-2.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-3.png"],
    "prompt": "Santa Claus and the bear hug by the lakeside.",
    "duration": 5,
    "seed": 0,
    "aspect_ratio": "3:4",
    "resolution": "540p",
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
credits	Int	本次调用使用的积分数
watermark	Bool	本次提交任务是否使用水印
created_at	String	任务创建时间
{
  "task_id": "your_task_id_here",
  "state": "created",
  "model": "viduq2",
  "images": ["https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-1.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-2.png","https://prod-ss-images.s3.cn-northwest-1.amazonaws.com.cn/vidu-maas/template/reference2video-3.png"],
  "prompt": "Santa Claus and the bear hug by the lakeside.",
  "duration": 5,
  "seed": random_number,
  "aspect_ratio": "3:4",
  "resolution": "540p",
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
注1：字符长度不能超过 2000 个字符
注2：若使用is_rec推荐提示词参数，模型将不考虑此参数所输入的提示词
is_rec	Bool	可选	是否使用推荐提示词
- true：是，由系统自动推荐提示词，并使用提示词内容生成视频，推荐提示词数量=1
- false：否，根据输入的prompt生成视频
注意：启用推荐提示词后，每个任务多消耗10积分
duration	Int	可选	视频时长参数，默认值依据模型而定：
- viduq2-pro 默认为 5，可选：1、2、3、4、5、6、7、8
- viduq2-turbo 默认为 5，可选：1、2、3、4、5、6、7、8
- viduq1 和 viduq1-classic 默认 5 秒，可选：5
- vidu2.0 默认 4 秒，可选：4、8
- vidu1.5 默认 4 秒，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
- viduq2-pro 1-8秒：默认 720p，可选：540p、720p、1080p
- viduq2-turbo 1-8秒：默认 720p，可选：540p、720p、1080p
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
wm_position	Int	可选	水印位置，表示水印出现在图片的位置，可选项为：
1：左上角
2：右上角
3：右下角
4：左下角
默认为：3
wm_url	String	可选	水印内容，此处为图片URL
不传时，使用默认水印：内容由AI生成
meta_data	String	可选	元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：
{
"Label": "your_label","ContentProducer": "yourcontentproducer","ContentPropagator": "your_content_propagator","ProduceID": "yourproductid", "PropagateID": "your_propagate_id","ReservedCode1": "yourreservedcode1", "ReservedCode2": "your_reserved_code2"
}
该参数为空时，默认使用vidu生成的元数据标识
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
credits	Int	本次调用使用的积分数
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
可选值：viduq2 、viduq1 、vidu1.5
- viduq2：最新模型
- viduq1：画面清晰，平滑转场，运镜稳定
- vidu1.5：动态幅度大
style	String	可选	风格
默认 general，可选值：general、anime
general：通用风格，可以通过提示词来控制风格
anime：动漫风格，仅在动漫风格表现突出，可以通过不同的动漫风格提示词来控制
注：使用q2模型时该参数不生效
prompt	String	是	文本提示词
生成视频的文本描述。
注：字符长度不能超过 2000 个字符
duration	Int	可选	视频时长参数，默认值依据模型而定：
- viduq2 : 默认5秒，可选：1-8
- viduq1 : 默认5秒，可选：5
- vidu1.5 : 默认4秒，可选：4、8
seed	Int	可选	随机种子
当默认不传或者传0时，会使用随机数替代
手动设置则使用设置的种子
aspect_ratio	String	可选	比例
默认 16:9，可选值：16:9、9:16、3:4、4:3、1:1
注：3:4、4:3仅支持q2模型
resolution	String	可选	分辨率参数，默认值依据模型和视频时长而定：
- viduq2(1-8秒)：默认 720p，可选：540p、720p、1080p
- viduq1(5秒)：默认 1080p，可选：1080p
- vidu1.5(4秒)：默认 360p，可选：360p、720p、1080p
- vidu1.5(8秒)：默认 720p，可选：720p
movement_amplitude	String	可选	运动幅度
默认 auto，可选值：auto、small、medium、large
注：使用q2模型时该参数不生效
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
wm_position	Int	可选	水印位置，表示水印出现在图片的位置，可选项为：
1：左上角
2：右上角
3：右下角
4：左下角
默认为：3
wm_url	String	可选	水印内容，此处为图片URL
不传时，使用默认水印：内容由AI生成
meta_data	String	可选	元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：
{
"Label": "your_label","ContentProducer": "yourcontentproducer","ContentPropagator": "your_content_propagator","ProduceID": "yourproductid", "PropagateID": "your_propagate_id","ReservedCode1": "yourreservedcode1", "ReservedCode2": "your_reserved_code2"
}
该参数为空时，默认使用vidu生成的元数据标识
callback_url	String	可选	Callback 协议
需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务有状态变化时，Vidu 将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致
回调返回的"status"包括以下状态：
- processing 任务处理中
- success 任务完成（如发送失败，回调三次）
- failed 任务失败（如发送失败，回调三次）
Vidu采用回调签名算法进行认证，详情见：回调签名算法
curl -X POST -H "Authorization: Token {your_api_key}" -H "Content-Type: application/json" -d '
{
    "model": "viduq2",
    "style": "general",
    "prompt": "In an ultra-realistic fashion photography style featuring light blue and pale amber tones, an astronaut in a spacesuit walks through the fog. The background consists of enchanting white and golden lights, creating a minimalist still life and an impressive panoramic scene.",
    "duration": 5,
    "seed": 0,
    "aspect_ratio": "4:3",
    "resolution": "540p",
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
credits	Int	本次调用使用的积分数
watermark	Bool	本次提交任务是否使用水印
created_at	String	任务创建时间

{
  "task_id": "your_task_id_here",
  "state": "created",
  "model": "viduq2",
  "style": "general",
  "prompt": "In an ultra-realistic fashion photography style featuring light blue and pale amber tones, an astronaut in a spacesuit walks through the fog. The background consists of enchanting white and golden lights, creating a minimalist still life and an impressive panoramic scene.",
  "duration": 5,
  "seed": random_number,
  "aspect_ratio": "4:3",
  "resolution": "540p",
  "movement_amplitude": "auto",
  "payload":"",
  "off_peak": false,
  "credits": credits_number,
  "created_at": "2025-01-01T15:41:31.968916Z"
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
      "cover_url": "your_generated_results_cover_url",
      "watermarked_url": "your_generated_results_watermarked_url"
    }
  ]
}

产品定价
Vidu 开放平台提供丰富的视频生成 API 能力，以满足企业和开发者的多样需求。目前，API 调用采用积分计价方式。
积分单价
积分	价格（RMB）
1	0.03125
可选充值套餐价格
商务定制，大量采购，请联系 platform@vidu.cn 我们将竭诚为您服务
金额（RMB）	购得积分数	可用并发数	有效期
500	16000 积分	≤ 5	1年
1000	32000 积分	≤ 5	1年
2000	64000 积分	≤ 5	1年
5000	160000 积分	≤ 5	1年
20000	640000 积分	≤ 5	1年
注意事项：
生成失败的视频不计积分消耗（包含内容审核拒绝导致的失败）；
可用并发数≤ 5，指最多可以使用5个并发任务。在极少数情况下，系统负载可能导致您的并发量低于最大值。
积分消耗明细
视频生成能力
ViduQ2
ViduQ2定价一句话版本如下，需要查询具体定价信息可以查看文档：Q2定价信息
错峰价格为正常生成价格的一半，如果积分存在小数则向上取整
能力	模型	分辨率	定价
图生&首尾帧	Q2-turbo	540P	6积分起，每秒+2积分
图生&首尾帧	Q2-turbo	720P	8积分起，第二秒10积分，第三秒开始每秒+10积分
图生&首尾帧	Q2-turbo	1080P	35积分起，每秒+10积分
图生&首尾帧	Q2-pro	540P	8积分起，第二秒10积分，第三秒开始每秒+5积分
图生&首尾帧	Q2-pro	720P	15积分起，每秒+10积分
图生&首尾帧	Q2-pro	1080P	55积分起，每秒+15积分
文生	Q2	540p	10积分起，每秒+2积分
文生	Q2	720p	15积分起，每秒+5积分
文生	Q2	1080p	20积分起，每秒+10积分
参考生	Q2	540p	15积分起，每秒+5积分
参考生	Q2	720p	25积分起，每秒+5积分
参考生	Q2	1080p	75积分起，每秒+10积分
ViduQ1
能力	模型版本	时长	清晰度	风格	积分消耗	错峰积分消耗
参考-视频	viduq1	5S	1080p	通用	80	40
图-视频	viduq1	5S	1080p	通用	80	40
图-视频	viduq1-classic	5S	1080p	通用	80	40
首尾帧	viduq1	5S	1080p	通用	80	40
首尾帧	viduq1-classic	5S	1080p	通用	80	40
文-视频	viduq1	5S	1080p	通用	80	40
文-视频	viduq1	5S	1080p	动漫	80	40
Vidu2.0
能力	模型版本	时长	清晰度	风格	积分消耗	错峰积分消耗
图-视频	vidu2.0	4S	360p	通用	20	10
图-视频	vidu2.0	4S	720p	通用	40	20
图-视频	vidu2.0	4S	1080p	通用	100	50
图-视频	vidu2.0	8S	720p	通用	100	50
首尾帧	vidu2.0	4S	360p	通用	20	10
首尾帧	vidu2.0	4S	720p	通用	40	20
首尾帧	vidu2.0	4S	1080p	通用	100	50
首尾帧	vidu2.0	8S	720p	通用	100	50
参考-视频	vidu2.0	4S	360p	通用	80	40
参考-视频	vidu2.0	4S	720p	通用	80	40
Vidu1.5
能力	模型版本	时长	清晰度	风格	积分消耗	错峰积分消耗
图-视频	vidu1.5	4S	360P	通用	40	20
图-视频	vidu1.5	4S	720P	通用	100	50
图-视频	vidu1.5	4S	1080P	通用	200	100
图-视频	vidu1.5	8S	720P	通用	200	100
首尾帧	vidu1.5	4S	360P	通用	40	20
首尾帧	vidu1.5	4S	720P	通用	100	50
首尾帧	vidu1.5	4S	1080P	通用	200	100
首尾帧	vidu1.5	8S	720P	通用	200	100
参考-视频	vidu1.5	4S	360P	通用	80	40
参考-视频	vidu1.5	4S	720P	通用	200	100
参考-视频	vidu1.5	4S	1080P	通用	400	200
参考-视频	vidu1.5	8S	720P	通用	400	200
文-视频	vidu1.5	4S	360P	通用	40	20
文-视频	vidu1.5	4S	720P	通用	100	50
文-视频	vidu1.5	4S	1080P	通用	200	100
文-视频	vidu1.5	8S	720P	通用	200	100
文-视频	vidu1.5	4S	360P	动漫	40	20
文-视频	vidu1.5	4S	720P	动漫	100	50
文-视频	vidu1.5	4S	1080P	动漫	200	100
文-视频	vidu1.5	8S	720P	动漫	200	100