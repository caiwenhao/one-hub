// Mock data for HotModels components
// 说明：为还原热门模型页面截图中的卡片效果，这里为“明星模型”补充了
// 1）更贴近设计稿的名称与描述
// 2）`useCases` 应用场景（用于 featured 卡片展示）
// 3）`iconColor` 作为卡片与标签的主色
// 4）`tag`（如：🔥 火爆）用于右上角胶囊标签
export const categories = [
  {
    id: 'text',
    name: '文本生成',
    description: '强大的文本理解与生成能力',
    icon: 'solar:document-text-bold'
  },
  {
    id: 'image',
    name: '图像生成',
    description: '创意图像生成与编辑',
    icon: 'solar:camera-bold'
  },
  {
    id: 'video',
    name: '视频生成',
    description: '智能视频创作与编辑',
    icon: 'solar:video-frame-play-vertical-bold'
  },
  {
    id: 'audio',
    name: '音频处理',
    description: '语音合成与音频生成',
    icon: 'solar:microphone-bold'
  },
  {
    id: 'embedding',
    name: '向量嵌入',
    description: '文本向量化与相似度计算',
    icon: 'solar:code-bold'
  }
];

// 模型数据
const baseModels = [
  // 文本 - 与截图一致的三张明星卡片
  {
    id: 'gpt-5',
    name: 'GPT-5',
    provider: 'OpenAI',
    abbr: 'O',
    category: 'text',
    description:
      '2025 年最强通用模型，混合多模态架构按任务难度自动路由，SWE-bench 74.9%。在复杂前端开发与多步推理上表现卓越，幻觉率降至约 1.4%。',
    price: '—',
    rating: 4.9,
    tags: ['推荐', '最新'],
    useCases: ['复杂编程任务', '前端原型开发', '深度推理科研/数学', '企业级 API 集成'],
    icon: 'simple-icons:openai',
    iconColor: '#3B82F6',
    tag: { type: 'hot', label: '🔥 火爆' }
  },
  {
    id: 'claude-sonnet-4',
    name: 'Claude Sonnet 4.5',
    provider: 'Anthropic',
    abbr: 'A',
    category: 'text',
    description:
      '面向自主智能体的顶级编码模型，长时专注与工具调度能力突出，支持 64K 输出 token，在计算机使用与浏览器自动化方面领先。',
    price: '—',
    rating: 4.8,
    tags: ['安全', '推理'],
    useCases: ['全生命周期软件开发', '漏洞自动修补', '长时运行智能体', '金融分析与风控'],
    icon: 'simple-icons:anthropic',
    iconColor: '#F59E0B',
    tag: { type: 'hot', label: '🔥 火爆' }
  },
  {
    id: 'gemini-2-5-pro',
    name: 'Gemini 2.5 Pro',
    provider: 'Google',
    abbr: 'G',
    category: 'text',
    description:
      '原生多模态与 100 万 token 上下文，擅长超长文档与代码理解；WebDev Arena 排名领先，适配企业级复杂场景。',
    price: '—',
    rating: 4.7,
    tags: ['多模态', '性价比'],
    useCases: ['Web 应用开发', '大规模代码库分析', '视频内容理解', '企业文档处理'],
    icon: 'simple-icons:google',
    iconColor: '#6366F1',
    tag: { type: 'hot', label: '🔥 火爆' }
  },
  // 新增文本生成模型
  {
    id: 'deepseek-v3-1',
    name: 'deepseek-v3.1',
    provider: 'DeepSeek',
    abbr: 'DS',
    category: 'text',
    description:
      '支持“思考/非思考”模式切换与 100 万 token 上下文，多步推理显著增强，覆盖 100+ 语言，幻觉率显著降低。',
    rating: 4.7,
    tags: ['推理', '高效'],
    icon: 'solar:cpu-bolt-bold',
    iconColor: '#0EA5E9',
    tag: { type: 'recommended', label: '⭐ 推荐' }
  },
  {
    id: 'kimi-k2',
    name: 'kimi-k2',
    provider: 'Moonshot AI',
    abbr: 'K2',
    category: 'text',
    description:
      '1T 参数 MoE 架构（激活约 32B），128K 上下文，面向工具调用与自主任务执行优化，具备极高性价比。',
    rating: 4.6,
    tags: ['长文本', '对话'],
    icon: 'solar:chat-line-bold',
    iconColor: '#8B5CF6',
    tag: { type: 'new', label: '🆕 上新' }
  },
  {
    id: 'gpt-4o',
    name: 'GPT-4o',
    provider: 'OpenAI',
    abbr: 'O',
    category: 'text',
    description:
      '多模态旗舰，实时处理音频/视觉/文本，低延迟语音交互与摄像头场景理解出色，覆盖广泛通用场景。',
    rating: 4.8,
    tags: ['多模态', '通用'],
    icon: 'simple-icons:openai',
    iconColor: '#16a34a',
    tag: { type: 'recommended', label: '⭐ 推荐' }
  },
  {
    id: 'gpt-image-1',
    name: 'GPT-Image-1',
    provider: 'OpenAI',
    abbr: 'GI',
    category: 'image',
    description: '2025 年 4 月发布，基于 GPT-4o 多模态架构，采用自回归替代扩散，最高 4096×4096。GIE-Bench 功能正确性领先，擅长精准文本渲染与细粒度编辑。',
    rating: 4.7,
    tags: ['文本到图像', '多模态'],
    features: ['文本到图像', '图像编辑', '变体生成', '高分辨率'],
    icon: 'simple-icons:openai',
    iconColor: '#22c55e',
    tag: { type: 'recommended', label: '⭐ 推荐' }
  },
  {
    id: 'midjourney',
    name: 'Midjourney V7',
    provider: 'Midjourney',
    abbr: 'MJ',
    category: 'image',
    description: '2025 年 4 月完全重构并于 6 月成为默认模型，架构与数据集全面升级。文本渲染与人体解剖显著提升，新增 Draft 与 Omni Reference 能力。',
    price: '¥0.08/图片',
    rating: 4.9,
    tags: ['艺术', '热门'],
    features: ['艺术风格', '细节丰富', '创意设计', '多样化'],
    icon: 'simple-icons:discord',
    iconColor: '#ef4444'
  },
  {
    id: 'viduq1',
    name: 'Vidu Q1',
    provider: 'Vidu',
    abbr: 'VQ',
    category: 'image',
    description: '新一代视频生成模型，稳定输出 5 秒 24 帧 1080P，单个视频约 $0.4。显著改善清晰度，缓解手部畸变与抖动，照片级渲染。',
    rating: 4.5,
    tags: ['高质量', '快速'],
    features: ['文本到图像', '多风格', '高分辨率'],
    icon: 'solar:camera-bold',
    iconColor: '#8B5CF6',
    tag: { type: 'new', label: '🆕 上新' }
  },
  {
    id: 'doubao-seedream-4-0',
    name: 'Doubao-SeeDream-4.0',
    provider: 'Doubao',
    abbr: 'DB',
    category: 'image',
    description: '2025 年 9 月发布的多模态图像引擎，文生图与编辑评测双榜首。最高 4K 输出，2K 生成约 1.8 秒，统一架构兼顾生成与编辑。',
    rating: 4.6,
    tags: ['高效', '稳定'],
    features: ['文本到图像', '高分辨率', '一致性强'],
    icon: 'solar:camera-bold',
    iconColor: '#0ea5e9'
  },
  {
    id: 'gpt-audio',
    name: 'gpt-audio',
    provider: 'OpenAI',
    abbr: 'GA',
    category: 'audio',
    description: 'OpenAI 音频多模态模型，兼顾识别、合成与指令理解。',
    rating: 4.6,
    tags: ['识别', '合成'],
    features: ['语音识别', '语音合成', '多语言'],
    icon: 'solar:microphone-bold',
    iconColor: '#22d3ee',
    tag: { type: 'recommended', label: '⭐ 推荐' }
  },
  {
    id: 'doubao-simultaneous',
    name: 'doubao-同声传译',
    provider: 'Doubao',
    abbr: 'DB',
    category: 'audio',
    description: '面向实时翻译与会议场景的同声传译模型，延迟低、稳定性高。',
    rating: 4.5,
    tags: ['同传', '实时'],
    features: ['实时翻译', '多语言', '低延迟'],
    icon: 'solar:microphone-bold',
    iconColor: '#0ea5e9'
  },
  {
    id: 'klingai-audio',
    name: 'klingai',
    provider: 'Kling AI',
    abbr: 'KA',
    category: 'audio',
    description: 'Kling AI 音频处理模型，覆盖识别、合成与增强。',
    rating: 4.4,
    tags: ['识别', '合成'],
    features: ['语音识别', '语音合成', '音频增强'],
    icon: 'solar:microphone-bold',
    iconColor: '#8B5CF6'
  },
  {
    id: 'speech-2-5-hd-preview',
    name: 'speech-2.5-hd-preview',
    provider: 'OpenAI',
    abbr: 'S2',
    category: 'audio',
    description: '高清语音合成预览版本，更自然的音色与更低噪声。',
    rating: 4.6,
    tags: ['高清', '合成'],
    features: ['高质量TTS', '多音色', '低噪声'],
    icon: 'solar:microphone-bold',
    iconColor: '#10b981'
  },
  // 向量嵌入（按需更新）
  {
    id: 'text-embedding-3-small',
    name: 'text-embedding-3-small',
    provider: 'OpenAI',
    abbr: 'O',
    category: 'embedding',
    description: 'OpenAI 小尺寸嵌入模型，性价比高，适合大规模检索。',
    rating: 4.5,
    tags: ['向量化', '检索'],
    features: ['文本向量化', '语义搜索', '相似度计算'],
    icon: 'simple-icons:openai',
    iconColor: '#94a3b8'
  },
  {
    id: 'text-embedding-3-large',
    name: 'text-embedding-3-large',
    provider: 'OpenAI',
    abbr: 'O',
    category: 'embedding',
    description: 'OpenAI 大尺寸嵌入模型，更高精度，适合高要求场景。',
    rating: 4.7,
    tags: ['向量化', '高精度'],
    features: ['文本向量化', '语义搜索', '知识库构建'],
    icon: 'simple-icons:openai',
    iconColor: '#64748b'
  },
  {
    id: 'text-embedding-ada-002',
    name: 'text-embedding-ada-002',
    provider: 'OpenAI',
    abbr: 'O',
    category: 'embedding',
    description: '经典嵌入模型，兼顾质量与成本。',
    rating: 4.4,
    tags: ['向量化', '通用'],
    features: ['文本向量化', '语义搜索', '相似度计算'],
    icon: 'simple-icons:openai',
    iconColor: '#94a3b8'
  },
  {
    id: 'doubao-embedding',
    name: 'doubao-embedding',
    provider: 'Doubao',
    abbr: 'DB',
    category: 'embedding',
    description: 'Doubao 嵌入模型，中文表现优秀，适合检索与匹配。',
    rating: 4.5,
    tags: ['中文优化', '检索'],
    features: ['文本向量化', '语义理解', '相似度计算'],
    icon: 'solar:translate-bold',
    iconColor: '#10b981'
  },
  {
    id: 'doubao-embedding-large',
    name: 'doubao-embedding-large',
    provider: 'Doubao',
    abbr: 'DB',
    category: 'embedding',
    description: 'Doubao 大尺寸嵌入模型，更高精度与稳定性。',
    rating: 4.6,
    tags: ['中文优化', '高精度'],
    features: ['文本向量化', '语义搜索', '知识库构建'],
    icon: 'solar:translate-bold',
    iconColor: '#059669'
  }
];

// 分类模型数据
export const mockModels = {
  // 截图中的“明星模型”仅展示前三个
  featured: baseModels.slice(0, 3),
  textModels: baseModels.filter(model => model.category === 'text'),
  imageModels: baseModels.filter(model => model.category === 'image'),
  videoModels: [
    // 指定顺序：sora-2 / Veo 3.1 / doubao-seedance / kling / Hailuo-02 / vidu2.0
    {
      id: 'sora-2',
      name: 'Sora-2',
      provider: 'OpenAI',
      category: 'video',
      description: '2025 年 9 月发布，最长 20 秒电影级视频并同步对话/音效/环境声；物理一致性强（如真实反弹），支持 Cameos 将特定人物与声音融入场景。',
      rating: 4.8,
      tags: ['长镜头', '高保真'],
      features: ['文本到视频', '物理一致性', '高分辨率'],
      icon: 'solar:video-frame-play-vertical-bold',
      iconColor: '#3B82F6'
    },
    {
      id: 'veo-3-1',
      name: 'Veo 3.1',
      provider: 'Google',
      category: 'video',
      description: '2025 年 10 月发布，最长 2 分钟 1080p，快速/标准双模式；支持 1–3 张参考图保证角色一致，场景扩展可无缝续接片段。',
      rating: 4.6,
      tags: ['创意', '多风格'],
      features: ['文本到视频', '风格控制', '快速渲染'],
      icon: 'solar:video-frame-play-vertical-bold',
      iconColor: '#0EA5E9'
    },
    {
      id: 'doubao-seedance',
      name: 'Doubao-Seedance',
      provider: 'Doubao',
      category: 'video',
      description: '2025 年 6 月旗舰，Video Arena 榜首；5 秒 1080p ≈ 41.4 秒（L20），原生多镜头叙事，统一架构支持文生/图生并以专用 RLHF 精准对齐提示。',
      rating: 4.4,
      tags: ['角色动作', '连贯性'],
      features: ['视频生成', '角色舞蹈', '动作迁移', '风格化'],
      icon: 'solar:video-frame-play-vertical-bold',
      iconColor: '#F59E0B'
    },
    {
      id: 'kling',
      name: 'Kling',
      provider: 'Kling',
      category: 'video',
      description: '最新 2.5 Turbo（1.6 版性能 +195%），文本/图像/视频多模态输入；支持对象级编辑（增删替换）并集成音效与 Kolors 2.0。',
      rating: 4.5,
      tags: ['视频生成', '高清输出'],
      features: ['文本到视频', '图像到视频', '高清输出', '长镜头'],
      icon: 'solar:video-frame-play-vertical-bold',
      iconColor: '#22c55e'
    },
    {
      id: 'hailuo-02',
      name: 'Hailuo-02',
      provider: 'Hailuo',
      category: 'video',
      description: '2025 年 6 月电影级模型，图生视频榜第二（仅次于 Seedance）；NCR 架构训练/推理提速约 2.5×，更大参数与数据带来更高保真。',
      rating: 4.3,
      tags: ['一致性', '细节'],
      features: ['文本到视频', '一致性', '高分辨率'],
      icon: 'solar:video-frame-play-vertical-bold',
      iconColor: '#8B5CF6'
    },
    {
      id: 'vidu-2-0',
      name: 'Vidu2.0',
      provider: 'Vidu',
      category: 'video',
      description: '2025 年 3 月发布，闪电/电影双模式；专注 2–8 秒短视频，10–20 秒出片，单图生成连续动作与表情，插值保证运动平滑。',
      rating: 4.4,
      tags: ['高质量', '快速'],
      features: ['文本到视频', '场景控制', '快速渲染', '多风格'],
      icon: 'solar:video-frame-play-vertical-bold',
      iconColor: '#ef4444'
    }
  ],
  audioModels: baseModels.filter(model => model.category === 'audio'),
  embeddingModels: baseModels.filter(model => model.category === 'embedding')
};

export default mockModels;
