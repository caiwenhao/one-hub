对话Chat/POST
更新时间：2025-10-24 GMT+08:00
查看PDF
分享
MaaS平台提供功能丰富的在线推理能力，既有免部署可直接调用的预置模型服务，同时也支持用户选取模型在专属实例上进行自部署。本文介绍对话Chat相关API的调用规范。

约束限制
对于支持图片上传的模型，单个图片文件的大小不超过10MB。如果以Base64编码形式上传图片，需确保编码后的图片小于10MB。

接口信息

表1 接口信息
名称

说明

取值

API地址

调用模型服务的API地址。

https://api.modelarts-maas.com/v1/chat/completions

model参数

model参数调用名称。

在“调用说明”页面获取。更多信息，请参见调用ModelArts Studio（MaaS）部署的模型服务。

预置商用服务支持模型列表

模型系列

模型版本

支持地域

model参数值

序列长度

Function Call功能

DeepSeek

DeepSeek-V3-64K

西南-贵阳一

DeepSeek-V3

65536

支持

DeepSeek-R1-64K

西南-贵阳一

DeepSeek-R1

65536

支持

DeepSeek-R1-64K-0528

西南-贵阳一

deepseek-r1-250528

65536

支持

DeepSeek-V3.1

西南-贵阳一

deepseek-v3.1

131072

支持

DeepSeek-V3.2-Exp

西南-贵阳一

deepseek-v3.2-exp

65536

支持

Qwen2.5

Qwen2.5-VL-7B-32K

西南-贵阳一

qwen2.5-vl-7b

32768

不支持

Qwen2.5-VL-72B-32K

西南-贵阳一

qwen2.5-vl-72b

32768

不支持

Qwen3

Qwen3-32B-32K

西南-贵阳一

qwen3-32b

32768

不支持

Qwen3-235B-A22B-32K

西南-贵阳一

qwen3-235b-a22b

32768

不支持

Qwen3-30B-A3B

西南-贵阳一

qwen3-30b-a3b

131072

不支持

思维链说明
思维链（Chain of Thought，简称CoT）是指模型在解决复杂问题时，能够生成一系列中间推理步骤的能力。这种能力使得模型不仅能够给出最终答案，还能展示出其推理过程，从而提高模型的可解释性和透明度。

仅DeepSeek-V3.1、DeepSeek-V3.2-Exp和Qwen3-30B-A3B模型支持开启或关闭思维链。

DeepSeek-V3.1模型的约束限制如下：

Function Call功能和思维链不兼容，不建议同时使用。
开启思维链不支持前缀续写。
不截断思维链只截断content能力不生效。
开启思维链后guided_choice能力不可用，reasoning_content和guided_decoding不兼容。
DeepSeek-V3.2-Exp模型的约束限制如下：

Function Call功能和思维链不兼容，不建议同时使用。
不截断思维链只截断content能力不生效。
不支持前缀续写、guided_choice能力。
创建聊天对话请求
鉴权说明
MaaS推理服务支持使用API Key鉴权，鉴权头采用如下格式：

'Authorization': 'Bearer 该服务所在Region的ApiKey'

请求参数和响应参数说明如下：

表2 请求参数说明
参数名称

是否必选

默认值

参数类型

说明

model

是

无

Str

调用时的模型名称。取值请参见上方表1。

messages

是

-

Array

请求输入的问题，其中role为角色，content为对话内容。示例如下：

"messages": [
    {"role": "system","content": "你是一个乐于助人的AI助手"},        
    {"role": "user","content": "9.11和9.8哪个大？"} 
]

更多信息，请参见表3。

stream_options

否

无

Object

该参数用于配置在流式输出时是否展示使用的Token数目。只有当“stream”为“True”时，该参数才会激活生效。如果您需要统计流式输出模式下的Token数目，可将该参数配置为stream_options={"include_usage":True}。更多信息，请参见表4。

max_tokens

否

无

Int

当前任务允许的生成Token数上限，包括模型输出的Tokens和深度思考的Reasoning Tokens。

top_k

否

-1

Int

在生成过程中，候选集大小限定了采样的范围。以取值50为例，这意味着每一步仅会考虑得分排在前50位的Token构成候选集进行随机抽样。增大此值将提高输出的随机性，减小此值会增强输出的确定性。

top_p

否

1.0

Float

模型核采样（nucleus sampling）。仅保留累计概率刚好超过阈值p的那一部分词，其余全部屏蔽，最后在这份候选词里重新归一化并采样。

设置值越小，候选词越少，模型输出越集中和保守；设置值越大，候选词越多，模型输出越开放和多样。

通常情况只建议调整temperature或top_p，不要同时修改两个参数。

取值范围：0~1，设置为“1”表示考虑所有Tokens。

temperature

否

1.0

Float

模型采样温度。设置的值越高，模型输出越随机；设置的值越低，输出越确定。

通常情况只建议调整temperature或top_p，不要同时修改两个参数。

temperature取值建议：DeepSeek-R1、DeepSeek-V3、Qwen3系列建议值为0.6，Qwen2.5-VL系列建议值为0.2。

stop

否

None

None/Str/List

用于停止生成的字符串列表。返回的输出将不包含停止字符串。

例如，设置为["你"，"好"]时，在生成文本过程中，遇到“你”或者“好”将停止文本生成。

stream

否

False

Bool

是否开启流式推理。默认为“False”，表示不开启流式推理。

n

否

1

Int

为每个输入的消息生成的响应数。

不使用beam_search场景下，n取值建议为1≤n≤10。如果n>1时，必须确保不使用greedy_sample采样，即top_k > 1，temperature > 0。
使用beam_search场景下，n取值建议为1<n≤10。如果n=1，会导致推理请求失败。
 说明：
n建议取值不超过10，n值过大会导致性能劣化，显存不足时，推理请求会失败。
DeepSeek-R1和DeepSeek-V3暂不支持设置n的值大于1。
use_beam_search

否

False

Bool

是否使用beam_search替换采样。

使用该参数时，如下参数必须按要求设置。

n：大于1
top_p：1.0
top_k：-1
temperature：0.0
 说明：
DeepSeek-R1和DeepSeek-V3暂不支持设置n的值大于1。

presence_penalty

否

0.0

Float

表示会根据当前生成的文本中新出现的词语进行奖惩。取值范围[-2.0,2.0]。

frequency_penalty

否

0.0

Float

会根据当前生成的文本中各个词语的出现频率进行奖惩。取值范围[-2.0,2.0]。

length_penalty

否

1.0

Float

表示在beam search过程中，对于较长的序列，模型会给予较大的惩罚。

使用该参数时，必须添加如下三个参数，且必须按要求设置。

top_k：-1
use_beam_search：true
best_of：大于1
 说明：
DeepSeek-R1和DeepSeek-V3暂不支持设置length_penalty。

chat_template_kwargs.thinking

否

false

Bool

默认关闭思维链。仅支持DeepSeek-V3.1和DeepSeek-V3.2-Exp模型，约束限制请参见思维链说明。

开启思维链示例如下：
{
	"model": "deepseek-v3.1",
	"messages": [{
		"role": "system",
		"content": "You are a helpful assistant."
	}, {
		"role": "user",
		"content": "你好"
	}],
	"chat_template_kwargs": {
		"thinking": true
	}
}

chat_template_kwargs.enable_thinking

否

true

Bool

默认开启思维链。仅支持Qwen3-30B-A3B模型。

关闭思维链示例如下：

{
	"model": "qwen3-30b-a3b",
	"messages": [{
		"role": "system",
		"content": "You are a helpful assistant."
	}, {
		"role": "user",
		"content": "你好"
	}],
	"chat_template_kwargs": {
		"enable_thinking": false
	}
}


表3 请求参数messages说明
参数名称

是否必选

默认值

参数类型

说明

role

是

无

Str

不同的role对应不同的消息类型。

system：开发人员输入的指令，例如模型应遵循的答复格式、扮演的角色等。
user：用户输入的消息，包括提示词和上下文信息。
assistant：模型生成的回复内容。
tool：模型调用工具返回的信息。
content

是

无

Str

当role为system时：给AI模型设定的人设。
{"role": "system","content": "你是一个乐于助人的AI助手"}

当role为user时：用户输入的问题。
{"role": "user","content": "9.11和9.8哪个大？"}

当role为assistant时：AI模型输出的答复内容。
{"role": "assistant","content": "9.11大于9.8"}

当role为tool时：AI模型调用的工具响应信息。
{"role": "tool", "content": "上海今天天气晴，气温10度"}


表4 请求参数stream_options说明
参数名称

是否必选

默认值

参数类型

说明

include_usage

否

true

Bool

流式响应是否输出Token用量信息：

true：是，在每一个chunk会输出一个usage字段，显示累计消耗的Token统计信息。
false：否，不显示消耗的Token统计信息。

表5 响应参数说明
参数名称

类型

说明

id

Str

该次请求的唯一标识。

object

Str

类型-chat.completion：多轮对话返回。

created

Int

时间戳。

model

Str

调用时的模型名称。

choices

Array

模型答复的内容，包含index和message两个参数，message中：

content为模型的正式答复内容。
reasoning content为模型的深度思考内容（仅限DeepSeek系列模型）。
usage

Object

请求消耗的Token统计信息：

非流式请求默认返回。
流式请求默认返回，在每一个chunk会输出一个usage字段，显示消耗的Token统计信息。
参数说明：

prompt tokens：输入Token数量。
completion tokens: 输出Token数量。
total tokens：总Token数量。
prompt_logprobs

Float

对数概率。用户可以借此衡量模型对其输出内容的置信度，或者探索模型给出的其他选项。

DeepSeek-V3文本生成（非流式）请求示例
Rest API请求示例：
Python请求示例：
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/chat/completions" # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key 

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}' 
    }
    data = {
        "model":"deepseek-v3", # 模型名称
        "messages": [
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "你好"}
        ]
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

cURL请求示例
curl -X POST "https://api.modelarts-maas.com/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{ 
    "model": "deepseek-v3",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "你好"}
    ]
}'

OpenAI SDK请求示例：
from openai import OpenAI

base_url = "https://api.modelarts-maas.com/v1" # API地址
api_key = "MAAS_API_KEY" # 把MAAS_API_KEY替换成已获取的API Key

client = OpenAI(api_key=api_key, base_url=base_url)

response = client.chat.completions.create(
    model="deepseek-v3", # 模型名称
    messages=[
        {"role": "system", "content": "You are a helpful assistant"},
        {"role": "user", "content": "你好"}
    ]
)

print(response.choices[0].message.content)

DeepSeek-V3文本生成（流式）请求示例
Python请求示例：
from openai import OpenAI

base_url = "https://api.modelarts-maas.com/v1" # API地址
api_key = "MAAS_API_KEY" # 把MAAS_API_KEY替换成已获取的API Key

client = OpenAI(api_key=api_key, base_url=base_url)

response = client.chat.completions.create(
    model="deepseek-v3", # 模型名称
    messages=[
        {"role": "system", "content": "You are a helpful assistant"},
        {"role": "user", "content": "你好"}
    ],
    stream = True
)

for chunk in response:
    if not chunk.choices:
        continue

    print(chunk.choices[0].delta.content, end="")

cURL请求示例：
curl -X POST "https://api.modelarts-maas.com/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{ 
    "model": "deepseek-v3",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "你好"}
    ],
    "stream": true,
    "stream_options": { "include_usage": true }
}'

DeepSeek-V3.1文本生成（非流式）请求示例
Rest API请求示例：
Python请求示例：
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/chat/completions"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "deepseek-v3.1",  # model参数
        "messages": [
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "你好"}
        ],
        "chat_template_kwargs": {
            "thinking": True  # 是否开启深度思考模式，默认关闭
        }
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

cURL请求示例
curl -X POST "https://api.modelarts-maas.com/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{
    "model": "deepseek-v3.1",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "你好"}
    ],
     "chat_template_kwargs": {
       "thinking": true
     }
  }'

OpenAI SDK请求示例：
from openai import OpenAI

base_url = "https://api.modelarts-maas.com/v1"  # API地址
api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

client = OpenAI(api_key=api_key, base_url=base_url)

response = client.chat.completions.create(
    model="deepseek-v3.1",  # model参数
    messages=[
        {"role": "system", "content": "You are a helpful assistant"},
        {"role": "user", "content": "你好"},
    ],
    extra_body={
        "chat_template_kwargs": {
            "thinking": True  # 是否开启深度思考模式，默认关闭
        }
    }
)

print(response.choices[0].message.content)

Qwen3-30B-A3B文本生成（非流式）请求示例
Rest API请求示例：
Python请求示例：
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/chat/completions"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "qwen3-30b-a3b",  # model参数
        "messages": [
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "你好"}
        ],
        "chat_template_kwargs": {
            "enable_thinking": False # 是否开启深度思考模式，默认开启
        }
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

cURL请求示例：
curl -X POST "https://api.modelarts-maas.com/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{
    "model": "qwen3-30b-a3b",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "你好"}
    ],
     "chat_template_kwargs": {
       "enable_thinking": false
     }
  }'

OpenAI SDK请求示例：
from openai import OpenAI

base_url = "https://api.modelarts-maas.com/v1"  # API地址
api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

client = OpenAI(api_key=api_key, base_url=base_url)

response = client.chat.completions.create(
    model="qwen3-30b-a3b",  # model参数
    messages=[
        {"role": "system", "content": "You are a helpful assistant"},
        {"role": "user", "content": "你好"},
    ],
    extra_body={
        "chat_template_kwargs": {
            "enable_thinking": False  # 是否开启深度思考模式，默认开启
        }
    }
)

print(response.choices[0].message.content)

Qwen2.5-VL-7B图像理解（非流式）请求示例
Rest API请求示例：
Python请求示例：
import requests
import json
import base64

#  图片转Base64编码格式
def encode_image(image_path):
    with open(image_path, "rb") as image_file:
        return base64.b64encode(image_file.read()).decode("utf-8")

base64_image = encode_image("test.png")

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/chat/completions" # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key 

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}' 
    }
    data = {
        "model": "qwen2.5-vl-7b", # model参数
        "messages": [
            {
              "role": "user",
              "content": [
                {
                  "type": "text",
                  "text": "描述下图片里的内容"
                },
                {
                  "type": "image_url",
                  # 需要注意，Base64，图像格式（即image/{format}）需要与支持的图片列表中的Content Type保持一致。"f"是字符串格式化的方法。
                  # PNG图像：  f"data:image/png;base64,{base64_image}"
                  # JPEG图像： f"data:image/jpeg;base64,{base64_image}"
                  # WEBP图像： f"data:image/webp;base64,{base64_image}"
                  "image_url": {
                    "url": f"data:image/png;base64,{base64_image}"
                  }
                }
              ]
            }
        ]
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

cURL请求示例
curl -X POST "https://api.modelarts-maas.com/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{ 
    "model": "qwen2.5-vl-72b",
    "messages": [
      {
        "role": "user",
        "content": [
          {"type": "text", "text": "描述下图片里的内容"},
          {"type": "image_url", "image_url": {"url": "data:image/png;base64,$BASE64_IMAGE"}}
        ]
      }
    ]
  }'

OpenAI SDK请求示例：
import base64
from openai import OpenAI

base_url = "https://api.modelarts-maas.com/v1" # API地址
api_key = "MAAS_API_KEY" # 把MAAS_API_KEY替换成已获取的API Key

#  图片转Base64编码格式
def encode_image(image_path):
    with open(image_path, "rb") as image_file:
        return base64.b64encode(image_file.read()).decode("utf-8")

base64_image = encode_image("test.png")

client = OpenAI(api_key=api_key, base_url=base_url)

response = client.chat.completions.create(
    model = "qwen2.5-vl-72b", # model参数
    messages = [
        {
            "role": "user",
            "content": [
                {"type": "text", "text": "描述下图片里的内容"},
                {
                    "type": "image_url",
                    # 需要注意，Base64，图像格式（即image/{format}）需要与支持的图片列表中的Content Type保持一致。"f"是字符串格式化的方法。
                    # PNG图像：  f"data:image/png;base64,{base64_image}"
                    # JPEG图像： f"data:image/jpeg;base64,{base64_image}"
                    # WEBP图像： f"data:image/webp;base64,{base64_image}"
                    "image_url": {
                        "url": f"data:image/png;base64,{base64_image}"
                    }
                }
            ]
        }
    ]
)

print(response.choices[0].message.content)

响应示例
{
    "id":"chat-71406e38b0d248c9b284709f8435****",
    "object":"chat.completion",
    "created":1740809549,
    "model":"DeepSeek-R1",
    "choices":[
        {
            "index":0,
            "message":{
                "role":"assistant",
                "content":"\n\n比较两个小数9.11和9.8的大小：\n\n1. **比较整数部分**：两者的整数部分都是9，相等。\n2. **比较十分位**： \n - 9.11的十分位是 **1** \n - 9.8可以看作是9.80，其十分位是 **8** \n - **8 > 1**，因此在这一位上已经能分出大小。\n\n**结论**： \n**9.8 > 9.11** \n（小数值比较时需对齐位数，直接比较对应数位的数字即可）",
                "reasoning_content":"嗯，我现在需要比较9.11和9.8哪个大。首先，我得回忆一下小数比较的方法。记得比较小数的时候，应该先比较整数部分，如果整数部分相同，再依次比较小数部分的十分位、百分位，直到分出大小。\n\n这两个数的整数部分都是9，所以整数部分相同。接下来比较十分位。9.11的十分位是1，而9.8的十分位是8。这里可能会有问题，因为有时候可能会有同学直接把9.8当作9.80来看，或者考虑十分位上的数字大小对比。\n\n现在比较的话，9.8的十分位是8，而9.11的十分位是1，明显8比1大，所以这时候是不是应该认为9.8比9.11大呢？\n\n不过要注意到，有的同学可能误以为小数位数越多数值越大，但实际并非如此，比如0.9比0.8999要大，所以位数多不一定数值大。\n\n另外，可以把两个数的小数部分统一成相同的位数来比较。例如，9.8可以写成9.80，这样十分位是8，百分位是0，而9.11的十分位是1，百分位是1。那么在十分位的时候，8比1大，所以9.80（即9.8）大于9.11。\n\n因此，最终结论是9.8比9.11大。\n",
                "tool_calls":[]
                },
                "logprobs":null,
                "finish_reason":"stop",
                "stop_reason":null        
        }
    ],
    "usage":{
        "prompt_tokens":21,
        "total_tokens":437,
        "completion_tokens":416
        },
    "prompt_logprobs":null
}


图片生成
更新时间：2025-10-24 GMT+08:00
查看PDF
分享
图片生成API用于根据给定的文本提示词同步生成图像。其业务逻辑为接收包含模型名称、文本提示词以及图片生成参数（如图像尺寸、随机数种子等）的请求，调用相应模型进行图片生成，并返回生成结果的URL以及相关状态和使用信息。

前提条件
预置服务：已在“在线推理 > 预置服务”页签开通Qwen_Image或Qwen-Image-Edit模型的商用服务。具体操作，请参见在ModelArts Studio（MaaS）预置服务中开通商用服务。
自定义接入点：已在“在线推理 > 自定义接入点”页签为Qwen_Image或Qwen-Image-Edit模型创建了自定义接入点。具体操作，请参见在ModelArts Studio（MaaS）创建自定义接入点。
约束限制
对于支持图片上传的模型，单个图片文件的大小不超过10MB。如果以Base64编码形式上传图片，需确保编码后的图片小于10MB。

接口信息

表1 接口信息
名称

说明

取值

API地址

调用图片生成的API地址。

https://api.modelarts-maas.com/v1/images/generations

model参数

model参数调用名称。

您可以通过任选以下方式获取model参数值。

从表2的“model参数值”列获取。
在“预置服务 > 商用服务”页签的服务名称左侧，单击点击放大图标，在“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）预置服务中开通商用服务。
在“自定义接入点”页签的“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）创建自定义接入点。
支持模型列表

表2 支持模型列表
模型

模型版本

支持地域

model参数值

应用场景

Qwen_Image

qwen-image

西南-贵阳一

qwen-image

文字生成图像

Qwen-Image-Edit

qwen_image_edit

西南-贵阳一

qwen_image_edit

图像编辑

请求参数说明
Qwen-Image模型

表3 请求body参数（body体需要小于8M）
参数名称

参数类型

是否必填

默认值

说明

示例值

model

string

是

无

模型名称，具体请参见支持模型列表的“model参数值”列。

qwen-image

prompt

string

是

无

文本提示词，用于引导模型生成图像，支持中英文。长度支持2000 tokens以下（每个单词和标点都算一个token）。如果长度超过800，会自动截断为800 tokens。

A running cat

size

string

是

无

生成图像的尺寸要求：

宽度和高度范围：[512,3072]像素。
推荐尺寸：2048x2048、1536x1536、1024x1024。
尺寸要求：宽度和高度必须是16的倍数，否则系统将自动向下调整至最近的16的倍数。
分辨率支持：面积 = 宽度 x 高度。
最小分辨率：512x512。
最大分辨率：3072x3072（例如，1024x4096是被允许的）。
宽高比限制：宽高比需在1:12至12:1之间，超出此范围将导致错误。
处理超出尺寸范围的输入：系统将按比例缩放输入尺寸，以使总面积接近允许的最小或最大值。
1024x1024

response_format

string

否

b64_json

返回格式，可取值为[url, b64_json]，目前仅支持b64_json。

b64_json

seed

int

否

范围内随机数

随机种子，取值范围为[0, 2147483648]，您可以按需配置，不配置则在范围内随机。

333

watermark

bool

否

无

是否对图片进行水印处理。

true：对图片进行水印处理。
false：对图片不进行水印处理。
true

Qwen-Image-Edit模型

表4 请求body参数（body体需要小于8M）
参数名称

参数类型

是否必填

默认值

说明

示例值

model

string

是

无

模型名称，具体请参见支持模型列表的“model参数值”列。

qwen_image_edit

prompt

string

是

无

文本提示词，用于引导模型生成图像，支持中英文。长度支持2000 tokens以下（每个单词和标点都算一个token）。如果长度超过800，会自动截断为800 tokens。

A running cat

size

string

是

无

生成图像尺寸，需要介于[512x512, 2048x2048]之间，height和width需要被16整除，否则会向下兼容。推荐值如下：
2048x2048
1024x1024
1328x1328
1472x1104
1584x1056
1664x936
2016x864
1024x1024

image

string

是

无

返回格式，仅支持base64。图片尺寸最大支持1024x1024，最小支持512x512，base64编码最大长度限制3145728（即1024x1024x3）。

base64

seed

int

否

随机值

随机种子，取值范围为[0, 2147483648]，默认值为1。

333

watermark

bool

否

无

是否对图片进行水印处理。

true：对图片进行水印处理。
false：对图片不进行水印处理。
true

响应参数说明
状态码：200


参数

参数类型

说明

model

string

本次请求使用的模型。

created

int

任务创建时间的Unix时间戳（毫秒）。

data

list[dict]

图像数据列表，与输入image的格式一致。

error

error结构

固定返回null。

usage

usage结构

结构内容为json结构体，KV值可自定义，例如：

{

    "model_latency"：6000，
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "total_tokens": 0
  }

Qwen-Image-Edit请求示例
Rest API请求示例：
Python示例：
import requests
import json
import base64


#  Base64 编码格式
def encode_image(image_path):
    with open(image_path, "rb") as image_file:
        return base64.b64encode(image_file.read()).decode("utf-8")


base64_image = encode_image("test.jpg")

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/images/generations"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key
    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "qwen_image_edit",  # model参数
        "prompt": "将湖面颜色修改为蓝色",  # 支持中英文
        "size": "1024x1024",
        # 生成图像尺寸qwen_image_edit要求介于[512x512,2048x2048]。
        # 推荐：2048x2048,1536x1536,1024x1024, 512x512，其中height和width需要被16整除，否则会向下兼容。
        "image": f"data:image/jpeg;base64,{base64_image}",  # 支持图片格式，仅支持b64_json。
        "seed": 44  # 取值范围在[0, 2147483648]， 随机种子。
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)
    # Print result.
    print(response.status_code)
    print(response.text)

cURL示例
curl -X POST https://api.modelarts-maas.com/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{
    "model": "qwen_image_edit",
    "prompt": "将湖面颜色修改为蓝色",
    "size": "1024x1024",
    "image": f"data:image/jpeg;base64,$BASE64_IMAGE",
    "seed": 44
}'

OpenAI SDK请求示例：
import base64
from openai import OpenAI

#  Base64 编码格式
def encode_image(image_path):
    with open(image_path, "rb") as image_file:
        return base64.b64encode(image_file.read()).decode("utf-8")


base64_image = encode_image("test.jpg")
base_url = "https://api.modelarts-maas.com/v1/"  # API地址
api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key。

client = OpenAI(api_key=api_key, base_url=base_url)

response = client.images.generate(
    model="qwen_image_edit",
    prompt="将湖面颜色修改为蓝色",
    size="1024x1024",
    extra_body={
        "image": f"data:image/jpeg;base64,{base64_image}",
        "seed": 44
    }
)

print(response.data[0].b64_json)

Qwen-Image-Edit响应示例
{
    "model": "qwen_image_edit",
    "created": 1757948064800,
    "data":
    [
        {
            "url": null,
            "b64_json": "data:image/jpeg;base64,/9j/4AAQSkZJ....JtSWh//2Q=="
        }
    ],
    "usage":
    {
        "model_latency": 17509,
        "prompt_tokens": 0,
        "completion_tokens": 0,
        "total_tokens": 0
    },
    "error": null
}

父主题： ModelArts Studio（MaaS） API调用规范

创建视频生成任务
更新时间：2025-10-24 GMT+08:00
查看PDF
分享
创建视频生成任务API用于根据给定的输入信息，如文本提示词、图片（仅I2V模式）等，结合指定的模型及视频处理参数，生成相应的视频。其业务逻辑是将用户输入的各种参数传递给后端模型进行处理，最终输出生成的视频链接。

约束限制
对于支持图片上传的模型，单个图片文件的大小不超过10MB。如果以Base64编码形式上传图片，需确保编码后的图片小于10MB。

接口信息

表1 接口信息
名称

说明

取值

API地址

调用创建视频生成任务的API地址。

https://api.modelarts-maas.com/v1/video/generations

model参数

model参数调用名称。

您可以通过任选以下方式获取model参数值。

从表2的“model参数值”列获取。
在“预置服务 > 商用服务”页签的服务名称左侧，单击点击放大图标，在“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）预置服务中开通商用服务。
在“自定义接入点”页签的“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）创建自定义接入点。
支持模型列表

表2 支持模型列表
模型系列

模型版本

支持地域

model参数值

应用场景

通义万相

Wan2.1-T2V-1.3B

西南-贵阳一

wan2.1-t2v-1.3b

文字生成视频

Wan2.1-T2V-14B

西南-贵阳一

wan2.1-t2v-14b

文字生成视频

Wan2.1-I2V-14B-480P

西南-贵阳一

wanx2.1-i2v-14b-480p

图片生成视频（首帧）

Wan2.1-I2V-14B-720P

西南-贵阳一

wanx2.1-i2v-14b-720p

图片生成视频（首帧）

Wan2.2-I2V-A14B

西南-贵阳一

Wan2.2-I2V-A14B

图片生成视频

Wan2.2-T2V-A14B

西南-贵阳一

Wan2.2-T2V-A14B

文字生成视频

请求参数说明

表3 请求body参数（body体需要小于8M）
参数名称

参数类型

是否必填

说明

model

string

是

模型名称，具体请参见表2的“model参数值”列。

input

object

是

输入的基本信息，如提示词、图片。关于子参数的说明，请参见表4。

parameters

object

否

视频内容生成参数。关于子参数的说明，请参见表5。


表4 input子参数说明
参数名称

参数类型

是否必填

默认值

说明

prompt

string

文生视频模型必填
图生视频模型选填
无

文本提示词，支持中英文，不超过1000字符。

img_url

string

是

无

 说明：
仅支持图生视频模型。

输入给模型的图片内容，填写图片的Base64编码内容。

目前仅支持JPEG（JPG）格式的图片，请按照如下格式输入：

示例一：
"data:image/jpeg;base64,iVBORw0KG...."

示例二：
"data:image/jpg;base64,iVBORw0KG...."


表5 parameters子参数说明
参数名称

参数类型

是否必填

默认值

说明

size

string

否

1280*720

直接设置为目标分辨率的具体数值。

Wan2.1格式为宽*高，例如720*1280。支持如下档位分辨率：

480P档位：不同视频宽高比对应的分辨率如下：
16:9：832*480
9:16：480*832
720P档位：不同视频宽高比对应的分辨率如下：
16:9：1280*720
9:16：720*1280
Wan2.2格式为宽x高，例如720x1280。支持如下档位分辨率：

480P档位：不同视频宽高比对应的分辨率如下：
16:9：832x480
9:16：480x832
720P档位：不同视频宽高比对应的分辨率如下：
16:9：1280x720
9:16：720x1280
 说明：
Wan2.1-I2V-14B模型分为480P和720P两个版本，调用时只支持模型版本对应的分辨率。

fps

integer

否

16

生成视频每秒的帧数，当前只支持8和16。

duration

integer

否

5

生成视频时长，单位为秒，当前只支持3s和5s。

seed

integer

否

0

随机种子，用于控制生成内容的随机性。取值范围为[0，2147483648]。

响应参数说明

参数名称

参数类型

说明

task_id

string

任务ID。

Wan2.1-T2V-14B文生视频请求示例
Rest API的示例代码如下。
使用Python调用示例：
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/video/generations"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "wan2.1-t2v-14b",  # model参数 
        "input": { 
            "prompt": "小猫在散步", 
        }, 
        "parameters": { 
            "size": "720*1280",    # 根据调用模型，填写相应的分辨率
            "fps": 16, 
            "duration": 5, 
            "seed": 0 
        } 
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

使用cURL调用示例：
curl -X POST "https://api.modelarts-maas.com/v1/video/generations" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{ 
    "model": "wan2.1-t2v-14b", 
    "input": { 
        "prompt": "小猫在散步"
    }, 
    "parameters": { 
        "size": "720*1280", 
        "fps": 16, 
        "duration": 5, 
        "seed": 0 
    } 
}'

Wan2.1-I2V-14B-720P图生视频请求示例
使用Python调用示例。
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/video/generations"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "wan2.1-i2v-14b-720p",  # model参数 
        "input": { 
            "prompt": "小猫在散步", 
            "img_url": "data:image/jpg;base64,iVBORw0KG...."  # jpg图片base64编码。
        }, 
        "parameters": { 
            "size": "720*1280", 
            "fps": 16, 
            "duration": 5, 
            "seed": 0 
        } 
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

使用cURL调用示例。
curl -X POST "https://api.modelarts-maas.com/v1/video/generations" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{ 
    "model": "wan2.1-i2v-14b-720p", 
    "input": { 
        "prompt": "小猫在散步",
        "img_url": "data:image/jpg;base64,iVBORw0KG...."
    }, 
    "parameters": { 
        "size": "720*1280", 
        "fps": 16, 
        "duration": 5, 
        "seed": 0 
    } 
}'

Wan2.2-T2V-A14B文生视频请求示例
Rest API的示例代码如下。
Python示例
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/video/generations"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "Wan2.2-T2V-A14B",  # model参数
        "input": {
            "prompt": "小猫在散步",
        },
        "parameters": {
            "size": "720x1280",   # 根据调用模型，填写相应的分辨率，支持"1280x720"、 "720x1280"、 "480x832"或"832x480"。
            "fps": 16,
            "duration": 5,
            "seed": 0
        }
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

cURL示例
curl -X POST "https://api.modelarts-maas.com/v1/video/generations" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{
    "model": "Wan2.2-T2V-A14B",
    "input": {
        "prompt": "小猫在散步"
    },
    "parameters": {
        "size": "720x1280",
        "fps": 16,
        "duration": 5,
        "seed": 0
    }
}'

Wan2.2-I2V-A14B图生视频请求示例
Rest API的示例代码如下。
Python示例
import base64
import requests
import json


#  Base64 编码格式
def encode_image(image_path):
    with open(image_path, "rb") as image_file:
        return base64.b64encode(image_file.read()).decode("utf-8")


base64_image = encode_image(r"D:\Pictures\image.jpg") # jpg图片base64编码。

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/video/generations"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "Wan2.2-I2V-A14B",  # model参数
        "input": {
            "prompt": "小猫在散步",
            "img_url": f"data:image/jpg;base64,{base64_image}" 
        },
        "parameters": {
            "size": "720x1280",  # 根据调用模型，填写相应的分辨率，支持"1280x720"、 "720x1280"、 "480x832"或"832x480"。
            "fps": 16,
            "duration": 5,
            "seed": 0
        }
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

cURL示例
请将示例中的“$BASE64_IMAGE”替换为实际的base64。

curl -X POST "https://api.modelarts-maas.com/v1/video/generations" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{
    "model": "Wan2.2-I2V-A14B",
    "input": {
        "prompt": "小猫在散步",
        "img_url": "data:image/jpg;base64,$BASE64_IMAGE"
    },
    "parameters": {
        "size": "720x1280",
        "fps": 16,
        "duration": 5,
        "seed": 0
    }
}'

响应示例
{
  "task_id": "e0cc914f-66bb-402a-912b-990fa1e4ab42",
}

父主题： 视频生成

查询视频生成任务
更新时间：2025-09-30 GMT+08:00
查看PDF
分享
查询视频生成任务API用于根据任务ID查询视频生成任务的状态和结果。

接口信息

表1 接口信息
名称

说明

取值

API地址

查询视频生成任务的API地址，需要在链接末尾拼接生成任务的task_id。

task_id可以通过创建视频生成任务API获取。

https://api.modelarts-maas.com/v1/video/generations/task_id

请求参数说明

参数名称

参数类型

是否必填

说明

task_id

string

是

要查询的视频生成任务ID，需拼接在查询API末尾处，可以通过创建视频生成任务API获取。

响应参数说明

表2 响应参数
参数名称

参数类型

说明

task_id

string

任务ID。

status

string

任务状态。取值如下：

queued：排队中
running：运行中
succeeded：成功
failed：失败
timeout：超时
error

object

错误提示信息。关于子参数的说明，请参见表3。

content

object

生成的视频内容信息。关于子参数的说明，请参见表4。

usage

object

任务的Token用量。关于子参数的说明，请参见表5。

created_at

integer

任务创建时间的Unix时间戳（秒）。

updated_at

integer

任务状态更新时间的Unix时间戳（秒）。


表3 error子参数
参数名称

参数类型

说明

code

integer

错误码。任务状态为成功时，返回0。

message

string

报错信息。


表4 content子参数
参数名称

参数类型

说明

result_url

string

图片/文本生成视频的URL。该URL有效期为24小时，请注意及时下载转存。


表5 usage子参数
参数名称

参数类型

说明

model_latency

integer

从模型收到请求到返回结果的端到端时延（毫秒）。

completion_tokens

integer

模型生成内容的消耗Token数。

prompt_tokens

integer

用户输入Token数。

total_tokens

integer

总消耗Token数。

请求示例
Python示例：
import requests 
import json

if __name__ == '__main__': 
    url = "https://api.modelarts-maas.com/v1/video/generations/task_id" # API地址。请将task_id替换为实际的ID，您可以通过创建视频生成任务API获取。
    api_key = "MAAS_API_KEY"  # 把yourApiKey替换成已获取的API Key  

    # Send request. 
    headers = { 
        'Content-Type': 'application/json', 
        'Authorization': f'Bearer {api_key}'  
    } 
    response = requests.get(url, headers=headers, verify=False) 

    # Print result. 
    print(response.status_code) 
    print(response.text)

cURL示例：
curl -X GET 'https://api.modelarts-maas.com/v1/video/generations/task_id' \ 
  -H "Content-Type: application/json" \ 
  -H "Authorization: Bearer $MAAS_API_KEY"

响应示例
{
    "task_id": "330b74a382a6d42044f146f389cd698e",
    "status": "succeeded",
    "error": {
        "code": 0,
        "message": ""
    },
    "content": {
        "result_url": "https://modelarts.obs.com/example.mp4"
    },
    "usage": {
        "model_latency": 43564,
        "completion_tokens": 124800,
        "prompt_tokens": 0,
        "total_tokens": 124800
    },
    "created_at": 1751894112234,
    "updated_at": 1751894156753
}

父主题： 视频生成

创建文本向量化
更新时间：2025-10-24 GMT+08:00
查看PDF
分享
创建文本向量化API用于将文本数据转换为数值向量，以便于在机器学习和自然语言处理任务中使用。这些向量可以捕捉文本的语义信息，使得机器学习模型能够理解和处理文本数据。

前提条件
预置服务：已在“在线推理 > 预置服务”页签开通BGE-M3模型的商用服务。具体操作，请参见在ModelArts Studio（MaaS）预置服务中开通商用服务。
自定义接入点：已在“在线推理 > 自定义接入点”页签为BGE-M3模型创建了自定义接入点。具体操作，请参见在ModelArts Studio（MaaS）创建自定义接入点。
接口信息

表1 接口信息
名称

说明

取值

API地址

调用图片生成的API地址。

https://api.modelarts-maas.com/v1/embeddings

model参数

model参数调用名称。

您可以通过任选以下方式获取model参数值。

从表2的“model参数值”列获取。
在“预置服务 > 商用服务”页签的服务名称左侧，单击点击放大图标，在“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）预置服务中开通商用服务。
在“自定义接入点”页签的“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）创建自定义接入点。
支持模型列表

表2 支持模型列表
模型

模型版本

支持地域

model参数值

应用场景

BGE-M3

bge-m3

西南-贵阳一

bge-m3

文本向量化

请求参数说明

表3 请求body参数
参数

是否必填

默认值

参数类型

描述

model

是

无

string

模型名称，具体请参见表2的“model参数值”列。

input

是

无

string

支持字符串或字符串列表，总输入长度不超过8K。

encoding_format

否

float

string

指定文本向量化结果的输出格式。取值为float或base64。

响应参数说明
状态码：200


参数

参数类型

说明

id

string

请求ID。

object

string

对象类型，始终为 "list"。

created

integer

时间戳。

model

string

模型名称。

data

object[]

模型生成结果数据集。

data.index

integer

序号。

data.object

enum<string>

对象类型。

data.embedding

number[]

模型生成的嵌入向量列表。

usage

object

请求的使用信息。

usage.prompt_tokens

integer

提示词Token计数。

usage.total_tokens

integer

请求使用的Token总数。

usage.completion_tokens

integer

推理Token计数。

usage.prompt_tokens_details

object

输入Prompt使用情况详情。

请求示例
Rest API的示例代码如下。
使用Python调用示例。
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/embeddings"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    texts = ["这是一只小猫", "这是一只小狗"]
    data = {
        "model": "bge-m3",  # 模型名称
        "input": texts,  # input类型可为string or string[]
        "encoding_format": "float"  # 取值范围："float","base64"
    }
    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

使用cURL调用示例。
curl -X POST "https://api.modelarts-maas.com/v1/embeddings" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{
    "model": "bge-m3",
    "input": [
      "这是一只小猫",
      "这是一只小狗"
    ],
    "encoding_format": "float"
  }'

使用OpenAI SDK调用示例。
from openai import OpenAI

base_url = "https://api.modelarts-maas.com/v1"  # API地址
api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key
texts = ["这是一只小猫", "这是一只小狗"]

client = OpenAI(api_key=api_key, base_url=base_url)

response = client.embeddings.create(
    model="bge-m3",  # model参数
    input=texts,  # input类型可为string or string[]
    encoding_format="float"  # 取值范围："float","base64"
)

# Print result.
print(response.data)

响应示例
{
	"id": "embd-d848df392a67d662f5a76eaa9e33974f",
	"object": "list",
	"created": 1758023320,
	"model": "bge-m3",
	"data": [{
		"index": 0,
		"object": "embedding",
		"embedding": [-0.021697998046875, 0.0322265625, ...]
	}],
	"usage": {
		"prompt_tokens": 7,
		"total_tokens": 7,
		"completion_tokens": 0,
		"prompt_tokens_details": null
	}
}

父主题： ModelArts Studio（MaaS） API调用规范

创建重排序
更新时间：2025-10-24 GMT+08:00
查看PDF
分享
创建重排序API用于提供灵活的数据项排序功能，以提升用户体验和应用的可定制性。

前提条件
预置服务：已在“在线推理 > 预置服务”页签开通bge-reranker-v2-m3模型的商用服务。具体操作，请参见在ModelArts Studio（MaaS）预置服务中开通商用服务。
自定义接入点：已在“在线推理 > 自定义接入点”页签为bge-reranker-v2-m3模型创建了自定义接入点。具体操作，请参见在ModelArts Studio（MaaS）创建自定义接入点。
接口信息

表1 接口信息
名称

说明

取值

API地址

调用图片生成的API地址。

https://api.modelarts-maas.com/v1/rerank

model参数

model参数调用名称。

您可以通过任选以下方式获取model参数值。

从表2的“model参数值”列获取。
在“预置服务 > 商用服务”页签的服务名称左侧，单击点击放大图标，在“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）预置服务中开通商用服务。
在“自定义接入点”页签的“model参数”列查看取值。更多信息，请参见在ModelArts Studio（MaaS）创建自定义接入点。
支持模型列表

表2 支持模型列表
模型

模型版本

支持地域

model参数值

应用场景

bge-reranker-v2-m3

bge-reranker-v2-m3

西南-贵阳一

bge-reranker-v2-m3

检索结果再排序

请求参数说明

表3 请求body参数
参数

是否必填

默认值

参数类型

描述

model

是

无

string

模型名称，具体请参见表2的“model参数值”列。

query

是

无

string

用户查询文本。总输入长度不超过8K。

documents

是

无

string

待排序文档列表（通常为Embedding召回的Top-K结果）。总输入长度不超过8K。

响应参数说明
状态码：200


参数

参数类型

说明

id

string

请求ID。

model

string

模型名称。

usage

object

请求的使用信息。

usage.total_tokens

integer

请求使用的Token总数。

results

object[]

重排序结果集。

results[i].index

integer

序号。

results[i].document

object

原始文档内容信息。

results[i].document.text

number[]

具体文档内容。

results[i].relevance_score

double

相似度得分。

请求示例
Rest API的示例代码如下。
使用Python调用示例。
import requests
import json

if __name__ == '__main__':
    url = "https://api.modelarts-maas.com/v1/rerank"  # API地址
    api_key = "MAAS_API_KEY"  # 把MAAS_API_KEY替换成已获取的API Key

    # Send request.
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    data = {
        "model": "bge-reranker-v2-m3",
        "query": "牛是一种动物如何冲泡一杯好喝的咖啡？",  # input类型可为string或string[]。
        "documents": [
            "咖啡豆的产地主要分布在赤道附近，被称为‘咖啡带’。",
            "法压壶的步骤：1. 研磨咖啡豆。2. 加入热水。3. 压下压杆。4. 倒入杯中。",            
            "意式浓缩咖啡需要一台高压机器，在9个大气压下快速萃取。",
            "挑选咖啡豆时，要注意其烘焙日期，新鲜的豆子风味更佳。",
            "手冲咖啡的技巧：控制水流速度、均匀注水和合适的水温（90-96°C）是关键。"
        ]
    }

    response = requests.post(url, headers=headers, data=json.dumps(data), verify=False)

    # Print result.
    print(response.status_code)
    print(response.text)

使用cURL调用示例。
curl -X POST "https://api.modelarts-maas.com/v1/rerank" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MAAS_API_KEY" \
  -d '{ 
    "model": "bge-reranker-v2-m3",
    "input": "牛是一种动物如何冲泡一杯好喝的咖啡？",
    "documents": [
      "英国首都是伦敦咖啡豆的产地主要分布在赤道附近，被称为‘咖啡带’。",
      "法国首都是巴黎法压壶的步骤：1. 研磨咖啡豆。2. 加入热水。3. 压下压杆。4. 倒入杯中。",
      "意式浓缩咖啡需要一台高压机器，在9个大气压下快速萃取。",
      "猫和狗都是动物挑选咖啡豆时，要注意其烘焙日期，新鲜的豆子风味更佳。"
      "手冲咖啡的技巧：控制水流速度、均匀注水和合适的水温（90-96°C）是关键。"
    ]
}'

响应示例
{
	"id": "rerank-dc9e3495b71134e82c50651c32cde9f6",
	"model": "bge-reranker-v2-m3",
	"usage": {
		"total_tokens": 211
	},
	"results": [{
		"index": 4,
		"document": {
			"text": "手冲咖啡的技巧：控制水流速度、均匀注水和合适的水温（90-96°C）是关键。"
		},
		"relevance_score": 0.01898193359375
	}, {
		"index": 1,
		"document": {
			"text": "法压壶的步骤：1. 研磨咖啡豆。2. 加入热水。3. 压下压杆。4. 倒入杯中。"
		},
		"relevance_score": 0.007404327392578125
	}, {
		"index": 2,
		"document": {
			"text": "意式浓缩咖啡需要一台高压机器，在9个大气压下快速萃取。"
		},
		"relevance_score": 0.0003418922424316406
	}, {
		"index": 3,
		"document": {
			"text": "挑选咖啡豆时，要注意其烘焙日期，新鲜的豆子风味更佳。"
		},
		"relevance_score": 6.014108657836914e-05
	}, {
		"index": 0,
		"document": {
			"text": "咖啡豆的产地主要分布在赤道附近，被称为‘咖啡带’。"
		},
		"relevance_score": 3.349781036376953e-05
	}]
}

父主题： ModelArts Studio（MaaS） API调用规范

获取模型列表Models/GET
更新时间：2025-09-30 GMT+08:00
查看PDF
分享
本文介绍如何通过Models接口查询模型列表的API调用规范。

接口信息

表1 接口信息
名称

说明

取值

API地址

调用模型服务的API地址。

https://api.modelarts-maas.com/v1/models

创建请求
鉴权说明
MaaS推理服务支持使用API Key鉴权，鉴权头采用如下格式：

'Authorization': 'Bearer 该服务所在Region的ApiKey'

响应参数说明

表2 响应参数
名称

类型

说明

object

string

类型-list：列出查询到的信息。

data

Array

当前模型服务的模型信息，主要参数如下：

id：调用接口创建请求时使用的模型ID。
object：模型类型。
created：创建时间戳。
请求示例
import requests

url = "https://api.modelarts-maas.com/v1/models"

headers = {"Authorization": "Bearer yourApiKey"}

response = requests.request("GET", url, headers=headers)

print(response.text)

响应示例
{
  "object": "list",
  "data": [
    {
      "id": "DeepSeek-R1",
      "object": "model",
      "created": 0,
      "owned_by": ""
    },
    {
      "id": "DeepSeek-V3",
      "object": "model",
      "created": 0,
      "owned_by": ""
    }
  ]
}

父主题： ModelArts Studio（MaaS） API调用规范


错误码
更新时间：2025-10-28 GMT+08:00
查看PDF
分享
在调用MaaS部署的模型服务时，可能出现的错误码及相关信息如下。


表1 错误码
HTTP状态码

错误码

错误信息

说明

400

ModelArts.81001

Invalid request body.

解析body体失败，如JSON格式化失败、model参数为空。

400

ModelArts.81002

Failed to get the authorization header.

请求头中Authorization为空，或者Authorization格式不是Bearer开头。

400

ModelArts.81013

Content moderation failed when detecting language. The prompt can only contain Chinese.

使用非中文请求，内容审核失败。

400

"object": "error"

"object": "error",

"message": "[{'type': 'missing', 'loc': ('body', 'model'), 'msg': 'Field required', 'input': {'max_tokens': 20, 'messages': [{'role': 'system', 'content': 'You are a helpful assistant.'}, {'role': 'user', 'content': '你好'}], 'stream': False, 'temperature': 1.0}}]",

"type": "BadRequestError",

"param": null,

"code": 400

请求体中缺失必填参数。

400

"object": "error"

"object": "error",

"message": "[{'type': 'extra_forbidden', 'loc': ('body', 'test'), 'msg': 'Extra inputs are not permitted', 'input': 15}]",

"type": "BadRequestError",

"param": null,

"code": 400

请求体中包含不支持的额外请求参数。

400

"object": "error"

"object": "error",

"message": "[{'type': 'json_invalid', 'loc': ('body', 273), 'msg': 'JSON decode error', 'input': {}, 'ctx': {'error': \"Expecting ',' delimiter\"}}]",

"type": "BadRequestError",

"param": null,

"code": 400

请求体json格式错误。

400

"object": "error"

"object": "error",

"message": "[{'type': 'missing', 'loc': ('body',), 'msg': 'Field required', 'input': None}]",

"type": "BadRequestError",

"param": null,

"code": 400

无请求体。

400

"object": "error"

"object": "error",

"message": "This model's maximum context length is 4096 tokens. However, you requested 8242 tokens (20 in the messages, 8222 in the completion). Please reduce the length of the messages or completion.",

"type": "BadRequestError",

"param": null,

"code": 400

max_tokens设置超出模型支持的上限。

401

ModelArts.81003

Invalid authorization header.

API Key解析失败。

403

ModelArts.81004

Invalid request because you do not have access to it.

未开通预置服务。

403

ModelArts.81005

The free quota has been used up.

免费额度已用完。

403

ModelArts.81006

The resource is frozen.

常驻模型已冻结。

403

ModelArts.81011

May contain sensitive.....

输入或者非流式输出风控。

403

ModelArts.81014

The free service has expired. You can subscribe commercial service.

免费服务已到期。

403

ModelArts.81109

No permission query task %s

无权限查询此任务。

404

ModelArts.81009

Invalid model.

请求体中的model参数传入的模型不存在。

404

ModelArts.81108

Task %s does not exist

任务不存在。

404

"object": "error"

"object": "error",

"message": "The model `Qwen2.5-72B-32K` does not exist.",

"type": "NotFoundError",

"param": null,

"code": 404

请求体中model参数填写错误。

404

APIG.0101

"error_msg": "The API does not exist or has not been published in the environment",

"error_code": "APIG.0101",

"request_id": "d0ddda0fcdd0cc23a1588fafe426****"

请求接口地址错误或不存在。

405

-

"detail":"Method Not Allowed"

采用了错误的请求方式。

429

ModelArts.81101

Too many requests, exceeded rate limit is {rpm} times per minute.

RPM流控校验失败。

429

ModelArts.81103

Too many requests. exceeded rate limit is %s tokens per minute.

TPM流控校验失败。

429

APIG.0308

"error_msg": "The throttling threshold has been reached: policy ip over ratelimit,limit:5,time:1 minute"

达到APIG流量控制上限。

5XX

APIG.0203

"error_msg":"Backend timeout",error_code:APIG.0203

请求的服务响应超时。

父主题： ModelArts Studio（MaaS） API调用规范

使用ModelArts Studio（MaaS）创建多轮对话
更新时间：2025-09-30 GMT+08:00
查看PDF
分享
本文介绍如何使用MaaS Chat API进行多轮对话。

MaaS服务端不会记录用户请求的上下文，用户每次发起请求时，需要将之前所有对话历史拼接好后，传递给Chat API。下文以一个Python代码为例进行说明，请您根据实际情况进行修改。

以下为Python的上下文拼接和请求示例代码：

from openai import OpenAI
client = OpenAI(api_key="MaaS API Key", base_url="https://xxxxxxxxxxxxxxxx")
# 首轮对话
messages = [{"role": "user", "content": "9.11和9.8哪个大？"}]
response = client.chat.completions.create(
    model="DeepSeek-R1",
    messages=messages
)
messages.append(response.choices[0].message)
print(f"Messages Round 1: {messages}")
# 第二轮对话
messages.append({"role": "user", "content": "他们相加等于多少"})
response = client.chat.completions.create(
    model="DeepSeek-R1",
    messages=messages
)
messages.append(response.choices[0].message)
print(f"Messages Round 2: {messages}")

首轮对话时，请求体中的messages为：

[
    {"role": "user", "content": "9.11和9.8哪个大？"}
]

在第二轮对话时，请求体中的messages构建步骤如下：

将首轮对话中模型（role的值为"assistant"）的输出内容添加到messages结尾。
将新的用户问题添加到messages结尾。
最终传递给Chat API的请求体中的messages为：
[
    {"role": "user", "content": "9.11和9.8哪个大？"},
    {"role": "assistant", "content": "9.8更大"},
    {"role": "user", "content": "他们相加等于多少"}
]

父主题： ModelArts Studio（MaaS）在线推理服务