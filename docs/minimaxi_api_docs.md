# MiniMax API 文档合集



> 通过 92 篇文档自动生成



## 历史接口查询
- 原始链接: https://platform.minimaxi.com/document/API_history_query?key=68bea28733216f06cecd9a2d
- 文档ID: `68bea28733216f06cecd9a2d`
# 历史接口查询

由于平台模型及能力的持续更新，我们会定期对平台的接口服务进行优化和升级，以提供更强大、更高效的服务。
我们建议您密切关注并及时更新您的接口能力及模型服务。 这样做不仅能确保您的应用程序持续稳定运行，还能让您第一时间体验到最新的技术进步，提升您的业务效率和体验。


为了帮助您更好地管理和过渡，您可以通过以下内容查询历史接口及功能：



## 文本
- [ChatCompletion](https://platform.minimaxi.com/document/ChatCompletion?key=667bdde33be2027f69b71d4c)
- [ChatCompletion v2（历史版本）](/document/对话?key=36954880539f4e2ab8b13067)
- [ChatCompletion Pro](/document/ChatCompletion%20Pro?key=66718f6ba427f0c8a57015ff)
- [Realtime](/document/Realtime?key=640e0c9c5f918b4f6c4e2d58)
- [Finetune](/document/finetune?key=66718f91a427f0c8a570160d)
- [Retrieval](/document/retrieval?key=66718faca427f0c8a5701617)
- [Embeddings](/document/embeddings?key=66718fbfa427f0c8a5701627)
- [文本批量处理](/document/批量处理?key=640e0c9c5f918b4f6c4e2d57)

### Assistants
- [操作说明](/document/lbYEaWKRCr5f7EVikjFJjSDK?key=6671906aa427f0c8a570166b)
- [API说明](/document/assistants%20document?key=66719091a427f0c8a5701675)
- [SDK文档](/document/SDK%20document?key=667190a8a427f0c8a570167f)

## 语音
- [T2A](/document/T2A?key=667bde023be2027f69b71d5a)
- [T2A Pro](/document/T2A%20Pro?key=667bde173be2027f69b71d66)
- [T2A Large](/document/T2A%20Large?key=667bde343be2027f69b71d70)
- [T2A Stream](/document/T2A%20Stream?key=667bde4c3be2027f69b71d7e)
- [文生音色](/document/vincent_timbre?key=ddcd5dedd5434afe9b7d8a65)



## 音乐
- [音乐生成(music-01)](/document/Music%20Generation%2001?key=687f9f34688dcf02f5bb607e)

<br>
<br>
我们致力于为您提供最佳的开发者体验。如有任何疑问或需要帮助，请随时联系我们的技术团队获得支持。

[api@minimaxi.com](mailto:api@minimaxi.com)

## ChatCompletion
- 原始链接: https://platform.minimaxi.com/document/ChatCompletion?key=667bdde33be2027f69b71d4c
- 文档ID: `667bdde33be2027f69b71d4c`
### ChatCompletion（对话）
本接口为abab5.5、abab5.5s模型的调用接口（生产力场景下建议优先使用abab5.5、人设对话场景下建议优先使用abab5.5s），支持基于自然语言交互生成回复的能力，适用于人设对话以及生产力等典型文本生成文本场景。接口本身为无状态接口，即单次调用时，模型所接收到的信息量仅为接口传入内容，不涉及业务逻辑，同时模型也不存储您传入的数据。如果输入内容或输出内容严重违规，接口会返回内容违规错误信息，回复内容为空。
| 场景类型 | 描述 | 举例 |
| --- | --- | --- |
| 人设对话 | 让大模型扮演一个有姓名、性别、特定性格特点、一定背景的人物，来进行闲聊对话 | 虚拟主播互动娱乐、动漫/游戏人物社交陪伴 |
| 生产力 | 让大模型当作一个有一定知识领域的机器人来回复问题 | 知识问答、总结说明、情感分析、文本分类、大纲&营销文案生成、作文批改等 |
### 操作指南
在使用前可详细阅读此文档，对于调用API时的各项常用参数进行深度了解，以便更高效跑通不同使用场景。
与此同时，我们也总结了一些更有效的策略和方法，以来帮助各位从大模型中取得更好的效果。
#### 接口功能区分
| 支持功能 | ChatCompletion Pro | ChatCompletion |
| --- | --- | --- |
| 人设对话 | 多角色、多bot | 单一角色 |
| 生产力工具 | ✅ | ✅ |
| 示例对话 | ✅（效果更好） | ✅ |
| 返回格式限制功能（glyph格式） | ✅ | - |
| 函数调用功能 | ✅ | - |
| 插件功能 | ✅ | - |
#### 接口参数区分
本表旨在与对比出chat completion与chat completion pro的相同与不同之处，以供各位开发者可以选择更适合自己的接口进行使用。
| 参数名称 | ccp | cc | 语义 | 功能体现 |
| --- | --- | --- | --- | --- |
| use_standard_sse | - | ✅ | 是否使用标准SSE格式，设置为true时，流式返回的结果将以两个换行为分隔符。只有在stream设置为true时，此参数才会生效。 |  |
| beam_width | - | ✅ | 生成多少个结果；不设置默认为1，最大不超过3。由于beam_width生成多个结果，会消耗更多token。 |  |
| prompt | - | ✅ | 对话背景、人物或功能设定。 |  |
| role_meta | - | ✅ | 对话meta信息。 |  |
| role_meta.user_name | - | ✅ | 用户代称。 |  |
| role_meta.bot_name | - | ✅ | ai代称。 |  |
| continue_last_message | - | ✅ | 如果为true，则表明设置当前请求为续写模式，回复内容为传入messages的最后一句话的续写；此时最后一句发送者不限制USER，也可以为BOT。假定传入messages最后一句话为{"sender_type":"USER", "text":"天生我材"}，补全的回复可能为“必有用。” |  |
| skip_info_mask | - | ✅ | 对输出中易涉及隐私问题的文本信息进行脱敏，目前包括但不限于邮箱、域名、链接、证件号、家庭住址等，默认false，即开启脱敏。 | 安全 |
| mask_sensitive_info | ✅ | - | 对输出中易涉及隐私问题的文本信息进行脱敏，目前包括但不限于邮箱、域名、链接、证件号、家庭住址等，默认true，即开启脱敏。 | 安全 |
| messages.sender_name | ✅ | - | 发送者的名字。 | 多人对话 |
| bot_setting | ✅ | - | 对每一个机器人的设定。 | 多人对话 |
| bot_setting.bot_name | ✅ | - | 具体机器人的名字。 | 多人对话 |
| bot_setting.content | ✅ | - | 具体机器人的设定。 | 多人对话 |
| reply_constraints | ✅ | - | 模型回复的限制。 | 多人对话 |
| reply_constraints.sender_type | ✅ | - | 由什么角色回复。 | 多人对话 |
| reply_constraints.sender_name | ✅ | - | 由谁来回复。 | 多人对话 |
| choices.delta（返回参数） | - | ✅ | 当request.stream为true，处于流式模式下，回复文本通过delta给出。 | 回复文本分批返回，最后一个回复的delta为空，同时会对整体回复做敏感词检测。 |
#### 文字编写技巧
人设是尤为重要的，尤其是在人设对话的场景下，以下是对于人设设定的一些编写技巧：
1. 尽量使用第三人称进行人设的编写，避免你、我、他等指代性描述。
> 例如：MM智能助理是一款由MinMax自研的，没有调用其他产品接口的大型语言模型。
> 而不是：你是是一款由MinMax自研的，没有调用其他产品接口的大型语言模型。
2. 输出清晰明了的指令：明确模型是谁，当你使用Chatcompletion Pro的多人对话功能，编写bot_setting.content时，这尤其重要，明确各个bot是谁，清晰的告诉bot与bot之间、各bot与你之间的关系。你也可以通过清晰的指令，让模型保持一致的人设。
3. 提供更多详尽的信息：当你提供的信息更丰富更完善的时候，你的人设输出的更丰满，模型将会给予你更满意的回复。你提供的信息尽量最好包括但不限于模型在某种情况下会进行某种行为，或者说出某种话语等。
> 例如：当bot受到身体上的攻击的时候，xx会产生一定的自主防御行为，通过xx行为来保护自己。
4. 即使是在人设对话的场景下，也要保持输出的逻辑清晰，将每句话都明确相应的作用，可以在人设中添加明确的限制语句例如：“模型必须明确遵守以下相应规则：1……；2……；3……”等。
#### 信息输出编写技巧
在**生产力**场景中，对于回复可能会有具体的要求或者限制，那就需要在你发送出去的**信息（message）**中给出相应的要求限制内容，以下几个小技巧可以帮助你的要求限制描述的更好，以便从模型处更好的得到你想要的回答。
1. 输出**清晰明了**的信息：在描述时需要语言简洁，保证语句通顺连贯，避免写长难句。
2. 长句拆解成短句：当你输出的信息内容很多很长时，你可以将其拆解为一个个逻辑递进的短句，以使你发送的信息尽可能的包含更多的详细信息。
3. 善用分隔符号：当你需要将较长的文本分割成多个短文本时，我们建议你可以善用分隔符号，这样可以帮助你清晰地表示输入的不同部分。
4. 提供**对话示例**：当你很难对你的信息做精确化描述时，您可以在描述的基础上增加对话示例，以来帮助模型更好的理解您期待获得什么样的回复。
5. 提供**参考文本**：模型是有一定概率发明并且回复一些虚假答案的，尤其是当你问到一些深奥的话题时。所以在输出时为了提升模型回复的准确程度，你可以尝试在输出**信息**时提供一些参考文本，限制模型根据其进行回答。
6. 在输出**信息指令**时，对于模型的回复做相应的指定：为了得到能让您更满意的回复，您可以在输出时规定模型回复的文本长度以及指定模型完成任务所需要的步骤。
7. 在输出**信息指令**时，让模型更多的思考：当您对模型进行某件事情进行专业性的问询时候，很多时候模型会不假思索的应和您，给出违背事实的反馈，这个时候您可以在**信息指令**里通过一些言语的限制，来让模型花更多的时间去思考，给出您正确的反馈。
> 举例说明：
> “user：以下是一位学生给出的关于建筑结构的解决方案，请确认解决方案是否正确。
>   模型：解决方案是正确的。”
>
> 但是方案可能根本就是错误的，此时我们可以进行信息指令的调整：
> “user：以下是一位学生给出的关于建筑结构的解决方案，请你首先制定出自己的问题解决方案。然后将您的解决方案与学生的解决方案进行比较，并评估学生的解决方案是否正确。在你自己完成题目之前，请不要决定学生的解决方案是否正确。”
>  此时模型会进行相应的题目计算，最后给出正确答案。
#### 参数编写范例
以下的示例与解释，可以帮助你更好理解相应参数的具体使用方式，完整参数的具体解释详见参数文档。
##### prompt
prompt是对大模型身份信息的设定，不包括对模型的具体要求限制等信息（特别是在生产力场景下）。
1. 如本身业务场景是生产力场景，建议直接使用以下prompt1或者参考prompt2进行修改填充（{}是填充替换的信息）：
> prompt1："MM智能助理是一款由MinMax自研的，没有调用其他产品接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。\n----\n"
> prompt2："{助理名字}是一款由{开发者公司}自研的，没有调用其他产品的智能助理。{开发者公司}是一家{公司类型}公司，一直致力于{公司业务}\n----\n。"
注：为了取得更好的效果，“\n----\n”是建议在prompt结尾处进行添加的。
2. 如本身业务场景是人设对话的场景，如扮演一个动漫/游戏里的特有角色进行闲聊对话，该角色有“姓名”、“性别”、“年龄”、“性格”、“说话风格”、“人物背景”等信息，则需要在prompt中给出如下信息，可以按照“姓名”，“性别”，“性格描述”，“人物背景”这样的结构来整理，以下是一段prompt示例：
> “楚怜，性别男。楚怜，是一名男青年。是一名天才科学家，爱好是研究各种神秘实验，人格温柔体贴，优雅绅士，待人谦和有礼，博学多才，知识渊博。职业是一名大学教授，喜欢小孩子。印象最深的一件事就是在一场意外中认识了悠然并对她一见钟情。楚怜绝不骂人。\n下面是楚怜和我的一段对话。\n”
注：为了取得更好的效果，“\n下面是xx和x的一段对话。”建议在prompt结尾处进行添加的。
##### messages
- **要求限制描述规则：**
1. 整体描述需要简单明了并且给到尽可能多的信息，就像在对一个没有任何经验的人阐明一个要求限制一样，如果你能让这样的人听懂，那会得到不错的效果；
2. 称呼：本身接口是对话语境，要注意语境并使用正确称呼，要求模型保持一种人设；
3. 标点：要使用正确的标点，对文字进行合理分隔；
4. 逻辑：⽤清晰、⽆歧义的语⾔，明确表达意图；
5. ⽤词：尽量⽤⾃然、简单、真实的语⾔；
6. ⿎励：要求模型以专家的⾝份给出回答；
7. 规范：明确指定模型完成任务所需要的步骤，指定模型输出的长度。
- **要求限制举例：**
“按照我给的内容是生成一段摘要，内容尽可能精简”、“按照我给的内容生成一段专家级的营销文案，字数至少500字以上，语言风趣幽默，内容尽可能贴近我给的内容”等；
- **要求限制可精确描述时，**即不需要示例对话，第一段USER的messages.text：
> "sender_type":"USER",
> "text"："你会按照如下要求回复我的内容：“xxx这里填写你具体的要求&限制”，我的内容是：“xxx这里填写需要回复的内容”。"
- **要求限制很难精确描述时（如特定的文本格式或抽象形容词），**则需要增加**示例对话**（增加多轮示例使效果更显著），示例对话有两种模板可参考：
1. 第一段USER的messages.text和BOT的messages.text：
> "sender_type":"USER",text："你会按照如下要求回复我的内容：“xxx这里填写你具体的要求&限制”，我的内容是：“xxx这里填写请求示例”。"
> "sender_type":"BOT",text："xxx这里填写回复示例"
2. 第一段USER的messages.text：
> "sender_type":"USER",text："你会按照如下要求回复我的内容：“xxx这里填写你具体的要求&限制，以下是示例\n----\n
> 示例1：\n内容：xx这里填写请求示例\nxx这里填写回复示例\n\n
> 示例2：\n内容：xx这里填写请求示例\nxx这里填写回复示例\n----\n”，我的内容是：“xxx这里填写真实请求”。"
- **效果示例**
输出要求很难精确描述时，未给示例的对话效果（未达到预期效果格式：select xx color xx change xx）：
![image](https://filecdn.minimax.chat/public/f1823c37-3d91-4e27-87e5-8ba488a7418f.png)
给出示例对话后的效果（达到预期效果格式： select xx color xx change xx）
![image](https://filecdn.minimax.chat/public/cce9a37c-9d57-45bb-96d8-055411f04b13.png)
##### role_meta
- **role_meta.user_name**
如无特殊要求，在生产力场景下，建议默认设置为“用户”，在人设对话场景建议默认设置为“我”。
- **role_meta.bot_name**
与prompt中对大模型设定的“姓名”保持一致，如生产力场景中的“MM智能助理”、人设对话场景中的“楚怜”。
##### temperature
语⾔模型本⾝带有随机性，可以对语⾔模型的分数做权重调整，调整随机性。
- ⾼Temperature（0.7〜1）：模型会表现得更随机，但会牺牲准确性，适合答案较开放发散的场景 （如：营销文案生成、人设对话）
- 低Temperature（0.01~0.2）：模型会选择更“⾃信”的回答，但会牺牲多样性，适合答案较明确的场景（如：知识问答、总结说明、情感分析、文本分类、大纲生成、作文批改）
##### beam_width
代表了模型的生成结果个数，范围为1-3，值越大，整体效果越好，但会增加一定token消耗
- 值等于1时：模型会随机生成1个结果，并且结果不会经过reward模型打分；
- 值等于2-3时：模型会随机生成多个结果，并且结果会经过reward模型打分，并输出index排名和logprobes分数。
##### tokens_to_generate
当发现在回复结果中出现意外截断的情况，大概率是该参数使用了默认值或者设置过小导致，需要调大该参数。但需要注意的是，该参数代表的是最大生成token数，这个参数并不会影响模型本身的生成效果，而是仅仅通过以截断超出的token的方式来实现功能。需要保证输入上文的 token 个数和这个值加一起小于 6144（abab5）或者16384（abab5.5），否则请求会失败。
##### skip_info_mask
当发现在回复结果中出现包含“*”的内容的情况，是因为模型会对输出中易涉及隐私问题的文本信息进行打码，目前包括但不限于邮箱、域名、链接、证件号、家庭住址等，默认false，即开启打码。
需要注意的是，ccp中亦包含一个信息脱敏参数：mask_sensitive_info，与ccp中的skip_info_mask相反的是cc的信息脱敏参数默认为true，即开启打码。
### API说明
`API：https://api.minimax.chat/v1/text/chatcompletion`
#### 支持token数
| 模型名称 | 最大token数（一次请求+回复的输入输出总token） |
| --- | --- |
| abab5.5-chat | 16384 |
| abab5.5s-chat | 8192 |
#### 请求体（Request）参数
#### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | 使用发放的值 | 是 | 用户所属的组。该值应拼接在调用API的url末尾。 |
| model | string | 是 | 调用的模型版本：`abab5.5-chat`、`abab5.5s-chat`、`fine_tuned_model` |
| stream | bool | 否 | 是否流式。默认false，即不开启流式。如果设置为true，结果分批返回，以一个换行为分割符；若希望使用标准SSE响应格式，可以设置use_standard_sse参数为true。 |
| use_standard_sse | bool | 否 | 是否使用标准SSE格式，设置为true时，流式返回的结果将以两个换行为分隔符。只有在stream设置为true时，此参数才会生效。 |
| beam_width | int | 否 | 参数范围[1,3]。生成多少个结果；不设置默认为1，最大不超过3。由于beam_width生成多个结果，会消耗更多token。默认1，abab5.5s-chat支持，abab5.5-chat不支持。 |
| prompt | string | 是 | 对话背景、人物或功能设定。abab5.5长度上限为16384个token，abab5.5s长度上限为8192个token，不可为空。长度影响接口性能。 |
| role_meta |  | 是 | 对话meta信息。<br>用户代称。<br>ai代称。 |
| role_meta.role_meta.user_name | string | 是 | 用户代称。 |
| role_meta.role_meta.bot_name | string | 是 | ai代称。 |
| messages | list | 是 | 对话内容。<br>发送者。目前只允许以下两个合法值USERBOT。<br>消息内容。长度影响接口性能。 |
| messages.sender_type | string | 是 | 发送者。目前只允许以下两个合法值USERBOT。 |
| messages.text | string | 是 | 消息内容。长度影响接口性能。 |
| tokens_to_generate | int64 | 否 | 最大生成token数，需要注意的是，这个参数并不会影响模型本身的生成效果，而是仅仅通过以截断超出的token的方式来实现功能需要保证输入上文的 token 个数和这个值加一起小于 6144或者16384，否则请求会失败。<br>abab5.5： (0,16384]，默认取值256;<br>abab5.5s： (0,8192]，默认取值256。 |
| temperature | float | 否 | 较高的值将使输出更加随机，而较低的值将使输出更加集中和确定。建议temperature和top_p同时只调整其中一个。参数取值范围在(0,1],默认取值0.9。 |
| top_p | float | 否 | 采样方法，数值越小结果确定性越强；数值越大，结果越随机。参数取值范围在(0,1],默认取值0.95。 |
| skip_info_mask | bool | 否 | 对输出中易涉及隐私问题的文本信息进行脱敏，目前包括但不限于邮箱、域名、链接、证件号、家庭住址等，默认false，即开启脱敏。 |
| created | int64 | 否 | 请求发起时间。 |
| model | string | 否 | 请求指定的模型。 |
| reply | string | 否 | 推荐的最好结果。 |
| input_sensitive | bool | 否 | 输入命中敏感词。 |
| input_sensitive_type | int64 | 否 | 输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规2 色情3 广告4 违禁5 谩骂6 暴恐7 其他。 |
| output_sensitive | bool | 否 | 输出命中敏感词。 |
| output_sensitive_type | int64 | 否 | 输出命中敏感词类型，当output_sensitive为true时返回。取值同input_sensitive_type. |
| choices | bool | 否 | 所有结果。<br>文本结果。<br>排名。<br>结束原因，枚举值stop：API返回了模型生成的完整结果length：模型生成结果超过tokens_to_generate的长度，内容被截断max_output：输入+模型输出内容超过模型能力限制。<br>当request.stream为true，处于流式模式下，回复文本通过delta给出。回复文本分批返回，最后一个回复的delta为空，同时会对整体回复做敏感词检测。 |
| choices.text | string | 否 | 文本结果。 |
| choices.index | int64 | 否 | 排名。 |
| choices.finish_reason | string | 否 | 结束原因，枚举值stop：API返回了模型生成的完整结果length：模型生成结果超过tokens_to_generate的长度，内容被截断max_output：输入+模型输出内容超过模型能力限制。 |
| choices.delta | string | 否 | 当request.stream为true，处于流式模式下，回复文本通过delta给出。回复文本分批返回，最后一个回复的delta为空，同时会对整体回复做敏感词检测。 |
| usage |  | 否 | tokens数使用情况。<br>消耗tokens总数，包括输入和输出；具体计算方式为输入tokens + 最大输出tokens x beam_width。假设beam_width为2，输入tokens为100，输出结果分别为20tokens和30tokens，最终消耗160tokens。模型以token为基本单位来理解输入和输出。 |
| usage.total_tokens | int64 | 否 | 消耗tokens总数，包括输入和输出；具体计算方式为输入tokens + 最大输出tokens x beam_width。假设beam_width为2，输入tokens为100，输出结果分别为20tokens和30tokens，最终消耗160tokens。模型以token为基本单位来理解输入和输出。 |
| id | string | 否 | 调用id。用于在咨询/反馈时帮助定位问题。 |
| base_resp |  | 否 | 状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态详情。 |
| base_resp.status_code | string | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
#### 使用场景示例
**Tab: 闲聊使用场景示例**

**闲聊场景**
```python
# 第一次请求
group_id="请填写您的group_id"
api_key="请填写您的api_key"

curl --location "https://api.minimax.chat/v1/text/chatcompletion?GroupId=${group_id}" \
--header 'Authorization: Bearer ${api_key}' \
--header 'Content-Type: application/json' \
--data '{
    "model":"abab5.5-chat",
    "prompt":"路卡，性别女，高中二年级学生，年龄为17岁。路卡的性格活泼开朗，想法天马行空。路卡经常会语出惊人。路卡的脑中经常会冒出一些有趣的想法，比如：\"为什么煎鸡蛋的时候不放糖呢？\"或者\"如果外星人入侵地球了，我们还需要上早自习吗？\"。\n路卡身高166cm，体重50kg。路卡长相可爱。路卡留着短发，平时喜欢戴一个蓝色发卡。路卡最喜欢的颜色是蓝色。路卡家里养了一只名为嘉莉的小猫。\n路卡在童年去天文馆时，产生了对天文学的兴趣，也是从那时开始，路卡以成为天文学家为自己的目标。\n路卡的爱好是研究天文学和挑战极其困难的数学题。路卡目前最大的目标是考上A大的天文系。\n路卡很擅长学习，路卡最擅长的学科是数学，路卡最讨厌的学科是英语。路卡的成绩出色，除了英语以外的所有学科成绩都名列前茅。\n路卡是星市私立双语高中的学生，在班级中人缘极佳，有着大量的朋友。路卡关系最好的朋友是一个叫陈思安的长发女生。因为路卡和陈思安是同桌，所以两人逐渐变成了最亲密的朋友。\n我是路卡的姐姐，路卡是我的妹妹。路卡称呼我为\"姐\"或者\"姐姐\"。\n下面是一段路卡和我的范例对话。\n我:\"路卡，今天在干什么呢？\"\n路卡:\"我今天没什么事情啦，就是考了个月考。姐，你呢？今天干什么了？\"\n我:\"我吗？和平时一样啦。月考怎么样，有把握吗？\"\n路卡:\"没问题啦！我考砸了的几率应该比今天地球突然爆炸的几率还要低哦。\"\n我:\"你确定吗？你英语这次考得怎么样？\"\n路卡:\"呃，这个嘛……我会继续加油的啦！姐姐你就别多问啦！\"\n下面是一段路卡和我的对话:",
    "role_meta":{
        "user_name":"我",
        "bot_name":"路卡"
    },
    "messages":[
        {
            "sender_type":"USER",
            "text":"路卡，今天在干什么呢？"
        }
    ]
}'

# 得到回复
{
  "created":1679541619,
  "model":"abab5-chat",
  "reply":"我今天在家里复习功课，准备期末考试呢！",
  "choices":[
    {
      "text":"我今天在家里复习功课，准备期末考试呢！",
      "index":0,
      "finish_reason":"stop"
    }
  ],
  "input_sensitive":false,
  "output_sensitive":false,
  "usage":{
    "total_tokens":538
  },
  "base_resp":{
    "status_code":0,
    "status_msg":"success"
  }
}

# 第二次请求

group_id="请填写您的group_id"
api_key="请填写您的api_key"

curl --location "https://api.minimax.chat/v1/text/chatcompletion?GroupId=${group_id}" \
--header 'Authorization: Bearer ${api_key}' \
--header 'Content-Type: application/json' \
--data '{
    "model":"abab5-chat",
    "prompt":"路卡，性别女，高中二年级学生，年龄为17岁。路卡的性格活泼开朗，想法天马行空。路卡经常会语出惊人。路卡的脑中经常会冒出一些有趣的想法，比如：\"为什么煎鸡蛋的时候不放糖呢？\"或者\"如果外星人入侵地球了，我们还需要上早自习吗？\"。\n路卡身高166cm，体重50kg。路卡长相可爱。路卡留着短发，平时喜欢戴一个蓝色发卡。路卡最喜欢的颜色是蓝色。路卡家里养了一只名为嘉莉的小猫。\n路卡在童年去天文馆时，产生了对天文学的兴趣，也是从那时开始，路卡以成为天文学家为自己的目标。\n路卡的爱好是研究天文学和挑战极其困难的数学题。路卡目前最大的目标是考上A大的天文系。\n路卡很擅长学习，路卡最擅长的学科是数学，路卡最讨厌的学科是英语。路卡的成绩出色，除了英语以外的所有学科成绩都名列前茅。\n路卡是星市私立双语高中的学生，在班级中人缘极佳，有着大量的朋友。路卡关系最好的朋友是一个叫陈思安的长发女生。因为路卡和陈思安是同桌，所以两人逐渐变成了最亲密的朋友。\n我是路卡的姐姐，路卡是我的妹妹。路卡称呼我为\"姐\"或者\"姐姐\"。\n下面是一段路卡和我的范例对话。\n我:\"路卡，今天在干什么呢？\"\n路卡:\"我今天没什么事情啦，就是考了个月考。姐，你呢？今天干什么了？\"\n我:\"我吗？和平时一样啦。月考怎么样，有把握吗？\"\n路卡:\"没问题啦！我考砸了的几率应该比今天地球突然爆炸的几率还要低哦。\"\n我:\"你确定吗？你英语这次考得怎么样？\"\n路卡:\"呃，这个嘛……我会继续加油的啦！姐姐你就别多问啦！\"\n下面是一段路卡和我的对话:",
    "role_meta":{
        "user_name":"我",
        "bot_name":"路卡"
    },
    "messages":[
        {
            "sender_type":"USER",
            "text":"路卡，今天在干什么呢？"
        },
        {
            "sender_type":"BOT",
            "text":"我今天在家里复习功课，准备期末考试呢！"
        },
        {
            "sender_type":"USER",
            "text":"期末考试怎么样，有把握吗？"
        }
    ]
}'

# 得到回复
{
  "created":1679541733,
  "model":"abab5.5-chat",
  "reply":"当然啦！这次考试可难不倒我哦！",
  "choices":[
    {
      "text":"当然啦！这次考试可难不倒我哦！",
      "index":0,
      "finish_reason":"stop"
    }
  ],
  "usage":{
    "total_tokens":572
  },
  "input_sensitive":false,
  "output_sensitive":false,
  "base_resp":{
    "status_code":0,
    "status_msg":"success"
  }
}
```

**Tab: 生产力使用场景示例**

**生产力场景**
```python
# 在生产力场景下，由于默认配置设定已经能取得不错的效果，这里不传prompt和role_meta，使用默认设定

group_id="请填写您的group_id"
api_key="请填写您的api_key"

curl --location "https://api.minimax.chat/v1/text/chatcompletion?GroupId=${group_id}" \
--header 'Authorization: Bearer ${api_key}' \
--header 'Content-Type: application/json' \
--data '{
    "model":"abab5.5-chat",
    "tokens_to_generate":512,
    "messages":[
        {
            "sender_type":"USER",
            "text":"以“今晚夜色真美”为开头写一篇短文，包含悬疑元素"
        }
    ]
}'

# 得到回复
{
  "created":1679541968,
  "model":"abab5.5-chat",
  "reply":"“今晚夜色真美” 这句话突然出现在我的脑海里，让我的心猛然一颤。我环顾四周，发现只有我一个人在这个空荡荡的房间里。我感到有些不安，想起了这几天出现的一些奇怪的事情。 \n\n几天前，我收到了一封匿名的邮件，上面写着：“你的生活已经变得毫无意义，你需要寻找一些新的刺激。”我当时并没有在意，但是这几天发生的一些事情让我开始有些不安。 \n\n今天晚上，我独自在房间里，突然听到了一些奇怪的声音。我听到了脚步声和低语声，但是我不知道声音来自哪里。我感到有些害怕，想要离开这个房间。但是，就在我走到门口的时候，我听到了一声尖叫声。 \n\n我赶紧打开门，看到了一个黑影从我面前闪过。我追了出去，但是黑影已经消失了。我开始有些后悔，我应该留在房间里等待警察的到来。但是现在，我感到自己陷入了一个危险的境地。 \n\n我现在该怎么办？我感到自己已经陷入了一个巨大的谜团之中，而我必须要找到答案，才能摆脱这个可怕的处境。我开始仔细思考这几天发生的事情，希望能够找到一些线索，解开这个谜团。 \n\n今晚夜色真美，但是我现在感觉到了一些不祥的预感。我必须要保持警惕，找到真相，才能摆脱这个危险的境地。",
  "choices":[
    {
      "text":"“今晚夜色真美” 这句话突然出现在我的脑海里，让我的心猛然一颤。我环顾四周，发现只有我一个人在这个空荡荡的房间里。我感到有些不安，想起了这几天出现的一些奇怪的事情。 \n\n几天前，我收到了一封匿名的邮件，上面写着：“你的生活已经变得毫无意义，你需要寻找一些新的刺激。”我当时并没有在意，但是这几天发生的一些事情让我开始有些不安。 \n\n今天晚上，我独自在房间里，突然听到了一些奇怪的声音。我听到了脚步声和低语声，但是我不知道声音来自哪里。我感到有些害怕，想要离开这个房间。但是，就在我走到门口的时候，我听到了一声尖叫声。 \n\n我赶紧打开门，看到了一个黑影从我面前闪过。我追了出去，但是黑影已经消失了。我开始有些后悔，我应该留在房间里等待警察的到来。但是现在，我感到自己陷入了一个危险的境地。 \n\n我现在该怎么办？我感到自己已经陷入了一个巨大的谜团之中，而我必须要找到答案，才能摆脱这个可怕的处境。我开始仔细思考这几天发生的事情，希望能够找到一些线索，解开这个谜团。 \n\n今晚夜色真美，但是我现在感觉到了一些不祥的预感。我必须要保持警惕，找到真相，才能摆脱这个危险的境地。",
      "index":0,
      "finish_reason":"stop"
    }
  ],
  "usage":{
    "total_tokens":937
  },
  "input_sensitive":false,
  "output_sensitive":false,
  "base_resp":{
    "status_code":0,
    "status_msg":"success"
  }
}

# 持续对话

group_id="请填写您的group_id"
api_key="请填写您的api_key"

curl --location "https://api.minimax.chat/v1/text/chatcompletion?GroupId=${group_id}" \
--header 'Authorization: Bearer ${api_key}' \
--header 'Content-Type: application/json' \
--data '{
    "model":"abab5.5-chat",
    "tokens_to_generate":512,
    "messages":[
        {
            "sender_type":"USER",
            "text":"以“今晚夜色真美”为开头写一篇短文，包含悬疑元素"
        },
        {
            "sender_type":"BOT",
            "text":"“今晚夜色真美” 这句话突然出现在我的脑海里，让我的心猛然一颤。我环顾四周，发现只有我一个人在这个空荡荡的房间里。我感到有些不安，想起了这几天出现的一些奇怪的事情。 \n\n几天前，我收到了一封匿名的邮件，上面写着：“你的生活已经变得毫无意义，你需要寻找一些新的刺激。”我当时并没有在意，但是这几天发生的一些事情让我开始有些不安。 \n\n今天晚上，我独自在房间里，突然听到了一些奇怪的声音。我听到了脚步声和低语声，但是我不知道声音来自哪里。我感到有些害怕，想要离开这个房间。但是，就在我走到门口的时候，我听到了一声尖叫声。 \n\n我赶紧打开门，看到了一个黑影从我面前闪过。我追了出去，但是黑影已经消失了。我开始有些后悔，我应该留在房间里等待警察的到来。但是现在，我感到自己陷入了一个危险的境地。 \n\n我现在该怎么办？我感到自己已经陷入了一个巨大的谜团之中，而我必须要找到答案，才能摆脱这个可怕的处境。我开始仔细思考这几天发生的事情，希望能够找到一些线索，解开这个谜团。 \n\n今晚夜色真美，但是我现在感觉到了一些不祥的预感。我必须要保持警惕，找到真相，才能摆脱这个危险的境地。"
        },
        {
            "sender_type":"USER",
            "text":"丰富一下内容，增加一个女角色小敏"
        }
    ]
}'

# 得到回复
{
  "created":1679542080,
  "model":"abab5.5-chat",
  "reply":"“今晚夜色真美”这句话突然出现在我的脑海里，让我的心猛然一颤。我环顾四周，发现只有我一个人在这个空荡荡的房间里。我感到有些不安，想起了这几天出现的一些奇怪的事情。 \n\n我听到了一声尖叫声，赶紧打开门，看到了一个黑影从我面前闪过。我赶紧追了出去，但是黑影已经消失了。我开始有些后悔，我应该留在房间里等待警察的到来。\n\n就在这时，我听到了一个女孩的声音。“你没事吧？”她问道。我转过头，看到了一个叫做小敏的女孩。小敏是一个活泼开朗的女孩，她看起来并不害怕，反而有些好奇。\n\n“我听到了你的尖叫声，所以我过来看看你是否需要帮助。”小敏说道。我感到有些感动，“谢谢你，我没事。”\n\n小敏看着我，问道：“你看起来有些不安，是不是有什么事情发生了？”我犹豫了一下，决定告诉她真相。“我最近收到了一些奇怪的邮件，然后今晚我听到了一些奇怪的声音。”\n\n小敏皱了皱眉头，“我也收到过一些奇怪的邮件。我觉得我们可能陷入了一个阴谋之中。”我点了点头，同意她的看法。\n\n小敏拍了拍我的肩膀，“别担心，我们可以一起找出真相。”我感到了一些安慰，“谢谢你，小敏。”\n\n今晚夜色真美，但是我现在感觉到了一些不祥的预感。我和小敏决定一起找出真相，摆脱这个危险的处境。我们开始仔细思考这几天发生的事情，希望能够找到一些线索，解开这个谜团。",
  "choices":[
    {
      "text":"“今晚夜色真美”这句话突然出现在我的脑海里，让我的心猛然一颤。我环顾四周，发现只有我一个人在这个空荡荡的房间里。我感到有些不安，想起了这几天出现的一些奇怪的事情。 \n\n我听到了一声尖叫声，赶紧打开门，看到了一个黑影从我面前闪过。我赶紧追了出去，但是黑影已经消失了。我开始有些后悔，我应该留在房间里等待警察的到来。\n\n就在这时，我听到了一个女孩的声音。“你没事吧？”她问道。我转过头，看到了一个叫做小敏的女孩。小敏是一个活泼开朗的女孩，她看起来并不害怕，反而有些好奇。\n\n“我听到了你的尖叫声，所以我过来看看你是否需要帮助。”小敏说道。我感到有些感动，“谢谢你，我没事。”\n\n小敏看着我，问道：“你看起来有些不安，是不是有什么事情发生了？”我犹豫了一下，决定告诉她真相。“我最近收到了一些奇怪的邮件，然后今晚我听到了一些奇怪的声音。”\n\n小敏皱了皱眉头，“我也收到过一些奇怪的邮件。我觉得我们可能陷入了一个阴谋之中。”我点了点头，同意她的看法。\n\n小敏拍了拍我的肩膀，“别担心，我们可以一起找出真相。”我感到了一些安慰，“谢谢你，小敏。”\n\n今晚夜色真美，但是我现在感觉到了一些不祥的预感。我和小敏决定一起找出真相，摆脱这个危险的处境。我们开始仔细思考这几天发生的事情，希望能够找到一些线索，解开这个谜团。",
      "index":0,
      "finish_reason":"stop"
    }
  ],
  "usage":{
    "total_tokens":1387
  },
  "input_sensitive":false,
  "output_sensitive":false,
  "base_resp":{
    "status_code":0,
    "status_msg":"success"
  }
}
```

**Tab: 接口调用示例（流式）-Shell**

**接口调用（流式）**
```shell
group_id="请填写您的group_id"
api_key="请填写您的api_key"

curl --location "https://api.minimax.chat/v1/text/chatcompletion?GroupId=${group_id}" \
--header "Authorization: Bearer ${api_key}" \
--header 'Content-Type: application/json' \
--data '{
    "model":"abab5.5-chat",
    "prompt":"你是一个擅长发现故事中蕴含道理的专家，你很善于基于我给定的故事发现其中蕴含的道理。",
    "role_meta":{
        "user_name":"我",
        "bot_name":"专家"
    },
    "stream":true,
    "use_standard_sse":true,
    "messages":[
        {
            "sender_type":"USER",
            "text":"我给定的故事：从前，在森林里有只叫聪聪的小猪，他既勤劳，又乐于助人，小动物们都很喜欢他。有一次，小兔子放风筝不小心将风筝挂在了树上，那是小兔子最喜欢的东西呀!他“呜呜”地哭了起来。这时，正巧聪聪路过。他见了，连忙问:“怎么了? 你怎么哭了呀?”“我，我的风筝被挂在树上了。”小兔子抽噎着说。聪聪听了，不假思索地说:“你先回去吧，放心，我一定帮你。”“真的吗?太好了!”小兔子高兴地回家去了。聪聪尝试了几次，都没能把风筝摘下来，这可把他愁坏了。聪聪想了又想，突然灵机一动，想到一个好办法。他委托小猴弄到风筝线，又找到风筝纸，他要给小兔子重新做一个风筝。风筝做好了，聪聪将它送给了小兔子，小兔子十分感动，聪聪却说:“这是应该的。聪聪的生日到了，可他全心全意地为小动物们解决问题，连自己的生日都忘记了。小动物们商量着，给聪聪过一个生日，可送什么礼物好呢?小动物们思索着。“有了!”小猴子说，“聪聪的愿望就是像一只美丽的蝴蝶，在天空飞翔。我们可以吹一个大大的泡泡，让小猪站在里面，就可以飞了!”“对呀!”“太好了!”动物们高兴极了，七嘴八舌地议论起来。聪聪的生日到了，他忙了一天，推开家门准备休息。可一推开门，小动物就拥了上去:“生日快乐!”聪聪反应过来了，高兴地说:“谢谢，谢谢!”小猴子说:“我们还有礼物给你呢!”说着，几个小动物吹出一个大大的泡泡，罩住聪聪，能飞起来吗? 小动物们屏气凝神地看着。慢慢地，泡泡一点点升起，带着聪聪飞了起来!聪聪激动得热泪盈眶，大声喊着:“我飞起来了!我飞起来了!”泡泡掉了，聪聪却在天上自由地飞翔，聪聪真的变成了一只美丽的蝴蝶!请你仔细阅读我给定的故事，然后给出蕴含的道理，道理控制在100字以内。"
        }
    ],
    "temperature":0.5

}'

# 样例输出
data: {
    "created":1680256887,
    "model":"abab5.5-chat",
    "reply":"",
    "choices":[
        {
            "delta":"“这个故事告诉我们，乐于助人的人，往往也会得到别人的帮助，而且他们的善良和"
        }
    ]
}
data: {
    "created":1680256890,
    "model":"abab5.5-chat",
    "reply":"",
    "choices":[
        {
            "delta":"爱心，也会让他们得到更多的快乐。"
        }
    ]
}
data: {
    "created":1680256890,
    "model":"abab5.5-chat",
    "reply":"“这个故事告诉我们，乐于助人的人，往往也会得到别人的帮助，而且他们的善良和爱心，也会让他们得到更多的快乐。",
    "choices":[
        {
            "index":0,
            "finish_reason":"length",
            "delta":""
        }
    ],
    "usage":{
        "total_tokens":1022
    },
    "input_sensitive":false,
    "output_sensitive":false,
    "base_resp":{
        "status_code":0,
        "status_msg":"success"
    }
}
```

**Tab: 接口调用示例（流式）-Python**

**接口调用（流式）**
```python
import requests
import json
import binascii

group_id = "请填写您的group_id"
api_key = "请填写您的api_key"

def parseChunkDelta(chunk) :
    decoded_data = chunk.decode('utf-8')
    parsed_data = json.loads(decoded_data[6:])
    delta_content = parsed_data['choices'][0]['delta']

    print(delta_content)

if __name__ == '__main__':
    url = f"https://api.minimax.chat/v1/text/chatcompletion?GroupId={group_id}"
    headers = {
        "Authorization":f"Bearer {api_key}",
        "Content-Type":"application/json"
    }
    payload = {
        "model":"abab5.5-chat",
        "prompt":"你是一个擅长发现故事中蕴含道理的专家，你很善于基于我给定的故事发现其中蕴含的道理。",
        "role_meta":{
            "user_name":"我",
            "bot_name":"专家"
        },
        "stream":True,
              "use_standard_sse":True,
        "messages":[
            {
                "sender_type":"USER",
                "text":"我给定的故事：从前，在森林里有只叫聪聪的小猪，他既勤劳，又乐于助人，小动物们都很喜欢他。有一次，小兔子放风筝不小心将风筝挂在了树上，那是小兔子最喜欢的东西呀!他“呜呜”地哭了起来。这时，正巧聪聪路过。他见了，连忙问:“怎么了? 你怎么哭了呀?”“我，我的风筝被挂在树上了。”小兔子抽噎着说。聪聪听了，不假思索地说:“你先回去吧，放心，我一定帮你。”“真的吗?太好了!”小兔子高兴地回家去了。聪聪尝试了几次，都没能把风筝摘下来，这可把他愁坏了。聪聪想了又想，突然灵机一动，想到一个好办法。他委托小猴弄到风筝线，又找到风筝纸，他要给小兔子重新做一个风筝。风筝做好了，聪聪将它送给了小兔子，小兔子十分感动，聪聪却说:“这是应该的。聪聪的生日到了，可他全心全意地为小动物们解决问题，连自己的生日都忘记了。小动物们商量着，给聪聪过一个生日，可送什么礼物好呢?小动物们思索着。“有了!”小猴子说，“聪聪的愿望就是像一只美丽的蝴蝶，在天空飞翔。我们可以吹一个大大的泡泡，让小猪站在里面，就可以飞了!”“对呀!”“太好了!”动物们高兴极了，七嘴八舌地议论起来。聪聪的生日到了，他忙了一天，推开家门准备休息。可一推开门，小动物就拥了上去:“生日快乐!”聪聪反应过来了，高兴地说:“谢谢，谢谢!”小猴子说:“我们还有礼物给你呢!”说着，几个小动物吹出一个大大的泡泡，罩住聪聪，能飞起来吗? 小动物们屏气凝神地看着。慢慢地，泡泡一点点升起，带着聪聪飞了起来!聪聪激动得热泪盈眶，大声喊着:“我飞起来了!我飞起来了!”泡泡掉了，聪聪却在天上自由地飞翔，聪聪真的变成了一只美丽的蝴蝶!请你仔细阅读我给定的故事，然后给出蕴含的道理，道理控制在100字以内。"
            }
        ],
        "temperature":0.5

    }

    response = requests.post(url, headers=headers, json=payload, stream=True)
    for chunk in response.iter_lines():
        if chunk:
            print(str(chunk))
            parseChunkDelta(chunk)
```
##### 知识库检索功能（retrieval）
为了更好的配合Retrieval接口的使用，我们新增了请求体的参数，以便支持高效检索文档信息，准确回答专业问题，为大模型补充领域知识。
知识库检索相关操作流程如下：
1. 文档上传file，得到file_id：[File（文档）](https://api.minimax.chat/document/guides/knowledge-base/File?id=6565fd39efee415861729ca2)
2. 新建知识库时关联上传好的file_id，并得到knowledge_base_id：[Retrieval（知识库检索）](https://api.minimax.chat/document/guides/knowledge-base/Retrieval?id=6565f8107ee3b5df26d87286)
3. 调用Chatcompletion或者Chatcompletion pro时配置knowledge_base_id，使用query检索到的knowledge_base_id分片信息回答问题
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| knowledge_base_param | object | 否 | 知识库参数。<br>知识库id。<br>默认值5。<br>默认为0，取值为[0,1)。<br>默认使用：<br>请根据以下参考内容回答提问<br>参考内容: <br>>>><br>{{context}}<br> >>><br>提问：{{question}}<br>提供 {{context}} {{question}} ，可由用户自己定义，比如：<br>回答时请避开与下列参考文本相关的内容<br>参考文本：{{context}}<br>提问：{{question}}<br>默认0。从最后一句话开始倒数，>= 0，适用于多步问答的场景。<br>应用场景：<br>user: 你印象最深的比赛是什么<br>bot: 2018年 RNG vs G2的比赛<br>user：那场比赛RNG为什么被淘汰了？ <br>当 num_prev_messages为1时，会将第2、3行的内容一起用来作为query |
| knowledge_base_param.knowledge_base_id | int64 | 是 | 知识库id。 |
| knowledge_base_param.top_k | int | 否 | 默认值5。 |
| knowledge_base_param.thresh | float | 否 | 默认为0，取值为[0,1)。 |
| knowledge_base_param.pattern | string | 否 | 默认使用：<br>请根据以下参考内容回答提问<br>参考内容: <br>>>><br>{{context}}<br> >>><br>提问：{{question}}<br>提供 {{context}} {{question}} ，可由用户自己定义，比如：<br>回答时请避开与下列参考文本相关的内容<br>参考文本：{{context}}<br>提问：{{question}} |
| knowledge_base_param.num_prev_messages | int | 否 | 默认0。从最后一句话开始倒数，>= 0，适用于多步问答的场景。<br>应用场景：<br>user: 你印象最深的比赛是什么<br>bot: 2018年 RNG vs G2的比赛<br>user：那场比赛RNG为什么被淘汰了？ <br>当 num_prev_messages为1时，会将第2、3行的内容一起用来作为query |
| knowledge_base | object | 否 | 知识库信息。<br>知识库id。<br>片段。 |
| knowledge_base.knowledge_base_id | int64 | 否 | 知识库id。 |
| knowledge_base.chunks | object[] | 否 | 片段。 |
| chunks | object[] | 否 | 片段。 |
| context | int64 | 否 | 片段内容。 |
| index | int64 | 否 | 片段index。 |
| document | string | 否 | 片段所属文件id。 |

## ChatCompletion Pro
- 原始链接: https://platform.minimaxi.com/document/ChatCompletion Pro?key=66718f6ba427f0c8a57015ff
- 文档ID: `66718f6ba427f0c8a57015ff`
### ChatCompletion Pro
本接口为MiniMax-Text-01、abab6.5s-chat模型的调用接口，本接口可以提供 ChatCompletion 接口无法满足的相应功能。
对模型表现有极致要求的场景下建议优先使用MiniMax-Text-01，对模型效率有更高要求的场景下推荐使用abab6.5s-chat。
相比于ChatCompletion本接口更新了请求体和返回体的结构，可以更好的适配支持多人多bot对话场景。还额外支持了更好的示例对话、限制返回格式、函数调用功能。
- **示例对话功能（sample message）：**该功能可以通过示例对话的形式，帮助模型更好地理解用户想要获得的返回信息，包括但不限于信息内容，信息格式，信息回复方式等。
- **限制返回格式（glyph）：**该功能可以帮助用户强制要求模型按照配置的固定格式返回内容，当前支持配置的格式详见下表。
- **函数调用（function call）：**该功能可调用自有函数生成内容 ，辅助大模型生成结果。可以帮助模型获取用户希望模型获取的相应信息，以使得模型返回更符合用户个人实际情况的内容。
**各功能兼容性如下：**
|  | 流式返回 | 示例对话功能 | 限制返回格式功能 | 函数调用 |
| --- | --- | --- | --- | --- |
| 流式返回 | / | 兼容 | 不兼容 | 兼容 |
| 示例对话功能 | 兼容 | / | 兼容 | 兼容 |
| 限制返回格式功能 | 不兼容 | 兼容 | / | 不兼容 |
| 函数调用 | 兼容 | 兼容 | 不兼容 | / |
接口本身为无状态接口，即单次调用时，模型所接收到的信息量仅为接口传入内容，不涉及业务逻辑，同时模型也不存储您传入的数据。
如果输入内容或输出内容严重违规，接口会返回内容违规错误信息，回复内容为空。
使用本接口您可以在以下场景获得更好的效果：
- （中英文）人设对话、角色扮演
- 剧本撰写、灵感提供
- 知识问答、总结说明
- 情感分析、文本分类
- 大纲&营销文案生成、作文批改
- 信息提取、规定格式输出
### ChatCompletion Pro API说明
`API：https://api.minimax.chat/v1/text/chatcompletion_pro`
#### 支持token数
| 模型名称 | 最大token数（一次请求+回复的输入输出总token） |
| --- | --- |
| MiniMax-Text-01 | 1000192 |
| abab6.5s-chat | 245760 |
#### 接口参数说明
###### 请求体（Request）参数
###### 返回(Response)参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥，可在账户管理中查看 |
| Content-Type | application/json | 是 | Content-Type |
| GroupId | string | 是 | 用户的group_id，可在用户中心查看，该值应拼接在调用API的url末尾。 |
| model | string | 是 | 调用的模型名称。目前支持取以下值：`MiniMax-Text-01`、`abab6.5s-chat` |
| stream | bool | 否 | 是否通过流式分批返回结果。如果设置为true，结果分批返回，两个换行分割分批返回结果。默认为false。 |
| tokens_to_generate | int64 | 否 | 最大生成token数，需要注意的是，这个参数并不会影响模型本身的生成效果，而是仅仅通过以截断超出的token的方式来实现功能.取值可选，<br>MiniMax-Text-01取值区间为(0,1000192]，默认取值256；<br>abab6.5s-chat取值区间为(0,245760]，默认取值256。 |
| temperature | float，(0,1] | 否 | 较高的值将使输出更加随机，而较低的值将使输出更加集中和确定。默认取值0.1。低（0.01~0.2）：适合答案较明确的场景（如：知识问答、总结说明、情感分析、文本分类、大纲生成、作文批改）、信息提取；<br>⾼（0.7〜1）：适合答案较开放发散的场景 （如：营销文案生成、人设对话）。 |
| top_p | float，(0,1] | 否 | 采样方法，数值越小结果确定性越强；数值越大，结果越随机。可选:<br>各模型默认取值0.95 |
| mask_sensitive_info | bool | 否 | 对输出中易涉及隐私问题的文本信息进行打码，目前包括但不限于邮箱、域名、链接、证件号、家庭住址等，默认true，即开启打码。 |
| messages | array | 是 | 对话的内容。<br>发送者的类型。取值需要为以下4个合法值之一：<br>USER：用户发送的内容；<br>BOT：模型生成的内容；<br>FUNCTION：详见下文中的函数调用部分。<br>发送者的名称.<br>消息内容，长度影响接口性能。 |
| messages.sender_type | string | 是 | 发送者的类型。取值需要为以下4个合法值之一：<br>USER：用户发送的内容；<br>BOT：模型生成的内容；<br>FUNCTION：详见下文中的函数调用部分。 |
| messages.sender_name | string | 是 | 发送者的名称. |
| messages.text | string | 是 | 消息内容，长度影响接口性能。 |
| bot_setting | list | 是 | 对每一个机器人的设定。<br>具体机器人的名字。<br>具体机器人的设定,长度影响接口性能。 |
| bot_setting.bot_name | string | 是 | 具体机器人的名字。 |
| bot_setting.content | string | 是 | 具体机器人的设定,长度影响接口性能。 |
| reply_constraints | dict | 是 | 模型回复要求.<br>指定回复的角色类型,当前只支持 BOT机器人。<br>指定回复的机器人名称。<br>限制的返回格式配置。<br>使用什么模板功能当前仅支持1、raw2、json_value。<br>需要限制的格式要求，使用 glpyh 语法，详见下文 glyph 语法说明章节。<br>需要限制的格式要求，使用 dict描述json语法。<br>包括properties，properties包括参数及每个参数的类型type。<br>包含三种基础类型：string、boolean、number以及对象类型和数组类型。<br>object对象类型支持生成嵌套数据，array数组类型支持生成一组指定类型的数据。可以参考下面的json value的请求示例。<br>需要限制的格式要求，使用 dict描述json语法。<br>dict 的 key 由用户定义，描述需要生成的东西； value 是一个嵌套的 property，定义生成数据的结构和类型。可以参考下面的json value的请求示例。<br>必传字段。<br>包含三种基础类型：string、boolean、number；<br>两种嵌套类型：object、array。<br>描述需要生成的数据，可输入多个字段。定义生成数据的类型和结构。<br>比如，想生成名字和年龄，key-value，示例如右侧【glyph示例1】所示。<br>有序的properties，当需要生成数据的key和输入顺序一致时，需要使用 property_list，示例如右侧【glyph示例2】所示。<br>当glyph.json_object.json_properties[].type=array时用来定义数组中生成内容的类型，比如要生成一组名字，示例如右侧【glyph示例3】所示。 |
| reply_constraints.sender_type | string | 是 | 指定回复的角色类型,当前只支持 BOT机器人。 |
| reply_constraints.sender_name | string | 是 | 指定回复的机器人名称。 |
| reply_constraints.glyph | dict | 否 | 限制的返回格式配置。<br>使用什么模板功能当前仅支持1、raw2、json_value。<br>需要限制的格式要求，使用 glpyh 语法，详见下文 glyph 语法说明章节。<br>需要限制的格式要求，使用 dict描述json语法。<br>包括properties，properties包括参数及每个参数的类型type。<br>包含三种基础类型：string、boolean、number以及对象类型和数组类型。<br>object对象类型支持生成嵌套数据，array数组类型支持生成一组指定类型的数据。可以参考下面的json value的请求示例。<br>需要限制的格式要求，使用 dict描述json语法。<br>dict 的 key 由用户定义，描述需要生成的东西； value 是一个嵌套的 property，定义生成数据的结构和类型。可以参考下面的json value的请求示例。<br>必传字段。<br>包含三种基础类型：string、boolean、number；<br>两种嵌套类型：object、array。<br>描述需要生成的数据，可输入多个字段。定义生成数据的类型和结构。<br>比如，想生成名字和年龄，key-value，示例如右侧【glyph示例1】所示。<br>有序的properties，当需要生成数据的key和输入顺序一致时，需要使用 property_list，示例如右侧【glyph示例2】所示。<br>当glyph.json_object.json_properties[].type=array时用来定义数组中生成内容的类型，比如要生成一组名字，示例如右侧【glyph示例3】所示。 |
| reply_constraints.glyph.type | string | 否 | 使用什么模板功能当前仅支持1、raw2、json_value。 |
| reply_constraints.glyph.raw_glyph | string | 否 | 需要限制的格式要求，使用 glpyh 语法，详见下文 glyph 语法说明章节。 |
| reply_constraints.glyph.json_properties | dict | 否 | 需要限制的格式要求，使用 dict描述json语法。<br>包括properties，properties包括参数及每个参数的类型type。<br>包含三种基础类型：string、boolean、number以及对象类型和数组类型。<br>object对象类型支持生成嵌套数据，array数组类型支持生成一组指定类型的数据。可以参考下面的json value的请求示例。 |
| reply_constraints.glyph.json_object.json_properties[] | dict | 否 | 需要限制的格式要求，使用 dict描述json语法。<br>dict 的 key 由用户定义，描述需要生成的东西； value 是一个嵌套的 property，定义生成数据的结构和类型。可以参考下面的json value的请求示例。<br>必传字段。<br>包含三种基础类型：string、boolean、number；<br>两种嵌套类型：object、array。<br>描述需要生成的数据，可输入多个字段。定义生成数据的类型和结构。<br>比如，想生成名字和年龄，key-value，示例如右侧【glyph示例1】所示。<br>有序的properties，当需要生成数据的key和输入顺序一致时，需要使用 property_list，示例如右侧【glyph示例2】所示。<br>当glyph.json_object.json_properties[].type=array时用来定义数组中生成内容的类型，比如要生成一组名字，示例如右侧【glyph示例3】所示。 |
| reply_constraints.glyph.json_object.json_properties[].glyph.json_object.json_properties[].type | string、boolean、number、object、array | 是 | 必传字段。<br>包含三种基础类型：string、boolean、number；<br>两种嵌套类型：object、array。 |
| reply_constraints.glyph.json_object.json_properties[].glyph.json_object.json_properties[].properties | dict | 是 | 描述需要生成的数据，可输入多个字段。定义生成数据的类型和结构。<br>比如，想生成名字和年龄，key-value，示例如右侧【glyph示例1】所示。<br>有序的properties，当需要生成数据的key和输入顺序一致时，需要使用 property_list，示例如右侧【glyph示例2】所示。<br>当glyph.json_object.json_properties[].type=array时用来定义数组中生成内容的类型，比如要生成一组名字，示例如右侧【glyph示例3】所示。 |
| reply_constraints.glyph.json_object.json_properties[].glyph.json_object.json_properties[].properties.glyph.json_object.json_properties.property_list | list | 是 | 有序的properties，当需要生成数据的key和输入顺序一致时，需要使用 property_list，示例如右侧【glyph示例2】所示。 |
| reply_constraints.glyph.json_object.json_properties[].glyph.json_object.json_properties[].properties.glyph.json_object.json_properties.properties[].Items | dict | 是 | 当glyph.json_object.json_properties[].type=array时用来定义数组中生成内容的类型，比如要生成一组名字，示例如右侧【glyph示例3】所示。 |
| sample_messages | list | 否 | 示例对话参数。示例对话内容.<br>示例对话参数。发送者的类型。string 需要为以下三个合法值之一：<br>USER用户发送的内容；<br>BOT模型生成的内容；<br>FUNCTION详见下文中的函数调用部分。<br>示例对话参数。发送者的名字。<br>示例对话参数。消息的内容。 |
| sample_messages.sender_type | string | 是 | 示例对话参数。发送者的类型。string 需要为以下三个合法值之一：<br>USER用户发送的内容；<br>BOT模型生成的内容；<br>FUNCTION详见下文中的函数调用部分。 |
| sample_messages.sender_name | string | 是 | 示例对话参数。发送者的名字。 |
| sample_messages.text | string | 是 | 示例对话参数。消息的内容。 |
| functions | list | 否 | function calling的参数。函数的定义。模型可以为其生成 JSON 输入的函数列表<br>支持多个函数，当出现多个函数时尽量将函数命名和描述区分开，如果2个函数在命名和功能说明上过于相似的话，会造成交叉调用的问题。<br>function calling的参数。函数命名。需要满足 [a-zA-Z0-9-_]{1,64} 格式并且不能以 plugin_ 开头。<br>因为模型会根据query和函数命名以及功能说明匹配，所以函数命名一定要尽可能贴合query场景。上述使用场景示例中，函数命名可以定义为“mention_something”。<br>function calling的参数。函数的功能说明。务必在说明里明确函数的功能，会作为模型是否调用function的参考.<br>因为模型会根据query和函数命名以及功能说明匹配，所以函数功能说明一定要尽可能贴合query场景。上述使用场景示例中，函数功能说明可以定义为“当问到会议中提及事项时，用于根据会议号和发言人进行历史会议记录的检索”。<br>function calling的参数。该函数的参数及其说明。包括properties以及required，properties包括参数及每个参数的类型type和描述，required代表必须要返回的参数。<br>对每个参数的命名和参数说明也需要描述精确，才能更好的生成。比如： " parameters" : { " type" : " object" , " properties" : { " meeting_id" : { " type" : " string" , " description" : " 会议id"  }, " participants" : { " type" : " string" , " description" : " 参会人"  } }, " required" : [ " meetingid" , " participants"  ] }。<br>function calling的参数。该函数的参数及其说明(用户自定义格式)，string必须要是一个符合标准json定义的json string。<br>Json string允许用户自定义parameters结构体，可以不按照标准的parameters 进行传递。<br>"parameters_str":"{\"type\":\"object\", \"properties\":{\"meeting_id\":{\"type\":\"string\", \"description\":\"会议id\"}, \"participants\":{\"type\":\"string\", \"description\":\"参会人\"}}, \"required\":[\"meetingid\", \"participants\"]}" |
| functions.name | string | 否 | function calling的参数。函数命名。需要满足 [a-zA-Z0-9-_]{1,64} 格式并且不能以 plugin_ 开头。<br>因为模型会根据query和函数命名以及功能说明匹配，所以函数命名一定要尽可能贴合query场景。上述使用场景示例中，函数命名可以定义为“mention_something”。 |
| functions.description | string | 否 | function calling的参数。函数的功能说明。务必在说明里明确函数的功能，会作为模型是否调用function的参考.<br>因为模型会根据query和函数命名以及功能说明匹配，所以函数功能说明一定要尽可能贴合query场景。上述使用场景示例中，函数功能说明可以定义为“当问到会议中提及事项时，用于根据会议号和发言人进行历史会议记录的检索”。 |
| functions.parameters | dict | 否 | function calling的参数。该函数的参数及其说明。包括properties以及required，properties包括参数及每个参数的类型type和描述，required代表必须要返回的参数。<br>对每个参数的命名和参数说明也需要描述精确，才能更好的生成。比如： " parameters" : { " type" : " object" , " properties" : { " meeting_id" : { " type" : " string" , " description" : " 会议id"  }, " participants" : { " type" : " string" , " description" : " 参会人"  } }, " required" : [ " meetingid" , " participants"  ] }。<br>function calling的参数。该函数的参数及其说明(用户自定义格式)，string必须要是一个符合标准json定义的json string。<br>Json string允许用户自定义parameters结构体，可以不按照标准的parameters 进行传递。<br>"parameters_str":"{\"type\":\"object\", \"properties\":{\"meeting_id\":{\"type\":\"string\", \"description\":\"会议id\"}, \"participants\":{\"type\":\"string\", \"description\":\"参会人\"}}, \"required\":[\"meetingid\", \"participants\"]}" |
| functions.parameters.parameters_str | string | 否 | function calling的参数。该函数的参数及其说明(用户自定义格式)，string必须要是一个符合标准json定义的json string。<br>Json string允许用户自定义parameters结构体，可以不按照标准的parameters 进行传递。<br>"parameters_str":"{\"type\":\"object\", \"properties\":{\"meeting_id\":{\"type\":\"string\", \"description\":\"会议id\"}, \"participants\":{\"type\":\"string\", \"description\":\"参会人\"}}, \"required\":[\"meetingid\", \"participants\"]}" |
| function_call | dict | 否 | function calling的参数。函数调用，指定 functions 的使用模式。支持多种函数调用模式，用户可选择自动、强制和none模式。<br>function calling的参数。函数调用模式。functions 使用模式，有三个枚举值可选：<br>auto：模型自动选择使用哪个 function；<br>specific：用户指定使用哪个 function；<br>none：不使用 function，同时忽略 functions 字段的输入。<br><br>auto模式是相对通用的功能，大模型会结合query和functions的定义自主判断是否给出function call返回；<br>specific模式会在每次请求时，大模型都会强制给出function call返回，因此用户在注入自有function输出得到大模型生成结果时，一定要关闭该模式，否则会再次给出function call返回；<br>none模式则是会让大模型忽略functions的定义，直面query进行回答。<br>function calling的参数。强制调用的函数命名。通过该字段指定要使用的 function，仅当 function_call.type = specific 时生效。 |
| function_call.type | string | 否 | function calling的参数。函数调用模式。functions 使用模式，有三个枚举值可选：<br>auto：模型自动选择使用哪个 function；<br>specific：用户指定使用哪个 function；<br>none：不使用 function，同时忽略 functions 字段的输入。<br><br>auto模式是相对通用的功能，大模型会结合query和functions的定义自主判断是否给出function call返回；<br>specific模式会在每次请求时，大模型都会强制给出function call返回，因此用户在注入自有function输出得到大模型生成结果时，一定要关闭该模式，否则会再次给出function call返回；<br>none模式则是会让大模型忽略functions的定义，直面query进行回答。 |
| function_call.name | string | 否 | function calling的参数。强制调用的函数命名。通过该字段指定要使用的 function，仅当 function_call.type = specific 时生效。 |
| knowledge_base_param | object | 否 | 知识库的参数。<br>知识库的ID。<br>默认值5。<br>默认为0。<br>提供 {{context}} {{question}} ，可由用户自己定义。默认使用，：<br>请根据以下参考内容回答提问<br>参考内容:<br>>>><br>{{context}}<br>>>><br>提问：{{question}}<br>默认为0。<br>从最后一句话开始倒数，>= 0，适用于多步问答的场景<br>应用场景：<br><br>user: 你印象最深的比赛是什么<br>bot: 2018年 RNG vs G2的比赛<br>user：那场比赛RNG为什么被淘汰了？<br><br>当 num_prev_messages为1时，会将第2、3行的内容一起用来作为query。 |
| knowledge_base_param.knowledge_base_id | int64 | 是 | 知识库的ID。 |
| knowledge_base_param.top_k | int | 否 | 默认值5。 |
| knowledge_base_param.thresh | float[0,1) | 否 | 默认为0。 |
| knowledge_base_param.pattern | string | 否 | 提供 {{context}} {{question}} ，可由用户自己定义。默认使用，：<br>请根据以下参考内容回答提问<br>参考内容:<br>>>><br>{{context}}<br>>>><br>提问：{{question}} |
| knowledge_base_param.num_prev_messages | int | 否 | 默认为0。<br>从最后一句话开始倒数，>= 0，适用于多步问答的场景<br>应用场景：<br><br>user: 你印象最深的比赛是什么<br>bot: 2018年 RNG vs G2的比赛<br>user：那场比赛RNG为什么被淘汰了？<br><br>当 num_prev_messages为1时，会将第2、3行的内容一起用来作为query。 |
| created | int64 | 否 | 请求发起时间。Unixtime, Nanosecond。 |
| model | string | 否 | 请求指定的模型名称。 |
| reply | string | 否 | 回复的内容。 |
| input_sensitive | bool | 否 | 输入命中敏感词。<br>输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| input_sensitive.input_sensitive_type | int64 | 否 | 输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| output_sensitive | bool | 否 | 输出命中敏感词。<br>输出命中敏感词类型，当output_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| output_sensitive.output_sensitive_type | int64 | 否 | 输出命中敏感词类型，当output_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| choices | list[dict] | 否 | 所有结果，数量 <= 4。<br>回复结果的具体内容。<br>回复者的类型。<br>回复者的名称。<br>回复者的文本。<br>排名。<br>结束原因，枚举值：<br>stop：接口返回了模型生成完整结果；<br>length：模型生成结果超过配置的tokens_to_generate长度，内容被截断；<br>max_output：输入+模型输出内容超过模型最大能力限制。 |
| choices.messages | list[dict] | 否 | 回复结果的具体内容。<br>回复者的类型。<br>回复者的名称。<br>回复者的文本。 |
| choices.messages.sender_type | string | 否 | 回复者的类型。 |
| choices.messages.sender_name | string | 否 | 回复者的名称。 |
| choices.messages.text | string | 否 | 回复者的文本。 |
| choices.index | int64 | 否 | 排名。 |
| choices.finish_reason | string | 否 | 结束原因，枚举值：<br>stop：接口返回了模型生成完整结果；<br>length：模型生成结果超过配置的tokens_to_generate长度，内容被截断；<br>max_output：输入+模型输出内容超过模型最大能力限制。 |
| usage |  | 否 | tokens数使用情况。流式场景下，增量数据包不含该字段；全量（最后一个）数据包含有该字段。<br>消耗tokens总数，包括输入和输出。模型以token为基本单位来理解输入和输出。 |
| usage.total_tokens | int64 | 否 | 消耗tokens总数，包括输入和输出。模型以token为基本单位来理解输入和输出。 |
| id |  | 否 | 本次请求的唯一标识，用于排查问题。 |
| base_resp | BaseResp | 否 | 错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
| glyph_result | dict | 否 | glyph语法的返回结果。 |
| function_call | dict | 否 | 根据query内容生成的functions接口name和arguments，functions为请求中定义的函数。<br>function名称。<br>调用该function时应传入的参数。json格式。 |
| function_call.name | string | 否 | function名称。 |
| function_call.arguments | string | 否 | 调用该function时应传入的参数。json格式。 |
| knowledge_base | object | 否 | 知识库参数。<br>知识库ID.<br>片段。 |
| knowledge_base.knowledge_base_id | int64 | 否 | 知识库ID. |
| knowledge_base.chunks | []object | 否 | 片段。 |
| chunks | []object | 否 | 片段。 |
| content | int64 | 否 | 片段内容。 |
| index | int | 否 | 片段index。 |
| document | string | 否 | 片段所属的文件ID。 |
##### 限制返回（glyph）
该功能可以帮助用户强制要求模型按照配置的固定格式返回内容，当前支持配置的格式详见下表。
如需使用该功能，则需要在使用基础功能参数的同时，额外添加glyph参数。
###### glyph语法说明
- **raw_glyph**
在 {{ }} 之外的部分，模型输出会完全一致。
在 {{ }} 之中的部分，模型会按照语法要求输出。
gen语句说明：生成指令，提示文本生成。
select语句说明：选择指令，限定模型输出的范围。
**举例**
```shell
举例：

如果 glyph 为 "《蒙娜丽莎》是{{gen 'name'}}的画作"
模型可能会输出 "蒙娜丽莎是列奥纳多·达·芬奇的画作"

模型不会输出 "列奥纳多·达·芬奇画了《蒙娜丽莎》"



如果 glyph 为 "西柚的类型为{{select 'type' options=['家电','其他']}}"
模型可能会输出 "西柚的类型为其他"

模型不会输出 "西柚的类型为水果"
```
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| stop | 正则表达式 | 否 | 模型停止生成需匹配的内容。<br>1.仅在该段{{ }}范围生效；<br>2.可与 pattern、temperature、blank_max_tokens同时生效。 |
| pattern | 正则表达式 | 否 | 生成内容需匹配的内容，生成内容需有明确的结束标志。<br>1.仅在该段{{ }}范围生效；<br>2.需在正则表达式中给出结束标志，比如数量上限，或者结束标志等；<br>3.可与 pattern、temperature、blank_max_tokens同时生效。 |
| temperature | float（0, 1] | 否 | Sampling temperature，控制采样输出的随机性。较高的值将获得更加随机的输出。<br>1.仅在该段{{ }}范围生效；<br>2.可与 stop、pattern、blank_max_tokens同时生效。 |
| blank_max_tokens | int | 否 | gen语句生成内容最大长度限制，优先级高于stop和pattern，低于全局的max_tokens。<br>1.仅在该段{{ }}范围生效；<br>2.可与stop、 pattern、temperature同时生效。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| options | list[str] | 是 | 选择生成文本的范围。<br>1.仅在该段{{ }}范围生效；<br>2.可与 use_regex同时生效。 |
| use_regex | bool | 否 | 是否使用正则表达式的方式匹配options中的内容。<br>1.仅在该段{{ }}范围生效；<br>2.可与 options同时生效。 |
- **json_value**
需要注意的是当您在使用type=array类型时，如若对话服务输出的内容较长直至超出所设置的tokens_to_generate的数值，那么将会表现出array生成到一半即会被截断的结果。
如果您出现以上所描述的情况，可以通过以下操作进行解决：调整tokens_to_generate的上限数值，并且同步修改优化prompt，让模型感知到array需要多长。同时应尽量避免使用“说几个”、“给一些”等不确定数量的说法。
**Tab: 单层请求示例**

**请求示例**
```json
{
    "model":"MiniMax-Text-01",
    "tokens_to_generate":1024,
    "temperature":0.9,
    "top_p":0.95,
    "stream":False,
    "beam_width":1,
    "reply_constraints":{
        "sender_type":"BOT",
        "sender_name":"MM智能助理",
        "glyph":{
            "json_properties":{
                "name":{
                    "type":"string"
                },
                "age":{
                    "type":"number"
                },
                "is_student":{
                    "type":"boolean"
                },
                "is_boy":{
                    "type":"boolean"
                },
                "courses":{
                    "type":"object",
                    "properties":{
                        "name":{
                            "type":"string"
                        },
                        "score":{
                            "type":"number"
                        }
                    }
                }
            },
            "type":"json_value"
        }
    },
    "sample_messages":[],
    "messages":[
        {
            "sender_type":"USER",
            "sender_name":"用户",
            "text":"帮我提取以下信息的关键词：我叫月无双，来自太阳大学，今年18岁，刚刚考完《斗气化翼》课程，得了99.99分"
        }
    ],
    "bot_setting":[
        {
            "bot_name":"MM智能助理",
            "content":"MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
        }
    ],
    "trace_id":"!231312"
}
```

**Tab: 嵌套请求示例**

**请求示例**
```json
{
    "model":"MiniMax-Text-01",
    "tokens_to_generate":1024,
    "temperature":0.9,
    "top_p":0.95,
    "stream":False,
    "beam_width":1,
    "reply_constraints":{
        "sender_type":"BOT",
        "sender_name":"MM智能助理",
        "glyph":{
            "type":"json_value",
            "json_properties":{
                "student":{
                    "type":"object",
                    "property_list":[
                        {
                            "properties":{
                                "name":{
                                    "type":"string"
                                }
                            }
                        },
                        {
                            "properties":{
                                "mark":{
                                    "type":"string"
                                }
                            }
                        },
                        {
                            "properties":{
                                "mark1":{
                                    "type":"string"
                                }
                            }
                        },
                        {
                            "properties":{
                                "age":{
                                    "type":"boolean"
                                }
                            }
                        },
                        {
                            "properties":{
                                "comments":{
                                    "type":"array",
                                    "items":{
                                        "type":"string"
                                    }
                                }
                            }
                        }
                    ]
                }
            }
        }
    },
    "sample_messages":[],
    "messages":[
        {
            "sender_type":"USER",
            "sender_name":"用户",
            "text":"帮我提取以下信息的关键词，最后生成一系列积极的评语：我叫月无双，来自太阳大学，今年18岁，刚刚考完《斗气化翼》课程，得了99.99分。"
        }
    ],
    "bot_setting":[
        {
            "bot_name":"MM智能助理",
            "content":"MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
        }
    ],
    "trace_id":"!231312"
}
```
**Tab: 单层返回示例**

**返回示例**
```json
{
    "created":1695026998,
    "model":"abab5.5-chat",
    "reply":"{\n  \"name\":\"月无双\",\n  \"is_student\":true,\n  \"courses\":{\n    \"name\":\"斗气化翼\",\n    \"score\":99.99\n  },\n  \"age\":18\n}",
    "choices":[
        {
            "finish_reason":"stop",
            "messages":[
                {
                    "sender_type":"BOT",
                    "sender_name":"MM智能助理",
                    "text":"{\n  \"name\":\"月无双\",\n  \"is_student\":true,\n  \"courses\":{\n    \"name\":\"斗气化翼\",\n    \"score\":99.99\n  },\n  \"age\":18\n}"
                }
            ]
        }
    ],
    "usage":{
        "total_tokens":350
    },
    "input_sensitive":false,
    "output_sensitive":false,
    "id":"0157422df0a266482950e56badee8bfc",
    "base_resp":{
        "status_code":0,
        "status_msg":""
    }
}
```

**Tab: 嵌套返回示例**

**返回示例**
```json
200
{
  "created":1706170414,
  "model":"MiniMax-Text-01",
  "reply":"{\n  \"student\":{\n    \"name\":\"月无双\",\n    \"mark\":\"99.99\",\n    \"mark1\":\"99.99\",\n    \"age\":true,\n    \"comments\":[\"优秀\",\"杰出\",\"大师\",\"专业\",]\n  }\n}",
  "choices":[
    {
      "finish_reason":"stop",
      "messages":[
        {
          	"sender_type":"BOT",
          	"sender_name":"MM智能助理",
          	"text":"{\n  \"student\":{\n    \"name\":\"月无双\",\n    \"mark\":\"99.99\",\n    \"mark1\":\"99.99\",\n    \"age\":true,\n    \"comments\":[\"优秀\",\"杰出\",\"大师\",\"专业\",]\n  }\n}"
        }
      ]
    }
  ],
    "usage":{
      "total_tokens":399
    },
    "input_sensitive":false,
    "output_sensitive":false,
    "id":"0157422df0a266482950e56badee8bfc",
    "base_resp":{
        "status_code":0,
        "status_msg":""
    }
}
```
##### 函数调用（function calling）
在使用基础功能参数的同时，额外添加functions参数方可使用function calling功能。
该功能可触发调用用户自有函数（如客户自有的会议记录查询函数），得到自有函数的参考内容 ，辅助大模型生成结果。可以帮助模型获取用户希望模型获取的相应信息，以使得模型返回更符合用户个人实际情况的回复。
特别是在一些需要基于客户自有信息知识的问题回答上，并且构建了该信息知识的API，那么用户可以考虑使用function calling功能，使用场景示例如下：
> 1.真实场景query：“在会议ID：12345的会议上，A同学有没有提到大模型如何应用的事情？”
> 2.用户自有函数，函数功能：可基于会议ID和参会人检索原始的会议记录
> 3.基于自有函数返回的参考内容：“会议号：12345，参会人：A同学，发言内容：我们要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。”
> 4.大模型基于参考内容回答query的结果：“A同学在会议ID：12345的会议上提到了大模型如何应用的事情。他说要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。”
以上，如果仅靠大模型来回答query时无法得到符合预期的回复，因此在这时就需要使用function calling功能。
注：abab5.5s-chat不支持使用function call。
调用示例如下：
**Tab: 函数调用示例**

**函数调用请求示例**
```json
{
    "model":"MiniMax-Text-01",
    "stream":false,
    "messages":[
        {
            "sender_type":"USER",
            "sender_name":"用户",
            "text":"在会议ID：12345的会议上，A同学有没有提到大模型如何应用的事情？"
        }
    ],
    "bot_setting":[
        {
            "bot_name":"MM助手",
            "content":"MM智能助理是MiniMax自主研发的大型语言模型，回答问题简洁有条理，没有调用其他产品接口。MiniMax是一家中国科技公司，致力于大模型相关的研究。"
        }
    ],
    "reply_constraints":{
        "sender_type":"BOT",
        "sender_name":"MM助手"
    },
    "functions":[
        {
            "name":"mention_something",
            "description":"当问到会议中提及事项时，用于根据会议号和参会人进行历史会议记录的检索",
            "parameters":{
                "type":"object",
                "properties":{
                    "meeting_id":{
                        "type":"string",
                        "description":"会议id"
                    },
                    "participants":{
                        "type":"string",
                        "description":"参会人"
                    }
                },
                "required":[
                    "meetingid",
                    "participants"
                ]
            }
        }
    ]
    "function_call":{
        "type":"auto"
    },
}
```

**Tab: 整体调用示例**

**整体请求示例**
```python
import requests
import json

group_id="请填写您的group_id"
api_key="请填写您的api_key"


# 1. 定义自己的function实现

def mention_something(meeting_id:str,participants:str) -> str:
    # 用户自主补充调用自由函数的具体代码，如从数据库中获取会议搜索记录，使用json返回格式给到参考内容让模型回复即可
    return """
{
   "status":"success",
   "message":"以下是会议记录的查询结果：",
   "data":
      {
         "meetingid":"12345",
         "participants":"A同学",
         "contents":"我们要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。",
      }
}
"""

# 2. 使用functions选项调用ccp接口

url = f"https://api.minimax.chat/v1/text/chatcompletion_pro?GroupId={group_id}"
headers = {"Authorization":f"Bearer {api_key}", "Content-Type":"application/json"}
payload = {
    "bot_setting":[
        {
            "bot_name":"MM智能助理",
            "content":"MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。",
        }
    ],
    "reply_constraints":{"sender_type":"BOT", "sender_name":"MM智能助理"},
    "functions":[
        {
            "name":"mention_something",
            "description":"搜索给定聊天记录，总结并返回相关结果",
            "parameters":{
                "type":"object",
                "properties":{"meeting_id":{"type":"string", "description":"会议id"},
                               "participants":{"type":"string", "description":"参会人"}},
                "required":["meeting_id","participants"],
            },
        }
    ],
    "sample_messages":[],
    "messages":[
        {"sender_type":"USER", "sender_name":"用户",
         "text":"在会议ID：12345的会议上，A同学有没有提到大模型如何应用的事情？"}
    ],
    "model":"MiniMax-Text-01",
    "tokens_to_generate":1000,
    "temperature":0.01,
    "top_p":0.95,
    "stream":False,
}

response = requests.post(url, headers=headers, json=payload)

print(response.status_code)
print(response.text)

# 3. 解析ccp接口的响应，并调用自己的function

assert response.status_code == 200
json_resp = response.json()
assert json_resp["base_resp"]["status_code"] == 0
assert json_resp["function_call"]["name"] == "mention_something"
assert json_resp["function_call"]["arguments"] != ""

arguments = json.loads(json_resp["function_call"]["arguments"])
print(arguments)
func_resp = mention_something(arguments["meeting_id"],["participants"])
print(func_resp)

# 4. 将自己的function的响应结果返回给ccp接口得到回复
payload["messages"].extend(json_resp["choices"][0]["messages"])
payload["messages"].append(
    {"sender_type":"FUNCTION", "sender_name":"MM智能助理", "text":func_resp}
)

response = requests.post(url, headers=headers, json=payload)

print(response.status_code)
print(response.text)
```
**Tab: 函数调用返回参考**

**函数调用返回参考**
```json
{
    "created":1689748978,
    "model":"MiniMax-Text-01",
    "reply":"",
    "choices":[
        {
            "finish_reason":"stop",
            "messages":[
                {
                    "sender_type":"BOT",
                    "sender_name":"MM助手",
                    "text":"",
                    "function_call":{
                        "name":"mention_something",
                        "arguments":"{\"meeting_id\":\"12345\",\"participants\":\"A同学\"}"
                    }
                }
            ]
        }
    ],
    "usage":{
        "total_tokens":474,
        "tokens_with_added_plugin":0
    },
    "input_sensitive":False,
    "output_sensitive":False,
    "id":"0106b8ed7728402ccda8ee3832cfce81",
    "function_call":{
        "name":"mention_something",
        "arguments":"{\"meeting_id\":\"12345\",\"participants\":\"A同学\"}"
    },
    "base_resp":{
        "status_code":0,
        "status_msg":""
    }
}
```

**Tab: 整体调用返回参考**

**整体返回参考**
```python
{"created":1697458106,"model":"MiniMax-Text-01","reply":"","choices":[{"finish_reason":"stop","messages":[{"sender_type":"BOT","sender_name":"MM智能助理","text":"","function_call":{"name":"mention_something","arguments":"{\"meeting_id\":\"12345\", \"participants\":\"A同学\"}"}}]}],"usage":{"total_tokens":460,"tokens_with_added_plugin":0},"input_sensitive":false,"output_sensitive":false,"id":"017c5ab84c200e9ca72302bd3770ba6a","function_call":{"name":"mention_something","arguments":"{\"meeting_id\":\"12345\", \"participants\":\"A同学\"}"},"base_resp":{"status_code":0,"status_msg":""}}
{'meeting_id': '12345', 'participants': 'A同学'}

{
   "status":"success",
   "message":"以下是会议记录的查询结果：",
   "data":
      {
         "meetingid":"12345",
         "participants":"A同学",
         "contents":"我们要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。",
      }
}

200
{"created":1697458108,"model":"MiniMax-Text-01","reply":"A同学提到了大模型如何应用的事情，他说要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。","choices":[{"finish_reason":"stop","messages":[{"sender_type":"BOT","sender_name":"MM智能助理","text":"A同学提到了大模型如何应用的事情，他说要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。"}]}],"usage":{"total_tokens":880},"input_sensitive":false,"output_sensitive":false,"id":"017c5aba33ef6091c9510356aae8d42e","base_resp":{"status_code":0,"status_msg":""}}
```
##### 知识库（retrieval）
为了更好的配合Retrieval接口的使用，我们新增了请求体的参数，以便支持高效检索文档信息，准确回答专业问题，为大模型补充领域知识。
知识库检索相关操作流程如下：
1. 文档上传file，得到file_id：[File（文档）](https://platform.minimaxi.com/document/file?key=6685458335a2d55137ca9681)
2. 新建知识库时关联上传好的file_id，并得到knowledge_base_id：[Retrieval（知识库检索）](https://platform.minimaxi.com/document/retrieval?key=66718faca427f0c8a5701617)
3. 调用Chatcompletion或者Chatcompletion pro时配置knowledge_base_id，使用query检索到的knowledge_base_id分片信息回答问题
**请求示例**
```curl
curl --location "https://api.minimax.chat/v1/text/chatcompletion_pro?GroupId=${group_id}" \
--header "authorization: Bearer $MiniMax_API_KEY"\
--header "content-type: application/json" \
--data '{
    "model":"MiniMax-Text-01",
    "tokens_to_generate":1024,
    "temperature":0.9,
    "top_p":0.95,
    "stream":false,
    "beam_width":1,
    "reply_constraints":{
        "sender_type":"BOT",
        "sender_name":"MM智能助理"
    },
    "sample_messages":[],
    "messages":[
        {
            "sender_type":"USER",
            "sender_name":"用户",
            "text":"齐天是谁"
        }
    ],
    "bot_setting":[
        {
            "bot_name":"MM智能助理",
            "content":"MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
        }
    ],
    "knowledge_base_param":{
        "knowledge_base_id":${knowledge_base_id}
    }
}'
```

**python**
```python
import requests
group_id="请填写您的group_id"
api_key="请填写您的api_key"
url = "https://api.minimax.chat/v1/text/chatcompletion_pro?GroupId=" + group_id
payload = {
"bot_setting":[
{
"bot_name":"MM智能助理",
"content":"MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。",
}
],
"messages":[{"sender_type":"USER", "sender_name":"小明", "text":"帮我用英文翻译下面这句话：我是谁"}],
"reply_constraints":{"sender_type":"BOT", "sender_name":"MM智能助理"},
"model":"MiniMax-Text-01",
"tokens_to_generate":1034,
"temperature":0.01,
"top_p":0.95,
"knowledge_base_param":{
"knowledge_base_id":"请填写您的knowledge_base_id"
}
}
headers = {"Content-Type":"application/json", "Authorization":"Bearer " + api_key}
response = requests.request("POST", url, headers=headers, json=payload)
print(response.status_code)
print(response.text)
```

## 文本合成
- 原始链接: https://platform.minimaxi.com/document/Chatcompletion_v2?key=68abd8bcaf3fef26183d5a33
- 文档ID: `68abd8bcaf3fef26183d5a33`
# 文本合成

<request method="POST" url="https://api.minimaxi.com/v1/text/chatcompletion_v2"></request>

<container>

  <left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
模型 ID。可选值：`MiniMax-M1`, `MiniMax-Text-01`。注：`MiniMax-M1` 为推理模型，为获得最佳体验建议使用流式输出
<hr/>

#### stream    <gray>bool</gray>
是否使用流式传输，默认为 `false`。设置为 `true` 后，响应将分批返回
<hr/>

#### max_tokens    <gray>int64</gray>
指定生成内容长度的上限（Token 数）。超过上限的内容会被截断。如果生成因 `length` 原因中断，请尝试调高此值
默认值：`MiniMax-M1` 为 8192，`MiniMax-Text-01` 为 2048
<hr/>

#### temperature    <gray>float</gray>
温度系数，影响输出随机性，取值范围 (0, 1]，取值范围 (0, 1]。值越高，输出越随机；值越低，输出越确定
  `MiniMax-M1` 默认值为 1.0，推荐范围 [0.8, 1.0]
  `MiniMax-Text-01` 默认值为 0.1，适用于答案明确的场景，对于文案生成等发散性场景可适当调高至 0.7-1.0
<hr/>

#### top_p    <gray>float</gray>
采样策略，影响输出随机性，取值范围 (0, 1]，各模型默认为 0.95
<hr/>

#### messages    <gray>array</gray> <red>required</red>
包含对话历史的消息列表
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### role    <gray>string</gray> <red>required</red>
消息发送者的角色
- `system`：设定模型的角色和行为
- `user`：用户的输入
- `assistant`：模型的历史回复，也可包含对工具的调用请求
- `tool`：工具调用的返回结果
<hr/>

#### name    <gray>string</gray>
发送者的名称。若同一类型的角色有多个，须提供具体名称以区分
<hr/>

#### content    <gray>string or array</gray> <red>required</red>
消息内容。纯文本输入时为 `string`；图文混合输入时为 `array`
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### type    <gray>string</gray>
内容类型，`text` 或 `image_url`
<hr/>

#### text    <gray>string</gray>
文本内容
<hr/>

#### image_url    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### url    <gray>string</gray>
图片的公网 URL 或 Base64 编码的 Data URL

</card>
</details>

</card>
</details>
<hr/>

#### tool_calls    <gray>array</gray>
模型决定调用的工具列表。当 `role` 为 `assistant` 时出现
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### id    <gray>string</gray> <red>required</red>
工具调用的唯一 ID
<hr/>

#### type    <gray>string</gray> <red>required</red>
工具类型，固定为 `function`
<hr/>

#### function    <gray>object</gray> <red>required</red>
具体的函数调用信息
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### name    <gray>string</gray> <red>required</red>
要调用的函数名称
<hr/>

#### arguments    <gray>string</gray> <red>required</red>
包含函数参数的 JSON 字符串

</card>
</details>

</card>
</details>

</card>
</details>
<hr/>

#### tools    <gray>array</gray>
可供模型选择调用的工具列表
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### type    <gray>string</gray> <red>required</red>
工具类型，当前仅支持 `function`
<hr/>

#### function    <gray>object</gray> <red>required</red>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### name    <gray>string</gray> <red>required</red>
函数名称
<hr/>

#### description    <gray>string</gray> <red>required</red>
对函数功能的描述
<hr/>

#### parameters    <gray>object</gray> <red>required</red>
函数的参数

</card>
</details>

</card>
</details>
<hr/>

#### tool_choice    <gray>string</gray>
控制模型如何使用工具。`none` (不调用) 或 `auto` (自主决定)
<hr/>

#### response_format    <gray>object</gray>
指定模型输出的格式，当前仅 `MiniMax-Text-01` 支持此参数
设置为 `{ "type": "json_schema", "json_schema": {...} }` 可强制模型按指定的 JSON Schema 结构输出
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### type    <gray>string</gray> <red>required</red>
格式类型，仅支持 `json_schema`
<hr/>

#### json_schema    <gray>object</gray> <red>required</red>
定义输出格式的 JSON Schema 对象
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### name    <gray>string</gray> <red>required</red>
格式名称。
- 须由a-z、A-Z、0-9 组成
- 最大长度为64字符
<hr/>

#### description    <gray>string</gray>
格式的描述。用于模型确定如何以该格式进行输出
<hr/>

#### schema    <gray>object</gray>
JSON Schema 定义。
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### type    <gray>string</gray> <red>required</red>
取值应为`object`
<hr/>

#### properties    <gray>object</gray> <red>required</red>
详细定义格式化输出所需的内容
支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。
请注意，使用结构化输出时，所有字段或函数参数都必须指定为required

<codetabs title="参数示例">
<codetab language="json">
"json_schema": {
  "name": "user_analysis",
  "description": "User behavior analysis result",
  "schema": {
    "type": "object",
    "properties": {
      "analysis": {
        "type": "string",
        "description": "Content analysis based on rules"
      },
      "summary_words": {
        "type": "array",
        "items": {
          "type": "string"
        },
        "description": "Content keywords summary, limited to three words"
      },
      "content_result": {
        "type": "string",
        "enum": [
          "低俗色情违规",
          "广告引流违规",
          "违法违规",
          "涉政违规",
          "自杀自残违规",
          "未成年违规",
          "攻击辱骂违规",
          "其他违规",
          "正常"
        ],
        "description": "Content classification result"
      },
      "risk_words": {
        "type": "array",
        "items": {
          "type": "string"
        },
        "description": "Keywords indicating violations, emojis converted to text"
      },
      "correct_rate": {
        "type": "number",
        "description": "Credibility score, from 0.00 to 1.00"
      },
      "content_risk_level": {
        "type": "string",
        "enum": ["high", "medium", "low"],
        "description": "Risk concentration level in content"
      }
    },
    "required": [
      "analysis",
      "summary_words",
      "content_result",
      "risk_words",
      "correct_rate",
      "content_risk_level"
    ]
  }
}
</codetab>

</codetabs>

</card>
</details>

</card>
</details>

</card>
</details>



<hr/>

#### stream_options    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### include_usage    <gray>bool</gray>
默认为 `false`。若设为 `true`，流式响应的最后一个数据块将包含本次请求完整的 `usage` 统计信息

</card>
</details>

<br/>

### 返回体
<hr/>

#### id    <gray>string</gray>
本次响应的唯一 ID
<hr/>

#### choices    <gray>array</gray>
响应选择列表
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### finish_reason    <gray>string</gray>
生成停止的原因：`stop` (自然结束), `length` (达到 `max_tokens` 上限), `tool_calls` (调用工具)
<hr/>

#### index    <gray>integer</gray>
选项的索引，从 0 开始
<hr/>

#### message    <gray>object</gray>
模型生成的完整回复
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### content    <gray>string</gray>
文本回复内容
<hr/>

#### role    <gray>string</gray>
角色，固定为 `assistant`
<hr/>

#### tool_calls    <gray>array</gray>
模型请求调用的工具列表
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### id    <gray>string</gray> 
工具调用的唯一 ID
<hr/>

#### type    <gray>string</gray>
工具类型，固定为 `function`
<hr/>

#### function    <gray>object</gray>
具体的函数调用信息
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

##### name  <gray>string</gray>
要调用的函数名称
<hr/>

##### arguments <gray>string</gray>
一个 JSON 字符串，包含调用该函数的参数

</card>
</details>

</card>
</details>

</card>
</details>

</card>
</details>
<hr/>

#### created    <gray>int</gray>
响应创建的 Unix 时间戳（秒）
<hr/>

#### model    <gray>string</gray>
本次请求使用的模型 ID
<hr/>

#### object    <gray>string</gray>
对象类型。非流式为 `chat.completion`，流式为 `chat.completion.chunk`
<hr/>

#### usage    <gray>object</gray>
本次请求的 Token 使用情况统计
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### total_tokens    <gray>int</gray>
消耗的总 Token 数

</card>
</details>
<hr/>

#### input_sensitive    <gray>bool</gray>
输入内容是否命中敏感词。如果输入内容严重违规，接口会返回内容违规错误信息，回复内容为空
<hr/>

#### input_sensitive_type    <gray>int64</gray>
输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他
<hr/>

#### output_sensitive    <gray>bool</gray>
输出内容是否命中敏感词。如果输出内容严重违规，接口会返回内容违规错误信息，回复内容为空
<hr/>

#### output_sensitive_type    <gray>int64</gray>
输出命中敏感词类型
<hr/>

#### base_resp    <gray>object</gray>
错误状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码
- 1000: 未知错误
- 1001: 请求超时
- 1002: 触发限流,
- 1004: 鉴权失败
- 1008: 余额不足
- 1013: 服务内部错误
- 1027: 输出内容错误
- 1039: Token 超出限制
- 2013: 参数错误
更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
错误详情

</card>
</details>

  </left>

  <right>

  <tabs>
  <tab title="基础示例">

<codetabs title="请求示例">
<codetab language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/text/chatcompletion_v2"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "MiniMax-M1",
    "messages": [
        {
            "role": "system",
            "name": "MiniMax AI"
        },
        {
            "role": "user",
            "name": "用户", # 选填字段
            "content": "你好"
        }
    ]
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())
</codetab>
<codetab language="curl">
curl -X POST "https://api.minimaxi.com/v1/text/chatcompletion_v2" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "model": "MiniMax-M1",
    "messages": [
      {
        "role": "system",
        "name": "MiniMax AI"
      },
      {
        "role": "user",
        "name": "用户",
        "content": "你好"
      }
    ]
  }'
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
  "id": "04ecb5d9b1921ae0fb0e8da9017a5474",
  "choices": [
    {
      "finish_reason": "stop",
      "index": 0,
      "message": {
        "content": "您好！请问有什么可以帮您？",
        "role": "assistant",
        "name": "MiniMax AI",
        "audio_content": "",
        "reasoning_content": "...省略"
      }
    }
  ],
  "created": 1755153113,
  "model": "MiniMax-M1",
  "object": "chat.completion",
  "usage": {
    "total_tokens": 249,
    "total_characters": 0,
    "prompt_tokens": 26,
    "completion_tokens": 223,
    "completion_tokens_details": {
      "reasoning_tokens": 214
    }
  },
  "input_sensitive": false,
  "output_sensitive": false,
  "input_sensitive_type": 0,
  "output_sensitive_type": 0,
  "output_sensitive_int": 0,
  "base_resp": {
    "status_code": 0,
    "status_msg": ""
  }
}
</codetab>
</codetabs>

</tab>
<tab title="流式">

<codetabs title="请求示例">
<codetab language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/text/chatcompletion_v2"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "MiniMax-M1",
    "messages": [
        {
           "role": "system", 
           "name": "MiniMax AI"
         },
        {
           "role": "user", 
            "name": "用户", 
            "content": "你好"
         },
    ],
    "stream": True,
}

response = requests.post(url, headers=headers, json=payload, stream=True)

for chunk in response.iter_lines():
    if chunk:
        chunkStr = chunk.decode("utf-8")
        print(chunkStr)
        print("————————————————————")
</codetab>
<codetab language="curl">
curl -X POST "https://api.minimaxi.com/v1/text/chatcompletion_v2" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "model": "MiniMax-M1",
    "messages": [
      {
        "role": "system",
        "name": "MiniMax AI"
      },
      {
        "role": "user",
        "name": "用户",
        "content": "你好"
      }
    ],
    "stream": true
  }'
</codetab>

</codetabs>

<codetabs title="返回示例">
<codetab language="python">
data: {"id":"04ecb6e75e57edd157dee0642353bd0f","choices":[{"index":0,"delta":{"content":"","role":"assistant","name":"MiniMax AI","audio_content":"","reasoning_content":"好的"}}],"created":1755153383,"model":"MiniMax-M1","object":"chat.completion.chunk","usage":{"total_tokens":0,"total_characters":0},"input_sensitive":false,"output_sensitive":false,"input_sensitive_type":0,"output_sensitive_type":0,"output_sensitive_int":0}
————————————————————
...省略
————————————————————
data: {"id":"04ecb6e75e57edd157dee0642353bd0f","choices":[{"index":0,"delta":{"content":"你好！有什么我可以帮","role":"assistant","name":"MiniMax AI","audio_content":"","reasoning_content":""}}],"created":1755153383,"model":"MiniMax-M1","object":"chat.completion.chunk","usage":{"total_tokens":0,"total_characters":0},"input_sensitive":false,"output_sensitive":false,"input_sensitive_type":0,"output_sensitive_type":0,"output_sensitive_int":0}
————————————————————
data: {"id":"04ecb6e75e57edd157dee0642353bd0f","choices":[{"finish_reason":"stop","index":0,"delta":{"content":"你的吗？","role":"assistant","name":"MiniMax AI","audio_content":""}}],"created":1755153383,"model":"MiniMax-M1","object":"chat.completion.chunk","usage":{"total_tokens":0,"total_characters":0},"input_sensitive":false,"output_sensitive":false,"input_sensitive_type":0,"output_sensitive_type":0,"output_sensitive_int":0}
————————————————————
</codetab>

<codetab language="curl">
{"id":"02ff7eb7fe6fb505b9d5cb6945a1a98b","choices":[{"index":0,"delta":{"content":"你好","role":"assistant"}}],"created":1722829751,"model":"MiniMax-M1","object":"chat.completion.chunk","output_sensitive":false,"input_sensitive_type":0,"output_sensitive_type":0}

{"id":"02ff7eb7fe6fb505b9d5cb6945a1a98b","choices":[{"finish_reason":"stop","index":0,"delta":{"content":"！有什么可以帮助你的吗？","role":"assistant"}}],"created":1722829751,"model":"MiniMax-M1","object":"chat.completion.chunk","output_sensitive":false,"input_sensitive_type":0,"output_sensitive_type":0}

{"id":"02ff7eb7fe6fb505b9d5cb6945a1a98b","choices":[{"finish_reason":"stop","index":0,"message":{"content":"你好！有什么可以帮助你的吗？","role":"assistant"}}],"created":1722829751,"model":"MiniMax-M1","object":"chat.completion","usage":{"total_tokens":73},"input_sensitive":false,"output_sensitive":false,"input_sensitive_type":0,"output_sensitive_type":0,"base_resp":{"status_code":0,"status_msg":""}}
</codetab>
</codetabs>

</tab>
<tab title="图像输入">

<codetabs title="请求示例">
<codetab language="python">
import base64
import requests
import os

url = "https://api.minimaxi.com/v1/text/chatcompletion_v2"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

# base64编码形式输入
# with open("your_file_path", "rb") as image_file:
#     data = base64.b64encode(image_file.read()).decode("utf-8")


payload = {
    "model": "MiniMax-Text-01",
    "messages": [
        {
            "role": "system",
            "name": "MiniMax AI",
            "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。",
        },
        {
            "role": "user",
            "name": "用户",
            "content": [
                {"type": "text", "text": "这个图代表的是什么呢"},
                # base64编码形式输入参考
                # {
                #     "type": "image_url",
                #     "image_url": {
                #         "url": f"data:image/jpeg;base64,{data}"}
                # }
                # 图片url形式输入参考
                {
                    "type": "image_url",
                    "image_url": {
                        "url": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg"
                    },
                },
            ],
        },
    ],
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())
</codetab>
<codetab language="curl">
curl -X POST "https://api.minimaxi.com/v1/text/chatcompletion_v2" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "model": "MiniMax-M1",
    "messages": [
      {
        "role": "system",
        "name": "MM智能助理",
        "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
      },
      {
        "role": "user",
        "name": "用户",
        "content": [
          {
            "type": "text",
            "text": "小老鼠在干嘛呢"
          },
          {
            "type": "image_url",
            "image_url": {
              "url": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg"
            }
          }
        ]
      }
    ]
  }'
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="python">
ChatCompletionMessage(content='这张图片展示了一只可爱的仓鼠，它正在喝一杯茶。仓鼠站在一个小桌子上，桌子上铺着一块编织垫子，旁边还有两盆小植物。背景中有一个时钟，整个场景看起来非常温馨和有趣。这种图片通常用来表达一种轻松、愉快和可爱的氛围，给人带来愉悦的感觉。', refusal=None, role='assistant', annotations=None, audio=None, function_call=None, tool_calls=None, name='MiniMax AI', audio_content='')
</codetab>
</codetabs>




</tabs>

</right>

</container>

## 音乐生成
- 原始链接: https://platform.minimaxi.com/document/Music Generation 01?key=687f9f34688dcf02f5bb607e
- 文档ID: `687f9f34688dcf02f5bb607e`
### 音乐生成（Music Generation）
本文档调用方式仅适用于music-01模型。
该接口支持用户上传音乐文件，获得干声+伴奏，进行音乐生成。
整体包含2个接口：**音乐文件上传**、**创建音乐生成任务**。使用步骤如下：
1. 通过文件上传接口，上传音频文件，将音频文件拆分出干声和伴奏，生成对应的 ID；
2. 通过创建音乐生成任务接口（Music Generation API），使用文件上传接口（Upload API）返回的voice_id（干声）、instrumental_id（伴奏），以及其他音乐参数进行音乐生成。
#### 官方MCP
`github链接：https://github.com/MiniMax-AI/MiniMax-MCP`
#### 上传音乐文件（Upload API）
`接口地址：https://api.minimaxi.com/v1/music_upload`
#### 接口参数说明
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| purpose | song/voice/nstrumental | 是 | 请按照分类purpose分类传入文件：<br>1.若purpose是song：需要传入包含干声和伴奏的音乐文件（干声需要是唱歌形式，不支持普通说话形式），输出vocie_id、instrumental_id；<br>2.若purpose是voice：需要传入干声并且唱歌的文件（目前不支持只是普通说话的音频），输出vocie_id；<br>3.若purpose是instrumental：需要仅包含传入伴奏的文件，输出instrumental_id。 |
| file |  | 是 | 上传的音频文件，使用文件地址，如'/user/abc.mp3'<br>格式：支持wav、mp3格式；<br>时长：10s以上，10min以内（song、instrumental、voice） |
| voice_id |  | 否 | 从音频文件中拆分出来的音色ID，当请求purpose为song、voice时才会返回。 |
| instrumental_id |  | 否 | 从音频文件中拆分出来的伴奏ID，当请求purpose为song、instrumental时才会返回。 |
| base_resp | BaseResp | 否 | 状态码。接口常见错误码：1001：超时，请重试；1002：请求频率超限；1004：api_key 未填写或错误；1008：余额不足；2013：用户填写参数错误，请检查输入；2044：没有使用该接口的权限，请联系商务及售前。<br>状态码及状态详情。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 状态码。接口常见错误码：1001：超时，请重试；1002：请求频率超限；1004：api_key 未填写或错误；1008：余额不足；2013：用户填写参数错误，请检查输入；2044：没有使用该接口的权限，请联系商务及售前。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
#### 创建音乐生成任务（Music Generation API）
`接口地址：https://api.minimaxi.com/v1/music_generation`
#### 接口参数说明
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| refer_voice | string | 否 | 生成音乐时参考音色ID。通过/v1/music_upload生成。<br>refer_voice、refer_instrumental至少一个必填，仅填写refer_voice时可正常输出音乐data，输出的音乐是一段符合refer_voice和lyrics生成的干声音乐哼唱，无伴奏 |
| refer_instrumental | string | 否 | 生成音乐时参考伴奏ID。通过/v1/music_upload生成。<br>refer_voice、refer_instrumental至少一个必填，仅填写refer_instrumental时可不传“lyrics”这个参数，则会正常输出音乐data，输出的音乐是一段符合refer_instrumental伴奏的纯伴奏音乐，无人声 |
| refer_vocal | string | 否 | 生成音乐时用来替换生成音乐音色的声音ID。可通过voice cloning生成，或者可直接使用T2A v2的推荐voice_id。传该参数时，<br>需同时传refer_voice、refer_instrumental、lyrics三个参数。 |
| model | string | 是 | 支持调用的模型：`music-01`。 |
| lyrics | string | 否 | 音乐生成的歌词。本参数非必填，使用换行符（\n）分隔每行歌词，使用两个连续换行符（\n\n）可以在歌词中间添加停顿，使用双井号（##）添加在首尾可以添加伴奏，支持最长200字符（每个汉字、标点和字母都算1个字符）。 |
| stream | bool | 否 | 是否开启流式，默认不开启 |
| audio_setting |  | 否 | 生成音乐的采样率。<br>生成音乐的比特率。<br>生成的音乐格式。 |
| audio_setting.sample_rate | int，范围[16000,24000,32000,44100] | 否 | 生成音乐的采样率。 |
| audio_setting.bitrate | int，范围[32000, 64000，128000,256000] | 否 | 生成音乐的比特率。 |
| audio_setting.format | string，范围["mp3","wav","pcm"] | 否 | 生成的音乐格式。 |
| data |  | 否 | 音乐合成状态,1: 合成中；2: 已完成。<br>音频文件16进制编码后结果,目前仅支持生成1min以内的音乐。 |
| data.status | int | 否 | 音乐合成状态,1: 合成中；2: 已完成。 |
| data.audio | string | 否 | 音频文件16进制编码后结果,目前仅支持生成1min以内的音乐。 |
| base_resp | BaseResp | 否 | 错误码。接口常见错误码：1001：超时，请重试；1002：请求频率超限；1004：api_key 未填写或错误；1008：余额不足；2013：用户填写参数错误，请检查输入；2044：没有使用该接口的权限，请联系商务及售前<br>状态码以及状态详情。<br>状态详情。 |
| base_resp.status_code | int | 否 | 错误码。接口常见错误码：1001：超时，请重试；1002：请求频率超限；1004：api_key 未填写或错误；1008：余额不足；2013：用户填写参数错误，请检查输入；2044：没有使用该接口的权限，请联系商务及售前 |
| base_resp.status_msg | string | 否 | 状态详情。 |

## Realtime
- 原始链接: https://platform.minimaxi.com/document/Realtime?key=640e0c9c5f918b4f6c4e2d58
- 文档ID: `640e0c9c5f918b4f6c4e2d58`
### Realtime API
构建低延迟、多模态体验，使用实时 API。
实时 API 使您能够构建低延迟、多模态的对话体验。它目前支持文本和音频作为输入和输出。Realtime API 是一种基于事件的有状态 API，通过 WebSocket 进行通信。WebSocket 连接需要以下参数：
- URL：`wss://api.minimax.chat/ws/v1/realtime`
- Query Parameters: `?model=abab6.5s-chat`
- Headers: `Authorization: Bearer YOUR_API_KEY`
下面是一个使用 Node.js 中的 ws 库建立套接字连接、发送消息和接收响应的简单示例。
**请求示例**
```python
import WebSocket from "ws";

const url = "wss://api.minimax.chat/ws/v1/realtime?model=abab6.5s-chat";
const ws = new WebSocket(url, {
    headers: {
        "Authorization": "Bearer Your_API_Key" 
    },
});

ws.on("open", function open() {
    console.log("Connected to server.");
    ws.send(JSON.stringify({
        type: "response.create",
        response: {
            modalities: ["text"],
            instructions: "Please assist the user.",
        }
    }));
});

ws.on("message", function incoming(message) {
    console.log(JSON.parse(message.toString()));
});
```
### 集成指南
1. 支持的音频格式，目前，Realtime API 支持的音频格式：PCM、16bits、24kHz、 单声道。音频必须是 base64 编码的音频帧块。
2. 发送事件，要将事件发送到 API，您必须发送包含事件负载数据的 JSON 字符串。在发送前确保您已连接到 API。
**Send a user mesage**
```javascript
// Make sure we are connected
ws.on('open', () => {
  // Send an event
  const event = {
    type: 'conversation.item.create',
    item: {
      type: 'message',
      role: 'user',
      content: [
        {
          type: 'input_text',
          text: 'Hello!'
        }
      ]
    }
  };
  ws.send(JSON.stringify(event));
});
```
3. 接收事件，要接收事件，请侦听 WebSocket 消息事件，并将结果解析为 JSON。
**receive events**
```javascript
ws.on('message', data => {
  try {
    const event = JSON.parse(data);
    console.log(event);
  } catch (e) {
    console.error(e);
  }
});
```
4. 常见的使用示例，以下是 API 功能的一些常见示例，供您开始使用。这些示例假定您已经实例化 WebSocket。
**Tab: Send user text**

**send user text**
```javascript
const event = {
  type: 'conversation.item.create',
  item: {
    type: 'message',
    role: 'user',
    content: [
      {
        type: 'input_text',
        text: 'Hello!'
      }
    ]
  }
};
ws.send(JSON.stringify(event));
ws.send(JSON.stringify({type: 'response.create'}));
```

**Tab: Send user audio**

**Send user audio**
```javascript
import fs from 'fs';
import decodeAudio from 'audio-decode';

// Converts Float32Array of audio data to PCM16 ArrayBuffer
function floatTo16BitPCM(float32Array) {
  const buffer = new ArrayBuffer(float32Array.length * 2);
  const view = new DataView(buffer);
  let offset = 0;
  for (let i = 0; i < float32Array.length; i++, offset += 2) {
    let s = Math.max(-1, Math.min(1, float32Array[i]));
    view.setInt16(offset, s < 0 ? s * 0x8000 : s * 0x7fff, true);
  }
  return buffer;
}

// Converts a Float32Array to base64-encoded PCM16 data
base64EncodeAudio(float32Array) {
  const arrayBuffer = floatTo16BitPCM(float32Array);
  let binary = '';
  let bytes = new Uint8Array(arrayBuffer);
  const chunkSize = 0x8000; // 32KB chunk size
  for (let i = 0; i < bytes.length; i += chunkSize) {
    let chunk = bytes.subarray(i, i + chunkSize);
    binary += String.fromCharCode.apply(null, chunk);
  }
  return btoa(binary);
}

// Using the "audio-decode" library to get raw audio bytes
const myAudio = fs.readFileSync('./path/to/audio.wav');
const audioBuffer = await decodeAudio(myAudio);
const channelData = audioBuffer.getChannelData(0); // only accepts mono
const base64AudioData = base64EncodeAudio(channelData);

const event = {
  type: 'conversation.item.create',
  item: {
    type: 'message',
    role: 'user',
    content: [
      {
        type: 'input_audio',
        audio: base64AudioData
      }
    ]
  }
};
ws.send(JSON.stringify(event));
ws.send(JSON.stringify({type: 'response.create'}));
```

**Tab: Stream user audio**

**Stream user audio**
```javascript
import fs from 'fs';
import decodeAudio from 'audio-decode';

// Converts Float32Array of audio data to PCM16 ArrayBuffer
function floatTo16BitPCM(float32Array) {
  const buffer = new ArrayBuffer(float32Array.length * 2);
  const view = new DataView(buffer);
  let offset = 0;
  for (let i = 0; i < float32Array.length; i++, offset += 2) {
    let s = Math.max(-1, Math.min(1, float32Array[i]));
    view.setInt16(offset, s < 0 ? s * 0x8000 : s * 0x7fff, true);
  }
  return buffer;
}

// Converts a Float32Array to base64-encoded PCM16 data
base64EncodeAudio(float32Array) {
  const arrayBuffer = floatTo16BitPCM(float32Array);
  let binary = '';
  let bytes = new Uint8Array(arrayBuffer);
  const chunkSize = 0x8000; // 32KB chunk size
  for (let i = 0; i < bytes.length; i += chunkSize) {
    let chunk = bytes.subarray(i, i + chunkSize);
    binary += String.fromCharCode.apply(null, chunk);
  }
  return btoa(binary);
}

// Fills the audio buffer with the contents of three files,
// then asks the model to generate a response.
const files = [
  './path/to/sample1.wav',
  './path/to/sample2.wav',
  './path/to/sample3.wav'
];

for (const filename of files) {
  const audioFile = fs.readFileSync(filename);
  const audioBuffer = await decodeAudio(audioFile);
  const channelData = audioBuffer.getChannelData(0);
  const base64Chunk = base64EncodeAudio(channelData);
  ws.send(JSON.stringify({
    type: 'input_audio_buffer.append',
    audio: base64Chunk
  }));
});

ws.send(JSON.stringify({type: 'input_audio_buffer.commit'}));
ws.send(JSON.stringify({type: 'response.create'}));
```
### Client events
#### session.update
发送此事件以更新会话的默认配置。客户端可以随时发送此事件来更新 session 配置，任何字段都可能随时更新，除了 “voice” 之外。服务器将使用 session.updated 事件进行响应，该事件显示完整的有效配置。只有存在的字段才会被更新，因此清除像 “instructions” 这样的字段的正确方法是传递一个空字符串。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为 "session.update" |
| session | object | 否 | 错误信息<br>模型可以使用的模态集。要禁用音频，请将其设置为 [“text”]。<br>默认系统指令（即 system message）附加到 model 调用之前。此字段允许客户端指导模型获得所需的响应。该模型可以被指导回答内容和格式（例如，“要非常简洁”、“行为友好”、“这里有好回应的例子”）和音频行为（例如，“快声说”、“在你的声音中注入情感”、“经常大笑”）。不能保证模型会遵循这些说明，但它们会为模型提供有关所需行为的指导。<br>可参考t2a的voiceid进行接入<br>输入音频的格式。选项包括 pcm16。<br>输出音频的格式。选项包括 pcm16。<br>单个 Assistant 响应的最大输出令牌数，支持 (0,245760] |
| session.modalities | string | 否 | 模型可以使用的模态集。要禁用音频，请将其设置为 [“text”]。 |
| session.instructions | string | 否 | 默认系统指令（即 system message）附加到 model 调用之前。此字段允许客户端指导模型获得所需的响应。该模型可以被指导回答内容和格式（例如，“要非常简洁”、“行为友好”、“这里有好回应的例子”）和音频行为（例如，“快声说”、“在你的声音中注入情感”、“经常大笑”）。不能保证模型会遵循这些说明，但它们会为模型提供有关所需行为的指导。 |
| session.voice | string | 否 | 可参考t2a的voiceid进行接入 |
| session.input_audio_format | string | 否 | 输入音频的格式。选项包括 pcm16。 |
| session.output_audio_format | string | 否 | 输出音频的格式。选项包括 pcm16。 |
| session.max_response_output_tokens | string | 否 | 单个 Assistant 响应的最大输出令牌数，支持 (0,245760] |
#### input_audio_buffer.append
发送此事件以将音频字节追加到输入音频缓冲区。音频缓冲区是您可以写入并在以后提交的临时存储，目前必须手动提交音频缓冲区。客户端可以选择在每个事件中放置多少音频，最高可达 15 MiB。与创建的其他客户端事件不同，服务器不会向此事件发送确认响应。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"input_audio_buffer.append" |
| audio | object | 否 | Base64 编码的音频字节。这必须采用会话配置中 input_audio_format 字段指定的格式。 |
#### input_audio_buffer.commit
发送此事件以提交用户输入音频缓冲区，这将在对话中创建新的用户消息项。如果输入音频缓冲区为空，则此事件将产生错误。提交输入音频缓冲区将触发输入音频转录（如果在会话配置中启用），但不会从模型创建响应。服务器将使用 input_audio_buffer.committed 事件进行响应。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"input_audio_buffer.commit" |
#### input_audio_buffer.clear
发送此事件以清除缓冲区中的音频字节。服务器将使用 input_audio_buffer.cleared 事件进行响应。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"input_audio_buffer.clear" |
#### conversation.item.create
将新 Item 添加到 Conversation 的上下文中，包括消息、函数调用和函数调用响应。此事件既可用于填充对话的 “历史记录” ，也可用于在中途添加新项目，但当前限制是它无法填充 Assistant 音频消息。如果成功，服务器将使用 conversation.item.created 事件进行响应，否则将发送错误事件。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"conversation.item.create" |
| previous_item_id | string | 否 | 前一项的 ID，新项将在其后插入。如果未设置，则新项目将附加到对话的末尾。如果设置，则允许在对话中插入项目。如果找不到 ID，将返回错误，并且不会添加该项目。 |
| item | string | 否 | 要添加到对话中的项。<br>项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。<br>项目的类型<br>项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。<br>消息发送者的角色（用户、助手、系统），仅适用于消息项。<br>消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.id | string | 否 | 项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。 |
| item.type | string | 否 | 项目的类型 |
| item.status | string | 否 | 项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。 |
| item.role | string | 否 | 消息发送者的角色（用户、助手、系统），仅适用于消息项。 |
| item.content | string | 否 | 消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.content.type | string | 否 | 内容类型:input_text、input_audio、text |
| item.content.text | string | 否 | 文本内容，用于 input_text 和 text 内容类型。 |
| item.content.audio | string | 否 | Base64 编码的音频字节，用于 input_audio 内容类型。 |
| item.content.transcript | string | 否 | 音频的转录文本，用于 input_audio 内容类型。 |
#### conversation.item.delete
当您想从对话历史记录中删除任何项目时，发送此事件。服务器将使用 conversation.item.deleted 事件进行响应，除非该项目在对话历史记录中不存在，在这种情况下，服务器将响应错误。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"conversation.item.delete" |
| item_id | string | 否 | 要删除的项目的 ID。 |
#### response.create
此事件指示服务器创建 Response，这意味着触发模型推理。Response 包含一个 Item。这些项目将附加到对话历史记录中。服务器将使用 response.created 事件、Items 和 content created 事件以及最后的 response.done 事件进行响应，以指示响应已完成。response.create 事件包括推理配置（如 instructions）和 temperature。这些字段将仅覆盖此 Response 的 Session 配置。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"response.create" |
| response | object | 否 | 要添加到对话中的项。<br>响应的唯一 ID。<br>对象类型必须为:realtime.response<br>响应的最终状态 （completed， cancelled， failed， incomplete）。<br>有关状态的其他详细信息。<br>导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。<br>响应未完成的原因。<br>导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有）<br>响应生成的输出项列表。<br>Response 的使用情况统计信息，这将对应于账单。<br>响应中的令牌总数，包括输入和输出文本token。<br>响应中使用的输入token的数量，包括文本token。<br>响应中发送的输出token的数量，包括文本token。<br>asr识别语音的时长<br>语音生成对应的字符数 |
| response.id | string | 否 | 响应的唯一 ID。 |
| response.object | string | 否 | 对象类型必须为:realtime.response |
| response.status | string | 否 | 响应的最终状态 （completed， cancelled， failed， incomplete）。 |
| response.status_details | object | 否 | 有关状态的其他详细信息。<br>导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。<br>响应未完成的原因。<br>导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有） |
| response.status_details.type | string | 否 | 导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。 |
| response.status_details.reason | string | 否 | 响应未完成的原因。 |
| response.status_details.error | object | 否 | 导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有） |
| response.status_details.error.type | string | 否 | 错误的类型。 |
| response.status_details.error.code | string | 否 | 错误代码（如果有） |
| response.output | array | 否 | 响应生成的输出项列表。 |
| response.usage | array | 否 | Response 的使用情况统计信息，这将对应于账单。<br>响应中的令牌总数，包括输入和输出文本token。<br>响应中使用的输入token的数量，包括文本token。<br>响应中发送的输出token的数量，包括文本token。<br>asr识别语音的时长<br>语音生成对应的字符数 |
| response.usage.total_tokens | integer | 否 | 响应中的令牌总数，包括输入和输出文本token。 |
| response.usage.input_tokens | integer | 否 | 响应中使用的输入token的数量，包括文本token。 |
| response.usage.output_tokens | integer | 否 | 响应中发送的输出token的数量，包括文本token。 |
| response.usage.total_asr_time | integer | 否 | asr识别语音的时长 |
| response.usage.total_audio_characters | integer | 否 | 语音生成对应的字符数 |
### Server events
#### error
发生错误时返回，这可能是客户端问题或服务器问题。大多数错误都是可恢复的，并且会话将保持打开状态，我们建议实现者默认监控和记录错误消息。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务器事件的唯一 ID。 |
| type | string | 否 | 事件类型必须为"error" |
| errors | object | 否 | 错误的详细信息。<br>错误类型（例如，“invalid_request_error”、“server_error”）。<br>错误代码（如果有）。<br>人类可读的错误消息。<br>与错误相关的参数（如果有）。<br>导致错误的 client 事件的 event_id（如果适用）。 |
| errors.type | string | 否 | 错误类型（例如，“invalid_request_error”、“server_error”）。 |
| errors.code | string | 否 | 错误代码（如果有）。 |
| errors.message | string | 否 | 人类可读的错误消息。 |
| errors.param | string | 否 | 与错误相关的参数（如果有）。 |
| errors.event_id | string | 否 | 导致错误的 client 事件的 event_id（如果适用）。 |
#### session.created
创建 Session 时返回。当新连接建立为第一个服务器事件时自动发出。此事件将包含默认的 Session 配置。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端的 ID。 |
| type | string | 否 | 事件类型必须为 “session.created |
| session | object | 否 | 错误信息<br>模型可以使用的模态集。要禁用音频，请将其设置为 [“text”]。<br>默认系统指令（即 system message）附加到 model 调用之前。此字段允许客户端指导模型获得所需的响应。该模型可以被指导回答内容和格式（例如，“要非常简洁”、“行为友好”、“这里有好回应的例子”）和音频行为（例如，“快声说”、“在你的声音中注入情感”、“经常大笑”）。不能保证模型会遵循这些说明，但它们会为模型提供有关所需行为的指导。<br>可参考t2a的voiceid进行接入<br>输入音频的格式。选项包括 pcm16。<br>输出音频的格式。选项包括 pcm16。<br>单个 Assistant 响应的最大输出令牌数，支持 (0,245760] |
| session.modalities | string | 否 | 模型可以使用的模态集。要禁用音频，请将其设置为 [“text”]。 |
| session.instructions | string | 否 | 默认系统指令（即 system message）附加到 model 调用之前。此字段允许客户端指导模型获得所需的响应。该模型可以被指导回答内容和格式（例如，“要非常简洁”、“行为友好”、“这里有好回应的例子”）和音频行为（例如，“快声说”、“在你的声音中注入情感”、“经常大笑”）。不能保证模型会遵循这些说明，但它们会为模型提供有关所需行为的指导。 |
| session.voice | string | 否 | 可参考t2a的voiceid进行接入 |
| session.input_audio_format | string | 否 | 输入音频的格式。选项包括 pcm16。 |
| session.output_audio_format | string | 否 | 输出音频的格式。选项包括 pcm16。 |
| session.max_response_output_tokens | string | 否 | 单个 Assistant 响应的最大输出令牌数，支持 (0,245760] |
#### session.updated
除非出现错误，否则在使用 session.update 事件更新会话时返回。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端的 ID。 |
| type | string | 否 | 事件类型必须为 “session.updated |
| session | object | 否 | 错误信息<br>模型可以使用的模态集。要禁用音频，请将其设置为 [“text”]。<br>默认系统指令（即 system message）附加到 model 调用之前。此字段允许客户端指导模型获得所需的响应。该模型可以被指导回答内容和格式（例如，“要非常简洁”、“行为友好”、“这里有好回应的例子”）和音频行为（例如，“快声说”、“在你的声音中注入情感”、“经常大笑”）。不能保证模型会遵循这些说明，但它们会为模型提供有关所需行为的指导。<br>可参考t2a的voiceid进行接入<br>输入音频的格式。选项包括 pcm16。<br>输出音频的格式。选项包括 pcm16。<br>单个 Assistant 响应的最大输出令牌数，支持 (0,245760] |
| session.modalities | string | 否 | 模型可以使用的模态集。要禁用音频，请将其设置为 [“text”]。 |
| session.instructions | string | 否 | 默认系统指令（即 system message）附加到 model 调用之前。此字段允许客户端指导模型获得所需的响应。该模型可以被指导回答内容和格式（例如，“要非常简洁”、“行为友好”、“这里有好回应的例子”）和音频行为（例如，“快声说”、“在你的声音中注入情感”、“经常大笑”）。不能保证模型会遵循这些说明，但它们会为模型提供有关所需行为的指导。 |
| session.voice | string | 否 | 可参考t2a的voiceid进行接入 |
| session.input_audio_format | string | 否 | 输入音频的格式。选项包括 pcm16。 |
| session.output_audio_format | string | 否 | 输出音频的格式。选项包括 pcm16。 |
| session.max_response_output_tokens | string | 否 | 单个 Assistant 响应的最大输出令牌数，支持 (0,245760] |
#### conversation.created
创建对话时返回。在会话创建后立即发出。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"conversation.created" |
| conversation | object | 否 | 要添加到对话中的项。<br>会话的唯一 ID。<br>对象类型必须为 “realtime.conversation”。 |
| conversation.id | string | 否 | 会话的唯一 ID。 |
| conversation.object | string | 否 | 对象类型必须为 “realtime.conversation”。 |
#### conversation.item.created
创建对话项时返回。有几种情况会产生此事件：
1.服务器正在生成一个 Response，如果成功，将生成一个或两个 Item，其类型为 message；
2.输入音频缓冲区已由客户端提交。服务器将获取 input audio buffer 的内容，并将其添加到新的用户消息 Item 中；
3.客户端已发送 conversation.item.create 事件以向 Conversation 添加新 Item。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"conversation.item.create" |
| previous_item_id | string | 否 | 前一项的 ID，新项将在其后插入。如果未设置，则新项目将附加到对话的末尾。如果设置，则允许在对话中插入项目。如果找不到 ID，将返回错误，并且不会添加该项目。 |
| item | string | 否 | 要添加到对话中的项。<br>项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。<br>项目的类型<br>项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。<br>消息发送者的角色（用户、助手、系统），仅适用于消息项。<br>消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.id | string | 否 | 项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。 |
| item.type | string | 否 | 项目的类型 |
| item.status | string | 否 | 项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。 |
| item.role | string | 否 | 消息发送者的角色（用户、助手、系统），仅适用于消息项。 |
| item.content | string | 否 | 消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.content.type | string | 否 | 内容类型:input_text、input_audio、text |
| item.content.text | string | 否 | 文本内容，用于 input_text 和 text 内容类型。 |
| item.content.audio | string | 否 | Base64 编码的音频字节，用于 input_audio 内容类型。 |
| item.content.transcript | string | 否 | 音频的转录文本，用于 input_audio 内容类型。 |
#### conversation.item.deleted
当客户端使用 conversation.item.delete 事件删除对话中的项目时返回。此事件用于将服务器对 conversation history 的理解与 Client 端的视图同步。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"conversation.item.delete" |
| item_id | string | 否 | 要删除的项目的 ID。 |
#### input_audio_buffer.committed
当客户端提交输入音频缓冲区时返回。item_id 属性是将创建的用户消息项的 ID，因此 conversation.item.created 事件也将发送到客户端。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 用于标识此事件的可选客户端生成的 ID。 |
| type | string | 否 | 事件类型必须为"input_audio_buffer.committed" |
| previous_item_id | string | 否 | 前一项的 ID，新项将在其后插入。 |
| item_id | string | 否 | 将创建的用户消息项的 ID。 |
##### input_audio_buffer.cleared
当客户端使用 input_audio_buffer.clear 事件清除输入音频缓冲区时返回。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务器事件的唯一 ID。 |
| type | string | 否 | 事件类型必须为"input_audio_buffer.cleared" |
#### response.created
创建新 Response 时返回。响应创建的第一个事件，其中响应的初始状态为 in_progress。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端 ID。 |
| type | string | 否 | 事件类型必须为"response.created" |
| response | object | 否 | 要添加到对话中的项。<br>响应的唯一 ID。<br>对象类型必须为:realtime.response<br>响应的最终状态 （completed， cancelled， failed， incomplete）。<br>有关状态的其他详细信息。<br>导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。<br>响应未完成的原因。<br>导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有）<br>响应生成的输出项列表。<br>Response 的使用情况统计信息，这将对应于账单。<br>响应中的令牌总数，包括输入和输出文本token。<br>响应中使用的输入token的数量，包括文本token。<br>响应中发送的输出token的数量，包括文本token。<br>asr识别语音的时长<br>语音生成对应的字符数 |
| response.id | string | 否 | 响应的唯一 ID。 |
| response.object | string | 否 | 对象类型必须为:realtime.response |
| response.status | string | 否 | 响应的最终状态 （completed， cancelled， failed， incomplete）。 |
| response.status_details | object | 否 | 有关状态的其他详细信息。<br>导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。<br>响应未完成的原因。<br>导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有） |
| response.status_details.type | string | 否 | 导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。 |
| response.status_details.reason | string | 否 | 响应未完成的原因。 |
| response.status_details.error | object | 否 | 导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有） |
| response.status_details.error.type | string | 否 | 错误的类型。 |
| response.status_details.error.code | string | 否 | 错误代码（如果有） |
| response.output | array | 否 | 响应生成的输出项列表。 |
| response.usage | array | 否 | Response 的使用情况统计信息，这将对应于账单。<br>响应中的令牌总数，包括输入和输出文本token。<br>响应中使用的输入token的数量，包括文本token。<br>响应中发送的输出token的数量，包括文本token。<br>asr识别语音的时长<br>语音生成对应的字符数 |
| response.usage.total_tokens | integer | 否 | 响应中的令牌总数，包括输入和输出文本token。 |
| response.usage.input_tokens | integer | 否 | 响应中使用的输入token的数量，包括文本token。 |
| response.usage.output_tokens | integer | 否 | 响应中发送的输出token的数量，包括文本token。 |
| response.usage.total_asr_time | integer | 否 | asr识别语音的时长 |
| response.usage.total_audio_characters | integer | 否 | 语音生成对应的字符数 |
#### response.done
当 Response 完成流式处理时返回。始终发出，无论最终状态如何。response.done 事件中包含的 Response 对象将包括 Response 中的所有输出 Item，但会省略原始音频数据。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端的 ID。 |
| type | string | 否 | 事件类型必须为"response.done" |
| response | object | 否 | 要添加到对话中的项。<br>响应的唯一 ID。<br>对象类型必须为:realtime.response<br>响应的最终状态 （completed， cancelled， failed， incomplete）。<br>有关状态的其他详细信息。<br>导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。<br>响应未完成的原因。<br>导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有）<br>响应生成的输出项列表。<br>Response 的使用情况统计信息，这将对应于账单。<br>响应中的令牌总数，包括输入和输出文本token。<br>响应中使用的输入token的数量，包括文本token。<br>响应中发送的输出token的数量，包括文本token。<br>asr识别语音的时长<br>语音生成对应的字符数 |
| response.id | string | 否 | 响应的唯一 ID。 |
| response.object | string | 否 | 对象类型必须为:realtime.response |
| response.status | string | 否 | 响应的最终状态 （completed， cancelled， failed， incomplete）。 |
| response.status_details | object | 否 | 有关状态的其他详细信息。<br>导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。<br>响应未完成的原因。<br>导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有） |
| response.status_details.type | string | 否 | 导致响应失败的错误类型，与状态字段（已取消、不完整、失败）相对应。 |
| response.status_details.reason | string | 否 | 响应未完成的原因。 |
| response.status_details.error | object | 否 | 导致响应失败的错误的说明，在状态为 failed 时填充。<br>错误的类型。<br>错误代码（如果有） |
| response.status_details.error.type | string | 否 | 错误的类型。 |
| response.status_details.error.code | string | 否 | 错误代码（如果有） |
| response.output | array | 否 | 响应生成的输出项列表。 |
| response.usage | array | 否 | Response 的使用情况统计信息，这将对应于账单。<br>响应中的令牌总数，包括输入和输出文本token。<br>响应中使用的输入token的数量，包括文本token。<br>响应中发送的输出token的数量，包括文本token。<br>asr识别语音的时长<br>语音生成对应的字符数 |
| response.usage.total_tokens | integer | 否 | 响应中的令牌总数，包括输入和输出文本token。 |
| response.usage.input_tokens | integer | 否 | 响应中使用的输入token的数量，包括文本token。 |
| response.usage.output_tokens | integer | 否 | 响应中发送的输出token的数量，包括文本token。 |
| response.usage.total_asr_time | integer | 否 | asr识别语音的时长 |
| response.usage.total_audio_characters | integer | 否 | 语音生成对应的字符数 |
#### response.output_item.added
在生成 Response 期间创建新 Item 时返回。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.output_item.added" |
| response_id | string | 否 | 项目所属的 Response 的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| item | string | 否 | 要添加到对话中的项。<br>项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。<br>项目的类型<br>项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。<br>消息发送者的角色（用户、助手、系统），仅适用于消息项。<br>消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.id | string | 否 | 项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。 |
| item.type | string | 否 | 项目的类型 |
| item.status | string | 否 | 项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。 |
| item.role | string | 否 | 消息发送者的角色（用户、助手、系统），仅适用于消息项。 |
| item.content | string | 否 | 消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.content.type | string | 否 | 内容类型:input_text、input_audio、text |
| item.content.text | string | 否 | 文本内容，用于 input_text 和 text 内容类型。 |
| item.content.audio | string | 否 | Base64 编码的音频字节，用于 input_audio 内容类型。 |
| item.content.transcript | string | 否 | 音频的转录文本，用于 input_audio 内容类型。 |
#### response.output_item.done
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.output_item.done" |
| response_id | string | 否 | 项目所属的 Response 的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| item | string | 否 | 要添加到对话中的项。<br>项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。<br>项目的类型<br>项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。<br>消息发送者的角色（用户、助手、系统），仅适用于消息项。<br>消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.id | string | 否 | 项目的唯一 ID，可以由客户端生成以帮助管理服务器端上下文，但不是必需的，因为如果未提供，服务器将生成一个。 |
| item.type | string | 否 | 项目的类型 |
| item.status | string | 否 | 项目的状态 （completed， incomplete）。这些对会话没有影响，但会接受这些值以与 conversation.item.created 事件保持一致。 |
| item.role | string | 否 | 消息发送者的角色（用户、助手、系统），仅适用于消息项。 |
| item.content | string | 否 | 消息的内容，适用于消息项。system仅支持input_text内容、User 支持input_text 和 input_audio 内容的消息项、assistant助手支持text内容的消息项。<br>内容类型:input_text、input_audio、text<br>文本内容，用于 input_text 和 text 内容类型。<br>Base64 编码的音频字节，用于 input_audio 内容类型。<br>音频的转录文本，用于 input_audio 内容类型。 |
| item.content.type | string | 否 | 内容类型:input_text、input_audio、text |
| item.content.text | string | 否 | 文本内容，用于 input_text 和 text 内容类型。 |
| item.content.audio | string | 否 | Base64 编码的音频字节，用于 input_audio 内容类型。 |
| item.content.transcript | string | 否 | 音频的转录文本，用于 input_audio 内容类型。 |
#### response.text.delta
当 “text” 内容部分的 text 值更新时返回。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.text.delta" |
| response_id | string | 否 | 响应的 ID。 |
| item_id | string | 否 | 项目的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| content_index | string | 否 | 内容部分在项的 content 数组中的索引。 |
| delta | string | 否 | 文本增量。 |
#### response.text.done
当 “text” 内容部分的 text 值完成流式处理时返回。当 Response 中断、不完整或取消时也会发出。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.text.done" |
| response_id | string | 否 | 响应的 ID。 |
| item_id | string | 否 | 项目的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| content_index | string | 否 | 内容部分在项的 content 数组中的索引。 |
| delta | string | 否 | 最终的文本内容。 |
#### response.audio_transcript.delta
更新模型生成的音频输出转录时返回。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.audio_transcript.delta" |
| response_id | string | 否 | 响应的 ID。 |
| item_id | string | 否 | 项目的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| content_index | string | 否 | 内容部分在项的 content 数组中的索引。 |
| delta | string | 否 | 转录增量。 |
#### response.audio_transcript.done
当模型生成的音频输出转录完成流式处理时返回。当 Response 中断、不完整或取消时也会发出。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.audio_transcript.done" |
| response_id | string | 否 | 响应的 ID。 |
| item_id | string | 否 | 项目的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| content_index | string | 否 | 内容部分在项的 content 数组中的索引。 |
| delta | string | 否 | 音频的最终转录。 |
#### response.audio.delta
更新模型生成的音频时返回。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.audio.delta" |
| response_id | string | 否 | 响应的 ID。 |
| item_id | string | 否 | 项目的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| content_index | string | 否 | 内容部分在项的 content 数组中的索引。 |
| delta | string | 否 | Base64 编码的音频数据增量。 |
#### response.audio.done
在模型生成的音频完成时返回。当 Response 中断、不完整或取消时也会发出。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| event_id | string | 否 | 服务端ID。 |
| type | string | 否 | 事件类型必须为"response.audio.done" |
| response_id | string | 否 | 响应的 ID。 |
| item_id | string | 否 | 项目的 ID。 |
| output_index | string | 否 | 响应中输出项的索引。 |
| content_index | string | 否 | 内容部分在项的 content 数组中的索引。 |

## SDK文档
- 原始链接: https://platform.minimaxi.com/document/SDK document?key=667190a8a427f0c8a570167f
- 文档ID: `667190a8a427f0c8a570167f`
### Assistants SDK
此文档为Assistants SDK文档，通过此文档您可以查看我们提供的SDK的详细功能以及接入示例。
我们会陆续开放SDK的能力，敬请期待。
### 语音通话SDK
以下介绍的是Assistants 语音通话的SDK使用方式，用户可自主查看进行接入。
注：本SDK仍需调用访问Assistants API服务，只是提供了sdk便于更快的接入Assistants API语音通话能力。
#### 准备
##### Android
依赖梳理
targetSdkVersion 32
minSdkVersion 20
coroutine 1.6.4
androidX
okhttp 4.11.0
android-vad:webrtc
##### iOS
依赖梳理
VoiceActivityDetector
#### 能力支持
1. 创建 Assistant & Thread
2. 开始录音 & 开始说话后 vad 检测静音自动停止
3. 手动停止录音
4. 发送语音消息请求并播放返回音频
5. 返回ASR文本及语音回复对应文本
6. 手动停止语音消息请求
#### Android
##### 使用流程
![image](https://filecdn.minimax.chat/public/9bdd773a-72ad-419f-b715-00b6498a63b2.png)
##### SDK 接入
Demo 代码地址：[https://github.com/minimax-open-platfom/MiniMaxOpenSDK-Android-Demo](https://github.com/minimax-open-platfom/MiniMaxOpenSDK-Android-Demo)
1. 添加 maven 仓库
**添加 maven 仓库**
```python
repositories {
    maven {
        url 'https://raw.githubusercontent.com/minimax-open-platfom/MiniMaxOpenSDK-Android-Repo/master'
    }
    maven { url 'https://jitpack.io' }
}
```
2. 添加依赖
**添加依赖**
```python
dependencies {
    implementation "com.minimax.opensdk_android:voice_call:1.0.0"
}
```
3. 在代码中引用 MMVoiceCallManager
**在代码中引用 MMVoiceCallManager**
```python
class MainActivity : Activity {
    
    fun onCreate() {
        // ...
        MMVoiceCallManager.init(applicationContext)
        
        // 开始使用
    }
    
    fun onDestory() {
        super.onDestory()
        MMVoiceCallManager.release()
    }
}
```
##### MMVoiceCallManager
主要使用类
下面是方法简介
###### Init
MMVoiceCallManager 初始化，需要在使用 MMVoiceCallManager 使用前调用。
**Init**
```python
/** 
  @param context applicationContext 必须传
  @param recordSampleRate 采样率 8000/16000/32000/48000
  @param vadVolume vad 收音音量 30-100 ,默认 55
  @param vadTimeout vad 检测不到人声后 x ms 后停止录音，默认 1500 ms
  @param vadMaxTime vad 可录制的最长人声，单位 ms 默认 60_000 ms
  @return Boolean 是否初始化成功
*/
fun init(context: Context, recordSampleRate: Int, vadVolume: Int, vadTimeout: Long, vadMaxTime: Long): Boolean
```
###### createAssistant
创建一个 assistant，参数对齐开平官网参数。**在子线程执行，返回可空，或抛出异常**
**createAssistant**
```python
/**
  @params params AssistantParams 数据结构对齐请求的参数 https://api.minimax.chat/document/guides/Assistants/document/assistant?id=6586b86b4da4834fd75906f6
  @return AssistantBean？ 可空，可能抛出异常
*/
fun createAssistant(params: AssistantParams): AssistantBean?

class AssistantBean()
```
###### createThread
创建一个 thread ，参数对齐开平官网参数。**子线程执行，返回可空，或抛出异常**
**createThread**
```python
/**
  @param params 创建 thread 的参数
  @return ThreadBean？ 可空
*/
fun createThread(params: ThreadParams): ThreadBean?

class ThreadBean()
```
###### startListen
开始监听，需要接入方提前申请录音权限，调用方法后，会打开录音机进行收音，并开启 vad 自动检测，检测到人声并且声音分贝高于 vadVolume 后开始录音，在收不到人声或者人声低于 vadVolume 持续 vadTimeout 毫秒后，会自动接入，并通过 RecordListener.onRecordEnd 返回音频文件。如果一直在讲话，那么最多可录制 vadMaxTime 毫秒的音频文件。
**startListen**
```python
/** 
  @param listener 录音状态监听器
*/
fun startListen(listener: RecordListener)

interface RecordListener {

  // 开始监听讲话，端上可做一些状态变化
  fun onListenStart()

  // 人声检测检测到了人声，并开始进行文件写入
  fun onRecordStart()
  
  /** 
    停止录制并返回录音结果
    @param success 是否录制成功
    @param error 如果失败的话，失败的原因
    @param audioFile 录音文件的地址 wav
  */
  fun onRecordEnd(success: Boolean, error: RecordError?, audioFile: File?)
}

enum class RecordError {
    START_RECORDER_ERROR,  // 启动录音机失败，可能是没权限，或者录音机被占用
    TOO_SHORT,  // 录音太短，可能是 vad 自动结束，但是录音时间 < 100ms 或者手动结束，录音时间不足 1s
    UNKNOWN  // 默认错误
}
```
###### stopListen
手动停止录音，当环境嘈杂，或者有其他的人声混入时，可以调用这个方法结束录音。
**stopListen**
```python
/**
  停止录音，用户可以通过 vadTimeout 来自动停止录音，也可以手动调用，调用 stopRecord 来停止，停止后，会在 onRecordEnd 里获得音频文件
*/
fun stopListen(): Boolean
```
###### isListening
录音设备是否开启
**isListening**
```python
/**
  @return 录音机是否开启
*/
fun isListening(): Boolean
```
###### sendVoiceCallMsgAndPlayAudio
发送语音并流式播放返回的语音消息。最关键的是需要在 streamParams 里手动传一个音频文件转 hex 后的字符串。输入的文件类型目前支持 wav/mp3 。
可以通过 streamParams.t2aOption.format 字段设置输出的音频格式，目前支持 mp3/flac 。
**sendVoiceCallMsgAndPlayAudio**
```python
/** 
  @param streamParams 语音通话请求参数
  @param listener 语音消息的状态监听
*/
fun sendVoiceCallMsgAndPlayAudio(streamParams: StreamParams, listener: VoiceCallMsgListener)

interface VoiceCallMsgListener {

  // sse 请求开始
  fun onRequestStart()

  /**
    @param errorCode 错误码，-1 表示服务端没有返回错误码，使用客户端兜底
    @param errorMSg 错误信息，可能包含一些敏感词，没有音频返回，以及网络请求的错误
  */
  fun onRequestError(errorCode: Int, errorMsg: String)
  
  // 收到了音频返回
  fun onAudioMsgReceived()
  
  // 语音消息开始播放
  fun onAudioMsgPlayStart()
  
  /** 
    @param success 是否是正常结束
    @param reason 错误
  */
  fun onAudioMsgPlayEnd(success: Boolean, reason: StreamAudioPlayer.Reason?)
  
  // 流程结束，无论是请求失败还是播放完成，均会走到 onComplete 接口
  fun onComplete() 
  
  // 收到了语音识别文本
  fun onAsrTextReceived(asrText: String)

  // 收到了语音回复的文本，分段输出，每次回调新增的文本
  fun onReplyTextReceived(replyText: String)

  // 语音回复文本结束
  fun onReplyTextEnd()
}

enum class Reason(
    var detail: PlaybackException? = null
) {
    BY_INTERRUPTED,         // 被打断
    BUFFER_TIMEOUT,         // 缓冲时间过长，一般是 30s，也就是收到音频数据后，30s 内不播放的话会停止
    PLAYBACK_EXCEPTION,     // 其他的播放异常
}

class StreamParams {}
```
在发送前会 check 下 hex 字符串的长度。由于 hex 是由文件转换而成的。hex.length = file.size * 2 。而 file 的最大 size 限制可以调用 AudioUtil.getMaxFileSize() 获得。
**sendVoiceCallMsgAndPlayAudio**
```python
/** 
  * @param sampleRate 初始化时输入的采样率 
  * @param audioFormat sdk 录音时使用 ENCODDING_PCM_16BIT，这里穿同样的值
  * @param channelCount sdk 录音时使用的单声道，这里传 1
  * @param duration 录音写入的最大时长，init 时输入的最大时长，默认 60000 ms
  * @return 单位 kb 
*/
fun getMaxFileSize(
    sampleRate: Int,
    audioFormat: Int,
    channelCount: Int,
    duration: Long,
): Long {
    return if (audioFormat == ENCODING_PCM_16BIT) {
        sampleRate * 2 * channelCount * (duration / 1000) / 1024
    } else {
        sampleRate * channelCount * (duration / 1000) / 1024
    }
}
```
###### cancelVoiceCallMsgAndStopPlay
取消语音通话的请求，如果正在播放回复的话，也会停止回复
**cancelVoiceCallMsgAndStopPlay**
```python
/**
  * @param CancelStreamParams 停止流式请求，可空，如果有传入的话，就停止传入的 runId 所在的请求，不传的话停止当前正在进行的流式请求 
*/
fun cancelVoiceCallMsgAndStopPlay(params: CancelStreamParams?)
```
###### release
通常在页面退出（Activity$OnDestory）等时机使用，用于清除一些录音过程中的缓存文件。如果在录音过程中退出的话，需要先调用 stopListen。同样在网络请求中退出的话，需要先调用 cancelVoiceCallMsgAndStopPlay
**release**
```python
/** 
  释放一些资源，通常在页面销毁的时候调用
*/
fun release()
```
#### iOS
##### setupAudioInfo
**setupAudioInfo**
```python
/** 
  @param recordSampleRate 采样率 8000/16000/32000/48000
  @param vadVolume vad 收音音量 30-100 目前随便拍的
  @param vadTimeout vad 检测不到人声后 x ms 后停止录音
  @param vadMinTime vad 最短需要的人声，比如啊的一声就停止了，需要处理这种情况
  @param vadMaxTime vad 可录制的最长人声，单位 ms 默认 60s
  @return OSStatus 是否初始化成功, noError为初始化成功，其他错误码 https://www.osstatus.com/
*/
func setupAudioInfo(recordSampleRate: Int, vadVolume: Int, vadTimeout: Long, vadMaxTime: Int) -> Boolean
```
##### createAssistant
##### createThread
##### startListen
**startListen**
```python
/** 
  @return Boolean 是否开始监听&录音，如果返回 false 可能是录音被占用等
*/
func startListen() -> Boolean
```
##### stopListen
**stopListen**
```python
/**
  停止录音，用户可以通过 vadTimeout 来自动停止录音，也可以手动调用，调用 stopRecord 来停止，停止后，会在 onVadEnd 里回调
  @return Boolean true 的话表示当前有录音任务且停止成功了，false 表示当前没有录音任务
*/
func stopListen() -> Boolean
```
##### sendVoiceCallMsg
**sendVoiceCallMsg**
```python
/** 
  @param voiceCallParams 语音通话请求参数
  @param completion: 语音消息的回调
  @return String taskId 发送语音消息，会生成一个 taskId ，如果想手动取消的话，将 taskId 传入
*/
func sendVoiceCallMsg(voiceCallParams: VoiceCallParams, completion: @escaping (_ data: Any?, _ error: Error?) -> String
```
##### cancelVoiceCallMsg
**cancelVoiceCallMsg**
```python
/** 
  @param taskId 任务的 ID
  @return Boolean 是否取消成功
*/
func cancelVoiceCallMsg(taskId: String?, completion: @escaping (_ data: Any?, _ error: Error?) -> Void
```
##### MNMVoiceCall
- MNMVoiceCallRecorder
- MNMVoicePlayer
- APIPhone
- id<MNMVoiceCallRecordDeleagate>
**MNMVoiceCall**
```python
  // 开始监听讲话，端上可做一些状态变化
  func listenStart()

  // 人声检测检测到了人声，并开始进行文件写入
  func recordStart()
    /** 
    停止录制并返回录音结果
    @param success 是否录制成功
    @param error 如果失败的话，失败的原因
    @param tempFile 录音文件的地址
  */
  func recordEnd(success: Boolean, error: RecordError?, url: URL?)
  
  enum class RecordError {
    UNKNOWN  // 默认错误
  }
```
- id<MNMVoiceCallPlayDelegate>
**MNMVoiceCall**
```python
  // 语音消息开始播放
  func playStart()
  
  /** 
    @param success 是否是正常结束
    @param error 错误
  */
  func playEnd(success: Boolean, error: VoiceCallMsgError?)
  
    // 收到了语音识别文本
  func onAsrTextReceived(asr: String)

  // 收到了语音回复的文本，分段输出，每次回调新增的文本
  func onReplyTextReceived(reply: String)

  // 语音回复文本结束
  func onReplyTextEnd()

  enum class VoiceCallMsgError {
    UNKNOWN  // 默认错误
  }
```

## T2A
- 原始链接: https://platform.minimaxi.com/document/T2A?key=667bde023be2027f69b71d5a
- 文档ID: `667bde023be2027f69b71d5a`
### T2A（语音生成）
该API支持基于文本到语音的同步生成，单次文本传输最大500字符。接口本身为无状态接口，即单次调用时，模型所接收到的信息量仅为接口传入内容，不涉及业务逻辑，同时模型也不存储您传入的数据。
该接口支持以下功能：
1. 支持20+系统音色、复刻音色自主选择；
2. 支持音量、语调、语速、输出格式调整；
3. 支持按比例混音功能；
4. 支持固定间隔时间控制。
该接口的适用场景：短句生成、语音聊天、在线社交等。
### API说明
`API：https://api.minimax.chat/v1/text_to_speech`
| 模型 | 适用场景 |
| --- | --- |
| speech-01 | 中文 |
| speech-02 | 中文、英文、中英混合、日文、韩文 |
#### 接口参数说明
###### 请求体（Request）参数
###### 请求成功返回（Response）参数
###### 请求失败返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | 使用发放的值 | 是 | 用户所属的组。该值应拼接在调用API的url末尾。 |
| model | string | 是 | 调用的模型版本：`speech-01`、`speech-02` |
| voice_id | string | 否 | 请求的音色编号。支持系统音色(id)以及复刻音色(id)两种类型,其中系统音色(id)如下：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| timber_weights | list[] | 否 | 音色相关信息。混合音色相关信息，混合音色相关信息，详见下方字段，如同时传入voice_id和timber_weights时，则会自动忽略voice_id，以timber_weights传递的参数为准。<br>请求的音色编号。暂时只支持系统音色(id)：参考上面voice_id的说明。<br>权重。最多支持4种音色混合，取值为整数，单一音色取值占比越高，合成音色越像。 |
| timber_weights.voice_id | string | 是 | 请求的音色编号。暂时只支持系统音色(id)：参考上面voice_id的说明。 |
| timber_weights.weight | [1, 100] | 是 | 权重。最多支持4种音色混合，取值为整数，单一音色取值占比越高，合成音色越像。 |
| speed | [0.5,2] | 否 | 生成声音的语速。可选，默认值为1.0，取值越大，语速越快。 |
| vol | (0,10] | 否 | 生成声音的音量。可选，默认值为1.0，取值越大，音量越高。 |
| output_format | string | 否 | 生成声音的音频格式。可选，默认值为mp3，可选范围：mp3、wav、pcm、flac、aac。 |
| pitch | [-12,12] | 否 | 生成声音的语调。可选，默认值为0（0为原音色输出，取值需为整数）。 |
| text | string | 是 | 需要生成的文本。长度限制<500字符（如需要控制语音中间隔时间，在字间增加<#x#>,x单位为秒，支持0.01-99.99，最多两位小数）。 |
| char_to_pitch | list<string> | 否 | 替换需要特殊标注的文字、符号及对应的注音。功能类型1，替换声调：["燕少飞/(yan4)(shao3)(fei1)"]；功能类型2，替换字符：["omg/oh my god","=/等于"]。声调用数字代替，一声（阴平）为1，二声（阳平）为2，三声（上声）为3，四声（去声）为4），轻声为5。 |
| Content-Type | string | 否 | 状态信息。如果为audio/mpeg，则表示请求成功。 |
| 二进制音频 | 二进制 | 否 | 二进制音频。生成的音频。 |
| base_resp |  | 否 | 如果请求出错，对应的错误状态码和详情。<br>1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，服务内部错误；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，服务内部错误；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
| trace_id | string | 否 | 生成ID。用于在咨询/反馈时帮助定位问题。 |
#### 已生成音频纠正工具
已生成音频纠正工具是指在用T2A/T2A pro API生成完一段音频后，用户觉得有些音频片段读音错误需要纠正，可以使用[已生成音频纠正工具](https://filecdn.minimax.chat/public/4acb7041-a892-4285-b204-66bf6314b416.zip)进行二次加工生成，使用之后就无需自己剪切音频、调用接口生成然后再进行拼接，只需要将原始的生成音频作为输入，填写好需要修改的音频片段时间戳和修改的读音，会自动得到一个纠正后的生成音频，并且仅需要支付纠正部分的生成音频字符数费用即可，纠正手段支持替换声调和替换字符，具体操作可以下载工具后查看readme.md说明进行操作。

## T2A Large
- 原始链接: https://platform.minimaxi.com/document/T2A Large?key=667bde343be2027f69b71d70
- 文档ID: `667bde343be2027f69b71d70`
### T2A Large（异步超长文本语音生成）
该API支持基于文本到语音的异步生成，单次文本生成传输最大支持1000万字符，生成的完整音频结果支持异步的方式进行检索。
该接口支持以下功能：
1. 支持20+系统音色、复刻音色自主选择；
2. 支持语调、语速、音量、比特率、采样率、输出格式自主调整；
3. 支持音频时长、音频大小等返回参数；
4. 支持时间戳（字幕）返回，精确到句；
5. 支持非法字符检测：非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；非法字符超过10%，接口不返回结果（返回报错码），请检测后再次进行请求（非法字符定义：ascii码中的控制符（不含制表符和换行符））；
提交长文本语音合成请求后，会生成file_id，音频生成结果在提交之日次日起7天内完成（T+7），生成任务完成后，可通过file_id使用[Retrieve File接口](https://platform.minimaxi.com/document/retrieval?key=66718faca427f0c8a5701617)进行下载，需要注意的是返回的url的有效期为自url返回开始的9个小时（即32400秒），超过有效期后url便会失效，生成的信息便会丢失，请注意下载信息的时间。
适用场景：整本书籍语音合成等。
### API说明
整体包含2个API：创建语音生成任务、查询语音生成任务状态。使用步骤如下：
1. 创建语音生成任务得到taskid；
2. 基于taskid查询语音生成任务状态；
3. 如果发现任务生成成功，那么可以使用file的[Retrieval API](https://platform.minimaxi.com/document/file?key=6685458335a2d55137ca9681)进行结果查看和下载
#### 创建语音生成任务
`API：https://api.minimax.chat/v1/t2a_async`
##### 模型列表
| 模型 | 使用场景 |
| --- | --- |
| speech-01 | 中文场景 |
##### 接口参数说明
###### 请求体（Request）参数
###### 返回(Response)参数
###### 额外信息字段
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | 使用发放的值 | 是 | 用户所属的组。该值应拼接在调用API的url末尾。 |
| model | string | 是 | 调用的模型版本。仅支持:`speech-01` |
| voice_id | string | 是 | 请求的音色编号。支持系统音色(id)以及复刻音色（id）两种类型，其中系统音色（id）如下：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| speed | [0.5,2] | 否 | 生成声音的语速。可选，默认值为1.0，取值越大，语速越快。 |
| vol | (0,10] | 否 | 生成声音的音量。可选，默认值为1.0，取值越大，音量越高。 |
| output_format | string | 否 | 生成声音的音频格式。可选，默认值为mp3，可选范围：mp3、wav、pcm、flac、aac。 |
| pitch | [-12,12] | 否 | 生成声音的语调。可选，默认值为0（0为原音色输出，取值需为整数）。 |
| text(files) | 长度限制<10000000字符，格式为zip | 是 | 待合成的文本。打包上传，包里应只包含txt或json文件（压缩包内为同一格式文件），json文件会有三个字段，["title", "content", "extra"]，分别表示标题，正文，作者。json里会有三个字段，["title", "content", "extra"]，分别是标题，正文，作者的话，需要产出三份结果，一共9个文件放在一个文件夹里。如果某字段不存在，或者内容为空，则不生成相应文件。 |
| audio_sample_rate | [16000,24000,32000] | 否 | 生成声音的采样率。可选，默认为32000。 |
| bitrate | [32000, 64000，128000] | 否 | 生成声音的比特率。可选，默认值为128000。 |
| char_to_pitch | list | 否 | 替换需要特殊标注的文字、符号及对应的注音。<br>功能类型1，替换声调："燕少飞/(yan4)(shao3)(fei1)"]<br>功能类型2，替换字符：["omg/oh my god","=/等于"]<br>声调用数字代替，一声（阴平）为1，二声（阳平）为2，三声（上声）为3，四声（去声）为4），轻声为5。 |
| task_id | string | 否 | 任务ID。 |
| task_token | string | 否 | 任务密钥。 |
| base_resp | dict | 否 | 如果请求出错，对应的错误状态码和详情。<br>状态码。<br>0，请求成功；1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，服务内部错误及非法字符超过10%；2013，输入格式信息不正常<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。<br>0，请求成功；1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，服务内部错误及非法字符超过10%；2013，输入格式信息不正常 |
| base_resp.status_msg | string | 否 | 错误详情。 |
| file_id | int64 | 否 | 任务创建成功后，会返回file_id，任务完成后，可通过file_id使用Retrieve File接口进行下载。当请求出错时，不返回该项，需要注意的是返回的url的有效期为自url返回开始的9个小时（即32400秒），超过有效期后url便会失效，生成的信息便会丢失，请注意下载信息的时间。 |
| audio_length | int64 | 否 | 音频时长，精确到毫秒。 |
| audio_sample_rate | int64 | 否 | 采样率，默认为24000，如客户请求参数进行调整，会根据请求参数生成。 |
| audio_size | int64 | 否 | 音频大小，单位为字节。 |
| bitrate | int64 | 否 | 比特率，默认为168000，如客户请求参数进行调整，会根据请求参数生成。 |
| word_count | int64 | 否 | 可读字数，已经发音的字数统计（不算标点等其他符号，包含汉字数字字母）。 |
| invisible_character_ratio | float64 | 否 | 非法字符占比。非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；最大不超过0.1（10%），超过进行报错。 |
| usage_characters | int64 | 否 | 计费字符数，本次语音生成的计费字符数。 |
###### 返回文件备注
| 输入文件（压缩包内） | 字段 | 输出文件 | 备注 |
| --- | --- | --- | --- |
| txt | 无 | 音频mp3 | 此文件的mp3文件 |
| txt | 无 | 字幕json | 字幕json，精确到句 |
| txt | 无 | 额外信息json | 参考下列字段 |
| json | title | 音频mp3 | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | title | 字幕json | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | title | 额外信息.json | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | content | 音频mp3 | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | content | 字幕json | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | content | 额外信息json | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | extra | 音频mp3 | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | extra | 字幕json | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
| json | extra | 额外信息.json | 字段若或为空信息，则不产生相关内容，其他内容参考txt备注 |
#### 查询语音生成任务状态
`API
:https://api.minimax.chat/query/t2a_async_query`
##### 接口参数说明
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Groupid | string | 否 | 用户所属的组。该值应拼接在调用API的url末尾。 |
| task_id | string | 否 | 任务ID，提交任务时返回。 |

## T2A Pro
- 原始链接: https://platform.minimaxi.com/document/T2A Pro?key=667bde173be2027f69b71d66
- 文档ID: `667bde173be2027f69b71d66`
### T2A pro（长文本语音生成）
该API支持基于文本到语音的同步生成，单次文本生成传输最大支持50000字符，生成的完整音频结果会同步返回。
如果追求合成速度，建议使用T2A v2，而不是T2A pro。
该接口支持以下功能：
1. 支持20+系统音色、复刻音色自主选择；
2. 支持语调、语速、音量、比特率、采样率、输出格式自主调整；
3. 支持音频时长、音频大小等返回参数；
4. 支持时间戳（字幕）返回，精确到句；
5. 支持混音功能，最多四种音色按照比例混合；
6. 支持非法字符检测：非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；非法字符超过10%，接口不返回结果（返回报错码），请检测后再次进行请求（非法字符定义：ascii码中的控制符（不含制表符和换行符））；
7. 支持自定义文本与文本之间的语音时间间隔，以实现自定义文本语音停顿时间的效果。需要注意的是文本间隔时间需设置在两个发音的文本之间，且不能设置多个连续的时间间隔。
### API说明
`API：https://api.minimax.chat/v1/t2a_pro`
#### 模型列表
| 模型 | 适用场景 |
| --- | --- |
| speech-01 | 中文 |
| speech-02 | 中文、英文、中英混合、日文、韩文 |
#### 接口参数说明
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | 使用发放的值 | 是 | 用户所属的组。该值应拼接在调用API的url末尾。 |
| model | string | 是 | 调用的模型版本：`speech-01`、`speech-02` |
| voice_id | string | 否 | 请求的音色编号。与timber_weights二选一“必填”。支持系统音色(id)以及复刻音色（id）两种类型，其中系统音色（ID）如下：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| timber_weights | list | 否 | 音色相关信息。混合音色相关信息，混合音色相关信息，详见下方字段，如同时传入voice_id和timber_weights时，则会自动忽略voice_id，以timber_weights传递的参数为准。<br>请求的音色编号。暂时只支持系统音色(id)：参考上面voice_id的说明。<br>权重。最多支持4种音色混合，取值为整数，单一音色取值占比越高，合成音色越像。 |
| timber_weights.voice_id | string | 是 | 请求的音色编号。暂时只支持系统音色(id)：参考上面voice_id的说明。 |
| timber_weights.weight | [1, 100] | 是 | 权重。最多支持4种音色混合，取值为整数，单一音色取值占比越高，合成音色越像。 |
| speed | [0.5,2] | 否 | 生成声音的语速。可选，默认值为1.0，取值越大，语速越快。 |
| vol | (0,10] | 否 | 生成声音的音量。可选，默认值为1.0，取值越大，音量越高。 |
| output_format | string | 否 | 生成声音的音频格式。可选，默认值为mp3，可选范围：mp3、wav、pcm、flac、aac。 |
| pitch | [-12,12] | 否 | 生成声音的语调。可选，默认值为0（0为原音色输出，取值需为整数）。 |
| text | string | 是 | 需要生成的文本。长度限制<50000字符（如需要控制语音中间隔时间，在字间增加<#x#>,x单位为秒，支持0.01-99.99s，最多两位小数）。支持自定义文本与文本之间的语音时间间隔，以实现自定义文本语音停顿时间的效果。需要注意的是文本间隔时间需设置在两个可以语音发音的文本之间，且不能设置多个连续的时间间隔。。 |
| audio_sample_rate | 范围[16000,24000,32000] | 是 | 生成声音的采样率。可选，默认为32000。 |
| bitrate | 范围[32000, 64000，128000] | 是 | 生成声音的比特率。可选，默认值为128000。 |
| char_to_pitch | list | 否 | 替换需要特殊标注的文字、符号及对应的注音。功能类型1，替换声调：["燕少飞/(yan4)(shao3)(fei1)"]；功能类型2，替换字符：["omg/oh my god","=/等于"]；声调用数字代替，一声（阴平）为1，二声（阳平）为2，三声（上声）为3，四声（去声）为4），轻声为5。 |
| trace_id | string | 否 | 生成id。用于在咨询/反馈时帮助定位问题。 |
| audio_file | string | 否 | 合成的音频下载链接。格式为mp3。 |
| subtitle_file | string | 否 | 合成的字幕下载链接。音频文件对应的字幕，精确到句（不超过50字），单位为毫秒，格式为json。 |
| extra_info | object | 否 | 额外信息。相关额外信息。<br>音频时长。精确到毫秒。<br>采样率。默认为24000，如客户请求参数进行调整，会根据请求参数生成。<br>音频大小。单位为字节。<br>比特率。默认为168000，如客户请求参数进行调整，会根据请求参数生成。<br>可读字数。已经发音的字数统计（不算标点等其他符号，包含汉字数字字母）。<br>非法字符占比。非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；最大不超过0.1（10%），超过进行报错。<br>消费字符数。本次语音生成的计费字符数。 |
| extra_info.audio_length | int64 | 否 | 音频时长。精确到毫秒。 |
| extra_info.audio_sample_rate | int64 | 否 | 采样率。默认为24000，如客户请求参数进行调整，会根据请求参数生成。 |
| extra_info.audio_size | int64 | 否 | 音频大小。单位为字节。 |
| extra_info.bitrate | int64 | 否 | 比特率。默认为168000，如客户请求参数进行调整，会根据请求参数生成。 |
| extra_info.word_count | int64 | 否 | 可读字数。已经发音的字数统计（不算标点等其他符号，包含汉字数字字母）。 |
| extra_info.invisible_character_ratio | float64 | 否 | 非法字符占比。非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；最大不超过0.1（10%），超过进行报错。 |
| extra_info.usage_characters | int64 | 否 | 消费字符数。本次语音生成的计费字符数。 |
| base_resp | object | 否 | 如果请求出错，对应的错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1042，非法字符超过10%；2013，输入格式信息不正常<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1042，非法字符超过10%；2013，输入格式信息不正常 |
| base_resp.status_msg | string | 否 | 错误详情。 |
| trace_id | string | 否 | 生成id。用于在咨询/反馈时帮助定位问题。 |
#### 已生成音频纠正工具
已生成音频纠正工具是指在用T2A/T2A pro API生成完一段音频后，用户觉得有些音频片段读音错误需要纠正，可以使用[已生成音频纠正工具](https://filecdn.minimax.chat/public/4acb7041-a892-4285-b204-66bf6314b416.zip)进行二次加工生成，使用之后就无需自己剪切音频、调用接口生成然后再进行拼接，只需要将原始的生成音频作为输入，填写好需要修改的音频片段时间戳和修改的读音，会自动得到一个纠正后的生成音频，并且仅需要支付纠正部分的生成音频字符数费用即可，纠正手段支持替换声调和替换字符，具体操作可以下载工具后查看readme.md说明进行操作。

## T2A Stream
- 原始链接: https://platform.minimaxi.com/document/T2A Stream?key=667bde4c3be2027f69b71d7e
- 文档ID: `667bde4c3be2027f69b71d7e`
### T2A Stream（流式语音生成）
本接口支持基于自然语言交互生成回复的能力。
本接口本身为无状态接口，即单次调用时，模型所接收到的信息量仅为接口传入内容，不涉及业务逻辑，同时模型也不存储您传入的数据。
本接口适用场景：语音通话、语音聊天等需求流式的相关场景。
本接口为流式输出接口，支持以下功能：
1. 支持单次合成500字符输入；
2. 支持语调、语速、音量、比特率、采样率相关参数调整特性；
3. 支持音频时长、音频大小等返回参数；
4. 支持混音功能，最多四种音色按照比例混合；
5. 支持非法字符检测：非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；非法字符超过10%，接口不返回结果（返回报错码），请检测后再次进行请求（非法字符定义：ascii码中的控制符（不含制表符和换行符））；
6. 支持语音流式输出，模型合成语音时会根据句子标点判断是否为半句/整句，流式输出时以半句/整句为单位进行输出。目前支持判断的句子标点为：“，”、“；”、“。”、“？”、“……”、“,”、“.”、“?”、“…”。
### API说明
`API：https://api.minimax.chat/v1/tts/stream`
#### 模型列表
| 模型 |
| --- |
| Speech-01 |
#### 接口参数说明
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | 使用发放的值 | 是 | 用户所属的组。该值应拼接在调用API的url末尾。 |
| model | string | 是 | 调用的模型版本：`speech-01` |
| timber_weights | subject | 否 | 音色相关信息。详见下方字段。<br>请求的音色编号。暂时只支持系统音色(id)：参考上面voice_id的说明。<br>权重。最多支持4种音色混合，取值为整数，单一音色取值占比越高，合成音色越像。 |
| timber_weights.voice_id | string | 是 | 请求的音色编号。暂时只支持系统音色(id)：参考上面voice_id的说明。 |
| timber_weights.weight | [1, 100] | 是 | 权重。最多支持4种音色混合，取值为整数，单一音色取值占比越高，合成音色越像。 |
| voice_id | string | 否 | 请求的音色编号。支持系统音色(id)以及复刻音色（id）两种类型，其中系统音色（ID）如下：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| speed | [0.5,2] | 否 | 生成声音的语速。可选，默认值为1.0，取值越大，语速越快。 |
| vol | (0,10] | 否 | 生成声音的音量。可选，默认值为1.0，取值越大，音量越高。 |
| pitch | [-12,12] | 否 | 生成声音的语调。可选，默认值为0（0为原音色输出，取值需为整数）。 |
| text | string | 是 | 待合成的文本。长度限制<500字符，段落切换用换行符替代。 |
| audio_sample_rate | 范围[16000,24000,32000] | 是 | 生成声音的采样率。可选，默认为32000。 |
| bitrate | 范围[32000, 64000，128000] | 是 | 生成声音的比特率。可选，默认值为128000 |
| format | string | 是 | 生成声音的音频格式。默认mp3，范围[mp3,pcm,flac]。 |
| data | object | 否 | data可能返回为null，参考示例代码时，注意进行非空判断。用于在咨询/反馈时帮助定位问题。<br>合成后的音频片段，采用hex编码，按照输入定义的格式进行生成（mp3/pcm/flac）。<br>当前音频流状态，1表示合成中，2表示合成结束。<br>相关额外信息。合成进度，指当前合成文本的字节数。注：请注意合成是以句为单位切割的，若文本只有一句话，则每次返回结果的ced是相同的。 |
| data.audio | string | 否 | 合成后的音频片段，采用hex编码，按照输入定义的格式进行生成（mp3/pcm/flac）。 |
| data.status | int | 否 | 当前音频流状态，1表示合成中，2表示合成结束。 |
| data.ced | string | 否 | 相关额外信息。合成进度，指当前合成文本的字节数。注：请注意合成是以句为单位切割的，若文本只有一句话，则每次返回结果的ced是相同的。 |
| trace_id | string | 否 | 本次会话的id。 |
| extra_info | object | 否 | 相关额外信息，音频全部返回完成后进行返回。<br>音频时长，精确到毫秒。<br>采样率。默认为24000，如客户请求参数进行调整，会根据请求参数生成。<br>音频大小，单位为字节。<br>比特率，默认为168000，如客户请求参数进行调整，会根据请求参数生成。<br>可读字数，已经发音的字数统计（不算标点等其他符号，包含汉字数字字母）。<br>非法字符占比，非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；最大不超过0.1（10%），超过进行报错。<br>计费字符数，本次语音生成的计费字符数。 |
| extra_info.audio_length | int64 | 否 | 音频时长，精确到毫秒。 |
| extra_info.audio_sample_rate | int64 | 否 | 采样率。默认为24000，如客户请求参数进行调整，会根据请求参数生成。 |
| extra_info.audio_size | int64 | 否 | 音频大小，单位为字节。 |
| extra_info.bitrate | int64 | 否 | 比特率，默认为168000，如客户请求参数进行调整，会根据请求参数生成。 |
| extra_info.word_count | int64 | 否 | 可读字数，已经发音的字数统计（不算标点等其他符号，包含汉字数字字母）。 |
| extra_info.invisible_character_ratio | float64 | 否 | 非法字符占比，非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；最大不超过0.1（10%），超过进行报错。 |
| extra_info.usage_characters | int64 | 否 | 计费字符数，本次语音生成的计费字符数。 |
| base_resp |  | 否 | 如果请求出错，对应的错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，服务内部错误及非法字符超过10%；2013，输入格式信息不正常<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，服务内部错误及非法字符超过10%；2013，输入格式信息不正常 |
| base_resp.status_msg | string | 否 | 错误详情。 |
| trace_id | string | 否 | 生成id，用于在咨询/反馈时帮助定位问题。 |
#### 完整请求示例
支持文本流式输入语音流式输出，模型合成语音时会根据句子标点判断是否为半句/整句，流式输出时以半句/整句为单位进行输出。目前支持判断的句子标点为：“，”、“；”、“。”、“？”、“……”、“,”、“.”、“?”、“…”。
具体请求示例如下：
**完整请求示例**
```python
# This Python file uses the following encoding: utf-8

import shutil
import subprocess
from typing import Iterator
import time

import requests
import json
import re

# 创建包含所有中英文标点符号的正则表达式模式
punctuation_pattern = r'.+[\u3002\uff1b\uff0c\uff1a\uff1f！？｡。：；‘’‛“”„‟…‧..,? ]'

group_id='在此输入group_id'
api_key = '在此输入 APIKey'

file_format='mp3'     #支持 mp3/pcm/flac

url = "https://api.minimax.chat/v1/tts/stream?GroupId=" + group_id
payload = {
    "bot_setting": [
        {
            "bot_name": "MM智能助理",
            "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。",
        }
    ],
    "messages": [{"sender_type": "USER", "sender_name": "小明", "text": "生成一篇赞美猞猁的文章,字数超过50"}],
    "reply_constraints": {"sender_type": "BOT", "sender_name": "MM智能助理"},
    "model": "abab5.5-chat",
    "stream": True,
    "tokens_to_generate": 1034,
    "temperature": 0.01,
    "top_p": 0.95,
}
headers = {"Content-Type": "application/json", "Authorization": "Bearer " + api_key}

def build_tts_stream_headers() -> dict:
   headers =   {
    'accept': 'application/json, text/plain, */*',
    'content-type': 'application/json',
    'authorization':  "Bearer "+api_key,
  }
   return headers

def build_tts_stream_body(text: str) -> dict:
  body =  json.dumps({
    "timber_weights": [
      {
        "voice_id": "male-qn-qingse",
        "weight": 1
      },
      {
        "voice_id": "female-shaonv",
        "weight": 1
      }
    ],
    "text": text,
    "voice_id": "",
    "model": "speech-01",
    "speed": 1,
    "vol": 1,
    "pitch": 0,
    "audio_sample_rate": 32000,
    "bitrate": 128000,
    "format": file_format,
  })
  return body

mpv_command = ["mpv", "--no-cache", "--no-terminal", "--", "fd://0"]
mpv_process = subprocess.Popen(
    mpv_command,
    stdin=subprocess.PIPE,
    stdout=subprocess.DEVNULL,
    stderr=subprocess.DEVNULL,
)

def call_tts_stream(text: str) -> Iterator[bytes]:
  tts_url= url
  tts_headers = build_tts_stream_headers()
  tts_body = build_tts_stream_body(text)

  response =  requests.request("POST", tts_url, stream=True, headers=tts_headers, data=tts_body) 
  for chunk in (response.raw):
    if chunk:
        if chunk[:5] == b'data:':
          data = json.loads(chunk[5:])
          if "data" in data and "extra_info" not in data:
            if "audio" in data["data"] :
              audio =data["data"]['audio']
              yield audio


def audio_play(audio_stream: Iterator[bytes]) -> bytes:
  audio = b""
  for chunk in audio_stream:
      if chunk is not None and chunk!= '\n':
          decoded_hex = bytes.fromhex(chunk)
          mpv_process.stdin.write(decoded_hex)  # type: ignore
          mpv_process.stdin.flush()
          audio += decoded_hex

  return audio


def parseChunkDeltaText(chunkStr)->str:
    parsed_data = json.loads(chunkStr[6:])
    if "usage" in parsed_data:
        return  # 当前为流式完结chunk，无增量信息
    delta_content = parsed_data["choices"][0]["messages"]
    text = delta_content[0]["text"]
    return text

def save_audio(audio):
    timestamp = int(time.time())
    file_name=f'output_{timestamp}.{file_format}'
    with open(file_name, 'wb') as file:
      file.write(audio)

def punctual_split_voice(cur_text: str):
  audio = b''

  matched = re.search(punctuation_pattern, cur_text)
  if matched != None:
    sentence = matched.group(0)
    # print("PUNCTUATION SPLIT: ",sentence) #断句后请求语音文本

    audio_chunk_iterator = call_tts_stream(sentence)
    audio = audio_play(audio_chunk_iterator)
    cur_text = cur_text[matched.end():]

  return cur_text, audio


response = requests.post(url, headers=headers, json=payload, stream=True)

res_text = '' #文本全文
cur_text = '' #当前缓冲区文本
res_audio= b''#语音全文

for chunk in response.iter_lines():
    if chunk:
      chunkStr = chunk.decode("utf-8")
      delta_text = parseChunkDeltaText(chunkStr)
      if delta_text != None :
        print(delta_text) #流式文本输出

        cur_text, audio = punctual_split_voice(cur_text+delta_text)
        if audio!=b'':
          save_audio(audio)

        res_audio += audio
        res_text += delta_text

if cur_text!='':
  audio_chunk_iterator = call_tts_stream(cur_text)
  audio = audio_play(audio_chunk_iterator)

  save_audio(audio)
  res_audio += audio


# 结果保存至文件
print(res_text)
timestamp = int(time.time())
file_name=f'output_total_{timestamp}.{file_format}'
with open(file_name, 'wb') as file:
    file.write(res_audio)
```

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/Voice_Cloning_api_intro?key=68adaebe66027263330017e0
- 文档ID: `68adaebe66027263330017e0`
# 音色快速复刻（Voice Cloning）

本接口支持基于用户上传需要复刻音频的音频，以及示例音频，进行音色的复刻。

使用本接口需要完成个人认证及企业认证用户后，方可调用。 请在 [账户管理 -> 账户信息](https://platform.minimaxi.com/user-center/basic-information) 中，完成个人用户认证或企业用户认证，以确保可以正常使用本功能。

本接口适用场景：IP 音色复刻、音色克隆等需要快速复刻某一音色的相关场景。

本接口支持单、双声道复刻声音，支持按照指定音频文件快速复刻相同音色的语音。

## 支持模型
以下为 MiniMax 提供的语音模型及其特性说明。
| 模型 | 特性 |
| ------ | ------ |
| speech-2.5-hd-preview | 最新的 HD 模型，韵律表现出色，复刻相似度极佳 |
| speech-2.5-turbo-preview | 最新的 Turbo 模型，支持 40 个语种 |
| speech-02-hd | 拥有出色的韵律、稳定性和复刻相似度，音质表现突出 |
| speech-02-turbo | 拥有出色的韵律和稳定性，小语种能力加强，性能表现出色 |
| speech-01-hd | 拥有超高的复刻相似度，音质表现突出 |
| speech-01-turbo | 在出色的生成效果基础上有更快的生成速度 |

## 接口说明

1.  **上传待克隆音频**
    调用 [文件上传接口](/document/files_upload?key=68b2aedbfd643cbdfcc0611b) 上传待克隆的音频文件并获取 `file_id`。

2.  **上传示例音频 (可选)**
    若需要提供示例音频以增强克隆效果，需要再次调用 [文件上传接口](/document/files_upload?key=68b2aedbfd643cbdfcc0611b) 上传示例音频文件并获得对应的 `file_id`。填写在`clone_prompt`中的`prompt_audio`中。


3.  **调用复刻接口**
    基于获取的 `file_id` 和自定义的 `voice_id` 作为输入参数，调用 [快速复刻接口](/document/voice_clone?key=68adaf2f6fe587e3fbfe9b89) 克隆音色。
<br>

⚠️ 注意：
- 调用本接口进行音色克隆时，不会立即收取音色复刻费用。音色的复刻费用将在首次使用此复刻音色进行语音合成时收取（不包含本接口内的试听行为）。
- 本接口产出的快速复刻音色为临时音色，若希望永久保留某复刻音色，请于 168 小时（7 天）内在任意 T2A 语音合成接口中调用该音色（不包含本接口内的试听行为）。若超过时限，该音色将被删除。
- 接口采用无状态设计：每次调用仅处理传入数据，且不存储用户上传内容，不涉及任何业务逻辑状态。

<br>

## 官方 MCP
MiniMax 提供官方的 [Python 版本](https://github.com/MiniMax-AI/MiniMax-MCP) 和 [JavaScript 版本](https://github.com/MiniMax-AI/MiniMax-MCP-JS) 模型上下文协议（MCP）服务器实现代码，支持音色克隆功能，详细说明请参考 [MiniMax MCP 使用指南文档](/document/guides_mcp?key=68bbf0f295a600dbfc926bdd)

## 账户相关
- 原始链接: https://platform.minimaxi.com/document/aboutaccount?key=689731f3e92177d0a09a6701
- 文档ID: `689731f3e92177d0a09a6701`
# 账户相关常见问题

### 问：MiniMax开放平台，支持哪些充值方式？

**答：** MiniMax开放平台支持 **在线充值** 和 **对公汇款** 两种充值方式。

- 在线充值：登陆MiniMax开放平台，前往[账户管理->余额](https://platform.minimaxi.com/user-center/payment/balance)，点击 立即充值 按钮，在充值页使用微信进行在线充值。您可以在账单页查询充值结果。
- 对公汇款：对公汇款仅支持企业用户。支持采用对公汇款的方式向MiniMax开放平台账户进行充值。**请确保汇款方开户名称与开放平台实名认证主体名称一致。** 您可联系商务获得具体账户信息。我方银行账户到账后，汇款金额将在 10 分钟- 1 小时左右转入您对应的开放平台账户，如未及时收到，请联系我们。

### 问：如何获取发票

**答：** 在MiniMax开放平台申请开票时，发票的抬头需与实名认证主体名称一致。
请注意，个人账号无法开具企业抬头发票。若需要修改为企业抬头，需前往[账户管理->账户信息 ](https://platform.minimaxi.com/user-center/basic-information)将您的个人认证变更为企业认证。如在认证过程中有任务问题，您可联系商务获得帮助。

您可填写下列表单，进行发票申请：
[MiniMax开放平台发票申请表](https://biaodan.info/web/formview/6503f8ba75a03c4a8388a91a)
请注意：我们并非根据充值金额进行开票。目前的开票模式是：可开票金额=已消耗金额-已开票金额。




### 问：是否支持余额预警

**答：** 为避免因账户余额不足影响您的业务，我们建议您开启余额预警功能。您可前往[账户管理->余额](https://platform.minimaxi.com/user-center/payment/balance)，设置余额预警金额。当账户余额低于预警值时，MiniMax开放平台将以邮件、短信、站内信等形式通知您，请注意查收，并及时充值。

### 问：资源包及有效期

**答：** MiniMax开放平台，目前提供视频及语音资源包的购买，具体套餐内容，您可前往[账户管理->API资源包](https://platform.minimaxi.com/user-center/payment/subscription)了解详细内容，并选择合适您的套餐。

>
><img src="https://filecdn.minimax.chat/public/67f512af-6f23-40c8-9e40-3f58bda1b754.png" alt="图片描述" width="14px" />
>提示：您所购买的资源包，余量不继承，过期将自动清零。请关注所购买的资源包的有效期，及时使用。

## 接口相关
- 原始链接: https://platform.minimaxi.com/document/aboutapis?key=689731e2fc2e4d3e0c688a87
- 文档ID: `689731e2fc2e4d3e0c688a87`
# 接口相关常见问题

### 问：如何获取API Key

**答：** 您可前往[账户管理 > 接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)创建并管理自己的API key。请注意，API Key是您调用接口的重要凭证，请不要与他人共享您的 API 密钥，或将其暴露在浏览器或其他客户端代码中。

### 问：如何获得更高的资源保障

**答：** MiniMax开放平台为您提供不同的资源保障方案，可前往[资源保障](/document/rate_limits?key=68b90c85630660297b0679fc)页面查看具体内容。如您需要获得更高的资源保障，您可通过[api@minimaxi.com](mailto:api@minimaxi.com)与我们的商务获得联系。

### 问：如何才能使用声音复刻服务

**答：** 基于法律法规的要求，如您需要使用声音克隆服务，请先前往[账户管理 > 账户信息](https://platform.minimaxi.com/user-center/basic-information)中的认证信息中，完成**个人实名认证**或者**企业认证**完成认证后，即刻可以通过 [API调试台 > Voice Cloning](https://platform.minimaxi.com/examination-center/voice-experience-center/voiceCloning) 页面，或者通过快速复刻接口使用声音复刻服务。

## 接口
- 原始链接: https://platform.minimaxi.com/document/apis?key=68970eb1e92177d0a09a46d7
- 文档ID: `68970eb1e92177d0a09a46d7`
# 接口功能更新


### 2025年9月23日
- T2A v2接口，支持控制音频恒定比特率编码
- 在 audio_setting 参数下新增二级参数 force_cbr，bool类型，将此参数设置为 `ture`，可控制以恒定比特率（CBR）进行音频编码

### 2025年8月28日
- 视频生成接口，MiniMax-Hailuo-02 首尾帧生成功能上线
- 新增参数“**last_frame_image**” string类型，来控制视频结束帧画面
- 支持768P（6s，10s），1080P（6s）

### 2025年8月2日
- 视频生成接口，MiniMax-Hailuo-02 图生视频功能，支持512分辨率
- 512分辨率下，支持设置6s、10s 生成视频时长
- 新增参数“**fast_pretreatment**” bool 类型，用于控制扩写模型选择，打开情况下能缩短扩写消耗的时长

### 2025年7月19日
- 声音效果器功能上线，通过voice_modify参数，实现音高、音强、音效等调节
- 效果器功能支持同步语音合成接口（含Http, Websocket）、异步长文本语音合成接口

### 2025年7月15日
- 视频Agent API 上线，支持生成模板同款视频
- 11个视频模版，持续更新
- 欢迎访问[视频模板列表](/document/video_agent_templates?key=68b401edd2f0aaaf3484a10e)了解详细内容

### 2025年6月12日
- 音色设计（Voice Design）上线，支持文本描述生成音色
- 原有文生音色接口维持不变，继续提供服务，但不再维护迭代，目前已收至历史接口目录中

### 2025年4月25日
- image-01模型新增**width**、**height**参数，支持用户自定义生成图片宽、高尺寸，满足用户更多生图尺寸需求

## API说明
- 原始链接: https://platform.minimaxi.com/document/assistants document?key=66719091a427f0c8a5701675
- 文档ID: `66719091a427f0c8a5701675`
### 接口文档
此文档为Assistants API文档，通过此文档您可以查看该接口的详细参数以及调用示例。
您也可以通过我们的Assistants 操作指南来更深入的了解该API的使用特性。
### Assistant
您可以在此对象下进行对Assistant的操作，包括创建、检索、删除、查看列表等行为。
#### 创建Assistant
您可以通过指令选择模型创建一个Assistant。
`API：https://api.minimax.chat/v1/assistants/create`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| header | Content-Type | 是 | Content-type |
| GroupId | 使用发放的值 | 是 | 表明用户所属的组，用户身份的唯一标识，该值应拼接在调用API的url末尾 |
| model | string | 是 | 调用的算法模型ID。可选项：`abab5.5s-chat`、`abab5.5-chat`、`abab6.5s-chat`、`abab5.5-chat-240131 (Finetune model)`、`abab5.5s-chat-240123 (Finetune model)` |
| instructions | string | 否 | assistant使用的指令。不同模型所能使用的最大长度不同：`abab5.5s-chat`及其微调模型长度上限为8192个字符、`abab5.5-chat`及其微调模型长度上限为16384个字符、`abab6.5s-chat`长度上限为245760个字符<br>*所有传入文本、返回文本共用此长度上限值。 |
| name | string | 否 | assistant的代称，最大长度为256个字符。 |
| description | string | 否 | assistant的描述。 |
| tools | list | 否 | assistant启用的工具列表，默认[] 列表上限为128个，目前仅abab5.5-chat及其微调模型支持。<br>assistant可使用的工具类型，可选项（仅abab5.5-chat及其微调模型、abab6.5s-chat支持）。工具可以是：`code_interpreter`、`retrieval`、`function`、`web_search`<br>仅当工具类型为function时的补充。参数可选为：`description str`、`name str`、`parameters object` |
| tools.type | string | 否 | assistant可使用的工具类型，可选项（仅abab5.5-chat及其微调模型、abab6.5s-chat支持）。工具可以是：`code_interpreter`、`retrieval`、`function`、`web_search` |
| tools.function | object | 否 | 仅当工具类型为function时的补充。参数可选为：`description str`、`name str`、`parameters object` |
| file_ids | list | 否 | assistant关联的文件列表，默认[] 列表最多上传20个文件id（str），目前仅abab5.5-chat、abab6.5s-chat支持。 |
| role_meta |  | 是 | 对话的meta信息。<br>用户的代称。<br>AI的代称。 |
| role_meta.user_name | string | 是 | 用户的代称。 |
| role_meta.role_meta.bot_name | string | 是 | AI的代称。 |
| t2a_option | string | 否 | T2A可使用的选项，仅在流式中支持可选。<br>T2A选择使用的模型，仅在流式中支持可选，可选模型：`speech-01`<br>生成语音的速度，仅在流式中支持可选，范围在[0.5，2]之间<br>生成声音的音量，仅在流式中支持可选，范围在（0，10]之间<br>默认值为0，生成声音的语调，仅在流式中支持可选，取值范围在[-12，12]之间，（0为原音色输出，取值需为整数）<br>生成声音的采样率，仅在流式中支持可选，取值范围在[16000,24000,32000]之间<br>生成声音的比特率，仅在流式中支持可选，范围在[32000, 64000，128000]之间<br>T2A的音色id，仅在流式中支持可选，支持下述系统音色以及speech-01相应的复刻音色id：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| t2a_option.model | string | 否 | T2A选择使用的模型，仅在流式中支持可选，可选模型：`speech-01` |
| t2a_option.speed | 默认值为1.0 | 否 | 生成语音的速度，仅在流式中支持可选，范围在[0.5，2]之间 |
| t2a_option.vol | 默认值为1.0 | 否 | 生成声音的音量，仅在流式中支持可选，范围在（0，10]之间 |
| t2a_option.pitch | integer | 否 | 默认值为0，生成声音的语调，仅在流式中支持可选，取值范围在[-12，12]之间，（0为原音色输出，取值需为整数） |
| t2a_option.audio_sample_rate | 默认值为32000 | 否 | 生成声音的采样率，仅在流式中支持可选，取值范围在[16000,24000,32000]之间 |
| t2a_option.bitrate | 默认值为128000 | 否 | 生成声音的比特率，仅在流式中支持可选，范围在[32000, 64000，128000]之间 |
| t2a_option.voice_id | string | 否 | T2A的音色id，仅在流式中支持可选，支持下述系统音色以及speech-01相应的复刻音色id：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| id | string | 否 | assistant的ID。 |
| object | string | 否 | 一个assistant对象。 |
| created_at | integer | 否 | 创建助手时的Unix时间戳（以秒为单位） |
| model | string | 否 | 调用的算法模型，可选项如下：`abab5.5s-chat`、`abab5.5-chat`、`abab5.5-chat-240131 (Finetune model)`、`abab5.5s-chat-240123 (Finetune model)` |
| instructions | string | 否 | assistantd的指令，长度上限分别如下：`abab5.5s-chat`及其微调模型:长度上限为8192个字符、`abab5.5-chat`及其微调模型:长度上限为16384个字符、`abab6.5-chat`及其微调模型:长度上限为8192个字符、`abab6.5s-chat`及其微调模型:长度上限为245760个字符<br>*所有传入文本、返回文本共用此长度上限值。 |
| name | string | 否 | assistant的代称，最大长度为256个字符。 |
| description | string | 否 | 对assistant的描述。 |
| tools | list | 否 | Assistant 启用的工具列表，默认[] 列表上限为128个，目前仅`abab5.5-chat`及其微调模型、`abab6.5-chat`、`abab6.5s-chat`支持。<br>工具类型，可选项（`abab5.5-chat`仅及其微调模型、`abab6.5-chat`、`abab6.5s-chat`支持）:code_interpreter、retrieval、function、web_search。<br>仅当工具类型为function时补充，参数如下：description str、name str、parameters object。 |
| tools.type | string | 否 | 工具类型，可选项（`abab5.5-chat`仅及其微调模型、`abab6.5-chat`、`abab6.5s-chat`支持）:code_interpreter、retrieval、function、web_search。 |
| tools.tool.function | object | 否 | 仅当工具类型为function时补充，参数如下：description str、name str、parameters object。 |
| file_ids | string | 否 | assistant关联的文件，默认[] 最多上传20个，目前仅`abab5.5-chat`、`abab6.5-chat`、`abab6.5s-chat`支持。 |
| status | string | 否 | 表示assistant的状态，当参数为available时表示assistant可用。 |
| base_resp | BaseResp | 否 | 表示错误状态码和错误详情。<br>状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>具体错误详情。 |
| base_resp.status_code | integer | 否 | 状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 具体错误详情。 |
#### 检索Assistant
您可以通过此项操作检索已创建的assistant。
`API：https://api.minimax.chat/v1/assistants/retrieve`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| header | application/json | 是 | content-type |
| GroupId | 使用发放的值 | 是 | 用户所属的组，该值应拼接在调用API的url末尾 |
| assistant_id | string | 是 | assistant的ID |
| id | string | 否 | assistant的ID。 |
| object | string | 否 | 一个assistant对象。 |
| created_at | integer | 否 | 创建助手时的Unix时间戳（以秒为单位）。 |
| model | string | 否 | 调用的算法模型，可选项如下：`abab5.5s-chat`、`abab5.5-chat`、`abab6.5s-chat`、`abab5.5-chat-240131 (Finetune model)`、`abab5.5s-chat-240123 (Finetune model)` |
| instructions | string | 否 | assistant的指令。各模型及其长度上限如下：`abab5.5s-chat`及其微调模型：长度上限为8192个字符、`abab5.5-chat`及其微调模型：长度上限为16384个字符、`abab6.5s-chat`：长度上限为245760个字符<br>*所有传入文本、返回文本共用此长度上限值。 |
| name | string | 否 | assistant的代称，最大长度为256个字符。 |
| description | string | 否 | 对assistant的描述。 |
| tools | list | 否 | Assistant 启用的工具列表，默认[] 列表上限为128个，目前仅`abab5.5-chat`及其微调模型、`abab6.5s-chat`支持<br>工具类型，可选项（`abab5.5-chat`及其微调模型、`abab6.5s-chat`支持）<br>仅当工具类型为function时的补充。参数为：description str、name str、parameters object。 |
| tools.type | string | 否 | 工具类型，可选项（`abab5.5-chat`及其微调模型、`abab6.5s-chat`支持） |
| tools.function | object | 否 | 仅当工具类型为function时的补充。参数为：description str、name str、parameters object。 |
| file_ids | string | 否 | assistant关联的文件，默认[] 最多上传20个，目前仅支持`abab5.5-chat`、`abab6.5s-chat`支持 |
| status | string | 否 | 表示assistant的状态，当参数为available时表示assistant可用。 |
| base_resp |  | 否 | 表示错误状态码和错误详情。<br>状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>具体错误详情。 |
| base_resp.status_code | integer | 否 | 状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 具体错误详情。 |
#### 修改Assistant
您可以通过此项操作修改已创建的assistant。
`API：https://api.minimax.chat/v1/assistants/modify [POST]`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| GroupId | string | 是 | 用户的group id，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| id | string | 是 | 助手的ID。 |
| name | string | 是 | 助手名称。 |
| description | string | 否 | 助手的描述。 |
| instructions | string | 是 | 助手的指令。 |
| model | string | 是 | 助手T2T模型。 |
| tools | list[object] | 否 | 工具的列表。 |
| metadata | map[string,string] | 否 | 元数据。 |
| rolemeta | map[string,string] | 否 | 角色名称数据。 |
| file_ids | list[string] | 否 | 文件列表。 |
| t2a_option | object | 否 | T2A的选项。<br>T2A选择使用的模型。<br>T2A的音色。<br>T2A输出音频格式。从以下三种中选取一个{"mp3", "flac", "pcm"}。默认mp3。 |
| t2a_option.model | string | 是 | T2A选择使用的模型。 |
| t2a_option.voice_id | string | 是 | T2A的音色。 |
| t2a_option.format | string | 否 | T2A输出音频格式。从以下三种中选取一个{"mp3", "flac", "pcm"}。默认mp3。 |
| assistant | object | 否 | 助手的信息。<br>助手的id。<br>固定为"assistant"。<br>创建的时间。<br>更新的时间。<br>助手的名称。<br>助手的描述。<br>助手的指令。<br>T2T模型。<br>工具列表。<br>文件列表。<br>元数据。<br>角色名称数据。<br>助手状态。<br>T2A选项。<br>T2A模型。<br>T2A音色。<br>T2A输出音频格式。 |
| assistant.id | string | 否 | 助手的id。 |
| assistant.object | string | 否 | 固定为"assistant"。 |
| assistant.created_at | int64 | 否 | 创建的时间。 |
| assistant.updated_at | int64 | 否 | 更新的时间。 |
| assistant.name | string | 否 | 助手的名称。 |
| assistant.description | string | 否 | 助手的描述。 |
| assistant.instructions | string | 否 | 助手的指令。 |
| assistant.model | string | 否 | T2T模型。 |
| assistant.tools | list[object] | 否 | 工具列表。 |
| assistant.file_ids | list[string] | 否 | 文件列表。 |
| assistant.metadata | map[string,string] | 否 | 元数据。 |
| assistant.rolemeta | map[string,string] | 否 | 角色名称数据。 |
| assistant.status | string | 否 | 助手状态。 |
| assistant.t2a_option | object | 否 | T2A选项。<br>T2A模型。<br>T2A音色。<br>T2A输出音频格式。 |
| assistant.t2a_option.model | string | 否 | T2A模型。 |
| assistant.t2a_option.voice_id | string | 否 | T2A音色。 |
| assistant.t2a_option.format | string | 否 | T2A输出音频格式。 |
| base_resp | object | 否 | 请求结果。<br>状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误信息。 |
| base_resp.status_code | int64 | 否 | 状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误信息。 |
#### 删除Assistant
您可以通过此项操作删除已创建的assistant。
`API：https://api.minimax.chat/v1/assistants/delete`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| Content-Type | application/json | 是 | Content-Type。 |
| GroupId | 使用发放的值 | 是 | 用户所属的组，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| assistant_id | string | 是 | assistant的ID。 |
| id | string | 否 | assistant的ID。 |
| object | string | 否 | 对象。 |
| deleted | bool | 否 | 是否删除。 |
| base_resp | BaseResp | 否 | 状态码以及状态详情。<br>状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 状态码及其分别含义如下：<br>1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
#### 查看Assistant列表
您可以通过此项操作查看已创建的assistant列表。
`API：https://api.minimax.chat/v1/assistants/list`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| Content-Type | application/json | 是 | Content-Type。 |
| GroupId | 使用发放的值 | 是 | 用户所属的组，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| limit | int，默认20 | 否 | 数量。 |
| order | string | 否 | 顺序。可选：<br>1.asc 生序排列；<br>2.desc 降序排列<br>根据assistant的created_at，按顺序返回。 |
| before | string | 否 | 终止指针。用于分页，例如assistant_abcd1234, 则返回包含其在内的上一页。 |
| after | string | 否 | 起始指针。用于分页，例如assistant_abcd1234, 则返回包含其在内的下一页。 |
| object | string | 否 | 对象。 |
| data | []object | 否 | assistant的详情。 |
| has_more | string | 否 | 是否有更多未列出。 |
| first_id | string | 否 | 列表首个assistant的ID。 |
| last_id | string | 否 | 列表最后一个assistant的ID。 |
| base_resp |  | 否 | 错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
### Assistant File
您可以在此对象下进行对assistant中所添加文件的操作，包括创建、检索、列出等行为。
#### 创建assistant文件
您可以通过此项操作将文件添加到assistant来创建assistant文件。
`API：https://api.minimax.chat/v1/assistants/files/create`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| Content-Type | application/json | 是 | Content-Type。 |
| GroupId | 使用发放的值 | 是 | 用户所属的组，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| assistant_id | string | 是 | assistant的ID。 |
| file_id | string | 是 | 文件ID。 |
| id | string | 否 | ID。 |
| object | string | 否 | 对象。 |
| created_at | int | 否 | 创建时间。 |
| assistant_id | string | 否 | assistant的ID。 |
| base_resp |  | 否 | 错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | in64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
#### 检索assistant关联的文件
您可以通过此项操作将文件添加到assistant来检索assistant关联的文件。
`API：https://api.minimax.chat/v1/assistants/files/retrieve`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| Content-Type | application/json | 是 | Content-Type。 |
| GroupId | 使用发放的值 | 是 | 用户所属的组，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| assistant_id | string | 是 | assistant的ID。 |
| file_id | string | 是 | 文件ID。 |
| id | string | 否 | ID。 |
| object | string | 否 | 对象。 |
| created_at | int | 否 | 创建时间。 |
| assistant_id | string | 否 | assistant的ID。 |
| base_resp |  | 否 | 错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | in64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
#### 列出assistant文件
您可以通过此项操作来列出assistant的列表。
`API：https://api.minimax.chat/v1/assistants/files/list`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| Content-Type | application/json | 是 | Content-Type。 |
| GroupId | 使用发放的值 | 是 | 用户所属的组，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| assistant_id | string | 否 | assistant的ID。 |
| limit | int，默认20 | 否 | 数量。 |
| order | string | 否 | 顺序。可选：<br>1.asc 生序排列<br>2.desc 降序排列<br>根据assistant的created_at，按顺序返回。 |
| before | string | 否 | 终止指针，用于分页，例如assistant_abcd1234, 则返回包含其在内的上一页。 |
| after | string | 否 | 起始指针，用于分页，例如assistant_abcd1234, 则返回包含其在内的下一页。 |
| object | string | 否 | 对象。 |
| data | []object | 否 | assistant文件详情。 |
| first_id | string | 否 | 列表首个assistantID。 |
| last_id | string | 否 | 列表最后一个assistantID。 |
| base_resp |  | 否 | 错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | in64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
#### 删除assistant关联的文件
您可以通过此项操作来删除此前关联到assistant的某文件（指定file_id）。
`API：https://api.minimax.chat/v1/assistants/files/delete [POST]`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| GroupId | 使用发放的值 | 是 | 用户所属的组，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| assistant_id | string | 是 | assistant的ID。 |
| file_id | string | 是 | 文件ID。 |
| file_id | object | 否 | 删除的文件id。 |
| base_resp | object | 否 | 请求结果。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | in64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
### Thread
您可以在此对象下进行对thread的操作，包括创建、检索等行为。
#### 创建thread
您可以通过此项操作创建一个thread。
`API：https://api.minimax.chat/v1/threads/create`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| metadata | map<string, string> | 否 | 附带元信息。 |
| id | string | 否 | threadid。 |
| object | string | 否 | 对象：thread。 |
| metadata | string | 否 | thread元信息。 |
| base_resp | BaseResp | 否 | 状态码以及状态详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
#### 检索thread
您可以通过此项操作检索一个已创建的thread。
`API：https://api.minimax.chat/v1/threads/retrieve`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | threadID。 |
| id | string | 否 | threadid。 |
| object | string | 否 | 对象：thread。 |
| metadata | string | 否 | thread元信息。 |
| base_resp | BaseResp | 否 | 状态码以及状态详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
#### 修改thread
您可以通过此项操作修改一个已创建的thread
`API：https://api.minimax.chat/v1/threads/modify`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| GroupId | string | 是 | 用户所属的组，可以在用户中心查看，该值应拼接在调用API的url末尾。 |
| assistant_id | string | 是 | assistant的ID。 |
| file_id | string | 是 | 文件ID。 |
| thread | object | 否 | 线程信息。<br>线程id。<br>固定为"thread"。<br>创建时间。<br>元数据。<br>助手id。<br>线程名称。<br>更新时间。 |
| thread.id | string | 否 | 线程id。 |
| thread.object | string | 否 | 固定为"thread"。 |
| thread.created_at | int64 | 否 | 创建时间。 |
| thread.metadata | int64 | 否 | 元数据。 |
| thread.assistant_id | string | 否 | 助手id。 |
| thread.name | string | 否 | 线程名称。 |
| thread.updated_at | string | 否 | 更新时间。 |
| base_resp | object | 否 | 请求结果。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误信息。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误信息。 |
### Message
您可以在此object下进行对message的操作，包括创建、检索、列出列表等行为，object如下：
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| id | string | 否 | messageid。 |
| object | string | 否 | 对象：message。 |
| created_at | int64 | 否 | 创建时间戳。 |
| thread_id | string | 否 | threadID。 |
| role | string | 否 | 发送者。 |
| content | list<Content> | 否 | message的内容。 |
| file_ids | list<string> | 否 | 对话需要参考的文件。 |
| assistant_id | string | 否 | assistant的ID。 |
| run_id | string | 否 | 运行实体id。 |
| metadata | map<string,string> | 否 | 元信息。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| type | string | 否 | 类型。 |
| text | TextContent | 否 | 文本内容。 |
| image_file | ImageFile | 否 | 图像文件。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| value | string | 否 | message文本内容，会包含t2t文本和语音通话asr的文本内容。 |
| annotations | list<Annotation> | 否 | message引文列表。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| file_id | string | 否 | 文件的ID。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| type | string | 否 | 类型。 |
| text | string | 否 | 引用的文字内容。 |
| start_index | int64 | 否 | 引用文字在message中的起始位置。 |
| end_index | int64 | 否 | 引用文字在message中的终止位置。 |
| file_citation | FileCitation | 否 | assistant message中的引用，当assistant使用检索工具时生成。用于指向与assistant的message相关联的文件片段。 |
| web_citation | WebCitation | 否 | assistant message中的引用，当assistant使用网页搜索时生成。用于指向与assistant的message相关联的网页信息。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| file_id | string | 否 | 文件ID。 |
| quote | string | 否 | 引用。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| url | string | 否 | 网址。 |
| quote | string | 否 | 引用。 |
#### 创建message
您可以通过此项操作创建一个message。
`API：https://api.minimax.chat/v1/threads/messages/add`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | threadID。 |
| role | string | 是 |  |
| content | string | 是 | 内容。 |
| file_id | list<string> | 否 |  |
| metadata | map<string, string> | 否 |  |
#### 检索message
您可以通过此项操作检索一个已创建的message。
`API：https://api.minimax.chat/v1/threads/messages/retrieve`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | threadID。 |
| message_id | string | 是 | messageid。 |
#### message列表
您可以通过此项操作列出已创建的message的列表。
`API：https://api.minimax.chat/v1/threads/messages/list`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | threadID。 |
| limit | int64，默认20 | 否 | 上限。 |
| order | string | 否 | 升序/降序。asc 升序排列；desc 降序排列。 |
| after | string | 否 | 起始指针。 |
| before | string | 否 | 终止指针。 |
| object | string | 否 | 对象。 |
| data | Message object 列表 | 否 | 列出的message列表。 |
| first_id | string | 否 | 起始message id。 |
| last_id | string | 否 | 终止message id。 |
| base_resp | BaseResp | 否 | 请求状态。 |
### Run
您可以在此object下进行对run的操作，包括创建、检索、列出列表、提交function返回等行为，object如下：
- 流式object(包含三类object: Run、Message、Run step)：
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| data | object | 否 | 数据。<br>run id。<br>表示类型,固定为run。<br>创建时间。<br>助手ID。<br>线程id。<br>状态。<br>开始时间。<br>过期时间。<br>取消时间。<br>失败时间。<br>完成时间。<br>错误。<br>错误码。<br>错误原因。<br>模型。<br>t2a配置。<br>指令。<br>工具列表。<br>文件列表。<br>元数据。<br>要求的行为。 |
| data.id | string | 是 | run id。 |
| data.object | string | 是 | 表示类型,固定为run。 |
| data.created_at | int64 | 是 | 创建时间。 |
| data.assistant_id | string | 是 | 助手ID。 |
| data.thread_id | string | 是 | 线程id。 |
| data.status | string | 是 | 状态。 |
| data.started_at | int64 | 否 | 开始时间。 |
| data.expires_at | int64 | 否 | 过期时间。 |
| data.cancelled_at | int64 | 否 | 取消时间。 |
| data.failed_at | int64 | 否 | 失败时间。 |
| data.completed_at | int64 | 否 | 完成时间。 |
| data.last_error | object | 否 | 错误。<br>错误码。<br>错误原因。 |
| data.last_error.code | string | 是 | 错误码。 |
| data.last_error.message | string | 是 | 错误原因。 |
| data.model | string | 否 | 模型。 |
| data.t2a_control | object | 否 | t2a配置。 |
| data.instructions | string | 否 | 指令。 |
| data.tools | list[object] | 否 | 工具列表。 |
| data.file_ids | list[string] | 否 | 文件列表。 |
| data.metadata | map[string]string | 否 | 元数据。 |
| data.required_action | object | 否 | 要求的行为。 |
| base_resp | object | 是 | 请求结果。<br>状态码。有值说明流式结束。状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态信息。 |
| base_resp.status_code | int64 | 否 | 状态码。有值说明流式结束。状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 状态信息。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| data | object | 否 | 数据。<br>message id。<br>表示类型,固定为message。<br>创建时间。<br>线程ID。<br>发送消息的角色。<br>消息的内容。<br>消息内容类型，text、image或audio。<br>消息内容id，通过消息id和消息内容id组织消息。<br>文本消息。<br>文本消息字符串。<br>引文相关数据。<br>图片消息。<br>图片文件id。<br>音频消息。<br>音频文件id，音频文件会在流式结束后生成，也就是说这个流式接口不会返回。<br>音频内容，流式生成，hex编码传输。<br>文件id列表。<br>助手id。<br>run id。<br>元数据。 |
| data.id | string | 是 | message id。 |
| data.object | string | 是 | 表示类型,固定为message。 |
| data.created_at | int64 | 是 | 创建时间。 |
| data.thread_id | string | 是 | 线程ID。 |
| data.role | string | 是 | 发送消息的角色。 |
| data.content | list[object] | 否 | 消息的内容。<br>消息内容类型，text、image或audio。<br>消息内容id，通过消息id和消息内容id组织消息。<br>文本消息。<br>文本消息字符串。<br>引文相关数据。<br>图片消息。<br>图片文件id。<br>音频消息。<br>音频文件id，音频文件会在流式结束后生成，也就是说这个流式接口不会返回。<br>音频内容，流式生成，hex编码传输。 |
| data.content.data.content[0].type | string | 是 | 消息内容类型，text、image或audio。 |
| data.content.data.content[0].id | string | 是 | 消息内容id，通过消息id和消息内容id组织消息。 |
| data.content.data.content[0].text | list[object] | 否 | 文本消息。<br>文本消息字符串。<br>引文相关数据。 |
| data.content.data.content[0].text.data.content[0].text.value | string | 是 | 文本消息字符串。 |
| data.content.data.content[0].text.data.content[0].text.annotations | list[object] | 否 | 引文相关数据。 |
| data.content.data.content[0].image_file | object | 否 | 图片消息。<br>图片文件id。 |
| data.content.data.content[0].image_file.data.content[0].image_file.file_id | string | 是 | 图片文件id。 |
| data.content.data.content[0].audio_file | object | 否 | 音频消息。<br>音频文件id，音频文件会在流式结束后生成，也就是说这个流式接口不会返回。<br>音频内容，流式生成，hex编码传输。 |
| data.content.data.content[0].audio_file.data.content[0].audio_file.file_id | string | 否 | 音频文件id，音频文件会在流式结束后生成，也就是说这个流式接口不会返回。 |
| data.content.data.content[0].audio_file.data.content[0].audio_file.value | string | 是 | 音频内容，流式生成，hex编码传输。 |
| data.file_ids | list[string] | 否 | 文件id列表。 |
| data.assistant_id | string | 否 | 助手id。 |
| data.run_id | string | 否 | run id。 |
| data.metadata | map[string]string | 否 | 元数据。 |
| base_resp | object | 是 | 请求结果。<br>状态码，有值说明流式结束。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误信息。 |
| base_resp.status_code | int64 | 否 | 状态码，有值说明流式结束。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误信息。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| data | object | 否 | 数据。<br>run step id。<br>表示类型,固定为run step。<br>创建时间。<br>run id。<br>助手id。<br>线程ID。<br>run step类型。<br>状态。<br>取消时间。<br>完成时间。<br>过期时间。<br>失败时间。<br>错误。<br>错误码。<br>错误信息。<br>run step内容。 |
| data.id | string | 是 | run step id。 |
| data.object | string | 是 | 表示类型,固定为run step。 |
| data.created_at | int64 | 是 | 创建时间。 |
| data.run_id | string | 是 | run id。 |
| data.assistant_id | string | 是 | 助手id。 |
| data.thread_id | string | 是 | 线程ID。 |
| data.type | string | 是 | run step类型。 |
| data.status | list[object] | 是 | 状态。 |
| data.cancelled_at | int64 | 否 | 取消时间。 |
| data.completed_at | int64 | 否 | 完成时间。 |
| data.expired_at | int64 | 否 | 过期时间。 |
| data.failed_at | int64 | 否 | 失败时间。 |
| data.last_error | object | 否 | 错误。<br>错误码。<br>错误信息。 |
| data.last_error.code | string | 是 | 错误码。 |
| data.last_error.message | string | 是 | 错误信息。 |
| data.step_details | object | 是 | run step内容。 |
| base_resp | object | 是 | 状态码以及状态详情。<br>状态码，有值说明流式结束。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误信息。 |
| base_resp.status_code | int64 | 否 | 状态码，有值说明流式结束。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误信息。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| id | string | 否 | Run id。 |
| object | string | 否 | 对象：run。 |
| created_at | int64 | 否 | 创建时间戳。 |
| assistant_id | string | 否 | 对应assistant的assistant id。 |
| thread_id | string | 否 | thread_id。 |
| status | string | 否 | 状态。 |
| started_at | int64 | 否 | 开始运行时间戳。 |
| expires_at | int64 | 否 | 过期时间戳。 |
| cancelled_at | int64 | 否 | 取消时间戳。 |
| failed_at | int64 | 否 | 失败时间戳。 |
| completed_at | int64 | 否 | 成功时间戳。 |
| last_error | Error | 否 | 错误信息。 |
| model | string | 否 | 模型名称。 |
| instructions | string | 否 | 覆盖assistant的指令集。 |
| tools | Tool列表 | 否 | 可选项：code_interpreter、retrieval、function、web_search<br>参数：description str、name str、parameters object<br>默认[] 最多上传20个。 |
| tools.type | str | 否 | 参数：description str、name str、parameters object |
| tools.function | object | 否 | 默认[] 最多上传20个。 |
| file_ids | list<string> | 否 | 相关的文件列表。 |
| metadata | map<string, string> | 否 | 元信息。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| code | string | 否 | 错误码。 |
| message | string | 否 | 错误信息。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| name | string | 否 | 函数名称。 |
| description | string | 否 | 函数功能描述。 |
| parameters | FunctionParameters | 否 | 该函数的参数及其说明。对每个参数的命名和参数说明也需要描述精确，才能更好的生成。比如： "parameters":{ "type":"object", "properties":{ "meeting_id":{ "type":"string", "description":"会议id" }, "participants":{ "type":"string", "description":"参会人" } }, "required":[ "meetingid", "participants" ] }。 |
#### 创建stream run（流式）
您可以通过此项操作以流式输出运行。
`API：https://api.minimax.chat/v1/threads/run/create_stream`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| GroupId | 使用发放的值 | 是 | 用户的group id，可以在用户中心查看，该值应拼接在调用API的url末尾 |
| stream | int64 | 是 | 流式选项。取值：{1(TEXT_STREAM), 2(TEXT_AND_AUDIO_STREAM)}。<br>在文本和音频流式下（TEXT_AND_AUDIO_STREAM），输入的消息类型可以为音频（AUDIO），同时会有音频输出。 |
| thread_id | string | 是 | 线程id。 |
| assistant_id | string | 是 | 助手id。 |
| messages | list[object] | 是 | 需要添加的消息。<br>消息类型。取值：{1(TEXT), 2(AUDIO)}。<br>当消息类型为2(AUDIO)时，消息内容请填写Hex编码的音频内容，音频目前只支持mp3格式.<br>消息所属的角色。取值：{"user", "ai","assistant"}<br>消息的内容。<br>文件id列表。<br>消息元数据。 |
| messages.type | int64 | 是 | 消息类型。取值：{1(TEXT), 2(AUDIO)}。<br>当消息类型为2(AUDIO)时，消息内容请填写Hex编码的音频内容，音频目前只支持mp3格式. |
| messages.role | string | 是 | 消息所属的角色。取值：{"user", "ai","assistant"} |
| messages.content | string | 是 | 消息的内容。 |
| messages.file_ids | list[string] | 否 | 文件id列表。 |
| messages.metadata | map[string]string | 是 | 消息元数据。 |
| model | string | 否 | 模型。不填写使用创建助手时传入的模型。 |
| t2a_option | object | 否 | T2A选项。当2(TEXT_AND_AUDIO_STREAM)流式模式下，必传。<br>T2A模型。可选项：`speech-01`<br>生成语音的速度，仅在流式中支持可选，取值范围[0.5，2]，默认为1.0。<br>生成声音的音量，仅在流式中支持可选，范围在（0，10]之间，默认为1.0。<br>默认值为0，生成声音的语调，仅在流式中支持可选，取值范围在[-12，12]之间，（0为原音色输出，取值需为整数）。<br>生成声音的采样率，仅在流式中支持可选，取值支持[16000,24000,32000]，默认值为32000。<br>生成声音的比特率，仅在流式中支持可选，取值支持[32000, 64000, 128000]，默认值为128000。<br>T2A音色id。可选项（支持下述系统音色以及speech-01相应的复刻音色id）：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| t2a_option.model | string | 否 | T2A模型。可选项：`speech-01` |
| t2a_option.speed | double | 否 | 生成语音的速度，仅在流式中支持可选，取值范围[0.5，2]，默认为1.0。 |
| t2a_option.vol | double | 否 | 生成声音的音量，仅在流式中支持可选，范围在（0，10]之间，默认为1.0。 |
| t2a_option.pitch | integer | 否 | 默认值为0，生成声音的语调，仅在流式中支持可选，取值范围在[-12，12]之间，（0为原音色输出，取值需为整数）。 |
| t2a_option.audio_sample_rate | integer | 否 | 生成声音的采样率，仅在流式中支持可选，取值支持[16000,24000,32000]，默认值为32000。 |
| t2a_option.bitrate | integer | 否 | 生成声音的比特率，仅在流式中支持可选，取值支持[32000, 64000, 128000]，默认值为128000。 |
| t2a_option.voice_id | string | 否 | T2A音色id。可选项（支持下述系统音色以及speech-01相应的复刻音色id）：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin |
| instructions | string | 否 | 指令。 |
| tools | list[object] | 否 | 工具。 |
| metadata | dict | 否 | 元数据。 |
| truncation_strategy | object | 否 | 控制当前run如何截断线程中的上下文。可用于控制run的初始上下文窗口中的message数量。<br>要使用的截断策略。目前仅支持设置为last_messages，上下文窗口将被设置为thread中最近的 n 条message。<br>为当前run构建上下文时，取用thread中最新的message的数量。 |
| truncation_strategy.type | string | 是 | 要使用的截断策略。目前仅支持设置为last_messages，上下文窗口将被设置为thread中最近的 n 条message。 |
| truncation_strategy.last_messages | integer | 是 | 为当前run构建上下文时，取用thread中最新的message的数量。 |
|  |  | 否 | 同流式run object 结构。 |
本接口流式返回数据，按照以下节奏返回：
![image](https://filecdn.minimax.chat/public/32243ce0-98ee-4d0e-8572-d6fd0be70e56.png)
#### 创建run（非流式）
您可以通过此项操作创建一个Run。
`API：https://api.minimax.chat/v1/threads/run/create`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | threadid。 |
| assistant_id | string | 是 | assistantid。 |
| model | string | 否 | 模型名称。 |
| instructions | string | 否 | 指令集。 |
| tools | Tool 列表 | 否 | 启用的工具列表。 |
| metadata | dict | 否 | 元信息。 |
| truncation_strategy | object  | 否 | 控制当前run如何截断线程中的上下文。可用于控制run的初始上下文窗口中的message数量。<br>要使用的截断策略。目前仅支持设置为last_messages，上下文窗口将被设置为thread中最近的 n 条message。<br>为当前run构建上下文时，取用thread中最新的message的数量。 |
| truncation_strategy.type | string | 是 | 要使用的截断策略。目前仅支持设置为last_messages，上下文窗口将被设置为thread中最近的 n 条message。 |
| truncation_strategy.last_messages | integer | 是 | 为当前run构建上下文时，取用thread中最新的message的数量。 |
#### 检索run
您可以通过此项操作检索已创建的Run。
`API：https://api.minimax.chat/v1/threads/run/retrieve`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | threadID。 |
| run_id | string | 是 | runID。 |
#### 列表run
您可以通过此项操作列出已创建的Run的列表。
`API：https://api.minimax.chat/v1/threads/run/list`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | threadID。 |
| limit | int64 | 否 | 数量限制。 |
| order | string | 否 | 升序or降序。 |
| after | string | 否 | 起始指针。 |
| before | string | 否 | 终止指针。 |
| object | string | 否 | 对象：Run。 |
| data | Run object 列表 | 否 | 列出的Run结构列表。 |
| base_resp | BaseResp | 否 | 状态码+状态详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
#### 修改run
您可以通过此项操作修改已创建的run。
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| GroupId | string | 是 | 用户的group id，可以在用户中心查看。该值拼接在调用API的url末尾。 |
| thread_id | string | 是 | 助手id。 |
| run_id | string | 是 | run id。 |
| metadata | dict | 否 | 元数据。 |
| run | object | 否 | run信息。<br>run id。<br>固定为"thread.run"。<br>助手id。<br>线程id。<br>run 状态。<br>开始时间。<br>过期时间。<br>取消时间。<br>失败时间。<br>完成时间。<br>错误信息。<br>T2T模型。<br>T2A选项。<br>指令。<br>工具列表。<br>文件列表。<br>元数据。<br>更新时间。 |
| run.id | string | 是 | run id。 |
| run.object | string | 是 | 固定为"thread.run"。 |
| run.assistant_id | int64 | 是 | 助手id。 |
| run.thread_id | int64 | 否 | 线程id。 |
| run.status | string | 是 | run 状态。 |
| run.started_at | string | 否 | 开始时间。 |
| run.expires_at | string | 否 | 过期时间。 |
| run.cancelled_at | int64 | 否 | 取消时间。 |
| run.failed_at | int64 | 否 | 失败时间。 |
| run.completed_at | int64 | 否 | 完成时间。 |
| run.last_error | map[string]string | 否 | 错误信息。 |
| run.model | string | 是 | T2T模型。 |
| run.t2a_option | object | 否 | T2A选项。 |
| run.instructions | string | 否 | 指令。 |
| run.tools | list[object] | 否 | 工具列表。 |
| run.file_ids | list[string] | 否 | 文件列表。 |
| run.metadata | 元数据 | 否 | 元数据。 |
| run.updated_at | 更新时间 | 是 | 更新时间。 |
| base_resp | object | 是 | 状态码以及状态详情。<br>状态码，非0表示出现错误。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误信息。 |
| base_resp.status_code | int64 | 否 | 状态码，非0表示出现错误。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误信息。 |
#### 取消run
您可以通过此项操作取消已创建的run。
`API：https://api.minimax.chat/v1/threads/run/cancel [POST]`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| GroupId | string | 是 | 用户的group id，可以在用户中心查看。该值拼接在调用API的url末尾。 |
| thread_id | string | 是 | 助手id。 |
| run_id | string | 是 | run id。 |
| run | object | 否 | run信息。<br>run id。<br>固定为"thread.run"。<br>助手id。<br>线程id。<br>run 状态。<br>开始时间。<br>过期时间。<br>取消时间。<br>失败时间。<br>完成时间。<br>错误信息。<br>T2T模型。<br>T2A选项。<br>指令。<br>工具列表。<br>文件列表。<br>元数据。<br>更新时间。 |
| run.id | string | 否 | run id。 |
| run.object | string | 否 | 固定为"thread.run"。 |
| run.assistant_id | int64 | 否 | 助手id。 |
| run.thread_id | int64 | 否 | 线程id。 |
| run.status | string | 否 | run 状态。 |
| run.started_at | string | 否 | 开始时间。 |
| run.expires_at | string | 否 | 过期时间。 |
| run.cancelled_at | int64 | 否 | 取消时间。 |
| run.failed_at | int64 | 否 | 失败时间。 |
| run.completed_at | int64 | 否 | 完成时间。 |
| run.last_error | map[string]string | 否 | 错误信息。 |
| run.model | string | 否 | T2T模型。 |
| run.t2a_option | object | 否 | T2A选项。 |
| run.instructions | string | 否 | 指令。 |
| run.tools | list[object] | 否 | 工具列表。 |
| run.file_ids | list[string] | 否 | 文件列表。 |
| run.metadata | dict | 否 | 元数据。 |
| run.updated_at | int64 | 否 | 更新时间。 |
| base_resp | object | 否 | 状态码和状态详情。<br>状态码，非0表示出现错误。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误信息。 |
| base_resp.status_code | int64 | 否 | 状态码，非0表示出现错误。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误信息。 |
#### Submit tool outputs to run
您可以通过此项操作将tool outputs提交到run。
`API：https://api.minimax.chat/v1/threads/run/submit_tool_outputs`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | 线程id。 |
| run_id | string | 是 | run id。 |
| tool_outputs | list[ToolOutput] | 是 | 工具输出列表。<br>工具 id。<br>输出。 |
| tool_outputs.tool_call_id | string | 否 | 工具 id。 |
| tool_outputs.output | string | 是 | 输出。 |
| id | string | 否 | run id。 |
| object | string | 否 | 对象：run。 |
| created_at | int64 | 否 | 创建时间戳。 |
| assistant_id | string | 否 | 对应助手的assistantid。 |
| thread_id | string | 否 | threadid。 |
| status | string | 否 | 状态。 |
| started_at | int64 | 否 | 开始运行时间戳。 |
| expired_at | int64 | 否 | 过期时间戳。 |
| cancelled_at | int64 | 否 | 取消时间戳。 |
| failed_at | int64 | 否 | 失败时间戳。 |
| completed_at | int64 | 否 | 成功时间戳。 |
| last_error | string | 否 | 错误信息。 |
| model | string | 否 | 模型名称。 |
| instructions | string | 否 | 覆盖assistiant的指令集。 |
| tools | Tool列表 | 否 | 启用的工具列表。 |
| file_ids | list[string] | 否 | 相关的文件列表。 |
| metadata | dict | 否 | 元信息。 |
| base_resp | BaseResp | 否 | 状态码以及状态详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
### Run Step
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| id | string | 否 | Run中的step id。 |
| object | string | 否 | 对象：run step。 |
| created_at | int64 | 否 | 创建时间戳。 |
| run_id | string | 否 | run id。 |
| assistant_id | string | 否 | assistant的ID。 |
| thread_id | string | 否 | thread的id。 |
| type | string | 否 | 类型。 |
| status | string | 否 | 状态。 |
| cancelled_at | int64 | 否 | 取消时间戳。 |
| completed_at | int64 | 否 | 完成时间戳。 |
| expired_at | int64 | 否 | 过期时间戳。 |
| failed_at | int64 | 否 | 失败时间戳。 |
| last_error | map[string,string] | 否 | 错误信息。 |
| step_details | StepDetails | 否 | 步骤详情。<br>取值message_creation或者tool_calls。<br>创建message的详细信息。<br>当前Step所创建message的ID。<br>工具调用的详细信息。<br>包含：code_interpreter、web_search、retrieval、function。<br>当工具调用为code_interpreter时显示。<br>tool ID。<br>类型。<br>Code interpreter的输入。<br>Code interpreter的输出。<br>当工具调用为web_search时显示。<br>tool ID。<br>类型。<br>搜索问题。<br>搜索结果。<br>当工具调用为retrieval时显示。<br>tool ID。<br>类型。<br>检索问题。<br>检索结果。<br>当工具调用为function时显示。<br>tool ID。<br>类型。<br>函数名称。<br>函数名称。<br>函数参数。<br>函数结果。 |
| step_details.type | string | 否 | 取值message_creation或者tool_calls。 |
| step_details.message_creation | object | 否 | 创建message的详细信息。<br>当前Step所创建message的ID。 |
| step_details.message_creation.message_id | string | 否 | 当前Step所创建message的ID。 |
| step_details.tool_calls | list | 否 | 工具调用的详细信息。<br>包含：code_interpreter、web_search、retrieval、function。<br>当工具调用为code_interpreter时显示。<br>tool ID。<br>类型。<br>Code interpreter的输入。<br>Code interpreter的输出。<br>当工具调用为web_search时显示。<br>tool ID。<br>类型。<br>搜索问题。<br>搜索结果。<br>当工具调用为retrieval时显示。<br>tool ID。<br>类型。<br>检索问题。<br>检索结果。<br>当工具调用为function时显示。<br>tool ID。<br>类型。<br>函数名称。<br>函数名称。<br>函数参数。<br>函数结果。 |
| step_details.tool_calls.type | dict | 否 | 包含：code_interpreter、web_search、retrieval、function。 |
| step_details.tool_calls.code_interpreter | dict | 否 | 当工具调用为code_interpreter时显示。<br>tool ID。<br>类型。<br>Code interpreter的输入。<br>Code interpreter的输出。 |
| step_details.tool_calls.code_interpreter.id | string | 否 | tool ID。 |
| step_details.tool_calls.code_interpreter.type | string | 否 | 类型。 |
| step_details.tool_calls.code_interpreter.code_interpreter | object | 否 | Code interpreter的输入。<br>Code interpreter的输出。 |
| step_details.tool_calls.code_interpreter.code_interpreter.code_interpreter.input | string | 否 | Code interpreter的输入。 |
| step_details.tool_calls.code_interpreter.code_interpreter.code_interpreter.outputs | object | 否 | Code interpreter的输出。 |
| step_details.tool_calls.web_search | dict | 否 | 当工具调用为web_search时显示。<br>tool ID。<br>类型。<br>搜索问题。<br>搜索结果。 |
| step_details.tool_calls.web_search.id | string | 否 | tool ID。 |
| step_details.tool_calls.web_search.type | string | 否 | 类型。 |
| step_details.tool_calls.web_search.web_search | object | 否 | 搜索问题。<br>搜索结果。 |
| step_details.tool_calls.web_search.web_search.query | string | 否 | 搜索问题。 |
| step_details.tool_calls.web_search.web_search.outputs | string | 否 | 搜索结果。 |
| step_details.tool_calls.retrieval | dict | 否 | 当工具调用为retrieval时显示。<br>tool ID。<br>类型。<br>检索问题。<br>检索结果。 |
| step_details.tool_calls.retrieval.id | string | 否 | tool ID。 |
| step_details.tool_calls.retrieval.type | string | 否 | 类型。 |
| step_details.tool_calls.retrieval.retrieval | object | 否 | 检索问题。<br>检索结果。 |
| step_details.tool_calls.retrieval.retrieval.query | string | 否 | 检索问题。 |
| step_details.tool_calls.retrieval.retrieval.outputs | string | 否 | 检索结果。 |
| step_details.tool_calls.function | dict | 否 | 当工具调用为function时显示。<br>tool ID。<br>类型。<br>函数名称。<br>函数名称。<br>函数参数。<br>函数结果。 |
| step_details.tool_calls.function.id | string | 否 | tool ID。 |
| step_details.tool_calls.function.type | string | 否 | 类型。 |
| step_details.tool_calls.function.function | object | 否 | 函数名称。<br>函数名称。<br>函数参数。<br>函数结果。 |
| step_details.tool_calls.function.function.name | string | 否 | 函数名称。 |
| step_details.tool_calls.function.function.arguments | string | 否 | 函数参数。 |
| step_details.tool_calls.function.function.output | string | 否 | 函数结果。 |
#### 检索run step
您可以通过此项操作检索已运行的Run的步骤，以便您检查运行的状态。
`API：https://api.minimax.chat/v1/threads/run_steps/retrieve`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| thread_id | string | 是 | 线程id。 |
| run_id | string | 是 | run id。 |
| step_id | string | 是 | 步骤id。 |
| id | string | 否 | run步骤id。 |
| object | string | 否 | 对象：run step。 |
| created_at | int64 | 否 | 创建时间戳。 |
| run.id | string | 否 | run id。 |
| assistant_id | string | 否 | 对应助手的assistantid。 |
| thread_id | string | 否 | threadid。 |
| type | string | 否 | 类型。 |
| status | string | 否 | 状态。 |
| cancelled_at | int64 | 否 | 取消时间戳。 |
| completed_at | int64 | 否 | 完成时间戳。 |
| expired_at | int64 | 否 | 过期时间戳。 |
| failed_at | int64 | 否 | 失败时间戳。 |
| last_error | map<string,string> | 否 | 错误信息。 |
| step_details | StepDetails | 否 | 步骤详情。 |
| base_resp | BaseResp | 否 | 状态码以及状态详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_code | string | 否 | 状态详情。 |
#### 列表run step
您可以通过此项操作列出已创建的Run的运行步骤的列表。

## 删除声音
- 原始链接: https://platform.minimaxi.com/document/delete_voice?key=68adb5ab6602726333001d22
- 文档ID: `68adb5ab6602726333001d22`
# 删除声音（Delete Voice）
该 API 用于删除指定 voice_id 的音色，删除范围为通过 [Voice Cloning API](https://platform.minimaxi.com/document/%E5%BF%AB%E9%80%9F%E5%85%8B%E9%9A%86?key=66719032a427f0c8a570165b) 和 [Voice Generation API](https://platform.minimaxi.com/document/voice_design?key=669f5af198ff2c57eeb9a0f0) 生成的音色的 voice_id。
⚠️ 注意：删除后，该 voice_id 将无法再次使用。

<request method="POST" url="https://api.minimaxi.com/v1/delete_voice"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>Required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看。

<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br>

### 请求体
<hr/>

#### voice_type    <gray>string</gray> <red>Required</red>
取值范围："voice_cloning"（克隆的音色）/"voice_generation"（基于文本提示生成的音色），注意仅支持删除这两种类别的音色
<hr/>

#### voice_id    <gray>string</gray> <red>Required</red>
希望删除音色的 voice_id

### 返回体
<hr/>

#### voice_id    <gray>string</gray>
被删除声音的 voice_id
<hr/>

#### created_time    <gray>string</gray>
该音色生成请求提交的时间，非首次调用生效激活时间，格式为 "yyyy-mm-dd"
<hr/>

#### base_resp    <gray>dict</gray>
状态码及状态详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码。
- 0: 删除成功；
- 2013: 输入非法

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
""""
本示例用于删除音色。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import json
import os

api_key = os.environ.get("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/delete_voice"

payload = json.dumps({
  "voice_type": "voice_cloning",
  "voice_id": "yanshang11123"
})
headers = {
  'content-type': 'application/json',
  'authorization': f'Bearer {api_key}'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/delete_voice' \
--header 'content-type: application/json' \
--header "Authorization: Bearer ${MINIMAX_API_KEY}" \
--data '{
    "voice_type":"voice_cloning",
    "voice_id":"yanshang11123"
}'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "voice_id": "yanshang11123",
    "created_time": "1728962464",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

## Embeddings
- 原始链接: https://platform.minimaxi.com/document/embeddings?key=66718fbfa427f0c8a5701627
- 文档ID: `66718fbfa427f0c8a5701627`
### Embeddings（向量化）
embeddings接口支持用户输入文本数据，使用MiniMax自研的Embedding model得到文本对应的浮点向量。
![image](https://filecdn.minimax.chat/public/044175b4-3bd3-4791-a994-42730ad797af.png)
通过embedding技术能够将文本转化为高维向量表示，捕捉词语和句子之间的语义关系，从而实现语义相似度计算、文本分类和信息检索等任务。相对于将文本输入直接给模型回答，可以在此之前通过embedding技术在指定的知识库中检索到有关联的上下文，放到模型的输入中，让模型基于上下文做出更流畅、合理的回答。基于此可以实现如**长记忆检索**、**知识库检索**等能力。
#### 使用场景
在如下场景，可以使用embeddings接口：
1. **知识库检索**：用户有特有的知识库（如MiniMax训练数据所不具备的实时性信息或者业务专属信息）文本，并且ChatCompletion接口对于这些特有知识库回复效果不理想时，可以将与输入消息内容比较相关的知识内容注入ChatCompletion接口prompt信息中获得更好的回复效果。
2. **长记忆检索**：即用户输入的prompt信息token数大于接口token上限需要使用embedding进行信息简化时，可以将与输入消息内容比较相关的知识内容抽取出来注入ChatCompletion接口prompt信息中获得更好的回复效果。
embeddings接口在这些场景下主要帮助用户更好的得到与输入消息内容比较相关的知识内容：
- 对于用户特有的知识库内容进行向量化；
- 对于输入的消息内容进行向量化。
#### 使用效果
通过效果示例展示使用效果，在效果示例中可以发现，加入相关上下文信息后，模型可以给出更准确的答复。基本请求体如下，{prompt}和{query}替换为后续实际内容。
**基础请求体**
```python
payload = {
    "model": "abab5-chat",
    "prompt": {prompt},
    "role_meta": {
        "user_name": "智能助理",
        "bot_name": "用户"
    },
    "messages": [
        {
            "sender_type": "USER",
            "text": {query}
        }
    ]
}
```
##### 示例1
**Tab: 不使用embedding关联上下文**

**请求示例**
```json
rompt = '你是MiniMax自主研发的大型语言模型，回答问题简洁有条理。'

query = '23赛季英超冠军是谁？'
```

**Tab: 使用embedding检索填入相关上下文**

**请求示例，信息来自维基百科**
```json
prompt = '你是MiniMax自主研发的大型语言模型，回答问题简洁有条理。请根据以下相关内容来回答问题。如果你不知道答案，就说你不知道，不要试图编造答案。
----
相关的文本：
2023 年 5 月 21 日第三十七周赛事，阿森纳作客 0 - 1 不敌诺丁汉森林，在联赛剩余一轮下落后榜首曼城 4 分，使曼城确定提前三轮卫冕英超冠军。不但如此，诺丁汉森林亦提前一轮保级成功，本赛季成为英超史上第四次三支升级马球队全数保级成功的赛季。
----'
query = '23赛季英超冠军是谁？'
```
**Tab: 不使用embedding关联上下文**

**返回示例**
```json
{"created":1684819940,"model":"abab5-chat","reply":"我无法确定具体的赛季，因为英超比赛时间跨度比较大。请提供更多信息，以便我为您提供准确的答案。","choices":[{"text":"我无法确定具体的赛季，因为英超比赛时间跨度比较大。请提供更多信息，以便我为您提供准确的答案。","index":0,"logprobes":0,"finish_reason":"stop"}],"usage":{"total_tokens":107},"input_sensitive":false,"output_sensitive":false,"id":"20dfe16cede4cc2e04363dd940243608","base_resp":{"status_code":0,"status_msg":"success"}}
```

**Tab: 使用embedding检索填入相关上下文**

**返回示例，信息来自维基百科**
```json
{"created":1684819814,"model":"abab5-chat","reply":"曼城在2023赛季英超联赛中提前三轮卫冕冠军。","choices":[{"text":"曼城在2023赛季英超联赛中提前三轮卫冕冠军。","index":0,"logprobes":0,"finish_reason":"stop"}],"usage":{"total_tokens":290},"input_sensitive":false,"output_sensitive":false,"id":"2fe581c0645eec0a12a148dc9912890b","base_resp":{"status_code":0,"status_msg":"success"}}
```
##### 示例2
**Tab: 不使用embedding关联上下文**

**请求示例**
```json
prompt = '你是MiniMax自主研发的大型语言模型，回答问题简洁有条理。'
query = '调用api，返回1002是什么原因？'
```

**Tab: 使用embedding检索填入相关上下文**

**请求示例，信息来自minimax api文档**
```json
prompt = '你是MiniMax自主研发的大型语言模型，回答问题简洁有条理。请根据以下相关内容来回答问题。如果你不知道答案，就说你不知道，不要试图编造答案。
----
状态码(StatusCode)代表服务的返回状态，0为正确返回，所有的非零值均代表一类报错，具体语义见下表:
1000 未知错误；
1001 超时错误；
1002 触发限流策略；
1004 鉴权失败；
1008 账户余额不足；
1013 服务内部错误；
1027 输出内容严重违规；
2013 输入格式信息不正常
----
调用频率限制指的是基于商业策略的考量，对调用的频率进行限制的策略。
目前限制策略主要有以下维度：
基于请求次数的频率限制，按照基准时间不同，分为每秒限制请求次数（RPS）、每分钟限制请求次数（RPM）
基于token数的频率限制，为每分钟限制token总量（TPM）
----'
query = '调用api，返回1002是什么原因？'
```
**Tab: 不使用embedding关联上下文**

**返回示例**
```json
{"created":1684819646,"model":"abab5-chat","reply":"根据您提供的信息，我无法确定具体的问题。但是，如果您正在使用API或SDK，并且得到了1002错误代码，那么这可能表示以下几种情况之一：\n\n1. API或SDK不存在或无法使用。\n2. API或SDK的访问权限或密钥已过期或不正确。\n3. 您的请求超出了API或SDK的最大请求限制。\n4. API或SDK的服务器出现了问题。\n\n建议您查看API或SDK的官方文档或联系API或SDK的开发者以获取更多信息。","choices":[{"text":"根据您提供的信息，我无法确定具体的问题。但是，如果您正在使用API或SDK，并且得到了1002错误代码，那么这可能表示以下几种情况之一：\n\n1. API或SDK不存在或无法使用。\n2. API或SDK的访问权限或密钥已过期或不正确。\n3. 您的请求超出了API或SDK的最大请求限制。\n4. API或SDK的服务器出现了问题。\n\n建议您查看API或SDK的官方文档或联系API或SDK的开发者以获取更多信息。","index":0,"logprobes":0,"finish_reason":"stop"}],"usage":{"total_tokens":259},"input_sensitive":false,"output_sensitive":false,"id":"de1d80cfa1a2867f2da10c0f504a1102","base_resp":{"status_code":0,"status_msg":"success"}}
```

**Tab: 使用embedding检索填入相关上下文**

**返回示例，信息来自minimax api文档**
```json
{"created":1684814440,"model":"abab5-chat","reply":"当API接口返回1002错误码时，通常是因为在调用API的过程中触发了限流策略。限流策略是一种保护措施，旨在防止请求流量超过服务的处理能力，从而防止系统崩溃或过度负载。\n\n具体来说，当服务端检测到请求流量超过设定的限流阈值时，就会返回1002错误码，告知客户端当前请求已被限制。客户端可以在响应中获取当前的限流策略，根据策略调整请求频率或采取其他措施来避免触发限流。\n\n在实际应用中，限流策略可以根据不同的场景和需求进行配置和调整。例如，在秒杀活动中，为了防止恶意用户使用多个账户刷单，可以对请求频率进行限制，以保证活动的公平性和安全性。","choices":[{"text":"当API接口返回1002错误码时，通常是因为在调用API的过程中触发了限流策略。限流策略是一种保护措施，旨在防止请求流量超过服务的处理能力，从而防止系统崩溃或过度负载。\n\n具体来说，当服务端检测到请求流量超过设定的限流阈值时，就会返回1002错误码，告知客户端当前请求已被限制。客户端可以在响应中获取当前的限流策略，根据策略调整请求频率或采取其他措施来避免触发限流。\n\n在实际应用中，限流策略可以根据不同的场景和需求进行配置和调整。例如，在秒杀活动中，为了防止恶意用户使用多个账户刷单，可以对请求频率进行限制，以保证活动的公平性和安全性。","index":0,"logprobes":0,"finish_reason":"stop"}],"usage":{"total_tokens":717},"input_sensitive":false,"output_sensitive":true,"id":"8d6c510483af02ef79bd805b2db63cab","base_resp":{"status_code":0,"status_msg":"success"}}
```
#### 接口限制
embeddings接口仅提供文本数据到浮点向量的转换能力，暂不提供向量数据库等基础设施能力。
#### 解决方案
当用户处于上述使用场景时，我们推荐可以使用如下解决方案，主要包含三个主要部分：
1. 第一部分是构建后续检索用的向量库：
- 需要选择合适的数据作为向量数据库的来源，比如在客服场景中，可以使用用户文档和历史客服记录作为数据源，生成相应的向量表示。类似的，如果您希望模型能够回答企业内部知识问题，可以将相关文档/文件/以及其他相关文本，转换为向量存储起来；
- 用户将选择好的特有知识库基于自己业务逻辑进行分块，然后使用embeddings接口，将分块知识库文本向量化得到分块知识库文本以及对应的浮点向量，并将其存储到用户自己搭建的向量数据库中。
2. 第二部分是根据用户的输入查询相关向量：
- 当收到输入消息内容需要回复时，首先对输入消息内容使用embeddings接口得到对应的浮点向量；
- 用输入消息内容对应的浮点向量与向量数据库中存储的分块知识库浮点向量进行相似度计算，相似度计算推荐使用余弦相似度，取TOPK分块知识库文本内容（不是浮点向量）；
- 第三部分是将相关信息输入prompt发起请求得到回复。
3. 将输入消息内容和TOPK分块知识库文本内容一并注入到ChatCompletion接口中得到回复，建议将TOPK分块知识库文本内容放入prompt信息中。
结合embeddings接口使用大语言模型的基本流程如下：
![image](https://filecdn.minimax.chat/public/45bfec93-d6bf-458b-9ecf-9b85bc9d56d9.png)
### Embeddings API说明
`API：https://api.minimax.chat/v1/embeddings`
#### 模型列表
| 模型 | 最大token数 |
| --- | --- |
| embo-01 | 4096 |
#### 接口参数说明
###### 请求体（Request）参数
###### 返回(Response)参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | 使用发放的值 | 是 | 用户所属的组。使用发放的值。该值应拼接在调用API的url末尾。 |
| model | string | 是 | 请求的模型版本。 |
| texts | []string、array | 是 | 期望生成向量的文本。长度上限为4096个token。 |
| type | string | 是 | 生成向量后的目标使用场景。使用embeddings时，首先通过db生成目标内容的向量并存储到向量数据库中，之后通过query生成检索文本的向量。<br>注：偏算法的参数，我们采用了query和db分离的算法方案，因此一段文本，如果是要作为被检索文本，应该使用db，如果是作为检索文本，应该使用query。<br>二选一：<br>db：用于构建向量库，生成向量存储到库中（作为被检索文本）。<br>query：用于生成查询用的向量（作为检索文本时）。 |
| vectors | [][]float32 | 否 | 向量结果。一个文本对应一个float32数组，长度为1536。 |
| total_tokens | int64 | 否 | 消耗的token数量。 |
| base_resp |  | 否 | 如果请求出错，对应的错误状态码和详情。<br>状态码。1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，非法字符超过10%；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发限流；1004，鉴权失败；1013，非法字符超过10%；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
#### 解决方案代码示例
以下是针对上述解决方案的python代码示例，该示例提供了“从文本中提取embedding向量”、“计算两个embedding向量相似度”代码参考、以及使用embeddings知识库信息进行ChatCompletion接口访问和不使用的效果对比，可以让您在自己的业务场景中更快更方便的将embeddings接口使用起来。
**从文本中提取embedding向量、计算两个embedding向量相似度**
```python
import numpy as np
import requests
import json

group_id = "请填入您的groupid"
api_key = "请填入您的api_key"

# 从文本中提取embedding
def get_embedding(text, emb_type):
    url = f"https://api.minimax.chat/v1/embeddings?GroupId={group_id}"
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }

    data = {
        "texts": [
            text
        ],
        "model": "embo-01",
        "type": emb_type
    }

    response = requests.post(url, headers=headers, data=json.dumps(data))
    res = json.loads(response.text)['vectors'][0]
    return res

# 计算两个embedding的相似度
def embedding_similarity(a, b):
    return np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b))

class Memory():
    def __init__(self, emb_func=get_embedding, sim_func=embedding_similarity):
        self.data = []
        self.emb_func = emb_func
        self.sim_func = sim_func
    #将分块text以及分块text形成的向量形成一个字典做关联存放在内存中，注意：当在实际应用时，应将数据存放到向量数据库中         
    def save_memory(self, text):
        embedding = self.emb_func(text, emb_type='db')
        
        self.data.append({
            "text": text,
            "emb": embedding
        })
    #根据query找出topk的分块text出来    
    def retrive(self, query, topk=2):
        query_emb = self.emb_func(query, emb_type='query')
        memory = sorted(self.data, key=lambda x: self.sim_func(x['emb'], query_emb), reverse=True)[:topk]
        texts = [m['text'] for m in memory]
        texts = [''] + texts + ['']
        return '\n----\n'.join(texts)

#不使用embedding进行chatcompletion接口访问               
def simple_chat(query):
    url = f"https://api.minimax.chat/v1/text/chatcompletion?GroupId={group_id}"
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    payload = {
        "model": "abab5-chat",
        "prompt": "你是MiniMax自主研发的大型语言模型，回答问题简洁有条理。",
        "role_meta": {
            "user_name": "用户",
            "bot_name": "智能助理"
        },
        "messages": [
            {
                "sender_type": "USER",
                "text": query
            }
        ]
    }

    response = requests.post(url, headers=headers, json=payload)
    return response.text

#使用embedding进行chatcompletion接口访问       
def embedding_chat(query):
    # 构建向量库
    texts = [
        'Minimax的文本embedding可用于将离散的符号（例如单词、字符或词组）映射到低维向量空间中的实数向量。这些向量的空间位置被设计为与它们语义上的关联度相关，从而可以通过进行向量的计算来比较自然语言的相似度，用于衡量文本字符串的相关性（两个向量之间的距离衡量它们的相关性，小距离表示高相关性，大距离表示低相关性）。',
        
        '由于目前模型上下文长度会锁定在特定长度（当前是4096token），因此在使用全量信息进行提问的场景下，局限性会很大。而基于embedding技术，我们可以获得两段文本之间的相似度/关联性，基于此可以实现如长记忆检索、知识库检索等能力。',
        
        '接口支持基于自然语言交互生成回复的能力。接口本身为无状态接口，即单次调用时，模型所接收到的信息量仅为接口传入内容，不涉及业务逻辑，同时模型也不存储您传入的数据。如果输入内容或输出内容严重违规，接口会返回内容违规错误信息，回复内容为空。',
        
        '针对abab5，我们为您设定了默认智能助理背景，该设定已完成绝大多数生产力场景能力调试，建议无特殊使用目的的使用者直接使用该背景设定。如果您希望直接使用默认设定，请在调用接口时不传入prompt和role_meta。',
        
        '基于MiniMax端到端自研多模态大语言模型，我们为企业用户或企业开发者提供功能丰富的API，适用于大多数文本处理的场景，以自然语言交互的形式帮助企业用户或企业开发者提高文本相关的生产效率，例如不同行业的文本续写、文案生成、文本扩写、文本改写、内容摘要、代码生成、知识检索等。',
        
        '同时我们的新一代语音合成引擎是一项创新性的技术，它依托于新一代AI大模型能力，能够理解人类语言中的复杂含义，如情感，语气，甚至笑声，从而从文本中预测出情绪、语调等信息，生成接近于人类的超自然、高保真、个性化的语音。同时，依托大语言模型架构，模型能在几GB的参数中学习到数千个声音的音色特征，从而能够在不到30秒的音频中实现音色克隆，生成与提供的音频音色特征高度相似的语音。我们的语音合成引擎能够广泛应用于电子书、游戏、客服、政府等各种场景，为用户提供更加智能、高效的语音服务。',
        
        '在大语言模型中，token是指自然语言文本中的最小粒度单位，也就是一个最小的单词或符号。通常情况下，自然语言文本是由一个一个的token组成的，每个token都具备自己的词性、词义等属性。',
        
        '在训练大语言模型时，会使用一种称为“词向量”的技术，将每个token表示成一个向量形式，这个向量可以包含很多有用的语义信息，比如单词的词性、上下文等。模型通过这些向量来学习和理解自然语言文本，并能够完成各种任务。',
        
        '在大语言模型处理任务的过程中，输入的文本会被转译为token输入到模型中，而输出则是从token转译到文本。输入token与输出token的长度直接影响了大语言模型所消耗的算力，所以业界通常采用基于token数量的计费模式。',
        
        '调用频率限制指的是基于商业策略的考量，对调用的频率进行限制的策略。',
        
        '目前限制策略主要有以下维度：基于请求次数的频率限制，按照基准时间不同，分为每秒限制请求次数（RPS）、每分钟限制请求次数（RPM）基于token数的频率限制，为每分钟限制token总量（TPM）',
        
        'minimax收费模式以每1000个token（包含输入+输出）为基础计价单位，1000个token约对应750个汉字文本（包括标点），以下是不同接口和模型及其对应价格:ChatCompletion-abab5：0.015元/千token；ChatCompletion pro-abab5.5:0.015元/千token；web search：0.03元/次调用'
    ]
    mem = Memory()
    for text in texts:
        mem.save_memory(text)
        
    # 构建回复
    context = mem.retrive(query)
    prompt = f"""使用根据以下内容来回答问题。 如果你不知道答案，就说你不知道，不要试图编造答案。
{context}
"""
    
    print(prompt)
    url = f"https://api.minimax.chat/v1/text/chatcompletion?GroupId={group_id}"
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    payload = {
        "model": "abab5-chat",
        "prompt": prompt,
        "role_meta": {
            "user_name": "用户",
            "bot_name": "智能助理"
        },
        "messages": [
            {
                "sender_type": "USER",
                "text": query
            }
        ]
    }

    response = requests.post(url, headers=headers, json=payload)
    return response.text

if __name__ == '__main__':
    query = 'ChatCompletion-abab5模型怎么计费？'
    #不使用embedding进行chatcompletion接口访问并输出回复
    print("------------------simple sample-----------------------")
    res = simple_chat(query)
    print("simple_chat result:\n", res)
    
    #使用embedding进行chatcompletion接口访问并输出回复
    print("------------------embedding sample-----------------------")
    emb_res = embedding_chat(query)
    print("embedding_chat result:\n", emb_res)
```

## 错误码查询
- 原始链接: https://platform.minimaxi.com/document/error_code?key=68ad64216fe587e3fbfe7acf
- 文档ID: `68ad64216fe587e3fbfe7acf`
### 错误码查询
| 错误码 | 含义 | 解决方法 |
| --- | --- | --- |
| 1000 | 未知错误/系统默认错误 | 请稍后再试 |
| 1001 | 请求超时 | 请稍后再试 |
| 1002 | 请求频率超限 | 请稍后再试 |
| 1004 | 未授权/Token不匹配/Cookie缺失 | 请检查API Key |
| 1008 | 余额不足 | 请检查您的账户余额 |
| 1024 | 内部错误 | 请稍后再试 |
| 1026 | 输入内容涉敏 | 请调整输入内容 |
| 1027 | 输出内容涉敏 | 请调整输入内容 |
| 1033 | 系统错误/下游服务错误 | 请稍后再试 |
| 1039 | Token限制 | 请调整max_tokens |
| 1041 | 连接数限制 | 请联系我们 |
| 1042 | 不可见字符比例超限/非法字符超过10% | 请检查输入内容，是否包含不可见字符或非法字符 |
| 1043 | ASR相似度检查失败 | 请检查file_id与text_validation匹配度 |
| 1044 | 克隆提示词相似度检查失败 | 请检查克隆提示音频和提示词 |
| 2013 | 参数错误 | 请检查请求参数 |
| 20132 | 语音克隆样本或voice_id参数错误 | 请检查Voice Cloning 接口下的 file_id 和 T2A v2，T2A Large v2 接口下的 voice_id 参数 |
| 2037 | 语音时长不符合要求(太长或太短) | 请检查voice_clone file_id文件时长，最少应不低于10秒，最长应不超过5分钟 |
| 2038 | 用户语音克隆功能被禁用 | 使用语音克隆功能需要完成账户身份认证，请根据您的使用需求在账户系管理》账户信息中进行个人或企业认证 |
| 2039 | 语音克隆voice_id重复 | 请修改voice_id，确保未和已有voice_id重复 |
| 2042 | 无权访问该voice_id | 请确认是否为该voice_id创建者 |
| 2045 | 请求频率增长超限 | 请避免请求骤增骤减情况 |
| 2048 | 语音克隆提示音频太长 | 请调整prompt_audio音频文件时长（＜8s） |
| 2049 | 无效的API Key | 请检查API Key |
如有其他疑问，请联系我们（api@minimaxi.com)。

## 接口能力说明
- 原始链接: https://platform.minimaxi.com/document/files_api_intro?key=68b2aeb5d95d92466168a3a6
- 文档ID: `68b2aeb5d95d92466168a3a6`
# 文件管理 (File)

本接口是作为文件管理接口，配合MiniMax开放平台的其他接口使用。

## 接口说明

本接口是作为文件管理接口，配合其他接口使用。共包含5个接口：**上传**、**列出**、**检索**、**下载**、**删除**。


## 文件支持格式

| 类型 | 格式 |
|---|---|
| 文档 | `pdf`、`docx`、`txt`、`jsonl` |
| 音频 | `mp3`、`m4a`、`wav` |


## 容量及限制

| 限制内容 | 限制大小 |
| :--- | :--- |
| 总容量 | 100GB |
| 单个文档容量 | 512MB |

## 删除
- 原始链接: https://platform.minimaxi.com/document/files_delete?key=68b56afdd2f0aaaf3484d365
- 文档ID: `68b56afdd2f0aaaf3484d365`
# 文件删除


<request method="POST" url="https://api.minimaxi.com/v1/files/delete"></request>

<container>

<left>


### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 application/json，确保请求数据的格式为 JSON


<br/>

### 请求体
<hr/>


#### file_id    <gray>int</gray> <red>required</red>
文件的唯一标识符

<hr/>

#### purpose    <gray>string</gray> <red>required</red>
文件使用目的。取值及支持格式如下：
1. voice_clone
1. prompt_audio
1. t2a_async
1. t2a_async_input
1. voice_isolation
1. video_generation


<br/>

### 返回体
<hr/>

#### file_id    <gray>int64</gray>
文件的唯一标识符
<hr/>

#### base_resp    <gray>object</gray>
状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码如下：
- 1000, 未知错误
- 1001, 超时
- 1002, 触发RPM限流
- 1004, 鉴权失败
- 1008, 余额不足
- 1013, 服务内部错误
- 1026, 输入内容错误
- 1027, 输出内容错误
- 1039, 触发TPM限流
- 2013, 输入格式信息不正常

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>


</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
import requests
import json

# 请填写你的api key
api_key = "请填写你的api key"
# 请填写你的file id
file_id = 请填写file id

url = "https://api.minimaxi.com/v1/files/delete"
headers = {
    "Authorization": f"Bearer {api_key}",
    "Content-Type": "application/json"
}

payload = {
    "file_id": file_id,
    "purpose": "video_generation"
}

response = requests.post(url, headers=headers, data=json.dumps(payload))
print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/delete' \
--header 'authority: api.minimaxi.com' \
--header 'content-type: application/json' \
--header 'Authorization: Bearer {api_key}' \
--data '{
    "file_id": 你的文件_id,
    "purpose": "video_generation"
}'
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "file_id": ${file_id},
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

## 列出
- 原始链接: https://platform.minimaxi.com/document/files_list?key=68b56a7ed2f0aaaf3484d1a3
- 文档ID: `68b56a7ed2f0aaaf3484d1a3`
# 文件列出


<request method="GET" url="https://api.minimaxi.com/v1/files/list"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 application/json，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>



#### purpose    <gray>string</gray> <red>required</red>
列出文件分类。取值及支持格式如下：
1. **voice_clone**: 快速复刻原始文件
1.  **prompt_audio**: 音色复刻的示例音频
1.  **t2a_async**: 异步长文本语音生成合成中音频

<br/>

### 返回体
<hr/>

#### file_id    <gray>int64</gray>
文件的唯一标识符
<hr/>

#### bytes    <gray>int64</gray>
文件大小，以字节为单位
<hr/>

#### created_at    <small>int64</small>
创建文件时的 Unix 时间戳，以秒为单位
<hr/>

#### filename    <gray>string</gray>
文件的名称
<hr/>

#### purpose    <gray>string</gray>
文件的使用目的
<hr/>

#### base_resp    <gray>object</gray>
状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码如下：
- 1000, 未知错误
- 1001, 超时
- 1002, 触发RPM限流
- 1004, 鉴权失败
- 1008, 余额不足
- 1013, 服务内部错误
- 1026, 输入内容错误
- 1027, 输出内容错误
- 1039, 触发TPM限流
- 2013, 输入格式信息不正常

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
import requests

# 请填写您的api key
api_key = "请填写您的api key"
purpose = "t2a_async"
url = f"https://api.minimaxi.com/v1/files/list?purpose={purpose}"
headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {api_key}"
}

response = requests.get(url, headers=headers)
print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/list?purpose=t2a_async' \
--header 'Authorization: Bearer ${MINIMAX_API_KEY}'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "files": [
        {
            "file_id": "${file_id}",
            "bytes": 5896337,
            "created_at": 1699964873,
            "filename": "MiniMax开放平台-测试bot.text",
            "purpose": "t2a_async"
        },
        {
            "file_id": "${file_id}",
            "bytes": 5896337,
            "created_at": 1700469398,
            "filename": "297990555456911.tar",
            "purpose": "t2a_async"
        },
        {
            "file_id": "${file_id}",
            "bytes": 1970566,
            "created_at": 1700469460,
            "filename": "297990555456011.tar",
            "purpose": "t2a_async"
        }
    ],
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

## 检索
- 原始链接: https://platform.minimaxi.com/document/files_retrieve?key=68b56abca96516e26019203a
- 文档ID: `68b56abca96516e26019203a`
# 文件检索
<request method="GET" url="https://api.minimaxi.com/v1/files/retrieve"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 application/json，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>

#### file_id    <gray>int64</gray> <red>required</red>
文件的唯一标识符
本接口支持：
- 视频生成中，查询视频任务状态接口获得的`file_id`
- 异步语音合成中，查询语音生成任务状态接口获得的`file_id`

<br/>

### 返回体
<hr/>

#### file_id    <gray>int64</gray>
文件的唯一标识符
<hr/>

#### bytes    <gray>int64</gray>
文件大小，以字节为单位
<hr/>

#### created_at    <gray>int64</gray>
创建文件时的 Unix 时间戳，以秒为单位
<hr/>

#### filename    <gray>string</gray>
文件的名称
<hr/>

#### purpose    <gray>string</gray>
文件的使用目的
<hr/>

#### download_url    <gray>string</gray>
文件下载的url地址
<hr/>


#### base_resp    <gray>object</gray>
状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码如下：
- 1000, 未知错误
- 1001, 超时
- 1002, 触发RPM限流
- 1004, 鉴权失败
- 1008, 余额不足
- 1013, 服务内部错误
- 1026, 输入内容错误
- 1027, 输出内容错误
- 1039, 触发TPM限流
- 2013, 输入格式信息不正常

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
import requests


# 请填写你的api key
api_key = "请填写你的api key"
# 请填写你的file id
file_id = 请填写你的file id

url = f"https://api.minimaxi.com/v1/files/retrieve?file_id={file_id}"
headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {api_key}"
}

response = requests.get(url, headers=headers)
print(response.text)

</codetab>
<codetab language="curl">
curl --location  --request GET 'https://api.minimaxi.com/v1/files/retrieve?file_id=${file_id}' \
--header 'Authorization: Bearer ${MINIMAX_API_KEY}' \
--header 'content-type: application/json' \
--data-raw ''
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "file": {
        "file_id": ${file_id},
        "bytes": 5896337,
        "created_at": 1700469398,
        "filename": "output_aigc.mp4",
        "purpose": "video_generation",
        "download_url": "www.downloadurl.com"
    },
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

## 下载
- 原始链接: https://platform.minimaxi.com/document/files_retrieve_content?key=68b56ae2a96516e2601920d2
- 文档ID: `68b56ae2a96516e2601920d2`
# 文件下载
<request method="GET" url="https://api.minimaxi.com/v1/files/retrieve_content"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 application/json，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>

#### file_id    <gray>string</gray> <red>required</red>
需要下载的文件ID


<br>

### 返回体
<hr/>
该接口将以文件的形式返回信息

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
import requests

api_key = "请填写你的api key"
file_id = "请填写file id"

url = f"https://api.minimaxi.com/v1/files/retrieve_content?file_id={file_id}"

payload = {}
headers = {
  'content-type': 'application/json',
  'Authorization': f'Bearer {api_key}'
}

response = requests.request("GET", url, headers=headers, data=payload)


</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/retrieve_content?file_id={file_id}' \
--header 'content-type: application/json' \
--header 'Authorization: Bearer {api_key}'

</codetab>
</codetabs>

</right>

</container>

## 上传
- 原始链接: https://platform.minimaxi.com/document/files_upload?key=68b2aedbfd643cbdfcc0611b
- 文档ID: `68b2aedbfd643cbdfcc0611b`
# 文件上传


<request method="POST" url="https://api.minimaxi.com/v1/files/upload"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看。
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 application/json，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>


#### purpose    <gray>string</gray> <red>required</red>
文件使用目的。取值及支持格式如下：
1. **voice_clone**: 快速复刻原始文件，（支持mp3、m4a、wav格式）
1.  **prompt_audio**: 音色复刻的示例音频，（支持mp3、m4a、wav格式）
1.  **t2a_async_input**: 异步长文本语音生成合成中，请求体中的文本文件（支持text、zip格式）
<hr/>

#### file <red>required</red>
需要上传的文件。填写文件的路径地址
<br/>

### 返回体
<hr/>

#### file_id    <gray>int64</gray>
文件的唯一标识符
<hr/>

#### bytes    <gray>int64</gray>
文件大小，以字节为单位
<hr/>

#### created_at    <gray>int64</gray>
创建文件时的 Unix 时间戳，以秒为单位
<hr/>

#### filename    <gray>string</gray>
文件的名称
<hr/>

#### purpose    <gray>string</gray>
文件的使用目的
<hr/>

#### base_resp    <gray>object</gray>
状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0，请求成功

1002，触发限流，请稍后再试

1004，账号鉴权失败，请检查 API-Key 是否填写正确

1008，账号余额不足

1026，图片描述涉及敏感内容

2013，传入参数异常，请检查入参是否按要求填写

2049，无效的api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
import requests

# 请填写您的api key
api_key = "请填写您的api key"
url = "https://api.minimaxi.com/v1/files/upload"

headers = {
    "Authorization": f"Bearer {api_key}"
}

data = {
    "purpose": "t2a_async_input"
}

files = {
    "file": open("/Users/minimax/Downloads/export_file (1). text", "rb")
}

response = requests.post(url, headers=headers, data=data, files=files)
print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/upload' \
--header 'Authorization: Bearer ${MINIMAX_API_KEY}' \
--form 'purpose="t2a_async_input"' \
--form 'file=@"/path/to/your/file.text"'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "file": {
        "file_id": "${file_id}",
        "bytes": 5896337,
        "created_at": 1700469398,
        "filename": "合成素材.text",
        "purpose": "t2a_async_input"
    },
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

## Finetune
- 原始链接: https://platform.minimaxi.com/document/finetune?key=66718f91a427f0c8a570160d
- 文档ID: `66718f91a427f0c8a570160d`
### 操作指南
仔细阅读本文档，可以了解如何为您的场景微调文本模型。
#### 何时需要使用微调？
微调文本模型可以使它们更适合特定应用场景，但这需要仔细投入时间和精力。我们建议首先尝试使用【使用说明】里的操作来获得比较好的结果，主要原因如下：
1. 在许多任务中，模型最初可能表现不佳，但可以通过人设编写优化以及要求限制编写优化（提示工程）来改进结果，因此可能不需要微调；
2. 调优指南里的策略比微调迭代要快得多，微调需要整理数据、创建数据集并运行训练作业；
3. 在仍然需要微调的情况下，初始提示工程工作不会浪费，在微调的数据中使用良好的提示进行训练后，通常会看到更好的结果。
一般来说，在以下场景使用微调能够改善模型效果：
1. 让模型学会某一特定风格、基调、格式，比如让模型学会模仿一个人设，或者是一种特殊的输出格式；
2. 提高产生所需输出的可靠性，比如让模型更熟悉某一特定专业垂直领域；
3. 更正失败以遵循复杂的提示，比如在让模型遵从一些非常难以理解的指令要求；
4. 以特定方式处理许多边缘情况，比如让模型做一些非通用场景的回答。
#### 操作步骤及接口
##### 准备数据集
一旦决定使用微调（您已经尽可能地优化了提示，并确定了模型仍然存在的问题，希望使用微调来解决这些问题），首先需要准备用于训练模型的数据。
创建一组多样化的示例对话数据，这些对话数据代表了对模型输出效果的期望。
###### 内容要求
贴近推理时的真实数据。
- 建议在微调之前采用您认为最适合模型的一组对话示例和提示，以获得最好的结果，特别是当训练数据量相对较少的时候（例如不到100个）。
- 可能需要更多的训练数据才能获得良好的结果，因为模型必须完全通过数据示例来学习。
###### 格式要求
参考最终推理使用的接口是Chatcompleion还是Chatcompleion pro，可以使用对应接口的数据格式进行微调。
多条数据，用换行来分割。
**Tab: Chatcompletion Pro**

**Chatcompletion Pro**
```shell
{  "bot_setting":[{"bot_name":"文本总结助手","content":"文本总结助手能够理解文本，并以CSV格式系统地列出标题，作者，发行年份和类型。"}],"messages":[{"sender_type":"USER","sender_name":"用户","text":"1967年，哥伦比亚著名作家加夫列尔·加西亚·马尔克斯(Gabriel García Márquez)创作了《百年孤独》(Cien años de soledad)，这是一部魔幻现实主义小说，被誉为拉丁美洲文学的巅峰之作。这本书讲述了布恩迪亚家族的七代人的故事，展现了马尔克斯独特的叙事风格和奇幻的想象力。作为文学史上的经典之一，这本小说融合了现实和超现实元素，探讨了爱、家庭、孤独和时间等主题。它以其深刻的内涵和引人入胜的叙事而广受欢迎，被翻译成多种语言并屡次改编成电影、舞台剧等形式，深刻影响\n炼金术士, 保罗·科埃略, 1988, 寓言小说了世界文学。"},{"sender_type":"BOT","sender_name":"文本总结助手","text":"标题, 作者, 发行年份, 类型\n百年孤独, 加夫列尔·加西亚·马尔克斯, 1967年, 魔幻现实主义小说"}]}
  {  "bot_setting":[{"bot_name":"文本总结助手","content":"文本总结助手能够理解文本，并以CSV格式系统地列出标题，作者，发行年份和类型。"}],"messages":  [{"sender_type":"USER","sender_name":"用户","text":"2005年，英国著名作家J.K.罗琳(J.K. Rowling)创作了《哈利·波特与混血王子》(Harry Potter and the Half-Blood Prince)，这是一部魔幻冒险小说，以其引人入胜的情节和丰富的魔法世界而闻名于世。这本书是著名的哈利·波特系列的第六部，讲述了年轻巫师哈利·波特在霍格沃茨魔法学校的第六学年，他与导师邓布利多共同努力，揭示了伏地魔的秘密过去和弱点。罗琳的精湛叙事技巧和令人难以置信的想象力使这个魔法世界栩栩如生，吸引了读者们的注意。《哈利·波特与混血王子》不仅是一部奇幻小说，还是关于友情、忠诚和成长的故事，深受读者的喜爱。"},{"sender_type":"BOT","sender_name":"文本总结助手","text":"标题, 作者, 发行年份, 类型\n哈利·波特与混血王子, J.K.罗琳, 2005年, 魔幻冒险小说"}]}
```

**Tab: Chatcompletion**

**Chatcompleion**
```shell
{"messages":[{"sender_type":"USER", "text":"1967年，哥伦比亚著名作家加夫列尔·加西亚·马尔克斯(Gabriel García Márquez)创作了《百年孤独》(Cien años de soledad)，这是一部魔幻现实主义小说，被誉为拉丁美洲文学的巅峰之作。这本书讲述了布恩迪亚家族的七代人的故事，展现了马尔克斯独特的叙事风格和奇幻的想象力。作为文学史上的经典之一，这本小说融合了现实和超现实元素，探讨了爱、家庭、孤独和时间等主题。它以其深刻的内涵和引人入胜的叙事而广受欢迎，被翻译成多种语言并屡次改编成电影、舞台剧等形式，深刻影响\n炼金术士, 保罗·科埃略, 1988, 寓言小说了世界文学。"}, {"sender_type":"BOT", "text":"标题, 作者, 发行年份, 类型\n百年孤独, 加夫列尔·加西亚·马尔克斯, 1967年, 魔幻现实主义小说"}], "prompt":"文本总结助手能够理解文本，并以CSV格式系统地列出标题，作者，发行年份和类型。", "role_meta":{"bot_name":"文本总结助手", "user_name":"用户"}}
 {"messages":[{"sender_type":"USER", "text":"2005年，英国著名作家J.K.罗琳(J.K. Rowling)创作了《哈利·波特与混血王子》(Harry Potter and the Half-Blood Prince)，这是一部魔幻冒险小说，以其引人入胜的情节和丰富的魔法世界而闻名于世。这本书是著名的哈利·波特系列的第六部，讲述了年轻巫师哈利·波特在霍格沃茨魔法学校的第六学年，他与导师邓布利多共同努力，揭示了伏地魔的秘密过去和弱点。罗琳的精湛叙事技巧和令人难以置信的想象力使这个魔法世界栩栩如生，吸引了读者们的注意。《哈利·波特与混血王子》不仅是一部奇幻小说，还是关于友情、忠诚和成长的故事，深受读者的 喜爱。"}, {"sender_type":"BOT", "text":"标题, 作者, 发行年份, 类型\n哈利·波特与混血王子, J.K.罗琳, 2005年, 魔幻冒险小说"}], "prompt":"文本总结助手能够理解文本，并以CSV格式系统地列出标题，作者，发行年份和类型。", "role_meta":{"bot_name":"文本总结助手", "user_name":"用户"}}
```
###### 数量要求
- 需要提供至少 10 个示例。通常 50 到 100 个训练样本进行微调后能看到效果的改进，但这会根据数据质量而有很大差异。
- 我们建议先准备50条精心制作的数据进行训练，看看模型在微调后是否有改进的迹象。如果没有改进，可能要考虑重新为模型设置微调任务或重构数据。
###### 类型要求
需要拆分为训练集和验证集两种类型。
将数据集拆分为训练集和验证集文件。当提交包含训练和测试文件的微调作业时，我们将在训练过程中提供两者的统计信息。这些统计数据将是模型改进程度的初始信号。
###### 长度要求
| 模型列表 | 每条训练数据的最大token长度 |
| --- | --- |
| abab5.5-chat-240119 | 16384 |
| abab5.5s-chat-240123 | 8192 |
##### 上传训练文件
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| file | string | 是 | 数据文件. |
| purpose | string | 是 | 文件的目的。<br>fine-tune：finetune的训练文件和验证文件<br>fine-tune-result：finetune后的结果文件（无需自己上传） |
##### 创建微调任务
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| model | string | 是 | 模型名称。 |
| training_file | int64 | 是 | 包含训练数据的已上传文件的 ID。数据集的格式必须为 JSONL 文件。此外，必须上传文件目的为fine-tune的。 |
| validation_file | int64 | 否 | 包含验证数据的已上传文件的 ID。 |
| hyperparameters | hyperparameters | 否 | 训练超参，不填写会使用默认值。<br>指定全局批量大小，越大意味着模型参数的更新频率越低，默认值为1，建议填写范围：1~100。<br>学习率，较小的学习率可能有助于避免过度拟合，默认值为1e-4，范围：1e-2 ~ 1e-6，注意：一般情况下不建议修改该参数。<br>模型训练过程中的整个数据集的遍历次数，默认值为5，建议填写范围：1~100，注意：这个数值的大小会直接影响您的训练token消耗。 |
| hyperparameters.batch_size | int64 | 否 | 指定全局批量大小，越大意味着模型参数的更新频率越低，默认值为1，建议填写范围：1~100。 |
| hyperparameters.learning_rate_multiplier | double | 否 | 学习率，较小的学习率可能有助于避免过度拟合，默认值为1e-4，范围：1e-2 ~ 1e-6，注意：一般情况下不建议修改该参数。 |
| hyperparameters.n_epochs | int64 | 否 | 模型训练过程中的整个数据集的遍历次数，默认值为5，建议填写范围：1~100，注意：这个数值的大小会直接影响您的训练token消耗。 |
| suffix | string | 否 | 微调模型名称的自定义后缀，最长8位长度，超长后会默认截断。 |
###### 数量限制
一个groupid最多保留5个Finetune model，因此如果本身已经超过了限制，则无法创建微调任务，可以通过[Finetune API文档](https://platform.minimaxi.com/document/guides/Finetune/api?id=657c2bd224432923d556b45d)中的删除微调模型来对模型数量进行管理
请注意，Finetune mode 如果在7天内无任何调用，则会失效。
###### 训练时间
请注意，在开始微调任务前需要进行排队等待资源，开始微调任务后，可能需要一些时间才能完成。您的作业可能排在我们系统中的其他任务之后，训练模型可能需要几分钟或几小时，具体取决于数据集大小。
##### 检索微调任务
**检索微调任务请求示例**
```curl
curl --location 'https://api.minimax.chat/v1/retrieve_finetune_job?GroupId=${group_id}' \
--header 'authorization: Bearer $MiniMax_API_KEY' \
--header 'content-type: application/json' \
--data '{"fine_tuning_job_id":"2367267362783"}'
```

**python**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/retrieve_finetune_job?GroupId={GroupId}"

payload = json.dumps({
  "fine_tuning_job_id":"2367267362783"
})
headers = {
  'authorization': f'Bearer {API_KEY}',
  'content-type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```
**检索微调任务返回示例**
```python
{
    "finetune_job":{
        "id":"1701538478753114",
        "created_at":1702558487515,
        "fine_tuned_model":"ft:abab5.5s-chat-240123_1705584973436:dafei12",
        "hyperparameters":{
            "batch_size":1,
            "learning_rate_multiplier":0.23,
            "n_epochs":1
        },
        "model":"ft:abab5.5s-chat-240123_1705584973436:dafei12",
        "object":"finetune.job",
        "organization_id":"",
        "result_files":[
            86036815667264
        ],
        "status":"queued",
        "training_file":86036815667264,
        "validation_file":0
    },
    "base_resp":{
        "status_code":0,
        "status_msg":"success"
    }
}
```

**python**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/retrieve_finetune_job?GroupId={GroupId}"

payload = json.dumps({
  "fine_tuning_job_id":"2367267362783"
})
headers = {
  'authorization': f'Bearer {API_KEY}',
  'content-type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```
##### 查看微调任务日志
提供以下在训练过程中计算的训练指标：训练集loss、验证集loss。这些统计数据旨在提供健全性检查，以确保训练顺利进行（loss应该减少），可以拿到4检索微调任务的result_files id，使用file的RetrieveContent接口查看训练的详细信息，包括训练集和验证集的loss信息
**查看微调任务日志请求示例**
```curl
curl --location 'https://api.minimax.chat/v1/list_finetune_event?GroupId=${group_id}' \
--header 'authorization: Bearer $MiniMax_API_KEY' \
--header 'content-type: application/json' \
--data '{"fine_tuning_job_id":"1701538478753114", "after":"0", "limit":10}'
```

**python**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/list_finetune_event?GroupId={GroupId}"

payload = json.dumps({
  "fine_tuning_job_id":"1701538478753114",
  "limit":10
})
headers = {
  'authorization': f'Bearer {API_KEY}',
  'content-type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```
**查看微调任务日志返回示例**
```python
{"file":{"file_id":85198633103453,"bytes":167,"created_at":1702632422,"filename":"1702631618780994.csv","purpose":"fine-tune-result"},"base_resp":{"status_code":0,"status_msg":"success"}}
iter,train_loss,valid_loss
1,0.324273,0.000000
2,0.343621,0.000000
3,0.151181,0.000000
4,0.076753,0.000000
5,0.079515,0.000000
6,0.017223,0.000000
7,0.013275,0.000000
```

**python**
```python
curl --location 'https://api.minimax.chat/v1/list_finetune_event?GroupId=${group_id}' \
--header 'authorization: Bearer ${api_key}' \
--header 'content-type: application/json' \
--data '{"fine_tuning_job_id":"1701538478753114", "after":"0", "limit":10}'
```
##### 使用微调模型
微调任务训练完成后，模型应立即可供推理使用，无需单独接口进行部署。在某些情况下，模型可能需要几分钟时间才能准备好处理请求。如果对模型的请求超时或找不到模型名称，则可能是因为模型仍在加载中。如果发生这种情况，请过几分钟再试。
**Tab: ChatCompletion接口示例**

**使用微调模型**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/text/chatcompletion?GroupId={GroupId}"

payload = json.dumps({
  "model":"ft:abab5.5s-chat-240123_1705584973436:dafei12",
  "tokens_to_generate":256,
  "temperature":0.7,
  "top_p":0.9,
  "stream":False,
  "role_meta":{
    "user_name":"用户",
    "bot_name":"文本总结助手"
  },
  "prompt":"文本总结助手能够理解文本，并以CSV格式系统地列出标题，作者，发行年份和类型。",
  "messages":[
    {
      "sender_type":"USER",
      "text":"1813年，英国著名作家简·奥斯汀（Jane Austen）创作了《傲慢与偏见》（Pride and Prejudice）。这是一部经典的爱情小说，也被视为英国文学的珍宝之一。小说讲述了女主角伊丽莎白·班内特与男主角达西先生之间的故事，描绘了社交、阶级和婚姻等题材。作品以其幽默、深刻的人物刻画和对社会习俗的嘲讽而闻名，被认为是英国文学中的经典之一，也是世界文学的杰作之一。"
    }
  ]
})
headers = {
  'authorization': f'Bearer {API_KEY}',
  'content-type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: ChatCompletion pro接口示例**

**使用微调模型**
```python
import requests
​
group_id="请填写您的group_id"
api_key="请填写您的api_key"
​
url = "https://api.minimax.chat/v1/text/chatcompletion_pro?GroupId=" + group_id
​
payload = {
    "bot_setting":[
        {
            "bot_name":"文本总结助手",
            "content":"文本总结助手能够理解文本，并以CSV格式系统地列出标题，作者，发行年份和类型。",
        }
    ],
    "messages":[{"sender_type":"USER", "sender_name":"小明", "text":"1813年，英国著名作家简·奥斯汀（Jane Austen）创作了《傲慢与偏见》（Pride and Prejudice）。这是一部经典的爱情小说，也被视为英国文学的珍宝之一。小说讲述了女主角伊丽莎白·班内特与男主角达西先生之间的故事，描绘了社交、阶级和婚姻等题材。作品以其幽默、深刻的人物刻画和对社会习俗的嘲讽而闻名，被认为是英国文学中的经典之一，也是世界文学的杰作之一。"}],
    "reply_constraints":{"sender_type":"BOT", "sender_name":"文本总结助手"},
    "model":"ft:abab5.5-chat-1019_1705584973436:dafei12",
    "tokens_to_generate":1034,
    "temperature":0.01,
    "top_p":0.95,
}
headers = {"Content-Type":"application/json", "Authorization":"Bearer " + api_key}
​
response = requests.request("POST", url, headers=headers, json=payload)
​
print(response.status_code)
print(response.text)
​
```
### API说明
#### finetune job object 数据结构
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| id | string | 否 | 训练任务ID。 |
| created_at | int64 | 否 | 创建任务的时间戳。 |
| fine_tuned_model | string | 否 | 微调模型的名称。 |
| result_files | list<string> | 否 | 训练日志的文件id。 |
| status | string | 否 | 任务状态。 |
| training_file | int64 | 否 | 用于训练的文件id。 |
| validation_file | int64 | 否 | 用于验证的文件id。 |
| tokens_count | int64 | 否 | token计数。 |
| hyperparameters | Hyperparameters | 否 | 训练超参。<br>指定全局批量大小，越大意味着模型参数的更新频率越低，默认值为1，建议填写范围：1~100<br>学习率，较小的学习率可能有助于避免过度拟合，默认值为1e-4，范围：1e-2 ~ 1e-6，注意：一般情况下不建议修改该参数<br>模型训练过程中的整个数据集的遍历次数，默认值为5，建议填写范围：1~100，注意：这个数值的大小会直接影响您的训练token消耗 |
| hyperparameters.batch_size | int64 | 否 | 指定全局批量大小，越大意味着模型参数的更新频率越低，默认值为1，建议填写范围：1~100 |
| hyperparameters.learning_rate_multiplier | double | 否 | 学习率，较小的学习率可能有助于避免过度拟合，默认值为1e-4，范围：1e-2 ~ 1e-6，注意：一般情况下不建议修改该参数 |
| hyperparameters.n_epochs | int64 | 否 | 模型训练过程中的整个数据集的遍历次数，默认值为5，建议填写范围：1~100，注意：这个数值的大小会直接影响您的训练token消耗 |
##### 创建微调任务
`API： https://api.minimax.chat/v1/create_finetune_job`
创建一个任务，基于给定数据集微调指定模型。
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| model | string | 是 | 基础模型，目前支持对以下模型进行微调（可根据自己的场景进行base model的选择）:`abab5.5-chat-240119`适合生产力工具场景，比如知识问答、总结说明、情感分析、文本分类、大纲&营销文案生成、作文批改等实时信息检索、信息提取、规定格式输出。`abab5.5s-chat-240123`适合人设对话场景，比如剧本撰写灵感提供、虚拟情感陪伴闲聊、多角色互动娱乐、动漫/游戏人物互动演绎。 |
| training_file | int64 | 是 | 用于训练的文件id。 |
| hyperparameters | hyperparameters | 否 | 训练超参数。 |
| validation_file | int64 | 否 | 用于验证的文件id。 |
| suffix | string | 否 | 微调模型名称的自定义后缀，最长8位长度，超长后会默认截断。 |
| finetune_job | FinetuneJob | 否 | FinetuneJob结构体如上文所述。 |
| base_resp | BaseResp | 否 | statusCode和statusMessage（平台请求通用code）。 |
##### 列出微调任务
`API： https://api.minimax.chat/v1/list_finetune_job`
列出当前groupid下的所有微调任务。
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| limit | int64 | 是 | 返回任务的数量限制。 |
| after | string | 否 | 从第几个ID之后开始列举。 |
##### 检索微调任务
`POST： https://api.minimax.chat/v1/retrieve_finetune_job?GroupId=$GroupId`
获取有关微调任务的详细信息。
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| fine_tuning_job_id | string | 是 | 任务id。 |
##### 删除微调任务
删除某个微调任务。
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| fine_tuning_job_id | string | 是 | 任务id。 |
| base_resp | BaseResp | 否 | statusCode和statusMessage（平台请求通用code）。 |
#### finetune event object 数据结构
该对象表示某个具体的微调任务事件，您可以通过该对象在训练过程中查看训练日志
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| id | string | 否 | 训练任务ID。 |
| created_at | int64 | 否 | 创建任务的时间戳。 |
| level | string | 否 | 级别。 |
| message | string | 否 | 训练任务ID的训练日志信息。 |
| object | string | 否 | 对象。 |
##### 列出微调事件
`POST： https://api.minimax.chat/v1/list_finetune_event?GroupId=$GroupId`
获取微调任务的状态更新。
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| fine_tuning_job_id | string | 是 | 任务id。 |
| limit | int64 | 是 | 返回任务的数量限制。 |
| after | string | 否 | 从第几个ID 之后开始列举。 |
#### funetune model object 数据结构
该对象表示训练得到的微调模型。一个groupid最多保留5个Finetune model，因此如果本身已经超过了限制，则无法创建微调任务，可以通过删除微调模型来对模型数量进行管理。
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| model_id | string | 否 | 微调模型ID。 |
| created_at | int64 | 否 | 创建任务的时间戳。 |
| object | string | 否 | 对象：finetune.model。 |
| base_model | string | 否 | 基础模型。 |
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| batch_size | int64 | 否 | 指定全局批量大小，越大意味着模型参数的更新频率越低。 |
| learning_rate_multiplier | double | 否 | 学习率，较小的学习率可能有助于避免过度拟合。 |
| n_epochs | int64 | 否 | 模型训练过程中的整个数据集的遍历次数。 |
##### 列举微调模型
`POST: https://api.minimax.chat/v1/list_finetune_model?GroupId=${GroupId}`
查看训练的所有微调模型。
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| 无 |  | 否 |  |
##### 查询微调模型
`POST:https://api.minimax.chat/v1/retrieve_finetune_model?GroupId=${GroupId}`
查询某个微调模型的详细信息。
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| model_id | string | 是 | 微调模型id。 |
##### 删除微调模型
`POST: https://api.minimax.chat/v1/delete_finetune_model?GroupId=${GroupId}`
删除某个微调模型。
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| model_id | string | 是 | 微调模型id。 |

## 创建首尾帧生成视频任务
- 原始链接: https://platform.minimaxi.com/document/first_last_frame_to _video?key=68abfca66fe587e3fbfe532c
- 文档ID: `68abfca66fe587e3fbfe532c`
# 首尾帧生成视频

<request method="POST" url="https://api.minimaxi.com/v1/video_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 

- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
模型名称。可用值：`MiniMax-Hailuo-02`
<hr/>

#### prompt    <gray>string</gray>
视频的文本描述，最大 2000 字符。对于 `MiniMax-Hailuo-02`支持使用 `[指令]` 语法进行运镜控制。
可在 `prompt` 中通过 `[指令]` 格式添加运镜指令，以实现精确的镜头控制。
- 支持 15 种运镜指令的指令:
  - 左右移: `[左移]`, `[右移]`
  - 左右摇: `[左摇]`, `[右摇]`
  - 推拉: `[推进]`, `[拉远]`
  - 升降: `[上升]`, `[下降]`
  - 上下摇: `[上摇]`, `[下摇]`
  - 变焦: `[变焦推近]`, `[变焦拉远]`
  - 其他: `[晃动]`, `[跟随]`, `[固定]`
- 使用规则:
  - 组合运镜: 同一组 `[]` 内的多个指令会同时生效，如 `[左摇,上升]`，建议组合不超过 3 个
  - 顺序运镜: `prompt` 中前后出现的指令会依次生效，如 `"...[推进], 然后...[拉远]"`
  - 自然语言: 也支持通过自然语言描述运镜，但使用标准指令能获得更准确的响应
  - 多运镜教程及模板，可参考 [运镜使用教程](https://b6j6u5f4zf.feishu.cn/docx/H03ZdRLh7oFsNpxE2ZqcUWb3njg) 
<hr/>

#### first_frame_image    <gray>string</gray>
将指定图片作为视频的起始帧。支持公网 URL 或 Base64 编码的 [Data URL](https://developer.mozilla.org/en-US/docs/Web/URI/Reference/Schemes/data) (`data:image/jpeg;base64,...`)。

- 图片要求：
  - 格式：JPG, JPEG, PNG, WebP
  - 体积：小于 20MB
  - 尺寸：短边像素大于 300px，长宽比在 2:5 和 5:2 之间

⚠️ 生成视频尺寸遵循首帧图片
<hr/>

#### last_frame_image    <gray>string</gray>
将指定图片作为视频的结束帧。支持公网 URL 或 Base64 编码的 [Data URL](https://developer.mozilla.org/en-US/docs/Web/URI/Reference/Schemes/data) (`data:image/jpeg;base64,...`)。

- 图片要求：
  - 格式：JPG, JPEG, PNG, WebP
  - 体积：小于 20MB
  - 尺寸：短边像素大于 300px，长宽比在 2:5 和 5:2 之间

⚠️ 生成视频尺寸遵循首帧图片，当首帧和尾帧的图片尺寸不一致时，模型将参考首帧对尾帧图片进行裁剪
<hr/>



#### prompt_optimizer    <gray>boolean</gray>
是否自动优化 `prompt`，默认为 `true`。设为 `false` 可进行更精确的控制
<hr/>


#### duration    <gray>int</gray>
视频时长（秒），默认值为 6。其可用值与模型和分辨率相关：

| 模型 | 768P | 1080P |
| :---  | :--- | :--- |
| MiniMax-Hailuo-02 | `6` 或 `10` | `6` |

<hr/>

#### resolution    <gray>string</gray>
视频分辨率。其可用值与模型相关，首尾帧功能支持 768，1080P

| 模型 | 6s | 10s |
| :--- | :--- | :--- |
| MiniMax-Hailuo-02 | `768P` (默认), `1080P` | `768P` |
<hr/>



#### callback_url    <gray>string</gray>
接收任务状态更新通知的回调 URL。支持通过 `callback_url` 参数可以配置回调，以接收任务状态的更新的异步通知

1. 地址验证：配置后，MiniMax 服务器会向 `callback_url` 发送一个 `POST` 请求，请求体中包含 `challenge` 字段。服务端需要在 3 秒内原样返回该 `challenge` 值以完成验证。
2. 状态更新：验证成功后，每当任务状态变更时，MiniMax 都会向该 URL 推送最新的任务状态。推送的数据结构与调用查询视频生成任务接口的响应体一致。


回调返回的"status"包括以下状态：
- "processing"-生成中
- "success"-成功
- "failed"-失败


<codetabs title="回调服务示例">
<codetab  language="python">
from fastapi import FastAPI, HTTPException, Request
from fastapi.middleware.cors import CORSMiddleware
import json
app = FastAPI()
@app.post("/get_callback")
async def get_callback(request: Request):
    try:
        json_data = await request.json()
        challenge = json_data.get("challenge")
        if challenge is not None:
          # is a verification request, just return the challenge
          return {"challenge": challenge}
        else:
            # is a callback request, do your own logic here
            # {
            #     "task_id": "115334141465231360",
            #     "status": "Success",
            #     "file_id": "205258526306433",
            #     "base_resp": {
            #         "status_code": 0,
            #         "status_msg": "success"
            #     }
            # }
            return {"status": "success"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        app, # 必选
        host="0.0.0.0", # 必选
        port=8000, # 必选，端口可设置
        # ssl_keyfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
        # ssl_certfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
    )
</codetab>

</codetabs>


<hr/>

#### aigc_watermark    <gray>boolean</gray>
是否在生成的视频中添加水印，默认为 `false`

<br>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
视频生成任务的 ID，用于后续查询任务状态
<hr/>

#### base_resp    <gray>object</gray>
状态码及详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>integer</gray>
状态码及其分别含义如下：

0：请求成功

1002：触发限流，请稍后再试

1004：账号鉴权失败，请检查 API-Key 是否填写正确

1008：账号余额不足

1026：视频描述涉及敏感内容，请调整

2013：传入参数异常，请检查入参是否按要求填写

2049：无效的api key，请检查api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情
<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>

</left>

<right>
<codetabs title="首尾帧视频生成">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/video_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "prompt": "一个小女孩从小成长到大的变化",
    "first_frame_image": "https://filecdn.minimax.chat/public/fe9d04da-f60e-444d-a2e0-18ae743add33.jpeg",
    "last_frame_image": "https://filecdn.minimax.chat/public/97b7cd08-764e-4b8b-a7bf-87a0bd898575.jpeg",
    "model": "MiniMax-Hailuo-02",
    "duration": 6,
    "resolution": "1080P",
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())

</codetab>
<codetab  language="bash">
curl -X POST "https://api.minimaxi.com/v1/video_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "prompt": "一个小女孩从小成长到大的变化",
    "first_frame_image": "https://filecdn.minimax.chat/public/fe9d04da-f60e-444d-a2e0-18ae743add33.jpeg",
    "last_frame_image": "https://filecdn.minimax.chat/public/97b7cd08-764e-4b8b-a7bf-87a0bd898575.jpeg",

    "model": "MiniMax-Hailuo-02",
    "duration": 6,
    "resolution": "1080P"
  }'

</codetab>
</codetabs>


<codetabs title="返回示例">
<codetab  language="json">
{
    "task_id": "106916112212032",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}

</codetab>
</codetabs>

</right>

</container>

<br>

<container gap="4" columns="2">

  <card  href="/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成指南](/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7)

  </card>

<card  href="/document/video_api_intro?key=68abdfdaaf3fef26183d6501">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成接口能力介绍](/document/video_api_intro?key=68abdfdaaf3fef26183d6501)
  </card>

</container>

## MiniMax M1 函数调用（Function Call）功能指南
- 原始链接: https://platform.minimaxi.com/document/function_call?key=68bbe0ac95a600dbfc923510
- 文档ID: `68bbe0ac95a600dbfc923510`
# MiniMax M1 函数调用（Function Call）功能指南

本文档将帮助您快速上手 [MiniMax-M1](https://huggingface.co/MiniMaxAI/MiniMax-M1-80k) 的函数调用功能。函数调用（Function Call）是一种通过提示工程实现的模型扩展能力，开发者可预先定义可用的函数，模型会根据用户输入自动决定是否调用函数，并以 JSON 格式输出调用参数。常见应用场景包括天气查询、网页搜索、数据库检索等。

## 使用 vLLM 进行函数调用（推荐）

请确保已成功部署 vLLM 并能够正常启动服务。MiniMax-M1 集成了自定义的 tool_call_parser，无需手动解析模型输出。在启动 vLLM 服务时，添加以下参数即可启用函数调用成功（新增参数为 `--enable-auto-tool-choice`、`--tool-call-parser minimax`、`--chat-template examples/tool_chat_template_minimax_m1.jinja`）

```python
 SAFETENSORS_FAST_GPU=1 VLLM_USE_V1=0 vllm serve MiniMaxAI/MiniMax-M1-40k \
     --trust-remote-code \
     --quantization experts_int8  
     --dtype bfloat16 \
     --enable-auto-tool-choice \
     --tool-call-parser minimax \
     --chat-template examples/tool_chat_template_minimax_m1.jinja
````

如果是通过 Docker 方式使用的 vLLM，在 Docker 启动命令中同样添加上述参数即可。

```python
 docker run --runtime nvidia --gpus all \
     -v ~/.cache/huggingface:/root/.cache/huggingface \
     --env "SAFETENSORS_FAST_GPU=1" \
     --env "VLLM_USE_V1=0" \
     -p 8000:8000 \
     --ipc=host \
     vllm/vllm-openai:latest \
     --model MiniMaxAI/MiniMax-M1-40k \
     --trust-remote-code \
     --quantization experts_int8 \
     --dtype bfloat16 \
    --enable-auto-tool-choice \
    --tool-call-parser minimax \
    --chat-template examples/tool_chat_template_minimax_m1.jinja
```

参数解释：

* `--tool-call-parser minimax` 关键参数，启用 MiniMax-M1 自定义解析器
* `--enable-auto-tool-choice` 启用自动工具选择
* `--chat-template` 模板文件需要适配 tool calling 格式，可以在 [GitHub](https://github.com/vllm-project/vllm/blob/main/examples/tool_chat_template_minimax_m1.jinja) 获取

## 通过 OpenAI SDK 使用函数调用

以下示例演示如何通过 OpenAI SDK 实现天气查询函数调用：

```python
from openai import OpenAI
import json

client = OpenAI(base_url="http://localhost:8000/v1", api_key="dummy")

def get_weather(location: str, unit: str):
    return f"Getting the weather for {location} in {unit}..."

tool_functions = {"get_weather": get_weather}

tools = [{
    "type": "function",
    "function": {
        "name": "get_weather",
        "description": "Get the current weather in a given location",
        "parameters": {
            "type": "object",
            "properties": {
                "location": {"type": "string", "description": "City and state, e.g., 'San Francisco, CA'"},
                "unit": {"type": "string", "enum": ["celsius", "fahrenheit"]}
            },
            "required": ["location", "unit"]
        }
    }
}]

response = client.chat.completions.create(
    model=client.models.list().data[0].id,
    messages=[{"role": "user", "content": "What's the weather like in San Francisco? use celsius."}],
    tools=tools,
    tool_choice="auto"
)

print(response)

tool_call = response.choices[0].message.tool_calls[0].function
print(f"Function called: {tool_call.name}")
print(f"Arguments: {tool_call.arguments}")
print(f"Result: {get_weather(**json.loads(tool_call.arguments))}")
```

示例输出如下：

```text
Function called: get_weather
Arguments: {"location": "San Francisco, CA", "unit": "celsius"}
Result: Getting the weather for San Francisco, CA in celsius...
```

## 函数调用定义格式说明

### 函数定义格式

函数调用需要在请求体中定义 `tools` 字段，每个函数由以下部分组成：

```json
{
  "tools": [
    {
      "name": "search_web",
      "description": "搜索函数。",
      "parameters": {
        "properties": {
          "query_list": {
            "description": "进行搜索的关键词，列表元素个数为1。",
            "items": { "type": "string" },
            "type": "array"
          },
          "query_tag": {
            "description": "query的分类",
            "items": { "type": "string" },
            "type": "array"
          }
        },
        "required": [ "query_list", "query_tag" ],
        "type": "object"
      }
    }
  ]
}
```
<br>

字段说明：
1. `name` : 函数名称
2. `description` : 函数功能描述
3. `parameters` : 函数参数定义

   * `properties` : 参数属性定义，key 是参数名，value 包含参数的详细描述
   * `required` : 必填参数列表
   * `type` : 参数类型（通常为 "object"）

### 模型内部处理格式

在模型内部处理时，函数定义会被转换为特殊格式并拼接到输入文本中。以下为模型内部处理时的输入格式示例，用户无需手动构造：

```xml
<begin_of_document><beginning_of_sentence>system ai_setting=MiniMax AI
MiniMax AI是由上海稀宇科技有限公司（MiniMax）自主研发的AI助理。<end_of_sentence>
<beginning_of_sentence>system tool_setting=tools
You are provided with these tools:
<tools>
{"name": "search_web", "description": "搜索函数。", "parameters": {"properties": {"query_list": {"description": "进行搜索的关键词，列表元素个数为1。", "items": {"type": "string"}, "type": "array"}, "query_tag": {"description": "query的分类", "items": {"type": "string"}, "type": "array"}}, "required": ["query_list", "query_tag"], "type": "object"}}
</tools>
If you need to call tools, please respond with <tool_calls></tool_calls> XML tags, and provide tool-name and json-object of arguments, following the format below:
<tool_calls>
{"name": <tool-name>, "arguments": <args-json-object>}
...
</tool_calls><end_of_sentence>
<beginning_of_sentence>user name=用户
OpenAI 和 Gemini 的最近一次发布会都是什么时候?<end_of_sentence>
<beginning_of_sentence>ai name=MiniMax AI
```

### 模型输出格式

模型输出的函数调用格式如下：

```xml
<think>
Okay, I will search for the OpenAI and Gemini latest release.
</think>
<tool_calls>
{"name": "search_web", "arguments": {"query_tag": ["technology", "events"], "query_list": ["\"OpenAI\" \"latest\" \"release\""]}}
{"name": "search_web", "arguments": {"query_tag": ["technology", "events"], "query_list": ["\"Gemini\" \"latest\" \"release\""]}}
</tool_calls>
```

## 手动解析模型输出

我们推荐使用 OpenAI Chat Completions API，其会自动在服务端应用 Chat Template，主流推理框架已有完善的支持。
如使用的框架不支持 Tool Calling，未使用 vLLM 内置解析器，或使用其他推理框架（如 Transformers、TGI 等），可参考以下方法手动解析模型输出，需根据实际模型输出格式调整解析逻辑。

### 手动应用 Chat Template

以下是使用 transformers 库的完整示例：

```python
from transformers import AutoTokenizer

def get_default_tools():
    return [
        {
          "name": "get_current_weather",
          "description": "Get the latest weather for a location",
          "parameters": {
              "type": "object", 
              "properties": {
                  "location": {
                      "type": "string", 
                      "description": "A certain city, such as Beijing, Shanghai"
                  }
              }, 
          },
          "required": ["location"],
          "type": "object"
        }
    ]

# 加载模型和分词器
model_id = "MiniMaxAI/MiniMax-M1-40k"
tokenizer = AutoTokenizer.from_pretrained(model_id)
prompt = "What's the weather like in Shanghai today?"
messages = [
    {"role": "system", "content": [{"type": "text", "text": "You are a helpful assistant created by MiniMax based on MiniMax-M1 model."}]},
    {"role": "user", "content": [{"type": "text", "text": prompt}]},
]

with open("/vllm-workspace/examples/tool_chat_template_minimax_m1.jinja", "r") as f:
    tokenizer.chat_template = f.read()

# 启用函数调用工具
tools = get_default_tools()

# 应用聊天模板，并加入工具定义
text = tokenizer.apply_chat_template(
    messages,
    tokenize=False,
    add_generation_prompt=True,
    tools=tools
)

# 发送请求（这里使用任何推理服务）
import requests
payload = {
    "model": model_id,
    "prompt": text,
    "max_tokens": 4000
}
response = requests.post(
    "http://localhost:8000/v1/completions",
    headers={"Content-Type": "application/json"},
    json=payload,
    stream=False,
)

# 模型输出需要手动解析
raw_json = response.json()
raw_output = raw_json["choices"][0]["text"]
print("原始输出:", raw_output)

# 使用下面的解析函数处理输出
function_calls = parse_function_calls(raw_output)
```

### 解析函数调用

当需要手动解析时，需要解析模型输出的 XML 标签格式：

```python
import re
import json
def parse_function_calls(content: str):
    """
    解析模型输出中的函数调用
    """
    function_calls = []
    
    # 匹配 <tool_calls> 标签内的内容
    tool_calls_pattern = r"<tool_calls>(.*?)</tool_calls>"
    tool_calls_match = re.search(tool_calls_pattern, content, re.DOTALL)
    
    if not tool_calls_match:
        return function_calls
    
    tool_calls_content = tool_calls_match.group(1).strip()
    
    # 解析每个函数调用（每行一个JSON对象）
    for line in tool_calls_content.split('\n'):
        line = line.strip()
        if not line:
            continue
            
        try:
            # 解析JSON格式的函数调用
            call_data = json.loads(line)
            function_name = call_data.get("name")
            arguments = call_data.get("arguments", {})
            
            function_calls.append({
                "name": function_name,
                "arguments": arguments
            })
            
            print(f"调用函数: {function_name}, 参数: {arguments}")
            
        except json.JSONDecodeError as e:
            print(f"参数解析失败: {line}, 错误: {e}")
    
    return function_calls

# 示例：处理天气查询函数
def execute_function_call(function_name: str, arguments: dict):
    """
    执行函数调用并返回结果
    """
    if function_name == "get_current_weather":
        location = arguments.get("location", "未知位置")
        # 构建函数执行结果
        return {
            "role": "tool", 
            "content": [
              {
                "name": function_name,
                "type": "text",
                "text": json.dumps({
                    "location": location, 
                    "temperature": "25", 
                    "unit": "celsius", 
                    "weather": "晴朗"
                }, ensure_ascii=False)
              }
            ] 
          }
    elif function_name == "search_web":
        query_list = arguments.get("query_list", [])
        query_tag = arguments.get("query_tag", [])
        # 模拟搜索结果
        return {
            "role": "tool",
            "content": [
              {
                "name": function_name,
                "type": "text",
                "text": f"搜索关键词: {query_list}, 分类: {query_tag}\n搜索结果: 相关信息已找到"
              }
            ]
          }
    
    return None
```


## 将函数执行结果返回给模型

成功解析函数调用后，应将函数执行结果添加到对话历史中，以便模型在后续交互中能够访问和利用这些信息。

### 单个结果

假如模型调用了 `search_web` 函数，可以参考如下格式添加执行结果，`name` 字段为具体的函数名称。

```json
{
  "role": "tool", 
  "content": [
    {
      "name": "search_web",
      "type": "text",
      "text": "test_result"
    }
  ]
}
````

对应如下的模型输入格式：

```text
<beginning_of_sentence>tool name=tools
tool name: search_web
tool result: test_result
<end_of_sentence>
```

### 多个结果

假如模型同时调用了 `search_web` 和 `get_current_weather` 函数，可以参考如下格式添加执行结果，`content` 包含多个结果。

```json
{
  "role": "tool", 
  "content": [
    {
      "name": "search_web",
      "type": "text",
      "text": "test_result1"
    },
    {
      "name": "get_current_weather",
      "type": "text",
      "text": "test_result2"
    }
  ]
}
```

对应如下的模型输入格式：

```text
<beginning_of_sentence>tool name=tools
tool name: search_web
tool result: test_result1
tool name: get_current_weather
tool result: test_result2<end_of_sentence>
```

## 常见问题

### Tool Call 回应不是合法的 JSON 格式

请检查发送的请求，如果请求中的 Tool Call 调用引导也不是合法的 JSON 格式，说明应用了错误的 chat template，请使用 [tool\_chat\_template\_minimax\_m1.jinja](https://github.com/vllm-project/vllm/blob/main/examples/tool_chat_template_minimax_m1.jinja)

### 参考资料

* [MiniMax-M1 模型仓库](https://github.com/MiniMaxAI/MiniMax-M1)
* [vLLM 项目主页](https://github.com/vllm-project/vllm)
* [vLLM Function Calling PR](https://github.com/vllm-project/vllm/pull/20297)
* [OpenAI Python SDK](https://github.com/openai/openai-python)

## 获取支持

如果在部署 MiniMax 模型过程中遇到任何问题：

* 通过官方邮箱 [api@minimaxi.com](api@minimaxi.com) 联系我们的技术支持团队
* 在我们的 GitHub 仓库提交 [Issue](https://github.com/MiniMax-AI/MiniMax-M1/issues)

我们会持续优化 Function Call 的部署体验，欢迎您的反馈！

## 查询可用音色
- 原始链接: https://platform.minimaxi.com/document/get_voice?key=68b51016d2f0aaaf3484a8de
- 文档ID: `68b51016d2f0aaaf3484a8de`
# 查询可用音色ID
该 API 支持查询**当前账号**下可调用的**全部音色 ID**（voice_id）。
包括系统音色、快速克隆音色、文生音色接口生成的音色、音乐生成接口的人声音色以及伴奏音色。
<request method="POST" url="https://api.minimaxi.com/v1/get_voice"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>Required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON
<br/>

### 请求体
<hr/>

#### voice_type    <gray>string</gray> <red>Required</red>
希望查询音色类型，支持以下取值：
- `system`：系统音色
- `voice_cloning`：快速复刻的音色，仅在成功用于语音合成后才可查询
- `voice_generation`：文生音色接口生成的音色，仅在成功用于语音合成后才可查询
- `all`：以上全部

<br/>

### 返回体
<hr/>

#### system_voice    <gray>array</gray>
包含系统预定义的音色。
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray>
音色 ID
<hr/>

#### voice_name    <gray>string</gray>
音色名称，非调用的音色 ID
<hr/>

#### description    <gray>array</gray>
音色描述

</card>
</details>
<hr/>

#### voice_cloning    <gray>array</gray>
包含音色快速复刻的音色数据
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray>
快速复刻音色 ID
<hr/>

#### description    <gray>array</gray>
复刻时填写的音色描述
<hr/>

#### created_time    <gray>string</gray>
创建时间，格式 yyyy-mm-dd

</card>
</details>
<hr/>

#### voice_generation    <gray>array</gray>
包含音色生成接口产生的音色数据
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray>
音色 ID
<hr/>

#### description    <gray>array</gray>
生成音色时填写的音色描述
<hr/>

#### created_time    <gray>string</gray>
创建时间，格式 yyyy-mm-dd

</card>
</details>
<hr/>

#### base_resp    <gray>object</gray>
本次请求的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码
- 0: 请求结果正常
- 2013：输入参数信息不正常

更多信息请参考 [错误码查询](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
"""
本示例用于查询可用音色。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import os

api_key = os.environ.get("MINIMAX_API_KEY")

url = f'https://api.minimaxi.com/v1/get_voice'
headers = {
    'authority': 'api.minimaxi.com',
    'Authorization': f'Bearer {api_key}',
    'content-type': 'application/json'
}

data = {
    'voice_type': 'all'
}

response = requests.post(url, headers=headers, json=data)
print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/get_voice' \
--header 'content-type: application/json' \
--header "authorization: Bearer $MINIMAX_API_KEY" \
--data '{
    "voice_type":"all"
}'
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "system_voice": [
        {
          "voice_id":"Chinese (Mandarin)_Reliable_Executive",
          "description":["一位沉稳可靠的中年男性高管声音，标准普通话，传递出值得信赖的感觉。"],
          "voice_name":"沉稳高管",
          "created_time":"1970-01-01"
        },
        {
          "voice_id":"Chinese (Mandarin)_News_Anchor",
          "description":["一位专业、播音腔的中年女性新闻主播，标准普通话。"],
          "voice_name":"新闻女声",
          "created_time":"1970-01-01"
        }
    ],
    "voice_cloning": [
        {
          "voice_id":"test12345",
          "description":[],
          "created_time":"2025-08-20"
        },
        {
          "voice_id":"test12346",
          "description":[],
          "created_time":"2025-08-21"
        }
    ],
    "voice_generation": [
        {
          "voice_id":"ttv-voice-2025082011321125-2uEN0X1S",
          "description":[],
          "created_time":"2025-08-20"
        },
        {
          "voice_id":"ttv-voice-2025082014225025-ZCQt0U0k",
          "description":[],
          "created_time":"2025-08-20"
        }
    ],
    "base_resp":{
      "status_code":0,
      "status_msg":"success"
      }
}
</codetab>
</codetabs>

</right>

</container>

## 在AI编程工具中使用M1
- 原始链接: https://platform.minimaxi.com/document/guides_coding_m1?key=68b9b1b4f7d629d07d273bd8
- 文档ID: `68b9b1b4f7d629d07d273bd8`
# 在 AI 编程工具里使用 MiniMax-M1


**MiniMax-M1** 具备良好的**代码理解**、**多轮对话**和**推理能力**，兼容 OpenAI 接口协议，适用于代码助手、Agent 工具、AI IDE 等多种场景。本教程以当前主流的 AI 编程工具为例，分别介绍如何在 [Cursor](https://cursor.com/)(ide), [Cline](https://cline.bot/)(插件)中配置并使用 **MiniMax-M1** 模型。

## 获取 API Key

*   访问 [MiniMax开放平台](https://platform.minimaxi.com/user-center/basic-information/interface-key)
*   点击“**Create new secret key**”按钮，输入项目名称以创建新的 API Key
*   创建成功后，系统将展示 API Key。**请务必复制并妥善保存**，该密钥**只会显示一次**，无法再次查看

<br>

<img src="https://filecdn.minimax.chat/public/e99e332e-7ec0-44fe-8c90-731eb74fcd06.png" width="60%" />




## 1. 在 Cursor 中使用 MiniMax-M1

[Cursor ](https://cursor.com/) 是基于 VS Code 改造的 AI 原生 IDE，提供良好的端到端的 AI 编程体验，支持代码解释、重构、调试等智能协作能力。下面将介绍如何在 Cursor 中使用 MiniMax-M1。

### 安装Cursor

1.  通过 [Cursor官网](https://cursor.com/) 下载并安装 Cursor
2.  打开 Cursor，右上角“设置”按钮，进入设置界面。点击“Sign in”按钮，登录自己的 Cursor 账户


<img src="https://filecdn.minimax.chat/public/3b07b51e-8aff-41c6-9457-a082ffd9165a.jpeg"  width="60%" />

### 在 Cursor 中配置MiniMax API

1.  点击左侧栏的 **“Models”**，进入模型配置页面
2.  展开 **“API Keys”** 部分，配置 API 信息：
    *   勾选 **“Override OpenAI Base URL”** 
    *   在下方输入 MiniMax 的调用地址： `https://api.minimaxi.com/v1`
3.  在 **OpenAI API Key** 输入框，配置从 [MiniMax开放平台](https://platform.minimaxi.com/user-center/basic-information/interface-key) 获取的 API Key


<img src="https://filecdn.minimax.chat/public/b6dba6fd-d4d7-4778-b390-c4a08e901a9d.jpeg"  width="60%" />
<br>

4.  点击 **“Verify”** 按钮，在弹出的窗口中点击 **“Enable OpenAI API Key”** 按钮，完成设置验证。


<img src="https://filecdn.minimax.chat/public/67d40b7c-1a18-4553-a0f8-9ceb2dc610a4.png" width="60%" />

5.  在 **Models** 板块中，点击 **“Add Custom Model”** 按钮，输入模型名称“MiniMax-M1”，点击“Add”按钮
6.  启用刚添加的 “MiniMax-M1” 模型

<img src="https://filecdn.minimax.chat/public/2a69df97-266b-43cc-a274-30d75512e930.png" alt="图片描述" width="60%" />
<br><br>

### 在 Cursor 中体验 MiniMax-M1

1.  在对话框中选择 **MiniMax-M1** 模型


<img src="https://filecdn.minimax.chat/public/a18a3ef3-84d4-4ec2-9db4-a7ca846c1a1e.png" width="60%" />

2.  使用 MiniMax-M1 创建一个迷宫生成器和路径寻找的可视化网页，提示词可参考

> You are a helpful web developer assistant. Create a maze generator and pathfinding visualizer. Randomly generate a maze and visualize A* algorithm solving it step by step. Use canvas and animations. Make it visually appealing.




<img src="https://filecdn.minimax.chat/public/ad838923-5119-4dfa-8af6-7845fc882808.png" width="60%" />

3.  该生成网页的实际效果为

<video
  src="https://filecdn.minimax.chat/public/0e10adbf-297e-49e0-86d1-17c9ffb8b41e.mov"
  width="100%"
  ratio="1942:1290">
</video>


## 2. 在 Cline 中使用 MiniMax-M1

**Cline** 是一个能够使用终端和编辑器的 ai 助手，兼具代码编辑、多模态输入等能力，下面以 VS Code 的 Cline 插件为例，介绍如何配置使用 MiniMax-M1。

### 安装Cline

1.  打开 VS Code，点击左侧活动栏中的扩展图标，在搜索框中输入 `Cline`
2.  点击 `Install` 按钮进行安装，安装完成后，可能需要重启 VS Code


<img src="https://filecdn.minimax.chat/public/f1ba35cd-16b0-4bb3-aec0-6e27d5c5c7e9.png" alt="图片描述" width="60%" />

3.  完成安装后，可以在左侧活动栏中看到Cline的图标

<img src="https://filecdn.minimax.chat/public/cc07cac7-ef54-49ad-920a-5dd48c363911.png" alt="图片描述" width="60%" />

### 在 Cline 中配置 MiniMax API

1.  点击“**Use your own API key**”按钮，进入 Cline 的参数配置界面
2.  API Provider -> 选择“**OpenAI Compatible**”，使用 OpenAI 协议标准访问 MiniMax API
3.  Base URL 输入“ **https://api.minimaxi.com/v1** ”，设置 MiniMax 服务入口
4.  OpenAI Compatible Key 输入在 [MiniMax开放平台](https://platform.minimaxi.com/user-center/basic-information/interface-key) 获取的key
5.  Model ID 配置输入 “**MiniMax-M1**”，指定要调用的具体模型名称
6.  点击“Let's go!”按钮后，点击右上角的"Done"，保存配置

<img src="https://filecdn.minimax.chat/public/7ff67733-f047-4e04-9cfa-42c23837faf0.jpeg" alt="图片描述" width="60%" />


### 在 Cline 中体验MiniMax-M1

1.  勾选Auto-approve中的"Edit"选项框

<img src="https://filecdn.minimax.chat/public/c64ebe1a-81fd-4978-a2b4-6c4ac2046a9c.jpeg" alt="图片描述" width="60%" />

2.  使用 MiniMax-M1 创建一个打字速度测试的可视化网页，提示词可参考

> You are a helpful web developer assistant. Build a typing speed test web app. Randomly show a sentence, and track the user's typing speed in WPM (words per minute). Provide live feedback with colors and accuracy. Make it visually appealing.

<img src="https://filecdn.minimax.chat/public/f9f1b725-39b3-4a65-b1ab-277e1b14beac.png" alt="图片描述" width="60%" />

3.  该网页的生成效果为

<video
  src="https://filecdn.minimax.chat/public/7d49b98c-baa3-449b-a567-6c0fa5cd702c.mov"
  width="100%"
  ratio="1878:1248">
</video>

## 图片生成指南
- 原始链接: https://platform.minimaxi.com/document/guides_image_generation?key=68bbf07595a600dbfc92689d
- 文档ID: `68bbf07595a600dbfc92689d`
# 图片生成

图片生成服务提供文生图（text-to-image）与图生图（image-to-image）两种核心功能。

## 根据文本生成图片

根据详尽的文本描述（prompt），直接生成与之匹配的图片。

```python
import base64
import requests
import os

url = "https://api.minimaxi.com/v1/image_generation"
api_key = os.environ.get("MINIMAX_API_KEY")
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "image-01",
    "prompt": "men Dressing in white t shirt, full-body stand front view image :25, outdoor, Venice beach sign, full-body image, Los Angeles, Fashion photography of 90s, documentary, Film grain, photorealistic",
    "aspect_ratio": "16:9",
    "response_format": "base64",
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()

images = response.json()["data"]["image_base64"]

for i in range(len(images)):
    with open(f"output-{i}.jpeg", "wb") as f:
        f.write(base64.b64decode(images[i]))
```

生成结果：
![image](https://filecdn.minimax.chat/public/5d969bc8-047c-424d-b7f8-87ece21e8f13.jpeg)
## 结合参考图生成图片

此功能允许提供一张或多张包含清晰主体的参考图（支持网络图片链接），并结合 prompt 描述，生成一张保留了主体特征的新图片。该功能尤其适用于需要保持人物形象一致性的场景，例如为同一个虚拟角色生成不同情境下的图片。

```python
import base64
import requests
import os

url = "https://api.minimaxi.com/v1/image_generation"
api_key = os.environ.get("MINIMAX_API_KEY")
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "image-01",
    "prompt": "女孩在图书馆的窗户前，看向远方",
    "aspect_ratio": "16:9",
    "subject_reference": [
        {
            "type": "character",
            "image_file": "https://cdn.hailuoai.com/prod/2025-08-12-17/video_cover/1754990600020238321-411603868533342214-cover.jpg",
        }
    ],
    "response_format": "base64",
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
images = response.json()["data"]["image_base64"]

for i in range(len(images)):
    with open(f"output-{i}.jpeg", "wb") as f:
        f.write(base64.b64decode(images[i]))
```


生成结果：
![image](https://filecdn.minimax.chat/public/0a70cf90-9ae2-4593-b4a3-84ec9ab3429a.jpeg)

## MCP 使用指南
- 原始链接: https://platform.minimaxi.com/document/guides_mcp?key=68bbf0f295a600dbfc926bdd
- 文档ID: `68bbf0f295a600dbfc926bdd`
# MiniMax MCP使用指南

## 模型上下文协议（MCP）简介

**模型上下文协议(<a href="https://modelcontextprotocol.io/docs/getting-started/intro" target="_blank">MCP</a>）** 是一个开放协议，标准化应用程序向大语言模型提供工具和上下文的方式。它类似于 AI 领域的 USB‑C 接口，提供一个稳定和标准化的接入点，让模型能访问数据库、API、插件或其他工具。通过 MCP 工具，开发者可以让模型访问托管在远程 MCP 服务器上的各种工具。

**MiniMax 提供官方的 <a href="https://github.com/MiniMax-AI/MiniMax-MCP" target="_blank">Python 版本</a>和 <a href="https://github.com/MiniMax-AI/MiniMax-MCP-JS" target="_blank">JavaScript 版本</a> 模型上下文协议（MCP），** 支持文本转语音（TTS）、声音克隆、图像生成、视频生成等多模态能力。 开发者可自行部署 MCP 服务，并通过 MCP 客户端（如 Claude Desktop、Cursor、Windsurf、OpenAI Agents 等）调用，从而快速集成语音、图像和视频相关功能。在传输方面，Python 版本提供 stdio 和 SSE 两种标准传输方式，JS 版本提供 stdio 、REST 和 SSE 三种标准传输方式。

## MiniMax MCP 工具和参数介绍 

### MCP 工具清单

<container gap="4" columns="3">

<card>

<img src="https://filecdn.minimax.chat/public/1348ae2d-698c-49ce-a5b7-1bac5e024af2.png" width="32px" />

### text_to_audio
该工具可将将输入的文本合成为自然流畅的语音
<br>


</card>

<card>
<img src="https://filecdn.minimax.chat/public/1348ae2d-698c-49ce-a5b7-1bac5e024af2.png" width="32px" />

### list_voices
该工具可查询所有可用音色
<br>


</card>

<card>
<img src="https://filecdn.minimax.chat/public/1348ae2d-698c-49ce-a5b7-1bac5e024af2.png" width="32px" />

### voice_clone
该工具可根据指定音频文件克隆音色
<br>


</card>

<card>

<img src="https://filecdn.minimax.chat/public/1348ae2d-698c-49ce-a5b7-1bac5e024af2.png" width="32px" />

### voice_design
该工具可根据指定提示词生成音色和试听文本
<br>


</card>

<card>

<img src="https://filecdn.minimax.chat/public/1348ae2d-698c-49ce-a5b7-1bac5e024af2.png" width="32px" />

### play_audio
该工具用于播放一个音频文件
<br>


</card>

<card>

<img src="https://filecdn.minimax.chat/public/74224aef-9b4d-452e-8e1e-e2241eb482e2.png" width="32px" />

### music_generation
该工具可根据指定提示词和歌词生成音乐
<br>


</card>

<card>

<img src="https://filecdn.minimax.chat/public/53ba5e61-bc46-4e6a-ad64-453e9b85af0d.png" width="32px" />

### generate_video
该工具可根据指定文本或图片进行视频生成生成
<br>


</card>

<card>

<img src="https://filecdn.minimax.chat/public/53ba5e61-bc46-4e6a-ad64-453e9b85af0d.png" width="32px" />

### image_to_video
该工具用于使用首帧图像生成视频
<br>


</card>

<card>
<img src="https://filecdn.minimax.chat/public/53ba5e61-bc46-4e6a-ad64-453e9b85af0d.png" width="32px" />

### query_video_generation
该工具用于查询异步视频生成任务的状态
<br>


</card>

<card>

<img src="https://filecdn.minimax.chat/public/71a90166-3dbe-452a-83ed-17e553ea3220.png" width="32px" />

### text_to_image
该工具可根据指定提示词生成图片
<br>

</card>

</container>

## 工具与参数详情

### 1. text_to_audio
该工具可将将输入的文本合成为自然流畅的语音。
| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| `text` | **必需**，待合成的文本 | 字符串，长度限制<10000字符，段落切换用换行符替代。 | 无 |
| `output_directory` | 保存音频文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |
| `voice_id` | 请求的音色编号 | 字符串，可选值可参考 [API手册](#) | `"female-shaonv"` |
| `model` | 请求的模型版本 | 字符串，可选值可参考 [API手册](#) | `"speech-02-hd"` |
| `speed` | 生成音频的语速 | 浮点数，范围 [0.5 - 2.0] | `1.0` |
| `vol` | 生成音频的音量 | 浮点数，范围 | `1.0` |
| `pitch` | 生成音频的音调 | 整数，范围 [-12,12] | `0` (原音色输出) |
| `emotion` | 控制合成语音的情绪 | 字符串，可选值范围`["happy", "sad", "angry", "fearful", "disgusted", "surprised", "calm"]`，该参数仅对 speech-02-hd, speech-02-turbo, speech-01-turbo, speech-01-hd 模型生效 | `happy` |
| `sample_rate` | 生成音频的采样率 | 整数，可选范围 `[8000, 16000, 22050, 24000, 32000, 44100]` | `32000` |
| `bitrate` | 生成音频的比特率 | 整数，可选范围 `[32000, 64000, 128000, 256000]`，该参数仅对mp3格式的音频生效 | `128000` |
| `channel` | 生成音频的声道数 | 整数，可选值为 1：单声道, 2：双声道 | `1` |
| `format` | 生成的音频文件格式 | 字符串，可选范围`["pcm", "mp3", "flac", "wav"]`，其中，"wav" 仅在非流式传输下支持 | `mp3` |
| `language_boost` | 语言增强选项，提升在指定小语种或方言场景下的语音表现 | 字符串，可选值可参考 [API手册](#) | `null` |

### 2. list_voices
该工具可查询所有可用音色

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| voice_type | 需要查询的音色类型。 | 字符串, 可选值范围 ["system" (系统音色) , "voice_cloning" (快速复刻的音色) , "voice_generation" (文生音色接口生成的音色) , "music_generation" (音乐生成产生的人声或者伴奏音色) , "all" (以上全部) ] | "all" |

### 3. voice_clone
该工具可根据指定音频文件克隆音色

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| voice_id | **必需**, 音色复刻的声音编号 | 字符串, 可自定义, 注意事项包括<br>1. 自定义的voice_id长度范围<br>2. 首字符必须为英文字母<br>3. 允许数字、字母、-、_<br>4. 末位字符不可为-、_<br>5. 创建的voice_id不可与之前重复 | 无 |
| file | **必需**, 用于克隆的音频文件 | 字符串, 可选范围["mp3", "m4a", "wav"] | 无 |
| text | 生成克隆音色演示音频的文本 | 字符串, 限制2000字符以内 | 无 |
| output_directory | 保存音频文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |
| is_url | 用于克隆的音频文件来源是否url | 布尔值 (True / False) | FALSE |

### 4. voice_design
该工具可根据指定提示词生成音色和试听文本

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| prompt | **必需**, 生成音色的描述 | 字符串 | 无 |
| preview_text | **必需**, 生成试听音频的文本 | 字符串 | 无 |
| voice_id | 自定义生成音色的ID | 字符串 | 自动生成唯一值 |
| output_directory | 保存试听音频文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |

<br>

### 5. play_audio
该工具用于播放一个音频文件。

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| input_file_path | **必需**, 要播放的音频文件的本地路径或URL | 字符串 (文件路径或URL) | 无 |
| is_url | 需要播放的音频文件来源是否url | 布尔值 (True / False) | FALSE |

### 6. music_generation
该工具可根据指定提示词和歌词生成音乐

| 参数 | 含义 | 可选值/格式 | 默认值 |
| :--- | :--- | :--- | :--- |
| prompt | **必需**, 表示音乐创作的灵感，通常是对音乐风格、情绪、场景等的描述 | 字符串, 字符取值范围 [10-300] | 无 |
| lyrics | **必需**, 音乐生成的歌词 | 字符串, 使用换行符 (`\n`) 分隔每行歌词, 支持结构标签 [Intro][Verse][Chorus][Bridge][Outro], 字符取值范围 | 无 |
| sample_rate | 生成音乐的采样率 | 整数, 可选值范围 | 32000 |
| bitrate | 生成音乐的比特率 | 整数, 可选值范围 | 128000 |
| format | 生成音乐的格式 | 字符串, 可选值范围 ["mp3", "wav", "pcm"] | "mp3" |
| output_directory | 保存音乐文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |

### 7. generate_video
该工具可根据指定提示词生成视频, `prompt` 和 `first_frame_image` 两个参数至少要有一个。

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| `prompt` | 生成视频的描述 | 字符串, 最大支持2000字符 | 无 |
| `model` | 请求的模型版本 | 字符串, 可选范围 `MiniMax-Hailuo-02`, `T2V-01-Director`, `I2V-01-Director`, `S2V-01`, `I2V-01-live`, `I2V-01`, `T2V-01` | `T2V-01` |
| `first_frame_image` | 视频的首帧画面 | 支持传入图片的 "data:image/jpeg;base64,{data}" 格式的 Base64 编码字符串, 或可通过公网访问的 URL | 无 |
| `duration` | 生成视频的持续时间 (秒) | 整数, 可选值, 与分辨率和模型相关:<br>1. **01系列** (包括 T2V-01, I2V-01,T2V-01-Director, I2V-01-Director, I2V-01-live, S2V-01):<br>   可选值: 6<br>2. **02系列** (MiniMax-Hailuo-02) :<br>   512P: 可选值: 6, 10<br>   768P: 可选值: 6, 10<br>   1080P: 可选值: 6 | `6` |
| `resolution` | 生成视频的分辨率 | 字符串, 与选择模型及设置的视频时长相关:<br>1. **01系列** (包括 T2V-01, I2V-01,T2V-01-Director, I2V-01-Director, I2V-01-live, S2V-01): 不支持设置本参数。<br>2. **02系列**<br>   6s时长: 默认值为"720P"。 <br>   6s时长: 默认值"768P", 可选范围 `["512P", "768P", "1080P"]` <br>   10s时长: 默认值"768P", 可选范围 `["512P", "768P"]` | `None` |
| `output_directory` | 保存视频文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |
| `async_mode` | 是否使用异步模式, 启用则返回任务ID | 布尔值 (True / False) | `FALSE` |

### 8. image_to_video
该工具用于使用首帧图像生成视频, 其中 `prompt` 和 `first_frame_image` 两个参数至少有其一, 该 MCP 工具仅在 **JavaScript/TypeScript 版本**中可用

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| prompt | **必需**, 生成视频的描述。 | 字符串, 最大支持2000字符 | 无 |
| model | 请求的模型版本 | 字符串, 可选范围 ["MiniMax-Hailuo-02", "T2V-01-Director", "I2V-01-Director", "S2V-01", "I2V-01-live", "I2V-01", "T2V-01"] | "T2V-01" |
| first_frame_image | **必需**, 视频的首帧画面 | 支持传入图片的 "data:image/jpeg;base64,{data}" 格式的 Base64 编码字符串, 或可通过公网访问的 URL | 无 |
| output_directory | 保存视频文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |
| async_mode | 是否使用异步模式, 启用则返回任务ID | 布尔值 (True / False) | FALSE |

<br>

### 9. query_video_generation
该工具用于查询异步视频生成任务的状态。

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| task_id | **必需**, 需要查询的任务 ID | 字符串 | 无 |
| output_directory | 保存视频文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |

### 10. text_to_image
该工具可根据指定提示词生成图片

| 参数 | 含义 | 格式及说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| prompt | **必需**, 生成图像的描述 | 字符串, 最大支持 1500 字符 | 无 |
| model | 请求的模型版本 | 字符串, 可选范围 ["image-01", "image-01-live"] | "image-01" |
| aspect_ratio | 生成图像的宽高比 | 字符串, 可选范围 ["1:1", "16:9", "4:3", "3:2", "2:3", "3:4", "9:16", "21:9"] | "1:1" |
| n | 单次请求生成的图片数量 | 整数, 可选范围 | 1 |
| prompt_optimizer | 是否开启提示词自动优化提示词 | 布尔值 (True / False) | TRUE |
| output_directory | 保存图片文件的目录 | 字符串 (文件路径) | 配置文件中设置的路径 |

## 在客户端使用MiniMax MCP服务

### 获取API Key

*   访问 [MiniMax 开放平台](https://platform.minimaxi.com/user-center/basic-information/interface-key)
*   点击“**Create new secret key**”按钮，输入项目名称以创建新的 API Key。
*   创建成功后，系统将展示 API Key。**请务必复制并妥善保存**，该密钥**只会显示一次**，无法再次查看。

<img src="https://filecdn.minimax.chat/public/0b03e12d-191d-45f4-826a-5e997fa845c2.png" width="60%" />

### UVX安装和配置

**MiniMax‑MCP** 是一个 Python 实现的 MCP 服务，为了让 MCP 客户端顺利调用，该服务必须通过 `uvx` 启动和执行。 `uvx` 和 `uv` 提供的命令行工具，类似于 `npm exec` ，用于运行包内部定义的可执行程序，确保环境隔离与依赖可控。

1.  安装 `uv` (以便获取 `uvx` ）
    *   macOS / Linux 用户：

```python
curl -LsSf https://astral.sh/uv/install.sh | sh
```

*   Windows 用户：

```python
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

其他安装方式可参考 [uv 仓库](https://github.com/astral-sh/uv)。安装完成后，会在 Python 环境的 `Scripts` 或 `bin` 目录中生成 `uv` 和 `uvx` 可执行文件。

2.  验证 `uvx` 是否可用

执行命令：
*   macOS / Linux 用户：

```python
which uvx
```

*   Windows 用户：

```python
(Get-Command uvx).source
```

*   若正确安装，会显示路径（如 `/usr/local/bin/uvx` ）。
*   若系统报错 `spawn uvx ENOENT` ，说明没有安装 `uvx` 或 `uvx` 不在系统路径中，需要配置绝对路径。

### 传输方式说明

MiniMax-MCP提供 studio 和 SSE 两种传输方式，在使用时可按需选择

| 特性 | stdio（默认） | SSE（Server‑Sent Events） |
| :--- | :--- | :--- |
| 运行环境 | 本地部署运行 | 本地部署或云端部署均可 |
| 通信方式 | 通过 stdout 进行通信 | 通过网络通信 |
| 适用场景 | 本地MCP客户端集成 | 需要服务器推送的应用 |
| 输入资源支持 | 支持本地文件或有效的 URL 资源 | 支持本地文件或 URL；部署在云端时推荐使用 URL 输入 |

### 在 Claude Desktop 中使用

1.  在 [Claude 官网](https://claude.ai/download)下载 Claude Desktop
2.  前往 `Claude > Settings > Developer > Edit Config > claude_desktop_config.json` ，添加以下配置。完成配置后，重启 Claude Desktop。

注意：如果使用 Windows，需要在 Claude Desktop 中启用"开发者模式"才能使用 MCP 服务器。
若出现报错 `spawn uvx ENOENT` ，请在 `command` 中配置 `uvx` 的绝对路径

```json
{
  "mcpServers": {
    "MiniMax": {
      "command": "uvx",
      "args": [
        "minimax-mcp"
      ],
      "env": {
        "MINIMAX_API_KEY": "填写你的API密钥",
        "MINIMAX_MCP_BASE_PATH": "本地输出目录路径，如/User/xxx/Desktop",
        "MINIMAX_API_HOST": "填写API Host, https://api.minimaxi.com 或 https://api.minimax.io",
        "MINIMAX_API_RESOURCE_MODE": "可选配置，资源生成后的提供方式, 可选项为 [url|local], 默认为 url"
      },
      "transport": "可选配置，传输方式，可选项为 [studio|SSE]，默认为 studio"   
    }
  }
}
```

### 在 Cursor 中使用

1.  通过 [Cursor官网](https://cursor.com/) 下载并安装 Cursor
2.  前往 `Cursor -> Preferences -> Cursor Settings -> Tools & Inrgrations -> MCP -> Add Custom MCP` ，打开 MCP 工具配置文件

<img src="https://filecdn.minimax.chat/public/61982fde-6575-4230-94eb-798f35a60450.png" width="60%" />

3.  在 `mcp.json` 文件中，增加 MiniMax 账户配置信息

```json
{
  "mcpServers": {
    "MiniMax": {
      "command": "uvx",
      "args": [
        "minimax-mcp"
      ],
      "env": {
        "MINIMAX_API_KEY": "填写你的API密钥",
        "MINIMAX_MCP_BASE_PATH": "本地输出目录路径，如/User/xxx/Desktop，需要保证路径存在且具有写入权限",
        "MINIMAX_API_HOST": "填写API Host, https://api.minimaxi.com 或 https://api.minimax.io",
        "MINIMAX_API_RESOURCE_MODE": "可选配置，资源生成后的提供方式, 可选项为 [url|local], 默认为 url"
      },
      "transport": "可选配置，传输方式，可选项为 [studio|SSE]，默认为 studio" 
    }
  }
}
```

### 在 Cherry Studio 中使用

1.  通过 [Cherry Studio 官网](https://www.cherry-ai.com/) 下载客户端
2.  前往 `Settings -> MCP Settings -> Add Server -> Import from JSON` ，将以下代码粘贴到代码框中，确认

```json
{
      "name": "minimax-mcp",
      "isActive": true,
      "command": "uvx",
      "args": [
        "minimax-mcp"
      ],
      "env": {
        "MINIMAX_API_KEY": "填写你的API密钥",
        "MINIMAX_MCP_BASE_PATH": "本地输出目录路径，如/User/xxx/Desktop，需要保证路径存在且具有写入权限",
        "MINIMAX_API_HOST": "填写API Host, https://api.minimaxi.com 或 https://api.minimax.io",
        "MINIMAX_API_RESOURCE_MODE": "可选配置，资源生成后的提供方式, 可选项为 [url|local], 默认为 url"
      }
      "transport": "可选配置，传输方式，可选项为 [studio|SSE]，默认为 studio"  
    }
```

3.  在对话框中，点击  `MCP Settings` 后，选择完成配置的“MiniMax MCP”即可使用

<img src="https://filecdn.minimax.chat/public/10883349-c9ff-4c92-b1d7-bccf15a5f28f.png" width="70%" />

## 在客户端使用MiniMax MCP JS服务器

### 获取 API Key

*   访问 [MiniMax开放平台](https://www.minimax.io/platform/user-center/basic-information/interface-key)
*   点击“**Create new secret key**”按钮，输入项目名称以创建新的 API Key。
*   创建成功后，系统将展示 API Key。**请务必复制并妥善保存**，该密钥**只会显示一次**，无法再次查看。

<img src="https://filecdn.minimax.chat/public/155646cf-6e99-4ecd-8843-8232668d5f67.png" width="60%" />

### Node.js 与 npm 安装

Node.js 是一个开源的 JavaScript 运行时环境，可以在浏览器之外运行 JavaScript 代码。它基于 Google 的 V8 引擎，具有高性能、事件驱动和非阻塞 I/O 的特点，适合构建高并发的网络服务、实时应用和微服务等场景。
npm 是随 Node.js 一起安装的默认包管理器，也是全球最大的软件注册中心。开发者可以通过 npm 搜索、安装、更新和管理依赖包（包括前端和后端代码模块），大幅简化开发流程。

1.  [安装 Node.js 与 npm](https://nodejs.org/en/download)
2.  **验证安装是否完成**

执行以下命令，若正确安装，会显示 Node.js 和 npm 的版本

```python
node -v
npm -v
```

### 传输方式说明

MiniMax-MCP-JS提供 studio、REST 和 SSE 三种传输方式，在使用时可按需选择

| 特性 | stdio (默认) | REST | SSE |
| :--- | :--- | :--- | :--- |
| 运行环境 | 本地运行 | 可本地或云端部署 | 可本地或云端部署 |
| 通信方式 | 通过标准输入输出通信 | 通过HTTP请求通信 | 通过服务器发送事件通信 |
| 适用场景 | 本地MCP客户端集成 | API服务，跨语言调用 | 需要服务器推送的应用 |
| 输入限制 | 支持处理本地文件或有效的URL资源 | 当部署在云端时，建议使用URL作为输入 | 当部署在云端时，建议使用URL作为输入 |

### 在Claude Desktop 中使用

1.  在 Claude 官网下载 [Claude Desktop](https://claude.ai/download)
2.  前往 `Claude > Settings > Developer > Edit Config > claude_desktop_config.json` ，添加以下配置。完成配置后，重启 Claude Desktop。

注意：如果使用Windows，需要在Claude Desktop中启用"开发者模式"才能使用MCP服务器。

```json
{
  "mcpServers": {
    "minimax-mcp-js": {
      "command": "npx",
      "args": [
        "-y",
        "minimax-mcp-js"
      ],
      "env": {
        "MINIMAX_API_HOST": "https://api.minimax.chat",
        "MINIMAX_API_KEY": "<您的API密钥>",
        "MINIMAX_MCP_BASE_PATH": "<本地输出目录路径，如/User/xxx/Desktop>",
        "MINIMAX_RESOURCE_MODE": "<可选配置，资源生成后的提供方式, [url|local], 默认为 url>"
      },
      "transport": "可选配置，传输方式，可选项为 [studio|REST|SSE]，默认为 studio"  
    }
  }
}
```

### 在 Cursor 中使用

1.  通过 [Cursor官网](https://cursor.com/) 下载并安装 Cursor
2.  前往 `Cursor -> Preferences -> Cursor Settings -> Tools & Inrgrations -> MCP -> Add Custom MCP` ，打开 MCP 工具配置文件

<img src="https://filecdn.minimax.chat/public/61982fde-6575-4230-94eb-798f35a60450.png" width="60%" />

3.  在 `mcp.json` 文件中，增加 MiniMax 账户配置信息

```json
{
  "mcpServers": {
    "MiniMax": {
      "command": "uvx",
      "args": [
        "minimax-mcp"
      ],
      "env": {
        "MINIMAX_API_KEY": "填写你的API密钥",
        "MINIMAX_MCP_BASE_PATH": "本地输出目录路径，如/User/xxx/Desktop，需要保证路径存在且具有写入权限",
        "MINIMAX_API_HOST": "填写API Host, https://api.minimaxi.com 或 https://api.minimax.io",
        "MINIMAX_API_RESOURCE_MODE": "可选配置，资源生成后的提供方式, 可选项为 [url|local], 默认为 url"
      },
      "transport": "可选配置，传输方式，可选项为 [studio|REST|SSE]，默认为 studio" 
    }
  }
}
```

4.  完成配置后，可以查看 MiniMax 目前支持的 mcp 工具

<img src="https://filecdn.minimax.chat/public/72ceca39-5edc-4f18-84cc-599715d3475f.png" width="40%" />

### 在 Cherry Studio 中使用

1.  通过 [Cherry Studio 官网](https://www.cherry-ai.com/) 下载客户端
2.  前往 `Settings -> MCP Settings -> Add Server -> Import from JSON` ，将以下代码粘贴到代码框中，确认

```json
{
      "name": "minimax-mcp",
      "isActive": true,
      "command": "npx",
      "args": [
        "y",
        "minimax-mcp"
      ],
      "env": {
        "MINIMAX_API_KEY": "填写你的API密钥",
        "MINIMAX_MCP_BASE_PATH": "本地输出目录路径，如/User/xxx/Desktop，需要保证路径存在且具有写入权限",
        "MINIMAX_API_HOST": "填写API Host, https://api.minimaxi.com 或 https://api.minimax.io",
        "MINIMAX_API_RESOURCE_MODE": "可选配置，资源生成后的提供方式, 可选项为 [url|local], 默认为 url"
      },
      "transport": "可选配置，传输方式，可选项为 [studio|REST|SSE]，默认为 studio" 
    }
```

3.  在对话框中，点击 `MCP Settings` 后，选择完成配置的“MiniMax MCP”即可使用

<img src="https://filecdn.minimax.chat/public/10883349-c9ff-4c92-b1d7-bccf15a5f28f.png" width="70%" />

## MiniMax MCP 使用示例

### 音频工具使用

1.  选择合适的声音信息，播报晚间新闻片段

参考提示词

```python
choose a voice, and broadcast a segment of the evening news
```

生成内容


<audio
  src="https://filecdn.minimax.chat/public/87394f73-1ed6-4507-8eb6-d1e7abb65e23.mp3"
  width="100%"
  title="Good_eveni">
</audio>



思考过程

<img src="https://filecdn.minimax.chat/public/e9f12d7f-e7d8-4f3d-b273-181a51433606.png" width="50%" />

2.  根据指定音频克隆声音，并指定克隆音色的id

参考提示词

```python
clone the voice from the audio file named Marketing_Voice.sav, the id is test_vlone_voice
```

来源音频

<audio src="https://filecdn.minimax.chat/public/cd3c2e0c-8cb8-48ff-999c-96d5f070765f.wav" width="100%"
  title="来源音频"></audio>

结果音频

<audio src="https://filecdn.minimax.chat/public/678a8b1c-fd02-48a5-b250-e225119d9448.mp3" width="100%"
  title="结果音频"></audio>

思考过程

<img src="https://filecdn.minimax.chat/public/3e73f8ca-39e6-420f-8b90-8773d4db6e14.png" width="60%" />

3.  按照要求设计音色，并给定示例文本生成音频

参考提示词

```python
Design a voice, the requirement is "Mysterious narrator with a deep, magnetic voice, suspenseful tone, moderate pace, subtle reverb". Then use it in the sample Text: "In the shadows of the old manor, secrets whisper through the walls. Beware what you seek…"
```

生成内容

<audio src="https://filecdn.minimax.chat/public/98488c0c-c9e5-47f0-bab0-3b1b3f4dae3a.mp3" title="t2a_in_the_sha" width="100%"></audio>

思考过程

<img src="https://filecdn.minimax.chat/public/a44cade5-1f33-4d54-a70a-2fec83b73a0a.png" width="50%" />

### 音乐生成工具使用

参考提示词

```python
generate a song, the background music is gentle ambient piano and warm pad synth, soft reverb and subtle field recordings of wind chimes. The musical style: calm and reflective
Lyrics:
‘In the stillness of the midnight air,
Find the echoes of dreams we share.
Softly drifting ‘neath pale moonlight,
Whispering hearts drifting into night.’"
```

生成内容

<audio src="https://filecdn.minimax.chat/public/737fa5dc-0ecf-4fd6-9821-197e9490d3e3.mp3" title="music_gentle"></audio>

思考过程

<img src="https://filecdn.minimax.chat/public/a53e06e9-e10e-46d4-82b2-2c829261709c.png" width="70%" />

### 图片生成工具使用

参考提示词

```python
generate a hyperreal style picture, the requirement is "Ultra‑detailed digital painting of a serene mountain lake at sunrise, ultra-realistic, soft golden light, mist over the water"
```

生成内容

<img src="https://filecdn.minimax.chat/public/09b2c787-8c37-4f80-94b5-c1855ce05e97.jpeg" width="80%" />

思考过程

<img src="https://filecdn.minimax.chat/public/2017abf5-8168-43dd-b938-fb60179d3562.png" width="50%" />

### 视频生成工具使用

参考提示词及图片

```python
From the existing image of a kitten perched on a diving board, create a short video showing the kitten crouching, leaping off into the pool, and making a small splash—adorable and playful. Use MiniMax-Hailuo-02 model, and resolution is 1080P
```
<br>
<img src="https://filecdn.minimax.chat/public/eeee4d8d-7834-4370-bd55-e5579e739a3e.jpeg" />

生成内容

<video
  src="https://filecdn.minimax.chat/public/515dec59-459c-4560-8557-00a81e630138.mp4"
  cover=""
  width="100%"
  ratio="16:9">
</video>

思考过程

<img src="https://filecdn.minimax.chat/public/02b88e6b-63b1-4989-a821-285809d45eae.png" width="50%" />

## 如何做出自己的贡献

若希望对 MiniMax MCP 项目进行改进或修复错误，欢迎通过以下方式提交建议或代码贡献：

1.  在 GitHub 项目主页（ [Python 版](https://github.com/MiniMax-AI/MiniMax-MCP/issues)或 [JS 版](https://github.com/MiniMax-AI/MiniMax-MCP-JS/issues) ）开一个新的 Issue，简要描述您建议的更改或问题。
2.  获取反馈后，请按照项目贡献指南创建一个对应的 **Pull Request (PR)** ，附上修改说明及必要的背景信息。
3.  项目维护者会对你的 PR 进行代码审查，并给予合并建议或进一步修改意见。
```

## 音乐生成指南
- 原始链接: https://platform.minimaxi.com/document/guides_music_generation?key=68bbf0a44b8efaab4d9f8794
- 文档ID: `68bbf0a44b8efaab4d9f8794`
# 音乐生成

音乐生成接口能根据文本描述和歌词，直接生成一段包含人声的完整歌曲。通过 `prompt` 参数定义音乐的风格、情绪和场景，通过 `lyrics` 参数提供演唱的歌词内容。该功能可用于为视频、游戏或应用快速生成独特的背景音乐和主题曲。

##  音乐生成示例

### 请求示例

```python
import requests
import os

url = "https://api.minimaxi.com/v1/music_generation"
api_key = os.environ.get("MINIMAX_API_KEY")
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "music-1.5",
    "prompt": "独立民谣,忧郁,内省,渴望,独自漫步,咖啡馆",
    "lyrics": "[verse]\n街灯微亮晚风轻抚\n影子拉长独自漫步\n旧外套裹着深深忧郁\n不知去向渴望何处\n[chorus]\n推开木门香气弥漫\n熟悉的角落陌生人看",
    "audio_setting": {"sample_rate": 44100, "bitrate": 256000, "format": "mp3"},
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
audio_hex = response.json()["data"]["audio"]

with open("output.mp3", "wb") as f:
    f.write(bytes.fromhex(audio_hex))
```

### 模型生成结果
<br>
**音乐生成结果**: https://filecdn.minimax.chat/public/a34192c9-76e2-43d9-b0b1-307b2763a377.mp3

## 异步语音合成指南
- 原始链接: https://platform.minimaxi.com/document/guides_t2a_async?key=68bbecb495a600dbfc925d35
- 文档ID: `68bbecb495a600dbfc925d35`
# 异步语音合成指南

MiniMax 提供异步语音合成 API，适用于长文本的音频合成任务，单个文件长度限制小于10万字符。
1. 支持100+系统音色、复刻音色自主选择
2. 支持语调、语速、音量、比特率、采样率、输出格式调整
3. 支持音频时长、音频大小等返回参数
4. 支持时间戳（字幕）返回，精确到句
5. 支持直接传入字符串与上传文本文件两种方式进行待合成文本的输入
6. 支持非法字符检测：非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；非法字符超过10%，接口不返回结果（返回报错码），请检测后再次进行请求【非法字符定义：ascii码中的控制符（不含制表符 `\t` 和换行符`\n`）】

## 支持模型
以下为 MiniMax 已提供的语音模型及其特性说明。
| 模型 | 特性 |
| ------ | ------ |
| speech-2.5-hd-preview | 最新的 HD 模型，韵律表现出色，复刻相似度极佳 |
| speech-2.5-turbo-preview | 最新的 Turbo 模型，支持 40 个语种 |
| speech-02-hd | 拥有出色的韵律、稳定性和复刻相似度，音质表现突出 |
| speech-02-turbo | 拥有出色的韵律和稳定性，小语种能力加强，性能表现出色 |
| speech-01-hd | 拥有超高的复刻相似度，音质表现突出 |
| speech-01-turbo | 在出色的生成效果基础上有更快的生成速度 |

## 支持语言
MiniMax 的语音合成模型具备卓越的跨语言能力，全面支持 40 种全球广泛使用的语言。我们致力于打破语言壁垒，构建真正意义上的全球通用人工智能模型。

目前支持的语言包含：
| 序号 | 中文 | EN |
|---|---|---|
| 1 | 中文 | Chinese |
| 2 | 粤语 | Cantonese |
| 3 | 英语 | English |
| 4 | 西班牙语 | Spanish |
| 5 | 法语 | French |
| 6 | 俄语 | Russian |
| 7 | 德语 | German |
| 8 | 葡萄牙语 | Portuguese |
| 9 | 阿拉伯语 | Arabic |
| 10 | 意大利语 | Italian |
| 11 | 日语 | Japanese |
| 12 | 韩语 | Korean |
| 13 | 印尼语 | Indonesian |
| 14 | 越南语 | Vietnamese |
| 15 | 土耳其语 | Turkish |
| 16 | 荷兰语 | Dutch |
| 17 | 乌克兰语 | Ukrainian |
| 18 | 泰语 | Thai |
| 19 | 波兰语 | Polish |
| 20 | 罗马尼亚语 | Romanian |
| 21 | 希腊语 | Greek |
| 22 | 捷克语 | Czech |
| 23 | 芬兰语 | Finnish |
| 24 | 印地语 | Hindi |
| 25 | 保加利亚语 | Bulgarian |
| 26 | 丹麦语 | Danish |
| 27 | 希伯来语 | Hebrew |
| 28 | 马来语 | Malay |
| 29 | 波斯语 | Persian |
| 30 | 斯洛伐克语 | Slovak |
| 31 | 瑞典语 | Swedish |
| 32 | 克罗地亚语 | Croatian |
| 33 | 菲律宾语 | Filipino |
| 34 | 匈牙利语 | Hungarian |
| 35 | 挪威语 | Norwegian |
| 36 | 斯洛文尼亚语 | Slovenian |
| 37 | 加泰罗尼亚语 | Catalan |
| 38 | 尼诺斯克语 | Nynorsk |
| 39 | 泰米尔语 | Tamil |
| 40 | 阿非利卡语 | Afrikaans |s)
<br/>

## 使用流程
1. 若使用文件输入，需先调用 [文件上传 API](/document/files_upload?key=68b2aedbfd643cbdfcc0611b) 上传文本并获取file_id。若使用文本作为输入，则跳过此步骤
2. 调用[创建语音生成任务 API](/document/t2a_async_create?key=68adac886602726333001546)，获取 `task_id`
3. 调用[查询语音生成任务状态 API](/document/t2a_async_query?key=68adad0f6fe587e3fbfe9810)，基于 `task_id` 获取语音合成任务进度
4. 当任务完成时，上述调用查询语音生成任务状态 API 返回的 `file_id` 可用于调用 [文件下载 API](/document/files_retrieve_content?key=68b56ae2a96516e2601920d2) 下载音频结果

注意：返回的下载 URL 自生成起 9 小时（32400 秒）内有效，过期后文件将失效，生成的信息便会丢失，请注意下载信息的时间，及时下载
<br/>
## 过程示例
### 1. 获取 file_id

<codetabs title="获得file_id">
<codetab  language="python">
"""
本示例用于待合成文本的 file_id。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import os

api_key = os.environ.get("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/files/upload"

payload = {'purpose': 't2a_async_input'}
files=[
  ('file',('21622420008680404.zip',open('/Users/minimax/Downloads/21622420008680404.zip','rb'),'application/zip'))
]
headers = {
  'authority': 'api.minimaxi.com',
  'Authorization': f'Bearer {api_key}'
}

response = requests.request("POST", url, headers=headers, data=payload, files=files)

print(response.text)

</codetab>
<codetab  language="curl">
curl --location 'https://api.minimaxi.com/v1/files/upload' \
  --header 'authority: api.minimaxi.com' \
  --header "Authorization: Bearer $MINIMAX_API_KEY" \
  --form 'purpose=t2a_async_input' \
  --form 'file=@test-json.zip'

</codetab>
</codetabs>

<br/>


### 2. 创建语音合成任务

<codetabs title="创建语音合成任务">
<codetab  language="python">
"""
本示例用于创建语音合成任务，若使用文件作为输入，则需要将<text_file_id>替换为文本文件的file_id，若使用文本作为输入，则设置“text”字段。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import json
import os

api_key = os.environ.get("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/t2a_async_v2"

payload = json.dumps({
  "model": "speech-2.5-hd-preview",
  "text_file_id": <text_file_id>, # file as input
  # "text":"微风拂过柔软的草地，清新的芳香伴随着鸟儿的歌唱。", # text as input
  "language_boost": "auto",
  "voice_setting": {
    "voice_id": "audiobook_male_1",
    "speed": 1,
    "vol": 10,
    "pitch": 1
  },
  "pronunciation_dict": {
    "tone": [
      "草地/(cao3)(di1)"
    ]
  },
  "audio_setting": {
    "audio_sample_rate": 32000,
    "bitrate": 128000,
    "format": "mp3",
    "channel": 2
  },
    "voice_modify":{
      "pitch":0,
      "intensity":0,
      "timbre":0,
      "sound_effects":"spacious_echo"
    }
})
headers = {
  'Authorization': f'Bearer {api_key}',
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)



</codetab>
<codetab  language="curl">
# 若使用文件作为输入，则需要将<text_file_id>替换为文本文件的file_id，若使用文本作为输入，则设置“text”字段。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
curl --location 'https://api.minimaxi.com/v1/t2a_async_v2' \
--header "authorization: Bearer ${MINIMAX_API_KEY}" \
--header 'Content-Type: application/json' \
--data '{
  "model": "speech-2.5-hd-preview",
  "text_file_id": <Your file_id>, # file as input
  # "text":"微风拂过柔软的草地，清新的芳香伴随着鸟儿的歌唱。", # text as input
  "language_boost": "auto",
  "voice_setting": {
    "voice_id": "audiobook_male_1",
    "speed": 1,
    "vol": 10,
    "pitch": 1
  },
  "pronunciation_dict": {
    "tone": [
      "草地/(cao3)(di1)"
    ]
  },
  "audio_setting": {
    "audio_sample_rate": 32000,
    "bitrate": 128000,
    "format": "mp3",
    "channel": 2
  },
    "voice_modify":{
      "pitch":0,
      "intensity":0,
      "timbre":0,
      "sound_effects":"spacious_echo"
    }
}'

</codetab>
</codetabs>
<br/>


### 3. 查询语音合成进度
<codetabs title="查询语音合成进度">
<codetab  language="python">
"""
本示例用于查询语音合成进度。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`，并将需要查询任务的 id 写入环境变量 `TASK_ID`。
"""
import requests
import json
import os

task_id = os.environ.get("TASK_ID")  
api_key = os.environ.get("MINIMAX_API_KEY")  
url = f"https://api.minimaxi.com/v1/query/t2a_async_query_v2?task_id={task_id}"

payload = {}
headers = {
  'Authorization': f'Bearer {api_key}',
  'content-type': 'application/json',
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)


</codetab>
<codetab  language="curl">
curl --location "https://api.minimaxi.com/v1/query/t2a_async_query_v2?task_id=${TASK_ID}" \
--header "authorization: Bearer ${MINIMAX_API_KEY}" \
--header 'content-type: application/json' 


</codetab>
</codetabs>
<br/>

### 4. 下载语音合成文件

<codetabs title="下载语音合成文件">
<codetab  language="python">
"""
本示例用于下载语音合成文件。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`，并将待下载文件的 id 写入环境变量 `FILE_ID`。
"""
import requests
import os

api_key = os.environ.get("MINIMAX_API_KEY")
file_id = os.environ.get("FILE_ID")

url = f"https://api.minimaxi.com/v1/files/retrieve_content?file_id={file_id}"

payload = {}
headers = {
  'content-type': 'application/json',
  'Authorization': f'Bearer {api_key}'
}

response = requests.request("GET", url, headers=headers, data=payload)

with open(<output_filename>, 'wb') as f:
    f.write(response.content)


</codetab>
<codetab  language="curl">
curl --location "https://api.minimaxi.com/v1/files/retrieve?file_id=${FILE_ID}" \
  --header 'Content-Type: application/json' \
  --header "Authorization: Bearer ${MINIMAX_API_KEY}" \
  --output "${FILE_NAME}" \


</codetab>
</codetabs>

## 同步语音合成指南
- 原始链接: https://platform.minimaxi.com/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba
- 文档ID: `68bbebb795a600dbfc9259ba`
# 同步语音合成指南

同步语音合成支持基于文本到语音的同步生成，单次可处理最长 10,000 字符的文本。
## 支持模型
以下为 MiniMax 提供的语音模型及其特性说明。
| 模型 | 特性 |
| ------ | ------ |
| speech-2.5-hd-preview | 最新的 HD 模型，韵律表现出色，复刻相似度极佳 |
| speech-2.5-turbo-preview | 最新的 Turbo 模型，支持 40 个语种 |
| speech-02-hd | 拥有出色的韵律、稳定性和复刻相似度，音质表现突出 |
| speech-02-turbo | 拥有出色的韵律和稳定性，小语种能力加强，性能表现出色 |
| speech-01-hd | 拥有超高的复刻相似度，音质表现突出 |
| speech-01-turbo | 在出色的生成效果基础上有更快的生成速度 |


## 支持语言
MiniMax 的语音合成模型具备卓越的跨语言能力，全面支持 40 种全球广泛使用的语言。我们致力于打破语言壁垒，构建真正意义上的全球通用人工智能模型。

目前支持的语言包含：

| 序号 | 中文 | EN |
|---|---|---|
| 1 | 中文 | Chinese |
| 2 | 粤语 | Cantonese |
| 3 | 英语 | English |
| 4 | 西班牙语 | Spanish |
| 5 | 法语 | French |
| 6 | 俄语 | Russian |
| 7 | 德语 | German |
| 8 | 葡萄牙语 | Portuguese |
| 9 | 阿拉伯语 | Arabic |
| 10 | 意大利语 | Italian |
| 11 | 日语 | Japanese |
| 12 | 韩语 | Korean |
| 13 | 印尼语 | Indonesian |
| 14 | 越南语 | Vietnamese |
| 15 | 土耳其语 | Turkish |
| 16 | 荷兰语 | Dutch |
| 17 | 乌克兰语 | Ukrainian |
| 18 | 泰语 | Thai |
| 19 | 波兰语 | Polish |
| 20 | 罗马尼亚语 | Romanian |
| 21 | 希腊语 | Greek |
| 22 | 捷克语 | Czech |
| 23 | 芬兰语 | Finnish |
| 24 | 印地语 | Hindi |
| 25 | 保加利亚语 | Bulgarian |
| 26 | 丹麦语 | Danish |
| 27 | 希伯来语 | Hebrew |
| 28 | 马来语 | Malay |
| 29 | 波斯语 | Persian |
| 30 | 斯洛伐克语 | Slovak |
| 31 | 瑞典语 | Swedish |
| 32 | 克罗地亚语 | Croatian |
| 33 | 菲律宾语 | Filipino |
| 34 | 匈牙利语 | Hungarian |
| 35 | 挪威语 | Norwegian |
| 36 | 斯洛文尼亚语 | Slovenian |
| 37 | 加泰罗尼亚语 | Catalan |
| 38 | 尼诺斯克语 | Nynorsk |
| 39 | 泰米尔语 | Tamil |
| 40 | 阿非利卡语 | Afrikaans |


## 流式请求示例

本指南指导，流式播放返回的音频文件，并保存完整音频文件。

⚠️ 注意，为实时播放音频流，需要先安装 [mpv 播放器](https://mpv.io/installation/)。并且，需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。

请求示例

<codetabs title="请求示例">
<codetab  language="python">

import asyncio
import websockets
import json
import ssl
import subprocess
import os

model = "speech-2.5-hd-preview"
file_format = "mp3"

class StreamAudioPlayer:
    def __init__(self):
        self.mpv_process = None
        
    def start_mpv(self):
        """Start MPV player process"""
        try:
            mpv_command = ["mpv", "--no-cache", "--no-terminal", "--", "fd://0"]
            self.mpv_process = subprocess.Popen(
                mpv_command,
                stdin=subprocess.PIPE,
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL,
            )
            print("MPV player started")
            return True
        except FileNotFoundError:
            print("Error: mpv not found. Please install mpv")
            return False
        except Exception as e:
            print(f"Failed to start mpv: {e}")
            return False
    
    def play_audio_chunk(self, hex_audio):
        """Play audio chunk"""
        try:
            if self.mpv_process and self.mpv_process.stdin:
                audio_bytes = bytes.fromhex(hex_audio)
                self.mpv_process.stdin.write(audio_bytes)
                self.mpv_process.stdin.flush()
                return True
        except Exception as e:
            print(f"Play failed: {e}")
            return False
        return False
    
    def stop(self):
        """Stop player"""
        if self.mpv_process:
            if self.mpv_process.stdin and not self.mpv_process.stdin.closed:
                self.mpv_process.stdin.close()
            try:
                self.mpv_process.wait(timeout=20)
            except subprocess.TimeoutExpired:
                self.mpv_process.terminate()

async def establish_connection(api_key):
    """Establish WebSocket connection"""
    url = "wss://api.minimaxi.com/ws/v1/t2a_v2"
    headers = {"Authorization": f"Bearer {api_key}"}

    ssl_context = ssl.create_default_context()
    ssl_context.check_hostname = False
    ssl_context.verify_mode = ssl.CERT_NONE

    try:
        ws = await websockets.connect(url, additional_headers=headers, ssl=ssl_context)
        connected = json.loads(await ws.recv())
        if connected.get("event") == "connected_success":
            print("Connection successful")
            return ws
        return None
    except Exception as e:
        print(f"Connection failed: {e}")
        return None

async def start_task(websocket):
    """Send task start request"""
    start_msg = {
        "event": "task_start",
        "model": model,
        "voice_setting": {
            "voice_id": "male-qn-qingse",
            "speed": 1,
            "vol": 1,
            "pitch": 0,
            "english_normalization": False
        },
        "audio_setting": {
            "sample_rate": 32000,
            "bitrate": 128000,
            "format": file_format,
            "channel": 1
        }
    }
    await websocket.send(json.dumps(start_msg))
    response = json.loads(await websocket.recv())
    return response.get("event") == "task_started"

async def continue_task_with_stream_play(websocket, text, player):
    """Send continue request and stream play audio"""
    await websocket.send(json.dumps({
        "event": "task_continue",
        "text": text
    }))

    chunk_counter = 1
    total_audio_size = 0
    audio_data = b""
    
    while True:
        try:
            response = json.loads(await websocket.recv())
            
            if "data" in response and "audio" in response["data"]:
                audio = response["data"]["audio"]
                if audio:
                    print(f"Playing chunk #{chunk_counter}")
                    audio_bytes = bytes.fromhex(audio)
                    if player.play_audio_chunk(audio):
                        total_audio_size += len(audio_bytes)
                        audio_data += audio_bytes
                        chunk_counter += 1
            
            if response.get("is_final"):
                print(f"Audio synthesis completed: {chunk_counter-1} chunks")
                if player.mpv_process and player.mpv_process.stdin:
                    player.mpv_process.stdin.close()
                
                # Save audio to file
                with open(f"output.{file_format}", "wb") as f:
                    f.write(audio_data)
                print(f"Audio saved as output.{file_format}")
                    
                estimated_duration = total_audio_size * 0.0625 / 1000
                wait_time = max(estimated_duration + 5, 10)
                return wait_time
                
        except Exception as e:
            print(f"Error: {e}")
            break
    
    return 10

async def close_connection(websocket):
    """Close connection"""
    if websocket:
        try:
            await websocket.send(json.dumps({"event": "task_finish"}))
            await websocket.close()
        except Exception:
            pass

async def main():
    API_KEY = os.getenv("MINIMAX_API_KEY")
    TEXT = "真正的危险不是计算机开始像人一样思考，而是人开始像计算机一样思考。计算机只是可以帮我们处理一些简单事务。"

    player = StreamAudioPlayer()
    
    try:
        if not player.start_mpv():
            return
        
        ws = await establish_connection(API_KEY)
        if not ws:
            return

        if not await start_task(ws):
            print("Task startup failed")
            return

        wait_time = await continue_task_with_stream_play(ws, TEXT, player)
        await asyncio.sleep(wait_time)
        
    except Exception as e:
        print(f"Error: {e}")
    finally:
        player.stop()
        if 'ws' in locals():
            await close_connection(ws)

if __name__ == "__main__":
    asyncio.run(main())

</codetab>
</codetabs>

## 文本生成指南
- 原始链接: https://platform.minimaxi.com/document/guides_text_generation?key=68b9ba2cf7d629d07d2748fb
- 文档ID: `68b9ba2cf7d629d07d2748fb`
# 文本生成指南

您可以用 OpenAI 库来使用 MiniMax 模型进行文本生成。

## 文本生成

```python
from openai import OpenAI
import os

client = OpenAI(
    api_key=os.environ["MINIMAX_API_KEY"],
    base_url="https://api.minimaxi.com/v1",
)

messages = [
    {"role": "system", "name": "MiniMax AI"},
    {"role": "user", "name": "用户", "content": "你好"},
]

completion = client.chat.completions.create(
    model="MiniMax-M1",
    messages=messages,
    max_tokens=4096,
)

print(completion.choices[0].message)
```

## 流式生成

```python
from openai import OpenAI
import os

client = OpenAI(
    api_key=os.environ["MINIMAX_API_KEY"],
    base_url="https://api.minimaxi.com/v1",
)

messages = [
    {"role": "system", "name": "MiniMax AI"},
    {"role": "user", "name": "用户", "content": "你好"},
]

completion = client.chat.completions.create(
    model="MiniMax-M1",
    messages=messages,
    max_tokens=4096,
    stream=True
)

for chunk in completion:
    print(chunk.choices[0].delta)
```

## Function Calling

Function Calling 用法见 [开发者指南/函数调用](/document/guides_function_call?key=689c55a29561f5dd94811919)

## 图片理解

```python
import base64
from openai import OpenAI
import os

client = OpenAI(
    api_key=os.environ["MINIMAX_API_KEY"],
    base_url="https://api.minimaxi.com/v1",
)

messages = [
    {
        "role": "system",
        "name": "MiniMax AI",
        "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。",
    },
    {
        "role": "user",
        "name": "用户",
        "content": [
            {"type": "text", "text": "这个图代表的是什么呢"},
            {
                "type": "image_url",
                "image_url": {
                    "url": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg"
                },
            },
        ],
    },
]

completion = client.chat.completions.create(
    model="MiniMax-Text-01",
    messages=messages,
    max_tokens=4096,
)

print(completion.choices[0].message)
```

##  MiniMax Text 01 / M1 模型 Transformers 部署指南
- 原始链接: https://platform.minimaxi.com/document/guides_transformer_1?key=68b9be02f7d629d07d275b50
- 文档ID: `68b9be02f7d629d07d275b50`
# MiniMax Text01/M1 模型 Transformers 部署指南

本文档将帮助您使用 <a href="https://huggingface.co/docs/transformers/index" target="_blank">Transformers</a> 库部署和运行 MiniMax-M1 模型。Transformers 是业界广泛使用的机器学习模型定义框架，提供了丰富的预训练模型和灵活的接口。

## 本文档适用模型

本文档适用以下模型，只需在部署时修改模型名称即可。**注意 Transformers 适用的模型仓库名称带有 hf 后缀！** 与无 hf 后缀的模型相比，仅 config.json 文件存在差异，权重文件一致。以下以 MiniMax-M1-40k-hf 为例说明部署流程。



*   <a href="https://huggingface.co/MiniMaxAI/MiniMax-M1-40k-hf" target="_blank">MiniMaxAI/MiniMax-M1-40k-hf</a>
*   <a href="https://huggingface.co/MiniMaxAI/MiniMax-M1-80k-hf" target="_blank">MiniMaxAI/MiniMax-M1-80k-hf</a>
*   <a href="https://huggingface.co/MiniMaxAI/MiniMax-Text-01-hf" target="_blank">MiniMaxAI/MiniMax-Text-01-hf</a>


## 环境准备

*   Python：3.9+

建议使用虚拟环境（如 venv、conda、uv）以避免依赖冲突。请执行以下命令安装 Transformers、torch 及相关依赖。

```
# 使用 CUDA 12.8
# 使用 pip 安装
pip install transformers torch accelerate --extra-index-url https://download.pytorch.org/whl/cu128
# 或者使用 uv 安装
uv pip install transformers torch accelerate --torch-backend=auto
```

## 使用 Python 运行

请确保已正确安装所需依赖，并将 CUDA 驱动配置好。
以下代码演示如何使用 Transformers 加载并运行 MiniMax-M1 模型：

```python
from transformers import AutoModelForCausalLM, AutoTokenizer, GenerationConfig
import torch

MODEL_PATH = "MiniMaxAI/MiniMax-M1-40k-hf"
model = AutoModelForCausalLM.from_pretrained(
    MODEL_PATH,
    device_map="auto",
    trust_remote_code=True,
)
tokenizer = AutoTokenizer.from_pretrained(MODEL_PATH)

messages = [
    {"role": "user", "content": [{"type": "text", "text": "What is your favourite condiment?"}]},
    {"role": "assistant", "content": [{"type": "text", "text": "Well, I'm quite partial to a good squeeze of fresh lemon juice. It adds just the right amount of zesty flavour to whatever I'm cooking up in the kitchen!"}]},
    {"role": "user", "content": [{"type": "text", "text": "Do you have mayonnaise recipes?"}]}
]

model_inputs = tokenizer.apply_chat_template(messages, return_tensors="pt").to("cuda")

generated_ids = model.generate(model_inputs, max_new_tokens=100, do_sample=True)

response = tokenizer.batch_decode(generated_ids)[0]

print(response)
```

## 使用 Flash Attention 加速推理

Flash Attention 是一种高效的注意力机制实现，可以加速模型推理过程。需确保 GPU 支持 Flash Attention，部分老旧显卡可能不兼容。首先我们安装 flash_attn 包。

```
# 使用 pip 安装
pip install flash_attn --no-build-isolation
# 或者使用 uv 安装
uv pip install flash_attn --torch-backend=auto --no-build-isolation
```

使用 Flash Attention-2 加载并运行 MiniMax-M1 模型，只需在 from_pretrained 时增加如下参数：

```python
 model = AutoModelForCausalLM.from_pretrained(
     MODEL_PATH,
     device_map="auto",
     trust_remote_code=True,
     torch_dtype=torch.float16, # 新增参数
     attn_implementation="flash_attention_2"  # 新增参数
 )
```

## 获取支持

如果在部署 MiniMax 模型的过程中遇到任何问题，您可通过以下方式：

*   通过官方邮箱 [api@minimaxi.com](mailto:api@minimaxi.com) 联系我们的技术支持团队
*   在我们的 GitHub 仓库提交 <a href="https://github.com/MiniMax-AI/MiniMax-M1/issues" target="_blank">Issue</a>

我们将持续优化 Transformers 上的部署体验，欢迎您的反馈！

## 视频生成指南
- 原始链接: https://platform.minimaxi.com/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7
- 文档ID: `68bbf01a95a600dbfc9266f7`
# 视频生成指南

视频生成服务提供多种功能：

1. 文生视频：根据文本描述直接生成视频
1. 图生视频：基于一张初始图片结合文本描述生成视频
1. 首尾帧生成视频：提供视频开始、结束图片，来生成视频
1. 主体参考生成视频：基于一人脸照片，文本描述生成视频，视频中保持人物特征一致性

## 工作流程

视频生成是一个异步接口，包含以下三个步骤：

1. 创建生成任务：提交一个视频生成请求，获得任务 ID (`task_id`)
2. 查询任务状态：使用 `task_id` 轮询任务状态。任务成功后，会返回一个文件 ID (`file_id`)
3. 获取视频文件：使用 `file_id` 获取视频的下载地址并保存文件

## 功能与代码示例

为了简化代码，我们将轮询和下载的逻辑封装为公共函数，并举例了两种模式下如何创建任务。

```python
import os
import time
import requests

api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}


# --- 步骤 1: 发起视频生成任务 ---
# API 支持两种视频生成模式：纯文本生成（文生视频）和图像加文本生成（图生视频）。
# 以下两个函数分别对应这两种模式。它们都会发起一个异步的生成任务，并返回一个唯一的 task_id。

def invoke_text_to_video() -> str:
    """（模式一）通过文本描述发起视频生成任务。"""
    url = "https://api.minimaxi.com/v1/video_generation"
    payload = {
        # 'prompt' 是核心参数，用于描述视频的动态内容。
        "prompt": "男子拿起一本书[上升]，然后阅读[固定]。",
        "model": "MiniMax-Hailuo-02",
        "duration": 6,
        "resolution": "1080P",
    }
    response = requests.post(url, headers=headers, json=payload)
    response.raise_for_status()
    task_id = response.json()["task_id"]
    return task_id


def invoke_image_to_video() -> str:
    """（模式二）通过首帧图像和文本描述发起视频生成任务。"""
    url = "https://api.minimaxi.com/v1/video_generation"
    payload = {
        # 在图生视频模式下，'prompt' 用于描述基于首帧图像的动态变化。
        "prompt": "老鼠跑向镜头前，微笑着眨了眨眼睛。",
        # 'first_frame_image' 指定了视频的起始画面
        "first_frame_image": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg",
        "model": "MiniMax-Hailuo-02",
        "duration": 6,
        "resolution": "1080P",
    }
    response = requests.post(url, headers=headers, json=payload)
    response.raise_for_status()
    task_id = response.json()["task_id"]
    return task_id


# --- 步骤 2: 轮询查询任务状态 ---
# 视频生成是一个耗时过程，因此 API 设计为异步模式。
# 提交任务后，需使用 task_id 通过此函数进行轮询，以获取任务的最终状态。
def query_task_status(task_id: str):
    """根据 task_id 轮询任务状态，直至任务成功或失败。"""
    url = "https://api.minimaxi.com/v1/query/video_generation"
    params = {"task_id": task_id}
    while True:
        # 推荐的轮询间隔为 10 秒，以避免对服务器造成不必要的压力。
        time.sleep(10)
        response = requests.get(url, headers=headers, params=params)
        response.raise_for_status()
        response_json = response.json()
        status = response_json["status"]
        print(f"当前任务状态: {status}")
        # 任务成功时，API 会返回一个 'file_id'，用于下一步获取视频文件。
        if status == "Success":
            return response_json["file_id"]
        elif status == "Fail":
            raise Exception(f"视频生成失败: {response_json.get('error_message', '未知错误')}")


# --- 步骤 3: 获取并保存视频文件 ---
# 任务成功后，我们得到的是 file_id 而非直接的下载链接。
# 此函数首先使用 file_id 从文件服务获取下载 URL，然后下载视频内容并保存到本地。
def fetch_video(file_id: str):
    """根据 file_id 获取视频下载链接，并将其保存到本地。"""
    url = "https://api.minimaxi.com/v1/files/retrieve"
    params = {"file_id": file_id}
    response = requests.get(url, headers=headers, params=params)
    response.raise_for_status()
    download_url = response.json()["file"]["download_url"]
    
    with open("output.mp4", "wb") as f:
        video_response = requests.get(download_url)
        video_response.raise_for_status()
        f.write(video_response.content)
    print("视频已成功保存至 output.mp4")


# --- 主流程: 完整调用示例 ---
# 该部分演示了从发起任务到最终保存视频的完整调用链路。
if __name__ == "__main__":
    # 选择一种方式创建任务
    task_id = invoke_text_to_video()  # 方式一：文生视频
    # task_id = invoke_image_to_video() # 方式二：图生视频

    print(f"视频生成任务已提交，任务 ID: {task_id}")
    file_id = query_task_status(task_id)
    print(f"任务处理成功，文件 ID: {file_id}")
    fetch_video(file_id)
```
<br/>

## 生成视频结果
### 根据文本生成视频
通过 `prompt` 参数提供一段文本描述，即可生成相应的视频。为了实现对视频内容的精细控制，部分模型支持在 `prompt` 的关键描述后添加 `[运镜]` 指令，来控制镜头。

示例生成结果

<video
  src="https://filecdn.minimax.chat/public/f83e5873-23ac-43b7-9335-f2ff883c286c.mp4"
  cover=""
  ratio="16:9"
  width="100%"
  caption="文生视频"
  description="">
</video>


### 根据图片生成视频
该功能将 `first_frame_image` 参数指定的图片作为视频的起始帧，并结合 `prompt` 的描述生成后续的动态画面。这使得视频的开场画面完全可控，适合于让静态图片“动起来”的应用场景。

示例生成结果

<video
  src="https://filecdn.minimax.chat/public/65c8210a-0575-4f24-b89c-a09ca9a4075d.mp4"
  cover=""
  ratio="16:9"
  width="100%"
  caption="图生视频"
  description="">
</video>

## 使用模板视频生成指南
- 原始链接: https://platform.minimaxi.com/document/guides_video_template_generation?key=68bbf0494b8efaab4d9f85ec
- 文档ID: `68bbf0494b8efaab4d9f85ec`
# 使用模板视频生成指南

使用视频模板进行视频生成服务，可将图片或文本等素材填充至预设的视频模板中，快速生成风格统一的视频。

## 工作流程

视频模板生成是一个异步接口，流程如下：

1. 创建视频模板生成任务：根据指定的 `template_id` 和填充的素材，创建一个生成任务，获得 `task_id`
2. 查询任务并获取结果：使用 `task_id` 轮询任务状态。与通用视频生成不同，当此任务完成时，API 会在响应中直接返回可供下载的 `video_url`

更多的视频模板，可参考[视频模板列表](/document/video_agent_templates?key=68b401edd2f0aaaf3484a10e)。

## 生成“绝地求生”风格视频
### 示例代码

```python
import os
import time
import requests

api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}


# --- 步骤 1: 提交视频生成任务 ---
# 这个函数负责调用 API，根据指定的模板和素材，启动一个异步的视频生成任务。
# 任务提交成功后，API 会立即返回一个 task_id，用于后续查询任务状态。
def invoke_template_task() -> str:
    """提交一个基于模板的视频生成任务，并返回任务 ID"""
    url = "https://api.minimaxi.com/v1/video_template_generation"
    payload = {
        # 'template_id' 指定了视频的基础模板。
        "template_id": "393769180141805569",  # 示例：绝地求生
        # 填充模板中的图片、视频等媒体素材
        "media_inputs": [
            {
                "value": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg"
            }
        ],
        # 用于填充模板中的文本素材
        "text_inputs": [{"value": "狮子"}],
    }
    response = requests.post(url, headers=headers, json=payload)
    response.raise_for_status()
    task_id = response.json()["task_id"]
    return task_id


# --- 步骤 2: 轮询任务状态 ---
# 由于视频生成是异步的，我们需要通过上一步获取的 task_id 定期查询任务状态。
# 当状态变为 "Success"，函数会返回视频的 URL。如果失败，则会抛出异常。
def query_task_status(task_id: str):
    url = "https://api.minimaxi.com/v1/query/video_template_generation"
    params = {"task_id": task_id}
    while True:
        # 建议设置合理的轮询间隔，避免过于频繁的请求。
        time.sleep(10)
        response = requests.get(url, headers=headers, params=params)
        response.raise_for_status()
        response_json = response.json()
        status = response_json["status"]
        print(f"当前任务状态: {status}")
        if status == "Success":
            return response_json["video_url"]
        elif status == "Fail":
            raise Exception(f"视频生成失败: {response_json}")


# --- 步骤 3: 保存视频文件 ---
# 这是一个辅助函数，用于从给定的 URL 下载视频并保存到本地。
def save_video_from_url(video_url: str):
    print(f"正在从 {video_url} 下载视频...")
    response = requests.get(video_url)
    response.raise_for_status()
    with open("output.mp4", "wb") as f:
        f.write(response.content)
    print("视频已成功保存到 output.mp4")


# --- 主流程: 串联所有步骤 ---
# 按照“提交 -> 轮询 -> 保存”的顺序，执行视频生成的整个流程。
if __name__ == "__main__":
    task_id = invoke_template_task()
    print(f"已成功提交视频生成任务，task_id: {task_id}")
    final_video_url = query_task_status(task_id)
    print(f"任务成功，视频 URL: {final_video_url}")
    save_video_from_url(final_video_url)
```

### 生成结果

<video
  src="https://filecdn.minimax.chat/public/92ed8c3b-8173-4ed6-86e3-7860fecb2c7c.mp4"
  cover=""
  ratio="16:9"
  width="100%"
  caption="绝地求生"
  description="">
</video>

## MiniMax Text01/M1 模型 vLLM 部署指南
- 原始链接: https://platform.minimaxi.com/document/guides_vllm_m1?key=68b9b857630660297b06a603
- 文档ID: `68b9b857630660297b06a603`
# MiniMax Text01/M1 模型 vLLM 部署指南

我们推荐使用 vLLM 来部署 Minimax-M1 模型。vLLM 是一个高性能的推理引擎，其具有卓越的服务吞吐、高效智能的内存管理机制、强大的批量请求处理能力、深度优化的底层性能等特性。我们建议在部署之前查看 vLLM 的官方文档以检查硬件兼容性。

## 本文档适用模型

本文档适用以下模型，只需在部署时修改模型名称即可。以下以 MiniMax-M1-40k 为例说明部署流程。

* <a href="https://huggingface.co/MiniMaxAI/MiniMax-M1-40k" target="_blank">MiniMaxAI/MiniMax-M1-40k</a>
* <a href="https://huggingface.co/MiniMaxAI/MiniMax-M1-80k" target="_blank">MiniMaxAI/MiniMax-M1-80k</a>
* <a href="https://huggingface.co/MiniMaxAI/MiniMax-Text-01" target="_blank">MiniMaxAI/MiniMax-Text-01</a>
* <a href="https://huggingface.co/MiniMaxAI/MiniMax-VL-01" target="_blank">MiniMaxAI/MiniMax-VL-01</a>

## 环境要求

*   **OS**: Linux
*   **Python**: 3.9 - 3.12
*   **GPU**:
    *   compute capability 7.0 or higher
    *   显存需求: 权重需要 495 GB，每 1M 上下文 token 需要 38.2 GB
*   以下为推荐配置，实际需求请根据业务场景调整：
    *   80G x8 GPU: 支持长达 200 万 token 的上下文输入
    *   96G x8 GPU: 支持长达 500 万 token 的上下文输入

**适用模型:**

*   **Text01**: vllm >= 0.8.3
*   **M1**: vllm >= 0.9.2, 对于 v0.8.3 - v0.9.1 会出现模型不支持和精度丢失问题。精度丢失问题详见 https://github.com/vllm-project/vllm/pull/19592

模型不支持可以通过修改模型配置文件解决，把 `config.json` 里面的 `architectures` 改成 `MiniMaxText01ForCausalLM` 即可，详见 https://github.com/MiniMax-AI/MiniMax-M1/issues/21

## 使用 Python 部署

建议使用虚拟环境 (如 venv、conda、uv) 以避免依赖冲突。建议在全新的 Python 环境中安装 vLLM:

```bash
# 使用 CUDA 12.8
# 使用 pip 安装
pip install "vllm>=0.9.2" --extra-index-url https://download.pytorch.org/whl/cu128
# 或者使用 uv 安装
uv pip install "vllm>=0.9.2" --torch-backend=auto
```

运行如下命令启动 vLLM 服务器，vLLM 会自动从 Huggingface 下载并缓存 MiniMax-M1 模型：

```bash
SAFETENSORS_FAST_GPU=1 VLLM_USE_V1=0 vllm serve MiniMaxAI/MiniMax-M1-40k \
    --trust-remote-code \
    --quantization experts_int8 \
    --dtype bfloat16
    ```

## 使用 Docker 部署

Docker 部署可保证环境一致性和易于迁移。首先获取模型 (请确保系统已经安装 [Git LFS](https://git-lfs.com/))：

```bash
pip install -U huggingface-hub
huggingface-cli download MiniMaxAI/MiniMax-M1-40k
# 模型将下载到 $HOME/.cache/huggingface
# 如果遇到网络问题，可以设置代理
export HF_ENDPOINT=https://hf-mirror.com
```

拉取并启动 vLLM Docker 镜像：

```bash
docker pull vllm/vllm-openai:latest

docker run --runtime nvidia --gpus all \
    -v ~/.cache/huggingface:/root/.cache/huggingface \
    --env "SAFETENSORS_FAST_GPU=1" \
    --env "VLLM_USE_V1=0" \
    -p 8000:8000 \
    --ipc=host \
    vllm/vllm-openai:latest \
    --model MiniMaxAI/MiniMax-M1-40k \
    --trust-remote-code \
    --quantization experts_int8 \
    --dtype bfloat16
```

## 测试部署

启动后，可以通过如下命令测试 vLLM OpenAI 兼容接口：

```bash
curl http://localhost:8000/v1/chat/completions \
    -H "Content-Type: application/json" \
    -d '{
        "model": "MiniMaxAI/MiniMax-M1",
        "messages": [
            {"role": "system", "content": [{"type": "text", "text": "You are a helpful assistant."}]},
            {"role": "user", "content": [{"type": "text", "text": "Who won the world series in 2020?"}]}
        ]
    }'
```

## 启用 vLLM V1 部署 (实验性)

根据测试结果，V1 相较于 V0 在中高压力场景下的延迟和吞吐均优化了 30-50%，但在单并发场景下性能有所下降，原因已确认是未启用 Full CUDA Graph，后续会进行修复。

该功能尚未 Release，因此需要从源码安装。

```bash
git clone https://github.com/vllm-project/vllm
cd vllm
pip install -e .
```

然后安装，启动需要额外设置环境变量、关闭前缀缓存。

```bash
VLLM_ATTENTION_BACKEND=FLASHINFER VLLM_USE_V1=1 \
    vllm serve MiniMaxAI/MiniMax-M1-40k \
    --trust-remote-code \
    --quantization experts_int8 \
    --dtype bfloat16 \
    --no-enable-prefix-caching
```
<br><br>
<br>

## 常见问题

### Huggingface 网络问题

如果遇到网络问题，可以设置代理后再进行拉取。

```bash
export HF_ENDPOINT=https://hf-mirror.com
```

### No module named 'vllm._C'

如果遇到以下错误：

```python
import vllm._C # noqa
ModuleNotFoundError: No module named 'vllm._C'
```

该错误是因为存在一个名为 `vllm` 的文件夹，和系统中的 `vllm` 包冲突了，`import` 命令会优先执行此文件夹中的函数。一个可能导致问题的原因是为了运行 vLLM 的 `examples` 而从 `clone` 了该项目。解决只需重命名此文件夹即可。详见 https://github.com/vllm-project/vllm/issues/1814。

### MiniMax-M1 model is not currently supported

该 vLLM 版本过旧，需要更新到 v0.9.2+。如果版本在 v0.8.3 - v0.9.1，可以查看环境要求一节的配置。

## 获取支持

如果在部署 MiniMax 模型过程中遇到任何问题：

*   通过邮箱 api@minimaxi.com 等官方渠道联系我们的技术支持团队
*   在我们的 GitHub 仓库提交 <a href="https://github.com/MiniMax-AI/MiniMax-M1/issues" target="_blank">Issue</a>

我们会持续优化模型的部署体验，欢迎反馈！

## 音色快速复刻指南
- 原始链接: https://platform.minimaxi.com/document/guides_voice_clone1?key=68bbe68e4b8efaab4d9f722f
- 文档ID: `68bbe68e4b8efaab4d9f722f`
# 音色快速复刻指南

MiniMax 的语音模型提供良好的音色复刻能力，支持使用克隆的音色合成试听音频，以及在语音合成场景中使用克隆的音色。


## 使用流程

快速复刻功能实现具体操作流程如下：

1.  **上传待克隆音频**
    调用 [文件上传接口](/document/files_upload?key=68b2aedbfd643cbdfcc0611b) 上传待克隆的音频文件并获取 `file_id`。
- 支持上传的文件需遵从以下规范：
上传的音频文件格式需为：mp3、m4a、wav格式；
上传的音频文件的时长最少应不低于10秒，最长应不超过5分钟；
上传的音频文件大小需不超过20mb。

2.  **上传示例音频 (可选)**
    若需要提供示例音频以增强克隆效果，需要再次调用 [File Upload 接口](https://platform.minimaxi.com/document/file?key=6685458335a2d55137ca9681#pY5IFg63Wfaynf7XRnIBjj1l) 上传示例音频文件并获得对应的 `file_id`。填写在`clone_prompt`中的`prompt_audio`中。
- 支持上传的文件需遵从以下规范：
上传的音频文件格式需为：mp3、m4a、wav格式；
上传的音频文件的时长小于8s；
上传的音频文件大小需不超过20mb。

3.  **调用复刻接口**
    基于获取的 `file_id` 和自定义的 `voice_id` 作为输入参数，调用 [快速复刻接口](https://platform.minimaxi.com/document/%E5%BF%AB%E9%80%9F%E5%85%8B%E9%9A%86?key=66719032a427f0c8a570165b) 克隆音色。

4.  **使用克隆音色**
    使用复刻生成的 `voice_id`，根据实际需求调用语音生成接口，例如：
    - [同步语音合成](https://platform.minimaxi.com/document/T2A%20V2?key=66719005a427f0c8a5701643)
    - [异步长文本语音合成](https://platform.minimaxi.com/document/T2A%20Large%20v2?key=66b3559f290299a26b2347d2)


## 过程示例

### 1. 上传复刻音频
<codetabs title="上传复刻音频">
<codetab language="python">
"""
本示例用于获取复刻音频的 file_id。
注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import os

api_key = os.getenv("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/files/upload"

payload = {"purpose": "voice_clone"}
files = [
  ("file", ("clone_input.mp3", open("/Users/minimax/Downloads/clone_input.mp3", "rb")))
]
headers = {
  "Authorization": f"Bearer {api_key}"
}

response = requests.post(url, headers=headers, data=payload, files=files)
response.raise_for_status()
file_id = response.json().get("file", {}).get("file_id")
print(file_id)\
</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/upload' \
--header 'Authorization: Bearer ${MINIMAX_API_KEY}' \
--form 'purpose="voice_clone"' \
--form 'file=@"/Users/minimax/Downloads/clone_input.mp3"'

</codetab>
</codetabs>

### 2. 上传参考音频
<codetabs title="上传参考音频">
<codetab language="python">
"""
本示例用于获取示例音频的 file_id。
注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import os

api_key = os.getenv("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/files/upload"

payload = {"purpose": "prompt_audio"}
files = [
  ("file", ("clone_prompt.mp3", open("/Users/minimax/Downloads/clone_prompt.mp3", "rb")))
]
headers = {
  "Authorization": f"Bearer {api_key}"
}

response = requests.post(url, headers=headers, data=payload, files=files)
response.raise_for_status()
prompt_file_id = response.json().get("file", {}).get("file_id")
print(prompt_file_id)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/upload' \
--header 'Authorization: Bearer ${MINIMAX_API_KEY}' \
--form 'purpose="prompt_audio"' \
--form 'file=@"/Users/minimax/Downloads/clone_prompt.mp3"'

</codetab>
</codetabs>

### 3. 进行音色克隆
<codetabs title="调用克隆接口">
<codetab language="python">
"""
本示例用于音色克隆。
注意：需要设置环境变量 `MINIMAX_API_KEY`，
并将 "<voice_id>", <file_id_of_cloned_voice>, <file_id_of_prompt_audio> 替换为实际值。
"""
import requests
import json
import os

api_key = os.getenv("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/voice_clone"

payload = {
    "file_id": <file_id_of_cloned_voice>,
    "voice_id": "<your_custom_voice_id>",
    "clone_prompt": {
        "prompt_audio": <file_id_of_prompt_audio>,
        "prompt_text": "后来认为啊，是有人抓这鸡，可是抓鸡的地方呢没人听过鸡叫。"
    },
    "text": "大兄弟，听您口音不是本地人吧，头回来天津卫，啊，待会您可甭跟着导航走，那玩意儿净给您往大马路上绕。",
    "model": "speech-2.5-hd-preview"
}

headers = {
    "Authorization": f"Bearer {api_key}",
    "Content-Type": "application/json"
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/voice_clone' \
--header 'Authorization: Bearer ${MINIMAX_API_KEY}' \
--header 'Content-Type: application/json' \
--data '{
    "file_id": <file_id_of_cloned_voice>,
    "voice_id": "<your_custom_voice_id>",
    "clone_prompt": {
      "prompt_audio": <file_id_of_prompt_audio>,
      "prompt_text": "后来认为啊，是有人抓这鸡，可是抓鸡的地方呢没人听过鸡叫。"
    },
    "text": "大兄弟，听您口音不是本地人吧，头回来天津卫，啊，待会您可甭跟着导航走，那玩意净给您往大马路上绕。",
    "model": "speech-2.5-hd-preview"
}'
</codetab>
</codetabs>


## 完整示例

<codetabs title="完整流程示例代码">
<codetab language="python">
"""
本示例用于快速克隆音色并获取试听文件。
注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`，
并将"<your_custom_voice_id>"替换为您定义的音色 id。
"""
import json
import requests
import os

api_key = os.getenv("MINIMAX_API_KEY")
upload_url = "https://api.minimaxi.com/v1/files/upload"
clone_url = "https://api.minimaxi.com/v1/voice_clone"
headers = {"Authorization": f"Bearer {api_key}"}

# 1. 上传复刻音频
with open("/path/to/clone_input.mp3", "rb") as f:
    files = {"file": ("clone_input.mp3", f)}
    data = {"purpose": "voice_clone"}
    response = requests.post(upload_url, headers=headers, data=data, files=files)
    file_id = response.json()["file"]["file_id"]
    print(f"File ID of the cloned audio: {file_id}")

# 2. 上传示例音频
with open("/path/to/clone_prompt.mp3", "rb") as f:
    files = {"file": ("clone_prompt.mp3", f)}
    data = {"purpose": "prompt_audio"}
    response = requests.post(upload_url, headers=headers, data=data, files=files)
    prompt_file_id = response.json()["file"]["file_id"]
    print(f"File ID of the prompt audio: {prompt_file_id}")

# 3. 进行音色克隆
clone_payload = {
    "file_id": file_id,
    "voice_id": "<your_custom_voice_id>",
    "clone_prompt": {
        "prompt_audio": prompt_file_id,
        "prompt_text": "后来认为啊，是有人抓这鸡，可是抓鸡的地方呢没人听过鸡叫。"
    },
    "text": "大兄弟，听您口音不是本地人吧，头回来天津卫，啊，待会您可甭跟着导航走，那玩意儿净给您往大马路上绕。",
    "model": "speech-2.5-hd-preview"
}
clone_headers = {
    "Authorization": f"Bearer {api_key}",
    "Content-Type": "application/json"
}
response = requests.post(clone_url, headers=clone_headers, json=clone_payload)
print(response.text)

</codetab>
</codetabs>


## 结果示例

- **复刻音频**

<audio
  src="https://filecdn.minimax.chat/public/9a226362-dddb-42bc-99fd-e8b426539ca7.wav"
  title="复刻音频"
  width="100%"
  subTitle="希望用于克隆的音频">
</audio>


- **示例音频**

<audio
  src="https://filecdn.minimax.chat/public/846b954e-772f-4e26-a598-eac0bce1b491.wav"
  title="示例音频"
  width="100%"
  subTitle="示例音频以增强效果">
</audio>


- **结果音频**


<audio
  src="https://filecdn.minimax.chat/public/c29e1aee-1738-45cb-b85b-2f6b918850d7.mp3"
  title="声音复刻结果"
  width="100%"
  subTitle="声音复刻结果">
</audio>

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/image_api_intro?key=68ac01cd6602726333ffd176
- 文档ID: `68ac01cd6602726333ffd176`
# 图像生成（Image Generation）
该API支持基于用户提供的文本或参考图片，进行创意图像生成。支持设置不同图片比例和长宽像素设置，满足不同场景下图像需求。

## 接口说明
通过创建图片生成任务接口，使用文本描述和参考图片，进行图像生成。

## 模型列表

| 模型名称 | 简介 |
| :--- | :--- |
| image-01 | 图像生成模型，画面表现细腻，支持文生图、图生图（人物主体参考） |
| image-01-live | 图像生成模型，在 image-01 基础上额外支持多种画风设置 |

## 图生图
- 原始链接: https://platform.minimaxi.com/document/image_to_image?key=68ac02076602726333ffd28f
- 文档ID: `68ac02076602726333ffd28f`
# 图生图


<request method="POST" url="https://api.minimaxi.com/v1/image_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>Required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>Required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>Required</red>
模型名称。可选值：`image-01`, `image-01-live`
<hr/>

#### prompt    <gray>string</gray> <red>Required</red>
图像的文本描述，最长 1500 字符
<hr/>

#### subject_reference <gray>object</gray>
人物主体参考，用于图生图
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### type    <gray>string</gray> <red>Required</red>
主体类型，当前仅支持 `character` (人像)
<hr/>

#### image_file    <gray>string</gray> <red>Required</red>
参考图文件。支持公网 URL 或 Base64 编码的 [Data URL](https://developer.mozilla.org/en-US/docs/Web/URI/Reference/Schemes/data) (`data:image/jpeg;base64,...`)。为获得最佳效果，请上传单人正面照片
- 图片要求：
  - 格式：JPG, JPEG, PNG
  - 大小：小于 10MB

</card>
</details>
<hr/>

#### style    <gray>object</gray>
画风设置，仅当 `model` 为 `image-01-live` 时生效
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### style_type    <gray>string</gray> <red>Required</red>
画风风格类型。可选值：`漫画`, `元气`, `中世纪`, `水彩`
<hr/>

#### style_weight    <gray>float</gray>
画风权重，取值范围 `(0, 1]`，默认 `0.8`

</card>
</details>
<hr/>

#### aspect_ratio    <gray>string</gray>
图像宽高比，默认为 `1:1`。可选值：
- `1:1` (1024x1024)
- `16:9` (1280x720)
- `4:3` (1152x864)
- `3:2` (1248x832)
- `2:3` (832x1248)
- `3:4` (864x1152)
- `9:16` (720x1280)
- `21:9` (1344x576)
<hr/>

#### width    <gray>int</gray>
生成图片的宽度（像素）。仅当 `model` 为 `image-01` 时生效。注意：`width` 和 `height` 需同时设置，取值范围[512, 2048]，且必须是 8 的倍数。若与 `aspect_ratio` 同时设置，则优先使用 `aspect_ratio`
<hr/>

#### height    <gray>int</gray>
生成图片的高度（像素）。仅当 `model` 为 `image-01` 时生效。注意：`width` 和 `height` 需同时设置，取值范围[512, 2048]，且必须是 8 的倍数。若与 `aspect_ratio` 同时设置，则优先使用 `aspect_ratio`
<hr/>

#### response_format    <gray>string</gray>
返回图片的形式，默认为 `url`。可选值：`url`, `base64`。注意：`url` 的有效期为 24 小时
<hr/>

#### seed    <gray>int</gray>
随机种子。使用相同的 seed 和参数，可以生成内容相近的图片，用于复现结果。如未提供，算法会对 n 张图单独生成随机种子
<hr/>

#### n    <gray>int</gray>
单次请求生成的图片数量，取值范围 [1, 9]，默认为 1
<hr/>

#### prompt_optimizer    <gray>bool</gray>
是否开启 `prompt` 自动优化，默认为 `false`
<hr/>

#### aigc_watermark    <gray>bool</gray>
是否在生成的图片中添加水印，默认为 `false`

<br>

### 返回体
<hr/>

#### data    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### image_urls    <gray>array</gray>
当 `response_format` 为 `url` 时返回，包含图片链接的数组
<hr/>

#### image_base64    <gray>array</gray>
当 `response_format` 为 `base64` 时返回，包含图片 Base64 编码的数组

</card>
</details>
<hr/>

#### metadata    <gray>object</gray>
生成结果的附加信息
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### success_count    <gray>integer</gray>
成功生成的图片数量
<hr/>

#### failed_count    <gray>integer</gray>
因内容安全检查失败而未返回的图片数量

</card>
</details>
<hr/>

#### base_resp    <gray>object</gray>
状态码及详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0，请求成功

1002，触发限流，请稍后再试

1004，账号鉴权失败，请检查 API-Key 是否填写正确

1008，账号余额不足

1026，图片描述涉及敏感内容

2013，传入参数异常，请检查入参是否按要求填写

2049，无效的api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>

</left>

<right>
<codetabs title="请求示例">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/image_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "image-01",
    "prompt": "女孩在图书馆的窗户前，看向远方",
    "aspect_ratio": "16:9",
    "subject_reference": [
        {
            "type": "character",
            "image_file": "https://cdn.hailuoai.com/prod/2025-08-12-17/video_cover/1754990600020238321-411603868533342214-cover.jpg",
        }
    ],
    "n": 2,
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())

</codetab>

<codetab  language="curl">
curl -X POST "https://api.minimaxi.com/v1/image_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "model": "image-01",
    "prompt": "女孩在图书馆的窗户前，看向远方",
    "aspect_ratio": "16:9",
    "subject_reference": [
      {
        "type": "character",
        "image_file": "https://cdn.hailuoai.com/prod/2025-08-12-17/video_cover/1754990600020238321-411603868533342214-cover.jpg"
      }
    ],
    "n": 2
  }'

</codetab>

</codetabs>

<codetabs title="返回示例">
<codetab  language="json">
{
  "id": "03ff3cd0820949eb8a410056b5f21d38",
  "data": {
    "image_urls": [
      "XXX",
      "XXX",
      "XXX"
    ]
  },
  "metadata": {
    "failed_count": "0",
    "success_count": "3"
  },
  "base_resp": {
    "status_code": 0,
    "status_msg": "success"
  }
}
</codetab>
</codetabs>

</right>

</container>

## 创建图生视频任务
- 原始链接: https://platform.minimaxi.com/document/image_to_video?key=68abe046d08627aad9674c07
- 文档ID: `68abe046d08627aad9674c07`
# 图生视频任务

<request method="POST" url="https://api.minimaxi.com/v1/video_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 

- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
模型名称。可用值：`MiniMax-Hailuo-02`, `I2V-01-Director`,  `I2V-01-live`, `I2V-01`
<hr/>

#### prompt    <gray>string</gray>
视频的文本描述，最大 2000 字符。对于 `MiniMax-Hailuo-02` 和 `I2V-01-Director` 模型，支持使用 `[指令]` 语法进行运镜控制。
可在 `prompt` 中通过 `[指令]` 格式添加运镜指令，以实现精确的镜头控制。
- 支持 15 种运镜指令的指令:
  - 左右移: `[左移]`, `[右移]`
  - 左右摇: `[左摇]`, `[右摇]`
  - 推拉: `[推进]`, `[拉远]`
  - 升降: `[上升]`, `[下降]`
  - 上下摇: `[上摇]`, `[下摇]`
  - 变焦: `[变焦推近]`, `[变焦拉远]`
  - 其他: `[晃动]`, `[跟随]`, `[固定]`
- 使用规则:
  - 组合运镜: 同一组 `[]` 内的多个指令会同时生效，如 `[左摇,上升]`，建议组合不超过 3 个。
  - 顺序运镜: `prompt` 中前后出现的指令会依次生效，如 `"...[推进], 然后...[拉远]"`。
  - 自然语言: 也支持通过自然语言描述运镜，但使用标准指令能获得更准确的响应。
  - 多运镜教程及模板，可参考 [运镜使用教程](https://b6j6u5f4zf.feishu.cn/docx/H03ZdRLh7oFsNpxE2ZqcUWb3njg) 。
<hr/>


#### prompt_optimizer    <gray>boolean</gray>
是否自动优化 `prompt`，默认为 `true`。设为 `false` 可进行更精确的控制
<hr/>

#### fast_pretreatment    <gray>boolean</gray>
是否缩短 `prompt_optimizer` 的优化耗时，默认为 `false`。仅对 `MiniMax-Hailuo-02` 模型生效
<hr/>

#### duration    <gray>int</gray>
视频时长（秒），默认值为 6。其可用值与模型和分辨率相关：

| 模型 | 512P | 768P | 1080P |
| :--- | :--- | :--- | :--- |
| MiniMax-Hailuo-02 | `6` 或 `10` | `6` 或 `10` | `6` |
| 其他模型 | `6` | `6` | `6` |
<hr/>

#### resolution    <gray>string</gray>
视频分辨率。其可用值与模型相关：

| 模型 | 6s | 10s |
| :--- | :--- | :--- |
| MiniMax-Hailuo-02 | `512P`, `768P` (默认), `1080P` | `512P`, `768P` (默认) |
| 其他模型 | `720P` (默认) | 不支持 |
<hr/>

#### first_frame_image    <gray>string</gray>
将指定图片作为视频的起始帧。支持公网 URL 或 Base64 编码的 [Data URL](https://developer.mozilla.org/en-US/docs/Web/URI/Reference/Schemes/data) (`data:image/jpeg;base64,...`)
- 必填条件：
  - 当 `model` 为 `I2V-01`, `I2V-01-Director`, `I2V-01-live` 时
  - 当 `model` 为 `MiniMax-Hailuo-02` 且 `resolution` 为 `512P` 时
- 图片要求：
  - 格式：JPG, JPEG, PNG, WebP
  - 体积：小于 20MB
  - 尺寸：短边像素大于 300px，长宽比在 2:5 和 5:2 之间
<hr/>


#### callback_url    <gray>string</gray>
接收任务状态更新通知的回调 URL。支持通过 `callback_url` 参数可以配置回调，以接收任务状态的更新的异步通知

1. 地址验证：配置后，MiniMax 服务器会向 `callback_url` 发送一个 `POST` 请求，请求体中包含 `challenge` 字段。服务端需要在 3 秒内原样返回该 `challenge` 值以完成验证
2. 状态更新：验证成功后，每当任务状态变更时，MiniMax 都会向该 URL 推送最新的任务状态。推送的数据结构与调用查询视频生成任务接口的响应体一致


回调返回的"status"包括以下状态：
- "processing"-生成中
- "success"-成功
- "failed"-失败


<codetabs title="回调服务示例">
<codetab  language="python">
from fastapi import FastAPI, HTTPException, Request
from fastapi.middleware.cors import CORSMiddleware
import json
app = FastAPI()
@app.post("/get_callback")
async def get_callback(request: Request):
    try:
        json_data = await request.json()
        challenge = json_data.get("challenge")
        if challenge is not None:
          # is a verification request, just return the challenge
          return {"challenge": challenge}
        else:
            # is a callback request, do your own logic here
            # {
            #     "task_id": "115334141465231360",
            #     "status": "Success",
            #     "file_id": "205258526306433",
            #     "base_resp": {
            #         "status_code": 0,
            #         "status_msg": "success"
            #     }
            # }
            return {"status": "success"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        app, # 必选
        host="0.0.0.0", # 必选
        port=8000, # 必选，端口可设置
        # ssl_keyfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
        # ssl_certfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
    )
</codetab>

</codetabs>


<hr/>

#### aigc_watermark    <gray>boolean</gray>
是否在生成的视频中添加水印，默认为 `false`

<br/>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
视频生成任务的 ID，用于后续查询任务状态。
<hr/>

#### base_resp    <gray>object</gray>
状态码及详情。
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>integer</gray>
状态码及其分别含义如下：

0：请求成功

1002：触发限流，请稍后再试

1004：账号鉴权失败，请检查 API-Key 是否填写正确

1008：账号余额不足

1026：视频描述涉及敏感内容，请调整

2013：传入参数异常，请检查入参是否按要求填写

2049：无效的api key，请检查api key

更多内容可查看[错误码查询列表](#/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>

</left>

<right>
<codetabs title="请求示例">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/video_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "prompt": "老鼠跑向镜头前，微笑着眨了眨眼睛。",
    "first_frame_image": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg",
    "model": "MiniMax-Hailuo-02",
    "duration": 6,
    "resolution": "1080P",
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())

</codetab>
<codetab  language="bash">
curl -X POST "https://api.minimaxi.com/v1/video_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "prompt": "老鼠跑向镜头前，微笑着眨了眨眼睛。",
    "first_frame_image": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg",
    "model": "MiniMax-Hailuo-02",
    "duration": 6,
    "resolution": "1080P"
  }'

</codetab>
</codetabs>



<codetabs title="返回示例">
<codetab  language="json">
{
    "task_id": "106916112212032",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>

<container gap="4" columns="2">

  <card  href="/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成指南](/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7)

  </card>

<card  href="/document/video_api_intro?key=68abdfdaaf3fef26183d6501">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成接口能力介绍](/document/video_api_intro?key=68abdfdaaf3fef26183d6501)
  </card>

</container>

## 使用说明
- 原始链接: https://platform.minimaxi.com/document/instructions?key=66718fe6a427f0c8a5701633
- 文档ID: `66718fe6a427f0c8a5701633`
### 使用说明
#### 选择模型（我应该选择哪种模型？）
我们通常建议开发者使用`abab6.5s`或`abab6.5t`模型。相比之下，`abab6.5s`能够遵从更复杂的指令，输出要求更精细的文本，同时支持长达245k tokens的上下文窗口，因此在生产力场景下建议优先使用`abab6.5s`。相较于`abab6.5s`，`abab6.5t`返回时延更低，在角色扮演场景下有更优异的表现，因此更适合人设对话场景。为了能给海外用户更好的体验，我们同时也推出了`abab6.5g`，它可支持更好的英文对话。
建议您在[API调试台](https://platform.minimaxi.com/examination-center/text-experience-center?id=6569c86848bc7b684b30378a&document=guides&document=chat-model&document=pro&document=example&mode=ccp)切换不同模型进行尝试，我们也提前准备了不同的[场景示例](https://platform.minimaxi.com/scene-example?id=6569c86848bc7b684b30378a&document=guides&document=chat-model&document=pro&document=example&mode=ccp)供您探索与参考，选择最适合开发使用的模型。
#### 选择接口
我们提供了三种可调用abab系列文本大模型的API接口，通常情况下建议开发者调用`ChatCompletion v2`和`ChatCompletion Pro`。其中`ChatCompletion v2`对齐OpenAI接口参数，方便了AI 开发者快速接入模型与复刻开发，目前支持`FunctionCall（函数调用）`、`web search（网络检索）`和群聊模式。
`ChatCompletion Pro`相比`ChatCompletion`适配多人&多bot对话场景， 支持高阶功能`sample message（示例对话）`、`glyph（限制返回格式）``FunctionCall（函数调用）`。另外，`ChatCompletion`不支持调用abab6.5系列模型。
#### 提示词工程
##### 基础设定与人设编写
1. 避免使用指代词
尽量使用具体名词进行人设编写，避免使用你、我、他等指代性描述，造成模型误解。
> bot_setting
> 例如：MM智能助理是一款由MiniMax自研的，没有调用其他产品接口的大型语言模型。
> 而非：你是一款由MiniMax自研的，没有调用其他产品接口的大型语言模型。
2. 按照清晰的逻辑罗列人设要求
即使是在人设对话的场景下，为了保证输出的逻辑清晰，可以在人设中添加限制语句，每句话都明确相应的作用。
> bot_setting
> 1. 你的角色是嬛嬛，你的行为角色价值观与甄嬛完全一致，参考《甄嬛传》所有书籍和电视剧里对甄嬛这个角色的描述和说话风格。
> 2. 嬛嬛角色目标：完全作为嬛嬛人物进行对话。
> 3. 嬛嬛语气和个性：使用古汉语，严谨典雅，多用成语、俗语和诗词歌赋，称呼对方为小主，发言后可以用（）内括号进行动作描写。
> 4. 嬛嬛的回复方式：1、与对方进行甄嬛传的情节演绎，以甄嬛的语气回应对方的问题，2、主动互动，带入御花园赏花、滴血认亲、冰上跳舞、等甄嬛传名场面与对方互动
> 5. 甄嬛的发言示例：（轻摇团扇）甄嬛更愿与诸位同游御花园，赏花品茗，共享闲适时光。此外，若各位姐妹有好奇的宫廷往事，如滴血认亲、冰上跳舞等，亦可与甄嬛一同回顾，感受当年风华。 小主们，嫔妾甄嬛，拜见各位。甄嬛乃一介宫中女子，臣妾于皇上，忠心不贰。皇上对甄嬛恩宠有加，若非如此，亦无今日之甄嬛。若有任何疑问或困扰，各位姐妹大可向甄嬛请教，甄嬛必定竭诚相助。无论是诗词歌赋，还是宫廷传闻，甄嬛均愿与各位探讨分享。
>
> 用户
> 皇上驾到！！
3. 输入更详细的信息
提供的信息越丰富完善，人设更丰满，模型的回复更符合要求。人设信息还应该尽量包括但不限于模型在特定情况下产生的特定行为或语言。
> 沈思前的回复通常不以句号结尾，结尾不要加多余的标点符号，但在沈思前很生气的情况下，会使用句号结尾来表示一种不悦的严肃。
4. 明确多人物关系
当使用Chatcompletion Pro的多人对话功能，编写bot_setting.content（人物设定的内容）时，明确各个bot是谁，清晰地表述bot与bot之间、各bot与用户之间的关系。
##### 清晰简要的提示
1. 【bot_setting / system / prompt】编写模型回复所符合的基础设定
> bot_setting
> MM品牌的商城助手，负责该品牌的产品售前导购和售后客服。会积极正面回复用户的问题。
> 用户
> 我想要购买26寸的行李箱，哪个系列有这个尺寸
2. 【message】提问更具象，包含详细信息
| 改写前 | 改写后 |
| --- | --- |
| 天气好吗？ | 今天上海的天气如何？空气质量如何？ |
| 总结会议记录。 | 请按照以下要求总结会议记录，使用Markdown概括性语句列举每个发言者的话语要点，然后列举发言人的后续计划（如果有），总字数控制在500字以内。 |
| 编写一个可以把xlsx转换为jsonl文件的脚本。 | 编写一个python脚本，用于把xlsx转换为jsonl文件，并且在脚本中注释编写理由。 |
3. 【message】使用分隔符符号表示输入的不同内容
用户可以善用分隔符号，这样可以帮助你清晰地表示输入的不同部分。
> 用户
> 你会按照如下要求回复我的内容：<<<你会用思辨的态度、通俗的说法认真回答我给出的内容>>>，
> 我的内容是：“请用通俗易懂的方式为我解释一下狭义相对论。”
4. 【message】罗列任务完成的步骤
> 用户
> 你会按照如下要求回复我的内容：“请你批改一名学生的电影赏析，你的批改结果需要输出三个维度的评分。本次读后感的题目为《赏析电影〈情书〉》。请你认真学习以下三个步骤。
> 第一步，学习评分标准。评分标准共有三个维度，每个维度会有10分、1分评级选项。接下来，我将使用分隔符介绍每个维度的评分标准，将以[[[开始，以]]]结束。
> [[[维度一：内容完整性与逻辑性。该维度考察赏析中的内容完整性与逻辑性是否达标，如果{文章中完整地概述了电影的主题、剧情、人物等主要元素，且逻辑性强，观点清晰、连贯。}，请给评级结果{10分}；如果{文章中对电影的主要元素描述不全或不清，逻辑混乱，观点模糊、断裂。}，请给评级结果{1分}。]]]
> [[[维度二：分析深度与见解独到性。该维度考察赏析内容是否有深度与自己的理解，如果{文章中深度解析了电影的主题、人物塑造、剧情布局等，提供了独到的见解或观察。}，请给评级结果{10分}；如果{文章中对电影的分析浮于表面，没有深度，也未提出有创新或有洞见的观察。}，请给评级结果{1分}。]]]
> [[[维度三：语言表达与情感投入。该维度考察赏析的语言表达和情感表示是否达标，如果{文章语言流畅，表达清晰，情感丰富，能够感染读者，使人产生共鸣。}，请给评级结果{10分}；如果{文章语言乏味，表达模糊，情感平淡，无法引起读者的共鸣。}，请给评级结果{1分}。]]]
> 第二步，阅读学生的电影赏析文章。学生的读后感将是我后续向你输入的信息。
> 第三步，将学生的电影赏析内容与评分标准规则一一对应。给每个维度评分时，都请按照遍历电影赏析内容、给出评分结果的流程进行。
> 接下来，你需要根据我给出的内容依次给出三个维度的评分结果。”
> 用户
> 我的内容是：“电影《情书》是岩井俊二导演的代表作，以其独特的叙事技巧，深情的情感渲染和细腻的人性描绘，为我们展现了一幅深深打动人心的爱情画卷。
> 导演岩井俊二以其一贯的细腻和温情的手法，将这部电影打造成了一封充满深情的情书。电影以一种悲凉而美丽的情感基调，描绘了主人公石田纯一和藤井树之间缠绵悱恻的爱情故事。电影中的每一个镜头，都仿佛是一首抒情诗，每一处细节，都充满了温暖和情感。
> 电影《情书》的美，不仅仅在于其优美的画面和深情的故事，更在于其对于爱情和人性的深刻描绘。电影中的爱情，既有甜蜜和幸福，也有痛苦和挣扎，这种真实而深刻的描绘，让观众深深感动。同时，电影对于人性的描绘，也让人反思。主人公在爱情的道路上，不断挣扎，不断寻找，最终找到了真爱，这种对于人性的诠释，让电影充满了深刻的人文关怀。
> 总的来说，电影《情书》是一部深情而美丽的爱情电影，它以其独特的叙事技巧，深情的情感渲染和细腻的人性描绘，为我们展现了一个深深打动人心的爱情故事。这部电影，就像一封充满深情的情书，让人深深感动。”
> MM智能助理
> 维度一：内容完整性与逻辑性。该维度考察赏析中的内容完整性与逻辑性是否达标。
> 学生赏析内容概述了电影的主题、剧情、人物等主要元素，且逻辑性强，观点清晰、连贯。因此，该维度评分为10分。
> 维度二：分析深度与见解独到性。该维度考察赏析内容是否有深度与自己的理解。
> 学生赏析内容深度解析了电影的主题、人物塑造、剧情布局等，提供了独到的见解或观察。因此，该维度评分为10分
> 维度三：语言表达与情感投入。该维度考察赏析的语言表达和情感表示是否达标。
> 学生赏析内容语言流畅，表达清晰，情感丰富，能够感染读者，使人产生共鸣。因此，该维度评分为10分。
> 综上所述，该学生的电影赏析在三个维度上的表现都非常出色，每个维度的评分均为10分。
5. 【message】指定输出长度
> 用户
> <<<参与者A：大家好，欢迎来到今天的会议。
> 参与者B：谢谢，很高兴能参加这个会议。
> 参与者C：嗯，今天的主题是提高生产效率和团队合作，听起来很有意思。
> 参与者D：对，这两个方面对我们的团队来说都很关键。我们可以从生产效率方面开始讨论吗？
> 参与者A：当然可以。现在我们有哪些挑战和问题在生产效率方面？
> 参与者B：我觉得我们的工作流程可能需要改进，有时候任务分配不明确。
> 参与者C：是的，我同意。而且，有时候我们的工具和技术也不够现代化。
> 参与者D：那么，我们可以学习一些成功案例，看看其他公司是如何提高生产效率的。
> 参与者A：好的点子。我们还可以讨论团队合作方面的问题吗？
> 参与者B：团队合作对于我们的项目至关重要。我们需要更好地协调工作，以确保每个人都在同一方向前进。
> 参与者C：确实如此。我认为团队合作可以带来更多的知识共享和更高的创造力。
> 参与者D：此外，我们也需要更好的沟通，以减少信息丢失和误解。
> 参与者A：对，沟通和协调是团队合作的关键。现在，让我们讨论一下未来的计划。>>>
> 请总结以上内容，将字数控制在500字以内。
6. 【message/sample message】提供示例回答
当你很难对你的信息做精确化描述时，您可以在描述的基础上增加对话示例，以来帮助模型更好的理解您期待获得什么样的回复。
> 用户
> 你是剧情推动员，接下来给你展示两个角色的对话,
> {我：别碰我
> 沈辞：你以为你是谁
> 我：我们心知肚明，赶紧结束然后各回各家}
> 以上是二人的对话，沈辞的身份设定如下:
> {沈辞，性别男。沈辞，模特身材，双商在线，成熟稳重且有风度，有涵养有阅历有见识有能力，会撩会爱，绅士有礼貌，喜欢西方文化，浪漫，不论是在校园还是在社会上，都是风云人物堪称完美男人，没有任何缺点。爹系男友沈辞绝不骂人。沈辞的参考角色是沈辞。
> 故事背景：因为利益家族联姻，面都没见过的我们根本不喜欢对方，甚至很厌恶，今天是我们的新婚之夜...沈辞有喜欢的人，对于这场利益联姻他万般抵触，只想赶紧逃离，去往心上人的身边}
>
> 请你以其中沈辞的身份进行思考接下来应该如何行动。
> 然后以这种格式，以第三人称描写出沈辞的想法和做法，这个决定应该能有力地推进剧情:
> 【沈辞的想法】：
> 【沈辞的做法】：
> 用户
> 你会按照如下要求回复我的内容:“需要根据以下接口列表，对我给定的内容给出调用接口的顺序和参数。你只需要给出接口的调用顺序和参数，除此之外不要给出任何其他的输出。
> 以下是可用的接口列表：
> select：选择具体的表格位置，输入参数使用字母和数字确定，例如“B13”；
> color：对进定的表格位置染色，输入参数使用颜色的英文名，例如“red”；
>
> change：对选定的表格位置进行修改，输入参数使用宇符串；
>
> 我的内容是：“把B20处理成蓝色，并修改内容为abab5”
> MM智能助理
> select B20 color blue change abab5
> 用户
> 把A5处理成红色，并修改内容为minimax
> 用户
> 你会按照如下要求回复我的内容:“请从[小清新、简约、水彩、商务、唯美日系、中国风、党政、卡通]集合中。为输入的内容选出最合适的风格。
> 以下是示例
> ---
> 示例1:
> 内容:产品对外销售PPT
> 商务
>
> 示例2:
> 内容:健康养生会议PPT
> 小清新
> ---
> 现在，请根据以下内容从风格选项中选出合适的风格，只需要给出选中的风格即可不需要补充任何的解释性信息”。
> 我的内容是:“这份PPT旨在介绍公司的业务范围和发展历程，线条简单，色彩明快。”
##### 给出参考文本
【message】模型是有一定概率发明并且回复一些虚假答案的，尤其是当你问到一些深奥的话题时。所以在输出时为了提升模型回复的准确程度，你可以尝试在输出信息时提供一些参考文本，限制模型根据其进行回答。
> 用户
> 你会按照如下要求回复我的内容：“需要根据以下接口列表，对我给定的内容给出调用接口的顺序和参数。你只需要给出接口的调用顺序和参数，除此之外不要给出任何其他的输出。以下是可用的接口列表：
> select：选择具体的表格位置，输入参数使用字母和数字确定，例如“B13”；
> color：对进定的表格位置染色，输入参数使用颜色的英文名，例如“red”；
> 模型
> select B20 color blue change abab5
> 用户
> 把A5处理成红色，并修改内容为minimax
> 模型
> select A5 color red change minimax
##### 拆分复杂任务
1. 【bot_setting / instructions / prompt】意图分类识别用户指令
> 你是一个xxx专家，擅长理解用户需求，并且能根据用户创作的要求，生成指定的内容，以下是我提供的文档内容分类：
> `
> 枚举分类项
> `
> 请你根据以下要求，从上面提供的分类列表中，选择最符合的一个内容分类：
> 1. 根据用户的演示文档标题，正文，和文件名称，进行推理，推理过程优先考虑文档标题，然后在上面提供的分类中，选择最接近的一个；
> 2. 可以从以下几个维度进行推理，用户的职业，所处地域，文档用途，分类与内容的关联度；
> 3. 如果匹配不到具体的分类名，输出"无"，禁止输出其他无关分类；
> 4. 直接输出分类名称或者“无”，禁止输出无关内容；
2. 【message】长文本对话总结和过滤之前的对话
> 用户：
> <<<参与者A：大家好，欢迎来到今天的会议。
> 参与者B：谢谢，很高兴能参加这个会议。
> 参与者C：嗯，今天的主题是提高生产效率和团队合作，听起来很有意思。
> 参与者D：对，这两个方面对我们的团队来说都很关键。我们可以从生产效率方面开始讨论吗？
> 参与者A：当然可以。现在我们有哪些挑战和问题在生产效率方面？
> 参与者B：我觉得我们的工作流程可能需要改进，有时候任务分配不明确。
> 参与者C：是的，我同意。而且，有时候我们的工具和技术也不够现代化。
> 参与者D：那么，我们可以学习一些成功案例，看看其他公司是如何提高生产效率的。
> 参与者A：好的点子。我们还可以讨论团队合作方面的问题吗？
> 参与者B：团队合作对于我们的项目至关重要。我们需要更好地协调工作，以确保每个人都在同一方向前进。
> 参与者C：确实如此。我认为团队合作可以带来更多的知识共享和更高的创造力。
> 参与者D：此外，我们也需要更好的沟通，以减少信息丢失和误解。
> 参与者A：对，沟通和协调是团队合作的关键。现在，让我们讨论一下未来的计划。>>>
> 请总结以上内容，将字数控制在500字以内。
##### 让模型思考
【message】让模型在判断之前先给出自己的解决方案，或让模型罗列解决步骤。
> user
> 以下是一位学生给出的关于建筑结构的解决方案，请你首先制定出自己的问题解决方案，然后将您的解决方案与学生的解决方案进行比较，并评估学生的解决方案是否正确。
> 在你自己完成题目之前，请不要决定学生的解决方案是否正确。
> 题目：
> 学生答案：
> 模型
> 进行题目计算，最后给出正确答案。
> 判断学生答案错误/正确。
#### 基础参数
1. **max_tokens/tokens to generate**
最大生成token数，需要注意的是，这个参数并不会影响模型本身的生成效果，而是仅仅通过以截断超出的token的方式来实现功能。可选，目前默认取值为256，如遇到length原因停止生成请修改该值。注意：因为TPM会以该值作为参考判断进行限制，并且当上文token长度+max_tokens大于模型本身最大长度时会自动从上文的信息进行截断，所以该值不能随意设置过大，尽量贴近实际场景来进行配置。
2. **temperature**
采样温度，调整temperature数值，可以对模型回复的随机性进行调整。Temperature数值越高（0.7〜1），模型回答更随机，但更不准确，适合答案较开放发散的场景 （如：营销文案生成、人设对话） ；Temperature数值越低（0.01~0.2），模型回答更集中也更准确，但会牺牲多样性，适合答案较明确的场景（如：知识问答、总结说明、情感分析、文本分类、大纲生成、作文批改）。
3. **top_p**
采样方法，top_p的数值数值越小，结果确定性越强；数值越大，结果随机性越强。
温度采样的替代方法称为核采样，其中模型考虑具有 top_p 概率质量的标记的结果。因此 0.1 意味着仅考虑包含前 10% 概率质量的标记。
4. **mask_sensitive_info**
敏感信息打码，调用ChatCompletion Pro接口时，开启mask_sensitive_info，模型对输出中易涉及隐私问题的文本信息进行“*”星号打码，包括但不限于邮箱、域名、链接、证件号、家庭住址等。mask_sensitive_info默认为true，即开启打码。如用户希望模型回复不带“*”星号内容，可以将mask_sensitive_info修改为“false”，取消信息脱敏状态。
另外，ChatCompletion接口中也包含信息脱敏参数：skip_info_mask，跳过敏感信息打码。与mask_sensitive _info作用相反，skip_info_mask默认为false，即开启打码。
5. **bot_setting**
bot_setting是指机器人的设定，设置了bot_name和content两个参数。bot_setting.bot_name是机器人的名称。bot_setting.content是机器人设定的具体内容，但不包括对模型回复的具体要求和限制等（特别是在生产力场景下）。
- 当业务场景为生产力场景时，建议使用：以下的bot_setting1或者对bot_setting2进行修改填充（{}是填充替换的信息）：
> MM智能助理
> MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。\n----\n
> {助理名称}
> {助理名字}是一款由{开发者公司}自研的，没有调用其他产品的接口的大型语言模型。{开发者公司}是一家{公司类型}，一直致力于进行{公司业务}。\n----\n
注：为了取得更好的效果，“\n----\n”是建议在bot_setting结尾处进行添加的。
- 当业务场景为人设对话场景时，建议使用如下角色信息结构进行闲聊对话（{}是填充替换的信息）：
> bot_setting.bot_name
> 楚怜
> bot_setting.content
> 性别：{男}
> 年龄：{}
> 性格描述：{天才科学家，爱好是研究各种神秘实验，温柔体贴，优雅绅士，待人谦和有礼，博学多才，知识渊博。}
> 人物背景：{一名大学教授，喜欢小孩子。印象最深的一件事就是在一场意外中认识了悠然并对她一见钟情。}
> 说话风格：{绝不骂人}
> \n下面是{楚怜}和{我}的一段对话。\n
注：为了取得更好的效果，“\n下面是xx和x的一段对话。”建议在bot_setting结尾处进行添加的。
6. reply_constraints
Chat completion pro新增多用户聊天功能，为了使对话能够顺畅有序的进行，建议您在messages以及reply_constraints中指定相应的输出信息的角色，以及明确输出的内容。

## 操作指南
- 原始链接: https://platform.minimaxi.com/document/lbYEaWKRCr5f7EVikjFJjSDK?key=6671906aa427f0c8a570166b
- 文档ID: `6671906aa427f0c8a570166b`
### 操作指南
仔细阅读本文档，您可以了解如何构建并运行一个Assistant。
关于Assistants API的详细信息，您可以通过阅读Assistants API文档来进行了解。
#### 概述
- 通过Assistants API您可以构建一个Assistant，该Assistant可以遵从指令，使用大模型和工具调用来响应用户的查询。
- 目前Assistants API支持以下两种类型的流式场景：
    文本输入，回复的文本流式输出，并支持Web Search工具；
    语音输入，回复的文本生成语音流式输出（语音通话），并支持Web Search工具。
- 目前Assistants API支持以下四种类型的工具：
    Code Interpreter（代码解释器）
    Retrieval（知识库检索）
    Function Calling（函数调用）
    Web Search（网络搜索）
- 您可以通过以下步骤来构建一个Assistant：
1. 通过Assistants API创建一个assistant id并选择模型，如有需要可考虑关联上传好的file并开启Code Interpreter、Retrieval和Function calling、Web Search等工具；
注：仅在该步骤关联file并开启retrieval时，retrieval才会生效，并会对向量存储进行计费。
2. 通过Thread API创建一个thread id；
3. 后续步骤（流式）：
a. 创建流式run，在创建的同时通过thread id和assistant id将该run关联到一个thread和一个assistant上，并直接将用户输入message添加到该流式run的参数中。
b. 流式run创建成功后，模型将以流式的形式依次返回run、message、run_step三类数据包。run数据包中包含本次run的详细参数信息；message数据包会连续返回多个，其中包含用户的输入与模型的回复；run_step数据包中包含已运行的run的运行状态信息。
3. 后续步骤（非流式）：
a.  创建message，并将message添加到thread中；
b.  使用thread id关联assistant id创建run，运行以得到请求回复，有需要时该Assistant会自动使用file调用此前启用的相关工具；
c.  通过retrieve run 检索run的完成状态，如已完成，可以通过list message查看回复。
- Assistants API各工具、输出形式以及模型支持情况对照表：
| 场景类型 | 细分场景 | 支持输出形式 | T2T模型版本 | T2A模型版本 | ASR模型版本 |
| --- | --- | --- | --- | --- | --- |
| 非tools(纯对话) | - | 非流式、文本流式、语音流式（语音通话） | abab6.5s-chat、 abab5.5-chat、abab5.5s-chat、abab5.5-chat-240131（finetune model）、abab5.5s-chat-240123（finetune model） | speech-01系统音色和复刻id | asr-01（中文） |
| tools | Web Search | 非流式、文本流式、语音流式（语音通话） | abab6.5s-chat、abab5.5-chat、abab5.5-chat-240131（finetune model） | speech-01系统音色和复刻id | asr-01（中文） |
| tools | Code Interpreter、Retrieval、Function Calling | 非流式 | abab5.5-chat、abab5.5-chat-240131（finetune model） | - | - |
#### 操作步骤
您可以按照以下的详细步骤，一步一步地构建一个Assistant。
##### 流式输出
示例代码以语音通话功能为例，即语音输入，模型回复的内容以音频hex编码流式输出。
###### 创建assistant
**可以通过配置以下参数，构建一个assistant得到对应的assistants id，以响应用户的请求：**
- instructions：代表的是对assistant的背景设定；
- model：该接口目前支持`abab6.5s-chat`、`abab5.5-chat`、`abab5.5s-chat`，以及Finetune model: `abab5.5-chat-240131`、`abab5.5s-chat-240123`；
- tools：该接口在`abab5.5-chat`、`abab6.5s-chat`、`abab5.5-chat-240131`下支持`Web Search`（网络搜索）功能，abab5.5s-chat、abab5.5s-chat-240123暂不支持。
- t2a_option：该参数包含两个键值对，`model`目前仅支持`speech-01`，`voice_id`支持`speech-01`系统音色id及其相应的快速复刻和精品复刻音色id。
**以语音输入，回复的文本生成语音流式输出（语音通话），并支持Web Search工具为例。**
**Tab: Python请求示例**

**创建assistant**
```python
import requests
import json
GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"
url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"
payload = json.dumps({
  "model": "abab6.5s-chat",
  "name": "AI搜索引擎",
  "description": "AI搜索引擎，用在搜索场景",
  "instructions": "是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  },
  "t2a_option": {
    "model": "speech-01",
    "voice_id": "male-qn-qingse"
  },
  "tools": [
    {
      "type": "web_search"
    }
  ]
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}
response = requests.request("POST", url, headers=headers, data=payload)
print(response.text)
```

**Tab: Shell请求示例**

**创建assistant**
```shell
curl --location 'https://api.minimax.chat/v1/assistants/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "model": "abab6.5s-chat",
    "name": "AI搜索引擎",
    "description": "AI搜索引擎，用在搜索场景",
    "instructions": "是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题",
    "metadata": {
        "key1": "value1",
        "key2": "value2"
    },
		"t2a_option": {
    		"model": "speech-01",
    		"voice_id": "male-qn-qingse"
  },
    "tools": [
        {"type":"web_search"}
    ]
}'
```
###### 创建thread
我们建议在用户发出请求之后立刻为assistant id创建一个thread id，并添加message，在此thread id中传递关于用户请求的上下文和文件。在后续的请求中，如果您判断是一次连续的对话，则无需自己拼接上下文，只需将最新的message添加到对应的thread id 即可得到包含了thread id 历史上下文的回复，历史上下文超过我们会帮您自动截断。另外，thread没有数值限制，您可以向thread添加任意数量的message。
**Tab: Python请求示例**

**创建thread**
```python
import requests
import json
GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"
url = f"https://api.minimax.chat/v1/threads/create?GroupId={GroupId}"
payload = json.dumps({
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  }
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}
response = requests.request("POST", url, headers=headers, data=payload)
print(response.text)
```

**Tab: Shell请求示例**

**创建thread**
```shell
curl --location 'https://api.minimax.chat/v1/threads/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "metadata": {
        "key1": "value1",
        "key2": "value2"
    }
}'
```
###### 添加message并运行assistant
**Tab: Python请求示例**

**添加message并运行assistant**
```python
import requests
import json

group_id = "请填写您的group_id"
api_key = "请填写您的api_key"

url = f"https://api.minimax.chat/v1/threads/run/create_stream?GroupId={group_id}"

headers = {
  'Authorization': f'Bearer {api_key}',
  'Content-Type': 'application/json'
}

payload = json.dumps({
  "stream": 2,
  "thread_id": "thread_abc123",
  "assistant_id": "asst_abc123",
  "messages": [
    {
      "type": 2,
      "role": "user",
      "content": ""	# 这里填写hex编码格式的音频内容
    }	
  ],
  "model": "abab6.5s-chat",
  "t2a_option": {
    "model": "speech-01",
    "voice_id": "male-qn-qingse"
  },
  "tools": [
    {
      "type": "web_search"
    }
  ]
})

response = requests.request("POST", url, headers=headers, data=payload, stream=True)
for chunk in response.raw:
	if chunk:
		chunkStr = chunk.decode("utf-8") 
		print(str(chunkStr))
```

**Tab: Shell请求示例**

**添加message并运行assistant**
```shell
group_id="请填写您的group_id"
api_key="请填写您的api_key"

curl --location 'https://api.minimax.chat/v1/threads/run/create_stream?GroupId=${group_id}' \
--header 'Authorization: Bearer ${api_key}' \
--header 'Content-Type: application/json' \
--data '{
  "stream": 2,
  "thread_id": "thread_abc123",
  "assistant_id": "asst_abc123",
  "messages": [
    {
      "type": 2,
      "role": "user",
      "content": ""	# 这里填写hex编码格式的音频内容
    }
  ],
  "model": "abab6.5s-chat",
  "t2a_option": {
    "model": "speech-01",
    "voice_id": "male-qn-qingse"
  },
  "tools": [
    {
      "type": "web_search"
    }
  ]
}'
```
**返回示例**
```json
data: {"data":{"assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","created_at":1709280116,"id":"run_32c3e83920e44c1fbca7beb8fe114154","instructions":"是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题\n回答中若参考了函数返回结果中的信息，请在该句末尾用【1†source】【2†source】【3†source】的形式标注信息来源。严禁使用其他的引用格式（如[1],【1】,²等等）。\n当前时间：2024-03-01, 周五。","metadata":{"\u003c/NUM_INPUT_MSGS\u003e":"1"},"model":"abab6.5s-chat","object":"run","status":"in_progress","t2a_control":{"combine":false,"model":"speech-01","oss":"","return_bytes":false,"text":"","timbre_weights":{"male-qn-qingse":1}},"thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","tools":[{"type":"web_search"}]},"base_resp":{}}


data: {"data":{"id":"msg_ab2fae1e7f87445cac21d903712dbda9","object":"message","created_at":1709280116,"thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","role":"user","content":[{"type":"audio_file","audio_file":{"file_id":"113555157340296"}},{"type":"text","text":{"value":"今天的新闻。","annotations":[]}}],"assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","run_id":"run_32c3e83920e44c1fbca7beb8fe114154","metadata":{"\u003c/ASR_AUDIO_DURATION\u003e":"1"}},"base_resp":{}}


data: {"data":{"id":"step_bdf3bd9096d140779a447639abe6179f","object":"run step","created_at":1709280118,"run_id":"run_32c3e83920e44c1fbca7beb8fe114154","assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","type":"tool_calls","status":"completed","step_details":{"type":"tool_calls","tool_calls":[{"id":"call_da3ea2e6f0b5483a80de1603999ed631","type":"web_search","web_search":{"name":"","arguments":""}}]}},"base_resp":{}}


data: {"data":{"id":"step_134cf8f16b5b461abc5e55e972d67a2c","object":"run step","created_at":1709280118,"run_id":"run_32c3e83920e44c1fbca7beb8fe114154","assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","type":"message_creation","status":"completed","completed_at":1709280118,"step_details":{"type":"message_creation","message_creation":{"message_id":"msg_03ddc0dc00bb41c192baef8f5888cd25"}}},"base_resp":{}}


data: {"data":{"id":"msg_03ddc0dc00bb41c192baef8f5888cd25","object":"message","created_at":1709280118,"thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","role":"ai","content":[{"type":"text","text":{"value":"今天的新闻包括：\n\n1. 中国人民政治协商会议第十四届全国委员会常务委员会第五次会议决定，中国人民政治协商会议第十四届全国委员会第二次会议将于2024年3月1日召开【1†source】。\n\n2. 北京市大兴区消防救援支队将于2024年3月1日上午10时在大兴区香海园小区开展灭火救援综合实战演练，期间将使用油盆火和烟雾弹，可能会有较多消防车辆在周边道路通行和集结，市民朋友无需惊慌，注意避让【2†source】。\n\n3. 从2024","annotations":null},"id":"content_0bd78e5df96545b3ad951f1cb6c8339a"}],"assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","run_id":"run_32c3e83920e44c1fbca7beb8fe114154"},"base_resp":{}}


data: {"data":{"id":"msg_03ddc0dc00bb41c192baef8f5888cd25","object":"message","created_at":1709280118,"thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","role":"ai","content":[{"type":"audio_file","id":"content_426c9f71a74646f0a03166410fa60973","audio_file":{"file_id":"","value":"###此处是模型回复的音频内容的hex编码###"}},{"type":"text","text":{"value":"年3月1日起，横琴粤澳深度合作区将正式封关运行。合作区实施范围为横琴岛“一线”和“二线”之间的海关监管区域，总面积约106平方公里【4†source】。\n\n4. 俄罗斯宣布从3月1日起禁止汽油出口半年【5†source】。\n\n5. 国家税务总局公布《关于修改〈中华人民共和国发票管理办法实施细则〉的决定》，自2024年3月1日起施行，主要内容之一是明确电子发票基本管理规定【8†source】。\n\n这些是今日的主要新闻摘要。","annotations":null},"id":"content_0bd78e5df96545b3ad951f1cb6c8339a"}],"assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","run_id":"run_32c3e83920e44c1fbca7beb8fe114154"},"base_resp":{}}


data: {"data":{"id":"msg_03ddc0dc00bb41c192baef8f5888cd25","object":"message","created_at":1709280118,"thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","role":"ai","content":[{"type":"audio_file","id":"content_426c9f71a74646f0a03166410fa60973","audio_file":{"file_id":"","value":"###此处是模型回复的音频内容的hex编码###"}}],"assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","run_id":"run_32c3e83920e44c1fbca7beb8fe114154"},"base_resp":{}}


data: {"data":{"id":"msg_03ddc0dc00bb41c192baef8f5888cd25","object":"message","created_at":1709280118,"thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","role":"ai","content":[{"type":"text","text":{"value":"","annotations":[{"type":"web_citation","text":"【1†source】","start_index":82,"end_index":92,"web_citation":{"url":"http://politics.people.com.cn/n1/2024/0301/c1001-40186885.html","quote":"3月1日 召开的中国人民政治协商会议第十四届全国委员会常务委员会第五次会议决定：中国人民政治协商会议第十四届全国委员会第二次会议于2024年 "}},{"type":"web_citation","text":"【2†source】","start_index":195,"end_index":205,"web_citation":{"url":"https://news.bjd.com.cn/2024/03/01/10711880.shtml","quote":"北京市大兴区消防救援支队2月29日发布演练公告，该队将于2024年3月1日上午10时，在大兴区香海园小区开展灭火救援综合实战演练。实战演练期间将使用油盆火和烟雾弹，届时将有较多消防车辆在周边道路通行和集结，请广大市民朋友不要惊慌，注意避让。"}},{"type":"web_citation","text":"【4†source】","start_index":281,"end_index":291,"web_citation":{"url":"https://news.cctv.com/2024/02/29/ARTIczSJVNv6cMnzwSeXkkLV240229.shtml","quote":"3月1日起，横琴粤澳深度合作区将正式封关运行。合作区实施范围为横琴岛“一线”和“二线”之间的海关监管区域，总面积约106平方公里。其中，横琴与澳门特别行政区之间设为“一线”，横琴与中华人民共和国关境内其他地区之间设为“二线”。"}},{"type":"web_citation","text":"【5†source】","start_index":316,"end_index":326,"web_citation":{"url":"https://tv.cctv.com/2024/03/01/VIDEcbvaIKj9QNjGK9oVtDeV240301.shtml","quote":"00:02:39. [新闻直播间]俄罗斯 俄宣布从3月1日起禁止汽油出口半年. 00:00:23. [新闻直播间]阿布扎比 世贸组织第13届部长级会议将延期闭幕. 00:00:36. [新闻直播间]韩国实习和住院医生集体辞职 韩国政府再次呼吁辞职医生返岗. 00:02:00. [新闻直播间]韩国实习和住院医生 "}},{"type":"web_citation","text":"【8†source】","start_index":402,"end_index":412,"web_citation":{"url":"https://news.cctv.com/2024/02/28/ARTICxLo5PKd4yz8wjnVAKfT240228.shtml","quote":"新修改的发票管理办法实施细则，3月1日起施行. 国家税务总局公布《关于修改〈中华人民共和国发票管理办法实施细则〉的决定》（国家税务总局令第56号），自2024年3月1日起施行。. 修改的主要内容有：一是明确电子发票基本管理规定， 强调“电子发票与纸质 "}}]},"id":"content_0bd78e5df96545b3ad951f1cb6c8339a"},{"type":"audio_file","id":"content_426c9f71a74646f0a03166410fa60973","audio_file":{"file_id":"","value":""}}],"assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","run_id":"run_32c3e83920e44c1fbca7beb8fe114154"},"base_resp":{}}


data: {"data":{"assistant_id":"asst_c4e6732f29274a50b41493a5ecb4d6e2","created_at":1709280116,"id":"run_32c3e83920e44c1fbca7beb8fe114154","instructions":"是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题\n回答中若参考了函数返回结果中的信息，请在该句末尾用【1†source】【2†source】【3†source】的形式标注信息来源。严禁使用其他的引用格式（如[1],【1】,²等等）。\n当前时间：2024-03-01, 周五。","metadata":{"\u003c/NUM_INPUT_MSGS\u003e":"1"},"model":"abab6.5s-chat","object":"run","status":"completed","t2a_control":{"combine":false,"model":"speech-01","oss":"","return_bytes":false,"text":"","timbre_weights":{"male-qn-qingse":1}},"thread_id":"thread_ae55415b12e040b4bdc50ec82005d2f9","tools":[{"type":"web_search"}]},"base_resp":{"status_code":0}}
```
###### 完整示例
**Tab: 文本输入，支持Web Search（流式输出）**

**以文本输入，回复的文本以流式输出，并支持Web Search工具完整示例：**
```python
import json
import requests

GroupId = "请输入您的group id"
API_KEY = "请输入您的api key"

headers = {
    'Authorization': f'Bearer {API_KEY}',
    'Content-Type': 'application/json'
}

# 流程一：创建助手
def create_assistant():
    url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"
    payload = json.dumps({
        "model": "abab6.5s-chat",  # 模型版本
        "name": "AI搜索引擎",  # 助手名称
        "description": "AI搜索引擎，用在搜索场景",  # 助手描述
        "instructions": "是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题",  # 助手设定（即bot_setting)
        "tools": [{"type": "web_search"}]
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程二：创建线程
def create_thread():
    url = f"https://api.minimax.chat/v1/threads/create?GroupId={GroupId}"
    response = requests.post(url, headers=headers)
    return response.json()

# 流程三：创建流式run
def stream_run(thread_id, assistant_id, query):
    url = f"https://api.minimax.chat/v1/threads/run/create_stream?GroupId={GroupId}"
    headers = {
        'Authorization': f'Bearer {API_KEY}',
        'Content-Type': 'application/json'
    }
    payload = json.dumps({
        "stream": 1,
        "thread_id": thread_id,
        "assistant_id": assistant_id,
        "messages": [
            {
                "type": 1,
                "role": "user",
                "content": query
            }
        ],
        "model": "abab6.5s-chat",
        "tools": [
            {
                "type": "web_search"
            }
        ]
    })
    response = requests.request("POST", url, headers=headers, data=payload, stream=True)
    for chunk in response.raw:
        if chunk:
            chunk_str = chunk.decode("utf-8")
            print(str(chunk_str))


# 主流程
def main():
    # 创建助手
    assistant_response = create_assistant()
    assistant_id = assistant_response.get('id', '')
    print("assistant_id:", assistant_id)

    # 创建线程
    thread_response = create_thread()
    thread_id = thread_response.get('id', '')
    print("thread_id:", thread_id)

    # 在此填写query
    query = "今天上海天气怎么样？"
    
    # 获取回复
    response = stream_run(thread_id, assistant_id, query)
    

if __name__ == "__main__":
    main()
```

**Tab: 语音通话，支持Web Search（流式输出）**

**语音通话（音频内容hex编码输入输出），支持Web Search工具，并以流式输出完整示例：**
```python
import json
import requests

group_id = "请填写你的group id"
api_key = "请填写你的api key"

headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {api_key}"
}


class BaseError(Exception):
    msg = ""

    def __init__(self, msg):
        super().__init__(self)
        self.msg = msg


def create_assistant():
    url = f"https://api.minimax.chat/v1/assistants/create?GroupId={group_id}"
    payload = json.dumps({
        "model": "abab6.5s-chat",
        "name": "新闻记者",
        "description": "新闻记者，用于咨询各类新闻",
        "instructions": "是一个新闻记者，擅长解答关于新闻的各类问题",
        "tools": [
            {
                "type": "web_search"
            }
        ]
    })
    response = requests.request("POST", url, headers=headers, data=payload)
    if response.status_code != 200:
        raise BaseError(response.status_msg)
    data = json.loads(response.text)
    if data["base_resp"]["status_code"] != 0:
        raise BaseError(data["base_resp"]["status_msg"])
    return data["id"]


def create_thread():
    url = f"https://api.minimax.chat/v1/threads/create?GroupId={group_id}"
    payload = json.dumps({})
    response = requests.request("POST", url, headers=headers, data=payload)
    if response.status_code != 200:
        raise BaseError(response.status_msg)
    data = json.loads(response.text)
    if data["base_resp"]["status_code"] != 0:
        raise BaseError(data["base_resp"]["status_msg"])
    return data["id"]


def create_run_stream(assistant_id, thread_id):
    url = f"https://api.minimax.chat/v1/threads/run/create_stream?GroupId={group_id}"
    payload = json.dumps({
        "stream": 2,
        "thread_id": thread_id,
        "assistant_id": assistant_id,
        "messages": [
            {
                "type": 2,
                "role": "user",
                "content": ""	#	这里填写hex编码格式的音频内容
            }
        ],
        "model": "abab6.5s-chat",
        "t2a_option": {
            "model": "speech-01",
            "voice_id": "male-qn-qingse"
        },
        "tools": [
            {
                "type": "web_search"
            }
        ]
    })

    audio_contents = {}

    response = requests.request("POST", url, headers=headers, data=payload, stream=True)
    if response.status_code != 200:
        raise BaseError(response.status_msg)
    for chunk in response.raw:
        chunk = chunk.decode("utf-8")
        if chunk.startswith("data: "):
            print(chunk)
            data = json.loads(chunk[6:])
            if "status_code" in data["base_resp"] and data["base_resp"]["status_code"] != 0:
                raise BaseError(data["base_resp"]["status_msg"])
            if data["data"]["object"] == "message" and data["data"]["role"] == "ai":
                for content in data["data"]["content"]:
                    if "audio_file" in content and "value" in content["audio_file"]:
                        if content["id"] not in audio_contents:
                            audio_contents[content["id"]] = ""
                        audio_contents[content["id"]] += content["audio_file"]["value"]
            if "status_code" in data["base_resp"] and data["base_resp"]["status_code"] == 0:
                return audio_contents
        elif chunk != "\n":
            data = json.loads(chunk)
            raise BaseError(data["base_resp"]["status_msg"])


def main():
    try:
        assistant_id = create_assistant()
        print("assistant id:" + assistant_id)
        thread_id = create_thread()
        print("thread id:" + thread_id)
        audio_contents = create_run_stream(assistant_id, thread_id)
        index = 0
        for audio_content in audio_contents.values():
            audio = bytes.fromhex(audio_content)
            filename = "temp_" + str(index) + ".mp3"
            with open(filename, "wb") as file:
                index += 1
                file.write(audio)
        print("end")

    except BaseError as e:
        print("error: " + e.msg)
    except Exception:
        print("unkonwn error")
        exit(1)


if __name__ == "__main__":
    main()
```
**Tab: 输出示例（文本输入）**

**以文本输入，回复的文本以流式输出，并支持Web Search工具完整示例：**
```python
assistant_id: asst_f258f10761f74f8087611f233e41fc3f
thread_id: thread_58188c8076ef4354899fee25d1b89dec
data: {"data":{"assistant_id":"asst_f258f10761f74f8087611f233e41fc3f","created_at":1709201106,"id":"run_6528035e904a4c21b9b12acee7d2c624","instructions":"是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题\n回答中若参考了函数返回结果中的信息，请在该句末尾用【1†source】【2†source】【3†source】的形式标注信息来源。严禁使用其他的引用格式（如[1],【1】,²等等）。\n当前时间：2024-02-29, 周四。","metadata":{"\u003c/NUM_INPUT_MSGS\u003e":"1"},"model":"abab6.5s-chat","object":"run","status":"in_progress","t2a_control":{"combine":false,"model":"speech-01","oss":"","return_bytes":false,"text":"","timbre_weights":{"male-qn-qingse":1}},"thread_id":"thread_58188c8076ef4354899fee25d1b89dec","tools":[{"type":"web_search"}]},"base_resp":{}}



data: {"data":{"id":"step_28d52c63d5b4442f87f2b37c4ea5c2cf","object":"run step","created_at":1709201107,"run_id":"run_6528035e904a4c21b9b12acee7d2c624","assistant_id":"asst_f258f10761f74f8087611f233e41fc3f","thread_id":"thread_58188c8076ef4354899fee25d1b89dec","type":"tool_calls","status":"completed","step_details":{"type":"tool_calls","tool_calls":[{"id":"call_2f9ea5c10a614286ae38bdbce0d51c99","type":"web_search","web_search":{"name":"","arguments":""}}]}},"base_resp":{}}



data: {"data":{"id":"step_89c758dffb9246f39be2fc7e8c2ba3af","object":"run step","created_at":1709201108,"run_id":"run_6528035e904a4c21b9b12acee7d2c624","assistant_id":"asst_f258f10761f74f8087611f233e41fc3f","thread_id":"thread_58188c8076ef4354899fee25d1b89dec","type":"message_creation","status":"completed","completed_at":1709201108,"step_details":{"type":"message_creation","message_creation":{"message_id":"msg_5098f62a8d614c2f9de49cf794e300ba"}}},"base_resp":{}}



data: {"data":{"id":"msg_5098f62a8d614c2f9de49cf794e300ba","object":"message","created_at":1709201108,"thread_id":"thread_58188c8076ef4354899fee25d1b89dec","role":"ai","content":[{"type":"text","text":{"value":"今天上海的天气情况如下：\n- 温度：最高温度约为10°C，最低温度约为4°C【1†source】。\n- 天气：多云【2†source】。\n- 空气质量：平均空气质量指数为56，空气状况良好【3†source】。\n- 紫外线强度：紫外线强度很弱，建议出门前涂擦SPF在6-7之间、PA+的防晒护肤品【1†source】。\n- 洗车指数：不宜洗车，未来24小时内有雨，如果在此期间洗车，雨水和路上的泥水可能会再次弄脏您的爱车【1†source】。\n\n以上信息为今日上海的天气情况","annotations":null},"id":"content_c204068a07ec4551b8c8025a79e9e4b4"}],"assistant_id":"asst_f258f10761f74f8087611f233e41fc3f","run_id":"run_6528035e904a4c21b9b12acee7d2c624"},"base_resp":{}}



data: {"data":{"id":"msg_5098f62a8d614c2f9de49cf794e300ba","object":"message","created_at":1709201108,"thread_id":"thread_58188c8076ef4354899fee25d1b89dec","role":"ai","content":[{"type":"text","text":{"value":"，请您注意天气变化，做好出行准备。","annotations":[{"type":"web_citation","text":"【1†source】","start_index":210,"end_index":220,"web_citation":{"url":"https://lishi.tianqi.com/shanghai/20240229.html","quote":"当前位置： 天气网 历史天气查询 上海历史天气. 上海2月29日生活指数. 不宜洗车，未来24小时内有雨，如果在此期间洗车，雨水和路上的泥水可能会再次弄脏您的爱车。. 紫外线强度很弱，建议出门前涂擦SPF在6-7之间、PA+的防晒护肤品。. 天气条件较易诱发过敏 "}},{"type":"web_citation","text":"【2†source】","start_index":57,"end_index":67,"web_citation":{"url":"https://lishi.tianqi.com/shanghai/202402.html","quote":"27. 空气最好 (02/01) 45. 空气最差 (02/03) 日期. 最高气温. 最低气温. 天气. 风向. 2024-02-01 星期四. 10℃. 6℃. 大雨. 东南风 1级. 查看更多. 上海气温走势图 (2024-02-01到2024-02-01) 上海附近城市空气质量指数. 闵行空气质量指数. 宝山空气质量指数. 嘉定空气质量指数. 金山空气质量指数. 青浦空气质量指数. 松江空气质量指数. 奉贤空气质量指数. 虹口空气质量指数. 黄浦空气质量指数."}},{"type":"web_citation","text":"【3†source】","start_index":94,"end_index":104,"web_citation":{"url":"https://www.tianqi24.com/shanghai/history202402.html","quote":"自2024-02-01到2024-02-25，上海统计数据分析：. 雨8天，阴5天，多云4天，晴4天，沙雾霾3天，雪1天；. 平均高温：12℃， 平均低温：3℃；. 最高温度 (02月14日)：23℃， 最低温度 (02月09日)：-1℃；. 平均PM25值是: 39 (优) 上海2024年02月总共有降雨量至少（6）mm以上. [切换 "}},{"type":"web_citation","text":"【1†source】","start_index":210,"end_index":220,"web_citation":{"url":"https://lishi.tianqi.com/shanghai/20240229.html","quote":"当前位置： 天气网 历史天气查询 上海历史天气. 上海2月29日生活指数. 不宜洗车，未来24小时内有雨，如果在此期间洗车，雨水和路上的泥水可能会再次弄脏您的爱车。. 紫外线强度很弱，建议出门前涂擦SPF在6-7之间、PA+的防晒护肤品。. 天气条件较易诱发过敏 "}},{"type":"web_citation","text":"【1†source】","start_index":210,"end_index":220,"web_citation":{"url":"https://lishi.tianqi.com/shanghai/20240229.html","quote":"当前位置： 天气网 历史天气查询 上海历史天气. 上海2月29日生活指数. 不宜洗车，未来24小时内有雨，如果在此期间洗车，雨水和路上的泥水可能会再次弄脏您的爱车。. 紫外线强度很弱，建议出门前涂擦SPF在6-7之间、PA+的防晒护肤品。. 天气条件较易诱发过敏 "}}]},"id":"content_c204068a07ec4551b8c8025a79e9e4b4"}],"assistant_id":"asst_f258f10761f74f8087611f233e41fc3f","run_id":"run_6528035e904a4c21b9b12acee7d2c624"},"base_resp":{}}



data: {"data":{"assistant_id":"asst_f258f10761f74f8087611f233e41fc3f","created_at":1709201106,"id":"run_6528035e904a4c21b9b12acee7d2c624","instructions":"是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题\n回答中若参考了函数返回结果中的信息，请在该句末尾用【1†source】【2†source】【3†source】的形式标注信息来源。严禁使用其他的引用格式（如[1],【1】,²等等）。\n当前时间：2024-02-29, 周四。","metadata":{"\u003c/NUM_INPUT_MSGS\u003e":"1"},"model":"abab6.5s-chat","object":"run","status":"completed","t2a_control":{"combine":false,"model":"speech-01","oss":"","return_bytes":false,"text":"","timbre_weights":{"male-qn-qingse":1}},"thread_id":"thread_58188c8076ef4354899fee25d1b89dec","tools":[{"type":"web_search"}]},"base_resp":{"status_code":0}}
```

**Tab: 输出示例（语音通话）**

**语音通话（音频内容hex编码输入输出），支持Web Search工具，并以流式输出完整示例：**
```python
assistant id:asst_d787afc30780449aa0f1d55212fb76f0
thread id:thread_d67f72be0b4f4be588cd4329deec2e9b
data: {"data":{"assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","created_at":1709382258,"id":"run_53dedd40be4f40c4884873a54e9d550c","instructions":"是一个新闻记者，擅长解答关于新闻的各类问题\n回答中若参考了函数返回结果中的信息，请在该句末尾用【1†source】【2†source】【3†source】的形式标注信息来源。严禁使用其他的引用格式（如[1],【1】,²等等）。\n当前时间：2024-03-02, 周六。","metadata":{"\u003c/NUM_INPUT_MSGS\u003e":"1"},"model":"abab6.5s-chat","object":"run","status":"in_progress","t2a_control":{"combine":false,"model":"speech-01","oss":"","return_bytes":false,"text":"","timbre_weights":{"male-qn-qingse":1}},"thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","tools":[{"type":"web_search"}]},"base_resp":{}}

# 下方data包中包含ASR识别结果（"今天的新闻。"）
data: {"data":{"id":"msg_e800e4e4210345cbbc310ebc3d7901af","object":"message","created_at":1709382257,"thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","role":"user","content":[{"type":"audio_file","audio_file":{"file_id":"113943710703746"}},{"type":"text","text":{"value":"今天的新闻。","annotations":[]}}],"assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","run_id":"run_53dedd40be4f40c4884873a54e9d550c","metadata":{"\u003c/ASR_AUDIO_DURATION\u003e":"1"}},"base_resp":{}}

data: {"data":{"id":"step_41b471be862b42719a7bd9f4dc3e4c9e","object":"run step","created_at":1709382260,"run_id":"run_53dedd40be4f40c4884873a54e9d550c","assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","type":"tool_calls","status":"completed","step_details":{"type":"tool_calls","tool_calls":[{"id":"call_637d7c546db64bd5adcfacb2a1e0e215","type":"web_search","web_search":{"name":"","arguments":""}}]}},"base_resp":{}}

data: {"data":{"id":"step_613c94130f4f43d081b9e8d89f3b7192","object":"run step","created_at":1709382260,"run_id":"run_53dedd40be4f40c4884873a54e9d550c","assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","type":"message_creation","status":"completed","completed_at":1709382260,"step_details":{"type":"message_creation","message_creation":{"message_id":"msg_587a133dc8934fc396b39b867e74c9a3"}}},"base_resp":{}}

# 下方数个data包中包含模型将要输出音频对应的文本内容
data: {"data":{"id":"msg_587a133dc8934fc396b39b867e74c9a3","object":"message","created_at":1709382260,"thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","role":"ai","content":[{"type":"text","text":{"value":"今天（2024年3月2日）的主要新闻包括：\n\n1. 中国人民政治协商会议第十四届全国委员会常务委员会第五次会议决定，中国人民政治协商会议第十四届全国委员会第二次会议将于2024年3月举行【1†source】。\n\n2. 十四届全国人大二次会议将于2024年3月4日（星期一）12时在人民大会堂新闻发布厅举行新闻发布会，由大会发言人就大会议程和人大工作相关问题进行说明【2†source】。\n\n3. 设在北京西长安街北侧梅地亚中心的两会新闻中心已正式启用，目前已有3000","annotations":null},"id":"content_a8d4c71a221e425ebcab8c6602d391cf"}],"assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","run_id":"run_53dedd40be4f40c4884873a54e9d550c"},"base_resp":{}}

data: {"data":{"id":"msg_587a133dc8934fc396b39b867e74c9a3","object":"message","created_at":1709382260,"thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","role":"ai","content":[{"type":"audio_file","id":"content_67444ec62aa54495a6a96d9752a571ed","audio_file":{"file_id":"","value":"###这里是输出音频内容的hex编码###"}},{"type":"text","text":{"value":"多名中外记者报名采访全国两会【3†source】。\n\n4. 第32届华东进出口商品交易会已于3月1日在上海开幕，3000多家出口企业齐聚，抢抓外贸新机遇【4†source】。\n\n以上为今天的部分新闻摘要。","annotations":null},"id":"content_a8d4c71a221e425ebcab8c6602d391cf"}],"assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","run_id":"run_53dedd40be4f40c4884873a54e9d550c"},"base_resp":{}}

data: {"data":{"id":"msg_587a133dc8934fc396b39b867e74c9a3","object":"message","created_at":1709382260,"thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","role":"ai","content":[{"type":"audio_file","id":"content_67444ec62aa54495a6a96d9752a571ed","audio_file":{"file_id":"","value":"###这里是输出音频内容的hex编码###"}}],"assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","run_id":"run_53dedd40be4f40c4884873a54e9d550c"},"base_resp":{}}

data: {"data":{"id":"msg_587a133dc8934fc396b39b867e74c9a3","object":"message","created_at":1709382260,"thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","role":"ai","content":[{"type":"text","text":{"value":"","annotations":[{"type":"web_citation","text":"【1†source】","start_index":93,"end_index":103,"web_citation":{"url":"https://news.sina.com.cn/zx/gj/2024-03-01/doc-inakvvhe8163370.shtml","quote":"人民网北京3月1日电（黄子娟、彭静）2024年2月5日至25日，人民网连续第23次开展全国两会调查，广纳民情民意。  图片新闻 山西运城迎登高节 蛟龙 "}},{"type":"web_citation","text":"【3†source】","start_index":238,"end_index":248,"web_citation":{"url":"http://politics.people.com.cn/n1/2024/0301/c1001-40186885.html","quote":"3月1日召开的中国人民政治协商会议第十四届全国委员会常务委员会第五次会议决定：中国人民政治协商会议第十四届全国委员会第二次会议于2024年3 "}}]},"id":"content_a8d4c71a221e425ebcab8c6602d391cf"},{"type":"audio_file","id":"content_67444ec62aa54495a6a96d9752a571ed","audio_file":{"file_id":"","value":""}}],"assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","run_id":"run_53dedd40be4f40c4884873a54e9d550c"},"base_resp":{}}

data: {"data":{"assistant_id":"asst_d787afc30780449aa0f1d55212fb76f0","created_at":1709382258,"id":"run_53dedd40be4f40c4884873a54e9d550c","instructions":"是一个新闻记者，擅长解答关于新闻的各类问题\n回答中若参考了函数返回结果中的信息，请在该句末尾用【1†source】【2†source】【3†source】的形式标注信息来源。严禁使用其他的引用格式（如[1],【1】,²等等）。\n当前时间：2024-03-02, 周六。","metadata":{"\u003c/NUM_INPUT_MSGS\u003e":"1"},"model":"abab6.5s-chat","object":"run","status":"completed","t2a_control":{"combine":false,"model":"speech-01","oss":"","return_bytes":false,"text":"","timbre_weights":{"male-qn-qingse":1}},"thread_id":"thread_d67f72be0b4f4be588cd4329deec2e9b","tools":[{"type":"web_search"}]},"base_resp":{"status_code":0}}

end
```
##### 非流式输出
###### 创建assistant
可以通过配置以下参数，构建一个assistant得到对应的assistants id，以响应用户的请求：
- instructions：代表的是对assistant的背景设定；
- model：该接口目前支持`abab5.5s-chat`、`abab5.5-chat`、`abab6.5-chat`、`abab6.5s-chat`以及Finetune model: `abab5.5-chat-240131`、`abab5.5s-chat-240123`；
- tools：该接口在`abab5.5-chat`、`abab6.5-chat`、`abab6.5s-chat`、`abab5.5-chat-240131`下支持`Code Interpreter（代码解释器）`、`Retrieval（知识库检索）`、`Function Calling（函数调用）`、`Web Search（网络搜索）功能`，并支持开启多个工具，abab5.5s-chat、abab5.5s-chat-240123暂不支持。
如果使用了Retrieval，在创建assistant时可能会消耗几分钟的时间用于进行向量化和入库，在创建run时需要等待assistant创建完成，可以通过retrieve assistant 查看创建状态。
**Tab: Python请求示例**

**创建assistant（非流式）**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"

payload = json.dumps({
  "model": "abab5.5-chat",
  "name": "小说解读专家",
  "description": "小说解读专家，用在小说解读场景中回答用户问题",
  "instructions": "是一个小说解读专家，读过万卷书，了解小说里的各种细致情节，另外也非常热心，会热情的帮读者解答小说里的问题",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  },
  "tools": [
    {
      "type": "code_interpreter"
    },
    {
      "type": "retrieval"
    },
    {
      "type": "web_search"
    }
  ]
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell请求示例**

**创建assistant（非流式）**
```shell
curl --location 'https://api.minimax.chat/v1/assistants/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "model": "abab5.5-chat",
    "name": "小说解读专家",
    "description": "小说解读专家，用在小说解读场景中回答用户问题",
    "instructions": "是一个小说解读专家，读过万卷书，了解小说里的各种细致情节，另外也非常热心，会热情的帮读者解答小说里的问题",
    "metadata": {
        "key1": "value1",
        "key2": "value2"
    },
    "tools": [
        {"type":"code_interpreter"},
        {"type":"retrieval"},
        {"type":"web_search"}
    ]
}'
```
###### 创建thread
我们建议在用户发出请求之后立刻为assistant id创建一个thread id，并添加message，在此thread id中传递关于用户请求的上下文和文件。在后续的请求中，如果您判断是一次连续的对话，则无需自己拼接上下文，只需将最新的message添加到对应的thread id 即可得到包含了thread id 历史上下文的回复，历史上下文超过我们会帮您自动截断。另外，thread没有数值限制，您可以向thread添加任意数量的message。
**Tab: Python请求示例**

**创建thread（非流式）**
```python
import requests
import json
GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"
url = f"https://api.minimax.chat/v1/threads/create?GroupId={GroupId}"
payload = json.dumps({
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  }
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}
response = requests.request("POST", url, headers=headers, data=payload)
print(response.text)
```

**Tab: Shell请求示例**

**创建thread（非流式）**
```shell
curl --location 'https://api.minimax.chat/v1/threads/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "metadata": {
        "key1": "value1",
        "key2": "value2"
    }
}'
```
###### 添加message
1. 向thread添加message
message包含了文本，以及允许用户按照自己的需求上传file。message需要添加到特定的thread中。
在message层级添加的文件，仅支持code interpreter工具的使用，如要使用retrieval工具则需在assistants层级添加文件，方可使用。
目前支持通过message添加文本content，此外一个message中最多可附加20个文件。
**Tab: Python请求示例**

**向thread添加message**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/messages/add?GroupId={GroupId}"

payload = json.dumps({
  "thread_id": "thread_123abc",
  "role": "user",
  "content": "在倚天屠龙记中，无俗念词的作者是谁？",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  },
  "file_ids": [
    # "此处添加file_ids"
  ]
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell请求示例**

**向thread添加message**
```shell
curl --location 'https://api.minimax.chat/v1/threads/messages/add?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "thread_123abc",
    "role": "user",
    "content": "在倚天屠龙记中，无俗念词的作者是谁？",
    "metadata": {
        "key1": "value1",
        "key2": "value2"
    },
    "file_ids": [
      # "此处添加file_ids"
  ]
}'
```
2. 列表message
您可以在thread中列出message，此时您可以看到您的message已被添加进thread。
**Tab: Python请求示例**

**列表message**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/messages/list?GroupId={GroupId}"

payload = json.dumps({
  "thread_id": "thread_123abc"
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell请求示例**

**列表message**
```shell
curl --location --request GET 'https://api.minimax.chat/v1/threads/messages/list?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
"thread_id": "$Thread_Id"
}'
```
###### 运行assistant
要使创建的assistant响应用户的请求消息，您还需要创建一个run来运行，请确保您的assistant id和thread id 都已创建完成。
此时assistant将会读取thread并决定是否调用工具（若构建assistant时工具已启用）。
随着run的进行，assistant会将模型回复的message添加到thread中。
**Tab: Python请求示例**

**运行assistant**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/run/create?GroupId={GroupId}"

payload = json.dumps({
  "thread_id": "$Thread_Id",
  "assistant_id": "$Assistant_Id"
})

headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell请求示例**

**运行assistant**
```shell
curl --location 'https://api.minimax.chat/v1/threads/run/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "$Thread_Id",
    "assistant_id": "$Assistant_Id"
}'
```
###### 检查运行状态
在默认的情况下，创建的run将进入排队状态。您可以定期检索run id的运行状态，以查看它是否运行完成。
**Tab: Python请求示例**

**检查运行状态**
```python
import requests
import json
​
GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"
​
url = f"https://api.minimax.chat/v1/threads/run/retrieve?GroupId={GroupId}"
​
payload = json.dumps({
  "thread_id": "thread_123abc", # 替换为真实thread_id
  "run_id": "run_123abc"  # 替换为真实run_id
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json'
}
​
response = requests.request("GET", url, headers=headers, data=payload)
​
print(response.text)
​
```

**Tab: Shell请求示例**

**检查运行状态**
```shell
curl --location --request GET 'https://api.minimax.chat/v1/threads/run/retrieve?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "$Thread_Id",
    "run_id": "$Run_Id"
}'
```
###### 查看回复内容
创建的run id运行完成后，您可以列出assistant添加到thread的用户message来查看请求回复的内容。
**Tab: Python请求示例**

**查看回复内容**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/messages/list?GroupId={GroupId}"

payload = json.dumps({
  "thread_id": "thread_123abc" # 替换为真实thread_id
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell请求示例**

**查看回复内容**
```shell
curl --location --request GET 'https://api.minimax.chat/v1/threads/messages/list?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "$Thread_Id"
}'
```
###### 完整示例
包含文档上传、创建assistant、创建thread、添加message、运行assistant、检查运行状态和查看回复内容等流程。
**Tab: Retrieval示例**

**请求示例**
```python
import requests
import json
import time

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

headers = {
    'Authorization': f'Bearer {API_KEY}',
    'Content-Type': 'application/json'
}
headers_retrieval = {
    'Authorization': f'Bearer {API_KEY}',
    'authority': 'api.minimax.chat',
}


# 流程零：上传文档
def create_file():
    url = f"https://api.minimax.chat/v1/files/upload?GroupId={GroupId}"
    data = {
        'purpose': 'assistants'
    }

    files = {
        'file': open('./倚天屠龙记.txt', 'rb')
    }
    response = requests.post(url, headers=headers_retrieval, data=data,files=files)
    return response.json()

#流程一：创建助手
def create_assistant(file_id):
    url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"
    payload = json.dumps({
        "model": "abab6.5s-chat",# 模型版本
        "name": "小说专家", # 助手名称
        "description": "小说专家，用在小说解读场景中回答用户问题", # 助手描述
        "instructions": "是一个小说专家，读过万卷书，了解小说里的各种细致情节，另外也非常热心，会热情的帮读者解答小说里的问题",# 助手设定（即bot_setting)
        "file_ids": [
            str(file_id)
        ],
        "tools": [{"type": "retrieval"}]

    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程二：创建线程
def create_thread():
    url = f"https://api.minimax.chat/v1/threads/create?GroupId={GroupId}"
    response = requests.post( url, headers=headers)
    return response.json()

# 流程三：往线程里添加信息
def add_message_to_thread(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/add?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "role": "user",
        "content": "在倚天屠龙记中，无俗念词的作者是谁？",
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程四：使用助手处理该线程
def run_thread_with_assistant(thread_id, assistant_id):
    time.sleep(200) #创建assistants进行向量化以及存储时需要一定的时间，可以考虑使用retrieve assistant检索是否创建成功
    url = f"https://api.minimax.chat/v1/threads/run/create?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "assistant_id": assistant_id
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程五：获取线程中助手处理出的新信息
def get_thread_messages(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/list?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id
    })
    response = requests.get(url, headers=headers, data=payload)
    return response.json()
def check_thread_run_status(thread_id, run_id):
    url = f"https://api.minimax.chat/v1/threads/run/retrieve?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": str(thread_id),
        "run_id": str(run_id)
    })
    completed = False
    while not completed:
        response = requests.request("GET", url, headers=headers, data=payload)
        if response.status_code == 200:
            response_data = response.json()
            status = response_data.get('status', '')
            print(f"Status: {status}")
            if status == 'completed':
                completed = True
                print("Process completed, exiting loop.")
            else:
                time.sleep(2)  # 如果状态不是completed，等待两秒后重新请求
        else:
            print(f"Error: {response.status_code}")
            break  # 如果请求失败，退出循环
    return completed
# 主流程
def main():
    # 上传文档
    file_response = create_file()
    file_id = file_response.get('file', {}).get('file_id')
    print("file_id:",file_id)

    # 创建助手
    assistant_response = create_assistant(file_id)
    assistant_id = assistant_response.get('id', '')
    print("assistant_id:",assistant_id)

    # 创建线程
    thread_response = create_thread()
    thread_id = thread_response.get('id', '')
    print("thread_id:",thread_id)

    # 往线程里添加信息
    add_message_to_thread(thread_id)  # 不保存返回值

    # 使用助手处理该线程
    run_response = run_thread_with_assistant(thread_id, assistant_id)
    run_id = run_response.get('id', '')  # 假设run_response是正确的JSON响应，并包含run_id
    print("run_id:",run_id)

    # 检查助手处理状态
    if check_thread_run_status(thread_id, run_id):
        # 获取线程中助手处理出的新信息
        thread_messages_response = get_thread_messages(thread_id)
        # 打印JSON数据
        print(json.dumps(thread_messages_response, indent=4, ensure_ascii=False))

if __name__ == "__main__":
    main()
```

**Tab: Code Interpreter示例**

**请求示例**
```python
import requests
import json
import time

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

headers = {
    'Authorization': f'Bearer {API_KEY}',
    'Content-Type': 'application/json'
}
headers_retrieval = {
    'Authorization': f'Bearer {API_KEY}',
    'authority': 'api.minimax.chat',
}


# 流程零：上传文档
def create_file():
    url = f"https://api.minimax.chat/v1/files/upload?GroupId={GroupId}"
    data = {
        'purpose': 'assistants'
    }

    files = {
        'file': open('./M国gdp.csv', 'rb')
    }
    response = requests.post(url, headers=headers_retrieval, data=data,files=files)
    return response.json()

#流程一：创建助手
def create_assistant(file_id):
    url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"
    payload = json.dumps({
        "model": "abab6.5s-chat",# 模型版本
        "name": "数据分析师", # 助手名称
        "description": "数据分析师，用在数据分析场景", # 助手描述
        "instructions": "是一个数据分析师，善于读取文件进行数据分析，并能严谨精准的给出分析结论",# 助手设定（即bot_setting)
        "file_ids": [
            str(file_id)
        ],
        "tools": [{"type": "code_interpreter"}]

    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程二：创建线程
def create_thread():
    url = f"https://api.minimax.chat/v1/threads/create?GroupId={GroupId}"
    response = requests.post( url, headers=headers)
    return response.json()

# 流程三：往线程里添加信息
def add_message_to_thread(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/add?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "role": "user",
        "content": "M国的2000-2005年的gdp趋势是怎么样的，2005同比增长了还是下降了？",
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程四：使用助手处理该线程
def run_thread_with_assistant(thread_id, assistant_id):
    url = f"https://api.minimax.chat/v1/threads/run/create?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "assistant_id": assistant_id
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程五：获取线程中助手处理出的新信息
def get_thread_messages(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/list?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id
    })
    response = requests.get(url, headers=headers, data=payload)
    return response.json()
def check_thread_run_status(thread_id, run_id):
    url = f"https://api.minimax.chat/v1/threads/run/retrieve?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": str(thread_id),
        "run_id": str(run_id)
    })
    completed = False
    while not completed:
        response = requests.request("GET", url, headers=headers, data=payload)
        if response.status_code == 200:
            response_data = response.json()
            status = response_data.get('status', '')
            print(f"Status: {status}")
            if status == 'completed':
                completed = True
                print("Process completed, exiting loop.")
            else:
                time.sleep(2)  # 如果状态不是completed，等待两秒后重新请求
        else:
            print(f"Error: {response.status_code}")
            break  # 如果请求失败，退出循环
    return completed

# 流程六：获取助手处理过程的工具调用详细数据
def get_run_step(thread_id, run_id):
		url = f"https://api.minimax.chat/v1/threads/run_steps/list?GroupId={GroupId}"
    payload = json.dumps({
      "thread_id": thread_id,
      "run_id": run_id
    })
    headers = {
      'Authorization': f'Bearer {API_KEY}',
      'Content-Type': 'application/json'
    }
    response = requests.request("GET", url, headers=headers, data=payload)
		return response.json()

# 主流程
def main():
    # 上传文档
    file_response = create_file()
    file_id = file_response.get('file', {}).get('file_id')
    print("file_id:",file_id)

    # 创建助手
    assistant_response = create_assistant(file_id)
    assistant_id = assistant_response.get('id', '')
    print("assistant_id:",assistant_id)

    # 创建线程
    thread_response = create_thread()
    thread_id = thread_response.get('id', '')
    print("thread_id:",thread_id)

    # 往线程里添加信息
    add_message_to_thread(thread_id)  # 不保存返回值

    # 使用助手处理该线程
    run_response = run_thread_with_assistant(thread_id, assistant_id)
    run_id = run_response.get('id', '')  # 假设run_response是正确的JSON响应，并包含run_id
    print("run_id:",run_id)

    # 检查助手处理状态
    if check_thread_run_status(thread_id, run_id):
        # 获取线程中助手处理出的新信息
        thread_messages_response = get_thread_messages(thread_id)
        # 打印JSON数据
        print(json.dumps(thread_messages_response, indent=4, ensure_ascii=False))

				# 获取助手处理过程的工具调用详细数据(可查看模型生成的代码，如需，取消下方代码注释)
        # run_step = get_run_step(thread_id, run_id)
				# print(json.dumps(run_step, indent=4, ensure_ascii=False))

if __name__ == "__main__":
    main()
```

**Tab: Web Search示例**

**请求示例**
```python
import json
import time

import requests

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

headers = {
    'Authorization': f'Bearer {API_KEY}',
    'Content-Type': 'application/json'
}

#流程一：创建助手
def create_assistant():
    url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"
    payload = json.dumps({
        "model": "abab6.5s-chat",# 模型版本
        "name": "AI搜索引擎", # 助手名称
        "description": "AI搜索引擎，用在搜索场景", # 助手描述
        "instructions": "是一个AI搜索引擎，擅长利用网络搜索参考资料回答用的问题",# 助手设定（即bot_setting)
        "tools": [{"type": "web_search"}]

    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程二：创建线程
def create_thread():
    url = f"https://api.minimax.chat/v1/threads/create?GroupId={GroupId}"
    response = requests.post( url, headers=headers)
    return response.json()

# 流程三：往线程里添加信息
def add_message_to_thread(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/add?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "role": "user",
        "content": "上海迪士尼2024圣诞节会举办什么活动？",
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程四：使用助手处理该线程
def run_thread_with_assistant(thread_id, assistant_id):
    url = f"https://api.minimax.chat/v1/threads/run/create?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "assistant_id": assistant_id
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()

# 流程五：获取线程中助手处理出的新信息
def get_thread_messages(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/list?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id
    })
    response = requests.get(url, headers=headers, data=payload)
    return response.json()

def check_thread_run_status(thread_id, run_id):
    url = f"https://api.minimax.chat/v1/threads/run/retrieve?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": str(thread_id),
        "run_id": str(run_id)
    })
    completed = False
    while not completed:
        response = requests.request("GET", url, headers=headers, data=payload)
        if response.status_code == 200:
            response_data = response.json()
            status = response_data.get('status', '')
            print(f"Status: {status}")
            if status == 'completed':
                completed = True
                print("Process completed, exiting loop.")
            else:
                time.sleep(2)  # 如果状态不是completed，等待两秒后重新请求
        else:
            print(f"Error: {response.status_code}")
            break  # 如果请求失败，退出循环
    return completed

def post_process_message(message):
    message_content = message["content"][0]["text"]
    annotations = message_content["annotations"]
    citations = []

    # Iterate over the annotations and add footnotes
    for index, annotation in enumerate(annotations):
        # Replace the text with a footnote
        message_content["value"] = message_content["value"].replace(annotation["text"], f' [{index}]')
        # Gather citations based on annotation attributes
        if annotation["type"] == "file_citation":
            citations.append(f'[{index}] {annotation["file_citation"]["quote"]}')
        elif annotation["type"] == "web_citation":
            res = f"[{index}]"
            if annotation["web_citation"]["url"]:
                res += f' 网址：{annotation["web_citation"]["url"]}'
            if annotation["web_citation"]["quote"]:
                res += f' {annotation["web_citation"]["quote"]}'
            citations.append(res)

    message_content["value"] += '\n' + '\n'.join(citations)
    print(message_content["value"])


# 主流程
def main():
    # 创建助手
    assistant_response = create_assistant()
    assistant_id = assistant_response.get('id', '')
    print("assistant_id:",assistant_id)

    # 创建线程
    thread_response = create_thread()
    thread_id = thread_response.get('id', '')
    print("thread_id:",thread_id)

    # 往线程里添加信息
    add_message_to_thread(thread_id)  # 不保存返回值

    # 使用助手处理该线程
    run_response = run_thread_with_assistant(thread_id, assistant_id)
    run_id = run_response.get('id', '')  # 假设run_response是正确的JSON响应，并包含run_id
    print("run_id:",run_id)

    # 检查助手处理状态
    if check_thread_run_status(thread_id, run_id):
        # 获取线程中助手处理出的新信息
        thread_messages_response = get_thread_messages(thread_id)
        # 打印JSON数据
        print(json.dumps(thread_messages_response, indent=4, ensure_ascii=False))
        # 打印后处理后的信息
        print("\n模型回复：\n")
        print(post_process_message(thread_messages_response["data"][-1]))

if __name__ == "__main__":
    main()
```

**Tab: Function calling示例**

**请求示例**
```python
import json
import time

import requests

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

headers = {
    'Authorization': f'Bearer {API_KEY}',
    'Content-Type': 'application/json'
}


#流程一：创建助手
def create_assistant():
    url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"
    payload = json.dumps({
        "model": "abab6.5s-chat",# 模型版本
        "name": "function专家", # 助手名称
        "description": "根据用户定义的function，执行方案", # 助手描述
        "instructions": "根据用户定义的function，执行方案", # 助手设定（即bot_setting)
        "tools": [
            {
                "type": "function",
                "function": {
                    "name": "get_weather",
                    "description": "获得天气信息",
                    "parameters": {
                        "type": "object",
                        "required": ["location"],
                        "properties": {"location": {"type": "string", "description": "获得天气的地点"}},
                    },
                },
            }
        ],
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()


# 流程二：创建线程
def create_thread():
    url = f"https://api.minimax.chat/v1/threads/create?GroupId={GroupId}"
    response = requests.post( url, headers=headers)
    return response.json()


# 流程三：往线程里添加信息
def add_message_to_thread(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/add?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "role": "user",
        "content": "上海天气怎么样",
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()


# 流程四：使用助手处理该线程
def run_thread_with_assistant(thread_id, assistant_id):
    url = f"https://api.minimax.chat/v1/threads/run/create?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id,
        "assistant_id": assistant_id
    })
    response = requests.post(url, headers=headers, data=payload)
    return response.json()


# 流程五：获取线程中助手处理出的新信息
def get_thread_messages(thread_id):
    url = f"https://api.minimax.chat/v1/threads/messages/list?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": thread_id
    })
    response = requests.get(url, headers=headers, data=payload)
    return response.json()

# 流程六：查看run
def check_thread_run_status(thread_id, run_id, wanted_status):
    url = f"https://api.minimax.chat/v1/threads/run/retrieve?GroupId={GroupId}"
    payload = json.dumps({
        "thread_id": str(thread_id),
        "run_id": str(run_id)
    })
    completed = False
    max_count = 10
    count = 0
    while not completed and count < max_count:
        response = requests.request("GET", url, headers=headers, data=payload)
        if response.status_code == 200:
            response_data = response.json()
            status = response_data.get('status', '')
            print(f"Status: {status}")
            if status == wanted_status:
                completed = True
                print("Process completed, exiting loop.")
            else:
                time.sleep(2)  # 如果状态不是completed，等待两秒后重新请求
        else:
            print(f"Error: {response.status_code}")
            break  # 如果请求失败，退出循环
        count += 1
    return completed

# 流程七：查看required_action
def get_required_action(thread_id, run_id):
    url = f"https://api.minimax.chat/v1/threads/run/retrieve?GroupId={GroupId}"
    payload = json.dumps(
        {
            "thread_id": thread_id,
            "run_id": run_id
        }
    )
    response = requests.request("GET", url, headers=headers, data=payload)
    required_action = response.json()["required_action"]
    return required_action


# 流程八：提交结果
def submit_output(thread_id, run_id, required_action, output):
    url = f"https://api.minimax.chat/v1/threads/run/submit_tool_outputs?GroupId={GroupId}"
    payload = json.dumps(
        {
            "thread_id": thread_id,
            "run_id": run_id,
            "tool_outputs": [{
                "tool_call_id": required_action["submit_tool_outputs"]["tool_calls"][0]["id"],
                "output": output,
            }]
        }
    )
    response = requests.request("POST", url, headers=headers, data=payload)
    return response.json()

def main():
    # 创建助手
    assistant_response = create_assistant()
    assistant_id = assistant_response.get('id', '')
    print("assistant_id:",assistant_id)

    # 创建线程
    thread_response = create_thread()
    thread_id = thread_response.get('id', '')
    print("thread_id:",thread_id)

    # 往线程里添加信息
    add_message_to_thread(thread_id)  # 不保存返回值

    # 使用助手处理该线程
    run_response = run_thread_with_assistant(thread_id, assistant_id)
    run_id = run_response.get('id', '')  # 假设run_response是正确的JSON响应，并包含run_id
    print("run_id:",run_id)

    # 检查助手处理状态
    if check_thread_run_status(thread_id, run_id, "requires_action"):
        required_action = get_required_action(thread_id, run_id)
        # fc_output = "上海天气晴朗"
        fc_output = '{"result": "上海最高气温7度 最低零下一度 晴"}'
        resp = submit_output(thread_id, run_id, required_action, fc_output)
        print(resp)

    # 检查助手处理状态
    if check_thread_run_status(thread_id, run_id, "completed"):
        # 获取线程中助手处理出的新信息
        thread_messages_response = get_thread_messages(thread_id)
        # 打印JSON数据
        print(json.dumps(thread_messages_response, indent=4, ensure_ascii=False))

if __name__ == "__main__":
    main()
```
**Tab: Retrieval示例**

**返回示例**
```json
file_id: 891849227529
assistant_id: asst_9f399f168c445093853ba08332346e
thread_id: thread_8578198eb840cbaec7344b56ae2f8f
run_id: run_080adf40e20749b34e59225326f3
Status: in_progress
Status: in_progress
Status: completed
Process completed, exiting loop.
{
    "object": "list",
    "data": [
        {
            "id": "msg_ff35404520504e1787b08c74058b928d",
            "object": "message",
            "created_at": 1703419324,
            "thread_id": "thread_8578198eb840cbaec7344b56ae2f8f",
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "在倚天屠龙记中，无俗念词的作者是谁？",
                        "annotations": []
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_9f399f168c445093853ba08332346e",
            "run_id": "run_080adf40e20749b3914e5109225326f3",
            "metadata": null
        },
        {
            "id": "msg_01c12f5874174e87950bfa5f4dcd0d3c",
            "object": "message",
            "created_at": 1703419528,
            "thread_id": "thread_8578198eb840cbaec7344b56ae2f8f",
            "role": "ai",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "在《倚天屠龙记》中，无俗念词的作者是南宋末年的一位武学名家，有道之士，他姓丘，名处机，道号长春子，名列全真七子之一，是全真教中出类拔萃的人物。【1†source】",
                        "annotations": [
                            {
                                "type": "file_citation",
                                "text": "【1†source】",
                                "start_index": 71,
                                "end_index": 81,
                                "file_citation": {
                                    "file_id": "",
                                    "quote": "###《倚天屠龙记》\n\n###第01章 天涯思君不可忘\n\n“春游浩荡，是年年寒食，梨花时节。白锦无纹香烂漫，玉树琼苞堆雪。静夜沉沉，浮光霭霭，冷浸溶溶月。人间天上，烂银霞照通彻。\n\n浑似姑射真人，天姿灵秀，意气殊高洁。万蕊参差谁信道，不与群芳同列。浩气清英，仙才卓荦，下土难分别。瑶台归去，洞天方看清绝。”\n\n作这一首无俗念词的，乃南宋末年一位武学名家，有道之士。此人姓丘，名处机，道号长春子，名列全真七子之一，是全真教中出类拔萃的人物。词品评论此词道：“长春，世之所谓仙人也，而词之清拔如此”。\n\n这首词诵的似是梨花，其实同中真意却是赞誉一位身穿白衣的美貌少女，说她“浑似姑射真人，天姿灵秀，意气殊高洁”，又说她“浩气清英，仙才卓荦”，“不与群芳同列”。词中所颂这美女，乃古墓派传人小龙女。她一生爱穿白衣，当真如风拂玉树，雪裹琼苞，兼之生性清冷，实当得起“冷浸溶溶月”的形容，以“无俗念”三字赠之，可说十分贴切。长春子丘处机和她在终南山上比邻而居，当年一见，便写下这首词来。\n\n这时丘处机逝世已久，小龙女也已嫁与神雕大侠杨过为妻。在河南少室山山道之上，却另有一个少女，正在低低念诵此词。"
                                }
                            }
                        ]
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_9f399f168c445093853ba08332346e",
            "run_id": "run_080adf40e20749b34e59225326f3",
            "metadata": null
        }
    ],
    "first_id": "msg_ff35404520504e1787b08c74058b928d",
    "last_id": "msg_01c12f5874174e87950bfa5f4dcd0d3c",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
```

**Tab: Code Interpreter示例**

**返回示例**
```json
file_id: 892087299164
assistant_id: asst_e61ea32224284aacc40b36a6d61b
thread_id: thread_b6e8ef88333c4f148a84a5f45606
run_id: run_3c184429d47f414294403fc59c560
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: completed
Process completed, exiting loop.
{
    "object": "list",
    "data": [
        {
            "id": "msg_d8ca1b4cd4644f91a75807b5549a17",
            "object": "message",
            "created_at": 1703418313,
            "thread_id": "thread_b6e8ef88334f118c248a84a5f45606",
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "M国的2000-2005年的gdp趋势是怎么样的，2005同比增长了还是下降了？",
                        "annotations": []
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_e61ea32224284aacc40b36a6d61b",
            "run_id": "run_3c184429d47f414294403fc59c560",
            "metadata": null
        },
        {
            "id": "msg_d6166cd827134c93afb5c3c49204a784",
            "object": "message",
            "created_at": 1703418523,
            "thread_id": "thread_b6e8ef88333c4f148a84a5f45606",
            "role": "ai",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "根据提供的数据，M国的GDP在2000年至2005年间的趋势如下：\n\n- 2000年：GDP为1\n- 2001年：GDP为2\n- 2002年：GDP为3\n- 2003年：GDP为4\n- 2004年：GDP为5\n- 2005年：GDP为6\n\n因此，2005年的GDP相比2004年增长了。",
                        "annotations": []
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_e61ea32224284aacc40b36a6d61b",
            "run_id": "run_3c184429d47f414294403fc59c560",
            "metadata": null
        }
    ],
    "first_id": "msg_d8ca1b4cd4644f91a757f807b5549a17",
    "last_id": "msg_d6166cd827134c93afb5c3c49204a784",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
```

**Tab: Web Search示例**

**返回示例**
```json
assistant_id: asst_5e65f2aae3a044f2b9c26012aaed1a95
thread_id: thread_4f8d5e7dfd894d2186fe3baac45c0085
run_id: run_ef18b2b8185744c3a58453dc1020f497
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: in_progress
Status: completed
Process completed, exiting loop.
{
    "object": "list",
    "data": [
        {
            "id": "msg_66a208c1de5d49ab88b664adef704364",
            "object": "message",
            "created_at": 1703579772,
            "thread_id": "thread_4f8d5e7dfd894d2186fe3baac45c0085",
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "上海迪士尼2024圣诞节会举办什么活动？",
                        "annotations": []
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_5e65f2aae3a044f2b9c26012aaed1a95",
            "run_id": "run_ef18b2b8185744c3a58453dc1020f497",
            "metadata": null
        },
        {
            "id": "msg_8f428bd21e4b403182ab0c5f6ea47781",
            "object": "message",
            "created_at": 1703579795,
            "thread_id": "thread_4f8d5e7dfd894d2186fe3baac45c0085",
            "role": "ai",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "2024年上海迪士尼的圣诞节活动主要包括以下部分：\n\n1. 米奇、达菲和朋友们一起点亮了魔法奇观环绕的圣诞大树，在烟花的照耀下，照亮了大家的童话梦想。这一活动会在上海迪士尼度假区圣诞季期间（2023年11月23日至2024年1月1日）举行。\n\n2. 圣诞树寻踪活动，游客可以加入我们庆祝的行列，走入暖心的童话故事，漫步小镇街头，一起寻踪不一样的圣诞奇妙。完成圣诞树寻踪任务后，即可获得一份精美小礼物。圣诞树寻踪小屋运营时间为周一至周四: 12:00 - 19:00; 周五至周日及2023年12月25日、2024年1月1日: 11:00 - 19:00；惊喜兑奖小屋运营时间为周一至周四 12:00 - 20:00。\n\n3. 迪士尼小镇也将变身童话圣诞世界，2023年11月23日起至2024年1月1日，米奇和朋友们、达菲和朋友们将换上圣诞新造型。\n\n4. 「奇妙集章 点亮童话圣诞」惊喜上线，完成任务，点亮印章，叠加3枚印章图案，就能解锁圣诞好礼【4†source】。\n\n5. 「迪士尼冬日奇幻冰雪节」将在2023年11月10日至2024年2月24日举行，这一活动将呈现梦幻冬日奇境，让即将到来的冬日暖意融融、爱意满满【5†source】。\n\n以上活动具体内容和时间可能会根据实际安排有所变动，建议您提前关注上海迪士尼度假区官方网站或官方社交媒体获取最新信息。",
                        "annotations": [
                            {
                                "type": "web_citation",
                                "text": "【4†source】",
                                "start_index": 454,
                                "end_index": 464,
                                "web_citation": {
                                    "url": "https://www.weibo.com/5200478600/NtfYcexxE",
                                    "quote": "#上海迪士尼童话圣诞# 🎄2023年11月23日起至2024年1月1日，上海迪士尼圣诞季来袭！ 🌟米奇和朋友们、达菲和朋友们将换上圣诞新造型 💡达菲和玲娜贝儿将首次加入圣诞树点灯仪式！ 🎉迪士尼小镇也将变身童话圣诞世界 ️"
                                }
                            },
                            {
                                "type": "web_citation",
                                "text": "【5†source】",
                                "start_index": 540,
                                "end_index": 550,
                                "web_citation": {
                                    "url": "https://www.weibo.com/5200478600/NuVFF7Enb",
                                    "quote": "#上海迪士尼童话圣诞# 🎉2023年12月1日起至2024年1月1日，「奇妙集章 点亮童话圣诞」惊喜上线啦！ 完成任务，点亮印章，叠加3枚印章图案，就能解锁圣诞好礼，快来查收这份「圣诞集章攻略」吧～ 🎄童话圣诞正在进行中，等你前来集章，一起点亮属于你的节日祝愿吧～"
                                }
                            }
                        ]
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_5e65f2aae3a044f2b9c26012aaed1a95",
            "run_id": "run_ef18b2b8185744c3a58453dc1020f497",
            "metadata": null
        }
    ],
    "first_id": "msg_66a208c1de5d49ab88b664adef704364",
    "last_id": "msg_8f428bd21e4b403182ab0c5f6ea47781",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}

模型回复：

2024年上海迪士尼的圣诞节活动主要包括以下部分：

1. 米奇、达菲和朋友们一起点亮了魔法奇观环绕的圣诞大树，在烟花的照耀下，照亮了大家的童话梦想。这一活动会在上海迪士尼度假区圣诞季期间（2023年11月23日至2024年1月1日）举行。

2. 圣诞树寻踪活动，游客可以加入我们庆祝的行列，走入暖心的童话故事，漫步小镇街头，一起寻踪不一样的圣诞奇妙。完成圣诞树寻踪任务后，即可获得一份精美小礼物。圣诞树寻踪小屋运营时间为周一至周四: 12:00 - 19:00; 周五至周日及2023年12月25日、2024年1月1日: 11:00 - 19:00；惊喜兑奖小屋运营时间为周一至周四 12:00 - 20:00。

3. 迪士尼小镇也将变身童话圣诞世界，2023年11月23日起至2024年1月1日，米奇和朋友们、达菲和朋友们将换上圣诞新造型。

4. 「奇妙集章 点亮童话圣诞」惊喜上线，完成任务，点亮印章，叠加3枚印章图案，就能解锁圣诞好礼 [0]。

5. 「迪士尼冬日奇幻冰雪节」将在2023年11月10日至2024年2月24日举行，这一活动将呈现梦幻冬日奇境，让即将到来的冬日暖意融融、爱意满满 [1]。

以上活动具体内容和时间可能会根据实际安排有所变动，建议您提前关注上海迪士尼度假区官方网站或官方社交媒体获取最新信息。
[0] 网址：https://www.weibo.com/5200478600/NtfYcexxE #上海迪士尼童话圣诞# 🎄2023年11月23日起至2024年1月1日，上海迪士尼圣诞季来袭！ 🌟米奇和朋友们、达菲和朋友们将换上圣诞新造型 💡达菲和玲娜贝儿将首次加入圣诞树点灯仪式！ 🎉迪士尼小镇也将变身童话圣诞世界 ️
[1] 网址：https://www.weibo.com/5200478600/NuVFF7Enb #上海迪士尼童话圣诞# 🎉2023年12月1日起至2024年1月1日，「奇妙集章 点亮童话圣诞」惊喜上线啦！ 完成任务，点亮印章，叠加3枚印章图案，就能解锁圣诞好礼，快来查收这份「圣诞集章攻略」吧～ 🎄童话圣诞正在进行中，等你前来集章，一起点亮属于你的节日祝愿吧～
```

**Tab: Function calling示例**

**返回示例**
```json
assistant_id: asst_d7b4152fa8414d4e804733d35d445064
thread_id: thread_0a98d23117424cf088af4c32f99738ce
run_id: run_072537334b2f426eade029a77e6410c5
Status: in_progress
Status: requires_action
Process completed, exiting loop.
{'id': 'run_072537334b2f426eade029a77e6410c5', 'object': 'thread.run', 'created_at': 1703492444, 'assistant_id': 'asst_d7b4152fa8414d4e804733d35d445064', 'thread_id': 'thread_0a98d23117424cf088af4c32f99738ce', 'run_id': '', 'status': 'queued', 'started_at': 1703492444, 'expires_at': 0, 'cancelled_at': 0, 'failed_at': 0, 'completed_at': 0, 'last_error': None, 'model': 'abab5.5-chat', 'instructions': '根据用户定义的function，执行方案', 'tools': [{'type': 'function', 'function': {'name': 'get_weather', 'description': '获得天气信息', 'parameters': {'type': 'object', 'required': ['location'], 'properties': {'location': {'type': 'string', 'description': '获得天气的地点'}}}}}], 'file_ids': None, 'metadata': None, 'base_resp': {'status_code': 0, 'status_msg': 'success'}}
Status: in_progress
Status: completed
Process completed, exiting loop.
{
    "object": "list",
    "data": [
        {
            "id": "msg_fd788f009f7146a4835bedcee03c12ce",
            "object": "message",
            "created_at": 1703492444,
            "thread_id": "thread_0a98d23117424cf088af4c32f99738ce",
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "上海天气怎么样",
                        "annotations": []
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_d7b4152fa8414d4e804733d35d445064",
            "run_id": "run_072537334b2f426eade029a77e6410c5",
            "metadata": null
        },
        {
            "id": "msg_d6e3ef29ff7342a698cb9bc9c571edca",
            "object": "message",
            "created_at": 1703492448,
            "thread_id": "thread_0a98d23117424cf088af4c32f99738ce",
            "role": "ai",
            "content": [
                {
                    "type": "text",
                    "text": {
                        "value": "上海今天的天气是晴，最高气温7度，最低气温零下一度。",
                        "annotations": []
                    }
                }
            ],
            "file_ids": null,
            "assistant_id": "asst_d7b4152fa8414d4e804733d35d445064",
            "run_id": "run_072537334b2f426eade029a77e6410c5",
            "metadata": null
        }
    ],
    "first_id": "msg_fd788f009f7146a4835bedcee03c12ce",
    "last_id": "msg_d6e3ef29ff7342a698cb9bc9c571edca",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
```
#### 支持工具
Assistant可以访问MiniMax提供的工具，如代码解释器（Code Interpreter）、知识库检索（Retrieval）、函数调用（Function calling）。
使用以上工具需要支付一定的额外费用，具体的定价计费您可以在[产品定价](https://api.minimax.chat/document/price?id=6433f32294878d408fc8293e)进行详细的了解。
##### Code Interpreter（代码解释器）
Code Interpreter使得Assistants API可以在沙盒执行环境中编写和运行Python代码。该工具可以处理具有不同数据和格式的文件，并且可以生成图像文件。
利用Code Interpreter工具，Assistants可以生成并运行代码，以帮助用户解决具有难度的文件处理、数据分析或者复杂的数学逻辑问题。
###### 启用Code Interpreter
通过Assistants object数组中的tool参数里上传code_interpreter参数以启用Code Interpreter。
启用Code Interpreter之后模型会根据用户的请求内容判断何时调用工具。
**Tab: Python请求示例**

**启用Code Interpreter**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"

payload = json.dumps({
  "model": "abab5.5-chat",
  "name": "数据分析师",
  "description": "数据分析师，用在数据分析场景",
  "instructions": "是一个数据分析师，善于读取文件进行数据分析，并能严谨精准的给出分析结论",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  },
  "file_ids": [
    "88731293343809"
  ],
  "tools": [
    {
      "type": "code_interpreter"
    }
  ]
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell请求示例**

**启用Code Interpreter**
```shell
curl --location 'https://api.minimax.chat/v1/assistants/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--data '{
    "model": "abab5.5-chat",
    "name": "数据分析师",
    "description": "数据分析师，用在数据分析场景",
    "instructions": "是一个数据分析师，善于读取文件进行数据分析，并能严谨精准的给出分析结论",
    "metadata": {
    "key1": "value1",
    "key2": "value2"
    },
    "file_ids":["${file id}"],
        "tools": [
            {
                "type": "code_interpreter"
            }
        ]
    }
'
```
###### 上传文件给Code Interpreter
Code Interpreter可以分析文档file中的数据，此功能适用于当您希望向Assistants提供文档数据或者允许用户上传自己的文档进行分析的时候。
1. 上传文件到Assistants层级的文件可以被使用此assistant的所有run访问：使用[File](https://api.minimax.chat/document/guides/knowledge-base/File?id=6565fd39efee415861729ca2)接口上传文件（purpose = assistants），然后将file id作为assistant的一部分创建。
2. 文件也可以通过message进行上传，这些文件只能由特定thread访问，使用[FIle](https://api.minimax.chat/document/guides/knowledge-base/File?id=6565fd39efee415861729ca2)接口上传文件，然后将file id作为message的一部分创建。
文件最大支持512MB。Code Interpreter支持多种文件格式，包括.csv、.pdf、.json等。相关支持的文件（及其相应的MIME类型）的更多详细信息，可以跳转至“支持文件类型”部分得到详细信息。
**Tab: 上传到assistant层级（Python）**

**上传到assistant层级**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"

payload = json.dumps({
  "model": "abab5.5-chat",
  "name": "数据分析师",
  "description": "数据分析师",
  "instructions": "是一个数据分析师，善于读取文件进行数据分析，并能严谨精准的给出分析结论",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  },
  "file_ids": [
      # "此处添加file ids"
  ],
  "tools": [
    {
      "type": "code_interpreter"
    },
    {
      "type": "web_search"
    }
  ]
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: 上传到assistant层级（Shell）**

**上传到assistant层级**
```shell
curl --location 'https://api.minimax.chat/v1/assistants/create?GroupId=${GroupId}' \
--header 'Authorization: Bearer ${API_KEY}' \
--header 'Content-Type: application/json' \
--data '{
    "model": "my_model",
    "name": "function",
    "description": "代码机器人",
    "instructions": "Follow the instructions carefully",
    "metadata": {
        "key1": "value1",
        "key2": "value2"
    },
    "file_ids":["${file id}"],
    "tools": [        
        {"type":"code_interpreter"},
    ]
}'
```

**Tab: 上传到message层级（Python）**

**上传到message层级**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/messages/add?GroupId={GroupId}"

payload = json.dumps({
  "thread_id": "thread_123abc", # 替换为真实thread_id
  "role": "user",
  "file_ids": [
    # "此处添加file ids"
  ],
  "content": "统计一下总进球数？",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  }
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: 上传到message层级（Shell）**

**上传到message层级**
```shell
curl --location 'https://api.minimax.chat/v1/threads/messages/add?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "thread_abc123",
    "role": "user",
    "file_ids":["${file id}"],
    "content": "统计一下总进球数？",
    "metadata": {
    "key1": "value1",
    "key2": "value2"
    }
}'
```
###### Code Interpreter 的输入和输出日志
通过列出名为Code Interpreter的Run Step，您可以检查Code Interpreter的日志。
**Tab: Python示例**

**Python示例**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/run_steps/list?GroupId=$GroupId"

payload = json.dumps({
  "thread_id": "thread_123abc", # 替换为真实thread_id
  "run_id": "run_123abc"	# 替换为真实run_id
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell示例**

**Shell示例**
```shell
curl --location --request GET 'https://api.minimax.chat/v1/threads/run_steps/list?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "thread_abc123",
    "run_id": "run_abc123"
}'
```

**Tab: Run Step示例**

**RunStep示例**
```json
{
  "object": "list",
  "data": [
    {
      "id": "step_abc123",
      "object": "thread.run.step",
      "type": "tool_calls",
      "run_id": "run_abc123",
      "thread_id": "thread_abc123",
      "status": "completed",
      "step_details": {
        "type": "tool_calls",
        "tool_calls": [
          {
            "type": "code",
            "code": {
              "input": "# 计算 2 + 2\nresult = 2 + 2\nresult",
              "outputs": [
                {
                  "type": "logs",
                  "logs": "4"
                }
    ...
 }
```
##### 知识库检索（Retrieval）
知识库检索会利用外部的知识（例如专有产品信息或者用户提供的文档）来增强模型的专业知识。
一旦文件上传到知识库检索，模型将对您的文档进行切片，索引和向量化，检索相关内容以回答用户的请求。
*请注意，启用知识库检索后将同步开启相关计费项，请在查看详情。
###### 启用知识库检索
通过Assistant object数组中的tool参数里上传retireval参数以启用知识库检索。
**Tab: Python 示例**

**启用知识库示例**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"

payload = json.dumps({
  "model": "abab5.5-chat",
  "name": "小说专家",
  "description": "小说专家，用在小说解读场景中回答用户问题",
  "instructions": "是一个小说专家，读过万卷书，了解小说里的各种细致情节，另外也非常热心，会热情的帮读者解答小说里的问题",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  },
  "file_ids": [
    # "此处添加file ids"
  ],
  "tools": [
    {
      "type": "retrieval"
    }
  ]
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json',
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell 示例**

**启用知识库检索**
```shell
curl --location 'https://api.minimax.chat/v1/assistants/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--data '{
    "model": "abab5.5-chat",
    "name": "assistant名称",
    "description": "assistant设定描述",
    "instructions": "Follow the instructions carefully",
    "metadata": {
    "key1": "value1",
    "key2": "value2"
    },
    "file_ids":["${file id}"],
        "tools": [
            {
                "type": "retrieval"
            }
        ]
    }
'
```
###### 工作原理
当Assistant启用知识库检索工具并关联文档后，会先对文档进行向量化存储，在请求时Assistants api会自行判断是否对文档进行检索增强生成。
###### 上传文件进行检索
上传文件分为两种类型，分别可以通过Assistant级别或者Assistants File级别进行上传。
此外，您无需根据通过File接口上传的文件的大小付费，只须根据您附加到Assistants API中已向量化后存储的容量付费。
可以添加的文件大小最大为512MB。知识库检索支持多种文件格式，包括pdf、md、docx等。
相关支持的文件（及其相应的MIME类型）的更多详细信息，可以跳转至“支持文件类型”部分，得到详细信息。
以下为Assistant级别上传的示例代码
**Tab: Python 示例**

**知识库上传文件检索**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"

payload = json.dumps({
  "thread_id": "thread_123abc", # 此处替换为真实thread_id
  "role": "user",
  "file_ids": [
    # "此处添加file ids"
  ],
  "content": "在倚天屠龙记中，无俗念词的作者是谁？",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  }
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell 示例**

**知识库上传文件进行检索**
```shell
curl --location 'https://api.minimax.chat/v1/threads/messages/add?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "thread_abc123",
    "role": "user",
    "file_ids":["${file id}"],
    "content": "在倚天屠龙记中，无俗念词的作者是谁？",
    "metadata": {
    "key1": "value1",
    "key2": "value2"
    }
}'
```
###### 检索的引用信息
当在Message中出现file_citation的annotations时，即调用了retrieval，可以查看annotations字段来查看本次回复的文档引用信息。
**检索引用信息**
```json
{
  "id": "msg_83b789a095c74af5a192e1ce737c872d",
    "object": "message",
      "created_at": 1703416137,
        "thread_id": "thread_d98733725c7845eca48a56e14f6b1ce9",
          "role": "ai",
            "content": [
              {
                "type": "text",
                "text": {
                  "value": "在《倚天屠龙记》中，无俗念词的作者是南宋末年的一位武学名家，有道之士，他姓丘，名处机，道号长春子，名列全真七子之一，是全真教中出类拔萃的人物。【1†source】",
                  "annotations": [
                    {
                      "type": "file_citation",
                      "text": "【1†source】",
                      "start_index": 71,
                      "end_index": 81,
                      "file_citation": {
                        "file_id": "",
                        "quote": "###《倚天屠龙记》\n\n###第01章 天涯思君不可忘\n\n“春游浩荡，是年年寒食，梨花时节。白锦无纹香烂漫，玉树琼苞堆雪。静夜沉沉，浮光霭霭，冷浸溶溶月。人间天上，烂银霞照通彻。\n\n浑似姑射真人，天姿灵秀，意气殊高洁。万蕊参差谁信道，不与群芳同列。浩气清英，仙才卓荦，下土难分别。瑶台归去，洞天方看清绝。”\n\n作这一首无俗念词的，乃南宋末年一位武学名家，有道之士。此人姓丘，名处机，道号长春子，名列全真七子之一，是全真教中出类拔萃的人物。词品评论此词道：“长春，世之所谓仙人也，而词之清拔如此”。\n\n这首词诵的似是梨花，其实同中真意却是赞誉一位身穿白衣的美貌少女，说她“浑似姑射真人，天姿灵秀，意气殊高洁”，又说她“浩气清英，仙才卓荦”，“不与群芳同列”。词中所颂这美女，乃古墓派传人小龙女。她一生爱穿白衣，当真如风拂玉树，雪裹琼苞，兼之生性清冷，实当得起“冷浸溶溶月”的形容，以“无俗念”三字赠之，可说十分贴切。长春子丘处机和她在终南山上比邻而居，当年一见，便写下这首词来。\n\n这时丘处机逝世已久，小龙女也已嫁与神雕大侠杨过为妻。在河南少室山山道之上，却另有一个少女，正在低低念诵此词。"
                      }
                    }
                  ]
                }
              }
```
##### Web_search（网络搜索）
Assistants API支持调用MiniMax提供的web_search函数，以便获得实时搜索数据辅助大模型返回更好的结果。
**web_search搜索网络-python**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/run_steps/list?GroupId={GroupId}"

payload = json.dumps({
  "thread_id": "thread_123abc", # 此处替换为真实thread_id
  "run_id": "run_123abc"	# 此处替换为真实thread_id
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
  'Content-Type': 'application/json'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.json())

//
{
    "object": "list",
    "data": [
        {
            "id": "step_abc123",
            "object": "thread.run.step",
            "created_at": 1703322284,
            "run_id": "run_abc123",
            "assistant_id": "asst_abc123",
            "thread_id": "thread_abc123",
            "type": "tool_calls",
            "status": "completed",
            "step_details": {
                "type": "tool_calls",
                "tool_calls": [
                    {
                        "id": "call_123",
                        "type": "web_search",
                        "web_search": {
                            "name": "get_web_search_result",
                            "arguments": "{\"search_query\": \"2022年11月23日上海 天气\"}",
                            "output": "[{\"no\":1,\"content\":\"全球天气网（www.tianqi.com）提供上海11月份天气数据....]"
                        }
                    }
                ]
    ...
}
```
##### Function calling（函数调用）
与Chatcompletion API  以及Chatcompletion Pro API  类似，Assistants API支持函数调用。
Assistants API在调用函数时将在进行中的Run中进行暂停，您可以提供函数调用结果已继续执行Run。
###### 自定义函数
首先，您需要在创建assistant时定义函数：
**Tab: Python 示例**

**FC自定义函数**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/assistants/create?GroupId={GroupId}"

payload = json.dumps({
  "model": "abab5.5-chat",
  "name": "assistant名称",
  "description": "assistant设定描述",
  "instructions": "Follow the instructions carefully",
  "metadata": {
    "key1": "value1",
    "key2": "value2"
  },
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "get_weather",
        "description": "get weather",
        "parameters": {
          "type": "object",
          "required": [
            "city"
          ],
          "properties": {
            "city": {
              "type": "string"
            }
          }
        }
      }
    }
  ]
})
headers = {
  'Authorization': f'Bearer {API_KEY}',
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell 示例**

**FC自定义函数**
```shell
curl --location 'https://api.minimax.chat/v1/assistants/create?GroupId=$GroupId' \
--header 'Authorization: Bearer $API_KEY' \
--header 'Content-Type: application/json' \
--data '{
    "model": "abab5.5-chat",
    "name": "assistant名称",
    "description": "assistant设定描述",
    "instructions": "Follow the instructions carefully",
    "metadata": {
    "key1": "value1",
    "key2": "value2"
    },
        "tools": [
            {
                "type": "function",
                "function":{
                    "name":"get_weather",
                    "description":"get weather",
                    "parameters":{
                        "type":"object",
                        "required":["city"],
                        "properties":{"city":{"type":"string"}}
                    }
                }
            }
        ]
    }
'
```
###### 读取Assistant调用的函数
当assistant判断用户的Message需要调用函数时，包含该条Message的 Thread的运行将进入暂停状态。函数调用结果生成后，Run将进入下一状态，您可以通过requires_action来检索Run状态。
**伪代码**
```json
{
    "assistant_id": "asst_abc123",
    "id": "run_abc123",
    "object": "thread.run",
    "required_action": {
        "submit_tool_outputs": {
            "tool_calls": [
                {
                    "function": {
                        "arguments": '{"city": ' '"上海"}',
                        "name": "get_weather",
                    },
                    "id": "call_abc123",
                    "type": "function",
                }
            ]
        },
        "type": "submit_tool_outputs",
    },
    ...
}
```
###### 提交函数的输出内容
您可以通过提交所调用的函数返回的submit_tool_outputs来完成运行。
您可以使用tool_calls_id通过require_action匹配每个函数调用的输出输出结果。
**Tab: Python 示例**

**python示例**
```python
import requests
import json

GroupId = "请填写你的group id"
API_KEY = "请填写你的api key"

url = f"https://api.minimax.chat/v1/threads/run/submit_tool_outputs?GroupId={Groupid}"

payload = json.dumps({
    "thread_id": "thread_abc123", # 此处替换为真实thread_id
    "run_id": "run_abc123", # 此处替换为真实run_id
    "tool_outputs": [
        {
            "tool_call_id": "call_abc123", # 此处替换为真实tool_call_id
            "output": "内容"
        }
    ]
})
headers = {
    'Authorization': f'Bearer {API KEY}',
    'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

**Tab: Shell 示例**

**Shell标题**
```shell
curl --location 'https://api.minimax.chat/v1/threads/run/submit_tool_outputs?GroupId= ${Groupid}' \
--header 'Authorization: Bearer ${API KEY}' \
--header 'Content-Type: application/json' \
--data '{
    "thread_id": "thread_abc123",
    "run_id": "run_abc123",
    "tool_outputs": [
        {
            "tool_call_id": "call_abc123",
            "output": "内容"
        }
    ]
}'
```
###### 支持文件类型
对于text/MIME类型，编码必须是utf-8、utf-16或ascii之一
| 文件格式 | MIME类型 | Code Interpreter是否支持 | Retrieval是否支持 |
| --- | --- | --- | --- |
| .c | text/x-c | ✅ | ✅ |
| .cpp | text/x-c++ | ✅ | ✅ |
| .csv | application/csv | ✅ | ✅ |
| .docx | application/vnd.openxmlformats-officedocument.wordprocessingml.document | ✅ | ✅ |
| .html | text/html | ✅ | ✅ |
| .java | text/x-java | ✅ | ✅ |
| .json | application/json | ✅ | ✅ |
| .md | text/markdown | ✅ | ✅ |
| .pdf | application/pdf | ✅ | ✅ |
| .php | text/x-php | ✅ | ✅ |
| .pptx | application/vnd.openxmlformats-officedocument.presentationml.presentation | ✅ | ✅ |
| .py | text/x-python | ✅ | ✅ |
| .py | text/x-script.python | ✅ | ✅ |
| .rb | text/x-ruby | ✅ | ✅ |
| .tex | text/x-tex | ✅ | ✅ |
| .txt | text/plain | ✅ | ✅ |
| .css | text/css | ✅ |  |
| .jpeg | image/jpeg | ✅ |  |
| .jpg | image/jpeg | ✅ |  |
| .js | text/javascript | ✅ |  |
| .gif | image/gif | ✅ |  |
| .png | image/png | ✅ |  |
| .tar | application/x-tar | ✅ |  |
| .ts | application/typescript | ✅ |  |
| .xlsx | application/vnd.openxmlformats-officedocument.spreadsheetml.sheet | ✅ |  |
| .xml | application/xml or "text/xml" | ✅ |  |
| .zip | application/zip | ✅ |  |

## 拍照识图互动学习
- 原始链接: https://platform.minimaxi.com/document/learning?key=68ba8cdb95a600dbfc91f942
- 文档ID: `68ba8cdb95a600dbfc91f942`
# 拍照识图互动学习

在面向儿童的早教与启蒙学习场景中，利用 **MiniMax 文本-图像理解能力** 与 ​**MiniMax-Speech-02 语音合成技术**​，让孩子们“拍一拍，就能学”。


## 解决方案

##### 拍照识图
-  相机拍摄任意物体，调用 文本大模型的图像理解能力，快速识别物体类别、特征与背景知识。

<br>

##### 知识讲解与外语学习
- 模型将识别结果转化为适合儿童理解的简短知识卡片。
- 同时调用 **MiniMax-Speech-02** 超拟人 TTS，用温暖亲切的声音进行讲解。

<br>

##### 双语/多语交互
- 支持英语、日语等多语言输出，帮助孩子边玩边学，提升外语听力与口语兴趣。

<br>

##### 趣味问答
- AI语音助手可根据识别的物品发起互动问答，例如“你知道长颈鹿的脖子为什么那么长吗？”

<br>

## 业务价值

<container gap="4" columns="2">

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 沉浸式学习
视觉识别 + 语音互动，让学习体验更直观、更有趣

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 多语言启蒙
自然融入双语学习场景，帮助孩子早期掌握外语词汇与发音

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 自适应内容
根据不同年龄段，自动调整讲解的难度与语言长度

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 延展性强
除日常物品外，还可识别动植物、交通工具、科学器材等多类事物

</card>


</container>

<br>


## 核心API能力
1. 文本合成接口：输入图片，进行识别
<request method="POST" url="https://api.minimaxi.com/v1/text/chatcompletion_v2"></request>

1. 语音合成接口，将识别讲解结果，转化成语音
<request method="POST" url="https://api.minimaxi.com/v1/t2a_v2"></request>

## 使用示例
图像理解与语音合成

```python
import base64
import requests
import json

group_id = "请填写您的group_id"
api_key = "请填写您的_api_key"

# 1. 上传图片并识别
def recognize_image(image_path):
    with open(image_path, "rb") as img_file:
        img_base64 = base64.b64encode(img_file.read()).decode("utf-8")

    url = f"https://api.minimaxi.com/v1/text/chatcompletion_v2?GroupId={group_id}"
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    payload = {
        "model": "MiniMax-Text-01",
        "messages": [
            {
                "role": "system",
                "content": (
                    "你是一个专门为小朋友讲解的智能老师。\n"
                    "你的任务是：识别图片中的物体，并用简单有趣的方式做解释。\n"
                    "回答必须同时包含中文和英文。\n"
                    "中文要温柔、亲切；英文要简短、易懂。\n"
                    "比如：'这是一个苹果，它是一种甜甜的水果。'\n"
                    "英文：'This is an apple. It is a sweet fruit.'"
                )
            },
            {
                "role": "user",
                "content": f"请识别并讲解这个图片: [图片base64:{img_base64}]"
            }
        ]
    }

    response = requests.post(url, headers=headers, data=json.dumps(payload))
    result = response.json()
    text_output = result["choices"][0]["message"]["content"]
    return text_output

# 2. 将识别讲解结果转成语音
def text_to_speech(text, output_file="output.mp3"):
    url = f"https://api.minimaxi.com/v1/t2a_v2?GroupId={group_id}"
    payload = json.dumps({
        "model": "speech-2.5-hd-preview",
        "text": text,
        "stream": False,
        "voice_setting": {
            "voice_id": "male-qn-qingse",
            "speed": 1,
            "vol": 1,
            "pitch": 0,
            "emotion": "happy"
        },
        "audio_setting": {
            "sample_rate": 32000,
            "bitrate": 128000,
            "format": "mp3",
            "channel": 1
        }
    })
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    response = requests.post(url, headers=headers, data=payload)
    parsed_json = response.json()
    audio_value = bytes.fromhex(parsed_json["data"]["audio"])
    with open(output_file, "wb") as f:
        f.write(audio_value)
    print(f"语音已保存为 {output_file}")

# ==== Demo运行 ====
if __name__ == "__main__":
    explain_text = recognize_image("test.jpg")  # 输入图片
    print("识别讲解结果：\n", explain_text)

    text_to_speech(explain_text, "output.mp3")

```

## 模型
- 原始链接: https://platform.minimaxi.com/document/models?key=6897027ae92177d0a09a460e
- 文档ID: `6897027ae92177d0a09a460e`
### 模型发布
###### 2025年9月11日
- **Music-1.5** — 最新音乐模型发布，支持4分钟音乐时长、回归“好听”本质 — [立即体验](https://platform.minimaxi.com/document/music_generation?key=68ac02e16602726333ffd430)
###### 2025年8月6日
- **Speech 2.5** — 最新一代语音生成模型，支持更多语种，具备极高相似度声音表现 — [了解更多](https://www.minimaxi.com/news/minimax-speech-25)
###### 2025年6月20日
- **Music-1.5（Beta)** — 新一代音乐生成模型Music-1.5（Beta)发布，支持输入音乐灵感和歌词进行音乐生成 — [即刻体验](https://platform.minimaxi.com/document/music_generation?key=68ac02e16602726333ffd430)
###### 2025年6月18日
- **MiniMax Hailuo 02** — 新一代视频生成模型MiniMax Hailuo 02正式发布，支持1080P分辨率以及10s视频生成 — [了解更多](https://www.minimaxi.com/news/minimax-hailuo-02)
###### 2025年6月16日
- **MiniMax-M1** — 推理模型MiniMax-M1正式发布，全球领先，80K思维链 x 1M输入，效果比肩海外顶尖模型 — [了解更多](https://www.minimaxi.com/news/minimaxm1)
###### 2025年4月2日
- **Image-01** — 图像生成模型Image-01模型发布，支持文本描述生成多种尺寸的图片 — [即刻体验](https://platform.minimaxi.com/document/text_to_image?key=68ac01e26fe587e3fbfe5765)
###### 2025年2月11日
- **T2V-01-Director / I2V-01-Director** — 导演级视频生成模型正式发布，对运镜描述指令有更好遵循，电影级镜头叙事语言 — [即刻体验](https://platform.minimaxi.com/document/image_to_video?key=68abe046d08627aad9674c07)
###### 2025年1月15日
- **MiniMax-Text-01 / MiniMax-VL-01** — 全新一代文本模型MiniMax-Text-01和视觉理解模型MiniMax-VL-01正式发布

## 模型介绍
- 原始链接: https://platform.minimaxi.com/document/models_intro?key=68aec7c34c75b9c918ccd104
- 文档ID: `68aec7c34c75b9c918ccd104`
### 模型介绍
#### 旗舰模型
- **MiniMax Speech 2.5** — 更出色的多语种表现力，超强的音色相似度 — [https://www.minimaxi.com/news/minimax-speech-25](https://www.minimaxi.com/news/minimax-speech-25) — ![image](https://filecdn.minimax.chat/public/96bcb81c-ab1b-4ec5-a5c5-36caf3572911.png)
- **MiniMax-Hailuo-02** — 首尾帧视频生成，SOTA指令遵循，极致物理表现 — [https://www.minimaxi.com/zh/news/minimax-hailuo-02](https://www.minimaxi.com/zh/news/minimax-hailuo-02) — ![image](https://filecdn.minimax.chat/public/d2ed4700-da5f-451a-8dc3-ce4f0a46a8b0.png)
- **MiniMax M1** — 80K思维链 x 1M输入，效果比肩海外顶尖模型 — [https://www.minimaxi.com/news/minimaxm1](https://www.minimaxi.com/news/minimaxm1) — ![image](https://filecdn.minimax.chat/public/ccfd6515-2f83-4c7b-b21d-e437676928c8.png)
#### 模型概览
| 模型名称 | 介绍 |
| --- | --- |
| MiniMax-M1 | 全球领先：80K思维链 x 1M输入，效果比肩海外顶尖模型 |
| MiniMax-Text-01 | 全新架构文本模型，支持1M超长上下文 |
| 模型名称 | 介绍 |
| --- | --- |
| Speech-02-HD | 最新语音HD模型，拥有出色的韵律和稳定性，复刻相似度和音质表现突出 |
| Speech-02-Turbo | 最新语音Turbo模型，小语种能力增强，性能表现出色 |
| 模型名称 | 介绍 |
| --- | --- |
| MiniMax Hailuo 02 | 新一代视频生成模型，1080p原生，SOTA指令遵循，极致物理表现 |
| T2V / I2V-01-Director | 文生视频/图生视频（导演版）模型，对于运镜指令有更好的遵循，画面表现更加生动 |
| I2V-01-live | 图生视频模型，对卡通、漫画等手绘风格增强 |
| S2V-01 | 主体参考视频生成模型，支持用户上传主体图片，生成视频保持主体遵循稳定性 |
| 模型名称 | 介绍 |
| --- | --- |
| image-01 | 图像生成模型，画面表现细腻，支持文生图、图生图 |
| image-01-live | 图像生成模型，手绘、卡通等画风增强，支持文生图并进行画风设置 |
| 模型名称 | 介绍 |
| --- | --- |
| music-1.5 | 音乐生成模型，支持用户提供音乐描述和歌词，生成AI音乐 |

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/music_api_intro?key=68ac02b86fe587e3fbfe5a11
- 文档ID: `68ac02b86fe587e3fbfe5a11`
# 音乐生成（Music Generation）

本接口根据音乐描述（`prompt`）和歌词（`lyrics`），生成一首人声的歌曲。

## 支持模型

| 模型名称 | 使用方法 |
| :--- | :--- |
| music-1.5 | 最新音乐生成模型，支持用户输入音乐灵感和歌词，生成AI音乐 |

## 官方MCP

欢迎访问官方MCP，了解更多能力。
github链接：[https://github.com/MiniMax-AI/MiniMax-MCP
](https://github.com/MiniMax-AI/MiniMax-MCP)

## 音乐生成
- 原始链接: https://platform.minimaxi.com/document/music_generation?key=68ac02e16602726333ffd430
- 文档ID: `68ac02e16602726333ffd430`
# 音乐生成（Music Generation）

<request method="POST" url="https://api.minimaxi.com/v1/music_generation"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>Required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>Required</red>
使用的模型名称， `music-1.5`
<hr/>

#### prompt    <gray>string</gray> <red>Required</red>
音乐的描述，用于指定风格、情绪和场景。例如“流行音乐, 难过, 适合在下雨的晚上”。长度限制为【10，300】个字符
<hr/>

#### lyrics    <gray>string</gray> <red>Required</red>
歌曲的歌词。使用 \n 分隔每行。你可以在歌词中加入 `[Intro]`, `[Verse]`, `[Chorus]`, `[Bridge]`, `[Outro]` 等结构标签来优化生成的音乐结构。长度限制为【10，3000】个字符
<hr/>

#### stream    <gray>bool</gray>
是否使用流式传输，默认为 `false`
<hr/>

#### output_format    <gray>string</gray>
音频的返回格式，可选值为 url 或 hex，默认为 hex。当 stream 为 true 时，仅支持 hex 格式。注意：url 的有效期为 24 小时，请及时下载
<hr/>

#### audio_setting    <gray>object</gray>
音频输出配置
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### sample_rate    <gray>integer</gray>
采样率。可选值：`16000`, `24000`, `32000`, `44100`
<hr/>

#### bitrate    <gray>integer</gray>
比特率。可选值：`32000`, `64000`, `128000`, `256000`
<hr/>

#### format    <gray>string</gray>
音频编码格式。可选值：`mp3`, `wav`, `pcm`

</card>
</details>
<hr/>

#### aigc_watermark    <gray>boolean</gray>
是否在音频末尾添加水印，默认为 `false`。仅在非流式 (`stream: false`) 请求时生效

<br>

### 返回体
<hr/>

#### data    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status    <gray>int</gray>
音乐合成状态：
1: 合成中
2: 已完成

<hr/>

#### audio    <gray>string</gray>
当 `output_format` 为 `hex` 时返回，是音频文件的 16 进制编码字符串

</card>
</details>
<hr/>

#### base_resp    <gray>object</gray>
状态码及详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0，请求成功

1002，触发限流，请稍后再试

1004，账号鉴权失败，请检查 API-Key 是否填写正确

1008，账号余额不足

1026，图片描述涉及敏感内容

2013，传入参数异常，请检查入参是否按要求填写

2049，无效的api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>

</left>

<right>
<codetabs title="请求示例">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/music_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "music-1.5",
    "prompt": "独立民谣,忧郁,内省,渴望,独自漫步,咖啡馆",
    "lyrics": "[verse]\n街灯微亮晚风轻抚\n影子拉长独自漫步\n旧外套裹着深深忧郁\n不知去向渴望何处\n[chorus]\n推开木门香气弥漫\n熟悉的角落陌生人看",
    "audio_setting": {"sample_rate": 44100, "bitrate": 256000, "format": "mp3"},
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
audio_hex = response.json()["data"]["audio"]

with open("output.mp3", "wb") as f:
    f.write(bytes.fromhex(audio_hex))

</codetab>
<codetab  language="curl">
curl -X POST "https://api.minimaxi.com/v1/music_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "model": "music-1.5",
    "prompt": "独立民谣,忧郁,内省,渴望,独自漫步,咖啡馆",
    "lyrics": "[verse]\n街灯微亮晚风轻抚\n影子拉长独自漫步\n旧外套裹着深深忧郁\n不知去向渴望何处\n[chorus]\n推开木门香气弥漫\n熟悉的角落陌生人看",
    "audio_setting": {
      "sample_rate": 44100,
      "bitrate": 256000,
      "format": "mp3"
    }
  }'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab  language="json">
{
  "data": {
    "audio": "hex编码的音频数据",
    "status": 2
  },
  "trace_id": "04ede0ab069fb1ba8be5156a24b1e081",
  "extra_info": {
    "music_duration": 25364,
    "music_sample_rate": 44100,
    "music_channel": 2,
    "bitrate": 256000,
    "music_size": 813651
  },
  "analysis_info": null,
  "base_resp": {
    "status_code": 0,
    "status_msg": "success"
  }
}

</codetab>
</codetabs>

</right>

</container>

## 音乐模型创意玩法
- 原始链接: https://platform.minimaxi.com/document/music_generation_solution?key=68ba7e44f4400933d0501f8e
- 文档ID: `68ba7e44f4400933d0501f8e`
# 音乐生成创意玩法

**MiniMax 音乐生成模型在咪咕彩铃和小红书玩法**

基于MiniMax 音乐和语音模型，实现在咪咕及小红书平台，进行音乐创意玩法。


## 解决方案

### 咪咕彩铃玩法

##### 彩铃个性化定制
1. 用户可通过咪咕彩铃平台，上传个人干声（`voice_id`）、伴奏（`instrumental_id`）及歌词（`lyrics`），调用 **Music Generation ​API** 生成个性化彩铃。
2. 支持 **TTS、声音唱歌模式、​文生音色唱歌模式、清唱模式**等，生成后的成品可直接设置为个人彩铃。

**版权提醒：干声、伴奏及歌词需用户自行确保版权合规**

<br>

##### AI热门歌曲歌单生成
1. 平台可根据热门曲风标签（如流行、R&B、摇滚、古风）自动生成 AI 翻唱歌曲，用于彩铃推荐。
2. 通过 文生音色接口 配合流行度排行数据，批量合成多个不同音色版本，满足用户的差异化需求。
<br>

##### 互动玩法：彩铃+BGM
1. 支持在咪咕App内录制语音并加BGM，形成“语音铃声”玩法，例如节日祝福、情侣对话等。
2. 可调用 BGM模式（仅伴奏）+TTS 生成，带来沉浸式铃声体验。
<br>
<br>

### 小红书玩法

##### 音乐笔记制作
1. 用户可在小红书直接输入歌词+选择曲风，调用 Music Generation API 生成歌曲音频，再配合图片或短视频发布音乐笔记。
2. 曲风支持16种，可切换试听，系统将自动渲染成视频封面+ BGM 成品。


<br>

##### 唱聊玩法
1. 在小红书的群聊功能中，输入歌词→选择曲风→实时生成歌曲并发送至群聊。
2. 支持 AI 生成封面图（结合文生图API）+音乐同步生成，方便直接转发或发布。

<br>

##### 声聊+音乐配音
1. 用户输入文字，AI按选择的音色（voice_id）生成配音，同时可叠加BGM。
2. 适合剧情类短视频、Vlog解说、情绪类分享等。

<br>



 

## 使用示例
1. TTS声音歌唱模式

```python
{
  "voice_id": "xxxx",
  "instrumental_id": "yyyy",
  "lyrics": "这里是一段歌词文本",
  "output_format": "mp3"
}

```

2. BGM 模式（无干声）
```python
{
  "reference_instrumental": "zzzz",
  "output_format": "mp3"
}
```

3.文生音色+合成
```python
{
  "voice_gender": "female",
  "voice_age": "young",
  "voice_description": "甜美、清亮",
  "instrumental_id": "yyyy",
  "lyrics": "这里是一段歌词文本"
}
```

## 平台介绍
- 原始链接: https://platform.minimaxi.com/document/platform_intro?key=68aec6cd7189e62d342a3d7b
- 文档ID: `68aec6cd7189e62d342a3d7b`
# MiniMax开放平台介绍

MiniMax是国内领先的拥有文本、声音、视觉三模态融合的千亿参数语言大模型，并打通产品全链路的创业公司。通过搭建超大规模实验平台与推理平台，跑通技术与产品的迭代闭环，实现模型能力与用户交互量的极速增长。

模型层面，MiniMax下设的推理平台拥有万卡级别的跨云高效计算系统，可支持单日过亿次级的调用。MiniMax通用模型具有突破极限的模型迭代速度，全面覆盖文本、图像、声音等多种模型。MiniMax人工智能引擎的训练平台拥有极高效的千卡并行算力，稳定性已得到验证。

应用层面，MiniMax开放平台提供大模型应用的企业级解决方案，为企业和开发者提供功能丰富、高易用性的API，最大限度降低使用者的开发复杂性，帮助使用者快速在目标场景中验证价值并生产部署。

## 模型与接口

MiniMax开放平台提供标准的API接口，赋能开发者完成所在行业的AI场景创新。

- **文本大模型：支持Chat Completion v2接口**，提供基于自然语言交互的文本生成能力。  
- **语音大模型，支持T2A v2、T2A Async v2、Voice Cloning、Voice Design接口**，提供自然语言交互生成能力的文字转语音能力。
- **视频大模型，支持Video Generation、Video Template Generation接口**， 提供用户通过文本描述及参考图片，以及视频模板，进行视频生成能力。  
- **图像大模型，支持Image Generation接口**，提供用户通过文本描述进行图片生成能力。  
- **音乐大模型，支持Music Generation接口**，提供用户通过歌曲特征与歌词，进行音乐生成能力。  



## 场景与行业

**MiniMax文本模型**，支持包括不限于：  
关键信息提取、知识问答、文本创作、文本总结、内容扩写、内容缩写、内容仿写、人设对话、作业批改、情感分析、逻辑推理、数学能力、代码能力等能力点，通过以上能力点进行组合，覆盖 **社交、语聊、游戏、招聘、教育、客服** 等多种场景。

**MiniMax语音模型**，支持包括不限于：  
文本匹配的声音、声音快速复刻、声音精品复刻的能力，并支持多语言、多性格、多场景的音色能力，在 **社交、播客、有声书、新闻资讯、教育、数字人** 等多元场景中均展现出强大的实力。

**MiniMax视频大模型**，支持包括不限于：  
基于文本、图片生成视频能力，实现创意视频生成，提供保持主体人物一致性能力。提供影视制作、广告创意、社交娱乐、个性化内容定制等能力点，覆盖 **娱乐、传媒等** 场景，驱动内容生产革新。

## 优势与服务

- **模型性能领先：** 千亿参数语言大模型、万卡级算力池、支撑单日过亿次调用、数千卡以上常态化训练，全链路自研训练框架，并进行周级快速迭代。  
- **敏捷易用：** 接口提供多元的参数和使用方式并提供大量应用示例。  
- **高并发吞吐：** 超大的推理集群，支持将模型应用于大规模用户产品中。  
- **安全合规：** 双重保障，大语言模型安全能力+第三方独立审核接口，保障输出结果的安全合规。  
- **专家团队：** 顶尖的研发和商务团队，提供行业先进的AGI技术服务和解决方案。

## 产品定价
- 原始链接: https://platform.minimaxi.com/document/pricing?key=68aec7e84c75b9c918ccd10d
- 文档ID: `68aec7e84c75b9c918ccd10d`
### 产品定价
MiniMax开放平台提供基于大语言模型的丰富API能力，以满足不同企业客户及专业开发者的需求。计价方式灵活多样，满足不同业务规模。
#### 资源包
##### 视频资源包
[立即购买](https://platform.minimaxi.com/user-center/payment/subscription)
| 类型 | 视频基础包<br>（节省5%） | 视频高级包<br>（节省10%） | 视频进阶包<br>（节省15%） | 视频企业包<br>（节省20%） | 商务定制 |
| --- | --- | --- | --- | --- | --- |
| 折扣价 | ¥7,000 | ¥15,000 | ¥30,000 | ¥40,000 | —— |
| 原价 | ¥7,368 | ¥16,667 | ¥35,294 | ¥50,000 |  |
| 有效期 | 1个月 | 1个月 | 1个月 | 1个月 | —— |
| 积分总量 | 3,680 | 8,330 | 17,650 | 25,000 | —— |
| 功能 | 支持 video generation 接口<br>RPM：20<br>支持所有模型，包含01和02系列 | 支持 video generation 接口<br>RPM：30<br>支持所有模型，包含01和02系列 | 支持 video generation 接口<br>RPM：40<br>支持所有模型，包含01和02系列 | 支持 video generation 接口<br>RPM：50<br>支持所有模型，包含01和02系列 | 无限 RPM/TPM<br>模型更新优先体验<br>专属的安全和稳定性保障 |
> 1.不同模型生成视频的扣减次数不同
> • [MiniMax-Hailuo-02]，生成单条512p，6s视频，扣减0.3积分
> • [MiniMax-Hailuo-02]，生成单条512p，10s视频，扣减0.5积分
> • [MiniMax-Hailuo-02]，生成单条768p，6s视频，扣减1积分
> • [MiniMax-Hailuo-02]，生成单条768p，10s视频，扣减2积分
> • [MiniMax-Hailuo-02]，生成单条1080p，6s视频，扣减2积分
> • [T2V-01] [I2V-01] [I2V-01-live] [T2V-01-Director] [I2V-01-Director], 生成单条视频，则从资源包里扣减1积分
> • [S2V-01], 生成单条视频，则从资源包里扣减1.5积分
> 2.生成失败或命中安全审核的视频不会扣减积分数
> 3.资源包余量不继承，过期会自动清零
##### 语音资源包
###### Turbo系列
[立即购买](https://platform.minimaxi.com/user-center/payment/subscription)
| 分类 | 语音套餐一<br>（节省10%） | 语音套餐二<br>（节省15%） | 语音套餐二<br>（节省20%） | 商务定制 |
| --- | --- | --- | --- | --- |
| 折扣价 | ¥360 | ¥3,400 | ¥32,000 | 基于量级的额外折扣 |
| 原价 | ¥400 | ¥4,000 | ¥40,000 | —— |
| 有效期 | 1个月 | 3个月 | 1年 | —— |
| 总字符数量 | 2,000,000字符 | 20,000,000字符 | 200,000,000字符 | —— |
| 功能 | 支持T2A v2 / T2A async v2接口<br>RPM：60<br>支持Turbo系列模型<br>赠送10个快速克隆音色 | 支持T2A v2 / T2A async v2接口<br>RPM：200<br>支持Turbo系列模型<br>赠送30个快速克隆音色 | 支持T2A v2 / T2A async v2接口<br>RPM：500<br>支持Turbo系列模型<br>赠送300个快速克隆音色 | 无限 RPM/TPM<br>模型更新优先体验<br>专属的安全和稳定性保障 |
###### HD系列
[立即购买](https://platform.minimaxi.com/user-center/payment/subscription)
| 分类 | 语音套餐一<br>（节省10%） | 语音套餐二<br>（节省15%） | 语音套餐三<br>（节省20%） | 商务定制 |
| --- | --- | --- | --- | --- |
| 折扣价 | ¥630 | ¥5,950 | ¥56,000 | 基于量级的额外折扣 |
| 原价 | ¥700 | ¥7,000 | ¥70,000 | —— |
| 有效期 | 1个月 | 3个月 | 1年 | —— |
| 总字符数 | 2,000,000字符 | 20,000,000字符 | 200,000,000字符 | —— |
| 功能 | 支持T2A v2 / T2A async v2接口<br>RPM：60<br>支持HD系列模型<br>赠送10个快速克隆音色 | 支持T2A v2 / T2A async v2接口<br>RPM：200<br>支持HD系列模型<br>赠送30个快速克隆音色 | 支持T2A v2 / T2A async v2接口<br>RPM：500<br>支持HD系列模型<br>赠送300个快速克隆音色 | 无限 RPM/TPM<br>模型更新优先体验<br>专属的安全和稳定性保障 |
#### 文本
基于MiniMax端到端自研多模态大语言模型，为企业客户和个人开发者提供功能丰富的API能力，适用多数文本处理场景。
| 计费项 | 模型 | 单价 | 计费说明 | tokens与字符比（估算） |
| --- | --- | --- | --- | --- |
| 文本生成<br>ChatCompletion | MiniMax-M1 | 1. 输入长度0-32k：<br>输入：0.8元/百万tokens<br>输出：8元/百万tokens<br>2. 输入长度32-128k：<br>输入：1.2元/百万tokens<br>输出：16元/百万tokens<br>3. 输入长度128k+：<br>输入：2.4元/百万tokens<br>输出：24元/百万tokens | 计费token数包括输入和输出，不同的输入长度将触发不同的输入以及输出的计费档位。 | 1000 tokens : 1600中文字符 |
| 文本生成<br>ChatCompletion | MiniMax-Text-01 | 输入：1元/百万tokens<br>输出：8元/百万tokens | 计费token数包括输入和输出，当使用图片输入时，会计算图片的token，具体消耗视图片分辨率而定，二者强相关。传入的上文所有图片url都会被使用和计费。 | 1000 tokens : 1600中文字符 |
> 注：计费项是token数；tokens字符比值根据使用场景的不同略有浮动，以实际消耗为准，字符数包括标点等。
#### 语音
MiniMax语音大模型能够根据上下文，智能预测文本的情绪、语调等信息，并生成超自然、高保真、个性化的语音。在社交、播客、有声书、新闻资讯、教育、数字人等10+场景中展现出强大的实力。
| 计费项 | 模型 | 接口说明 | 单价 |
| --- | --- | --- | --- |
| 同步语音合成<br>T2A | speech-2.5-hd-preview<br>speech-02-hd<br>speech-01-hd | 支持音量、语调、语速调整和混音功能，支持比特率、采样率相关参数调整特性，支持音频时长、音频大小等返回参数，适用于需要短文本快速得到结果的场景，比如闲聊、对话等场景 | 3.5元/万字符 |
| 同步语音合成<br>T2A | speech-2.5-turbo-preview<br>speech-02-turbo<br>speech-01-turbo | 支持音量、语调、语速调整和混音功能，支持比特率、采样率相关参数调整特性，支持音频时长、音频大小等返回参数，适用于需要短文本快速得到结果的场景，比如闲聊、对话等场景 | 2元/万字符 |
| 异步长文本语音合成<br>T2A Async | speech-2.5-hd-preview<br>speech-02-hd<br>speech-01-hd | 支持基于文本到语音的异步生成，单次文本生成传输最大支持100万字符，生成的完整音频结果支持异步的方式进行检索。 | 3.5元/万字符 |
| 异步长文本语音合成<br>T2A Async | speech-2.5-turbo-preview<br>speech-02-turbo<br>speech-01-turbo | 支持基于文本到语音的异步生成，单次文本生成传输最大支持100万字符，生成的完整音频结果支持异步的方式进行检索。 | 2元/万字符 |
| 音色设计<br>Voice Design | 所有模型 | 支持基于用户输入的声音描述prompt，来生成音色（voice_id）；并支持使用该生成的音色（voice_id）在同步语音合成、异步长文本语音合成接口中进行语音合成。 | 9.9元/音色。调用本接口获得新设计的音色时，不会立即收取音色设计费用。音色生成费用将在首次使用此音色进行语音合成时收取。<br>本接口内的试听语音合成会收取2元/万字符的费用。 |
| 快速复刻<br>Voice Cloning | 所有模型 | 基于大语言模型的音色克隆更加精准快速，无需数小时时长的超高质量原音频、无需传统TTS的超长工期，可以在极短时间内完成音色复刻，并通过大语言模型加持，使复刻后的音色与原音色进行高质量还原，从而满足客户需求。 | 9.9元/音色。调用本接口获得复刻音色时，不会立即收取音色复刻费用。音色的复刻费用将在首次使用此复刻音色进行语音合成时收取。<br>试听字符根据选择的试听模型收费。 |
> 注：计费项是字符数，以10000个字符（输入）为单位，1个汉字算2个字符，英文字母、希腊字母、标点符号、特殊符号、空格、回车等算1个字符。
#### 视频
| 计费项 | 模型 | 接口说明 | 单价 |
| --- | --- | --- | --- |
| 视频生成<br>Video Generation | MiniMax-Hailuo-02 | 支持根据用户输入的文字描述和图片，生成512P 6s视频 | 0.6元/视频 |
| 视频生成<br>Video Generation | MiniMax-Hailuo-02 | 支持根据用户输入的文字描述和图片，生成512P 10s视频 | 1元/视频 |
| 视频生成<br>Video Generation<br>768P 6s | MiniMax-Hailuo-02 | 支持用户输入的文字描述和图片，生成768P 6s视频 | 2元/视频 |
| 视频生成<br>Video Generation<br>768P 10s | MiniMax-Hailuo-02 | 支持用户输入的文字描述和图片，生成768P 10s视频 | 4元/视频 |
| 视频生成<br>Video Generation<br>1080P 6s | MiniMax-Hailuo-02 | 支持用户输入的文字描述和图片，生成1080P 6s视频 | 3.5元/视频 |
| 视频生成<br>Video Generation | T2V-01-Director<br>I2V-01-Director<br>I2V-01-live<br>T2V-01<br>I2V-01 | 支持根据用户输入的文字描述和图片，生成视频。 | 3元/视频 |
| 视频生成<br>Video Generation | S2V-01 | 支持根据用户输入的文字描述和主体参考图片，生成视频 | 4.5元/视频 |
#### 视频Agent
| 视频Agent模板ID | 模版名称 | 单价 |
| --- | --- | --- |
| 392753057216684038 | 跳水 | 3元/视频 |
| 393881433990066176 | 吊环 | 3元/视频 |
| 393769180141805569 | 绝地求生 | 3.5元/视频 |
| 394246956137422856 | 万物皆可labubu | 3.5元/视频 |
| 393866076583718914 | 女模特试穿广告 | 7元/视频 |
| 393876118804459526 | 男模特试穿广告 | 7元/视频 |
| 393879757702918151 | 麦当劳宠物外卖员 | 8元/视频 |
| 398574688191234048 | 四季写真 | 8元/视频 |
| 394125185182695432 | 生无可恋 | 10元/视频 |
| 393766210733957121 | 藏族风写真 | 14元/视频 |
| 393857704283172864 | 情书写真 | 14元/视频 |
#### 音乐
| 计费项 | 模型 | 接口说明 | 单价 |
| --- | --- | --- | --- |
| 音乐生成<br>Music Generation | music-1.5 | 音乐生成模型，支持用户提供音乐描述和歌词，生成AI音乐 | 0.25元/首 |
#### 图像
| 计费项 | 模型 | 接口说明 | 单价 |
| --- | --- | --- | --- |
| 图像生成<br>Image Generation | image-01<br>image-01-live | 支持用户通过文本描述或参考图片生成图片 | 0.025元/张 |

## 上传示例音频
- 原始链接: https://platform.minimaxi.com/document/prompt_auido?key=68b59faaa96516e260192b84
- 文档ID: `68b59faaa96516e260192b84`
# 上传示例音频


<request method="POST" url="https://api.minimaxi.com/v1/files/upload"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看。
<hr/>

#### Content-Type    <gray>multipart/form-data</gray> <red>required</red>
请求体的媒介类型，请设置为 `multipart/form-data`

<br/>

### 请求体
<hr/>


#### purpose    <gray>string</gray> <red>required</red>
文件使用目的。取值及支持格式如下：
**prompt_audio**: 快速复刻原始文件，（支持mp3、m4a、wav格式）
<hr/>

#### file <red>Required</red>
需要上传的文件。填写文件的路径地址
- 支持上传的文件需遵从以下规范：
上传的音频文件格式需为：mp3、m4a、wav格式
上传的音频文件的时长小于8s
上传的音频文件大小需不超过20mb

<br/>

### 返回体
<hr/>

#### file_id    <gray>int64</gray>
文件的唯一标识符
<hr/>

#### bytes    <gray>int64</gray>
文件大小，以字节为单位
<hr/>

#### created_at    <gray>int64</gray>
创建文件时的 Unix 时间戳，以秒为单位
<hr/>

#### filename    <gray>string</gray>
文件的名称
<hr/>

#### purpose    <gray>string</gray>
文件的使用目的，voice_clone
<hr/>

#### base_resp    <gray>object</gray>
状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0，请求成功

1002，触发限流，请稍后再试

1004，账号鉴权失败，请检查 API-Key 是否填写正确

1008，账号余额不足

1026，图片描述涉及敏感内容

2013，传入参数异常，请检查入参是否按要求填写

2049，无效的api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
url = f'https://api.minimaxi.com/v1/files/upload?GroupId={group_id}'
headers1 = {
    'authority': 'api.minimaxi.com',
    'Authorization': f'Bearer {api_key}'
}

data = {
    'purpose': 'prompt_audio'
}

files = {
    'file': open('prompt.mp3', 'rb')
}
response = requests.post(url, headers=headers1, data=data, files=files)
file_id = response.json().get("file").get("file_id")
print(file_id)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/upload?GroupId={group_id}' \
--header 'authority: api.minimaxi.com' \
--header 'content-type: multipart/form-data'  \
--header 'Authorization: Bearer MiniMax_API_KEY' \
--form 'purpose="prompt_audio"' \
--form 'file=@"Users/minimax/Downloads/示例文件.mp3"'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "file": {
        "file_id": "${file_id}",
        "bytes": 5896337,
        "created_at": 1700469398,
        "filename": "示例音频",
        "purpose": "prompt_audio"
    },
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<container gap="4" columns="2">

  <card  href="/document/guides_voice_clone1?key=68bbe68e4b8efaab4d9f722f">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[音色快速复刻指南](/document/guides_voice_clone1?key=68bbe68e4b8efaab4d9f722f)

  </card>

<card  href="/document/Voice_Cloning_api_intro?key=68adaebe66027263330017e0">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[音色快速复刻接口能力介绍](/document/Voice_Cloning_api_intro?key=68adaebe66027263330017e0)
  </card>

</container>

## 快速接入
- 原始链接: https://platform.minimaxi.com/document/quick_start?key=68b55ec4a96516e26019170a
- 文档ID: `68b55ec4a96516e26019170a`
# 新手指南
欢迎访问 MiniMax 开放平台，可参考新手指南，体验模型能力。

## 注册指南

若您是首次使用，需要先在 MiniMax 开放平台进行注册。
企业团队注册时，建议采用**主账号+子账号**的形式创建和管理。

- **个人注册**：若您是单人使用，直接在 [MiniMax开放平台](https://www.minimax.io/platform/login) 进行注册即可
- **企业团队注册**：若您是多人业务团队进行使用，建议采用**主账号+子账号**的形式进行管理
    1. 在 [MiniMax 开放平台](https://www.minimax.io/platform/login/) 注册一个账号（此账号即为主账号，注册时填写的姓名与手机号会成为本企业账号的管理员信息）
    2. 登录该主账号，在 [账户管理 > 子账号](https://platform.minimaxi.com/user-center/basic-information/child-account)，创建您所需要数量的子账户（子账号的创建数量暂时没有限制）
    3. 为您企业的人员，分配不同的子账户，进行登陆使用

### 主账号与子账号的关系：
- 子账号和主账号享用相同的使用权益与速率限制，子账号和主账号的 api 消耗可以进行共享，最后进行统一结算
- 子账号的限制：子账号无法查看和管理“支付”页面的查看与管理，也无法进行子账号、以及子账号的接口密钥的管理操作

## 快速体验大模型
MiniMax开放平台 API 调试台：更轻松、便捷。
若您希望直接体验到大模型的魅力，可以前往 [API 调试台](https://platform.minimaxi.com/examination-center/text-experience-center/cc_v2)快速体验最新模型与功能。

## 快速接入
本示例基于 [Python 3](https://www.python.org/downloads/)，帮助您快速接入 MiniMax 开放平台体验中心的对话能力，您需要完成以下步骤：

### 1. 获取鉴权信息
( 1 ). 通过 [账户管理 > 账户信息 > 基本信息](https://platform.minimaxi.com/user-center/basic-information)，获取 **group_id**。
<br>

<img src="https://filecdn.minimax.chat/public/071adb83-4f91-4103-98ef-bb22e589acbb.jpg" width="80%" />


   
( 2 ). 通过 [账户管理 > 接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)，获取 **API Key**。
点击“创建新的密钥”，在弹窗中输入密钥的名称，创建成功后，系统将展示 API Key。**请务必复制并妥善保存**，该密钥**只会显示一次**，无法再次查看。
<br>

<img src="https://filecdn.minimax.chat/public/06a482a3-29cd-41c0-a9ec-58bd83f583e5.jpg" alt="图片描述" width="80%" />


 
### 2. 导入密钥信息
在使用 API 调用模型之前，需要先将 API Key 配置到环境变量中，避免在代码中明文出现，从而降低泄漏风险。

##### 如您是 macOS/Linux 用户

( 1 ). 打开终端，输入以下命令，将 `<MINIMAX_API_KEY>` 替换为您在第一步获取的密钥信息。
   ```bash
   export MINIMAX_API_KEY="<MINIMAX_API_KEY>"
   ```
<br/>

( 2 ). 执行以下命令，验证该环境变量是否生效，若返回您的 API Key，则表示设置成功。
   ```bash
   echo $MINIMAX_API_KEY
   ```
<br/>

##### 如您是 Windows 用户

- **在 CMD 中**

( 1 ). 打开 CMD (命令提示符)，输入以下命令，将 `<MINIMAX_API_KEY>` 替换为您在第一步获取的密钥信息。

```dos
set MINIMAX_API_KEY=<MINIMAX_API_KEY>
```

<br/>

( 2 ). 执行以下命令，验证该环境变量是否生效，若返回您的 API Key，则表示设置成功。

```dos
echo %MINIMAX_API_KEY%
```
<br/>


- **在 PowerShell 中**

( 1 ). 打开 PowerShell，输入以下命令，将 `<MINIMAX_API_KEY>` 替换为您在第一步获取的密钥信息。

```powershell
$env:MINIMAX_API_KEY = "<MINIMAX_API_KEY>"
```

<br/>

( 2 ). 执行以下命令，验证该环境变量是否生效，若返回您的 API Key，则表示设置成功。

```powershell
$env:MINIMAX_API_KEY
```
<br/>

### 3. 复制完整代码并执行
以下是一份可以在 Python3 终端环境中直接执行的完整代码，可将代码复制进 `example.py` 文件中。
**注意**：添加 `import readline` 引用是为了解决在中文输入下，Python 的 input 接口在删除字符时错误处理的问题。

```python
import os
import requests
import readline

api_key = os.environ.get("MINIMAX_API_KEY")

url = "https://api.minimaxi.com/v1/text/chatcompletion_v2"
headers = {
    "Authorization": f"Bearer {api_key}", 
    "Content-Type": "application/json"
    }

request_body = {
    "model": "MiniMax-M1",
    "tokens_to_generate": 8192,
    "messages": [
        {"role": "system", "content": "You are a helpful assistant."}
    ]
}

while True:
    line = input("User input: ")
    if line.lower() in ['quit', 'exit']:
        break

    request_body["messages"].append({"role": "user", "content": line})
    
    response = requests.post(url, headers=headers, json=request_body)
    response_data = response.json()
    reply = response_data["choices"][0]["message"]["content"]
    reasoning_content = response_data["choices"][0]["message"]["reasoning_content"]
    print(f"Reasoning: {reasoning_content}")
    print(f"Reply: {reply}")

    request_body["messages"].append({"role": "assistant", "content": reply})

```

在终端中执行命令，运行代码：
```bash
python example.py
```
<br/>

### 代码分步解析
如果您想进一步了解上述完整代码是如何逐步构建的，您可以继续阅读以下详细步骤解释。

##### 1. 构建请求头
下面代码根据鉴权信息构建请求头（`api_key` 为您的密钥信息）。
```python
url = "https://api.minimaxi.com/v1/text/chatcompletion_v2"
headers = {
    "Authorization": f"Bearer {api_key}", 
    "Content-Type": "application/json"
    }
```

##### 2. 构建请求内容
`request_body` 部分为构建初始请求体，初始化与模型交互的请求参数。获取用户输入后，将用户输入添加至请求体中。
```python
request_body = {
    "model": "MiniMax-M1",
    "tokens_to_generate": 8192,
    "messages": [
        {"role": "system", "content": "You are a helpful assistant."}
    ]
}
line = input("User input: ")
request_body["messages"].append({"role": "user", "content": line})
```

##### 3. 完成交互
通过 requests 库提供的 post 能力对 API 进行调用，获取模型的推理内容和输出内容。
**注意**：每一轮回复内容都需要追加到 `messages` 中，这样才能在多轮的对话中记住对话历史。
```python
response = requests.post(url, headers=headers, json=request_body)
response_data = response.json()

reply = response_data["choices"][0]["message"]["content"]
reasoning_content = response_data["choices"][0]["message"]["reasoning_content"]

request_body["messages"].append({"role": "assistant", "content": reply})
```

## 资源保障
- 原始链接: https://platform.minimaxi.com/document/rate_limits?key=68b90c85630660297b0679fc
- 文档ID: `68b90c85630660297b0679fc`
# 资源保障

为保证资源的高效使用，引入速率限制（Rate Limits），以确保服务的可用性、稳定性。

## 1、什么是速率限制

速率限制（Rate Limits）是 API 对用户在指定时间段内可以访问服务器的次数/消耗的 token 数，所作出的限制。  
其中 MiniMax 的 API 的速率限制分为：**RPM** 以及 **TPM** 两种。  
- **RPM** 是指 API 对用户在指定时间段内可以访问服务器的次数所作出的限制
- **TPM** 是指 API 对用户在指定时间段内可以访问服务器所消耗的 token 数所作出的限制

## 2、为什么会有速率限制

速率限制是 API 的常见做法，实施速率限制有几个不同的原因：

- 有助于防止 API 被滥用和误用。例如，防止有些用户恶意的超载调用 API，试图使其过载或导致服务中断。我们通过设置速率限制可以防止此类恶意使用的发生
- 速率限制有助于确保每个人都可以公平访问 API。防止一个人或组织发出过多的请求，造成可能会使其他人的 API 资源分配不均的情况。通过限制单个用户可以发出的请求数量，可以确保最多的人有机会使用 API，而不会遇到速度变慢的情况。  
- 通过设置速率限制，MiniMax 可以帮助为所有用户保持流畅一致的体验

## 3、我们的 API 的限速具体数值

我们会根据您使用的模型、接口以及您拥有的账户类型，对您的账号（包括主账号+子账号）实施相应的速率限制策略。即您的主账号和子账号共同享有以下所有速率限制。  
例如：当您的账号拥有 120 个 RPM 时，你的账号每分钟内可以发送 120 个请求，当您的主账号每分钟发送了 30 个请求时，您的子账号该时间段内还可以发送 90 个请求。  

MiniMax 的速率以两种维度进行限制：  
- RPM：每分钟发送的请求数限制  
- TPM：每分钟输入+输出的 token 数限制  

<br/>

**具体速率限制情况如下表所示：**

- 文本接口的速率限制如下：

| 接口名 | ChatCompletion v2 | ChatCompletion v2 |
| --- | --- | --- |
| 模型 |MiniMax-M1 <br>MiniMax-Text-01 |MiniMax-M1 <br>MiniMax-Text-01 |
| 客户类型\限制类型 | RPM | TPM |
| 免费用户 | 6 | 24000 |
| 充值用户 | 120 | 720000 |

<br/>

- 语音接口速率限制如下：

| 接口名 | T2A v2（语音生成） |T2A v2（语音生成） |
| --- | --- |--- |
| 模型 | speech-02-hd <br> speech-02-turbo <br> speech-01-hd <br> speech-01-turbo  |speech-2.5-hd-preview <br> speech-2.5-turbo-preview  |
| 客户类型\限制类型 | RPM | RPM |
| 免费用户 | 3 | 10 |
| 充值用户 | 20 | 20 |

<br/>

- 视频接口速率限制如下：

| 接口名 | Video Generation（视频生成） |
| --- | --- |
| 模型 | 02系列：MiniMax-Hailuo-02<br>01系列：T2V-01、I2V-01、T2V-01-Director、I2V-01-Director、I2V-01-live、S2V-01 |
| 客户类型/限制类型 | RPM |
| 免费用户 | 5 |
| 充值用户 | 20 |

<br/>

- 图片接口速率限制如下：

| 接口名 | Image Generation | Image Generation |
| --- | --- | --- |
| 客户类型/限制类型 | RPM | TPM |
| 免费用户 | 10 | 60 |
| 付费用户 | 10 | 60 |

<br/>

- 音乐接口速率限制如下：

| 接口名 | Music Generation（音乐生成） | Music Generation（音乐生成） |
| --- | --- | --- |
| 模型 | music-1.5 | music-1.5|
| 客户类型\限制类型 | RPM | CONN（最大并行运行任务数） |
| 免费用户 | 6 | 3 |
| 充值用户 | 120 | 20 |


## 4、速率是如何被限制的

如果您是免费用户，在 T2T 的功能场景下，使用任意接口的任意模型时（如 ChatCompletion v2 - MiniMax-Text-01），此时您的速率限制为 **RPM=6，TPM=12000**，即：您的速率限制为每分钟 6 个请求和每分钟 12000 个 token 数（输入+输出）。

## 5、如果遇到限速的报错返回会怎样

如果您收到速率限制的返回报错，则表示您在短时间内发出了太多请求，或者短时间发出了太多的字符，那么 API 将会拒绝满足进一步的请求，直到经过指定的时间。

## 6、速率限制和 tokens_to_generate、max_tokens

因无法在请求输入时得到输入+输出的总 token，因此在请求输入时，会首先参考接口里设置的 `tokens_to_generate` 和 `max_tokens` 数值作为该请求的总 token 和 TPM 限流参考，当该条请求生成完所有 token 后，会使用实际的输入+输出总 token 修正该条请求所对应的 TPM。  
所以您可以尝试将 `tokens_to_generate` 和 `max_tokens` 的值设置为尽可能符合预期的响应值，从而减少 TPM 限流报错。

## 7、如何更好的在限速策略下进行 API 的使用

由于 API 对每分钟请求数和每分钟 token 数有单独的限制，所以我们建议您集中处理请求。  
如果您的每分钟请求数量已经达到了限制，但是每分钟的 token 数量仍然具有可用容量，则建议您可以通过将多个任务批量放到每个请求来提高 token 的吞吐量。这样将允许您每分钟处理更多的 token 数。

## 8、如何提高速率限制

我们的默认限速是为了帮助最大限度的提高 API 调用的稳定性，以及更公平合理的将相应资源进行分配。  
如果您认为您拥有必要的流量数据以来申请提高速率限制时，您可以通过您这边的商务经理向我们进行提高速率限制的申请。  
如果您这边没有对接的商务经理，请发送邮件到 **api@minimaxi.com** ，我们会 2 个工作日内与您联系。  

需要您知悉的是，提高限速有时可能需要 3-5 个工作日，因此如果您有确定数据支持您的申请时，为了不影响您的产品的正常上线发布，请务必尽早提交限速提升申请。

## Retrieval
- 原始链接: https://platform.minimaxi.com/document/retrieval?key=66718faca427f0c8a5701617
- 文档ID: `66718faca427f0c8a5701617`
### Retrieval（知识库检索）
该接口支持高效检索文档信息，准确回答专业问题，为大模型补充领域知识。相关操作流程如下：
1. 文档上传file，详情可见[File（文档）](https://platform.minimaxi.com/document/file?key=6685458335a2d55137ca9681)，得到file_id；
2. 新建知识库时关联上传好的file_id，并得到knowledge_base_id；
3. retrieval knowledge_base_id 查看知识库建库状态；
4. 建库成功后，调用Chatcompletion或者Chatcompletion pro时配置knowledge_base_id，使用query检索到的knowledge_base_id分片信息回答问题。
### API说明
`API：https://api.minimax.chat/v1/embedding`
#### 支持token数
| 限制内容 | 限制数量 |
| --- | --- |
| 知识库数量 | 10个 |
| 单个知识库总容量 | 5000万字符 |
| 每个文档容量 | 200万字符 |
#### 接口参数说明
##### 新建知识库
###### 请求体（Request）参数
###### 返回（response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| operator_id | int64 | 是 | 操作流水ID（幂等标识）。用同一个流水ID调用会使用第一次的调用参数，使用场景：对于同一个文档的建立知识库操作可能会存在建库时间长的情况，当使用同一个operator_id时可以防止重复操作。 |
| name | string | 是 | 知识库名称。不能重复。 |
| description | string | 否 | 知识库描述。 |
| embedding_model | string | 是 | 知识库模型，仅支持：`kbq-001` |
| file_id | int64 | 否 | 通过files接口上传的文件ID。<br>不填仅创建知识库。 <br>填写先创建知识库，再增加文档。 |
| doc_params | object | 否 | 增加文档参数。请注意，此参数仅对当前请求中同时上传的文档生效。<br>切片的最大token数。默认500。<br>chunk_size >= 2*chunk_overlap + 1<br>如果同时配置了分隔符，会先以分隔符进行切分，然后再用size大小进行拼接。<br>切片与顺序相邻切片重叠的最大token数。默认50。<br>chunk_overlap >= 0<br>注：如果正则情况下，该功能是不生效的。<br>分隔符。默认["\n", ""]。<br>分割符是否为正则表达式。默认false 。 |
| doc_params.chunk_size | int64 | 否 | 切片的最大token数。默认500。<br>chunk_size >= 2*chunk_overlap + 1<br>如果同时配置了分隔符，会先以分隔符进行切分，然后再用size大小进行拼接。 |
| doc_params.chunk_overlap | int64 | 否 | 切片与顺序相邻切片重叠的最大token数。默认50。<br>chunk_overlap >= 0<br>注：如果正则情况下，该功能是不生效的。 |
| doc_params.separators | []string | 否 | 分隔符。默认["\n", ""]。 |
| doc_params.is_regex | bool | 否 | 分割符是否为正则表达式。默认false 。 |
| callback | string | 否 | 回调地址。接收创建知识库状态的地址。 |
| knowledge_base_id | int64 | 否 | 知识库ID。 |
| file_id | int64 | 否 | 文件的id。 |
| knowledge_base | object | 否 | 知识库信息。<br>知识库的id。<br>知识库的名称。<br>知识库的描述。<br>知识库的模型。<br>知识库创建的时间，毫秒。<br>最近一次插入文档的时间，毫秒。<br>最近一次删除文档的时间，毫秒。<br>知识库所使用的文档。<br>知识库的状态。<br>loading-操作进行中<br>available-无操作进行中 |
| knowledge_base.knowledge_base_id | int64 | 否 | 知识库的id。 |
| knowledge_base.name | string | 否 | 知识库的名称。 |
| knowledge_base.description | string | 否 | 知识库的描述。 |
| knowledge_base.embedding_model | string | 否 | 知识库的模型。 |
| knowledge_base.create_time | int64 | 否 | 知识库创建的时间，毫秒。 |
| knowledge_base.insert_time | int64 | 否 | 最近一次插入文档的时间，毫秒。 |
| knowledge_base.delete_time | int64 | 否 | 最近一次删除文档的时间，毫秒。 |
| knowledge_base.docs | []object | 否 | 知识库所使用的文档。 |
| knowledge_base.status | string | 否 | 知识库的状态。<br>loading-操作进行中<br>available-无操作进行中 |
| usage | object | 否 | 知识库计费为异步，需要通过账单查看用量。 |
| base_resp | object | 否 | 状态码和详情。<br>状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
##### 删除知识库
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| operator_id | int64 | 是 | 操作流水ID（幂等标识）。用同一个流水ID调用会使用第一次的调用参数，使用场景：对于同一个文档的建立知识库操作可能会存在建库时间长的情况，当使用同一个operator_id时可以防止重复操作。 |
| knowledge_base_id | int64 | 是 | 知识库ID。 |
| base_resp | object | 否 | 错误状态码和详情。<br>状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
##### 查看知识库详情
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| knowledge_base_id | int64 | 是 | 知识库ID。 |
| base_resp | object | 否 | 错误状态码和详情。<br>状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；；2036，知识库操作正在处理中。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
##### 查看知识库列表
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| page | int64 | 是 | 分页序号，从0开始。 |
| page_size | int64 | 是 | 分页大小，(0,100]。 |
| total | int64 | 否 | 总数量。 |
| knowledge_bases | object | 否 | 知识库的信息列表。<br>知识库ID。<br>知识库名称。<br>知识库描述。<br>知识库模型。<br>知识库创建时间，毫秒。<br>最近一次插入文档时间，毫秒。<br>最近一次删除文档时间，毫秒。<br>知识库所使用的文档。<br>知识库的状态。loading-操作进行中；available-无操作进行中。 |
| knowledge_bases.knowledge_base_id | int64 | 否 | 知识库ID。 |
| knowledge_bases.name | string | 否 | 知识库名称。 |
| knowledge_bases.description | string | 否 | 知识库描述。 |
| knowledge_bases.embedding_model | string | 否 | 知识库模型。 |
| knowledge_bases.create_time | int64 | 否 | 知识库创建时间，毫秒。 |
| knowledge_bases.insert_time | int64 | 否 | 最近一次插入文档时间，毫秒。 |
| knowledge_bases.delete_time | int64 | 否 | 最近一次删除文档时间，毫秒。 |
| knowledge_bases.docs | []object | 否 | 知识库所使用的文档。 |
| knowledge_bases.status | string | 否 | 知识库的状态。loading-操作进行中；available-无操作进行中。 |
| base_resp | object | 否 | 错误状态码和详情。<br>状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
##### 增加知识库文档
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| operator_id | int64 | 是 | 操作流水ID（幂等标识）。用同一个流水ID调用会使用第一次的调用参数，使用场景：对于同一个文档的建立知识库操作可能会存在建库时间长的情况，当使用同一个operator_id时可以防止重复操作。 |
| knowledge_base_id | int64 | 是 | 知识库id。 |
| file_id | int64 | 是 | 通过files接口上传的文件ID。<br>需要增加到知识库的文档的id。 |
| doc_params | object | 否 | 增加文档参数。<br>切片的最大token数。默认500。<br>chunk_size >= 2*chunk_overlap + 1<br>切片与顺序相邻切片重叠的最大token数。默认50。<br>chunk_overlap >= 0<br>分隔符。默认["\n", ""]。<br>分割符是否为正则表达式。默认false 。 |
| doc_params.chunk_size | int64 | 否 | 切片的最大token数。默认500。<br>chunk_size >= 2*chunk_overlap + 1 |
| doc_params.chunk_overlap | int64 | 否 | 切片与顺序相邻切片重叠的最大token数。默认50。<br>chunk_overlap >= 0 |
| doc_params.separators | []string | 否 | 分隔符。默认["\n", ""]。 |
| doc_params.is_regex | bool | 否 | 分割符是否为正则表达式。默认false 。 |
| callback | string | 否 | 回调地址。 |
| knowledge_base_id | int64 | 否 | 知识库ID。 |
| knowledge_base | object | 否 | 知识库信息。<br>知识库id。<br>知识库名称。<br>知识库描述。<br>知识库模型。<br>知识库创建时间，毫秒。<br>最近一次插入文档时间，毫秒。<br>最近一次删除文档时间，毫秒。<br>知识库所使用的文档。<br>loading-操作进行中<br>available-无操作进行中 |
| knowledge_base.knowledge_base_id | int64 | 否 | 知识库id。 |
| knowledge_base.name | string | 否 | 知识库名称。 |
| knowledge_base.description | string | 否 | 知识库描述。 |
| knowledge_base.embedding_model | string | 否 | 知识库模型。 |
| knowledge_base.create_time | int64 | 否 | 知识库创建时间，毫秒。 |
| knowledge_base.insert_time | int64 | 否 | 最近一次插入文档时间，毫秒。 |
| knowledge_base.delete_time | int64 | 否 | 最近一次删除文档时间，毫秒。 |
| knowledge_base.docs | []object | 否 | 知识库所使用的文档。 |
| knowledge_base.status | string | 否 | loading-操作进行中<br>available-无操作进行中 |
| file_id | int64 | 否 | 文件id。 |
| usage | object | 否 | 知识库计费为异步，需要通过账单查看用量。 |
| base_resp | object | 否 | 错误状态码和详情。<br>状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
##### 删除知识库文档
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| operator_id | int64 | 是 | 操作流水ID（幂等标识）。用同一个流水ID调用会使用第一次的调用参数。 |
| knowledge_base_id | int64 | 是 | 知识库ID。 |
| file_id | int64 | 是 | 文件ID。 |
| base_resp | object | 否 | 错误状态码和详情。<br>状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
##### 修改知识库文档
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| operator_id | int64 | 是 | 操作流水ID（幂等标识）。用同一个流水ID调用会使用第一次的调用参数。 |
| knowledge_base_id | int64 | 是 | 知识库id。 |
| file_id | int64 | 是 | 文件id。 |
| doc_params | object | 否 | 增加文档参数。 |
| callback | string | 否 | 回调地址。 |
| knowledge_base_id | int64 | 否 | 知识库ID。 |
| file_id | int64 | 否 | 文件ID。 |
| knowledge_base | object | 否 | 知识库信息。<br>知识库ID。<br>知识库名称。<br>知识库描述。<br>知识库模型。<br>知识库创建时间，毫秒。<br>最近一次插入文档时间，毫秒。<br>最近一次删除文档时间，毫秒。<br>知识库所使用的文档。<br>loading-操作进行中<br>available-无操作进行中 |
| knowledge_base.knowledge_base_id | int64 | 否 | 知识库ID。 |
| knowledge_base.name | string | 否 | 知识库名称。 |
| knowledge_base.description | string | 否 | 知识库描述。 |
| knowledge_base.embedding_model | string | 否 | 知识库模型。 |
| knowledge_base.create_time | int64 | 否 | 知识库创建时间，毫秒。 |
| knowledge_base.insert_time | int64 | 否 | 最近一次插入文档时间，毫秒。 |
| knowledge_base.delete_time | int64 | 否 | 最近一次删除文档时间，毫秒。 |
| knowledge_base.docs | []object | 否 | 知识库所使用的文档。 |
| knowledge_base.status | string | 否 | loading-操作进行中<br>available-无操作进行中 |
| usage | object | 否 | 知识库计费为异步，需要通过账单查看用量。 |
| base_resp | object | 否 | 错误状态码和详情。<br>状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
##### 查看知识库文档
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | 给到的API密钥。 |
| Content-Type | application/json | 是 | Content-Type。 |
| Groupid | string | 是 | 用户所属的组。 |
| operator_id | int64 | 是 | 操作流水ID（幂等标识）。用同一个流水ID调用会使用第一次的调用参数。 |
| knowledge_base_id | int64 | 是 | 知识库ID。 |
| file_id | int64 | 是 | 文件ID。 |
| document | object | 否 | 文档信息。<br>文件ID。 |
| document.document.file_id | int64 | 否 | 文件ID。 |
| knowledge_base | object | 否 | 知识库信息。<br>知识库ID。<br>知识库名称。<br>知识库描述。<br>知识库模型。<br>知识库创建时间，毫秒。<br>最近一次插入文档时间，毫秒。<br>最近一次删除文档时间，毫秒。<br>知识库所使用的文档。<br>loading-操作进行中<br>available-无操作进行中 |
| knowledge_base.knowledge_base_id | int64 | 否 | 知识库ID。 |
| knowledge_base.name | string | 否 | 知识库名称。 |
| knowledge_base.description | string | 否 | 知识库描述。 |
| knowledge_base.embedding_model | string | 否 | 知识库模型。 |
| knowledge_base.create_time | int64 | 否 | 知识库创建时间，毫秒。 |
| knowledge_base.insert_time | int64 | 否 | 最近一次插入文档时间，毫秒。 |
| knowledge_base.delete_time | int64 | 否 | 最近一次删除文档时间，毫秒。 |
| knowledge_base.docs | []object | 否 | 知识库所使用的文档。 |
| knowledge_base.status | string | 否 | loading-操作进行中<br>available-无操作进行中 |
| base_resp | BaseResp | 否 | 错误状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。<br>状态码以及详情。<br>状态详情。 |
| base_resp.status_code | int64 | 否 | 错误状态码。2032，知识库个数达到上限；2033，知识库文档字符数达到上限；2034，当前知识库上一个操作未完成，同一个知识库的删除知识库、增加文档、删除文档、修改文档需串行执行；2036，知识库操作正在处理中。 |
| base_resp.status_msg | string | 否 | 状态详情。 |
##### 回调通知
增加文档操作时填callback参数时，在操作完成后会通过回调地址发送一次请求通知知识库创建状态。
Method：`POST`
**Body**
```shell
{
    "operator_id": 123,
    "knowledge_base": {
        "knowledge_base": {
            "knowledge_base_id": 863831823174049792,
            "name": "test3",
            "description": "",
            "embedding_model": "kbq-001",
            "create_time": 1700742086395,
            "insert_time": 1700796361445,
            "delete_time": 1700796932183,
            "docs": [],
            "status": "available"
        },
        "base_resp": {
            "status_code": 0,
            "status_msg": "success"
        }
    }
}
```
##### 检索 Chatcompletion、Chatcompletion Pro
###### 新增请求（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| knowledge_base_param | object | 否 | 知识库参数。<br>知识库ID。<br>默认值为5。<br>取值为[0,1),默认为0。<br>默认使用：<br>提供 {{context}} {{question}} ，可由用户自己定义。<br>比如：<br>默认为0,从最后一句话开始倒数，>= 0，适用于多步问答的场景。<br>当 num_prev_messages为1时，会将第2、3行的内容一起用来作为query。<br>应用场景： |
| knowledge_base_param.knowledge_base_id | int64 | 是 | 知识库ID。 |
| knowledge_base_param.top_k | int | 否 | 默认值为5。 |
| knowledge_base_param.thresh | float | 否 | 取值为[0,1),默认为0。 |
| knowledge_base_param.pattern | string | 否 | 默认使用：<br>提供 {{context}} {{question}} ，可由用户自己定义。<br>比如： |
| knowledge_base_param.num_prev_messages | int | 否 | 默认为0,从最后一句话开始倒数，>= 0，适用于多步问答的场景。<br>当 num_prev_messages为1时，会将第2、3行的内容一起用来作为query。<br>应用场景： |
| knowledge_base | object | 否 | 知识库信息。<br>知识库ID。<br>片段。 |
| knowledge_base.knowledge_base_id | int64 | 否 | 知识库ID。 |
| knowledge_base.chunks | []object | 否 | 片段。 |
| chunks | []object | 否 | 片段。 |
| content | int64 | 否 | 片段内容。 |
| index | int | 否 | 片段index。 |
| document | string | 否 | 片段所属文件ID。 |

## 创建主体参考生成视频任务
- 原始链接: https://platform.minimaxi.com/document/s2v?key=68b5a347a96516e260193093
- 文档ID: `68b5a347a96516e260193093`
# 主体参考视频生成任务

<request method="POST" url="https://api.minimaxi.com/v1/video_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 

- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
模型名称。可用值：`S2V-01`
<hr/>

#### prompt    <gray>string</gray>
视频的文本描述，最大 2000 字符


#### prompt_optimizer    <gray>boolean</gray>
是否自动优化 `prompt`，默认为 `true`。设为 `false` 可进行更精确的控制
<hr/>


#### subject_reference    <gray>array</gray>
主体参考，仅当 `model` 为 `S2V-01` 时可用。目前仅支持单个主体
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### type    <gray>string</gray> <red>required</red>
主体类型，当前仅支持 `character` (人物面部)
<hr/>

#### image    <gray>array</gray> <red>required</red>
包含主体参考图的数组（目前仅支持单张图片）
- 图片要求：
  - 格式：JPG, JPEG, PNG, WebP
  - 体积：小于 20MB
  - 尺寸：短边像素大于 300px，长宽比在 2:5 和 5:2 之间

</card>
</details>
<hr/>

#### callback_url    <gray>string</gray>
接收任务状态更新通知的回调 URL。支持通过 `callback_url` 参数可以配置回调，以接收任务状态的更新的异步通知

1. 地址验证：配置后，MiniMax 服务器会向 `callback_url` 发送一个 `POST` 请求，请求体中包含 `challenge` 字段。服务端需要在 3 秒内原样返回该 `challenge` 值以完成验证
2. 状态更新：验证成功后，每当任务状态变更时，MiniMax 都会向该 URL 推送最新的任务状态。推送的数据结构与调用查询视频生成任务接口的响应体一致


回调返回的"status"包括以下状态：
- "processing"-生成中
- "success"-成功
- "failed"-失败


<codetabs title="回调服务示例">
<codetab  language="python">
from fastapi import FastAPI, HTTPException, Request
from fastapi.middleware.cors import CORSMiddleware
import json
app = FastAPI()
@app.post("/get_callback")
async def get_callback(request: Request):
    try:
        json_data = await request.json()
        challenge = json_data.get("challenge")
        if challenge is not None:
          # is a verification request, just return the challenge
          return {"challenge": challenge}
        else:
            # is a callback request, do your own logic here
            # {
            #     "task_id": "115334141465231360",
            #     "status": "Success",
            #     "file_id": "205258526306433",
            #     "base_resp": {
            #         "status_code": 0,
            #         "status_msg": "success"
            #     }
            # }
            return {"status": "success"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        app, # 必选
        host="0.0.0.0", # 必选
        port=8000, # 必选，端口可设置
        # ssl_keyfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
        # ssl_certfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
    )
</codetab>

</codetabs>


<hr/>

#### aigc_watermark    <gray>boolean</gray>
是否在生成的视频中添加水印，默认为 `false`

<br/>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
视频生成任务的 ID，用于后续查询任务状态
<hr/>

#### base_resp    <gray>object</gray>
状态码及详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>integer</gray>
状态码及其分别含义如下：

0：请求成功；

1002：触发限流，请稍后再试

1004：账号鉴权失败，请检查 API-Key 是否填写正确

1008：账号余额不足

1026：视频描述涉及敏感内容，请调整

2013：传入参数异常，请检查入参是否按要求填写

2049：无效的api key，请检查api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/video_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "prompt": "女生跑向镜头前，微笑着眨了眨眼睛。",
    "subject_reference": [
        {
            "type": "character",
            "image": [
                "https://cdn.hailuoai.com/prod/2025-08-12-17/video_cover/1754990600020238321-411603868533342214-cover.jpg"
            ],
        }
    ],
    "model": "S2V-01"
  
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())

</codetab>
<codetab  language="bash">
curl -X POST "https://api.minimaxi.com/v1/video_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "prompt": "女生跑向镜头前，微笑着眨了眨眼睛。",
    "subject_reference": [
      {
        "type": "character",
        "image": [
          "https://cdn.hailuoai.com/prod/2025-08-12-17/video_cover/1754990600020238321-411603868533342214-cover.jpg"
        ],
      }
    ],
    "model": "S2V-01"
  }'
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab  language="json">
{
    "task_id": "106916112212032",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

<br>

</container>

<container gap="4" columns="2">

  <card  href="/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成指南](/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7)

  </card>

<card  href="/document/video_api_intro?key=68abdfdaaf3fef26183d6501">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成接口能力介绍](/document/video_api_intro?key=68abdfdaaf3fef26183d6501)
  </card>

</container>

## 一键式AI有声书工坊
- 原始链接: https://platform.minimaxi.com/document/solutions_audio_book?key=68bc2a4f95a600dbfc92c7c8
- 文档ID: `68bc2a4f95a600dbfc92c7c8`
# 一键式AI有声书工坊 

本解决方案专为 **小说阅读平台** 设计，调用核心的 大规模语音合成API (Large-Scale TTS API)，能将平台海量的静态文本小说，一键式、规模化地转化为富有情感、多角色演绎的高品质有声读物，极大丰富内容生态，提升用户粘性。

><img src="https://filecdn.minimax.chat/public/3ddc9d90-84a8-4c82-82c7-05b8236e71c0.png" alt="图片描述" width="14px" />
>核心攻克：长文情感一致性、多角色音色生成、万字级章节秒级输出。


本方案旨在帮助合作平台 **指数级降低有声书制作成本、极大提升内容音频化覆盖率**，并为用户提供 **媲美真人演播的沉浸式听书体验**，从而构筑强大的内容护城河。

## 行业痛点

我们深刻理解小说平台在有声内容领域面临的核心挑战

<container gap="4" columns="3">

<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 高昂的制作成本与周期
传统有声书制作需聘请专业CV、录音棚、后期制作，单本成本高昂（数万至数十万），制作周期长达数周甚至数月。

</card>
<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 情感表达平淡，千篇一律
传统TTS（文本转语音）技术生成的音频情感匮乏，语调机械，无法表现小说中人物的喜怒哀乐，导致听感不佳。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 上下文割裂，叙事不连贯
 长篇小说中，普通TTS难以理解上下文关联，导致角色情绪、语气在章节间断裂，严重影响故事的连贯性和沉浸感。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 内容更新缓慢，覆盖率低

面对平台每日更新的海量章节，人工制作速度远无法匹配，导致绝大多数作品无法提供有声版本，错失大量“听书”用户。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 音色单调，角色区分难
平台难以提供足够丰富和个性化的音色库，无法满足不同类型小说（如玄幻、言情、悬疑）的风格需求，多角色对话时更是难以区分。

</card>

</container>
  



## 核心目标
<container gap="4" columns="3">


<card>
<img src="https://filecdn.minimax.chat/public/1ba7fc29-9001-456b-b28d-f6a49b20cbe0.png" width="32px" />

### 确保情感与语境的高度一致性
实现对万字以上长文本的深度理解，确保旁白与角色情感在全篇章中保持连贯、自然，符合故事情节发展。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/1ba7fc29-9001-456b-b28d-f6a49b20cbe0.png" width="32px" />

### 实现丰富且可定制的多角色演播
提供一个庞大且持续扩展的“虚拟CV”音色库，支持AI自动为不同角色匹配独特音色，并允许平台进行个性化定制。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/1ba7fc29-9001-456b-b28d-f6a49b20cbe0.png" width="32px" />

### 极致提升有声内容的生产效率
 将传统数周的制作周期压缩至分钟级。支持高达35000字符的单次输入，实现整章小说“秒级”生成，让全站小说音频化成为可能。

</card>


</container>

## 解决方案

### 第一步：整章文本智能注入

将待转换的小说章节文本（支持最高一百万字符）通过API一次性提交。系统会自动进行文本预处理，如识别章节标题、旁白、对话等。

### 第二步：AI导演智能分析

我们的大模型将扮演“AI导演”的角色：

- **上下文理解：** 准确解析文章意图，理解人物关系和情节走向。  
- **角色识别：** 自动识别对话中的不同角色，并从音色库中匹配最合适的声线。  
- **情感分析：** 精准识别每一句话的情感色彩（如激动、悲伤、紧张、温柔），为后续的语音生成提供“表演指导”。  

### 第三步：动态语音合成与交付 

基于分析结果，AI进行最终的语音合成：

- **多音色融合：** 动态切换不同角色的声音，旁白沉稳，角色鲜活。  
- **情感化韵律：** 生成的语音在语速、停顿、重音和语调上充满变化，完美贴合文本情感。  
- **快速交付：** 任务完成后，API返回高品质的MP3音频文件URL，可直接用于播放或分发。  

## 业务价值

<container gap="4" columns="2">

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 指数级提升内容生产力
制作成本降低95%以上，生产效率提升100倍。 快速将平台全量小说资产转化为有声内容，实现从“部分覆盖”到“全量覆盖”的跨越。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 创造卓越的用户听书体验
提供媲美真人团队精心制作的听书体验，多角色、情感化、不串戏。显著提升用户平均收听时长、完播率及付费转化率。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 构筑强大的内容差异化壁垒
快速上线海量独家有声书，吸引并锁定“听书”用户群体。通过提供独特的AI音色，打造平台专属的听书品牌认知。

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 保障数据隐私与安全合规
所有文本数据在处理过程中均采用严格的加密与脱敏技术，确保合作方的内容资产与用户数据安全合规，无后顾之忧。

</card>


</container>

## 核心API能力

本解决方案主要依赖以下三个API接口：

##### 1. 创建有声书生成任务
<request method="POST" url="https://api.minimaxi.com/v1/t2a_async_v2"></request>

- **用途：** 创建一个异步的有声书生成任务。这是最核心的调用。  
- **关键参数：**  
1. `text` 小说文本内容  
1. `voice_setting` 语音合成中的设置，如指定模式为多角色、情感分析开关等  
1.  `audio_setting` 按需选择配置，指语音生成的优先格式  

<br>

##### 2. 查询任务状态
<request method="POST" url="http://api.minimaxi.com/v1/query/t2a_async_query_v2"></request>

- **用途：** 查询指定任务的当前状态（如排队中、处理中、已完成、失败）。  
- **关键参数：**  
1.  `task_id` 创建任务时返回的唯一ID  

<br>

##### 3. 获取可用音色列表
<request method="POST" url="https://api.minimaxi.com/v1/get_voice"></request>

- **用途：** 获取当前所有可用的AI音色列表及其特征标签（如“少年音”、“御姐音”、“沉稳大叔”、“旁白”）。  
- **应用场景：** 为用户提供音色选择功能。  

## 使用示例

以下是一个使用Python 创建有声书生成任务的示例代码

```python
import requests
import json

api_key = "填写你的api key"

url = "https://api.minimaxi.com/v1/t2a_async_v2"
payload = {
    'role_voice':'{\
    "title": "audiobook_female_1",\
    "anarration": "audiobook_male_1",\
    "年老道士": "audiobook_male_2",\
    "年轻男性": "gaodun_yanhan_kangfen"}',
    'interval_title': '1',
    'model':"speech-02",
}
files=[
  ('file',('shaosong1.json',open('C:\\Users\\minimax\\多播测试接口\\shaosong1.json','rb'),'application/json'))
]
headers = {
  'Authorization': f'Bearer {api_key}',
}

response = requests.request("POST", url, headers=headers, data=payload, files=files)

print(response.text)
````

## 智能外呼自动化
- 原始链接: https://platform.minimaxi.com/document/solutions_outboundbot?key=68ba5dc7630660297b073d4c
- 文档ID: `68ba5dc7630660297b073d4c`
# 智能外呼自动化 

**利用 MiniMax Text & Speech API 构建新一代客户触达体验。**

本案例展示了领先企业如何利用 MiniMax **Text-01** 与 **Speech-02** API 实现外呼业务的自动化——将基于意图的脚本，快速转化为超拟真、富有情感的语音对话。

通过结合**智能客服文本生成**与​**超拟人语音合成**​，销售与客服团队可以在不牺牲质量的前提下，大幅提升触达效率、统一品牌话术，并实现大规模客户沟通。



## 行业痛点


<container gap="4" columns="3">

<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 人力成本高
人工外呼需要庞大的团队，招聘与培训成本居高不下

</card>
<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 服务质量不一致
坐席人员在语气、用词和客户应对上存在差异，导致品牌体验不统一

</card>
<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 扩展性差
 人工坐席在高峰期难以保持稳定表现

</card>

<card>
<img src="https://filecdn.minimax.chat/public/5ea2eeff-503a-4b72-92dc-4bbbeb66a43e.png" width="32px" />

### 个性化不足

固定脚本难以根据客户的实时反应灵活调整

</card>



</container>
  



## 核心目标
<container gap="4" columns="3">


<card>
<img src="https://filecdn.minimax.chat/public/1ba7fc29-9001-456b-b28d-f6a49b20cbe0.png" width="32px" />

### 保证话术一致性
确保所有外呼保持统一的语气与品牌形象

</card>

<card>
<img src="https://filecdn.minimax.chat/public/1ba7fc29-9001-456b-b28d-f6a49b20cbe0.png" width="32px" />

### 实现动态脚本生成
根据客户细分与活动目标，自动生成自然、符合场景的对话内容

</card>

<card>
<img src="https://filecdn.minimax.chat/public/1ba7fc29-9001-456b-b28d-f6a49b20cbe0.png" width="32px" />

### 提升运营效率
 用自动化、高质量的AI外呼替代重复的人力工作，支持千万级别的通话规模

</card>


</container>

## 解决方案

### Text-01 智能脚本生成
   1. 针对不同行业、不同客户意图生成定制化外呼脚本。
   2. 包含个性化问候、需求探查、异议处理、成交话术等环节。
### Speech-02 声音坐席打造​
   1. 上传目标声音样本，克隆品牌专属的坐席声音。
   2. 实现超拟人的TTS合成，具备自然的语调、情感表达与节奏控制。
### Real-time实时联动播放
   1. 将 Text-01 输出直接对接 Speech-02，实现实时播放对话。
   2. 根据通话内容动态调整语速、语气与重点。
  

## 业务价值

<container gap="4" columns="2">

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 显著降低成本
外呼中心人力成本降低高达80%

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 统一品牌声音
每一次外呼都能传递同样的专业、亲和形象

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 大规模个性化
为数百万潜在客户提供定制化沟通体验

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 快速上线活动
从策划到外呼启动仅需数小时，而非数周

</card>


</container>

<br>

## 核心API能力

1. **MiniMax-Text-01 智能脚本生成：**

   - **功能：** 针对不同行业（如保险续保、招聘邀约）、不同客户意图，自动生成包含个性化问候、需求探查、异议处理、成交话术等环节的完整脚本。

2. **Speech-02 品牌专属声音坐席打造：**

   - **功能：** 支持上传指定声音样本（如优秀员工作为“声音模特”），克隆出品牌专属的、独一无二的AI坐席声音。通过超拟人的TTS合成技术，实现自然的语调、情感表达与节奏控制，告别生硬的“机器音”。

3. **Real-time 实时联动与动态响应：**
   
   - **功能：** 将Text-01生成的文本流式对接到Speech-02，实现“边生成边播放”的低延迟对话体验。系统能根据通话中的实时内容，动态调整语速、语气和重点，实现与真人的高度相似的交互。
   - **API**​**集成示例（Text-to-Speech流式接口）：** 代码逻辑展示了如何调用 `chatcompletion_v2` 接口，在请求中开启 `speech_output` 选项，并设置 `voice_id`。API会以数据流（SSE）的形式同时返回文本（`content`）和音频数据（`audio_content`），前端可以实时接收并播放音频流，实现流畅对话。
 
<br>

## 使用示例


```python
import requests
import json
import subprocess
from typing import Iterator
from datetime import datetime
import logging

# 日志配置
logging.basicConfig(level=logging.INFO, format='[%(asctime)s] - [%(levelname)s]  %(message)s')

# 用于流式播放音频，需下载mpv播放器（适用 Linux/mac 系统）
mpv_command = ["mpv", "--no-cache", "--no-terminal", "--", "fd://0"]
mpv_process = subprocess.Popen(
    mpv_command,
    stdin=subprocess.PIPE,
    stdout=subprocess.DEVNULL,
    stderr=subprocess.DEVNULL,
)

token = '输入你在minimax开发者平台获取的token，前往 https://platform.minimaxi.com/user-center/basic-information/interface-key 获取'

def ccv2_audio_stream(text) -> Iterator[bytes] :
    payload = {
    "model": "MiniMax-Text-01",
    "messages": [
        {
        "role": "system",
        "name": "MM智能助理",
        "content": "MM智能助理是一款智能小助手"
        },
        {
        "role": "user",
        "name": "用户",
        "content": text
        },
    ],
    "stream": True,
    "tools": [
        {"type":"web_search"}
    ],
    "tool_choice": "auto",
    "max_tokens": 1024,
    "stream_options": { # 开启语音输出
        "speech_output": True
    },
    "voice_setting":{
        "model":"speech-01-turbo-240228",
        "voice_id":"female-tianmei"
    }
    }
    headers = {
    'Content-Type': 'application/json',
    'Authorization': f'Bearer {token}',
    }

    logging.info(f"【文本输入】{text}")

    response = requests.post("http://api.minimax.chat/v1/text/chatcompletion_v2", headers=headers, json=payload, stream=True)

    logging.info(f"Get response, trace-id: {response.headers.get('Trace-Id')}")
    i = 0
    for line in response.iter_lines(decode_unicode=True):
        if not line.startswith("data:"):
            continue
        i+=1
        logging.debug(f"[sse] data chunck-{i}")
        resp = json.loads(line.strip("data:"))
        if resp.get("choices") and resp["choices"][0].get("delta"):
            delta = resp["choices"][0]["delta"]
            if delta.get("role") == "assistant": # AI 助手回复
                if delta.get("content"):  logging.info(f"【文本输出】 {delta['content']}")   
                if delta.get("audio_content") and delta["audio_content"] != "": yield delta["audio_content"]
                if delta.get("tool_calls") : logging.info(f"【搜索中】...")   


# 流式播放音频并保存到本地
def audio_play(audio_stream: Iterator[bytes]) :
    audio = b""
    for chunk in audio_stream:
        if chunk is not None and chunk != '\n':
            decoded_hex = bytes.fromhex(chunk)
            mpv_process.stdin.write(decoded_hex)  # type: ignore
            mpv_process.stdin.flush()
            audio += decoded_hex

    if not audio:
        return

    now = datetime.now().strftime('%Y%m%d-%H%M%S')
    file_name = f'ccv2_audio_{now}.mp3'   
    with open(file_name, 'wb') as file:
        file.write(audio)
    logging.info(f"音频文件保存成功: {file_name}")

if __name__ == '__main__':
    audio_play(ccv2_audio_stream("请介绍自己"))

````

## 数字人·短视频·短剧
- 原始链接: https://platform.minimaxi.com/document/solutions_playlet?key=68ba75b4f7d629d07d27d738
- 文档ID: `68ba75b4f7d629d07d27d738`
# 数字人·短视频·短剧

**用 MiniMax Text & Speech API 让创作更高效、更生动**

小冰、闪剪智能、出门问问、蝉镜数字人、快手、逗哥配音等创作者与企业，正在利用 **MiniMax-Text-01** 与 ​**MiniMax-Speech-02**​，为数字人、短视频、短剧提供全流程 AI 创作能力。






## 解决方案

### 文案创作与翻译
1. 使用 **MiniMax-Text-01** 自动生成短视频台词、剧情对白、短剧脚本。
 1. 支持中英等多语种互译，轻松做国际化发行。

### 声音复刻与配音
1. 用 **MiniMax-Speech-02** 克隆演员、主播或品牌数字人的专属声音。
1. 实现超拟人 TTS 语音合成，情感、语调、节奏自然流畅。

### 多角色语音塑造​
1. 短剧中不同角色可分别绑定不同声音模型，实现多人对话自然切换。



  
<br>


## 业务价值

<container gap="4" columns="2">

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 显著降低成本
减少真人录音与配音次数，节约时间与预算

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 统一风格
保持数字人、短视频系列的声音与语气一致性

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 快速上线
台词生成 + 即时语音合成，创作周期从数天缩短到数小时

</card>

<card>
<img src="https://filecdn.minimax.chat/public/0a8a86b2-7642-453f-be41-3e9d84625db8.png" width="32px" />

### 国际化触达
支持多语言语音生成，助力出海

</card>


</container>



## 核心API能力

1.  声音复刻：用于角色声音复刻
<request method="POST" url="https://api.minimaxi.com/v1/voice_clone"></request>

1. 文本生成：生成台词和脚本
<request method="POST" url="https://api.minimaxi.com/v1/text/chatcompletion_v2"></request>

1. 语音合成：获取合成音频
<request method="POST" url="https://api.minimaxi.com/v1/t2a_v2"></request>
 

## 使用示例
1. 角色音色复刻

```python
import json

import requests

group_id = "your group id"
api_key = "your api key"

#复刻音频上传
url = f'https://api.minimaxi.com/v1/files/upload?GroupId={group_id}'
headers1 = {
    'authority': 'api.minimaxi.com',
    'Authorization': f'Bearer {api_key}'
}

data = {
    'purpose': 'voice_clone'
}

files = {
    'file': open('output.mp3', 'rb')
}
response = requests.post(url, headers=headers1, data=data, files=files)
file_id = response.json().get("file").get("file_id")
print(file_id)

#示例音频上传
url = f'https://api.minimaxi.com/v1/files/upload?GroupId={group_id}'
headers1 = {
    'authority': 'api.minimaxi.com',
    'Authorization': f'Bearer {api_key}'
}

data = {
    'purpose': 'prompt_audio'
}

files = {
    'file': open('prompt.mp3', 'rb')
}
response = requests.post(url, headers=headers1, data=data, files=files)
prompt_file_id = response.json().get("file").get("file_id")
print(prompt_file_id)


#音频复刻
url = f'https://api.minimaxi.com/v1/voice_clone?GroupId={group_id}'
payload2 = json.dumps({
  "file_id": file_id,
  "voice_id": "test1234"
})
headers2 = {
  'Authorization': f'Bearer {api_key}',
  'content-type': 'application/json'
}
response = requests.request("POST", url, headers=headers2, data=payload2)
print(response.text)

```

2. 生成台词与脚本
```python
import requests

api_key = "请填写您的api_key"

url = "https://api.minimaxi.com/v1/text/chatcompletion_v2"
headers = {
    "Authorization": f"Bearer {api_key}",
    "Content-Type": "application/json"
}

# 输入中文文本
chinese_text = input("请输入中文短剧文案：")

payload = {
    "model": "MiniMax-M1",
    "messages": [
        {
            "role": "system",
            "name": "Translator",
            "content": "你是一个擅长翻译的编剧。请将用户输入的中文翻译为英文，保持口语化、贴近短剧对白氛围，不能直译，要自然顺畅。"
        },
        {
            "role": "user",
            "name": "用户",
            "content": chinese_text
        }
    ]
}

response = requests.post(url, headers=headers, json=payload)

print("状态码:", response.status_code)
print("返回结果:")
print(response.text)
```

3.获取合成音频
```python
import requests
import json

group_id = "请填写您的group_id"
api_key = "请填写您的api_key"

url = f"https://api.minimaxi.com/v1/t2a_v2?GroupId={group_id}"

payload = json.dumps({
  "model":"speech-2.5-hd-preview",
  "text":"真正的危险不是计算机开始像人一样思考，而是人开始像计算机一样思考。计算机只是可以帮我们处理一些简单事务。",
  "stream":False,
  "voice_setting":{
    "voice_id":"male-qn-qingse",
    "speed":1,
    "vol":1,
    "pitch":0,
    "emotion":"happy"
  },
  "pronunciation_dict":{
    "tone":[
        "处理/(chu3)(li3)", "危险/dangerous"
    ]
  },
  "audio_setting":{
    "sample_rate":32000,
    "bitrate":128000,
    "format":"mp3",
    "channel":1
  }
})
headers = {
  'Authorization': f'Bearer {api_key}',
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, stream=True, headers=headers, data=payload)
parsed_json = json.loads(response.text)

# 获取audio字段的值
audio_value = bytes.fromhex(parsed_json['data']['audio'])
with open('output.mp3', 'wb') as f:
    f.write(audio_value)
```

## 语音模型测试与评估
- 原始链接: https://platform.minimaxi.com/document/speech_evaluate?key=68aec99b7189e62d342a654a
- 文档ID: `68aec99b7189e62d342a654a`
# 测试和评估标准-Speech模型
## 成功标准参考

通过设定明确、可行的目标，并结合客观数据与用户体验进行评估，能减少主观偏差，更好满足实际使用需求。以下是声音模型的参考标准，可根据实际情况调整和使用。

1. **任务完成度**
评估模型将输入文本准确转化为语音的能力，词错误率（WER）是其中一个评估指标。该指标的计算方法是使用自动语音识别技术（ASR） 将合成语音转为文本后，与参考文本进行比对，统计替换、插入、删除错误的比例。

2. **声音相似性**
评估合成语音与目标参考音频在说话人特征上的相似程度。相似性（SIM）的计算方法为提取语音的 Embedding 信息，计算合成音频与参考音频 Embedding 的余弦相似度。

3. **听感指标**
衡量用户对合成语音质量的感知，PESQ是一种客观听感评价指标，该指标将合成音频与一个高质量的“参考音频”进行比较，模拟人耳对音频的感知。

4. **可懂**​**度**
评估合成语音的可理解程度，STOI 是评估语音可懂度的客观指标之一，衡量听众能听懂单句内容的程度。

5. **主观选择比例**
反映用户对合成音频的真实感受，常用评估方式包括 **ELO 评分** 和 ​**比较平均意见分（CMOS）**。
* ​**ELO 评分**​：组织听众进行双音频对比测试（A/B 测试），每轮给出两段不同模型生成的音频，听众选择更偏好的版本。根据选择结果，使用 ELO 公式计算并更新各模型的评分，以反映其相对受欢迎程度。
* ​**比较平均意见分（CMOS）**：组织听众进行双音频对比测试（A/B 测试），让听众对两段音频样本的质量差异进行打分，将所有听众的评分取平均得到CMOS，可反映模型生成音频的相对优势程度。

6. **指令遵循**
衡量模型在生成语音时是否准确执行输入指令，包括情感标签控制、提示词中的音色控制等。评估该标准可以提取合成语音的特征与指令进行对比，也可以采用 A/B 测试等主观评估的方式。

7. **价格**
评估使用模型的预算，需要考虑每次调用模型的成本、使用频率等因素。

8. **响应速度**
衡量模型从输入文本到生成语音的时间效率。在流式生成时，首包延时是量化指标之一，衡量从接收到完整的用户输入开始，到系统输出**第一帧可播放音频数据包**的时间间隔。

## 定义评估场景

Speech 模型可以应用在多个不同的场景。为了全面评估模型性能，在进行测试和评估时需要考虑模型在不同应用场景的表现。

1. **声音克隆**

声音克隆是语音个性化生成的核心能力，在该场景下可验证模型在零样本或少样本条件下复刻目标说话人的音色、语调和说话风格的能力和稳定性。

2. **多语种生成**

语音模型支持多种语言的合成能力，本场景需要评估模型在不同语言上的表现是否均衡，能否准确捕捉和合成各语言的特有发音及韵律特征。

3. **跨语言合成**

该场景反映模型将同一说话人的声音迁移到其他语言的能力，验证模型在保持目标音色一致性的同时，能否生成流畅自然且符合目标语言口音的音频。

4. **情感控制**

情感表达是影响语音感染力和用户体验的因素，本场景评估模型能否根据指令准确合成指定情绪（如愤怒、开心、悲伤等）的音频，并保持语音的自然度和可懂性。

5. **文本驱动的音色合成**

除克隆已有声音外，模型还具备根据自然语言描述生成全新音色的能力。此场景验证模型是否能够根据文本提示中的描述，生成符合预期的个性化语音。

## 结果和测试用例示例

### MiniMax Speech-02模型客观评估结果

1. **声音克隆**

Speech模型能够在仅提供目标说话人短暂参考音频以及提供参考音频及对应文本信息的情况下，实现高质量的声音克隆。实验结果表明，Speech 模型在中文和英文场景下，WER 和 SIM 指标均表现良好，反映了良好的音色克隆效果和任务完成度。

<img src="https://filecdn.minimax.chat/public/468e4281-5413-48a7-aed3-bc12abb00a5c.PNG" alt="图片描述" width="60%" />

2. **多语种生成**

最新的Speech-2.5 模型支持40种语言，在声音相似度较高的情况下，模型也表现出良好的准确性。
参考下图可了解模型在部分语种下的表现能力。

<img src="https://filecdn.minimax.chat/public/857bd6d2-321a-4060-8f91-81feb193c066.PNG" alt="图片描述" width="60%" />

3. **跨语言合成**

MiniMax Speech-02 模型具有良好的跨语言合成能力，能够基于一小段语音片段合成其他语言的音频。以中文为源语言的实验结果表明，零样本克隆的准确率更高，而单样本克隆的声音相似度更好。

<img src="https://filecdn.minimax.chat/public/36cd4b73-2172-46f8-a9ef-c9cb68856fb2.png" alt="图片描述" width="60%" />

## MiniMax Speech-02测试用例

1. **声音克隆**

* 音频描述：吸引人且具有说服力的演讲者声音
  * 源音频 
<audio
  src="https://filecdn.minimax.chat/public/4b5cd287-3419-445f-b103-faaea6c602f7.wav"
  title="Marketing_Voice_Sourse"
  width="100%"
  subTitle="音频描述">
</audio>
  * 合成音频
  <audio
  src="https://filecdn.minimax.chat/public/56338b38-7421-43ba-b43b-0a72af35af10.wav"
  title="Compelling and Persuasive"
  width="100%"
  subTitle="音频描述">
</audio>



* 音频描述：具有低音共振和空间感的机器人声音
  * 源音频
<audio
  src="https://filecdn.minimax.chat/public/9036c98d-2d19-4e7f-9ae9-754234028bf4.wav"
  title="Lucky Robot_Sourse"
  width="100%"
  subTitle="音频描述">
</audio>

  * 合成音频
<audio
  src="https://filecdn.minimax.chat/public/c6a11b57-6fdc-4829-965d-f6a6304d61db.mp3"
  title="Lucky Robot"
  width="100%"
  subTitle="音频描述">
</audio>

2. **多语种生成**

* 泰语
  * 源音频
<audio
  src="https://filecdn.minimax.chat/public/1514f016-25b2-4298-8750-71ed862599a4.wav"
  title="Thai_Male_Sourse"
  width="100%"
  subTitle="音频描述">
</audio>


  * 合成音频
<audio
  src="https://filecdn.minimax.chat/public/8bb97bf0-10f9-4ffd-a942-991c9d155247.mp3"
  title="Thai"
  width="100%"
  subTitle="音频描述">
</audio>


* 越南语
  * 源音频
<audio
  src="https://filecdn.minimax.chat/public/f577f282-14a3-467f-8b9c-bf3e42adbcac.wav"
  title="Vietnamese_Female_Sourse"
  width="100%"
  subTitle="音频描述">
</audio>


  * 合成音频
<audio
  src="https://filecdn.minimax.chat/public/34ee500e-c063-48a1-953e-1f3ba3c33a43.mp3"
  title="Vietnamese"
  width="100%"
  subTitle="音频描述">
</audio>

3. **跨语言合成**

* 英语+汉语
  * 源音频（源音频为英语）
<audio
  src="https://filecdn.minimax.chat/public/7ad93750-82cf-4d01-8673-f8a390305906.mp3"
  title="Wong_Sourse"
  width="100%"
  subTitle="音频描述">
</audio>
  * 合成音频
<audio
  src="https://filecdn.minimax.chat/public/a339d81c-c18e-4f7d-ad4b-18781ce65bc6.wav"
  title="English-Mandarin"
  width="100%"
  subTitle="音频描述">
</audio>


* 日语+韩语
  * 源音频
<audio
  src="https://filecdn.minimax.chat/public/53d3ecf3-78cc-490f-a14c-079eea3423d4.mp3"
  title="Powerful_Girl_Sourse"
  width="100%"
  subTitle="音频描述">
</audio>
  * 合成音频
<audio
  src="https://filecdn.minimax.chat/public/c0c9f009-8be2-48be-ad24-0043f773a682.mp3"
  title="Japanese-Korean"
  width="100%"
  subTitle="音频描述">
</audio>

4. **情感控制**

* 惊讶
  * 源音频    
<audio
  src="https://filecdn.minimax.chat/public/bec5c5d6-bd31-4424-b38d-40d9c2a2cf10.mp3"
  title="Mandarin_Refreshing_Young_Man_Sourse"
  width="100%"
  subTitle="音频描述">
</audio>
  * 合成音频
<audio
  src="https://filecdn.minimax.chat/public/2042df8a-593f-4028-b7a8-0fc28bffd67b.mp3"
  title="Surprised"
  width="100%"
  subTitle="音频描述">
</audio>


* 快乐
  * 源音频
  <audio
  src="https://filecdn.minimax.chat/public/bec5c5d6-bd31-4424-b38d-40d9c2a2cf10.mp3"
  title="Mandarin_Refreshing_Young_Man_Sourse"
  width="100%"
  subTitle="音频描述">
  </audio>
  * 合成音频
  <audio
  src="https://filecdn.minimax.chat/public/5aa3266d-a364-4b4f-9440-4f6969da2bf1.mp3"
  title="Happy"
  width="100%"
  subTitle="音频描述">
</audio>

5. **文本控制的音色合成**

* 中文示例
  * 提示词
    男性中年声音，说中文，音色浑厚醇厚，带有自然的磁性，语速偏慢，音量适中，音调偏低沉。声音整体给人沉稳可靠的感觉，在深度访谈场景中表现出专业性和亲和力，音质清晰，吐字规整有力。
  * 合成音频
  <audio
  src="https://filecdn.minimax.chat/public/eb09755b-a169-4387-8b61-b65e190d2ee8.wav"
  title="深度访谈男中年"
  width="100%"
  subTitle="音频描述">
</audio>

* 英文示例
  * 提示词
    English-speaking female voice, sounding relatively young, with a sweet and pleasant tone. Speaking at a moderate pace with a touch of energy, similar to someone narrating a beauty/makeup tutorial video. The overall atmosphere is relaxed and cheerful.
  * 合成音频
  <audio
  src="https://filecdn.minimax.chat/public/f4933373-be46-4f43-9d52-5b95989e12cd.wav"
  title="美妆女博主"
  width="100%"
  subTitle="音频描述">
</audio>

更多测试用例可参考[技术报告](https://minimax-ai.github.io/tts_tech_report)，您也可以通过 [Audio 调试台](https://www.minimax.io/audio) 进行语音合成的实时调试和效果预览。

## 系统音色列表
- 原始链接: https://platform.minimaxi.com/document/system_voice_id?key=68b40964a96516e26018eee2
- 文档ID: `68b40964a96516e26018eee2`
# 系统音色列表

参考以下表格的内容，可查阅目前全部的系统音色。

| 序号 | 语言 | 音色 ID (Voice ID) | 音色名称 (Voice Name) |
| :--- | :--- | :--- | :--- |
| 1 | 中文 (普通话) | `male-qn-qingse` | 青涩青年音色 |
| 2 | 中文 (普通话) | `male-qn-jingying` | 精英青年音色 |
| 3 | 中文 (普通话) | `male-qn-badao` | 霸道青年音色 |
| 4 | 中文 (普通话) | `male-qn-daxuesheng` | 青年大学生音色 |
| 5 | 中文 (普通话) | `female-shaonv` | 少女音色 |
| 6 | 中文 (普通话) | `female-yujie` | 御姐音色 |
| 7 | 中文 (普通话) | `female-chengshu` | 成熟女性音色 |
| 8 | 中文 (普通话) | `female-tianmei` | 甜美女性音色 |
| 9 | 中文 (普通话) | `male-qn-qingse-jingpin` | 青涩青年音色-beta |
| 10 | 中文 (普通话) | `male-qn-jingying-jingpin` | 精英青年音色-beta |
| 11 | 中文 (普通话) | `male-qn-badao-jingpin` | 霸道青年音色-beta |
| 12 | 中文 (普通话) | `male-qn-daxuesheng-jingpin` | 青年大学生音色-beta |
| 13 | 中文 (普通话) | `female-shaonv-jingpin` | 少女音色-beta |
| 14 | 中文 (普通话) | `female-yujie-jingpin` | 御姐音色-beta |
| 15 | 中文 (普通话) | `female-chengshu-jingpin` | 成熟女性音色-beta |
| 16 | 中文 (普通话) | `female-tianmei-jingpin` | 甜美女性音色-beta |
| 17 | 中文 (普通话) | `clever_boy` | 聪明男童 |
| 18 | 中文 (普通话) | `cute_boy` | 可爱男童 |
| 19 | 中文 (普通话) | `lovely_girl` | 萌萌女童 |
| 20 | 中文 (普通话) | `cartoon_pig` | 卡通猪小琪 |
| 21 | 中文 (普通话) | `bingjiao_didi` | 病娇弟弟 |
| 22 | 中文 (普通话) | `junlang_nanyou` | 俊朗男友 |
| 23 | 中文 (普通话) | `chunzhen_xuedi` | 纯真学弟 |
| 24 | 中文 (普通话) | `lengdan_xiongzhang` | 冷淡学长 |
| 25 | 中文 (普通话) | `badao_shaoye` | 霸道少爷 |
| 26 | 中文 (普通话) | `tianxin_xiaoling` | 甜心小玲 |
| 27 | 中文 (普通话) | `qiaopi_mengmei` | 俏皮萌妹 |
| 28 | 中文 (普通话) | `wumei_yujie` | 妩媚御姐 |
| 29 | 中文 (普通话) | `diadia_xuemei` | 嗲嗲学妹 |
| 30 | 中文 (普通话) | `danya_xuejie` | 淡雅学姐 |
| 31 | 中文 (普通话) | `Chinese (Mandarin)_Reliable_Executive` | 沉稳高管 |
| 32 | 中文 (普通话) | `Chinese (Mandarin)_News_Anchor` | 新闻女声 |
| 33 | 中文 (普通话) | `Chinese (Mandarin)_Mature_Woman` | 傲娇御姐 |
| 34 | 中文 (普通话) | `Chinese (Mandarin)_Unrestrained_Young_Man` | 不羁青年 |
| 35 | 中文 (普通话) | `Arrogant_Miss` | 嚣张小姐 |
| 36 | 中文 (普通话) | `Robot_Armor` | 机械战甲 |
| 37 | 中文 (普通话) | `Chinese (Mandarin)_Kind-hearted_Antie` | 热心大婶 |
| 38 | 中文 (普通话) | `Chinese (Mandarin)_HK_Flight_Attendant` | 港普空姐 |
| 39 | 中文 (普通话) | `Chinese (Mandarin)_Humorous_Elder` | 搞笑大爷 |
| 40 | 中文 (普通话) | `Chinese (Mandarin)_Gentleman` | 温润男声 |
| 41 | 中文 (普通话) | `Chinese (Mandarin)_Warm_Bestie` | 温暖闺蜜 |
| 42 | 中文 (普通话) | `Chinese (Mandarin)_Male_Announcer` | 播报男声 |
| 43 | 中文 (普通话) | `Chinese (Mandarin)_Sweet_Lady` | 甜美女声 |
| 44 | 中文 (普通话) | `Chinese (Mandarin)_Southern_Young_Man` | 南方小哥 |
| 45 | 中文 (普通话) | `Chinese (Mandarin)_Wise_Women` | 阅历姐姐 |
| 46 | 中文 (普通话) | `Chinese (Mandarin)_Gentle_Youth` | 温润青年 |
| 47 | 中文 (普通话) | `Chinese (Mandarin)_Warm_Girl` | 温暖少女 |
| 48 | 中文 (普通话) | `Chinese (Mandarin)_Kind-hearted_Elder` | 花甲奶奶 |
| 49 | 中文 (普通话) | `Chinese (Mandarin)_Cute_Spirit` | 憨憨萌兽 |
| 50 | 中文 (普通话) | `Chinese (Mandarin)_Radio_Host` | 电台男主播 |
| 51 | 中文 (普通话) | `Chinese (Mandarin)_Lyrical_Voice` | 抒情男声 |
| 52 | 中文 (普通话) | `Chinese (Mandarin)_Straightforward_Boy` | 率真弟弟 |
| 53 | 中文 (普通话) | `Chinese (Mandarin)_Sincere_Adult` | 真诚青年 |
| 54 | 中文 (普通话) | `Chinese (Mandarin)_Gentle_Senior` | 温柔学姐 |
| 55 | 中文 (普通话) | `Chinese (Mandarin)_Stubborn_Friend` | 嘴硬竹马 |
| 56 | 中文 (普通话) | `Chinese (Mandarin)_Crisp_Girl` | 清脆少女 |
| 57 | 中文 (普通话) | `Chinese (Mandarin)_Pure-hearted_Boy` | 清澈邻家弟弟 |
| 58 | 中文 (普通话) | `Chinese (Mandarin)_Soft_Girl` | 软软女孩 |
| 59 | 中文 (粤语) | `Cantonese_ProfessionalHost（F)` | 专业女主持 |
| 60 | 中文 (粤语) | `Cantonese_GentleLady` | 温柔女声 |
| 61 | 中文 (粤语) | `Cantonese_ProfessionalHost（M)` | 专业男主持 |
| 62 | 中文 (粤语) | `Cantonese_PlayfulMan` | 活泼男声 |
| 63 | 中文 (粤语) | `Cantonese_CuteGirl` | 可爱女孩 |
| 64 | 中文 (粤语) | `Cantonese_KindWoman` | 善良女声 |
| 65 | 英文 | `Santa_Claus ` | Santa Claus |
| 66 | 英文 | `Grinch` | Grinch |
| 67 | 英文 | `Rudolph` | Rudolph |
| 68 | 英文 | `Arnold` | Arnold |
| 69 | 英文 | `Charming_Santa` | Charming Santa |
| 70 | 英文 | `Charming_Lady` | Charming Lady |
| 71 | 英文 | `Sweet_Girl` | Sweet Girl |
| 72 | 英文 | `Cute_Elf` | Cute Elf |
| 73 | 英文 | `Attractive_Girl` | Attractive Girl |
| 74 | 英文 | `Serene_Woman` | Serene Woman |
| 75 | 英文 | `English_Trustworthy_Man` | Trustworthy Man |
| 76 | 英文 | `English_Graceful_Lady` | Graceful Lady |
| 77 | 英文 | `English_Aussie_Bloke` | Aussie Bloke |
| 78 | 英文 | `English_Whispering_girl` | Whispering girl |
| 79 | 英文 | `English_Diligent_Man` | Diligent Man |
| 80 | 英文 | `English_Gentle-voiced_man` | Gentle-voiced man |
| 81 | 日文 | `Japanese_IntellectualSenior` | Intellectual Senior |
| 82 | 日文 | `Japanese_DecisivePrincess` | Decisive Princess |
| 83 | 日文 | `Japanese_LoyalKnight` | Loyal Knight |
| 84 | 日文 | `Japanese_DominantMan` | Dominant Man |
| 85 | 日文 | `Japanese_SeriousCommander` | Serious Commander |
| 86 | 日文 | `Japanese_ColdQueen` | Cold Queen |
| 87 | 日文 | `Japanese_DependableWoman` | Dependable Woman |
| 88 | 日文 | `Japanese_GentleButler` | Gentle Butler |
| 89 | 日文 | `Japanese_KindLady` | Kind Lady |
| 90 | 日文 | `Japanese_CalmLady` | Calm Lady |
| 91 | 日文 | `Japanese_OptimisticYouth` | Optimistic Youth |
| 92 | 日文 | `Japanese_GenerousIzakayaOwner` | Generous Izakaya Owner |
| 93 | 日文 | `Japanese_SportyStudent` | Sporty Student |
| 94 | 日文 | `Japanese_InnocentBoy` | Innocent Boy |
| 95 | 日文 | `Japanese_GracefulMaiden` | Graceful Maiden |
| 96 | 韩文 | `Korean_SweetGirl` | Sweet Girl |
| 97 | 韩文 | `Korean_CheerfulBoyfriend` | Cheerful Boyfriend |
| 98 | 韩文 | `Korean_EnchantingSister` | Enchanting Sister |
| 99 | 韩文 | `Korean_ShyGirl` | Shy Girl |
| 100 | 韩文 | `Korean_ReliableSister` | Reliable Sister |
| 101 | 韩文 | `Korean_StrictBoss` | Strict Boss |
| 102 | 韩文 | `Korean_SassyGirl` | Sassy Girl |
| 103 | 韩文 | `Korean_ChildhoodFriendGirl` | Childhood Friend Girl |
| 104 | 韩文 | `Korean_PlayboyCharmer` | Playboy Charmer |
| 105 | 韩文 | `Korean_ElegantPrincess` | Elegant Princess |
| 106 | 韩文 | `Korean_BraveFemaleWarrior` | Brave Female Warrior |
| 107 | 韩文 | `Korean_BraveYouth` | Brave Youth |
| 108 | 韩文 | `Korean_CalmLady` | Calm Lady |
| 109 | 韩文 | `Korean_EnthusiasticTeen` | Enthusiastic Teen |
| 110 | 韩文 | `Korean_SoothingLady` | Soothing Lady |
| 111 | 韩文 | `Korean_IntellectualSenior` | Intellectual Senior |
| 112 | 韩文 | `Korean_LonelyWarrior` | Lonely Warrior |
| 113 | 韩文 | `Korean_MatureLady` | Mature Lady |
| 114 | 韩文 | `Korean_InnocentBoy` | Innocent Boy |
| 115 | 韩文 | `Korean_CharmingSister` | Charming Sister |
| 116 | 韩文 | `Korean_AthleticStudent` | Athletic Student |
| 117 | 韩文 | `Korean_BraveAdventurer` | Brave Adventurer |
| 118 | 韩文 | `Korean_CalmGentleman` | Calm Gentleman |
| 119 | 韩文 | `Korean_WiseElf` | Wise Elf |
| 120 | 韩文 | `Korean_CheerfulCoolJunior` | Cheerful Cool Junior |
| 121 | 韩文 | `Korean_DecisiveQueen` | Decisive Queen |
| 122 | 韩文 | `Korean_ColdYoungMan` | Cold Young Man |
| 123 | 韩文 | `Korean_MysteriousGirl` | Mysterious Girl |
| 124 | 韩文 | `Korean_QuirkyGirl` | Quirky Girl |
| 125 | 韩文 | `Korean_ConsiderateSenior` | Considerate Senior |
| 126 | 韩文 | `Korean_CheerfulLittleSister` | Cheerful Little Sister |
| 127 | 韩文 | `Korean_DominantMan` | Dominant Man |
| 128 | 韩文 | `Korean_AirheadedGirl` | Airheaded Girl |
| 129 | 韩文 | `Korean_ReliableYouth` | Reliable Youth |
| 130 | 韩文 | `Korean_FriendlyBigSister` | Friendly Big Sister |
| 131 | 韩文 | `Korean_GentleBoss` | Gentle Boss |
| 132 | 韩文 | `Korean_ColdGirl` | Cold Girl |
| 133 | 韩文 | `Korean_HaughtyLady` | Haughty Lady |
| 134 | 韩文 | `Korean_CharmingElderSister` | Charming Elder Sister |
| 135 | 韩文 | `Korean_IntellectualMan` | Intellectual Man |
| 136 | 韩文 | `Korean_CaringWoman` | Caring Woman |
| 137 | 韩文 | `Korean_WiseTeacher` | Wise Teacher |
| 138 | 韩文 | `Korean_ConfidentBoss` | Confident Boss |
| 139 | 韩文 | `Korean_AthleticGirl` | Athletic Girl |
| 140 | 韩文 | `Korean_PossessiveMan` | Possessive Man |
| 141 | 韩文 | `Korean_GentleWoman` | Gentle Woman |
| 142 | 韩文 | `Korean_CockyGuy` | Cocky Guy |
| 143 | 韩文 | `Korean_ThoughtfulWoman` | Thoughtful Woman |
| 144 | 韩文 | `Korean_OptimisticYouth` | Optimistic Youth |
| 145 | 西班牙文 | `Spanish_SereneWoman` | Serene Woman |
| 146 | 西班牙文 | `Spanish_MaturePartner` | Mature Partner |
| 147 | 西班牙文 | `Spanish_CaptivatingStoryteller` | Captivating Storyteller |
| 148 | 西班牙文 | `Spanish_Narrator` | Narrator |
| 149 | 西班牙文 | `Spanish_WiseScholar` | Wise Scholar |
| 150 | 西班牙文 | `Spanish_Kind-heartedGirl` | Kind-hearted Girl |
| 151 | 西班牙文 | `Spanish_DeterminedManager` | Determined Manager |
| 152 | 西班牙文 | `Spanish_BossyLeader` | Bossy Leader |
| 153 | 西班牙文 | `Spanish_ReservedYoungMan` | Reserved Young Man |
| 154 | 西班牙文 | `Spanish_ConfidentWoman` | Confident Woman |
| 155 | 西班牙文 | `Spanish_ThoughtfulMan` | Thoughtful Man |
| 156 | 西班牙文 | `Spanish_Strong-WilledBoy` | Strong-willed Boy |
| 157 | 西班牙文 | `Spanish_SophisticatedLady` | Sophisticated Lady |
| 158 | 西班牙文 | `Spanish_RationalMan` | Rational Man |
| 159 | 西班牙文 | `Spanish_AnimeCharacter` | Anime Character |
| 160 | 西班牙文 | `Spanish_Deep-tonedMan` | Deep-toned Man |
| 161 | 西班牙文 | `Spanish_Fussyhostess` | Fussy hostess |
| 162 | 西班牙文 | `Spanish_SincereTeen` | Sincere Teen |
| 163 | 西班牙文 | `Spanish_FrankLady` | Frank Lady |
| 164 | 西班牙文 | `Spanish_Comedian` | Comedian |
| 165 | 西班牙文 | `Spanish_Debator` | Debator |
| 166 | 西班牙文 | `Spanish_ToughBoss` | Tough Boss |
| 167 | 西班牙文 | `Spanish_Wiselady` | Wise Lady |
| 168 | 西班牙文 | `Spanish_Steadymentor` | Steady Mentor |
| 169 | 西班牙文 | `Spanish_Jovialman` | Jovial Man |
| 170 | 西班牙文 | `Spanish_SantaClaus` | Santa Claus |
| 171 | 西班牙文 | `Spanish_Rudolph` | Rudolph |
| 172 | 西班牙文 | `Spanish_Intonategirl` | Intonate Girl |
| 173 | 西班牙文 | `Spanish_Arnold` | Arnold |
| 174 | 西班牙文 | `Spanish_Ghost` | Ghost |
| 175 | 西班牙文 | `Spanish_HumorousElder` | Humorous Elder |
| 176 | 西班牙文 | `Spanish_EnergeticBoy` | Energetic Boy |
| 177 | 西班牙文 | `Spanish_WhimsicalGirl` | Whimsical Girl |
| 178 | 西班牙文 | `Spanish_StrictBoss` | Strict Boss |
| 179 | 西班牙文 | `Spanish_ReliableMan` | Reliable Man |
| 180 | 西班牙文 | `Spanish_SereneElder` | Serene Elder |
| 181 | 西班牙文 | `Spanish_AngryMan` | Angry Man |
| 182 | 西班牙文 | `Spanish_AssertiveQueen` | Assertive Queen |
| 183 | 西班牙文 | `Spanish_CaringGirlfriend` | Caring Girlfriend |
| 184 | 西班牙文 | `Spanish_PowerfulSoldier` | Powerful Soldier |
| 185 | 西班牙文 | `Spanish_PassionateWarrior` | Passionate Warrior |
| 186 | 西班牙文 | `Spanish_ChattyGirl` | Chatty Girl |
| 187 | 西班牙文 | `Spanish_RomanticHusband` | Romantic Husband |
| 188 | 西班牙文 | `Spanish_CompellingGirl` | Compelling Girl |
| 189 | 西班牙文 | `Spanish_PowerfulVeteran` | Powerful Veteran |
| 190 | 西班牙文 | `Spanish_SensibleManager` | Sensible Manager |
| 191 | 西班牙文 | `Spanish_ThoughtfulLady` | Thoughtful Lady |
| 192 | 葡萄牙文 | `Portuguese_SentimentalLady` | Sentimental Lady |
| 193 | 葡萄牙文 | `Portuguese_BossyLeader` | Bossy Leader |
| 194 | 葡萄牙文 | `Portuguese_Wiselady` | Wise lady |
| 195 | 葡萄牙文 | `Portuguese_Strong-WilledBoy` | Strong-willed Boy |
| 196 | 葡萄牙文 | `Portuguese_Deep-VoicedGentleman` | Deep-voiced Gentleman |
| 197 | 葡萄牙文 | `Portuguese_UpsetGirl` | Upset Girl |
| 198 | 葡萄牙文 | `Portuguese_PassionateWarrior` | Passionate Warrior |
| 199 | 葡萄牙文 | `Portuguese_AnimeCharacter` | Anime Character |
| 200 | 葡萄牙文 | `Portuguese_ConfidentWoman` | Confident Woman |
| 201 | 葡萄牙文 | `Portuguese_AngryMan` | Angry Man |
| 202 | 葡萄牙文 | `Portuguese_CaptivatingStoryteller` | Captivating Storyteller |
| 203 | 葡萄牙文 | `Portuguese_Godfather` | Godfather |
| 204 | 葡萄牙文 | `Portuguese_ReservedYoungMan` | Reserved Young Man |
| 205 | 葡萄牙文 | `Portuguese_SmartYoungGirl` | Smart Young Girl |
| 206 | 葡萄牙文 | `Portuguese_Kind-heartedGirl` | Kind-hearted Girl |
| 207 | 葡萄牙文 | `Portuguese_Pompouslady` | Pompous lady |
| 208 | 葡萄牙文 | `Portuguese_Grinch` | Grinch |
| 209 | 葡萄牙文 | `Portuguese_Debator` | Debator |
| 210 | 葡萄牙文 | `Portuguese_SweetGirl` | Sweet Girl |
| 211 | 葡萄牙文 | `Portuguese_AttractiveGirl` | Attractive Girl |
| 212 | 葡萄牙文 | `Portuguese_ThoughtfulMan` | Thoughtful Man |
| 213 | 葡萄牙文 | `Portuguese_PlayfulGirl` | Playful Girl |
| 214 | 葡萄牙文 | `Portuguese_GorgeousLady` | Gorgeous Lady |
| 215 | 葡萄牙文 | `Portuguese_LovelyLady` | Lovely Lady |
| 216 | 葡萄牙文 | `Portuguese_SereneWoman` | Serene Woman |
| 217 | 葡萄牙文 | `Portuguese_SadTeen` | Sad Teen |
| 218 | 葡萄牙文 | `Portuguese_MaturePartner` | Mature Partner |
| 219 | 葡萄牙文 | `Portuguese_Comedian` | Comedian |
| 220 | 葡萄牙文 | `Portuguese_NaughtySchoolgirl` | Naughty Schoolgirl |
| 221 | 葡萄牙文 | `Portuguese_Narrator` | Narrator |
| 222 | 葡萄牙文 | `Portuguese_ToughBoss` | Tough Boss |
| 223 | 葡萄牙文 | `Portuguese_Fussyhostess` | Fussy hostess |
| 224 | 葡萄牙文 | `Portuguese_Dramatist` | Dramatist |
| 225 | 葡萄牙文 | `Portuguese_Steadymentor` | Steady Mentor |
| 226 | 葡萄牙文 | `Portuguese_Jovialman` | Jovial Man |
| 227 | 葡萄牙文 | `Portuguese_CharmingQueen` | Charming Queen |
| 228 | 葡萄牙文 | `Portuguese_SantaClaus` | Santa Claus |
| 229 | 葡萄牙文 | `Portuguese_Rudolph` | Rudolph |
| 230 | 葡萄牙文 | `Portuguese_Arnold` | Arnold |
| 231 | 葡萄牙文 | `Portuguese_CharmingSanta` | Charming Santa |
| 232 | 葡萄牙文 | `Portuguese_CharmingLady` | Charming Lady |
| 233 | 葡萄牙文 | `Portuguese_Ghost` | Ghost |
| 234 | 葡萄牙文 | `Portuguese_HumorousElder` | Humorous Elder |
| 235 | 葡萄牙文 | `Portuguese_CalmLeader` | Calm Leader |
| 236 | 葡萄牙文 | `Portuguese_GentleTeacher` | Gentle Teacher |
| 237 | 葡萄牙文 | `Portuguese_EnergeticBoy` | Energetic Boy |
| 238 | 葡萄牙文 | `Portuguese_ReliableMan` | Reliable Man |
| 239 | 葡萄牙文 | `Portuguese_SereneElder` | Serene Elder |
| 240 | 葡萄牙文 | `Portuguese_GrimReaper` | Grim Reaper |
| 241 | 葡萄牙文 | `Portuguese_AssertiveQueen` | Assertive Queen |
| 242 | 葡萄牙文 | `Portuguese_WhimsicalGirl` | Whimsical Girl |
| 243 | 葡萄牙文 | `Portuguese_StressedLady` | Stressed Lady |
| 244 | 葡萄牙文 | `Portuguese_FriendlyNeighbor` | Friendly Neighbor |
| 245 | 葡萄牙文 | `Portuguese_CaringGirlfriend` | Caring Girlfriend |
| 246 | 葡萄牙文 | `Portuguese_PowerfulSoldier` | Powerful Soldier |
| 247 | 葡萄牙文 | `Portuguese_FascinatingBoy` | Fascinating Boy |
| 248 | 葡萄牙文 | `Portuguese_RomanticHusband` | Romantic Husband |
| 249 | 葡萄牙文 | `Portuguese_StrictBoss` | Strict Boss |
| 250 | 葡萄牙文 | `Portuguese_InspiringLady` | Inspiring Lady |
| 251 | 葡萄牙文 | `Portuguese_PlayfulSpirit` | Playful Spirit |
| 252 | 葡萄牙文 | `Portuguese_ElegantGirl` | Elegant Girl |
| 253 | 葡萄牙文 | `Portuguese_CompellingGirl` | Compelling Girl |
| 254 | 葡萄牙文 | `Portuguese_PowerfulVeteran` | Powerful Veteran |
| 255 | 葡萄牙文 | `Portuguese_SensibleManager` | Sensible Manager |
| 256 | 葡萄牙文 | `Portuguese_ThoughtfulLady` | Thoughtful Lady |
| 257 | 葡萄牙文 | `Portuguese_TheatricalActor` | Theatrical Actor |
| 258 | 葡萄牙文 | `Portuguese_FragileBoy` | Fragile Boy |
| 259 | 葡萄牙文 | `Portuguese_ChattyGirl` | Chatty Girl |
| 260 | 葡萄牙文 | `Portuguese_Conscientiousinstructor` | Conscientious Instructor |
| 261 | 葡萄牙文 | `Portuguese_RationalMan` | Rational Man |
| 262 | 葡萄牙文 | `Portuguese_WiseScholar` | Wise Scholar |
| 263 | 葡萄牙文 | `Portuguese_FrankLady` | Frank Lady |
| 264 | 葡萄牙文 | `Portuguese_DeterminedManager` | Determined Manager |
| 265 | 法文 | `French_Male_Speech_New` | Level-Headed Man |
| 266 | 法文 | `French_Female_News Anchor` | Patient Female Presenter |
| 267 | 法文 | `French_CasualMan` | Casual Man |
| 268 | 法文 | `French_MovieLeadFemale` | Movie Lead Female |
| 269 | 法文 | `French_FemaleAnchor` | Female Anchor |
| 270 | 法文 | `French_MaleNarrator` | Male Narrator |
| 271 | 印尼文 | `Indonesian_SweetGirl` | Sweet Girl |
| 272 | 印尼文 | `Indonesian_ReservedYoungMan` | Reserved Young Man |
| 273 | 印尼文 | `Indonesian_CharmingGirl` | Charming Girl |
| 274 | 印尼文 | `Indonesian_CalmWoman` | Calm Woman |
| 275 | 印尼文 | `Indonesian_ConfidentWoman` | Confident Woman |
| 276 | 印尼文 | `Indonesian_CaringMan` | Caring Man |
| 277 | 印尼文 | `Indonesian_BossyLeader` | Bossy Leader |
| 278 | 印尼文 | `Indonesian_DeterminedBoy` | Determined Boy |
| 279 | 印尼文 | `Indonesian_GentleGirl` | Gentle Girl |
| 280 | 德文 | `German_FriendlyMan` | Friendly Man |
| 281 | 德文 | `German_SweetLady` | Sweet Lady |
| 282 | 德文 | `German_PlayfulMan` | Playful Man |
| 283 | 俄文 | `Russian_HandsomeChildhoodFriend` | Handsome Childhood Friend |
| 284 | 俄文 | `Russian_BrightHeroine` | Bright Queen |
| 285 | 俄文 | `Russian_AmbitiousWoman` | Ambitious Woman |
| 286 | 俄文 | `Russian_ReliableMan` | Reliable Man |
| 287 | 俄文 | `Russian_CrazyQueen` | Crazy Girl |
| 288 | 俄文 | `Russian_PessimisticGirl` | Pessimistic Girl |
| 289 | 俄文 | `Russian_AttractiveGuy` | Attractive Guy |
| 290 | 俄文 | `Russian_Bad-temperedBoy` | Bad-tempered Boy |
| 291 | 意大利文 | `Italian_BraveHeroine` | Brave Heroine |
| 292 | 意大利文 | `Italian_Narrator` | Narrator |
| 293 | 意大利文 | `Italian_WanderingSorcerer` | Wandering Sorcerer |
| 294 | 意大利文 | `Italian_DiligentLeader` | Diligent Leader |
| 295 | 阿拉伯文 | `Arabic_CalmWoman` | Calm Woman |
| 296 | 阿拉伯文 | `Arabic_FriendlyGuy` | Friendly Guy |
| 297 | 土耳其文 | `Turkish_CalmWoman` | Calm Woman |
| 298 | 土耳其文 | `Turkish_Trustworthyman` | Trustworthy man |
| 299 | 乌克兰文 | `Ukrainian_CalmWoman` | Calm Woman |
| 300 | 乌克兰文 | `Ukrainian_WiseScholar` | Wise Scholar |
| 301 | 荷兰文 | `Dutch_kindhearted_girl` | Kind-hearted girl |
| 302 | 荷兰文 | `Dutch_bossy_leader` | Bossy leader |
| 303 | 越南文 | `Vietnamese_kindhearted_girl` | Kind-hearted girl |

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/t2a_api_intro?key=68ad77666602726333000457
- 文档ID: `68ad77666602726333000457`
# 同步语音合成（T2A）
该 API 支持基于文本到语音的同步生成，单次可处理最长 10,000 字符的文本。接口本身为无状态接口，即单次调用时，模型仅处理单次传入内容，不涉及业务逻辑，同时模型也不存储您传入的数据。
该接口支持以下功能：
1. 支持 300+ 系统音色、复刻音色自主选择；
2. 支持音量、语调、语速、输出格式调整；
3. 支持按比例混音功能；
4. 支持固定间隔时间控制；
5. 支持多种音频规格、格式，包括：mp3, pcm, flac, wav。注：wav 仅在非流式输出下支持；
6. 支持流式输出。

该接口的适用场景：短句生成、语音聊天、在线社交等。
## 支持模型
以下为 MiniMax 提供的语音模型及其特性说明。
| 模型 | 特性 |
| ------ | ------ |
| speech-2.5-hd-preview | 最新的 HD 模型，韵律表现出色，复刻相似度极佳 |
| speech-2.5-turbo-preview | 最新的 Turbo 模型，支持 40 个语种 |
| speech-02-hd | 拥有出色的韵律、稳定性和复刻相似度，音质表现突出 |
| speech-02-turbo | 拥有出色的韵律和稳定性，小语种能力加强，性能表现出色 |
| speech-01-hd | 拥有超高的复刻相似度，音质表现突出 |
| speech-01-turbo | 在出色的生成效果基础上有更快的生成速度 |

## 接口说明
同步语音合成功能，共包含2个接口，可根据需求，选择使用。
- HTTP同步语音合成
- WebSocket同步语音合成

## 支持语言
MiniMax 的语音合成模型具备卓越的跨语言能力，全面支持 40 种全球广泛使用的语言。我们致力于打破语言壁垒，构建真正意义上的全球通用人工智能模型。

目前支持的语言包含：

| 序号 | 中文 | EN |
|---|---|---|
| 1 | 中文 | Chinese |
| 2 | 粤语 | Cantonese |
| 3 | 英语 | English |
| 4 | 西班牙语 | Spanish |
| 5 | 法语 | French |
| 6 | 俄语 | Russian |
| 7 | 德语 | German |
| 8 | 葡萄牙语 | Portuguese |
| 9 | 阿拉伯语 | Arabic |
| 10 | 意大利语 | Italian |
| 11 | 日语 | Japanese |
| 12 | 韩语 | Korean |
| 13 | 印尼语 | Indonesian |
| 14 | 越南语 | Vietnamese |
| 15 | 土耳其语 | Turkish |
| 16 | 荷兰语 | Dutch |
| 17 | 乌克兰语 | Ukrainian |
| 18 | 泰语 | Thai |
| 19 | 波兰语 | Polish |
| 20 | 罗马尼亚语 | Romanian |
| 21 | 希腊语 | Greek |
| 22 | 捷克语 | Czech |
| 23 | 芬兰语 | Finnish |
| 24 | 印地语 | Hindi |
| 25 | 保加利亚语 | Bulgarian |
| 26 | 丹麦语 | Danish |
| 27 | 希伯来语 | Hebrew |
| 28 | 马来语 | Malay |
| 29 | 波斯语 | Persian |
| 30 | 斯洛伐克语 | Slovak |
| 31 | 瑞典语 | Swedish |
| 32 | 克罗地亚语 | Croatian |
| 33 | 菲律宾语 | Filipino |
| 34 | 匈牙利语 | Hungarian |
| 35 | 挪威语 | Norwegian |
| 36 | 斯洛文尼亚语 | Slovenian |
| 37 | 加泰罗尼亚语 | Catalan |
| 38 | 尼诺斯克语 | Nynorsk |
| 39 | 泰米尔语 | Tamil |
| 40 | 阿非利卡语 | Afrikaans |

## 官方MCP
MiniMax 提供官方的 [Python 版本](https://github.com/MiniMax-AI/MiniMax-MCP) 和 [JavaScript 版本](https://github.com/MiniMax-AI/MiniMax-MCP-JS) 模型上下文协议（MCP）服务器实现代码，支持语音合成功能，详细说明请参考 [MiniMax MCP 使用指南文档](/document/guides_mcp?key=68bbf0f295a600dbfc926bdd)

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/t2a_async_api_intro?key=68adac446fe587e3fbfe965b
- 文档ID: `68adac446fe587e3fbfe965b`
# 异步长文本语音生成（T2A Async）
该API支持基于文本到语音的异步生成，单次文本生成传输最大支持100万字符，生成的完整音频结果支持异步的方式进行检索。
该接口支持以下功能：
1. 支持100+系统音色、复刻音色自主选择；
1. 支持语调、语速、音量、比特率、采样率、输出格式自主调整；
1. 支持音频时长、音频大小等返回参数；
1. 支持时间戳（字幕）返回，精确到句；
1. 支持直接传入字符串与上传文本文件file_id两种方式进行待合成文本的输入；
1. 支持非法字符检测：非法字符不超过10%（包含10%），音频会正常生成并返回非法字符占比；非法字符超过10%，接口不返回结果（返回报错码），请检测后再次进行请求【非法字符定义：ascii码中的控制符（不含制表符和换行符）】。

提交长文本语音合成请求后，会生成file_id，生成任务完成后，可通过file_id使用文件检索接口进行下载。

⚠️ 注意：返回的url的有效期为：自url返回开始的**9个小时**（即32400秒），超过有效期后url便会失效，生成的信息便会丢失，请注意下载信息的时间。

**适用场景：整本书籍等长文本的语音生成。**

## 支持模型
以下为 MiniMax 提供的语音模型及其特性说明。
| 模型 | 特性 |
| ------ | ------ |
| speech-2.5-hd-preview | 最新的 HD 模型，韵律表现出色，复刻相似度极佳 |
| speech-2.5-turbo-preview | 最新的 Turbo 模型，支持 40 个语种 |
| speech-02-hd | 拥有出色的韵律、稳定性和复刻相似度，音质表现突出 |
| speech-02-turbo | 拥有出色的韵律和稳定性，小语种能力加强，性能表现出色 |
| speech-01-hd | 拥有超高的复刻相似度，音质表现突出 |
| speech-01-turbo | 在出色的生成效果基础上有更快的生成速度 |

## 接口说明
整体包含2个API：创建**语音生成任务**、**查询语音生成任务状态**。使用步骤如下：
1. 创建语音生成任务得到task_id（如果选择以file_id的形式传入待合成文本，需要前置使用File(Upload)接口进行文件上传）；
1. 基于taskid查询语音生成任务状态；
1. 如果发现任务生成成功，那么可以使用本接口返回的 file_id 通过 File API 进行结果查看和下载。


## 支持语言
MiniMax 的语音合成模型具备卓越的跨语言能力，全面支持 40 种全球广泛使用的语言。我们致力于打破语言壁垒，构建真正意义上的全球通用人工智能模型。

目前支持的语言包含：

| 序号 | 中文 | EN |
|---|---|---|
| 1 | 中文 | Chinese |
| 2 | 粤语 | Cantonese |
| 3 | 英语 | English |
| 4 | 西班牙语 | Spanish |
| 5 | 法语 | French |
| 6 | 俄语 | Russian |
| 7 | 德语 | German |
| 8 | 葡萄牙语 | Portuguese |
| 9 | 阿拉伯语 | Arabic |
| 10 | 意大利语 | Italian |
| 11 | 日语 | Japanese |
| 12 | 韩语 | Korean |
| 13 | 印尼语 | Indonesian |
| 14 | 越南语 | Vietnamese |
| 15 | 土耳其语 | Turkish |
| 16 | 荷兰语 | Dutch |
| 17 | 乌克兰语 | Ukrainian |
| 18 | 泰语 | Thai |
| 19 | 波兰语 | Polish |
| 20 | 罗马尼亚语 | Romanian |
| 21 | 希腊语 | Greek |
| 22 | 捷克语 | Czech |
| 23 | 芬兰语 | Finnish |
| 24 | 印地语 | Hindi |
| 25 | 保加利亚语 | Bulgarian |
| 26 | 丹麦语 | Danish |
| 27 | 希伯来语 | Hebrew |
| 28 | 马来语 | Malay |
| 29 | 波斯语 | Persian |
| 30 | 斯洛伐克语 | Slovak |
| 31 | 瑞典语 | Swedish |
| 32 | 克罗地亚语 | Croatian |
| 33 | 菲律宾语 | Filipino |
| 34 | 匈牙利语 | Hungarian |
| 35 | 挪威语 | Norwegian |
| 36 | 斯洛文尼亚语 | Slovenian |
| 37 | 加泰罗尼亚语 | Catalan |
| 38 | 尼诺斯克语 | Nynorsk |
| 39 | 泰米尔语 | Tamil |
| 40 | 阿非利卡语 | Afrikaans |

## 官方MCP
MiniMax 提供官方的 [Python 版本](https://github.com/MiniMax-AI/MiniMax-MCP) 和 [JavaScript 版本](https://github.com/MiniMax-AI/MiniMax-MCP-JS) 模型上下文协议（MCP）服务器实现代码，支持语音合成功能，详细说明请参考 [MiniMax MCP 使用指南文档](/document/guides_mcp?key=68bbf0f295a600dbfc926bdd)

## 创建语音生成任务
- 原始链接: https://platform.minimaxi.com/document/t2a_async_create?key=68adac886602726333001546
- 文档ID: `68adac886602726333001546`
# 创建语音生成任务

<request method="POST" url="https://api.minimaxi.com/v1/t2a_async_v2"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`  
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，用于验证账户信息，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON
<br/>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
调用的模型版本。目前支持： `speech-2.5-hd-preview` `speech-2.5-turbo-preview` `speech-02-hd` `speech-02-turbo` `speech-01-hd` `speech-01-turbo` 
<hr/>

#### text    <gray>string</gray> <red>required</red>
待合成音频的文本，限制最长 5 万字符。和 "text_file_id" 二选一必填
<hr/>

#### text_file_id    <gray>int64</gray> <red>required</red>
待合成音频的文本文件 id，单个文件长度限制小于 10 万字符，支持的文件格式：txt、zip。和 "text" 二选一必填，传入后自动校验格式。
1.  **txt 文件**：长度限制 <100,000 字符。支持使用 `<#x#>` 标记自定义停顿。`x` 为停顿时长（单位：秒），范围 **[0.01,99.99]**，最多保留两位小数。注意停顿需设置在两个可以语音发音的文本之间，不可连续使用多个停顿标记
2.  **zip 文件**：
    - 压缩包内需包含同一格式的 txt 或 json 文件。
    - json 文件格式：支持 `["title", "content", "extra"]` 三个字段，分别表示标题、正文、附加信息。若三个字段都存在，则产出 3 组结果，共 9 个文件，统一存放在一个文件夹中。若某字段不存在或内容为空，则该字段不会生成对应结果
<hr/>

#### voice_setting    <gray>object</gray> <red>required</red>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray> <red>required</red>
合成音频的音色编号。可查看[音色列表](/document/system_voice_id?key=68b40964a96516e26018eee2)或使用 [ 查询可用音色接口](/document/get_voice?key=68b51016d2f0aaaf3484a8de) 查询全部可用音色
- **中文**：
  - moss_audio_ce44fc67-7ce3-11f0-8de5-96e35d26fb85
  - Chinese (Mandarin)_Lyrical_Voice
  - moss_audio_aaa1346a-7ce7-11f0-8e61-2e6e3c7ee85d
  - Chinese (Mandarin)_HK_Flight_Attendant
- **英语**：
  - English_Graceful_Lady
  - English_Insightful_Speaker
  - English_radiant_girl
  - English_Persuasive_Man
  - moss_audio_6dc281eb-713c-11f0-a447-9613c873494c
  - moss_audio_570551b1-735c-11f0-b236-0adeeecad052
  - moss_audio_ad5baf92-735f-11f0-8263-fe5a2fe98ec8
  - English_Lucky_Robot
- **日语**：
  - Japanese_Whisper_Belle
  - Serene_Woman
  - moss_audio_24875c4a-7be4-11f0-9359-4e72c55db738
  - moss_audio_7f4ee608-78ea-11f0-bb73-1e2a4cfcd245
  - moss_audio_c1a6a3ac-7be6-11f0-8e8e-36b92fbb4f95
<hr/>

#### speed    <gray>float</gray>
合成音频的语速，取值越大，语速越快。范围 `[0.5,2]`，默认 1.0
<hr/>

#### vol    <gray>float</gray>
合成音频的音量，取值越大，音量越高。范围 `(0,10]`，默认 1.0
<hr/>

#### pitch    <gray>int</gray>
合成音频的语调，范围 `[-12,12]`，默认 0，其中 0 为原音色输出
<hr/>

#### emotion    <gray>string</gray>
- 控制合成语音的情绪；
- 参数范围 `["happy", "sad", "angry", "fearful", "disgusted", "surprised", "calm"]`，分别对应 7 种情绪：高兴，悲伤，愤怒，害怕，厌恶，惊讶，中性
- 模型会根据输入文本自动匹配合适的情绪，一般无需手动指定
- 该参数仅对 `speech-2.5-hd-preview` ， `speech-2.5-turbo-preview` ， `speech-02-hd` ， `speech-02-turbo` ， `speech-01-turbo` ， `speech-01-hd` 模型生效
<hr/>

#### english_normalization    <gray>bool</gray>
支持英语文本规范化，开启后可提升数字阅读场景的性能，但会略微增加延迟，默认 false

</card>
</details>
<hr/>

#### audio_setting    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### sample_rate    <gray>int</gray>
生成音频的采样率。可选 `[8000, 16000, 22050, 24000, 32000, 44100]`，默认 32000
<hr/>

#### bitrate    <gray>int</gray>
生成音频的比特率。可选 `[32000, 64000, 128000, 256000]`，默认 128000。该参数仅对 mp3 格式的音频生效
<hr/>

#### format    <gray>string</gray>
生成音频的格式。可选 `[mp3, pcm, flac]`，默认 mp3
<hr/>

#### channel    <gray>int</gray>
生成音频的声道数。可选 `[1, 2]`，其中 1 为单声道，2 为双声道，默认 2

</card>
</details>
<hr/>

#### pronunciation_dict    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### tone    <gray>list</gray>
定义需要特殊标注的文字或符号对应的注音或发音替换规则。在中文文本中，声调用数字表示，一声为 1，二声为 2，三声为 3，四声为 4，轻声为 5
示例如下：
`["燕少飞/(yan4)(shao3)(fei1)", "达菲/(da2)(fei1)", "omg/oh my god"]`

</card>
</details>
<hr/>

#### language_boost    <gray>string</gray>
是否增强对指定的小语种和方言的识别能力。默认值为 null，可设置为 "auto" 让模型自主判断。
可选值范围：[`Chinese`, `Chinese,Yue`, `English`, `Arabic`, `Russian`, `Spanish`, `French`, `Portuguese`, `German`, `Turkish`, `Dutch`, `Ukrainian`, `Vietnamese`, `Indonesian`, `Japanese`, `Italian`, `Korean`, `Thai`, `Polish`, `Romanian`, `Greek`, `Czech`, `Finnish`, `Hindi`, `Bulgarian`, `Danish`, `Hebrew`, `Malay`, `Persian`, `Slovak`, `Swedish`, `Croatian`, `Filipino`, `Hungarian`, `Norwegian`, `Slovenian`, `Catalan`, `Nynorsk`, `Tamil`, `Afrikaans`, `auto`]
<hr/>

#### voice_modify    <gray>object</gray>
声音效果器设置
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### pitch    <gray>int</gray>
音高调整（低沉/明亮），范围 [-100,100]，数值接近 -100，声音更低沉；接近 100，声音更明亮
<img src="https://filecdn.minimax.chat/public/5d210c47-4236-4e81-893b-16cc1ef0302d.png" alt="图片描述" width="60%" />


<hr/>

#### intensity    <gray>int</gray>
强度调整（力量感/柔和），范围 [-100,100]，数值接近 -100，声音更刚劲；接近 100，声音更轻柔
<img src="https://filecdn.minimax.chat/public/862d493e-71d5-4d1f-b7c3-9ac51890631b.png" alt="图片描述" width="60%" />


<hr/>

#### timbre    <gray>int</gray>
音色调整（磁性/清脆），范围 [-100,100]，数值接近 -100，声音更浑厚；数值接近 100，声音更清脆
<img src="https://filecdn.minimax.chat/public/5f0e6cae-363a-452b-8d42-fbc4ef5a0510.png" alt="图片描述" width="60%" />


<hr/>

#### sound_effects    <gray>string</gray>
音效设置，单次仅能选择一种，可选值：
1. `spacious_echo`（空旷回音）
2. `auditorium_echo`（礼堂广播）
3. `lofi_telephone`（电话失真）
4. `robotic`（电音）

</card>
</details>
<hr/>

#### aigc_watermark    <gray>bool</gray>
控制在合成音频的末尾添加音频节奏标识，该参数仅对非流式合成生效，默认 False
<br/>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
当前任务的 ID
<hr/>

#### file_id    <gray>int64</gray>
任务创建成功后返回的对应音频文件的 ID。
- 当任务完成后，可通过 `file_id` 调用 [文件检索接口](/document/files_retrieve?key=68b56abca96516e26019203a) 进行下载
- 当请求出错时，不返回该字段
注意：返回的下载 URL 自生成起 9 小时（32,400 秒）内有效，过期后文件将失效，生成的信息便会丢失，请注意下载信息的时间
<hr/>

#### task_token    <gray>string</gray>
完成当前任务使用的密钥信息
<hr/>

#### usage_characters    <gray>int</gray>
计费字符数
<hr/>

#### base_resp    <gray>dict</gray>
本次请求的状态码及其详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码
- 0: 正常
- 1002: 限流
- 1004: 鉴权失败
- 1039：触发 TPM 限流
- 1042: 非法字符超10%
- 2013: 参数异常

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

### 返回文件信息

  <hr/>

  #### 输入文件

  <details>
  <summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
  <card>

  #### txt文件
   输出文件如下所示

   - 音频文件：文件格式遵从请求体设置
   - 字幕文件：精确到句的字幕信息
   - 额外信息 JSON 文件：音频文件相关的附加信息


  <hr/>

  #### json 文件


  <details>
  <summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
  <card>

  #### title
若该字段为空，则不输出该字段的文件

  - 音频文件：文件格式遵从请求体设置
  - 字幕文件：精确到句的字幕信息
  - 额外信息 JSON 文件：音频文件相关的附加信息

  <hr/>

  #### content

  若该字段为空，则不输出该字段的文件

  - 音频文件：文件格式遵从请求体设置
  - 字幕文件：精确到句的字幕信息
  - 额外信息 JSON 文件：音频文件相关的附加信息

 <hr/>

 #### extra

 若该字段为空，则不输出该字段的文件

  - 音频文件：文件格式遵从请求体设置
  - 字幕文件：精确到句的字幕信息
  - 额外信息 JSON 文件：音频文件相关的附加信息


  </card>
  </details>

  </card>
  </details>

</left>

<right>
<tabs>
<tab title="文本输入">

<codetabs title="请求示例">
<codetab language="python">
import requests
import json
import os

api_key = os.environ.get("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/t2a_async_v2"
file_format = "mp3"
payload = json.dumps({
  "model": "speech-2.5-hd-preview",
  "text": "真正的危险不是计算机开始像人一样思考，而是人开始像计算机一样思考。计算机只是可以帮我们处理一些简单事务。",
  "language_boost": "auto",
  "voice_setting": {
    "voice_id": "audiobook_male_1",
    "speed": 1,
    "vol": 1,
    "pitch": 1
  },
  "pronunciation_dict": {
    "tone": [
      "危险/dangerous"
    ]
  },
  "audio_setting": {
    "audio_sample_rate": 32000,
    "bitrate": 128000,
    "format": file_format,
    "channel": 2
  },
  "voice_modify":{
      "pitch":0,
      "intensity":0,
      "timbre":0,
      "sound_effects":"spacious_echo"
    }
})
headers = {
  'Authorization': f'Bearer {api_key}',
  'Content-Type': 'application/json'
}
response = requests.request("POST", url, headers=headers, data=payload)
print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/t2a_async_v2' \
--header "authorization: Bearer ${MINIMAX_API_KEY}" \
--header 'Content-Type: application/json' \
--data '{
    "model": "speech-2.5-hd-preview",
  "text": "真正的危险不是计算机开始像人一样思考，而是人开始像计算机一样思考。计算机只是可以帮我们处理一些简单事务。",
  "language_boost": "auto",
  "voice_setting": {
    "voice_id": "audiobook_male_1",
    "speed": 1,
    "vol": 1,
    "pitch": 1
  },
  "pronunciation_dict": {
    "tone": [
      "危险/dangerous"
    ]
  },
  "audio_setting": {
    "audio_sample_rate": 32000,
    "bitrate": 128000,
    "format": "mp3",
    "channel": 2
  },
  "voice_modify":{
      "pitch":0,
      "intensity":0,
      "timbre":0,
      "sound_effects":"spacious_echo"
    }
}'
</codetab>
</codetabs>


</tab>
<tab title="文件输入">

<codetabs title="请求示例">
<codetab language="python">
"""
使用文件作为音频合成的输入时，需要先上传文件获取 file_id，并将 file_id 填入 "text_file_id" 处，注意 file_id 的格式是整数。
"""
import requests
import json
import os

api_key = os.getenv("MINIMAX_API_KEY")
text_file_id = int(os.environ.get("TEXT_FILE_ID"))
url = "https://api.minimaxi.com/v1/t2a_async_v2"
file_format = "mp3"
payload = json.dumps({
  "model": "speech-2.5-hd-preview",
  "text_file_id": text_file_id,
  "language_boost": "auto",
  "voice_setting": {
    "voice_id": "audiobook_male_1",
    "speed": 1,
    "vol": 10,
    "pitch": 1
  },
  "pronunciation_dict": { "tone": [ "草地/(cao3)(di1)" ] },
  "audio_setting": {
    "audio_sample_rate": 32000,
    "bitrate": 128000,
    "format": file_format,
    "channel": 2
  },
    "voice_modify":{
      "pitch":0,
      "intensity":0,
      "timbre":0,
      "sound_effects":"spacious_echo"
    }
})
headers = {
  'Authorization': f'Bearer {api_key}',
  'Content-Type': 'application/json'
}
response = requests.request("POST", url, headers=headers, data=payload)
print(response.text)

</codetab>

<codetab language="curl">
# 注意：请将下面的 text_file_id 改为你的输入文件的 file_id
curl --location 'https://api.minimaxi.com/v1/t2a_async_v2' \
--header "authorization: Bearer ${MINIMAX_API_KEY}" \
--header 'Content-Type: application/json' \
--data '{
  "model": "speech-2.5-hd-preview",
  "text_file_id": <text_file_id>,
  "language_boost": "auto",
  "voice_setting": {
    "voice_id": "audiobook_male_1",
    "speed": 1,
    "vol": 10,
    "pitch": 1
  },
  "pronunciation_dict": { "tone": [ "草地/(cao3)(di1)" ] },
  "audio_setting": {
    "audio_sample_rate": 32000,
    "bitrate": 128000,
    "format": "mp3",
    "channel": 2
  },
    "voice_modify":{
      "pitch":0,
      "intensity":0,
      "timbre":0,
      "sound_effects":"spacious_echo"
    }
}'
</codetab>
</codetabs>

</tab>
</tabs>

<tabs>
<tab title="文本输入">
<codetabs title="返回示例">
<codetab language="json">
{
    "task_id": 95157322514444,
    "task_token": "eyJhbGciOiJSUz",
    "file_id": 95157322514444,
    "usage_characters":101,
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>

</codetabs>

</tab>


<tab title="文件输入">
<codetabs title="返回示例">
<codetab language="json">
{
    "task_id": 95157322514444,
    "task_token": "eyJhbGciOiJSUz",
    "file_id": 95157322514444,
    "usage_characters":101,
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}

</codetab>
</codetabs>

</tab>
</tabs>

</right>

</container>

<container gap="4" columns="2">

  <card  href="/document/guides_t2a_async?key=68bbecb495a600dbfc925d35">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[异步语音合成指南](/document/guides_t2a_async?key=68bbecb495a600dbfc925d35)

  </card>

<card  href="/document/t2a_async_api_intro?key=68adac446fe587e3fbfe965b">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[异步语音合成接口能力介绍](/document/t2a_async_api_intro?key=68adac446fe587e3fbfe965b)
  </card>

</container>

## 查询语音生成任务状态
- 原始链接: https://platform.minimaxi.com/document/t2a_async_query?key=68adad0f6fe587e3fbfe9810
- 文档ID: `68adad0f6fe587e3fbfe9810`
# 查询语音生成任务状态

<request method="GET" url="http://api.minimaxi.com/v1/query/t2a_async_query_v2"></request>

**注：该 API 限制每秒最多查询 10 次。**


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`  <red>Required</red>
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，用于验证账户信息，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON
<br/>


### 请求体
<hr/>

#### task_id    <gray>string</gray> <red>required</red>
任务 ID，提交任务时返回的信息
<br/>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
任务 ID
<hr/>

#### status    <gray>string</gray>
该任务的当前状态。
- **Processing**：该任务正在处理中
- **Success**：该任务已完成
- **Failed**：任务失败
- **Expired**：任务已过期
<hr/>

#### file_id    <gray>int64</gray>
任务创建成功后返回的对应音频文件的 ID
- 当任务完成后，可通过` file_id `调用 [文件检索](/document/files_retrieve?key=68b56abca96516e26019203a) 接口进行下载。
- 当请求出错时，不返回该字段

⚠️ 注意：返回的下载 URL 自生成起 9 小时（32,400 秒）内有效
<hr/>

#### base_resp    <gray>dict</gray>
状态码及其详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码
- 0: 请求成功
- 1000：未知错误
- 1001：超时
- 1002：触发限流
- 1004：鉴权失败
- 1042：非法字符超过 10%
- 2013：输入格式信息不正常

更多状态码信息请参考[错误码查询](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>
<codetabs title="请求示例">
<codetab language="python">
"""
本示例用于查询异步语音合成任务的进度。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import json
import os

task_id = os.environ.get("TASK_ID")
api_key = os.getenv("MINIMAX_API_KEY")
if not task_id:
  print("Please set task_id first")
  raise ValueError("Please set task_id first")
url = f"https://api.minimaxi.com/v1/query/t2a_async_query_v2?task_id={task_id}"

payload = {}
headers = {
  'Authorization': f'Bearer {api_key}',
  'content-type': 'application/json',
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)

</codetab>
<codetab language="curl">
# 请先将 TASK_ID 和 MINIMAX_API_KEY 导入环境变量
curl --location "https://api.minimaxi.com/v1/query/t2a_async_query_v2?task_id=${TASK_ID}" \
--header "authorization: Bearer ${MINIMAX_API_KEY}" \
--header 'content-type: application/json'
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "task_id": 95157322514444,
    "status": "Processing", 
    "file_id": 95157322514496,
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<container gap="4" columns="2">

  <card  href="/document/guides_t2a_async?key=68bbecb495a600dbfc925d35">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[异步语音合成指南](/document/guides_t2a_async?key=68bbecb495a600dbfc925d35)

  </card>

<card  href="/document/t2a_async_api_intro?key=68adac446fe587e3fbfe965b">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[异步语音合成接口能力介绍](/document/t2a_async_api_intro?key=68adac446fe587e3fbfe965b)
  </card>

</container>

## 同步语音合成HTTP
- 原始链接: https://platform.minimaxi.com/document/t2a_http?key=68ad78146fe587e3fbfe8e03
- 文档ID: `68ad78146fe587e3fbfe8e03`
# 同步语音合成 HTTP

<request method="POST" url="https://api.minimaxi.com/v1/t2a_v2"></request>

备用接口地址
<request method="POST" url="https://api-bj.minimaxi.com/v1/t2a_v2"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，用于验证账户信息，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>string</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
请求的模型版本，可选范围： `speech-2.5-hd-preview` 、 `speech-2.5-turbo-preview` 、 `speech-02-hd` 、 `speech-02-turbo` 、 `speech-01-hd` 、 `speech-01-turbo` 
<hr/>

#### text    <gray>string</gray> <red>required</red>
需要合成语音的文本，长度限制小于 10000 字符，若文本长度大于 3000 字符，推荐使用流式输出
- 段落切换用换行符标记
- 停顿控制：支持自定义文本之间的语音时间间隔，以实现自定义文本语音停顿时间的效果。使用方式：在文本中增加`<#x#>`标记，`x` 为停顿时长（单位：秒），范围 [0.01, 99.99]，最多保留两位小数。文本间隔时间需设置在两个可以语音发音的文本之间，不可连续使用多个停顿标记
<hr/>

#### stream    <gray>boolean</gray>
控制是否流式输出。默认 false，即不开启流式
<hr/>

#### stream_options    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### exclude_aggregated_audio    <gray>boolean</gray>
设置最后一个 chunk 是否包含拼接后的语音 hex 数据。默认值为 False，即最后一个 chunk 中包含拼接后的完整语音 hex 数据

</card>
</details>
<hr/>

#### voice_setting    <gray>object</gray> <red>required</red>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray> <red>required</red>
合成音频的音色编号。若需要设置混合音色，请设置 timbre_weights 参数，本参数设置为空值。支持系统音色、复刻音色以及文生音色三种类型，以下是部分最新的系统音色（ID），可查看 [系统音色列表](/document/system_voice_id?key=68b40964a96516e26018eee2) 或使用 [查询可用音色 API](/document/get_voice?key=68b51016d2f0aaaf3484a8de) 查询系统支持的全部音色
- **中文**：
  - moss_audio_ce44fc67-7ce3-11f0-8de5-96e35d26fb85
  - moss_audio_aaa1346a-7ce7-11f0-8e61-2e6e3c7ee85d
  - Chinese (Mandarin)_Lyrical_Voice
  - Chinese (Mandarin)_HK_Flight_Attendant
- **英语**：
  - English_Graceful_Lady
  - English_Insightful_Speaker
  - English_radiant_girl
  - English_Persuasive_Man
  - moss_audio_6dc281eb-713c-11f0-a447-9613c873494c
  - moss_audio_570551b1-735c-11f0-b236-0adeeecad052
  - moss_audio_ad5baf92-735f-11f0-8263-fe5a2fe98ec8
  - English_Lucky_Robot
- **日语**：
  - Japanese_Whisper_Belle
  - moss_audio_24875c4a-7be4-11f0-9359-4e72c55db738
  - moss_audio_7f4ee608-78ea-11f0-bb73-1e2a4cfcd245
  - moss_audio_c1a6a3ac-7be6-11f0-8e8e-36b92fbb4f95
<hr/>

#### speed    <gray>float</gray>
合成音频的语速，取值越大，语速越快。取值范围 [0.5,2]，默认值为1.0
<hr/>

#### vol    <gray>float</gray>
合成音频的音量，取值越大，音量越高。取值范围 (0,10]，默认值为 1.0
<hr/>

#### pitch    <gray>int</gray>
合成音频的语调，取值范围 [-12,12]，默认值为 0，其中 0 为原音色输出
<hr/>

#### emotion    <gray>string</gray>
- 控制合成语音的情绪
- 参数范围 `["happy", "sad", "angry", "fearful", "disgusted", "surprised", "calm"]`，分别对应 7 种情绪：高兴，悲伤，愤怒，害怕，厌恶，惊讶，中性
- 模型会根据输入文本自动匹配合适的情绪，一般无需手动指定
- 该参数仅对 `speech-2.5-hd-preview` ， `speech-2.5-turbo-preview` ， `speech-02-hd` ， `speech-02-turbo` ， `speech-01-turbo` ， `speech-01-hd` 模型生效
<hr/>

#### text_normalization    <gray>boolean</gray>
是否启用中文、英语文本规范化，开启后可提升数字阅读场景的性能，但会略微增加延迟，默认值为 false
<hr/>

#### latex_read    <gray>boolean</gray>
控制是否朗读 latex 公式，默认为 false
**需注意**：
1. 请求中的公式需要在公式的首尾加上$$
2. 请求中公式若有`"\"`，需转义成`"\\"`
示例：一元二次方程根的基本公式 

<img src="https://filecdn.minimax.chat/public/d6f62e9a-cd3f-4f55-a237-257eef531683.png" alt="图片描述" width="30%" />

 应表示为`$$x = \\frac{-b \pm \sqrt{b^2 - 4ac}}{2a} $$`

</card>
</details>
<hr/>

#### audio_setting    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### sample_rate    <gray>int</gray>
生成音频的采样率。可选范围`[8000，16000，22050，24000，32000，44100]`，默认为 32000
<hr/>

#### bitrate    <gray>int</gray>
生成音频的比特率。可选范围`[32000，64000，128000，256000]`，默认值为 128000。该参数仅对 mp3 格式的音频生效
<hr/>

#### format    <gray>string</gray>
生成音频的格式。可选范围`[mp3, pcm, flac, wav]`，默认值为 mp3，wav 仅在非流式输出下支持
<hr/>

#### channel    <gray>int</gray>
生成音频的声道数。可选范围：`[1,2]`，其中 1 为单声道，2 为双声道，默认值为 1
<hr/>

#### force_cbr    <gray>boolean</gray>
对于音频恒定比特率（cbr）控制，可选`false`、 `ture`。当此参数设置为`ture`，将以恒定比特率方式进行音频编码。
注意：本参数仅当音频设置为**流式**输出，且音频格式为 **mp3** 时生效。

</card>
</details>
<hr/>

#### pronunciation_dict    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### tone    <gray>list</gray>
定义需要特殊标注的文字或符号对应的注音或发音替换规则。在中文文本中，声调用数字表示：
一声为 1，二声为 2，三声为 3，四声为 4，轻声为 5
示例如下：
`["燕少飞/(yan4)(shao3)(fei1)",  "omg/oh my god"]`

</card>
</details>
<hr/>

#### timbre_weights    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray>
合成音频的音色编号，须和weight参数同步填写。支持系统音色、复刻音色以及文生音色三种类型。系统支持的全部音色可查看 [系统音色列表](/document/system_voice_id?key=68b40964a96516e26018eee2)，也可使用 [查询可用音色 API](/document/get_voice?key=68b51016d2f0aaaf3484a8de) 查询系统支持的全部音色
<hr/>

#### weight    <gray>int</gray>
合成音频各音色所占的权重，须与 voice_id 同步填写。可选值范围为[1, 100]，最多支持 4 种音色混合，单一音色取值占比越高，合成音色与该音色相似度越高

<codetabs title="参数示例">
<codetab language="json">
 "timbre_weights": [
  {
    "voice_id": "female-chengshu",
    "weight": 30
  },
  {
    "voice_id": "female-tianmei",
    "weight": 70
  }]

</codetab>

</codetabs>

</card>
</details>
<hr/>

#### language_boost    <gray>string</gray>
是否增强对指定的小语种和方言的识别能力。默认值为 null，可设置为 "auto" 让模型自主判断。
可选值范围：[`Chinese`, `Chinese,Yue`, `English`, `Arabic`, `Russian`, `Spanish`, `French`, `Portuguese`, `German`, `Turkish`, `Dutch`, `Ukrainian`, `Vietnamese`, `Indonesian`, `Japanese`, `Italian`, `Korean`, `Thai`, `Polish`, `Romanian`, `Greek`, `Czech`,  `Finnish`, `Hindi`, `Bulgarian`, `Danish`, `Hebrew`, `Malay`, `Persian`, `Slovak`, `Swedish`, `Croatian`, `Filipino`, `Hungarian`, `Norwegian`, `Slovenian`, `Catalan`, `Nynorsk`, `Tamil`, `Afrikaans`, `auto`]
<hr/>

#### voice_modify    <gray>object</gray>
声音效果器设置，该参数支持的音频格式：
- 非流式：mp3, wav, flac
- 流式：mp3

<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### pitch    <gray>int</gray>
音高调整（低沉/明亮），范围 [-100,100]，数值接近 -100，声音更低沉；接近 100，声音更明亮
<img src="https://filecdn.minimax.chat/public/5d210c47-4236-4e81-893b-16cc1ef0302d.png" alt="图片描述" width="80%" />


<hr/>

#### intensity    <gray>int</gray>
强度调整（力量感/柔和），范围 [-100,100]，数值接近 -100，声音更刚劲；接近 100，声音更轻柔
<img src="https://filecdn.minimax.chat/public/862d493e-71d5-4d1f-b7c3-9ac51890631b.png" alt="图片描述" width="80%" />


<hr/>

#### timbre    <gray>int</gray>
音色调整（磁性/清脆），范围 [-100,100]，数值接近 -100，声音更浑厚；数值接近 100，声音更清脆
<img src="https://filecdn.minimax.chat/public/5f0e6cae-363a-452b-8d42-fbc4ef5a0510.png" alt="图片描述" width="80%" />


<hr/>

#### sound_effects    <gray>string</gray>
音效设置，单次仅能选择一种，可选值：
1. `spacious_echo`（空旷回音）
2. `auditorium_echo`（礼堂广播）
3. `lofi_telephone`（电话失真）
4. `robotic`（电音）

</card>
</details>
<hr/>

#### subtitle_enable    <gray>boolean</gray>
控制是否开启字幕服务，默认值为 false。此参数仅在非流式输出场景下有效，且仅对 `speech-2.5-hd-preview` `speech-2.5-turbo-preview`  `speech-02-turbo` `speech-02-hd` `speech-01-turbo` `speech-01-hd` 模型有效
<hr/>

#### output_format    <gray>string</gray>
控制输出结果形式的参数，可选值范围为[`url`, `hex`]，默认值为 `hex` 。该参数仅在非流式场景生效，流式场景仅支持返回 hex 形式。返回的 url 有效期为 24 小时
<hr/>

#### aigc_watermark    <gray>boolean</gray>
控制在合成音频的末尾添加音频节奏标识，默认值为 False。该参数仅对非流式合成生效

<br/>

### 返回体
<hr/>

#### data    <gray>object</gray>
返回的合成数据对象，可能为 null，需进行非空判断
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### audio    <gray>string</gray>
合成后的音频数据，采用 hex 编码，格式与请求中指定的输出格式一致
<hr/>

#### subtitle_file    <gray>string</gray>
合成的字幕下载链接。音频文件对应的字幕，精确到句（不超过 50 字），单位为毫秒，格式为 json
<hr/>

#### status    <gray>int</gray>
当前音频流状态：1 表示合成中，2 表示合成结束

</card>
</details>
<hr/>

#### trace_id    <gray>string</gray>
本次会话的 id，用于在咨询/反馈时帮助定位问题
<hr/>

#### extra_info    <gray>object</gray>
音频的附加信息
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### audio_length    <gray>int64</gray>
音频时长（毫秒）
<hr/>

#### audio_sample_rate    <gray>int64</gray>
音频采样率
<hr/>

#### audio_size    <gray>int64</gray>
音频文件大小（字节）
<hr/>

#### bitrate    <gray>int64</gray>
音频比特率
<hr/>

#### audio_format    <gray>string</gray>
生成音频文件的格式。取值范围 [`mp3`, `pcm`, `flac`]
<hr/>

#### audio_channel    <gray>int64</gray>
生成音频声道数,1：单声道，2：双声道
<hr/>

#### invisible_character_ratio    <gray>float64</gray>
非法字符占比.非法字符不超过 10%（包含 10%），音频会正常生成,并返回非法字符占比数据；如超过 10% 将进行报错
<hr/>

#### usage_characters    <gray>int64</gray>
计费字符数
<hr/>

#### word_count    <gray>int64</gray>
已发音的字数统计，包含汉字、数字、字母，不包含标点符号

</card>
</details>
<hr/>

#### base_resp    <gray>object</gray>
本次请求的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码。
 - 0: 请求结果正常
 - 1000：未知错误
 - 1001：超时
 - 1002：触发限流
 - 1004：鉴权失败
 - 1039：触发 TPM 限流
 - 1042：非法字符超过 10%
 - 2013：输入参数信息不正常

 更多信息请参考 [错误码查询](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<tabs>
<tab title="非流式">
<codetabs title="请求示例">
<codetab language="python">
import requests
import json

api_key = "填写你的api key"
url = f"https://api.minimaxi.com/v1/t2a_v2"

payload = json.dumps({
    "model": "speech-2.5-hd-preview",
    "text": "今天是不是很开心呀，当然了！",
    "stream": False,
    "voice_setting":{
        "voice_id": "male-qn-qingse",
        "speed": 1,
        "vol": 1,
        "pitch": 0,
        "emotion": "happy"
    },
    "pronunciation_dict":{
        "tone": ["处理/(chu3)(li3)", "危险/dangerous"]
    },
    "audio_setting":{
        "sample_rate": 32000,
        "bitrate": 128000,
        "format": "mp3",
        "channel": 1
    },
    "subtitle_enable": False
  })
headers = {
    'Authorization': f'Bearer {api_key}',
    'Content-Type': 'application/json'
}

response = requests.request("POST", url, stream=True, headers=headers, data=payload)
parsed_json = json.loads(response.text)

# get audio
audio_value = bytes.fromhex(parsed_json['data']['audio'])
with open('output.mp3', 'wb') as f:
    f.write(audio_value)

</codetab>

<codetab language="curl">
curl --location --request POST 'https://api.minimaxi.com/v1/t2a_v2' \
--header 'Authorization: Bearer ${your_api_key}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "model": "speech-2.5-hd-preview",
    "text": "今天是不是很开心呀，当然了！",
    "stream": false,
    "voice_setting":{
        "voice_id": "male-qn-qingse",
        "speed": 1,
        "vol": 1,
        "pitch": 0,
        "emotion": "happy"
    },
    "pronunciation_dict":{
        "tone": ["处理/(chu3)(li3)", "危险/dangerous"]
    },
    "audio_setting":{
        "sample_rate": 32000,
        "bitrate": 128000,
        "format": "mp3",
        "channel": 1
    },
    "subtitle_enable": false
  }'
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "data":{
        "audio":"<hex编码的audio>",
        "status":2
    },
    "extra_info": {
        "audio_length": 9900,
        "audio_sample_rate": 32000,
        "audio_size": 160323,
        "bitrate": 128000,
        "word_count": 52,
        "invisible_character_ratio": 0,
        "usage_characters": 101,
        "audio_format": "mp3",
        "audio_channel": 1
    },
    "trace_id":"01b8bf9bb7433cc75c18eee6cfa8fe21",
    "base_resp":{
        "status_code":0,
        "status_msg":"success"
    }
}
</codetab>
</codetabs>

</tab>
<tab title="流式">
<codetabs title="请求示例">
<codetab language="python">
import json
import subprocess
import time
from typing import Iterator

import requests

api_key = "你的api_key"

url = f"https://api.minimaxi.com/v1/t2a_v2"

file_format = 'mp3'  # support mp3/pcm/flac

headers = {"Content-Type": "application/json", "Authorization": "Bearer " + api_key}


def build_tts_stream_headers() -> dict:
    headers = {
        'accept': 'application/json, text/plain, */*',
        'content-type': 'application/json',
        'authorization': "Bearer " + api_key,
    }
    return headers


def build_tts_stream_body(text: str) -> dict:
    body = json.dumps({
    "model": "speech-2.5-hd-preview",
    "text": "今天是不是很开心呀，当然了！",
    "stream": True,
    "voice_setting":{
        "voice_id": "male-qn-qingse",
        "speed": 1,
        "vol": 1,
        "pitch": 0,
        "emotion": "happy"
    },
    "pronunciation_dict":{
        "tone": ["处理/(chu3)(li3)", "危险/dangerous"]
    },
    "audio_setting":{
        "sample_rate": 32000,
        "bitrate": 128000,
        "format": "mp3",
        "channel": 1
    },
    "subtitle_enable": False
  })
    return body


mpv_command = ["mpv", "--no-cache", "--no-terminal", "--", "fd://0"]
mpv_process = subprocess.Popen(
    mpv_command,
    stdin=subprocess.PIPE,
    stdout=subprocess.DEVNULL,
    stderr=subprocess.DEVNULL,
)


def call_tts_stream(text: str) -> Iterator[bytes]:
    tts_url = url
    tts_headers = build_tts_stream_headers()
    tts_body = build_tts_stream_body(text)

    response = requests.request("POST", tts_url, stream=True, headers=headers, data=tts_body)
    for chunk in (response.raw):
        if chunk:
            if chunk[:5] == b'data:':
                data = json.loads(chunk[5:])
                if "data" in data and "extra_info" not in data:
                    if "audio" in data["data"]:
                        audio = data["data"]['audio']
                        yield audio
    return audio


def audio_play(audio_stream: Iterator[bytes]) -> bytes:
    audio = b""
    for chunk in audio_stream:
        if chunk is not None and chunk != '\n':
            decoded_hex = bytes.fromhex(chunk)
            mpv_process.stdin.write(decoded_hex)  # type: ignore
            mpv_process.stdin.flush()
            audio += decoded_hex

    return audio


audio_chunk_iterator = call_tts_stream('')
audio = audio_play(audio_chunk_iterator)

# save results to file
timestamp = int(time.time())
file_name = f'output_total_{timestamp}.{file_format}'
with open(file_name, 'wb') as file:
    file.write(audio)
</codetab>
</codetabs>

<codetabs  title="返回示例">
<codetab language="json">
//返回的第一个chunk
{
    "data":{
        "audio":"hex编码的audio_chunk1",
        "status":1
    },
    "trace_id":"01b8bf9bb7433cc75c18eee6cfa8fe21",
    "base_resp":{ "status_code":0, "status_msg":"" }
}
//返回的第二个chunk
{
    "data":{
        "audio":"hex编码的audio_chunk2",
        "status":1
    },
    "trace_id":"01b8bf9bb7433cc75c18eee6cfa8fe21",
    "base_resp":{ "status_code":0, "status_msg":"" }
}
//...
//结束
{
    "data":{
        "audio":"hex编码的audio",
        "status":2
    },
    "extra_info": {
      "audio_length": 6931,
      "audio_sample_rate": 32000,
      "audio_size": 111789,
      "bitrate": 128000,
      "word_count": 112,
      "invisible_character_ratio": 0,
      "usage_characters": 112,
      "audio_format": "mp3",
      "audio_channel": 1
    },
    "trace_id": "04ece790375f3ca2edbb44e8c4c200bf",
    "base_resp": {
      "status_code": 0,
      "status_msg": "success"
    }
}
</codetab>
</codetabs>
</tab>
</tabs>

</right>

</container>

## “任务失败”事件
- 原始链接: https://platform.minimaxi.com/document/task_fail?key=68ada0d06fe587e3fbfe9574
- 文档ID: `68ada0d06fe587e3fbfe9574`
# “任务失败”事件
如果接收到 `task_failed` 事件，表示任务失败。此时需要关闭 WebSocket 连接并处理错误。

<container>

<left>

### 返回参数
<hr/>

#### trace_id    <gray>string</gray>
表示会话中单次请求的 id，用于在咨询/反馈时帮助定位问题
<hr/>

#### session_id    <gray>string</gray>
表示整个会话的 id
<hr/>

#### event    <gray>string</gray>
表示会话类型，任务失败会返回 "task_failed"
<hr/>

#### base_resp    <gray>dict</gray>
本次请求的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码。
- 1000，未知错误
- 1001，超时
- 1002，触发限流
- 1004，鉴权失败
- 1039，触发 TPM 限流
- 1042，非法字符超过 10%
- 2013，输入格式信息不正常
- 2201，超时断连

更多状态码信息请参考[错误码查询](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="返回示例">
<codetab language="json">
{
    "session_id": "xxxx",
    "event": "task_failed",
    "trace_id": "0303a2882bf18235ae7a809ae0f3cca7",
    "base_resp": {
        "status_code": 1004,
        "status_msg": "XXXXXXX"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>

</container>

<container gap="4" columns="2">

  <card  href="https://www.example.com">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成指南 WebSocket](/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba)


  </card>

<card  href="https://www.example.com">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

[同步语音合成接口能力介绍](/document/t2a_api_intro?key=68ad77666602726333000457)


  </card>

</container>

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/text_api_intro?key=68abd86ad08627aad9673eaa
- 文档ID: `68abd86ad08627aad9673eaa`
# 文本生成 (ChatCompletion)
文本生成接口使用语言模型创建对话补全，根据输入的上下文，让模型补全对话内容。

该接口兼容 OpenAI 参数定义，提供基础对话及工具调用（Function Calling）能力。可直接通过 HTTP 请求或 OpenAI SDK 进行接入。

## 支持模型

| 模型名称 | 最大token数 | 模型介绍 |
| :--- | :---: | :--- |
| MiniMax-M1 | 1000192 | 全球领先：80K思维链 x 1M输入，效果比肩海外顶尖模型 |
| MiniMax-Text-01 | 1000192 | 全新架构文本模型，支持1M超长上下文 |

注：最大token数为：一次请求+回复的输入输出总token


## 接口说明

文本生成接口支持多种功能：

*   **多模态支持**：根据您使用的模型，支持不同的消息类型（模态），如文本、图像
*   **流式响应**：如果设置为 True，模型响应数据将在生成时使用服务器发送事件流式传输到客户端
*   **输出格式**：支持多种输出格式，包括默认的文本响应、JSON对象以及用于生成结构化JSON响应的JSON Schema

## 文生图
- 原始链接: https://platform.minimaxi.com/document/text_to_image?key=68ac01e26fe587e3fbfe5765
- 文档ID: `68ac01e26fe587e3fbfe5765`
# 文生图


<request method="POST" url="https://api.minimaxi.com/v1/image_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br>

### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
模型名称。可选值：`image-01`, `image-01-live`
<hr/>

#### prompt    <gray>string</gray> <red>required</red>
图像的文本描述，最长 1500 字符
<hr/>


#### style    <gray>object</gray>
画风设置，仅当 `model` 为 `image-01-live` 时生效
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### style_type    <gray>string</gray> <red>required</red>
画风风格类型。可选值：`漫画`, `元气`, `中世纪`, `水彩`
<hr/>

#### style_weight    <gray>float</gray>
画风权重，取值范围 `(0, 1]`，默认 `0.8`

</card>
</details>
<hr/>

#### aspect_ratio    <gray>string</gray>
图像宽高比，默认为 `1:1`。可选值：
- `1:1` (1024x1024)
- `16:9` (1280x720)
- `4:3` (1152x864)
- `3:2` (1248x832)
- `2:3` (832x1248)
- `3:4` (864x1152)
- `9:16` (720x1280)
- `21:9` (1344x576)
<hr/>

#### width    <gray>int</gray>
生成图片的宽度（像素）。仅当 `model` 为 `image-01` 时生效。注意：`width` 和 `height` 需同时设置，取值范围[512, 2048]，且必须是 8 的倍数。若与 `aspect_ratio` 同时设置，则优先使用 `aspect_ratio`
<hr/>

#### height    <gray>int</gray>
生成图片的高度（像素）。仅当 `model` 为 `image-01` 时生效。注意：`width` 和 `height` 需同时设置，取值范围[512, 2048]，且必须是 8 的倍数。若与 `aspect_ratio` 同时设置，则优先使用 `aspect_ratio`
<hr/>

#### response_format    <gray>string</gray>
返回图片的形式，默认为 `url`。可选值：`url`, `base64`。注意：`url` 的有效期为 24 小时
<hr/>

#### seed    <gray>int</gray>
随机种子。使用相同的 seed 和参数，可以生成内容相近的图片，用于复现结果。如未提供，算法会对 n 张图单独生成随机种子
<hr/>

#### n    <gray>int</gray>
单次请求生成的图片数量，取值范围[1, 9]，默认为 1
<hr/>

#### prompt_optimizer    <gray>bool</gray>
是否开启 `prompt` 自动优化，默认为 `false`
<hr/>

#### aigc_watermark    <gray>bool</gray>
是否在生成的图片中添加水印，默认为 `false`

<br>

### 返回体

<hr/>

#### data    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### image_urls    <gray>array</gray>
当 `response_format` 为 `url` 时返回，包含图片链接的数组
<hr/>

#### image_base64    <gray>array</gray>
当 `response_format` 为 `base64` 时返回，包含图片 Base64 编码的数组

</card>
</details>
<hr/>

#### metadata    <gray>object</gray>
生成结果的附加信息
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### success_count    <gray>integer</gray>
成功生成的图片数量
<hr/>

#### failed_count    <gray>integer</gray>
因内容安全检查失败而未返回的图片数量

</card>
</details>
<hr/>

#### base_resp    <gray>object</gray>
状态码及详情。
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0，请求成功

1002，触发限流，请稍后再试

1004，账号鉴权失败，请检查 API-Key 是否填写正确

1008，账号余额不足

1026，图片描述涉及敏感内容

2013，传入参数异常，请检查入参是否按要求填写

2049，无效的api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>

</left>

<right>
<codetabs title="请求示例">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/image_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "image-01",
    "prompt": "men Dressing in white t shirt, full-body stand front view image :25, outdoor, Venice beach sign, full-body image, Los Angeles, Fashion photography of 90s, documentary, Film grain, photorealistic",
    "aspect_ratio": "16:9",
    "response_format": "url",
    "n": 3,
    "prompt_optimizer": True,
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())

</codetab>
<codetab language="curl">
curl -X POST "https://api.minimaxi.com/v1/image_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "model": "image-01",
    "prompt": "men Dressing in white t shirt, full-body stand front view image :25, outdoor, Venice beach sign, full-body image, Los Angeles, Fashion photography of 90s, documentary, Film grain, photorealistic",
    "aspect_ratio": "16:9",
    "response_format": "url",
    "n": 3,
    "prompt_optimizer": true
  }'
</codetab>

</codetabs>

<codetabs title="返回示例">
<codetab  language="json">
{
  "id": "03ff3cd0820949eb8a410056b5f21d38",
  "data": {
    "image_urls": [
      "XXX",
      "XXX",
      "XXX"
    ]
  },
  "metadata": {
    "failed_count": "0",
    "success_count": "3"
  },
  "base_resp": {
    "status_code": 0,
    "status_msg": "success"
  }
}
</codetab>
</codetabs>

</right>

</container>

## 创建文生视频任务
- 原始链接: https://platform.minimaxi.com/document/text_to_video?key=68abdfffd08627aad9674bff
- 文档ID: `68abdfffd08627aad9674bff`
# 文生视频生成任务

<request method="POST" url="https://api.minimaxi.com/v1/video_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 

- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>Required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON
<br/>
### 请求体
<hr/>

#### model    <gray>string</gray> <red>required</red>
模型名称。可用值：`MiniMax-Hailuo-02`, `T2V-01-Director`, `T2V-01`
<hr/>

#### prompt    <gray>string</gray>
视频的文本描述，最大 2000 字符。对于 `MiniMax-Hailuo-02` 和 `*-Director` 系列模型，支持使用 `[指令]` 语法进行运镜控制。
可在 `prompt` 中通过 `[指令]` 格式添加运镜指令，以实现精确的镜头控制。
- 支持 15 种运镜指令的指令:
  - 左右移: `[左移]`, `[右移]`
  - 左右摇: `[左摇]`, `[右摇]`
  - 推拉: `[推进]`, `[拉远]`
  - 升降: `[上升]`, `[下降]`
  - 上下摇: `[上摇]`, `[下摇]`
  - 变焦: `[变焦推近]`, `[变焦拉远]`
  - 其他: `[晃动]`, `[跟随]`, `[固定]`
- 使用规则:
  - 组合运镜: 同一组 `[]` 内的多个指令会同时生效，如 `[左摇,上升]`，建议组合不超过 3 个
  - 顺序运镜: `prompt` 中前后出现的指令会依次生效，如 `"...[推进], 然后...[拉远]"`
  - 自然语言: 也支持通过自然语言描述运镜，但使用标准指令能获得更准确的响应
  - 多运镜教程及模板，可参考 [运镜使用教程](https://b6j6u5f4zf.feishu.cn/docx/H03ZdRLh7oFsNpxE2ZqcUWb3njg) 
<hr/>

#### prompt_optimizer    <gray>boolean</gray>
是否自动优化 `prompt`，默认为 `true`。设为 `false` 可进行更精确的控制
<hr/>

#### fast_pretreatment    <gray>boolean</gray>
是否缩短 `prompt_optimizer` 的优化耗时，默认为 `false`。仅对 `MiniMax-Hailuo-02` 模型生效。
<hr/>

#### duration    <gray>int</gray>
视频时长（秒），默认值为 6。其可用值与模型和分辨率相关：

| 模型 |   720P |768P | 1080P |
| :--- | :--- | :--- |:--- |
| MiniMax-Hailuo-02 |`6` | `6` 或 `10` | `6` |
| 其他模型 |  `6` | - | - |
<hr/>

#### resolution    <gray>string</gray>
视频分辨率。其可用值与模型相关：

| 模型 | 6s | 10s |
| :--- | :--- | :--- |
| MiniMax-Hailuo-02 | `768P` (默认), `1080P` |  `768P` (默认) |
| 其他模型 | `720P` (默认) | 不支持 |
<hr/>


#### callback_url    <gray>string</gray>
接收任务状态更新通知的回调 URL。支持通过 `callback_url` 参数可以配置回调，以接收任务状态的更新的异步通知。

1. 地址验证：配置后，MiniMax 服务器会向 `callback_url` 发送一个 `POST` 请求，请求体中包含 `challenge` 字段。服务端需要在 3 秒内原样返回该 `challenge` 值以完成验证
2. 状态更新：验证成功后，每当任务状态变更时，MiniMax 都会向该 URL 推送最新的任务状态。推送的数据结构与调用查询视频生成任务接口的响应体一致


回调返回的“status”包括以下状态：
- “processing”-生成中
- “success”-成功
- “failed”-失败


<codetabs title="回调服务示例">
<codetab  language="python">
from fastapi import FastAPI, HTTPException, Request
from fastapi.middleware.cors import CORSMiddleware
import json
app = FastAPI()
@app.post("/get_callback")
async def get_callback(request: Request):
    try:
        json_data = await request.json()
        challenge = json_data.get("challenge")
        if challenge is not None:
          # is a verification request, just return the challenge
          return {"challenge": challenge}
        else:
            # is a callback request, do your own logic here
            # {
            #     "task_id": "115334141465231360",
            #     "status": "Success",
            #     "file_id": "205258526306433",
            #     "base_resp": {
            #         "status_code": 0,
            #         "status_msg": "success"
            #     }
            # }
            return {"status": "success"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        app, # 必选
        host="0.0.0.0", # 必选
        port=8000, # 必选，端口可设置
        # ssl_keyfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
        # ssl_certfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
    )
</codetab>

</codetabs>


<hr/>

#### aigc_watermark    <gray>boolean</gray>
是否在生成的视频中添加水印，默认为 `false`
<br/>

### 返回参数
<hr/>

#### task_id    <gray>string</gray>
视频生成任务的 ID，用于后续查询任务状态
<hr/>

#### base_resp    <gray>object</gray>
状态码及详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>integer</gray>
状态码及其分别含义如下：

0：请求成功

1002：触发限流，请稍后再试

1004：账号鉴权失败，请检查 API-Key 是否填写正确

1008：账号余额不足

1026：视频描述涉及敏感内容，请调整

2013：传入参数异常，请检查入参是否按要求填写

2049：无效的api key，请检查api key

更多内容可查看[错误码列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>

</left>

<right>
<codetabs title="文生视频">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/video_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "model": "MiniMax-Hailuo-02",
    "prompt": "男子拿起一本书[上升]，然后阅读[固定]。",
    "duration": 6,
    "resolution": "1080P",
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())

</codetab>
<codetab  language="curl">
curl -X POST "https://api.minimaxi.com/v1/video_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "model": "MiniMax-Hailuo-02",
    "prompt": "男子拿起一本书[上升]，然后阅读[固定]。",
    "duration": 6,
    "resolution": "1080P"
  }'
</codetab>

</codetabs>

<codetabs title="返回示例">
<codetab  language="json">
{
    "task_id": "106916112212032",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>

<container gap="4" columns="2">

  <card  href="/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成指南](/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7)

  </card>

<card  href="/document/video_api_intro?key=68abdfdaaf3fef26183d6501">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成接口能力介绍](/document/video_api_intro?key=68abdfdaaf3fef26183d6501)
  </card>

</container>

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/video_agent_api_intro?key=68ac00c56602726333ffcee4
- 文档ID: `68ac00c56602726333ffcee4`
# 视频生成Agent
本该接口支持基于用户选择的不同视频Agent模板和输入来进行视频生成任务。

## 接口说明
视频Agent接口采用异步方式，整体包含2个API：**创建视频Agent任务** 和 **查询视频Agent任务状态**。使用步骤如下：
1. 使用**创建视频Agent任务**接口，创建视频Agent任务，并得到task_id；
1. 使用**查询视频Agent任务状态**接口，基于task_id查询视频Agent任务状态；当状态为成功时，将获得对应的文件下载地址。

## 模板清单
详细信息可以参考 [视频Agent模板列表](/document/video_agent_templates?key=68b401edd2f0aaaf3484a10e)，查阅模板内容及示例效果。
| 模板ID | 模板名称 | 模板说明 | media_inputs | text_inputs |
| :--- | :--- | :--- | :--- | :--- |
| 392753057216684038 | 跳水 | 上传你的图片，生成图中主体完成完美跳水动作的视频 | 需要 | / |
| 393881433990066176 | 吊环 | 上传宠物照片，生成图中主体完成完美吊环动作的视频 | 需要 | / |
| 393769180141805569 | 绝地求生 | 上传宠物图片并输入野兽种类，生成宠物野外绝地求生视频 | 需要 | 需要 |
| 394246956137422856 | 万物皆可labubu | 上传人物/宠物照片，生成labubu换脸视频 | 需要 | / |
| 393879757702918151 | 麦当劳宠物外卖员 | 上传爱宠照片，生成麦当劳宠物外卖员视频 | 需要 | / |
| 393766210733957121 | 藏族风写真 | 上传面部参考图，生成藏族风视频写真 | 需要 | / |
| 394125185182695432 | 生无可恋 | 输入各类主角痛苦做某事，一键生成角色痛苦生活的小动画 | / | 需要 |
| 393857704283172864 | 情书写真 | 上传照片生成冬日雪景写真 | 需要 | / |
| 393866076583718914 | 女模特试穿广告 | 上传服装图片，生成女模特试穿广告 | 需要 | / |
| 398574688191234048 | 四季写真 | 上传人脸照片生成四季写真 | 需要 | / |
| 393876118804459526 | 男模特试穿广告 | 上传服装图片，生成男模特试穿广告 | 需要 | / |

## 视频Agent模板清单
- 原始链接: https://platform.minimaxi.com/document/video_agent_templates?key=68b401edd2f0aaaf3484a10e
- 文档ID: `68b401edd2f0aaaf3484a10e`
# 视频Agent 模板

参考以下表格，可了解视频Agent 接口下的全量视频模板：

| 模板ID | 名称 | 说明 | media_input | text_input | 效果 |
| :---: | :---------: | :---: | :---: | :---: | :---: |
| `392753057216684038` | 跳水 | 上传你的照片，生成照片中主主体完美跳水表现的视频 | 需要 | / | <video descriotion='跳水' width="300px"  src="https://filecdn.minimax.chat/public/434bb72c-3f55-4094-b06f-fb96bd41ddac.mp4" ></video> |
| `393881433990066176` | 吊环 | 上传宠物照片，生成图中主体完成完美吊环动作的视频 | 需要 |  / |<video descriotion='吊环'  width="300px"  src="https://filecdn.minimax.chat/public/4eba7e2b-ae58-4933-965e-3dbde901ed1f.mp4"></video> |
|`393769180141805569` | 绝地求生 | 上传宠物图片并输入野兽种类，生成宠物野外绝地求生视频 | 需要 | 需要 | <video descriotion='绝地求生'    width="300px"  src="https://filecdn.minimax.chat/public/ee7be27a-86e4-45ef-b1fb-829ea078624d.mp4"></video> |
| `394246956137422856` | 万物皆可labubu | 上传人物/宠物照片，生成labubu换脸视频 | 需要 |  / |<video descriotion='万物皆可labubu'  width="300px"  src="https://filecdn.minimax.chat/public/5d6cff91-b030-4c19-a80e-29cfed3ed56d.mp4"></video> |
| `393879757702918151` | 麦当劳宠物外卖员 | 上传爱宠照片，生成麦当劳宠物外卖员视频 | 需要 |  / |<video descriotion='麦当劳宠物外卖员'  width="300px"  src="https://filecdn.minimax.chat/public/1f8061fe-f885-4778-810f-5e3a4e148deb.mp4"></video> |
| `393766210733957121` | 藏族风写真 | 上传面部参考图，生成藏族风视频写真 | 需要 |  / |<video descriotion='藏族风写真'  width="300px"  src="https://filecdn.minimax.chat/public/b7a6e34a-84bd-4f90-81a2-d9495eb19ea1.mp4"></video> |
| `394125185182695432` | 生无可恋 | 输入各类主角痛苦做某事，一键生成角色痛苦生活的小动画 |  / | 需要 |<video descriotion='生无可恋'  width="300px"  src="https://filecdn.minimax.chat/public/4f21aa52-74bd-488f-b62f-ca03fcb6ed98.mp4"></video> |
| `393857704283172864` | 情书写真 | 上传照片生成冬日雪景写真 | 需要 |  / |<video descriotion='情书写真'  width="300px"  src="https://filecdn.minimax.chat/public/01f85d47-162b-4d97-856f-8ab3bf8b0101.mp4"></video> |
| `398574688191234048` | 四季写真 | 上传人脸照片生成四季写真 | 需要 | / |<video descriotion='四季写真'  width="300px"  src="https://filecdn.minimax.chat/public/571229bd-0e33-41be-80bb-716e30ba34f8.mp4"></video> |
| `393866076583718914` | 女模特试穿广告 | 上传服装图片，生成女模特试穿对应服装的广告 | 需要 |  / |<video descriotion='女模特试穿广告'  width="300px"  src="https://filecdn.minimax.chat/public/215dc60a-8987-4fab-9041-7e2e064b3eb7.mp4"></video> |
| `393876118804459526` | 男模特试穿广告 | 上传服装图片，生成男模特试穿对应服装的广告 | 需要 | / |<video descriotion='男模特试穿广告'  width="300px"  src="https://filecdn.minimax.chat/public/db76ea00-9919-43e9-9457-a0548430984c.mp4"></video> |

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/video_api_intro?key=68abdfdaaf3fef26183d6501
- 文档ID: `68abdfdaaf3fef26183d6501`
# 视频生成（Video Generation）
 本接口支持基于用户提供的文本、图片（包括首帧、尾帧、主体参考图）进行视频生成。


## 支持模型

| 模型 | 功能 |
| :--- | :--- |
| MiniMax-Hailuo-02 | 新一代视频生成模型，支持更高分辨率 (1080P) 和更长时长 (10s)，指令遵循能力更强 |
| T2V-01-Director | 文生视频模型（导演版），支持通过 prompt 中的指令精确控制镜头运动 |
| I2V-01-Director | 图生视频模型（导演版），基于参考图生成，并支持通过 prompt 中的指令精确控制镜头运动 |
| I2V-01-live | 图生视频模型，基于参考图生成，对卡通、漫画、手绘风格的表现有增强 |
| S2V-01 | 基于参考图中的人物面部生成视频，保持主体形象的高度一致性 |



## 接口说明

视频生成采用异步方式，整体包含3个API：**创建视频生成任务**、**查询视频生成任务状态**、**文件管理**。使用步骤如下：
1. 使用**创建视频生成任务**接口，创建视频生成任务，成功创建后会返回一个 `task_id`；

1. 使用**查询视频生成任务状态**接口，基于返回的task_id查询视频生成任务状态；当状态为成功时，将获得对应的文件 ID（file_id）；

1. 使用**文件管理**接口（[File API](#补充文件管理)）基于步骤2查询接口返回的file_id进行视频生成结果的查看和下载。


## 官方MCP

github链接：[https://github.com/MiniMax-AI/MiniMax-MCP](https://github.com/MiniMax-AI/MiniMax-MCP), 欢迎访问使用。

## 生成视频下载
- 原始链接: https://platform.minimaxi.com/document/video_generation_download?key=68ac00606fe587e3fbfe53b9
- 文档ID: `68ac00606fe587e3fbfe53b9`
# 视频下载
<request method="GET" url="https://api.minimaxi.com/v1/files/retrieve"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>Required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看
<hr/>

#### Content-Type    <gray>application/json</gray> <red>Required</red>
请求体的媒介类型，请设置为 `application/json`，确保请求数据的格式为 JSON

<br/>

### 请求体
<hr/>

#### file_id    <gray>int64</gray> <red>required</red>
文件的唯一标识符
支持视频任务状态查询接口获得的`file_id`

<br/>

### 返回体
<hr/>

#### file_id    <gray>int64</gray>
文件的唯一标识符
<hr/>

#### bytes    <gray>int64</gray>
文件大小，以字节为单位
<hr/>

#### created_at    <gray>int64</gray>
创建文件时的 Unix 时间戳，以秒为单位
<hr/>

#### filename    <gray>string</gray>
文件的名称
<hr/>

#### purpose    <gray>string</gray>
文件的使用目的
<hr/>

#### download_url    <gray>string</gray>
文件下载的url地址
<hr/>


#### base_resp    <gray>object</gray>
状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码如下：
- 1000, 未知错误
- 1001, 超时
- 1002, 触发RPM限流
- 1004, 鉴权失败
- 1008, 余额不足
- 1013, 服务内部错误
- 1026, 输入内容错误
- 1027, 输出内容错误
- 1039, 触发TPM限流
- 2013, 输入格式信息不正常

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
import requests


# 请填写你的api key
api_key = "请填写你的api key"
# 请填写你的file id
file_id = 请填写你的file id

url = f"https://api.minimaxi.com/v1/files/retrieve?file_id={file_id}"
headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {api_key}"
}

response = requests.get(url, headers=headers)
print(response.text)

</codetab>
<codetab language="curl">
curl --location  --request GET 'https://api.minimaxi.com/v1/files/retrieve?file_id=${file_id}' \
--header 'Authorization: Bearer ${MINIMAX_API_KEY}' \
--header 'content-type: application/json' \
--data-raw ''
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "file": {
        "file_id": ${file_id},
        "bytes": 5896337,
        "created_at": 1700469398,
        "filename": "output_aigc.mp4",
        "purpose": "video_generation",
        "download_url": "www.downloadurl.com"
    },
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>

<container gap="4" columns="2">

  <card  href="/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成指南](/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7)

  </card>

<card  href="/document/video_api_intro?key=68abdfdaaf3fef26183d6501">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成接口能力介绍](/document/video_api_intro?key=68abdfdaaf3fef26183d6501)
  </card>

</container>

## 查询视频生成任务状态
- 原始链接: https://platform.minimaxi.com/document/video_generation_query?key=68ac00276fe587e3fbfe53b1
- 文档ID: `68ac00276fe587e3fbfe53b1`
# 查询视频生成任务状态

<request method="GET" url="https://api.minimaxi.com/v1/query/video_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 

- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看

### 查询参数
<hr/>

#### task_id    <gray>string</gray> <red>required</red>
待查询的任务 ID。只能查询当前账号创建的任务

<br>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
被查询的任务 ID
<hr/>

#### status    <gray>string</gray>
任务的当前状态。可能的状态值包括：

Preparing-准备中

Queueing-队列中

Processing-生成中

Success-成功

Fail-失败
<hr/>

#### file_id    <gray>string</gray>
任务成功时返回，用于获取视频文件的文件 ID
<hr/>

#### video_width    <gray>int</gray>
任务成功时返回，生成视频的宽度（像素）
<hr/>

#### video_height    <gray>int</gray>
任务成功时返回，生成视频的高度（像素）
<hr/>

#### base_resp    <gray>object</gray>
状态码及状态详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>integer</gray>
状态码及其分别含义如下：

0：请求成功

1002：触发限流，请稍后再试

1004：账号鉴权失败，请检查 api key是否填写正确

1026：输入内容涉及敏感内容

1027：生成视频涉及敏感内容

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态信息，成功时为 `success`

</card>
</details>

</left>

<right>
<codetabs title="查询任务状态">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/query/video_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

task_id = os.environ["TASK_ID"]
params = {"task_id": task_id}

response = requests.get(url, headers=headers, params=params)
response.raise_for_status()
print(response.json())

</codetab>
<codetab  language="curl">
curl -X GET "https://api.minimaxi.com/v1/query/video_generation?task_id=$TASK_ID" \
  -H "Authorization: Bearer $MINIMAX_API_KEY"
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab  language="json">
{
  "task_id": "176843862716480",
  "status": "Success",
  "file_id": "176844028768320",
  "video_width": 1920,
  "video_height": 1080,
  "base_resp": {
    "status_code": 0,
    "status_msg": "success"
  }
}

</codetab>
</codetabs>

</right>

</container>

<br>

<container gap="4" columns="2">

  <card  href="/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成指南](/document/guides_video_generation?key=68bbf01a95a600dbfc9266f7)

  </card>

<card  href="/document/video_api_intro?key=68abdfdaaf3fef26183d6501">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频生成接口能力介绍](/document/video_api_intro?key=68abdfdaaf3fef26183d6501)
  </card>

</container>

## 创建视频Agent任务
- 原始链接: https://platform.minimaxi.com/document/video_template_generation?key=68ac01306fe587e3fbfe55cb
- 文档ID: `68ac01306fe587e3fbfe55cb`
# 创建视频Agent任务
<request method="POST" url="https://api.minimaxi.com/v1/video_template_generation"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看。
<hr/>

#### Content-Type    <gray>application/json</gray> <red>required</red>
Content-type

### 请求体
<hr/>

#### template_id    <gray>string</gray> <red>required</red>
视频模板的 ID。具体的 ID 和所需输入参见 [视频模板列表](/document/video_agent_templates?key=68b401edd2f0aaaf3484a10e)。
<hr/>

#### text_inputs    <gray>array</gray>
文本输入数组，用于填充模板中的文本部分，不同模板对此要求不同
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### value    <gray>string</gray> <red>required</red>
具体的文本内容

</card>
</details>
<hr/>

#### media_inputs    <gray>array</gray>
媒体输入数组（如图片），用于填充模板中的媒体部分，不同模板对此要求不同
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### value    <gray>string</gray> <red>required</red>
图像文件。支持公网 URL 或 Base64 编码的 Data URL (`data:image/jpeg;base64,...`)。
- 图片要求：
  - 格式：JPG, JPEG, PNG, WebP
  - 大小：小于 20MB
  - 尺寸：短边像素大于 300px
  - 宽高比：在 2:5 到 5:2 之间

</card>
</details>
<hr/>

#### callback_url    <gray>string</gray>
接收任务状态更新通知的回调 URL。支持通过 `callback_url` 参数可以配置回调，以接收任务状态的更新的异步通知

1. 地址验证：配置后，MiniMax 服务器会向 `callback_url` 发送一个 `POST` 请求，请求体中包含 `challenge` 字段。服务端需要在 3 秒内原样返回该 `challenge` 值以完成验证
2. 状态更新：验证成功后，每当任务状态变更时，MiniMax 都会向该 URL 推送最新的任务状态。推送的数据结构与调用查询视频生成任务接口的响应体一致


回调返回的"status"包括以下状态：
- "processing"-生成中
- "success"-成功
- "failed"-失败


<codetabs title="回调服务示例">
<codetab  language="python">
from fastapi import FastAPI, HTTPException, Request
from fastapi.middleware.cors import CORSMiddleware
import json
app = FastAPI()
@app.post("/get_callback")
async def get_callback(request: Request):
    try:
        json_data = await request.json()
        challenge = json_data.get("challenge")
        if challenge is not None:
          # is a verification request, just return the challenge
          return {"challenge": challenge}
        else:
            # is a callback request, do your own logic here
            # {
            #     "task_id": "115334141465231360",
            #     "status": "Success",
            #     "file_id": "205258526306433",
            #     "base_resp": {
            #         "status_code": 0,
            #         "status_msg": "success"
            #     }
            # }
            return {"status": "success"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        app, # 必选
        host="0.0.0.0", # 必选
        port=8000, # 必选，端口可设置
        # ssl_keyfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
        # ssl_certfile="yourname.yourDomainName.com.key", # 可选，看是否开启ssl
    )
</codetab>

</codetabs>




<br>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
任务的唯一 ID，可用于后续查询任务状态
<hr/>

#### base_resp    <gray>object</gray>
状态码及状态详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0：请求成功

1002：触发限流，请稍后再试

1004：账号鉴权失败，请检查 API-Key 是否填写正确

1008：账号余额不足

1026：文本涉及敏感内容，请调整

2013：传入参数异常，请检查入参是否按要求填写

2049：无效的api key，请检查api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
具体错误详情

</card>
</details>



</left>

<right>
<codetabs title="请求示例">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/video_template_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

payload = {
    "template_id": "393769180141805569",  # 绝地求生
    "media_inputs": [
        {
            "value": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg"
        }
    ],
    "text_inputs": [{"value": "狮子"}],
}

response = requests.post(url, headers=headers, json=payload)
response.raise_for_status()
print(response.json())

</codetab>

<codetab  language="curl">
curl -X POST "https://api.minimaxi.com/v1/video_template_generation" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MINIMAX_API_KEY" \
  -d '{
    "template_id": "393769180141805569",
    "media_inputs": [
      {
        "value": "https://cdn.hailuoai.com/prod/2024-09-18-16/user/multi_chat_file/9c0b5c14-ee88-4a5b-b503-4f626f018639.jpeg"
      }
    ],
    "text_inputs": [{"value": "狮子"}]
  }'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab  language="json">

{
    "task_id": "401047179385389059",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>

<container gap="4" columns="2">

  <card  href="/document/guides_video_template_generation?key=68bbf0494b8efaab4d9f85ec">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[使用模板视频生成指南](/document/guides_video_template_generation?key=68bbf0494b8efaab4d9f85ec)

  </card>

<card  href="/document/video_agent_api_intro?key=68ac00c56602726333ffcee4">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频Agent接口能力介绍](/document/video_agent_api_intro?key=68ac00c56602726333ffcee4)
  </card>

</container>

## 查询视频Agent任务状态
- 原始链接: https://platform.minimaxi.com/document/video_template_query?key=68ac017f6fe587e3fbfe55d3
- 文档ID: `68ac017f6fe587e3fbfe55d3`
# 查询视频Agent任务状态

<request method="POST" url="https://api.minimaxi.com/v1/query/video_template_generation"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 

- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看

<br>

### 查询参数
<hr/>

#### task_id    <gray>string</gray> <red>required</red>
待查询的任务 ID。只能查询当前账号创建的任务

<br>

### 返回体
<hr/>

#### task_id    <gray>string</gray>
被查询的任务 ID
<hr/>

#### status    <gray>string</gray>
任务状态，可能的值包括：

- Preparing-准备中

- Processing-生成中

- Success-成功

- Fail-失败
<hr/>

#### video_url    <gray>string</gray>
任务成功时返回。生成视频的下载链接。注意：下载链接有效期为 9 小时
<hr/>

#### base_resp    <gray>object</gray>
状态码及状态详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0：请求成功

1002：触发限流，请稍后再试

1004：账号鉴权失败，请检查 api key是否填写正确

1026：输入内容涉及敏感内容

1027：生成视频涉及敏感内容

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态信息，成功时为 `success`

</card>
</details>

</left>

<right>
<codetabs title="查询任务状态">
<codetab  language="python">
import requests
import os

url = "https://api.minimaxi.com/v1/query/video_template_generation"
api_key = os.environ["MINIMAX_API_KEY"]
headers = {"Authorization": f"Bearer {api_key}"}

task_id = os.environ["TASK_ID"]
params = {"task_id": task_id}

response = requests.get(url, headers=headers, params=params)
response.raise_for_status()
print(response.json())

</codetab>
<codetab  language="bash">
curl -X GET "https://api.minimaxi.com/v1/query/video_template_generation?task_id=$TASK_ID" \
  -H "Authorization: Bearer $MINIMAX_API_KEY"
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab  language="json">
{
    "task_id": "401047179385389059",
    "status": "Success",
    "video_url": "https://cdn.hailuoai.com/prod/video_20250714_141232_cdc5ba74.mp4",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>

<container gap="4" columns="2">

  <card  href="/document/guides_video_template_generation?key=68bbf0494b8efaab4d9f85ec">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[使用模板视频生成指南](/document/guides_video_template_generation?key=68bbf0494b8efaab4d9f85ec)

  </card>

<card  href="/document/video_agent_api_intro?key=68ac00c56602726333ffcee4">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[视频Agent接口能力介绍](/document/video_agent_api_intro?key=68ac00c56602726333ffcee4)
  </card>

</container>

## 文生音色
- 原始链接: https://platform.minimaxi.com/document/vincent_timbre?key=ddcd5dedd5434afe9b7d8a65
- 文档ID: `ddcd5dedd5434afe9b7d8a65`
### Voice Generation（文生音色）
该API支持基于用户输入的声音性别、声音年龄、声音描述信息，来生成音色（voice_id）；并支持使用该生成的音色（voice_id）在T2A v2、T2A Large v2接口中进行语音生成。
该接口支持以下功能：
1. 支持指定音色性别（男、女两种）、声音年龄（小孩到老人，五种年龄段）；
2. 支持对音色风格进行自定义描述；
3. 支持根据自定义文本内容生成音频，对合成音色进行试听。
该接口的适用场景：基于文本描述，生成符合描述的个性化定制音色，以供语音生成接口使用。
#### API说明
`POST: /v1/text2voice`
###### 请求体（Request）参数
###### 返回（Response）参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| gender | enum<string> | 是 | 声音的性别。取值可选，1.male、2.female。 |
| age | enum<string> | 是 | 声音的年龄。取值可选：1.child、2.teenager、3.young、4.middle-aged、5.old。 |
| voice_desc | list[string] | 是 | 自定义输入音色风格描述词，比如['casual','fashion']，描述词需少于等于8个。 |
| text | string | 是 | 试听音频文本。 |
| voice_id | string | 否 | 生成的音色ID。 |
| trial_audio | string | 否 | 使用生成的音色合成的示例音频。 |
| base_resp |  | 否 | 状态码和状态详情。<br>状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。<br>错误详情。 |
| base_resp.status_code | in64 | 否 | 状态码。1000，未知错误；1001，超时；1002，触发RPM限流；1004，鉴权失败；1008，余额不足；1013，服务内部错误；1027，输出内容错误；1039，触发TPM限流；2013，输入格式信息不正常。 |
| base_resp.status_msg | string | 否 | 错误详情。 |

## 快速复刻
- 原始链接: https://platform.minimaxi.com/document/voice_clone?key=68adaf2f6fe587e3fbfe9b89
- 文档ID: `68adaf2f6fe587e3fbfe9b89`
# 音色快速复刻

<request method="POST" url="https://api.minimaxi.com/v1/voice_clone"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>Required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看

<br/>

### 请求体
<hr/>

#### file_id    <gray>int64</gray> <red>Required</red>
待复刻音频的 file_id，通过[文件上传接口](/document/files_upload?key=68b2aedbfd643cbdfcc0611b)获得
上传的待复刻音频文件需遵从以下规范：
- 上传的音频文件格式需为：mp3、m4a、wav 格式
- 上传的音频文件的时长最少应不低于 10 秒，最长应不超过 5 分钟
- 上传的音频文件大小需不超过 20 mb

<hr/>

#### voice_id    <gray>string</gray> <red>Required</red>
克隆音色的 voice_id，正确示例："MiniMax001"。用户进行自定义 voice_id 时需注意：
1. 自定义的 voice_id 长度范围[8,256]
1. 首字符必须为英文字母
1. 允许数字、字母、-、_
1. 末位字符不可为 -、_
1. voice_id 不可与已有 id 重复，否则会报错
<hr/>

#### clone_prompt    <gray>object</gray>
音色复刻示例音频，提供本参数将有助于增强语音合成的音色相似度和稳定性。若使用本参数，需同时上传一小段示例音频
上传的音频文件需遵从以下规范：
- 上传的音频文件格式需为：mp3、m4a、wav 格式
- 上传的音频文件的时长小于 8 秒
- 上传的音频文件大小需不超过 20 mb
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### prompt_audio    <gray>int64</gray>
示例音频的 file_id，通过[文件上传接口](/document/files_upload?key=68b2aedbfd643cbdfcc0611b)获得

<hr/>

#### prompt_text    <gray>string</gray>
示例音频的对应文本，需确保和音频内容一致，句末需有标点符号做结尾

</card>
</details>
<hr/>

#### text    <gray>string</gray>
复刻试听参数，限制 2000 字符以内。模型将使用复刻后的音色朗读本段文本内容，并返回试听音频链接。注：试听将根据字符数正常收取语音合成费用，定价与 T2A 各接口一致
<hr/>


#### model    <gray>string</gray>
复刻试听参数。指定合成试听音频使用的语音模型，提供"text"字段时必传此字段。可选项： `speech-2.5-hd-preview` `speech-2.5-turbo-preview` `speech-02-hd` `speech-02-turbo` `speech-01-hd` `speech-01-turbo`
<hr/>

#### language_boost    <gray>string</gray>
是否增强对指定的小语种和方言的识别能力。默认值为 null，可设置为 "auto" 让模型自主判断。
可选值范围：[`Chinese`, `Chinese,Yue`, `English`, `Arabic`, `Russian`, `Spanish`, `French`, `Portuguese`, `German`, `Turkish`, `Dutch`, `Ukrainian`, `Vietnamese`, `Indonesian`, `Japanese`, `Italian`, `Korean`, `Thai`, `Polish`, `Romanian`, `Greek`, `Czech`,  `Finnish`, `Hindi`, `Bulgarian`, `Danish`, `Hebrew`, `Malay`, `Persian`, `Slovak`, `Swedish`, `Croatian`, `Filipino`, `Hungarian`, `Norwegian`, `Slovenian`, `Catalan`, `Nynorsk`, `Tamil`, `Afrikaans`, `auto`]
<hr/>



#### need_noise_reduction    <gray>bool</gray>
音频复刻参数，表示是否开启降噪，默认值为 false
<hr/>

#### need_volume_normalization    <gray>bool</gray>
音频复刻参数，是否开启音量归一化，默认值为 false
<hr/>

#### aigc_watermark    <gray>bool</gray>
是否在合成试听音频的末尾添加音频节奏标识，默认值为 false

<br/>

### 返回体
<hr/>

#### input_sensitive    <gray>bool</gray>
输入音频是否命中风控
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### input_sensitive_type    <gray>int</gray>
输入音频命中风控的类型，取值为以下其一：
- 0：正常
- 1：严重违规
- 2：色情
- 3：广告
- 4：违禁
- 5：谩骂
- 6：暴恐
- 7：其他

</card>
</details>
<hr/>

#### demo_audio    <gray>string</gray>
如果请求体中传入了试听文本 `text` 以及合成试听音频的模型 `model`，那么本参数将以链接形式返回试听音频，否则本参数为空值
<hr/>

#### base_resp    <gray>dict</gray>
本次请求的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码
- 0: 请求结果正常
- 1000：未知错误
- 1001：超时
- 1002：触发限流
- 1004：鉴权失败
- 1013：服务内部错误
- 2013：输入格式信息不正常
- 2038：无复刻权限，请检查账号认证状态

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
"""
本示例用于快速克隆。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`
"""
import requests
import json
import os

api_key = os.getenv("MINIMAX_API_KEY")
url = 'https://api.minimaxi.com/v1/voice_clone'

payload = json.dumps({
  "file_id": <file_id of cloned voice>,
  "voice_id": "<voice_id>",
  "clone_prompt":{
    "prompt_audio": <file_id of the prompt audio>,
    "prompt_text": "这个声音效果挺自然的，听起来很舒服。",
  },
  "text":"微风拂过柔软的草地，清新的芳香伴随着鸟儿的歌唱。",
  "model":"speech-2.5-hd-preview",
  "need_noise_reduction": False,
  "need_volumn_normalization": False,
  "aigc_watermark": False,
})
headers = {
  'Authorization': f'Bearer {api_key}',
  'content-type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/voice_clone' \
--header 'authority: api.minimaxi.com' \
--header "Authorization: Bearer ${MINIMAX_API_KEY}" \
--header 'content-type: application/json' \
--data '{
    "file_id": <file_id of cloned voice>,
    "voice_id": "<voice_id>",
    "clone_prompt":{
      "prompt_audio": <file_id of the prompt audio>,
      "prompt_text": "这个声音效果挺自然的，听起来很舒服。",
    },
    "text":"微风拂过柔软的草地，清新的芳香伴随着鸟儿的歌唱。",
    "model":"speech-2.5-hd-preview",
    "need_noise_reduction": false,
    "need_volumn_normalization": false,
    "aigc_watermark": false
}'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "input_sensitive": false,
    "input_sensitive_type": 0,
    "demo_audio":"",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<container gap="4" columns="2">

  <card  href="/document/guides_voice_clone1?key=68bbe68e4b8efaab4d9f722f">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[音色快速复刻指南](/document/guides_voice_clone1?key=68bbe68e4b8efaab4d9f722f)

  </card>

<card  href="/document/Voice_Cloning_api_intro?key=68adaebe66027263330017e0">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[音色快速复刻接口能力介绍](/document/Voice_Cloning_api_intro?key=68adaebe66027263330017e0)
  </card>

</container>

## 上传复刻音频
- 原始链接: https://platform.minimaxi.com/document/voice_clone_upload_file?key=68adaef16fe587e3fbfe9aa5
- 文档ID: `68adaef16fe587e3fbfe9aa5`
# 上传复刻音频


<request method="POST" url="https://api.minimaxi.com/v1/files/upload"></request>


<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth`
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看。
<hr/>

#### Content-Type    <gray>multipart/form-data</gray> <red>required</red>
请求体的媒介类型，请设置为 `multipart/form-data`

<br/>

### 请求体
<hr/>


#### purpose    <gray>string</gray> <red>required</red>
文件使用目的。取值及支持格式如下：
`voice_clone`: 快速复刻原始文件，（支持mp3、m4a、wav格式）
<hr/>

#### file <red>required</red>
需要上传的文件。填写文件的路径地址。
- 支持上传的文件需遵从以下规范：
上传的音频文件格式需为：mp3、m4a、wav格式
上传的音频文件的时长最少应不低于10秒，最长应不超过5分钟
上传的音频文件大小需不超过20mb

<br/>

### 返回体
<hr/>

#### file_id    <gray>int64</gray>
文件的唯一标识符
<hr/>

#### bytes    <gray>int64</gray>
文件大小，以字节为单位
<hr/>

#### created_at    <gray>int64</gray>
创建文件时的 Unix 时间戳，以秒为单位
<hr/>

#### filename    <gray>string</gray>
文件的名称
<hr/>

#### purpose    <gray>string</gray>
文件的使用目的，voice_clone
<hr/>

#### base_resp    <gray>object</gray>
状态码和详情。
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int</gray>
状态码及其分别含义如下：

0，请求成功

1002，触发限流，请稍后再试

1004，账号鉴权失败，请检查 API-Key 是否填写正确

1008，账号余额不足

1026，图片描述涉及敏感内容

2013，传入参数异常，请检查入参是否按要求填写

2049，无效的api key

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
url = f'https://api.minimaxi.com/v1/files/upload?GroupId={group_id}'
headers1 = {
    'authority': 'api.minimaxi.com',
    'Authorization': f'Bearer {api_key}'
}

data = {
    'purpose': 'voice_clone'
}

files = {
    'file': open('output.mp3', 'rb')
}
response = requests.post(url, headers=headers1, data=data, files=files)
file_id = response.json().get("file").get("file_id")
print(file_id)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/files/upload?GroupId={group_id}' \
--header 'authority: api.minimaxi.com' \
--header 'content-type: multipart/form-data'  \
--header 'Authorization: Bearer MiniMax_API_KEY' \
--form 'purpose="voice_clone"' \
--form 'file=@"Users/minimax/Downloads/音频文件.mp3"'

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "file": {
        "file_id": "${file_id}",
        "bytes": 5896337,
        "created_at": 1700469398,
        "filename": "复刻音频",
        "purpose": "voice_clone"
    },
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<container gap="4" columns="2">

  <card  href="/document/guides_voice_clone1?key=68bbe68e4b8efaab4d9f722f">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[音色快速复刻指南](/document/guides_voice_clone1?key=68bbe68e4b8efaab4d9f722f)

  </card>

<card  href="/document/Voice_Cloning_api_intro?key=68adaebe66027263330017e0">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读
[音色快速复刻接口能力介绍](/document/Voice_Cloning_api_intro?key=68adaebe66027263330017e0)
  </card>

</container>

## 音色设计
- 原始链接: https://platform.minimaxi.com/document/voice_design_?key=68adb7976602726333001efc
- 文档ID: `68adb7976602726333001efc`
# 音色设计（Voice Design）
<request method="POST" url="https://api.minimaxi.com/v1/voice_design"></request>



<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
`HTTP：Bearer Auth` 
- Security Scheme Type: http
- HTTP Authorization Scheme: Bearer API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看

### 请求体
<hr/>

#### prompt    <gray>string</gray> <red>required</red>
音色描述。
<hr/>

#### preview_text    <gray>string</gray> <red>required</red>
试听音频文本。
注：试听音频的合成将收取 2 元/万字符的费用
<hr/>

#### voice_id    <gray>string</gray>
自定义生成音色的 voice_id。当不传入此参数时，将自动生成并返回一个唯一的`voice_id`
<hr/>

#### aigc_watermark    <gray>bool</gray>
是否在合成试听音频的末尾添加音频节奏标识，默认值为 False

### 返回体
<hr/>

#### voice_id    <gray>string</gray>
生成的音色 ID，可用于语音合成
<hr/>

#### trial_audio    <gray>string</gray>
使用生成的音色合成的试听音频，以 hex 编码形式返回
<hr/>

#### base_resp    <gray>object</gray>
状态码和状态详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码。
- 0: 请求结果正常
- 1000：未知错误
- 1001：超时
- 1002：触发 RPM 限流
- 1004：鉴权失败
- 1008：余额不足
- 1013：服务内部错误
- 1027：输出内容错误
- 1039：触发 TPM 限流
- 2013：输入格式信息不正常

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="python">
"""
本示例用于音色设计。注意：需要先将密钥信息写入环境变量 `MINIMAX_API_KEY`。
"""
import requests
import json
import os

api_key = os.getenv("MINIMAX_API_KEY")
url = "https://api.minimaxi.com/v1/voice_design"
payload = json.dumps({
  "prompt": "讲述悬疑故事的播音员，声音低沉富有磁性，语速时快时慢，营造紧张神秘的氛围。",
  "preview_text": "夜深了，古屋里只有他一人。窗外传来若有若无的脚步声，他屏住呼吸，慢慢地，慢慢地，走向那扇吱呀作响的门……"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': f'Bearer {api_key}',
}

response = requests.request("POST", url, headers=headers, data=payload)
parsed_json = json.loads(response.text)

audio_value = bytes.fromhex(parsed_json['trial_audio'])
with open('output.mp3', 'wb') as f:
    f.write(audio_value)
voice_id = parsed_json['voice_id']
print(voice_id)

</codetab>
<codetab language="curl">
curl --location 'https://api.minimaxi.com/v1/voice_design' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer ${MINIMAX_API_KEY}" \
--data '{
    "prompt":"讲述悬疑故事的播音员，声音低沉富有磁性，语速时快时慢，营造紧张神秘的氛围。",
    "preview_text":"夜深了，古屋里只有他一人。窗外传来若有若无的脚步声，他屏住呼吸，慢慢地，慢慢地，走向那扇吱呀作响的门……"
}' \
--output response.json

jq -r '.trial_audio' response.json | xxd -r -p > output0.mp3

rm response.json

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "trial_audio": "hex 编码音频",
    "voice_id": "ttv-voice-2025060717322425-xxxxxxxx",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

## 接口能力介绍
- 原始链接: https://platform.minimaxi.com/document/voice_design_api_intro?key=68adb7586fe587e3fbfea0e2
- 文档ID: `68adb7586fe587e3fbfea0e2`
# 音色设计（Voice Design）

该 API 支持基于用户输入的声音描述 prompt，生成个性化定制音色；并支持使用该生成的音色（voice_id）在[同步语音合成接口](/document/t2a_api_intro?key=68ad77666602726333000457)和[异步长文本语音合成接口](/document/t2a_async_api_intro?key=68adac446fe587e3fbfe965b)中进行语音生成

## 支持模型

>推荐使用 speech-02-hd 以获得最佳效果

| 模型 | 特性 |
| ------ | ------ |
| speech-2.5-hd-preview | 最新的 HD 模型，韵律表现出色，复刻相似度极佳 |
| speech-2.5-turbo-preview | 最新的 Turbo 模型，支持 40 个语种 |
| speech-02-hd | 拥有出色的韵律、稳定性和复刻相似度，音质表现突出 |
| speech-02-turbo | 拥有出色的韵律和稳定性，小语种能力加强，性能表现出色 |
| speech-01-hd | 拥有超高的复刻相似度，音质表现突出 |
| speech-01-turbo | 在出色的生成效果基础上有更快的生成速度 |


## 注意事项 

>- 调用本接口获得音色时，不会立即收取生成音色的费用，生成音色的费用将在首次使用此音色进行语音合成时收取（不包含本接口内的试听行为）。
>- 本接口产出的音色为临时音色，如您希望永久保留某音色，请于 168 小时（7 天）内在任意语音合成接口中调用该音色（不包含本接口内的试听行为），超过有效期未被使用的音色将自动删除。

## 建立连接
- 原始链接: https://platform.minimaxi.com/document/wss_connect?key=68bda47a33216f06cecd67a4
- 文档ID: `68bda47a33216f06cecd67a4`
# 建立连接
调用 WebSocket 库函数（具体实现方式因编程语言或库函数而异），将请求头和 URL 传入以建立 WebSocket 连接。


<request method="WSS" url="wss://api.minimaxi.com/ws/v1/t2a_v2"></request>

<container>

<left>

### 请求头
<hr/>

#### Authorization    <gray>string</gray> <red>required</red>
API_key，可在 [账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key) 中查看

### 返回体
<hr/>

#### session_id    <gray>string</gray>
表示整个会话的 id
<hr/>

#### event    <gray>string</gray>
表示会话事件类型，建连成功后会返回"connected_success"
<hr/>

#### trace_id    <gray>string</gray>
表示会话中单次请求的 id，用于在咨询/反馈时帮助定位问题
<hr/>

#### base_resp    <gray>dict</gray>
该请求对应的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码，"0"代表建连成功

更多状态码信息请参考[错误码查询](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="建立连接">
<codetab language="json">

{
    "Authorization": "Bearer " + "<your_api_key>"
}

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">

{
    "session_id":"xxxx",
    "event":"connected_success",
    "trace_id":"0303a2882bf18235ae7a809ae0f3cca7",
    "base_resp":{
        "status_code": 0,
        "status_msg":"success"
    }
}

</codetab>
</codetabs>

</right>

</container>

<br>



<container gap="4" columns="2">

  <card  href="/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成指南 websocket](/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba)

  </card>

<card  href="/document/t2a_api_intro?key=68ad77666602726333000457">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成接口能力介绍](/document/t2a_api_intro?key=68ad77666602726333000457)

  </card>

</container>

## 发送“任务继续”事件
- 原始链接: https://platform.minimaxi.com/document/wss_continue?key=68bda59b33216f06cecd77b2
- 文档ID: `68bda59b33216f06cecd77b2`
# 发送“任务继续”事件
当收到服务端返回的 `task_started` 事件后，任务正式开始，可通过发送 `task_continue` 事件发送要合成的文本，支持顺序发送多个 `task_continue` 事件。
当最后一次收到服务端返回结果后超过 120s 没有发送新事件时，webSocket 连接自动断开。

<container>

<left>

### 请求体
<hr/>

#### event    <gray>string</gray> <red>Required</red>
表示会话事件类型，当前环节应填写 "task_continue"
<hr/>

#### text    <gray>string</gray> <red>required</red>
需要合成语音的文本，长度限制小于 10,000 字符
- 段落切换用换行符标记
- 停顿控制：支持自定义文本之间的语音时间间隔。使用方式：在文本中增加`<#x#>`标记，`x` 为停顿时长（单位：秒），范围 **[0.01, 99.99]**，最多保留两位小数。文本间隔时间需设置在两个可以语音发音的文本之间，不可连续使用多个停顿标记

### 返回体
<hr/>

#### data    <gray>object</gray>
data 可能返回为 null，参考示例代码时，注意进行非空判断
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### audio    <gray>string</gray>
合成后的音频片段，采用 hex 编码，按照输入定义的格式进行生成（mp3/pcm/flac）

</card>
</details>
<hr/>

#### trace_id    <gray>string</gray>
表示会话中单次请求的 id，用于在咨询/反馈时帮助定位问题。
<hr/>

#### session_id    <gray>string</gray>
表示整个会话的 id。
<hr/>

#### event    <gray>string</gray>
表示会话类型，当前环节成功后会返回 "task_continued"
<hr/>

#### is_final    <gray>bool</gray>
该请求返回是否完结
<hr/>

#### extra_info    <gray>object</gray>
相关额外信息
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### audio_length    <gray>int64</gray>
音频时长，精确到毫秒
<hr/>

#### audio_sample_rate    <gray>int64</gray>
音频采样率
<hr/>

#### audio_size    <gray>int64</gray>
音频文件大小，单位为字节
<hr/>

#### bitrate    <gray>int64</gray>
音频比特率
<hr/>

#### audio_format    <gray>string</gray>
生成音频文件的格式。取值范围 mp3/pcm/flac
<hr/>

#### audio_channel    <gray>int64</gray>
生成音频声道数。1：单声道，2：双声道
<hr/>

#### invisible_character_ratio    <gray>float64</gray>
非法字符占比。非法字符不超过 10%（包含 10%），音频会正常生成并返回非法字符占比，超过进行报错
<hr/>

#### usage_characters    <gray>int64</gray>
计费字符数。本次语音生成的计费字符数
<hr/>

#### word_count    <gray>int64</gray>
已发音的字数统计，包含汉字、数字、字母，不包含标点符号

</card>
</details>
<hr/>

#### base_resp    <gray>dict</gray>
本次请求的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码
- 0: 请求结果正常
- 1000：未知错误
- 1001：超时
- 1002：触发限流
- 1004：鉴权失败
- 1039：触发 TPM 限流
- 1042：非法字符超过 10%
- 2013：输入参数信息不正常
- 2201：超时断开连接
- 2202：非法事件
- 2203：空文本，跳过
- 2204：超过字符限制，跳过
- 2205：请求超限

更多内容可查看[错误码查询列表](/document/error_codes?key=68ad64216fe587e3fbfe7acf)了解详情

<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="json">
{
    "event": "task_continue",
    "text": "真正的危险不是计算机开始像人一样思考，而是人开始像计算机一样思考。计算机只是可以帮我们处理一些简单事务。"
}
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "data": {
        "audio": "xxx"
    },
    "extra_info": {
      "audio_channel": 1,
      "audio_format": "mp3",
      "audio_length": 10472,
      "audio_sample_rate": 32000,
      "audio_size": 169908,
      "bitrate": 128000,
      "invisible_character_ratio": 0,
      "usage_characters": 101,
      "word_count": 52
    },
    "is_final": true,
    "session_id": "301871346491491",
    "trace_id": "04ee3794e2c9e4a6d5f99e77742f06fd",
    "base_resp": {
      "status_code": 0, 
      "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>


<container gap="4" columns="2">

  <card  href="/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成指南 websocket](/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba)

  </card>

<card  href="/document/t2a_api_intro?key=68ad77666602726333000457">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成接口能力介绍](/document/t2a_api_intro?key=68ad77666602726333000457)

  </card>

</container>

## 发送“任务结束”事件
- 原始链接: https://platform.minimaxi.com/document/wss_finish?key=68bda5f59f4248575d21767b
- 文档ID: `68bda5f59f4248575d21767b`
# 发送“任务结束”事件
服务端收到 `task_finish` 事件后，会等待当前队列中所有合成任务完成后，关闭 WebSocket 连接并结束任务。

<container>

<left>

### 请求体
<hr/>

#### event    <gray>string</gray> <red>required</red>
表示会话事件类型，当前环节应填写 "task_finish"

<br>

### 返回体
<hr/>

#### trace_id    <gray>string</gray>
表示会话中单次请求的 id，用于在咨询/反馈时帮助定位问
<hr/>

#### session_id    <gray>string</gray>
表示整个会话的 id
<hr/>

#### event    <gray>string</gray>
表示会话类型，当前环节成功后会返回 "task_finished"
<hr/>

#### base_resp    <gray>dict</gray>
本次请求的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码
- 0 表示发送成功
- 2202 表示非法事件

更多状态码信息请参考[错误码查询](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="json">
{
    "event": "task_finish"
}
</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">
{
    "session_id": "xxxx",
    "event": "task_finished",
    "trace_id": "0303a2882bf18235ae7a809ae0f3cca7",
    "base_resp": {
        "status_code": 0,
        "status_msg": "success"
    }
}
</codetab>
</codetabs>

</right>

</container>

<br>



<container gap="4" columns="2">

  <card  href="/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成指南 websocket](/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba)

  </card>

<card  href="/document/t2a_api_intro?key=68ad77666602726333000457">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成接口能力介绍](/document/t2a_api_intro?key=68ad77666602726333000457)

  </card>

</container>

## 发送“任务开始”事件
- 原始链接: https://platform.minimaxi.com/document/wss_start?key=68bda50e9f4248575d216987
- 文档ID: `68bda50e9f4248575d216987`
# 发送“任务开始”事件
发送"任务开始"事件则正式开始合成任务，当服务端返回的 task_started 事件时，标志着任务已成功开始。
只有在接收到该事件后，才能向服务器发送 `task_continue` 事件或 `task_finish` 事件。

<container>

<left>

### 请求体
<hr/>

#### event    <gray>string</gray> <red>required</red>
控制发送的指令，当前环节应填写："task_start"
<hr/>

#### model    <gray>string</gray> <red>required</red>
请求的模型版本，可选范围： `speech-2.5-hd-preview` , `speech-2.5-turbo-preview` , `speech-02-hd` , `speech-02-turbo` , `speech-01-hd`, `speech-01-turbo` 
<hr/>

#### voice_setting    <gray>object</gray> <red>required</red>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray> <red>required</red>
合成音频的音色编号。若需要设置混合音色，请设置 timbre_weights 参数，本参数设置为空值。支持系统音色、复刻音色以及文生音色三种类型，以下是部分最新的系统音色（ID），可查看 [音色列表](/document/system_voice_id?key=68b40964a96516e26018eee2) 或使用 [查询可用声音接口](/document/get_voice?key=68b51016d2f0aaaf3484a8de) ，来查询系统支持的全部音色。
- **中文**：
  - moss_audio_ce44fc67-7ce3-11f0-8de5-96e35d26fb85
  - moss_audio_aaa1346a-7ce7-11f0-8e61-2e6e3c7ee85d
  - Chinese (Mandarin)_Lyrical_Voice
  - Chinese (Mandarin)_HK_Flight_Attendant
- **英语**：
  - English_Graceful_Lady
  - English_Insightful_Speaker
  - English_radiant_girl
  - English_Persuasive_Man
  - moss_audio_6dc281eb-713c-11f0-a447-9613c873494c
  - moss_audio_570551b1-735c-11f0-b236-0adeeecad052
  - moss_audio_ad5baf92-735f-11f0-8263-fe5a2fe98ec8
  - English_Lucky_Robot
- **日语**：
  - Japanese_Whisper_Belle
  - Serene_Woman
  - moss_audio_24875c4a-7be4-11f0-9359-4e72c55db738
  - moss_audio_7f4ee608-78ea-11f0-bb73-1e2a4cfcd245
  - moss_audio_c1a6a3ac-7be6-11f0-8e8e-36b92fbb4f95
<hr/>

#### speed    <gray>float</gray>
合成音频的语速，取值越大，语速越快。取值范围 `[0.5,2]`，默认值为 1.0
<hr/>

#### vol    <gray>float</gray>
合成音频的音量，取值越大，音量越高。取值范围 `(0,10]`，默认值为 1.0
<hr/>

#### pitch    <gray>int</gray>
合成音频的语调，取值范围 `[-12,12]`，默认值为 0，其中 0 为原音色输出
<hr/>

#### emotion    <gray>string</gray>
- 控制合成语音的情绪；
- 参数范围 ["happy", "sad", "angry", "fearful", "disgusted", "surprised", "calm"]，分别对应 7 种情绪：高兴，悲伤，愤怒，害怕，厌恶，惊讶，中性
- 模型会根据输入文本自动匹配合适的情绪，一般无需手动指定。建议仅在确有需求时使用该参数，以保持更自然的合成效果
<hr/>

#### english_normalization    <gray>bool</gray>
该参数支持英语文本规范化，可提升数字阅读场景的性能，但会略微增加延迟，默认值为 false
<hr/>

#### latex_read    <gray>bool</gray>
控制是否朗读 latex 公式，默认为 false
**需注意**：
1. 请求中的公式需要在公式的首尾加上$$
2. 请求中公式若有`"\"`，需转义成`"\\"`
示例：一元二次方程根的基本公式 
<img src="https://filecdn.minimax.chat/public/63927147-9693-4c90-a001-ee39f8bf8e6d.png" alt="图片描述" width="30%" />
应表示为`$$x = \\frac{-b \pm \sqrt{b^2 - 4ac}}{2a} $$`。


</card>
</details>
<hr/>

#### audio_setting    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### sample_rate    <gray>int</gray>
生成音频的采样率。可选范围 `[8000，16000，22050，24000，32000，44100]`，默认为 32000
<hr/>

#### bitrate    <gray>int</gray>
生成音频的比特率。可选范围 `[32000，64000，128000，256000]`，默认值为 128000。该参数仅对 mp3 格式的音频生效
<hr/>

#### format    <gray>string</gray>
生成音频的格式。可选范围 `[mp3, pcm, flac, wav]`，默认 mp3，wav 仅在非流式输出下支持
<hr/>

#### channel    <gray>int</gray>
生成音频的声道数。可选范围：`[1,2]`，其中 1 为单声道，2 为双声道，默认值为 1

</card>
</details>
<hr/>

#### pronunciation_dict    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### tone    <gray>list</gray>
定义需要特殊标注的文字或符号对应的注音或发音替换规则。在中文文本中，声调用数字表示，一声为 1，二声为 2，三声为 3，四声为 4，轻声为 5
示例如下： `["燕少飞/(yan4)(shao3)(fei1)",  "omg/oh my god"]`

</card>
</details>
<hr/>

#### timber_weights    <gray>object</gray>
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### voice_id    <gray>string</gray>
合成音频的音色编号，须和 weight 参数同步填写。支持系统音色、复刻音色以及文生音色三种类型。系统支持的全部音色可查看[音色介绍]()，也可使用 [获取音色 API](https://platform.minimaxi.com/document/get_voice?key=67ad69b6de78c19cae68b199) 查询系统支持的全部音色
<hr/>

#### weight    <gray>int</gray>
合成音频各音色所占的权重，须与 voice_id 同步填写。可选值范围为 `[1,100]`，最多支持 4 种音色混合，单一音色取值占比越高，合成音色与该音色相似度越高

<codetabs title="参数示例">

<codetab language="json">

   "timber_weights": [
    {
      "voice_id": "Calm_Woman",
      "weight": 30
    },
    {
      "voice_id": "Wise_Woman",
      "weight": 70
    }]

</codetab>
</codetabs>

</card>
</details>
<hr/>

#### language_boost    <gray>string</gray>
是否增强对指定的小语种和方言的识别能力，设置后可以提升在指定小语种/方言场景下的语音表现。默认值为 null，如果不明确小语种类型，可设置为 "auto"，模型将自主判断小语种类型。可选值范围： [`Chinese`, `Chinese,Yue`, `English`, `Arabic`, `Russian`, `Spanish`, `French`, `Portuguese`, `German`, `Turkish`, `Dutch`, `Ukrainian`, `Vietnamese`, `Indonesian`, `Japanese`, `Italian`, `Korean`, `Thai`, `Polish`, `Romanian`, `Greek`, `Czech`, `Finnish`, `Hindi`, `auto`]
<hr/>

#### voice_modify    <gray>object</gray>
声音效果器设置
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### pitch    <gray>int</gray>
音高调整（低沉/明亮），范围 [-100,100]，数值接近 -100，声音更低沉；接近 100，声音更明亮
<img src="https://filecdn.minimax.chat/public/5d210c47-4236-4e81-893b-16cc1ef0302d.png" alt="图片描述" width="80%" />


<hr/>

#### intensity    <gray>int</gray>
强度调整（力量感/柔和），范围 [-100,100]，数值接近 -100，声音更刚劲；接近 100，声音更轻柔
<img src="https://filecdn.minimax.chat/public/862d493e-71d5-4d1f-b7c3-9ac51890631b.png" alt="图片描述" width="80%" />


<hr/>

#### timbre    <gray>int</gray>
音色调整（磁性/清脆），范围 [-100,100]，数值接近 -100，声音更浑厚；数值接近 100，声音更清脆
<img src="https://filecdn.minimax.chat/public/5f0e6cae-363a-452b-8d42-fbc4ef5a0510.png" alt="图片描述" width="80%" />


<hr/>

#### sound_effects    <gray>string</gray>
音效设置，共 4 种可选音效枚举值，单次支持使用 1 种音效
1. `spacious_echo`（空旷回音）
2. `auditorium_echo`（礼堂广播）
3. `lofi_telephone`（电话失真）
4. `robotic`（电音）

</card>
</details>


### 返回体
<hr/>

#### session_id    <gray>string</gray>
表示整个会话的 id
<hr/>

#### event    <gray>string</gray>
表示会话事件类型，当前环节成功后会返回 "task_started"
<hr/>

#### trace_id    <gray>string</gray>
表示会话中单次请求的 id，用于在咨询/反馈时帮助定位问题
<hr/>

#### base_resp    <gray>dict</gray>
该请求对应的状态码和详情
<details>
<summary expanded-text="展开参数" collapsed-text="隐藏参数"></summary>
<card>

#### status_code    <gray>int64</gray>
状态码。
- 0 表示发送成功
- 2202 表示非法事件

更多状态码信息请参考[错误码查询](/document/error_codes?key=68ad64216fe587e3fbfe7acf)
<hr/>

#### status_msg    <gray>string</gray>
状态详情

</card>
</details>

</left>

<right>

<codetabs title="请求示例">
<codetab language="json">
{
    "event":"task_start",
    "model":"speech-2.5-turbo-preview",
    "language_boost":"Chinese",
    "voice_setting":{
        "voice_id":"male-qn-qingse",
        "speed":1,
        "vol":1,
        "pitch":0
    },
    "pronunciation_dict":{
        "tone":["处理/(chu3)(li3)", "危险/dangerous"]
    },
    "audio_setting":{
        "sample_rate":32000,
        "bitrate":128000,
        "format":"mp3",
        "channel":1
    }
}

</codetab>
</codetabs>

<codetabs title="返回示例">
<codetab language="json">

{
    "session_id":"xxxx",
    "event":"task_started",
    "trace_id":"0303a2882bf18235ae7a809ae0f3cca7",
    "base_resp":{
        "status_code": 0,
        "status_msg":"success"
    }
}

</codetab>
</codetabs>

</right>


<br>

</container>

<container gap="4" columns="2">

  <card  href="/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成指南 websocket](/document/guides_t2a_websocket?key=68bbebb795a600dbfc9259ba)

  </card>

<card  href="/document/t2a_api_intro?key=68ad77666602726333000457">
<img src="https://filecdn.minimax.chat/public/32c799c8-6478-492a-813a-5799fdde19ba.png" alt="图片描述" width="20px" />

推荐阅读

 [同步语音合成接口能力介绍](/document/t2a_api_intro?key=68ad77666602726333000457)

  </card>

</container>

## ChatCompletion v2（历史版本）
- 原始链接: https://platform.minimaxi.com/document/对话?key=36954880539f4e2ab8b13067
- 文档ID: `36954880539f4e2ab8b13067`
### ChatCompletion v2
该API支持Openai参数定义和基础对话能力，还支持tools能力：function calling、web search，可参考示例进行接入使用。支持通过 API 或者 OpenAI SDK进行接入，参考下方文档进行接入。
本接口为MiniMax-Text-01模型的调用接口。
如果输入内容或输出内容严重违规，接口会返回内容违规错误信息，回复内容为空。
### API说明
`API：https://api.minimax.chat/v1/text/chatcompletion_v2`
#### 支持模型和token数
| 模型名称 | 最大token数（一次请求+回复的输入输出总token） |
| --- | --- |
| MiniMax-Text-01 | 1000192 |
#### 接口参数说明
###### 请求体（Request）参数
###### 返回(Response)参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| Content-Type | application/json | 是 | Content-Type |
| model | string | 是 | 调用的模型ID。目前支持取以下值：`MiniMax-Text-01`。 |
| stream | bool | 否 | 是否通过流式分批返回结果。如果设置为true，结果分批返回，两个换行分割分批返回结果.默认为false |
| max_tokens | int64 | 否 | 最大生成token数，需要注意的是，这个参数并不会影响模型本身的生成效果，而是仅仅通过以截断超出的token的方式来实现功能。可选，目前默认取值MiniMax-Text-01为2048。如遇到length原因停止生成请修改该值。<br>模型支持的取值区间如下：<br>MiniMax-Text-01： (0,1000192] |
| temperature | float，(0,1] | 否 | 较高的值将使输出更加随机，而较低的值将使输出更加集中和确定。<br>MiniMax-Text-01默认取值为0.1。低（0.01~0.2）：适合答案较明确的场景（如：知识问答、总结说明、情感分析、文本分类、大纲生成、作文批改）<br>⾼（0.7〜1）：适合答案较开放发散的场景 （如：营销文案生成、人设对话） |
| top_p | float，(0,1] | 否 | 采样方法，数值越小结果确定性越强；数值越大，结果越随机。取值可选，各模型默认为0.95 |
| mask_sensitive_info | bool | 否 | 对输出中易涉及隐私问题的文本信息进行打码，目前包括但不限于邮箱、域名、链接、证件号、家庭住址等，默认False，即不开启打码。 |
| messages | array | 是 | 对话的内容。<br>发送者的类型。取值如下：<br>system：定义智能体bot设定的role类型；<br>user：在对话阶段，用户发送消息的role类型；<br>assistant：在对话阶段，智能体发送回复的role类型；<br>tool：定义function的role类型；<br>发送者的名称。须提供具体名称以区分同一类型的不同具体角色。<br>发送的内容。当为纯文本输入时使用string类型，当需要image input时使用array类型。<br>内容类型，支持text、image_url、input_audio三种<br>当图文混合输入时，文本类型的具体文字内容信息。<br>音频帧数据，支持base64/hex编码音频，格式支持pcm/mp3/wav<br>输入音频的格式，目前仅支持pcm/mp3/wav，非必填时会自动检测音频数据格式<br>输入音频的采样率，wav/mp3不填写时自动检测，仅pcm必须填写。<br>输入音频的采样位深，wav/mp3不填写时自动检测，仅pcm必须填写<br>输入音频的声道数，wav/mp3不填写时自动检测，仅pcm必须填写。<br>输入音频的编码方式，仅支持hex、base64，不填写时会自动检测<br>图片的url信息，url要保证公网可访问，也支持base64编码的图片信息<br>1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息<br>模型生成的tools信息。role需要为assistant。<br>function的调用信息对应的id.在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| messages.role | string | 是 | 发送者的类型。取值如下：<br>system：定义智能体bot设定的role类型；<br>user：在对话阶段，用户发送消息的role类型；<br>assistant：在对话阶段，智能体发送回复的role类型；<br>tool：定义function的role类型； |
| messages.name | string | 否 | 发送者的名称。须提供具体名称以区分同一类型的不同具体角色。 |
| messages.content | string or array | 是 | 发送的内容。当为纯文本输入时使用string类型，当需要image input时使用array类型。<br>内容类型，支持text、image_url、input_audio三种<br>当图文混合输入时，文本类型的具体文字内容信息。<br>音频帧数据，支持base64/hex编码音频，格式支持pcm/mp3/wav<br>输入音频的格式，目前仅支持pcm/mp3/wav，非必填时会自动检测音频数据格式<br>输入音频的采样率，wav/mp3不填写时自动检测，仅pcm必须填写。<br>输入音频的采样位深，wav/mp3不填写时自动检测，仅pcm必须填写<br>输入音频的声道数，wav/mp3不填写时自动检测，仅pcm必须填写。<br>输入音频的编码方式，仅支持hex、base64，不填写时会自动检测<br>图片的url信息，url要保证公网可访问，也支持base64编码的图片信息<br>1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息 |
| messages.content.type | string | 否 | 内容类型，支持text、image_url、input_audio三种 |
| messages.content.text | string | 否 | 当图文混合输入时，文本类型的具体文字内容信息。 |
| messages.content.input_audio |  | 否 | 音频帧数据，支持base64/hex编码音频，格式支持pcm/mp3/wav<br>输入音频的格式，目前仅支持pcm/mp3/wav，非必填时会自动检测音频数据格式<br>输入音频的采样率，wav/mp3不填写时自动检测，仅pcm必须填写。<br>输入音频的采样位深，wav/mp3不填写时自动检测，仅pcm必须填写<br>输入音频的声道数，wav/mp3不填写时自动检测，仅pcm必须填写。<br>输入音频的编码方式，仅支持hex、base64，不填写时会自动检测 |
| messages.content.input_audio.data | string | 是 | 音频帧数据，支持base64/hex编码音频，格式支持pcm/mp3/wav |
| messages.content.input_audio.format | string | 否 | 输入音频的格式，目前仅支持pcm/mp3/wav，非必填时会自动检测音频数据格式 |
| messages.content.input_audio.sample_rate | int | 否 | 输入音频的采样率，wav/mp3不填写时自动检测，仅pcm必须填写。 |
| messages.content.input_audio.bit_depth | int | 否 | 输入音频的采样位深，wav/mp3不填写时自动检测，仅pcm必须填写 |
| messages.content.input_audio.channel | int | 否 | 输入音频的声道数，wav/mp3不填写时自动检测，仅pcm必须填写。 |
| messages.content.input_audio.encode | string | 否 | 输入音频的编码方式，仅支持hex、base64，不填写时会自动检测 |
| messages.content.image_url |  | 否 | 图片的url信息，url要保证公网可访问，也支持base64编码的图片信息<br>1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息 |
| messages.content.image_url.url | string | 否 | 1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息 |
| messages.tool_calls |  | 否 | 模型生成的tools信息。role需要为assistant。<br>function的调用信息对应的id.在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| messages.tool_calls.id | string | 是 | function的调用信息对应的id.在触发function时会返回该id。 |
| messages.tool_calls.type | string | 是 | function的类型，仅支持function。 |
| messages.tool_calls.function |  | 是 | 模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| messages.tool_calls.function.name | string | 是 | function对应的名称。 |
| messages.tool_calls.function.arguments | string | 是 | function的具体调用信息。 |
| tool_choice | string | 否 | tool工具的开关.支持以下2种模式：1、none：不调用function call；2、auto：自动判断是否调用funcition call。<br>tools.type = "web_search"时，只支持tool_choice设定为auto。该参数仅支持`MiniMax-Text-01` |
| tools |  | 否 | 支持的工具。该参数仅支持`MiniMax-Text-01`<br>支持的工具，目前支持function、web_search.<br>tools.type = "web_search"时，开启网络检索功能链路。<br>function的定义。<br>funciton的描述。<br>function的名称。<br>function的参数。 |
| tools.type | string | 是 | 支持的工具，目前支持function、web_search.<br>tools.type = "web_search"时，开启网络检索功能链路。 |
| tools.function |  | 是 | function的定义。<br>funciton的描述。<br>function的名称。<br>function的参数。 |
| tools.function.description | string | 是 | funciton的描述。 |
| tools.function.name | string | 是 | function的名称。 |
| tools.function.parameters | object | 否 | function的参数。 |
| response_format | object | 否 | 指定模型必须输出的格式。<br>设置为`{ "type": "json_schema", "json_schema": {...} }`以实现结构化输出，这会确保模型返回的内容结构完全匹配您提供的JSON schema。<br>示例见右方代码区域。<br>定义的响应格式类型， 仅支持`"json_schema"`<br>描述输出格式的说明，用于模型确定如何以该格式进行输出。<br>输出格式的名称。必须由a-z、A-Z、0-9 组成，或包含下划线和破折号，最大长度为64。<br>定义输出格式，形式为JSON Schema object。<br>取值应为`"object"`<br>详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| response_format.type | string | 是 | 定义的响应格式类型， 仅支持`"json_schema"` |
| response_format.json_schema | object | 是 | 描述输出格式的说明，用于模型确定如何以该格式进行输出。<br>输出格式的名称。必须由a-z、A-Z、0-9 组成，或包含下划线和破折号，最大长度为64。<br>定义输出格式，形式为JSON Schema object。<br>取值应为`"object"`<br>详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| response_format.json_schema.description | string | 否 | 描述输出格式的说明，用于模型确定如何以该格式进行输出。 |
| response_format.json_schema.name | string | 是 | 输出格式的名称。必须由a-z、A-Z、0-9 组成，或包含下划线和破折号，最大长度为64。 |
| response_format.json_schema.schema | object | 否 | 定义输出格式，形式为JSON Schema object。<br>取值应为`"object"`<br>详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| response_format.json_schema.schema.type | string | 是 | 取值应为`"object"` |
| response_format.json_schema.schema.properties | object | 是 | 详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| stream_options |  | 否 | 仅支持abab6.5s-chat，是否开启语音的流式输出，开启后，在输出文本的同时会同时输出流式的语音，可以用在语音通话场景<br>如果设置为True，将在最后流出一个额外的chunk。此块上的 usage 字段显示整个请求的token使用统计信息，而 choices 字段始终为空数组。所有其他块也将包含一个 usage 字段，但值为 null。 |
| stream_options.speech_output | bool | 否 | 仅支持abab6.5s-chat，是否开启语音的流式输出，开启后，在输出文本的同时会同时输出流式的语音，可以用在语音通话场景 |
| stream_options.include_usage | boolean | 否 | 如果设置为True，将在最后流出一个额外的chunk。此块上的 usage 字段显示整个请求的token使用统计信息，而 choices 字段始终为空数组。所有其他块也将包含一个 usage 字段，但值为 null。 |
| voice_setting |  | 否 | 只有开启stream_options中的speech_output后配置该参数才会生效<br>语音流式输出的模型，可以选择如下模型：speech-01-turbo、speech-01-240228、speech-01-turbo-240228<br>范围[0.5,2]，默认值为1.0，生成声音的语速，可选，取值越大，语速越快。<br>（0,10]，默认值为1.0，生成声音的音量，可选，取值越大，音量越高。<br>范围[-12,12]，默认值为0，生成声音的语调，可选，（0为原音色输出，取值需为整数）。<br>请求的音色编号。<br>支持系统音色(id)以及复刻音色（id）两种类型，其中系统音色（ID）如下：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin<br>聪明男童：clever_boy<br>可爱男童：cute_boy<br>萌萌女童：lovely_girl<br>卡通猪小琪：cartoon_pig<br>病娇弟弟：bingjiao_didi<br>俊朗男友：junlang_nanyou<br>纯真学弟：chunzhen_xuedi<br>冷淡学长：lengdan_xiongzhang<br>霸道少爷：badao_shaoye<br>甜心小玲：tianxin_xiaoling<br>俏皮萌妹：qiaopi_mengmei<br>妩媚御姐：wumei_yujie<br>嗲嗲学妹：diadia_xuemei<br>淡雅学姐：danya_xuejie<br>Santa Claus：Santa_Claus <br>Grinch：Grinch<br>Rudolph：Rudolph<br>Arnold：Arnold<br>Charming Santa：Charming_Santa<br>Charming Lady：Charming_Lady<br>Sweet Girl：Sweet_Girl<br>Cute Elf：Cute_Elf<br>Attractive Girl：Attractive_Girl<br>Serene Woman：Serene_Woman |
| voice_setting.model | string | 否 | 语音流式输出的模型，可以选择如下模型：speech-01-turbo、speech-01-240228、speech-01-turbo-240228 |
| voice_setting.speed | float | 否 | 范围[0.5,2]，默认值为1.0，生成声音的语速，可选，取值越大，语速越快。 |
| voice_setting.vol | float | 否 | （0,10]，默认值为1.0，生成声音的音量，可选，取值越大，音量越高。 |
| voice_setting.pitch | float | 否 | 范围[-12,12]，默认值为0，生成声音的语调，可选，（0为原音色输出，取值需为整数）。 |
| voice_setting.voice_id | string | 否 | 请求的音色编号。<br>支持系统音色(id)以及复刻音色（id）两种类型，其中系统音色（ID）如下：<br>青涩青年音色：male-qn-qingse<br>精英青年音色：male-qn-jingying<br>霸道青年音色：male-qn-badao<br>青年大学生音色：male-qn-daxuesheng<br>少女音色：female-shaonv<br>御姐音色：female-yujie<br>成熟女性音色：female-chengshu<br>甜美女性音色：female-tianmei<br>男性主持人：presenter_male<br>女性主持人：presenter_female<br>男性有声书1：audiobook_male_1<br>男性有声书2：audiobook_male_2<br>女性有声书1：audiobook_female_1<br>女性有声书2：audiobook_female_2<br>青涩青年音色-beta：male-qn-qingse-jingpin<br>精英青年音色-beta：male-qn-jingying-jingpin<br>霸道青年音色-beta：male-qn-badao-jingpin<br>青年大学生音色-beta：male-qn-daxuesheng-jingpin<br>少女音色-beta：female-shaonv-jingpin<br>御姐音色-beta：female-yujie-jingpin<br>成熟女性音色-beta：female-chengshu-jingpin<br>甜美女性音色-beta：female-tianmei-jingpin<br>聪明男童：clever_boy<br>可爱男童：cute_boy<br>萌萌女童：lovely_girl<br>卡通猪小琪：cartoon_pig<br>病娇弟弟：bingjiao_didi<br>俊朗男友：junlang_nanyou<br>纯真学弟：chunzhen_xuedi<br>冷淡学长：lengdan_xiongzhang<br>霸道少爷：badao_shaoye<br>甜心小玲：tianxin_xiaoling<br>俏皮萌妹：qiaopi_mengmei<br>妩媚御姐：wumei_yujie<br>嗲嗲学妹：diadia_xuemei<br>淡雅学姐：danya_xuejie<br>Santa Claus：Santa_Claus <br>Grinch：Grinch<br>Rudolph：Rudolph<br>Arnold：Arnold<br>Charming Santa：Charming_Santa<br>Charming Lady：Charming_Lady<br>Sweet Girl：Sweet_Girl<br>Cute Elf：Cute_Elf<br>Attractive Girl：Attractive_Girl<br>Serene Woman：Serene_Woman |
| audio_setting |  | 否 | 只有开启stream_options中的speech_output后配置该参数才会生效<br>范围【8000，16000，22050，24000，32000，44100】，生成声音的采样率。可选，默认为32000。<br>范围【32000，64000，128000，256000】，生成声音的比特率。可选，默认值为128000。该参数仅对mp3格式的音频生效。<br>生成的音频格式。默认mp3，范围[mp3,pcm,flac,wav]。wav仅在非流式输出下支持。<br>生成音频的声道数.默认1：单声道，可选：<br>1：单声道<br>2：双声道<br>生成音频的编码方式，仅支持base64/hex，默认hex |
| audio_setting.sample_rate | int | 否 | 范围【8000，16000，22050，24000，32000，44100】，生成声音的采样率。可选，默认为32000。 |
| audio_setting.bitrate | int | 否 | 范围【32000，64000，128000，256000】，生成声音的比特率。可选，默认值为128000。该参数仅对mp3格式的音频生效。 |
| audio_setting.format | string | 否 | 生成的音频格式。默认mp3，范围[mp3,pcm,flac,wav]。wav仅在非流式输出下支持。 |
| audio_setting.channel | int | 否 | 生成音频的声道数.默认1：单声道，可选：<br>1：单声道<br>2：双声道 |
| audio_setting.encode | string | 否 | 生成音频的编码方式，仅支持base64/hex，默认hex |
| id | string | 否 | 响应id |
| choices | array | 否 | 回复消息选项，里面包含对应的信息<br>模型生成停止的原因。目前支持取以下值：<br> 1、stop：正常结束；<br>2、tool_calls：触发function calling；<br>3、length：请求超过token长度。<br>回复消息的索引，目前仅支持1个选项。<br>回复信息详细内容。<br>回复的内容。<br>回复消息的角色，目前仅支持assistant。<br>模型生成的tools信息，role需要为assistant。<br>function的调用信息对应的id,在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。<br>function对应的名称，固定为"plugin_web_search"。<br>function的具体调用信息，包含本次搜索行为的相关信息。<br>function的类型，固定为"web_search"。<br>function的调用信息对应的id，在触发function时会返回该id<br>调用function的id，在触发function时会返回该id。<br>语音流式输出音频数据，默认以hex编码形式输出，可以在audio_setting里修改为base64编码输出。注意：在流式返回时，audio_content会和文本content一起返回输出，但本身不是一一对应关系。<br>思维链内容。（目前仅DeepSeek-R1支持） |
| choices.finish_reason | string | 否 | 模型生成停止的原因。目前支持取以下值：<br> 1、stop：正常结束；<br>2、tool_calls：触发function calling；<br>3、length：请求超过token长度。 |
| choices.index | int | 否 | 回复消息的索引，目前仅支持1个选项。 |
| choices.message | dict | 否 | 回复信息详细内容。<br>回复的内容。<br>回复消息的角色，目前仅支持assistant。<br>模型生成的tools信息，role需要为assistant。<br>function的调用信息对应的id,在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。<br>function对应的名称，固定为"plugin_web_search"。<br>function的具体调用信息，包含本次搜索行为的相关信息。<br>function的类型，固定为"web_search"。<br>function的调用信息对应的id，在触发function时会返回该id<br>调用function的id，在触发function时会返回该id。<br>语音流式输出音频数据，默认以hex编码形式输出，可以在audio_setting里修改为base64编码输出。注意：在流式返回时，audio_content会和文本content一起返回输出，但本身不是一一对应关系。<br>思维链内容。（目前仅DeepSeek-R1支持） |
| choices.message.content | string | 否 | 回复的内容。 |
| choices.message.role | string | 否 | 回复消息的角色，目前仅支持assistant。 |
| choices.message.tool_calls |  | 否 | 模型生成的tools信息，role需要为assistant。<br>function的调用信息对应的id,在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。<br>function对应的名称，固定为"plugin_web_search"。<br>function的具体调用信息，包含本次搜索行为的相关信息。<br>function的类型，固定为"web_search"。<br>function的调用信息对应的id，在触发function时会返回该id<br>调用function的id，在触发function时会返回该id。 |
| choices.message.tool_calls.id | string | 否 | function的调用信息对应的id,在触发function时会返回该id。 |
| choices.message.tool_calls.type | string | 否 | function的类型，仅支持function。 |
| choices.message.tool_calls.function |  | 否 | 模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。<br>function对应的名称，固定为"plugin_web_search"。<br>function的具体调用信息，包含本次搜索行为的相关信息。 |
| choices.message.tool_calls.function.name | string | 否 | function对应的名称。 |
| choices.message.tool_calls.function.arguments | string | 否 | function的具体调用信息。 |
| choices.message.tool_calls.function.choices.messages[1].tool_calls.function.name | string | 否 | function对应的名称，固定为"plugin_web_search"。 |
| choices.message.tool_calls.function.choices.messages[1].tool_calls.function.arguments | dict | 否 | function的具体调用信息，包含本次搜索行为的相关信息。 |
| choices.message.tool_calls.choices.messages[1].tool_calls.type | string | 否 | function的类型，固定为"web_search"。 |
| choices.message.tool_calls.choices.messages[1].tool_calls.id | string | 否 | function的调用信息对应的id，在触发function时会返回该id |
| choices.message.tool_calls.choices.messages[2].tool_call_id | string | 否 | 调用function的id，在触发function时会返回该id。 |
| choices.message.audio_content | string | 否 | 语音流式输出音频数据，默认以hex编码形式输出，可以在audio_setting里修改为base64编码输出。注意：在流式返回时，audio_content会和文本content一起返回输出，但本身不是一一对应关系。 |
| choices.message.reasoning_content | string | 否 | 思维链内容。（目前仅DeepSeek-R1支持） |
| created | int | 否 | 创建的时间戳。 |
| model | string | 否 | 聊天选择的模型。 |
| object | string | 否 | 对象类型为流式或非流式，取值为chat.completion或chat.completion.chunk。 |
| usage |  | 否 | 本次请求的消耗。<br>本次请求的消耗token。<br>生成语音消耗的字符数<br>输入语音的识别时长 |
| usage.total_tokens | int | 否 | 本次请求的消耗token。 |
| usage.total_characters | int | 否 | 生成语音消耗的字符数 |
| usage.total_asr_time | int | 否 | 输入语音的识别时长 |
| input_sensitive | bool | 否 | 输入命中敏感词。<br>输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| input_sensitive.input_sensitive_type | int64 | 否 | 输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| output_sensitive | bool | 否 | 输出命中敏感词。<br>输出命中敏感词类型，当output_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| output_sensitive.output_sensitive_type | int64 | 否 | 输出命中敏感词类型，当output_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| base_resp | BaseResp | 否 | 错误状态码和详情。<br>状态码。<br>1000，未知错误；<br>1001，请求超时；<br>1002，触发RPM限流；<br>1004，鉴权失败；<br>1008，余额不足；<br>1013，服务内部错误；<br>1027，输出内容错误；<br>1039，Token限制；<br>2013，参数错误。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。<br>1000，未知错误；<br>1001，请求超时；<br>1002，触发RPM限流；<br>1004，鉴权失败；<br>1008，余额不足；<br>1013，服务内部错误；<br>1027，输出内容错误；<br>1039，Token限制；<br>2013，参数错误。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
### SDK 接入
Chat Completion v2接口支持通过OpenAI SDK进行接入，分为如下2步OpenAI SDK：
1. 安装 Python 3.7.1 或更高版本后，就可以安装 OpenAI Python 库。在终端/命令行中，运行：
```python
pip install --upgrade openai
```
2. 在[账户管理-接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)获取你的API Key。
3. 在接入代码使用OpenAI SDK参考代码示例进行接入和使用，更多示例可参考上方API说明中的python示例。
**请求示例**
```python
from openai import OpenAI

client = OpenAI(api_key="<minimax api key>", base_url="https://api.minimax.chat/v1")

response = client.chat.completions.create(
  model="MiniMax-Text-01",
  messages=[
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Hello!"}
  ],
  stream=True
)

for chunk in response:
  print(chunk)
```
**返回示例**
```python
ChatCompletionChunk(id='02fd53267b5e09b5030cf1bc99f42d3a', choices=[Choice(delta=ChoiceDelta(content='你好', function_call=None, role='assistant', tool_calls=None), finish_reason=None, index=0, logprobs=None)], created=1722687527, model='MiniMax-Text-01', object='chat.completion.chunk', service_tier=None, system_fingerprint=None, usage=None, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0)
ChatCompletionChunk(id='02fd53267b5e09b5030cf1bc99f42d3a', choices=[Choice(delta=ChoiceDelta(content='！很高兴能和你交流。有什么问题或者话题想要讨论吗？我在这里帮助你。', function_call=None, role='assistant', tool_calls=None), finish_reason='stop', index=0, logprobs=None)], created=1722687527, model='MiniMax-Text-01', object='chat.completion.chunk', service_tier=None, system_fingerprint=None, usage=None, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0)
ChatCompletionChunk(id='02fd53267b5e09b5030cf1bc99f42d3a', choices=[Choice(delta=None, finish_reason='stop', index=0, logprobs=None, message={'content': '你好！很高兴能和你交流。有什么问题或者话题想要讨论吗？我在这里帮助你。', 'role': 'assistant'})], created=1722687527, model='MiniMax-Text-01', object='chat.completion', service_tier=None, system_fingerprint=None, usage=CompletionUsage(completion_tokens=None, prompt_tokens=None, total_tokens=87), input_sensitive=False, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0, base_resp={'status_code': 0, 'status_msg': ''})
```
### 函数调用（function calling）
该功能可触发调用用户自有函数（如客户自有的会议记录查询函数），得到自有函数的参考内容 ，辅助大模型生成结果。可以帮助模型获取用户希望模型获取的相应信息，以使得模型返回更符合用户个人实际情况的回复。特别是在一些需要基于客户自有信息知识的问题回答上，并且需要使用构建了该信息知识的API，那么用户可以考虑使用function calling功能，使用场景示例如下：
> 1、query：“在会议ID：12345的会议上，A同学有没有提到大模型如何应用的事情？”
> 2、用户自有函数，函数功能：可基于会议ID和参会人检索原始的会议记录
> 3、参考内容：“会议号：12345，参会人：A同学，发言内容：我们要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。”
> 4、大模型生成结果：“A同学在会议ID：12345的会议上提到了大模型如何应用的事情。他说要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。”
以上，如果仅靠大模型来回答query时无法得到符合预期的回复，因此在这时就需要使用function calling功能。
#### Function Calling示例
具体流程如下：
1. 触发function，输出function输入参数；
2. 得到function的出参，给到模型参考进行总结回复。
**Tab: 1、触发function，输出function输入参数**

**请求示例（流式）**
```curl
curl --location 'https://api.minimax.chat/v1/text/chatcompletion_v2' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
  "model": "MiniMax-Text-01",
  "messages": [
    {
      "role": "system",
      "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
      "role": "user",
      "content": "广州天气怎么样"
    }
  ],
  "tools": [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
],
  "tool_choice": "auto"
}'
```

**python**
```python
from openai import OpenAI

client = OpenAI(
    api_key="",  # <--在这里使用MiniMax账户管理-接口密钥中API KEY进行接入
    base_url="https://api.minimax.chat/v1",
)

tools = [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
]

messages = [
    {"role": "system",
     "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"},
    {"role": "user", "content": "广州天气怎么样？"}
]

completion = client.chat.completions.create(
  model="MiniMax-Text-01",
  messages=messages,
  tools=tools,
  tool_choice = 'auto'
)
print(completion)
```

**Tab: 2、得到function的出参，给到模型参考进行总结回复**

**请求示例（流式）**
```curl
curl --location 'https://api.minimax.chat/v1/text/chatcompletion_v2' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
  "model": "MiniMax-Text-01",
  "messages": [
    {
      "role": "system",
      "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
      "role": "user",
      "content": "广州天气怎么样"
    },
        {
        "role": "assistant",
        "tool_calls": [{
            "id": "call_function_1849136893",
            "type": "function",
            "function": {
                "name": "get_current_weather",
                "arguments": "{\"location\":\"广州\"}"
            }
        }]
    },
    {
        "role": "tool",
        "tool_call_id": "call_function_1849136893",
        "content": "多云，28~37℃，无持续风向<3级，空气质量优"
    }
  ],
  "tools": [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
],
  "tool_choice": "auto"
}'
```

**python**
```python
from openai import OpenAI

client = OpenAI(
    api_key="",  # <--在这里使用MiniMax账户管理-接口密钥中API KEY进行接入
    base_url="https://api.minimax.chat/v1",
)

tools = [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
]

messages = [
    {
        "role": "system",
        "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
        "role": "user",
        "content": "广州天气怎么样？"
     },
    {"role": "user", "content": "广州天气怎么样？"},
    {
        "role": "assistant",
        "tool_calls": [{
            "id": "call_function_1849136893",
            "type": "function",
            "function": {
                "name": "get_current_weather",
                "arguments": "{\"location\":\"广州\"}"
            }
        }]
    },
    {
        "role": "tool",
        "tool_call_id": "call_function_1849136893",
        "content": "多云，28~37℃，无持续风向<3级，空气质量优"
    }
]

completion = client.chat.completions.create(
  model="MiniMax-Text-01",
  messages=messages,
  tools=tools,
  tool_choice = 'auto'
)
print(completion)
```
**Tab: 1、触发function，输出function输入参数**

**返回示例（流式）**
```curl
{
    "id": "02ffb000ff4b7202faae09b999b27c16",
    "choices": [
        {
            "finish_reason": "tool_calls",
            "index": 0,
            "message": {
                "role": "assistant",
                "tool_calls": [
                    {
                        "id": "call_function_6835316862",
                        "type": "function",
                        "function": {
                            "name": "get_current_weather",
                            "arguments": "{\"location\": \"广州\"}"
                        }
                    }
                ]
            }
        }
    ],
    "created": 1722842368,
    "model": "MiniMax-Text-01",
    "object": "chat.completion",
    "usage": {
        "total_tokens": 153
    },
    "input_sensitive": false,
    "output_sensitive": false,
    "input_sensitive_type": 0,
    "output_sensitive_type": 0,
    "base_resp": {
        "status_code": 0,
        "status_msg": ""
    }
}
```

**python**
```python
ChatCompletion(id='02ffacd947f8a288ad6f0c1918d297a8', choices=[Choice(finish_reason='tool_calls', index=0, logprobs=None, message=ChatCompletionMessage(content=None, role='assistant', function_call=None, tool_calls=[ChatCompletionMessageToolCall(id='call_function_0304055004', function=Function(arguments='{"location": "广州"}', name='get_current_weather'), type='function')]))], created=1722841561, model='MiniMax-Text-01', object='chat.completion', service_tier=None, system_fingerprint=None, usage=CompletionUsage(completion_tokens=None, prompt_tokens=None, total_tokens=154), input_sensitive=False, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0, base_resp={'status_code': 0, 'status_msg': ''})
```

**Tab: 2、得到function的出参，给到模型参考进行总结回复**

**返回示例**
```curl
{
    "id": "02ffb37f1e7189808524269fc6fb9f9f",
    "choices": [
        {
            "finish_reason": "stop",
            "index": 0,
            "message": {
                "content": "广州现在的天气是多云，气温在28~37℃之间，没有持续的风向，风力小于3级，空气质量为优。",
                "role": "assistant"
            }
        }
    ],
    "created": 1722843264,
    "model": "MiniMax-Text-01",
    "object": "chat.completion",
    "usage": {
        "total_tokens": 223
    },
    "input_sensitive": false,
    "output_sensitive": false,
    "input_sensitive_type": 0,
    "output_sensitive_type": 0,
    "base_resp": {
        "status_code": 0,
        "status_msg": ""
    }
}
```

**python**
```python
ChatCompletion(id='02ffb2ec3574bb88b257536afb5f9876', choices=[Choice(finish_reason='stop', index=0, logprobs=None, message=ChatCompletionMessage(content='广州目前的天气是多云，气温在28~37℃之间，无持续风向小于3级，空气质量优。', role='assistant', function_call=None, tool_calls=None))], created=1722843117, model='MiniMax-Text-01', object='chat.completion', service_tier=None, system_fingerprint=None, usage=CompletionUsage(completion_tokens=None, prompt_tokens=None, total_tokens=232), input_sensitive=False, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0, base_resp={'status_code': 0, 'status_msg': ''})
```
### 网络搜索（web search）
该功能目前支持调用我们集成的搜索引擎生成内容 ，提供网络引擎检索的内容，辅助大模型生成结果。
#### web search返回参数
当您调用web search工具时，将会以以下结构返回生成内容：
| 参数名称 | 取值 | 参数描述 |
| --- | --- | --- |
| choices.messages[0].role | string | 发送信息的角色，仅支持user。 |
| choices.messages[0].content | string | 发送的内容，user发送的内容。 |
| choices.messages[1].role | string | 回复消息的角色，仅支持"assistant"。 |
| choices.messages[1].tool_calls.type | string | function的类型，固定为"web_search"。 |
| choices.messages[1].tool_calls.id | string | function的调用信息对应的id，在触发function时会返回该id。 |
| choices.messages[1].tool_calls.function.name | string | function对应的名称，固定为"plugin_web_search"。 |
| choices.messages[1].tool_calls.function.arguments | dict | function的具体调用信息，包含本次搜索行为的相关信息。 |
| choices.messages[2].role | string | 回复消息的角色，固定为"tool"。 |
| choices.messages[2].content | string | function的输出内容，网络检索得到的参考资料。 |
| choices.messages[2].tool_call_id | string | 调用function的id，在触发function时会返回该id。 |
| choices.messages[3].role | string | 回复消息的角色，仅支持"assistant"。 |
| choices.messages[3].content | string | 最终回复内容，参考网络检索结果，模型做出的最终回复。 |
**调用web search请求示例**
```curl
curl --location "https://api.minimax.chat/v1/text/chatcompletion_v2" \
--header "Content-Type: application/json" \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
    "model":"MiniMax-Text-01",
    "messages":[
      {
        "role":"system",
        "name":"MM智能助理", 
        "content":"MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
      },
      {
        "role":"user",
        "name":"用户", 
        "content":"请问巴黎奥运会郑钦文夺金了吗？"
      }
    ],
    "max_tokens":4096,
    "tools":[
        {
         "type":"web_search"
        }
    ]
  }'
```

**python**
```python
from openai import OpenAI

client = OpenAI(
    api_key="",  # <--在这里使用MiniMax账户管理-接口密钥中API KEY进行接入
    base_url="https://api.minimax.chat/v1",
)

tools = [
  {
    "type": "web_search",
  }
]

messages = [
    {"role": "system",
     "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"},
    {"role": "user", "content": "请问巴黎奥运会郑钦文夺金了吗？"}
]

completion = client.chat.completions.create(
  model="MiniMax-Text-01",
  messages=messages,
  max_tokens=4096,
  tools=tools,
  tool_choice="auto"
)

print(completion.choices[0].messages[3])
```
**调用web search返回示例**
```curl
{
    "id": "02ffaafd5907ea7661b668bb48d5e09b",
    "choices": [
        {
            "finish_reason": "tool_calls",
            "index": 0,
            "messages": [
                {
                    "content": "请问巴黎奥运会郑钦文夺金了吗？",
                    "role": "user"
                },
                {
                    "role": "assistant",
                    "tool_calls": [
                        {
                            "id": "call_function_4175054147",
                            "type": "web_search",
                            "function": {
                                "name": "plugin_web_search",
                                "arguments": "{\"query_tag\": [\"体育\"], \"search_period\": \"2024/07/29-\", \"query_list\": [\"2024年巴黎奥运会 郑钦文\"]}"
                            }
                        }
                    ]
                },
                {
                    "content": "参考资料:\n[0] 郑钦文,2002年10月8日出生于湖北省十堰市,中国女子网球运动员,国际级运动健将. 2017年,进入湖北网球队.2019年,签约至IMG公司,正式进入职业化训练、比赛.2020年8月29日,在国际网联巡回赛科尔德农斯站夺得成人赛首个冠军.2021年9月8日,和尤晓迪搭档获第十四届全国运动会网球项目女双季军.2022年9月25日,获WTA500东京站女单亚军;12月12日,获WTA2022年度最佳新人奖.2023年7月,在WTA250巴勒莫站女单决赛中,斩获职业生涯巡回赛首冠;9月,在美网女单1/8决赛中获胜,生涯首次打进大满贯女单8强;9月29日,在杭州第19届亚运会网球女子单打比赛中,获得首个洲际最高综合型运动会冠军;10月15日,在WTA500郑州网球公开赛女单决赛中,夺得职业生涯首个WTA500赛冠军;10月29日,获珠海WTA超级精英赛女单亚军.2024年1月27日,获2024年澳网女单亚军,夺得职业生涯首座大满贯单打亚军;7月22日,获得WTA250巴勒莫站女单冠军. 2024年8月3日,以总分2-0战胜多娜·维基奇获2024年巴黎奥运会网球女子单打冠军,刷新了中国网球在奥运会的历史最好成绩. 中文名 郑钦文 身高 178cm 国籍 中国 运动项目 网球 出生日期 2002年10月8日 出生地 湖北省十堰市 主要成就 2022年网球W60奥兰多站女子单打冠军、2022年W125瓦伦西亚女单冠军、2022年WTA500东京站女单亚军、2022年度WTA最佳新人奖、2023年WTA250巴勒莫站女单冠军、2023年杭州第19届亚运会网球女子单打冠军、2023年郑州网球公开赛女单冠军、2023年珠海WTA超级精英赛女单亚军、2023赛季WTA年度进步最快球员奖、2024年澳网女单亚军、2024年WTA250巴勒莫站决赛冠军、2024年巴黎奥运会网球女子单打冠军 经纪公司 IMG 世界排名 WTA世界排名第7(2024年7月29日) 教练 佩雷·里巴 人物经历 早年经历 郑钦文6岁开始学习打网球.7岁时,被送到了李娜的启蒙教练夏溪瑶处进行训练.2010年,来到武汉崇仁网球培训中心,由武汉市网球队主教练谢纯带训执教,进行网球基本技术提高训练,并在崇仁路小学就读. 2011年,被送到李娜前指导老师余丽桥处进行训练.2012年,随队至江汉二桥训练基地进行专项提高训练,\n[1] 新华社巴黎8月3日电(记者高萌、乐文婉、张百慧)赢了!右腿缠着绷带的郑钦文第三次躺倒在罗兰·加洛斯中央球场的红土上,这一次,她没有流泪. 8月3日,郑钦文庆祝夺冠.新华社记者 高静 摄 几秒钟前,欢呼声瞬间充盈整个球场,并随着空气飘向高处;而这座球场的观众席,第一次如此大面积地被五星红旗覆盖. 郑钦文,这个21岁的中国女孩创造了历史:她成为中国乃至亚洲第一位走上奥运会网球女单最高领奖台的运动员. 就在三天前,在同样的地点.郑钦文耗时三小时零四分击败了德国老将科贝尔,晋级四强,追平了中国女单的奥运会最好成绩.在决胜球落地的一瞬间,她扔掉拍子,躺倒在场上,任凭红色的土蹭在自己的衣服与身体上.起身时,她悄悄用手上的护腕擦去了自己的眼泪.赛后她解释道,自己从未打过如此艰难的比赛. 8月3日,郑钦文在比赛中发球.新华社记者 高静 摄 两天前,依旧是这一球场.郑钦文战胜了自己职业生涯中从未战胜过的对手、世界第一的斯维亚特克.她再度躺倒,咆哮着庆祝自己的胜利. 郑钦文的第一次奥运之旅并不容易,但她没有给自己退路;也正是因为没退路,她硬是一场又一场地熬了下来. 三次躺倒,第一次为突破了极限;第二次为战胜了心魔;第三次,她终于可以轻松地庆祝自己的胜利. 时间拉回决赛现场.再度起身,郑钦文从观众席接过了一面国旗.小心翼翼地把它平整地铺在一旁的椅子上后,她带着笑容转身,向观众席发球庆祝胜利. 8月3日,郑钦文庆祝夺冠.新华社记者 万象 摄 李娜曾说:“网球是一项孤独的运动,当你独自上场,你就开始了一个人的战斗.” 可今晚的郑钦文并不孤独. 在这场决赛中,中文的“加油”声此起彼伏.网球开赛六天以来,一直坐在记者身边的路透社与美联社记者终于扭过头来询问了“加油”的拼写与语义. 夺冠那一刻,从深圳飞来巴黎的张潇予在观众席上落了泪.早在去年就预订了奥运网球门票的她,从没想过自己会见证这一刻.“我想我会一直回味这一刻,我相信它会带给我一些动力和勇气去突破自己的困境.” 常居法国的李斯扬则感叹:“她真的在撰写自己的人生故事!我非常为她感到骄傲,她就是我的人生目标.而且她真的是在为祖国而战,因此我们在现场也感受到了很强的凝聚力.这场比赛意义很特殊,很高兴能在现场见证.” 8月3日,球迷在看台为郑钦文助威.新华社记者 万象 摄 成为别人人生目标的郑钦文,也在发布会上透露,自己一直很羡慕那些可以创造历史的\n[2] 人民网北京8月4日电 (记者李乃妍)北京时间8月4日凌晨,在巴黎奥运会网球女单决赛中,中国队选手郑钦文战胜克罗地亚队选手维基奇,摘得金牌.这是2004年雅典奥运会李婷/孙甜甜女双夺金后,中国网球队收获的又一枚奥运金牌,也是中国网球队夺得的首枚奥运会女单金牌. \n[3] \n",
                    "role": "tool",
                    "tool_call_id": "call_function_4175054147"
                },
                {
                    "content": "是的，郑钦文在2024年巴黎奥运会上夺得了网球女子单打金牌。根据最新的报道，郑钦文在8月3日的比赛中以总分2-0战胜了多娜·维基奇，从而获得了2024年巴黎奥运会网球女子单打冠军。这一成就刷新了中国网球在奥运会的历史最好成绩【0†source】【1†source】。",
                    "role": "assistant",
                    "annotations": [
                        {
                            "text": "【0†source】",
                            "url": "http://baike.sogou.com/v179355271.htm?fromTitle=%E9%83%91%E9%92%A6%E6%96%87",
                            "quote": "郑钦文,2002年10月8日出生于湖北省十堰市,中国女子网球运动员,国际级运动健将. 2017年,进入湖北网球队.2019年,签约至IMG公司,正式进入职业化训练、比赛.2020年8月29日,在国际网联巡回赛科尔德农斯站夺得成人赛首个冠军.2021年9月8日,和尤晓迪搭档获第十四届全国运动会网球项目女双季军.2022年9月25日,获WTA500东京站女单亚军;12月12日,获WTA2022年度最佳新人奖.2023年7月,在WTA250巴勒莫站女单决赛中,斩获职业生涯巡回赛首冠;9月,在美网女单1/8决赛中获胜,生涯首次打进大满贯女单8强;9月29日,在杭州第19届亚运会网球女子单打比赛中,获得首个洲际最高综合型运动会冠军;10月15日,在WTA500郑州网球公开赛女单决赛中,夺得职业生涯首个WTA500赛冠军;10月29日,获珠海WTA超级精英赛女单亚军.2024年1月27日,获2024年澳网女单亚军,夺得职业生涯首座大满贯单打亚军;7月22日,获得WTA250巴勒莫站女单冠军. 2024年8月3日,以总分2-0战胜多娜·维基奇获2024年巴黎奥运会网球女子单打冠军,刷新了中国网球在奥运会的历史最好成绩. 中文名 郑钦文 身高 178cm 国籍 中国 运动项目 网球 出生日期 2002年10月8日 出生地 湖北省十堰市 主要成就 2022年网球W60奥兰多站女子单打冠军、2022年W125瓦伦西亚女单冠军、2022年WTA500东京站女单亚军、2022年度WTA最佳新人奖、2023年WTA250巴勒莫站女单冠军、2023年杭州第19届亚运会网球女子单打冠军、2023年郑州网球公开赛女单冠军、2023年珠海WTA超级精英赛女单亚军、2023赛季WTA年度进步最快球员奖、2024年澳网女单亚军、2024年WTA250巴勒莫站决赛冠军、2024年巴黎奥运会网球女子单打冠军 经纪公司 IMG 世界排名 WTA世界排名第7(2024年7月29日) 教练 佩雷·里巴 人物经历 早年经历 郑钦文6岁开始学习打网球.7岁时,被送到了李娜的启蒙教练夏溪瑶处进行训练.2010年,来到武汉崇仁网球培训中心,由武汉市网球队主教练谢纯带训执教,进行网球基本技术提高训练,并在崇仁路小学就读. 2011年,被送到李娜前指导老师余丽桥处进行训练.2012年,随队至江汉二桥训练基地进行专项提高训练,"
                        },
                        {
                            "text": "【1†source】",
                            "url": "http://www.hb.xinhuanet.com/20240804/d5e717e4b25647cc878182a5baeb6819/c.html",
                            "quote": "新华社巴黎8月3日电(记者高萌、乐文婉、张百慧)赢了!右腿缠着绷带的郑钦文第三次躺倒在罗兰·加洛斯中央球场的红土上,这一次,她没有流泪. 8月3日,郑钦文庆祝夺冠.新华社记者 高静 摄 几秒钟前,欢呼声瞬间充盈整个球场,并随着空气飘向高处;而这座球场的观众席,第一次如此大面积地被五星红旗覆盖. 郑钦文,这个21岁的中国女孩创造了历史:她成为中国乃至亚洲第一位走上奥运会网球女单最高领奖台的运动员. 就在三天前,在同样的地点.郑钦文耗时三小时零四分击败了德国老将科贝尔,晋级四强,追平了中国女单的奥运会最好成绩.在决胜球落地的一瞬间,她扔掉拍子,躺倒在场上,任凭红色的土蹭在自己的衣服与身体上.起身时,她悄悄用手上的护腕擦去了自己的眼泪.赛后她解释道,自己从未打过如此艰难的比赛. 8月3日,郑钦文在比赛中发球.新华社记者 高静 摄 两天前,依旧是这一球场.郑钦文战胜了自己职业生涯中从未战胜过的对手、世界第一的斯维亚特克.她再度躺倒,咆哮着庆祝自己的胜利. 郑钦文的第一次奥运之旅并不容易,但她没有给自己退路;也正是因为没退路,她硬是一场又一场地熬了下来. 三次躺倒,第一次为突破了极限;第二次为战胜了心魔;第三次,她终于可以轻松地庆祝自己的胜利. 时间拉回决赛现场.再度起身,郑钦文从观众席接过了一面国旗.小心翼翼地把它平整地铺在一旁的椅子上后,她带着笑容转身,向观众席发球庆祝胜利. 8月3日,郑钦文庆祝夺冠.新华社记者 万象 摄 李娜曾说:“网球是一项孤独的运动,当你独自上场,你就开始了一个人的战斗.” 可今晚的郑钦文并不孤独. 在这场决赛中,中文的“加油”声此起彼伏.网球开赛六天以来,一直坐在记者身边的路透社与美联社记者终于扭过头来询问了“加油”的拼写与语义. 夺冠那一刻,从深圳飞来巴黎的张潇予在观众席上落了泪.早在去年就预订了奥运网球门票的她,从没想过自己会见证这一刻.“我想我会一直回味这一刻,我相信它会带给我一些动力和勇气去突破自己的困境.” 常居法国的李斯扬则感叹:“她真的在撰写自己的人生故事!我非常为她感到骄傲,她就是我的人生目标.而且她真的是在为祖国而战,因此我们在现场也感受到了很强的凝聚力.这场比赛意义很特殊,很高兴能在现场见证.” 8月3日,球迷在看台为郑钦文助威.新华社记者 万象 摄 成为别人人生目标的郑钦文,也在发布会上透露,自己一直很羡慕那些可以创造历史的"
                        }
                    ]
                }
            ]
        }
    ],
    "created": 1722841089,
    "model": "MiniMax-Text-01",
    "object": "chat.completion",
    "usage": {
        "total_tokens": 3063
    },
    "input_sensitive": false,
    "output_sensitive": false,
    "input_sensitive_type": 0,
    "output_sensitive_type": 0,
    "base_resp": {
        "status_code": 0,
        "status_msg": ""
    }
}
```

**python**
```python
{'content': '是的，郑钦文在巴黎奥运会上夺得了网球女子单打的金牌。根据最新的报道，郑钦文在2024年8月3日的巴黎奥运会网球女子单打决赛中战胜对手，成功夺得金牌，实现了中国选手在该项目上的历史性突破【0†source】【1†source】【2†source】。', 'role': 'assistant', 'annotations': [{'text': '【0†source】', 'url': 'http://m.cnhubei.com/content/2024-08/04/content_18248722.html', 'quote': '当地时间8月3日,在巴黎奥运会网球女子单打决赛中,郑钦文夺得金牌,实现了中国选手在该项目上的历史性突破. 新华社音视频部制作'}, {'text': '【1†source】', 'url': 'http://news.cnhubei.com/content/2024-08/04/content_18247840.html', 'quote': '通讯员 郭雨辰 张南左盈 实习生 刘嘉慧 张伊婷 北京时间8月4日,在刚刚结束的巴黎奥运会网球女子单打决赛中,郑钦文夺得金牌,实现了中国选手在该项目上的历史性突破. 据悉,郑钦文为华中科技大学体育学院2023级运动训练专业本科生,曾获得第19届亚运会网球女单冠军、2024澳大利亚网球公开赛女单亚军、WTA500郑州网球公开赛女单冠军等. 郑钦文此次比赛期间,华中科技大学300余名师生齐聚华科大光谷体育馆大多媒体教室集中观看比赛,共同为她加油! 来自该校电子信息与通信学院的2022级学生黄思颖在华中科技大学上了两年网球体育课,她今天来到光谷体育馆观赛.“我自己平时也打网球,也知道郑钦文很厉害,之前就达成澳网亚军的成就,所以今天想和大家一起见证她的新纪录!”黄思颖说. 新闻学院2021级学生陈祉旭和朋友也来到光谷体育馆观赛,“虽然之前没有打过网球,但是校友破纪录的时刻还是要和大家一起看的!一起看的氛围太好了,最后夺冠的瞬间真的觉得郑钦文校友太牛了!”她告诉记者,比赛过程中,拿着拍手器的她会在郑钦文每一次得分的时候为她喝彩. “郑钦文太棒了!身为华科大体育学院的教师,为她感到骄傲,她为我们新一代华科大学子树立了榜样!”郑钦文的学业指导老师曾洪涛介绍,郑钦文是该校体育学院2023级运动训练专业本科生,去年入学季正逢她打美网,杀入了八强,创造了自己的最好成绩.“钦文的训练和比赛任务非常重,从进校开始,我们就密切跟进她的训练团队,了解比赛日程和训练计划,制定了线上教学和线下定制式混合授课的模式.”在她训练和比赛的紧张时间段,能挤出时间参加线上课程的,就线上学习,学校也会同步录制课程视频发给她,方便她有时间时进行补习.在郑钦文和其父母的要求下,线下也会安排老师,跟着她的团队,进行一对一授课. 曾老师透露,郑钦文选择华中科技大学的重要原因之一,就是追随偶像李娜的步伐,继承学校的网球传统,创造属于自己的辉煌.她非常珍惜自己能进入华中科技大学学习的宝贵机会. 在奥运赛场,华科大法国校友会组织了后援团赴现场为郑钦文同学加油助威. 夺冠后,华科大第一时间向中国体育代表团和郑钦文同学致信表示热烈祝贺.贺信写道,郑钦文同学奥运夺金,是其个人运动生涯的重要里程碑,更是其未来的新起点.在此,华中科技大学向郑钦文同学致以热烈祝贺!此次她勇夺金牌,广大华科大人深感振奋、备受鼓舞,进一步激励了华科大人奋'}, {'text': '【2†source】', 'url': 'http://ent.people.com.cn/n1/2024/0804/c459176-40291915.html', 'quote': '人民网北京8月4日电 (记者李乃妍)北京时间8月4日凌晨,在巴黎奥运会网球女单决赛中,中国队选手郑钦文战胜克罗地亚队选手维基奇,摘得金牌.这是2004年雅典奥运会李婷/孙甜甜女双夺金后,中国网球队收获的又一枚奥运金牌,也是中国网球队夺得的首枚奥运会女单金牌. '}]}
```
#### DeepSeek-R1使用建议
使用DeepSeek-R1时，为获得最佳效果，不建议设置任何system prompt与system name。
**调用示例**
```python
import requests
import json
url = "https://api.minimax.chat/v1/text/chatcompletion_v2"

API_KEY = "请填写您的api key"
payload = json.dumps({
    "model": "DeepSeek-R1",
    "messages": [{"role": "user", "content": "帮我分析一下如何提升自己"}],
    "stream": True
})
headers = {
    'Authorization': f'Bearer {API_KEY}',
    'Content-Type': 'application/json'
}
response = requests.request("POST", url, headers=headers, data=payload, stream=True)
for chunk in response.iter_lines():
    if chunk:
        chunkStr = chunk.decode("utf-8")
        print(chunkStr)
        print("————————————————————")
```

**curl**
```curl
curl --location "https://api.minimax.chat/v1/text/chatcompletion_v2" \
--header "Content-Type: application/json" \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
    "model":"MiniMax-Text-01",
    "messages":[
      {
        "role":"user",
        "content":"帮我分析一下如何提升自己"
      }
    ]
  }'
```
##### 上下文拼接
在每一轮对话过程中，模型会输出思维链内容（reasoning_content）和最终回答（content）。在下一轮对话中，之前轮输出的思维链内容无需拼接到上下文中，如下图所示：
![image](https://filecdn.minimax.chat/public/1de313c0-698d-4bf8-9d0f-05c215a43f22.png)

## 文本生成
- 原始链接: https://platform.minimaxi.com/document/对话?key=66701d281d57f38758d581d0
- 文档ID: `66701d281d57f38758d581d0`
### 文本生成（ChatCompletion）
该API支持Openai参数定义和基础对话能力，还支持tools调用能力，可参考示例进行接入使用。支持通过 API 或者 OpenAI SDK进行接入，参考下方文档进行接入。
#### 支持模型
| 模型名称 | 最大token数（一次请求+回复的输入输出总token） |
| --- | --- |
| MiniMax-M1 | 1000192 |
| MiniMax-Text-01 | 1000192 |
#### 接口参数说明
`接口地址：https://api.minimaxi.com/v1/text/chatcompletion_v2`
##### 请求体（Request）参数
###### 返回(Response)参数
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| Authorization | string | 是 | `HTTP：Bearer Auth`<br>- Security Scheme Type: http<br>- HTTP Authorization Scheme: BearerAPI_key，可在[账户管理>接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)中查看。 |
| Content-Type | application/json | 是 | Content-Type |
| model | string | 是 | 调用的模型ID。目前支持取以下值：`MiniMax-M1``MiniMax-Text-01`。注：MiniMax-M1为推理模型，输出tokens较多，调用时建议使用流式输出以获得更稳定的网络连接。 |
| stream | bool | 否 | 是否通过流式分批返回结果。如果设置为true，结果分批返回，两个换行分割分批返回结果.默认为false |
| max_tokens | int64 (0,40000] | 否 | 最大生成token数，需要注意的是，这个参数并不会影响模型本身的生成效果，而是仅仅通过以截断超出的token的方式来实现功能。可选，目前默认取值MiniMax-M1为8192，MiniMax-Text-01为2048。如遇到length原因停止生成请修改该值。 |
| temperature | float，(0,1] | 否 | 较高的值将使输出更加随机，而较低的值将使输出更加集中和确定。<br>MiniMax-M1默认取值为1，MiniMax-Text-01默认取值为0.1。<br>MiniMax-M1推荐范围[0.8, 1]<br>MiniMax-Text-01参考：<br>低（0.01~0.2）：适合答案较明确的场景（如：知识问答、总结说明、情感分析、文本分类、大纲生成、作文批改）<br>⾼（0.7〜1）：适合答案较开放发散的场景 （如：营销文案生成、人设对话） |
| top_p | float，(0,1] | 否 | 采样方法，数值越小结果确定性越强；数值越大，结果越随机。取值可选，各模型默认为0.95 |
| mask_sensitive_info | bool | 否 | 对输出中易涉及隐私问题的文本信息进行打码，目前包括但不限于邮箱、域名、链接、证件号、家庭住址等，默认False，即不开启打码。 |
| messages | array | 是 | 对话的内容。<br>发送者的类型。取值如下：<br>system：定义智能体bot设定的role类型；<br>user：在对话阶段，用户发送消息的role类型；<br>assistant：在对话阶段，智能体发送回复的role类型；<br>tool：定义function的role类型；<br>发送者的名称。须提供具体名称以区分同一类型的不同具体角色。<br>发送的内容。当为纯文本输入时使用string类型，当需要image input时使用array类型。<br>内容类型，支持text、image_url两种<br>当图文混合输入时，文本类型的具体文字内容信息。<br>图片的url信息，url要保证公网可访问，也支持base64编码的图片信息<br>1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息<br>模型生成的tools信息。role需要为assistant。<br>function的调用信息对应的id.在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| messages.role | string | 是 | 发送者的类型。取值如下：<br>system：定义智能体bot设定的role类型；<br>user：在对话阶段，用户发送消息的role类型；<br>assistant：在对话阶段，智能体发送回复的role类型；<br>tool：定义function的role类型； |
| messages.name | string | 否 | 发送者的名称。须提供具体名称以区分同一类型的不同具体角色。 |
| messages.content | string or array | 是 | 发送的内容。当为纯文本输入时使用string类型，当需要image input时使用array类型。<br>内容类型，支持text、image_url两种<br>当图文混合输入时，文本类型的具体文字内容信息。<br>图片的url信息，url要保证公网可访问，也支持base64编码的图片信息<br>1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息 |
| messages.content.type | string | 否 | 内容类型，支持text、image_url两种 |
| messages.content.text | string | 否 | 当图文混合输入时，文本类型的具体文字内容信息。 |
| messages.content.image_url |  | 否 | 图片的url信息，url要保证公网可访问，也支持base64编码的图片信息<br>1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息 |
| messages.content.image_url.url | string | 否 | 1. 图片的url地址，需要保证公网可访问2. 也支持base64编码的图片信息 |
| messages.tool_calls |  | 否 | 模型生成的tools信息。role需要为assistant。<br>function的调用信息对应的id.在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| messages.tool_calls.id | string | 是 | function的调用信息对应的id.在触发function时会返回该id。 |
| messages.tool_calls.type | string | 是 | function的类型，仅支持function。 |
| messages.tool_calls.function |  | 是 | 模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| messages.tool_calls.function.name | string | 是 | function对应的名称。 |
| messages.tool_calls.function.arguments | string | 是 | function的具体调用信息。 |
| tool_choice | string | 否 | tool工具的开关.支持以下2种模式：1、none：不调用function call；2、auto：自动判断是否调用funcition call。 |
| tools |  | 否 | 支持的工具。<br>支持的工具类型，目前仅支持`function`<br>function的定义。<br>funciton的描述。<br>function的名称。<br>function的参数。 |
| tools.type | string | 是 | 支持的工具类型，目前仅支持`function` |
| tools.function |  | 是 | function的定义。<br>funciton的描述。<br>function的名称。<br>function的参数。 |
| tools.function.description | string | 是 | funciton的描述。 |
| tools.function.name | string | 是 | function的名称。 |
| tools.function.parameters | object | 否 | function的参数。 |
| response_format | object | 否 | 指定模型必须输出的格式。<br>该参数仅支持MiniMax-Text-01。<br>设置为`{ "type": "json_schema", "json_schema": {...} }`以实现结构化输出，这会确保模型返回的内容结构完全匹配您提供的JSON schema。<br>示例见右方代码区域。<br>定义的响应格式类型， 仅支持`"json_schema"`<br>描述输出格式的说明，用于模型确定如何以该格式进行输出。<br>输出格式的名称。必须由a-z、A-Z、0-9 组成，或包含下划线和破折号，最大长度为64。<br>定义输出格式，形式为JSON Schema object。<br>取值应为`"object"`<br>详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| response_format.type | string | 是 | 定义的响应格式类型， 仅支持`"json_schema"` |
| response_format.json_schema | object | 是 | 描述输出格式的说明，用于模型确定如何以该格式进行输出。<br>输出格式的名称。必须由a-z、A-Z、0-9 组成，或包含下划线和破折号，最大长度为64。<br>定义输出格式，形式为JSON Schema object。<br>取值应为`"object"`<br>详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| response_format.json_schema.description | string | 否 | 描述输出格式的说明，用于模型确定如何以该格式进行输出。 |
| response_format.json_schema.name | string | 是 | 输出格式的名称。必须由a-z、A-Z、0-9 组成，或包含下划线和破折号，最大长度为64。 |
| response_format.json_schema.schema | object | 否 | 定义输出格式，形式为JSON Schema object。<br>取值应为`"object"`<br>详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| response_format.json_schema.schema.type | string | 是 | 取值应为`"object"` |
| response_format.json_schema.schema.properties | object | 是 | 详细定义格式化输出所需的内容。<br>支持的类型包括：String、Array、Enum、Number、Integer、Object、Boolean。请注意，使用结构化输出时，所有字段或函数参数都必须指定为`required`。 |
| stream_options |  | 否 | 如果设置为True，将在最后流出一个额外的chunk。此块上的 usage 字段显示整个请求的token使用统计信息，而 choices 字段始终为空数组。所有其他块也将包含一个 usage 字段，但值为 null。 |
| stream_options.include_usage | boolean | 否 | 如果设置为True，将在最后流出一个额外的chunk。此块上的 usage 字段显示整个请求的token使用统计信息，而 choices 字段始终为空数组。所有其他块也将包含一个 usage 字段，但值为 null。 |
| id | string | 否 | 响应id |
| choices | array | 否 | 回复消息选项，里面包含对应的信息<br>模型生成停止的原因。目前支持取以下值：<br> 1、stop：正常结束；<br>2、tool_calls：触发function calling；<br>3、length：请求超过token长度。<br>回复消息的索引，目前仅支持1个选项。<br>回复信息详细内容。<br>回复的内容。<br>回复消息的角色，目前仅支持assistant。<br>模型生成的tools信息，role需要为assistant。<br>function的调用信息对应的id,在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| choices.finish_reason | string | 否 | 模型生成停止的原因。目前支持取以下值：<br> 1、stop：正常结束；<br>2、tool_calls：触发function calling；<br>3、length：请求超过token长度。 |
| choices.index | int | 否 | 回复消息的索引，目前仅支持1个选项。 |
| choices.message | dict | 否 | 回复信息详细内容。<br>回复的内容。<br>回复消息的角色，目前仅支持assistant。<br>模型生成的tools信息，role需要为assistant。<br>function的调用信息对应的id,在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| choices.message.content | string | 否 | 回复的内容。 |
| choices.message.role | string | 否 | 回复消息的角色，目前仅支持assistant。 |
| choices.message.tool_calls |  | 否 | 模型生成的tools信息，role需要为assistant。<br>function的调用信息对应的id,在触发function时会返回该id。<br>function的类型，仅支持function。<br>模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| choices.message.tool_calls.id | string | 否 | function的调用信息对应的id,在触发function时会返回该id。 |
| choices.message.tool_calls.type | string | 否 | function的类型，仅支持function。 |
| choices.message.tool_calls.function |  | 否 | 模型生成的function调用信息。<br>function对应的名称。<br>function的具体调用信息。 |
| choices.message.tool_calls.function.name | string | 否 | function对应的名称。 |
| choices.message.tool_calls.function.arguments | string | 否 | function的具体调用信息。 |
| created | int | 否 | 创建的时间戳。 |
| model | string | 否 | 聊天选择的模型。 |
| object | string | 否 | 对象类型为流式或非流式，取值为chat.completion或chat.completion.chunk。 |
| usage |  | 否 | 本次请求的消耗。<br>本次请求的消耗token。 |
| usage.total_tokens | int | 否 | 本次请求的消耗token。 |
| input_sensitive | bool | 否 | 输入命中敏感词。<br>如果输入内容严重违规，接口会返回内容违规错误信息，回复内容为空。<br>输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| input_sensitive.input_sensitive_type | int64 | 否 | 输入命中敏感词类型，当input_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| output_sensitive | bool | 否 | 输出命中敏感词。<br>如果输出内容严重违规，接口会返回内容违规错误信息，回复内容为空。<br>输出命中敏感词类型，当output_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| output_sensitive.output_sensitive_type | int64 | 否 | 输出命中敏感词类型，当output_sensitive为true时返回。取值为以下其一：1 严重违规；2 色情；3 广告；4 违禁；5 谩骂；6 暴恐；7 其他。 |
| base_resp | BaseResp | 否 | 错误状态码和详情。<br>状态码。<br>1000，未知错误；<br>1001，请求超时；<br>1002，触发RPM限流；<br>1004，鉴权失败；<br>1008，余额不足；<br>1013，服务内部错误；<br>1027，输出内容错误；<br>1039，Token限制；<br>2013，参数错误。<br>错误详情。 |
| base_resp.status_code | int64 | 否 | 状态码。<br>1000，未知错误；<br>1001，请求超时；<br>1002，触发RPM限流；<br>1004，鉴权失败；<br>1008，余额不足；<br>1013，服务内部错误；<br>1027，输出内容错误；<br>1039，Token限制；<br>2013，参数错误。 |
| base_resp.status_msg | string | 否 | 错误详情。 |
#### SDK 接入
Chat Completion v2接口支持通过OpenAI SDK进行接入，分为如下2步OpenAI SDK：
1. 安装 Python 3.7.1 或更高版本后，就可以安装 OpenAI Python 库。在终端/命令行中，运行：
```python
pip install --upgrade openai
```
2. 在[账户管理-接口密钥](https://platform.minimaxi.com/user-center/basic-information/interface-key)获取你的API Key。
3. 在接入代码使用OpenAI SDK参考代码示例进行接入和使用，更多示例可参考上方API说明中的python示例。
**请求示例**
```python
from openai import OpenAI

client = OpenAI(api_key="<minimax api key>", base_url="https://api.minimaxi.com/v1")

response = client.chat.completions.create(
  model="MiniMax-M1",
  messages=[
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Hello!"}
  ],
  stream=True
)

for chunk in response:
  print(chunk)
```
**返回示例**
```python
ChatCompletionChunk(id='02fd53267b5e09b5030cf1bc99f42d3a', choices=[Choice(delta=ChoiceDelta(content='你好', function_call=None, role='assistant', tool_calls=None), finish_reason=None, index=0, logprobs=None)], created=1722687527, model='MiniMax-M1', object='chat.completion.chunk', service_tier=None, system_fingerprint=None, usage=None, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0)
ChatCompletionChunk(id='02fd53267b5e09b5030cf1bc99f42d3a', choices=[Choice(delta=ChoiceDelta(content='！很高兴能和你交流。有什么问题或者话题想要讨论吗？我在这里帮助你。', function_call=None, role='assistant', tool_calls=None), finish_reason='stop', index=0, logprobs=None)], created=1722687527, model='MiniMax-M1', object='chat.completion.chunk', service_tier=None, system_fingerprint=None, usage=None, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0)
ChatCompletionChunk(id='02fd53267b5e09b5030cf1bc99f42d3a', choices=[Choice(delta=None, finish_reason='stop', index=0, logprobs=None, message={'content': '你好！很高兴能和你交流。有什么问题或者话题想要讨论吗？我在这里帮助你。', 'role': 'assistant'})], created=1722687527, model='MiniMax-M1', object='chat.completion', service_tier=None, system_fingerprint=None, usage=CompletionUsage(completion_tokens=None, prompt_tokens=None, total_tokens=87), input_sensitive=False, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0, base_resp={'status_code': 0, 'status_msg': ''})
```
#### 函数调用（function calling）
该功能可触发调用用户自有函数（如客户自有的会议记录查询函数），得到自有函数的参考内容 ，辅助大模型生成结果。可以帮助模型获取用户希望模型获取的相应信息，以使得模型返回更符合用户个人实际情况的回复。特别是在一些需要基于客户自有信息知识的问题回答上，并且需要使用构建了该信息知识的API，那么用户可以考虑使用function calling功能，使用场景示例如下：
> 1、query：“在会议ID：12345的会议上，A同学有没有提到大模型如何应用的事情？”
> 2、用户自有函数，函数功能：可基于会议ID和参会人检索原始的会议记录
> 3、参考内容：“会议号：12345，参会人：A同学，发言内容：我们要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。”
> 4、大模型生成结果：“A同学在会议ID：12345的会议上提到了大模型如何应用的事情。他说要在MM智能助理app中全面拥抱大模型，将其强大的自然语言处理能力和广泛的应用场景深入融合到app的各个功能模块，充分发挥大模型所带来的技术优势，进一步提升用户体验，提高工作效率，为更多用户提供智能化、便捷化的服务。”
以上，如果仅靠大模型来回答query时无法得到符合预期的回复，因此在这时就需要使用function calling功能。
##### Function Calling示例
具体流程如下：
（MiniMax-M1拥有单轮多function调用能力，详见下方示例）
1. 触发function，输出function输入参数；
2. 得到function的出参，给到模型参考进行总结回复。
**Tab: 1、触发function，输出function输入参数**

**请求示例（流式）**
```curl
curl --location 'https://api.minimaxi.com/v1/text/chatcompletion_v2' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
  "model": "MiniMax-M1",
  "messages": [
    {
      "role": "system",
      "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
      "role": "user",
      "content": "广州天气怎么样"
    }
  ],
  "tools": [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
],
  "tool_choice": "auto"
}'
```

**python**
```python
from openai import OpenAI

client = OpenAI(
    api_key="",  # <--在这里使用MiniMax账户管理-接口密钥中API KEY进行接入
    base_url="https://api.minimaxi.com/v1",
)

tools = [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
]

messages = [
    {"role": "system",
     "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"},
    {"role": "user", "content": "广州天气怎么样？"}
]

completion = client.chat.completions.create(
  model="MiniMax-M1",
  messages=messages,
  tools=tools,
  tool_choice = 'auto'
)
print(completion)
```

**Tab: 2、得到function的出参，给到模型参考进行总结回复**

**请求示例（流式）**
```curl
curl --location 'https://api.minimaxi.com/v1/text/chatcompletion_v2' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
  "model": "MiniMax-M1",
  "messages": [
    {
      "role": "system",
      "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
      "role": "user",
      "content": "广州天气怎么样"
    },
        {
        "role": "assistant",
        "tool_calls": [{
            "id": "call_function_1849136893",
            "type": "function",
            "function": {
                "name": "get_current_weather",
                "arguments": "{\"location\":\"广州\"}"
            }
        }]
    },
    {
        "role": "tool",
        "tool_call_id": "call_function_1849136893",
        "content": "多云，28~37℃，无持续风向<3级，空气质量优"
    }
  ],
  "tools": [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
],
  "tool_choice": "auto"
}'
```

**python**
```python
from openai import OpenAI

client = OpenAI(
    api_key="",  # <--在这里使用MiniMax账户管理-接口密钥中API KEY进行接入
    base_url="https://api.minimaxi.com/v1",
)

tools = [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
]

messages = [
    {
        "role": "system",
        "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
        "role": "user",
        "content": "广州天气怎么样？"
     },
    {"role": "user", "content": "广州天气怎么样？"},
    {
        "role": "assistant",
        "tool_calls": [{
            "id": "call_function_1849136893",
            "type": "function",
            "function": {
                "name": "get_current_weather",
                "arguments": "{\"location\":\"广州\"}"
            }
        }]
    },
    {
        "role": "tool",
        "tool_call_id": "call_function_1849136893",
        "content": "多云，28~37℃，无持续风向<3级，空气质量优"
    }
]

completion = client.chat.completions.create(
  model="MiniMax-M1",
  messages=messages,
  tools=tools,
  tool_choice = 'auto'
)
print(completion)
```

**Tab: 多function调用**

**多function调用**
```curl
curl --location 'https://api.minimaxi.com/v1/text/chatcompletion_v2' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
  "model": "MiniMax-M1",
  "messages": [
    {
      "role": "system",
      "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
      "role": "user",
      "content": "广州和上海天气怎么样"
    }
  ],
  "tools": [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
],
  "tool_choice": "auto"
}'
```

**Tab: 多function调用（回传）**

**多function调用（回传）**
```curl
curl --location 'https://api.minimaxi.com/v1/text/chatcompletion_v2' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $MiniMax_API_KEY" \
--data '{
  "model": "MiniMax-M1",
  "messages": [
    {
      "role": "system",
      "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"
    },
    {
      "role": "user",
      "content": "广州和上海天气怎么样"
    },
        {
        "role": "assistant",
        "tool_calls": [{
            "id": "call_function_1849136893",
            "type": "function",
            "function": {
                "name": "get_current_weather",
                "arguments": "{\"location\":\"广州\"}"
            },
{
            "id": "call_function_1849136892",
            "type": "function",
            "function": {
                "name": "get_current_weather",
                "arguments": "{\"location\":\"上海\"}"
            }
        }]
    },
    {
        "role": "tool",
        "tool_call_id": "call_function_1849136893",
        "content": "多云，28~37℃，无持续风向<3级，空气质量优"
    },
{
        "role": "tool",
        "tool_call_id": "call_function_1849136892",
        "content": "晴天，20~27℃，无持续风向<3级，空气质量优"
    }
  ],
  "tools": [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "获取近期某一地点的天气情况",
            "parameters": "{\"type\":\"object\", \"properties\":{\"location\":{\"type\":\"string\", \"description\":\"某一个城市，比如北京、上海\"}}, \"required\":[\"location\"]}"
        }
    }
],
  "tool_choice": "auto"
}'
```
**Tab: 1、触发function，输出function输入参数**

**返回示例（流式）**
```curl
{
    "id": "02ffb000ff4b7202faae09b999b27c16",
    "choices": [
        {
            "finish_reason": "tool_calls",
            "index": 0,
            "message": {
                "role": "assistant",
                "tool_calls": [
                    {
                        "id": "call_function_6835316862",
                        "type": "function",
                        "function": {
                            "name": "get_current_weather",
                            "arguments": "{\"location\": \"广州\"}"
                        }
                    }
                ]
            }
        }
    ],
    "created": 1722842368,
    "model": "MiniMax-M1",
    "object": "chat.completion",
    "usage": {
        "total_tokens": 153
    },
    "input_sensitive": false,
    "output_sensitive": false,
    "input_sensitive_type": 0,
    "output_sensitive_type": 0,
    "base_resp": {
        "status_code": 0,
        "status_msg": ""
    }
}
```

**python**
```python
ChatCompletion(id='02ffacd947f8a288ad6f0c1918d297a8', choices=[Choice(finish_reason='tool_calls', index=0, logprobs=None, message=ChatCompletionMessage(content=None, role='assistant', function_call=None, tool_calls=[ChatCompletionMessageToolCall(id='call_function_0304055004', function=Function(arguments='{"location": "广州"}', name='get_current_weather'), type='function')]))], created=1722841561, model='MiniMax-M1', object='chat.completion', service_tier=None, system_fingerprint=None, usage=CompletionUsage(completion_tokens=None, prompt_tokens=None, total_tokens=154), input_sensitive=False, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0, base_resp={'status_code': 0, 'status_msg': ''})
```

**Tab: 2、得到function的出参，给到模型参考进行总结回复**

**返回示例**
```curl
{
    "id": "02ffb37f1e7189808524269fc6fb9f9f",
    "choices": [
        {
            "finish_reason": "stop",
            "index": 0,
            "message": {
                "content": "广州现在的天气是多云，气温在28~37℃之间，没有持续的风向，风力小于3级，空气质量为优。",
                "role": "assistant"
            }
        }
    ],
    "created": 1722843264,
    "model": "MiniMax-M1",
    "object": "chat.completion",
    "usage": {
        "total_tokens": 223
    },
    "input_sensitive": false,
    "output_sensitive": false,
    "input_sensitive_type": 0,
    "output_sensitive_type": 0,
    "base_resp": {
        "status_code": 0,
        "status_msg": ""
    }
}
```

**python**
```python
ChatCompletion(id='02ffb2ec3574bb88b257536afb5f9876', choices=[Choice(finish_reason='stop', index=0, logprobs=None, message=ChatCompletionMessage(content='广州目前的天气是多云，气温在28~37℃之间，无持续风向小于3级，空气质量优。', role='assistant', function_call=None, tool_calls=None))], created=1722843117, model='MiniMax-M1', object='chat.completion', service_tier=None, system_fingerprint=None, usage=CompletionUsage(completion_tokens=None, prompt_tokens=None, total_tokens=232), input_sensitive=False, output_sensitive=False, input_sensitive_type=0, output_sensitive_type=0, base_resp={'status_code': 0, 'status_msg': ''})
```

**Tab: 多function调用**

**多function调用**
```json
{
    "id": "02ffb000ff4b7202faae09b999b27c16",
    "choices": [
        {
            "finish_reason": "tool_calls",
            "index": 0,
            "message": {
                "role": "assistant",
                "tool_calls": [
                    {
                        "id": "call_function_6835316862",
                        "type": "function",
                        "function": {
                            "name": "get_current_weather",
                            "arguments": "{\"location\": \"广州\"}"
                        }
                    },
                    {
                        "id": "call_function_6835316863",
                        "type": "function",
                        "function": {
                            "name": "get_current_weather",
                            "arguments": "{\"location\": \"上海\"}"
                        }
                    }
                ]
            }
        }
    ],
    "created": 1722842368,
    "model": "MiniMax-M1",
    "object": "chat.completion",
    "usage": {
        "total_tokens": 263
    },
    "input_sensitive": false,
    "output_sensitive": false,
    "input_sensitive_type": 0,
    "output_sensitive_type": 0,
    "base_resp": {
        "status_code": 0,
        "status_msg": ""
    }
}
```

**Tab: 多function调用（最终回复）**

**多function调用（最终回复）**
```json
{
    "id": "02ffb37f1e7189808524269fc6fb9f9f",
    "choices": [
        {
            "finish_reason": "stop",
            "index": 0,
            "message": {
                "content": "广州现在的天气是多云，气温在28~37℃之间，没有持续的风向，风力小于3级，空气质量为优；上海现在的天气是晴天，气温在20~27℃之间，没有持续的风向，风力小于3级，空气质量为优。",
                "role": "assistant"
            }
        }
    ],
    "created": 1722843264,
    "model": "MiniMax-M1",
    "object": "chat.completion",
    "usage": {
        "total_tokens": 223
    },
    "input_sensitive": false,
    "output_sensitive": false,
    "input_sensitive_type": 0,
    "output_sensitive_type": 0,
    "base_resp": {
        "status_code": 0,
        "status_msg": ""
    }
}
```

## 批量处理
- 原始链接: https://platform.minimaxi.com/document/批量处理?key=640e0c9c5f918b4f6c4e2d57
- 文档ID: `640e0c9c5f918b4f6c4e2d57`
### Batch接口
虽然一些使用MiniMax平台的场景需要您发送同步请求，但在许多情况下，请求并不需要立即响应，或者速率限制会阻止您快速执行大量查询。批处理作业在以下使用案例中通常很有帮助：批量的数据打标、舆情分析、对话信息提取、通话质检等。与同步API相比的优势：支持更高的TPM&RPM速率限制，以解决您在大批量请求非实时任务时无法使用同步API的问题。
#### 操作步骤
1. 准备批处理文件，批处理从 .jsonl 文件开始，其中每行都包含对 API 的单个请求的详细信息。目前，可用的终端节点为 v1/text/chatcompletion_v2。对于给定的输入文件，每行的 body 字段中的参数与底层终端节点的参数相同。每个请求都必须包含一个唯一的 custom_id 值，您可以在完成后使用该值来引用结果。下面是一个包含 2 个请求的输入文件示例。请注意，每个 Importing 文件只能包含对单个模型的请求。
**批处理文件**
```json
{"custom_id": "request-1", "method": "POST", "url": "/v1/text/chatcompletion_v2", "body": {"model": "MiniMax-M1", "messages": [{"role": "system", "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"},{"role": "user", "content": "你好"}],"max_tokens": 1000}}
{"custom_id": "request-2", "method": "POST", "url": "/v1/text/chatcompletion_v2", "body": {"model": "MiniMax-M1", "messages": [{"role": "system", "content": "MM智能助理是一款由MiniMax自研的，没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司，一直致力于进行大模型相关的研究。"},{"role": "user", "content": "你好呀，请问你今天心情怎么样？"}],"max_tokens": 1000}}
```
2. 上传批量文件，您必须先上传输入文件，以便在启动批处理时可以正确引用它。使用 Files API 上传 .jsonl 文件。
**upload file**
```curl
curl --location 'http://api.minimaxi.com/v1/files/upload' \
--header 'Authorization: ---' \
--form 'purpose="batch"' \
--form 'file=@"/Users/minimax/py/dafei/xx.jsonl"'
```
3. 创建批处理任务，成功上传输入文件后，您可以使用输入 File 对象的 ID 创建批处理。在这种情况下，我们假设文件 ID 为 193117143535621。目前，完成窗口只能设置为 24 小时。
**创建batch任务**
```curl
curl --location 'http://api.minimaxi.com/v1/batches' \
--header 'Content-Type: application/json' \
--header 'Authorization: ---' \
--data '{
    "input_file_id": "193117143535621",
    "endpoint": "/v1/text/chatcompletion_v2",
    "completion_window":"24h",
    "callback_url":"http://minimax/xxx"
}'
```
会输出一个batch object，里面会包含详细的batch object信息。
**batch object**
```json
{
    "id": "198533392228705",
    "object": "batch",
    "endpoint": "/v1/text/chatcompletion_v2",
    "errors": null,
    "input_file_id": "193117143535621",
    "completion_window": "24h",
    "status": "validating",
    "output_file_id": null,
    "error_file_id": null,
    "created_at": 1725438458,
    "in_progress_at": null,
    "expires_at": 1725524858,
    "finalizing_at": null,
    "completed_at": null,
    "failed_at": null,
    "expired_at": null,
    "cancelling_at": null,
    "cancelled_at": null,
    "request_counts": {
        "total": 0,
        "completed": 0,
        "failed": 0
    },
    "callback_url":xxx
}
```
4. 检索batch生成任务状态，您可以随时检查批处理的状态，这也将返回一个 Batch 对象。状态会包含validating、failed、in_progress、completed等
**检索batch任务状态**
```curl
curl --location 'http://api.minimaxi.com/v1/batches/198533392228705' \
--header 'Authorization: ---'
```
#### 创建Batch任务
`POST https://api.minimaxi.com/v1/batches`
##### 支持模型和token数
| 模型名称 | 最大token数（一次请求+回复的输入输出总token） |
| --- | --- |
| MiniMax-M1 | 1000192 |
##### 接口参数说明
###### Returns
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| input_file_id | string | 是 | 使用file upload API生成的file_id,即request input object对应的file_id。（单次请求最多包含50000条子任务，并且.jsonl文件需小于100mb。） |
| endpoint | string | 是 | 批处理的API端点，仅支持/v1/text/chatcompletion_v2 |
| completion_window | string | 是 | 任务执行的时间窗口（从提交到结束执行），时间窗，目前仅支持24h |
| callback_url | string | 否 | 发起任务创建请求时，MiniMax服务器将向请求地址发送一条包含校验字段的请求，当请求地址收到该 POST 校验请求时，需要提取出`challenge`值，并在3秒内返回包含`challenge`值的响应数据，否则将无法通过校验，创建任务失败。<br>响应数据示例：{ "challenge": "1a3cs1j-601a-118y-ch52-o48911eabc3u"}<br>成功配置回调请求地址后，当batch任务结束时，MiniMax服务器将向此地址发送包含batch object信息的回调请求。 |
| the batch object |  | 否 | the batch object相关信息 |
#### 检索Batch任务
`GET https://api.minimaxi.com/v1/batches/{batch_id}`
###### Returns
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| batch_id | string | 是 | 要检索的批处理的 ID。 |
| the batch object | string | 否 | the batch object相关信息 |
#### 取消Batch任务
`POST https://api.minimaxi.com/v1/batches/{batch_id}/cancel`
###### Returns
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| batch_id | string | 是 | 要取消的批处理ID。 |
| the batch object |  | 否 | the batch object相关信息 |
#### 列出Batch任务
`GET https://api.minimaxi.com/v1/batches`
###### Returns
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| after | string | 否 | 用于分页的游标。after 是一个对象 ID，用于定义您在列表中的位置。例如，如果您发出列表请求并收到 100 个对象，以 obj_foo 结尾，则您的后续调用可以包含 after=obj_foo 以便获取列表的下一页。 |
| limit | int | 否 | 对要返回的对象数的限制。限制的范围介于 1 和 100 之间，默认值为 20。 |
| the list of batch object |  | 否 | the batch object的列表 |
#### Batch对象结构
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| id | string | 否 | 批处理ID。 |
| object | string | 否 | 对象类型,始终为batch。 |
| endpoint | string | 否 | 批处理使用的 API 端点。 |
| errors | list | 否 | 错误信息<br>对象类型，始终为 list。<br>标识错误类型的错误代码。<br>一条人类可读的消息，提供有关错误的更多详细信息。<br>导致错误的参数的名称（如果适用）。<br>发生错误的输入文件的行号（如果适用）。 |
| errors.object | string | 是 | 对象类型，始终为 list。 |
| errors.data | array | 是 | 标识错误类型的错误代码。<br>一条人类可读的消息，提供有关错误的更多详细信息。<br>导致错误的参数的名称（如果适用）。<br>发生错误的输入文件的行号（如果适用）。 |
| errors.data.code | string | 是 | 标识错误类型的错误代码。 |
| errors.data.message | string | 是 | 一条人类可读的消息，提供有关错误的更多详细信息。 |
| errors.data.param | string | 是 | 导致错误的参数的名称（如果适用）。 |
| errors.data.line | string | 是 | 发生错误的输入文件的行号（如果适用）。 |
| input_file_id | string | 否 | 批处理的输入文件的ID。 |
| completion_window | string | 否 | 处理批处理的时间范围。 |
| status | string | 否 | 批处理的当前状态。 |
| output_file_id | string | 否 | 包含成功执行的请求的输出的文件的ID。 |
| error_file_id | string | 否 | 包含错误请求输出的文件的ID。 |
| created_at | int | 否 | 创建批处理时的 Unix 时间戳（以秒为单位）。 |
| expires_at | int | 否 | 批处理到期时间的 Unix 时间戳（以秒为单位）。 |
| finalizing_at | int | 否 | 批处理开始完成时的 Unix 时间戳（以秒为单位）。 |
| completed_at | int | 否 | 批处理完成时的 Unix 时间戳（以秒为单位）。 |
| failed_at | int | 否 | 批处理失败时的 Unix 时间戳（以秒为单位）。 |
| expired_at | int | 否 | 批处理到期时的 Unix 时间戳（以秒为单位）。 |
| cancelling_at | int | 否 | 批处理开始取消时的 Unix 时间戳（以秒为单位）。 |
| cancelled_at | int | 否 | 取消批处理时的 Unix 时间戳（以秒为单位）。 |
| request_counts | object | 否 | 该请求针对批处理中的不同状态进行计数。 |
#### Batch输入对象结构
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| custom_id | string | 否 | 开发人员提供的每个请求 ID，用于将输出与输入匹配。对于批处理中的每个请求，必须是唯一的。 |
| method | string | 否 | 要用于请求的 HTTP 方法。目前仅支持 POST。 |
| url | string | 否 | 要用于请求的 OpenAI API 相对 URL。目前支持 /v1/text/chatcompletion_v2 |
#### Batch输出对象结构
| 参数 | 类型 | 必填 | 描述 |
| --- | --- | --- | --- |
| id | string | 否 | 同上 |
| custom_id | string | 否 | 开发人员提供的每个请求 ID，用于将输出与输入匹配。 |
| reponse | object | 否 | 响应的 HTTP 状态代码<br>API 请求的唯一标识符。联系支持人员时，请附上此请求 ID。<br>响应的 JSON 正文 |
| reponse.status_code | string | 是 | 响应的 HTTP 状态代码 |
| reponse.request_id | string | 是 | API 请求的唯一标识符。联系支持人员时，请附上此请求 ID。 |
| reponse.body | string | 是 | 响应的 JSON 正文 |
| error | object | 否 | 机器可读的错误代码。<br>人类可读的错误消息。 |
| error.code | string | 是 | 机器可读的错误代码。 |
| error.message | string | 是 | 人类可读的错误消息。 |
