import PropTypes from 'prop-types';
import { useState, useEffect, useRef, useMemo } from 'react';
import { CHANNEL_OPTIONS } from 'constants/ChannelConstants';
import { useTheme } from '@mui/material/styles';
import { API } from 'utils/api';
import { showError, showSuccess, trims, copy } from 'utils/common';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
  Divider,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  OutlinedInput,
  ButtonGroup,
  Container,
  Autocomplete,
  FormHelperText,
  Checkbox,
  Switch,
  FormControlLabel,
  Typography,
  Tooltip,
  Collapse,
  Box,
  Chip,
  useMediaQuery
} from '@mui/material';
import { Formik } from 'formik';
import * as Yup from 'yup';
import { defaultConfig, typeConfig } from '../type/Config'; //typeConfig
import { createFilterOptions } from '@mui/material/Autocomplete';
import CheckBoxOutlineBlankIcon from '@mui/icons-material/CheckBoxOutlineBlank';
import CheckBoxIcon from '@mui/icons-material/CheckBox';
import { useTranslation } from 'react-i18next';
import useCustomizeT from 'hooks/useCustomizeT';
import { PreCostType } from '../type/other';
import MapInput from './MapInput';
import ListInput from './ListInput';
import ModelSelectorModal from './ModelSelectorModal';
import pluginList from '../type/Plugin.json';
import { Icon } from '@iconify/react';

const icon = <CheckBoxOutlineBlankIcon fontSize="small" />;
const checkedIcon = <CheckBoxIcon fontSize="small" />;

const filter = createFilterOptions();
const getValidationSchema = (t) =>
  Yup.object().shape({
    is_edit: Yup.boolean(),
    // is_tag: Yup.boolean(),
    name: Yup.string().required(t('channel_edit.requiredName')),
    type: Yup.number().required(t('channel_edit.requiredChannel')),
    key: Yup.string().when('is_edit', { is: false, then: Yup.string().required(t('channel_edit.requiredKey')) }),
    other: Yup.string(),
    proxy: Yup.string(),
    test_model: Yup.string(),
    models: Yup.array().min(1, t('channel_edit.requiredModels')),
    groups: Yup.array().min(1, t('channel_edit.requiredGroup')),
    base_url: Yup.string().when(['type', 'openai_upstream'], {
      is: (type, upstream) => [3, 8].includes(type) || (type === 1 && (upstream === 'mountsea' || upstream === 'sutui')),
      then: Yup.string().required(t('channel_edit.requiredBaseUrl')), // base_url 必填：Azure/自定义/或 OpenAI+MountSea 上游
      otherwise: Yup.string()
    }),
    model_mapping: Yup.array(),
    model_headers: Yup.array(),
    custom_parameter: Yup.string().nullable(),
    openai_upstream: Yup.string()
  });

// MiniMax 上游供应商选项（后续可按渠道类型扩展）
const MINIMAX_UPSTREAM_OPTIONS = [
  { value: '', label: '默认（官方）' },
  { value: 'official', label: '官方（直连）' },
  { value: 'ppinfra', label: 'PPInfra（聚合）' },
  { value: 'polloi', label: 'Polloi（聚合）' }
];

// OpenAI 上游供应商选项
const OPENAI_UPSTREAM_OPTIONS = [
  { value: '', label: '默认（官方）' },
  { value: 'official', label: '官方（OpenAI）' },
  { value: 'openrouter', label: 'OpenRouter（聚合）' },
  { value: 'mountsea', label: 'MountSea' },
  { value: 'sutui', label: '速推 Sutui' },
  { value: 'apimart', label: 'Apimart（Sora 聚合）' }
];

// Gemini 上游供应商选项（简化版：只保留常用的供应商）
const GEMINI_UPSTREAM_OPTIONS = [
  { value: 'google', label: 'Google（官方）' },
  { value: 'sutui', label: 'Sutui（速推）' },
  { value: 'apimart', label: 'Apimart' },
  { value: 'ezlinkai', label: 'EzlinkAI' }
];

const EditModal = ({ open, channelId, onCancel, onOk, groupOptions, isTag, modelOptions, prices }) => {
  const { t } = useTranslation();
  const { t: customizeT } = useCustomizeT();
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  // const [loading, setLoading] = useState(false);
  const [initialInput, setInitialInput] = useState(defaultConfig.input);
  const [inputLabel, setInputLabel] = useState(defaultConfig.inputLabel); //
  const [inputPrompt, setInputPrompt] = useState(defaultConfig.prompt);
  const [batchAdd, setBatchAdd] = useState(false);
  const [hasTag, setHasTag] = useState(false);
  const [expanded, setExpanded] = useState(false);
  const [inputValue, setInputValue] = useState('');
  const [parameterFocused, setParameterFocused] = useState(false);
  const parameterInputRef = useRef(null);
  const removeDuplicates = (array) => [...new Set(array)];
  const [modelSelectorOpen, setModelSelectorOpen] = useState(false);
  const [tempFormikValues, setTempFormikValues] = useState(null);
  const [tempSetFieldValue, setTempSetFieldValue] = useState(null);

  // 供应商关键词（最小必要别名，便于搜索）
  const PROVIDER_KEYWORDS = useMemo(
    () => ({
      1: ['openai', 'gpt'],
      3: ['azure', 'azure openai', 'aoai'],
      55: ['azure v1', 'azure openai v1'],
      14: ['anthropic', 'claude'],
      15: ['baidu', '文心', '千帆'],
      16: ['zhipu', '智谱', 'glm'],
      17: ['ali', '阿里', '通义', 'qwen'],
      20: ['openrouter', 'router', 'or'],
      25: ['google', 'gemini', 'palm'],
      27: ['minimax', 'mini max', '海螺', 'hailuo', 'ppinfra', 'polloi'],
      31: ['groq'],
      36: ['cohere'],
      37: ['stability', 'sd3', 'stable'],
      39: ['ollama'],
      40: ['hunyuan', '混元', 'tencent'],
      41: ['suno'],
      42: ['vertex', 'vertexai', 'google'],
      45: ['siliconflow', '硅基'],
      47: ['jina'],
      49: ['github', 'gh models'],
      51: ['recraft'],
      52: ['replicate'],
      53: ['kling'],
      56: ['xai', 'grok'],
      57: ['vidu'],
      58: ['volc', 'ark', '火山', '方舟', 'doubao'],
      59: ['huawei', 'maas', 'modelarts', '华为']
    }),
    []
  );

  const providerOptions = useMemo(() =>
    Object.values(CHANNEL_OPTIONS).map((opt) => ({
      label: opt.text,
      value: opt.value,
      keywords: PROVIDER_KEYWORDS[opt.value] || []
    })),
  [PROVIDER_KEYWORDS]);

  const initChannel = (typeValue) => {
    if (typeConfig[typeValue]?.inputLabel) {
      setInputLabel({ ...defaultConfig.inputLabel, ...typeConfig[typeValue].inputLabel });
    } else {
      setInputLabel(defaultConfig.inputLabel);
    }

    if (typeConfig[typeValue]?.prompt) {
      setInputPrompt({ ...defaultConfig.prompt, ...typeConfig[typeValue].prompt });
    } else {
      setInputPrompt(defaultConfig.prompt);
    }

    return typeConfig[typeValue]?.input;
  };

  const handleTypeChange = (setFieldValue, typeValue, values) => {
    // 处理插件事务
    if (pluginList[typeValue]) {
      const newPluginValues = {};
      const pluginConfig = pluginList[typeValue];
      for (const pluginName in pluginConfig) {
        const plugin = pluginConfig[pluginName];
        const oldValve = values['plugin'] ? values['plugin'][pluginName] || {} : {};
        newPluginValues[pluginName] = {};
        for (const paramName in plugin.params) {
          const param = plugin.params[paramName];
          newPluginValues[pluginName][paramName] = oldValve[paramName] || (param.type === 'bool' ? false : '');
        }
      }
      setFieldValue('plugin', newPluginValues);
    }

    const newInput = initChannel(typeValue);

    if (newInput) {
      Object.keys(newInput).forEach((key) => {
        if (
          (!Array.isArray(values[key]) && values[key] !== null && values[key] !== undefined && values[key] !== '') ||
          (Array.isArray(values[key]) && values[key].length > 0)
        ) {
          return;
        }

        if (key === 'models') {
          setFieldValue(key, initialModel(newInput[key]));
          return;
        }
        setFieldValue(key, newInput[key]);
      });
    }

    if (typeValue === 57) {
      const hasModels = Array.isArray(values.models) ? values.models.length > 0 : !!values.models;
      if (!hasModels) {
        const defaultModels = basicModels(typeValue);
        if (defaultModels.length > 0) {
          setFieldValue('models', defaultModels);
        }
      }
    }
  };

  const basicModels = (channelType) => {
    // Vidu：仅返回基础模型，避免自动填充变种模型
    if (channelType === 57) {
      const group = typeConfig[channelType]?.modelGroup || 'Vidu';
      const bases = ['viduq2-pro', 'viduq2-turbo', 'viduq1', 'viduq1-classic', 'vidu2.0', 'vidu1.5'];
      return bases.map((id) => ({ id, group }));
    }
    // Gemini：默认填充 Veo 官方视频模型
    if (channelType === 25) {
      const group = typeConfig[channelType]?.modelGroup || 'Google Gemini';
      const bases = ['veo-3.1-generate-preview', 'veo-3.1-fast-generate-preview'];
      return bases.map((id) => ({ id, group }));
    }

    let modelGroup = typeConfig[channelType]?.modelGroup || defaultConfig.modelGroup;
    // 循环 modelOptions，找到 modelGroup 对应的模型
    let modelList = [];
    modelOptions.forEach((model) => {
      if (model.group === modelGroup) {
        modelList.push(model);
      }
    });
    return modelList;
  };

  const handleModelSelectorConfirm = (selectedModels, overwriteModels) => {
    if (tempSetFieldValue && tempFormikValues) {
      if (overwriteModels) {
        // 覆盖模式：清空现有的模型列表，使用选择器中的模型
        tempSetFieldValue('models', selectedModels);
      } else {
        // 追加模式：合并现有模型和新选择的模型，避免重复
        const existingModels = tempFormikValues.models || [];
        const existingModelIds = new Set(existingModels.map((model) => model.id));

        // 过滤掉已存在的模型，避免重复
        const newModels = selectedModels.filter((model) => !existingModelIds.has(model.id));

        // 合并模型列表
        tempSetFieldValue('models', [...existingModels, ...newModels]);
      }
    }
  };

  const submit = async (values, { setErrors, setStatus, setSubmitting }) => {
    setSubmitting(true);
    values = trims(values);
    if (values.base_url && values.base_url.endsWith('/')) {
      values.base_url = values.base_url.slice(0, values.base_url.length - 1);
    }
    if (values.type === 3 && values.other === '') {
      values.other = '2024-05-01-preview';
    }
    if (values.type === 18 && values.other === '') {
      values.other = 'v2.1';
    }
    let res;

    let modelMappingModel = [];

    if (values.model_mapping) {
      try {
        const modelMapping = values.model_mapping.reduce((acc, item) => {
          if (item.key && item.value) {
            acc[item.key] = item.value;
          }
          return acc;
        }, {});
        const cleanedMapping = {};

        for (const [key, value] of Object.entries(modelMapping)) {
          if (key && value && !(key in cleanedMapping)) {
            cleanedMapping[key] = value;
            modelMappingModel.push(key);
          }
        }

        values.model_mapping = JSON.stringify(cleanedMapping, null, 2);
      } catch (error) {
        showError('Error parsing model_mapping:' + error.message);
      }
    }
    let modelHeadersKey = [];

    if (values.model_headers) {
      try {
        const modelHeader = values.model_headers.reduce((acc, item) => {
          if (item.key && item.value) {
            acc[item.key] = item.value;
          }
          return acc;
        }, {});
        const cleanedHeader = {};

        for (const [key, value] of Object.entries(modelHeader)) {
          if (key && value && !(key in cleanedHeader)) {
            cleanedHeader[key] = value;
            modelHeadersKey.push(key);
          }
        }

        values.model_headers = JSON.stringify(cleanedHeader, null, 2);
      } catch (error) {
        showError('Error parsing model_headers:' + error.message);
      }
    }

    // 合并“上游供应商”到 custom_parameter 顶层（仅当有选择时写入；为空则不写入/移除）
    if (values.custom_parameter) {
      try {
        JSON.parse(values.custom_parameter);
      } catch (error) {
        showError('Error parsing custom_parameter: ' + error.message);
        setStatus({ success: false });
        setErrors({ submit: error.message });
        setSubmitting(false);
        return;
      }
    }

    if (values.disabled_stream) {
      values.disabled_stream = removeDuplicates(values.disabled_stream);
    }

    // 获取现有的模型 ID
    const existingModelIds = values.models.map((model) => model.id);

    // 找出在 modelMappingModel 中存在但不在 existingModelIds 中的模型
    const newModelIds = modelMappingModel.filter((id) => !existingModelIds.includes(id));

    // 合并现有的模型 ID 和新的模型 ID，并去重
    const allUniqueModelIds = Array.from(new Set([...existingModelIds, ...newModelIds]));

    // 创建新的 modelsStr
    const modelsStr = allUniqueModelIds.join(',');
    values.group = values.groups.join(',');

    let baseApiUrl = '/api/channel/';

    if (isTag) {
      baseApiUrl = '/api/channel_tag/' + encodeURIComponent(channelId);
    }

    // 删除前端临时字段，避免传给后端

    try {
      if (channelId) {
        res = await API.put(baseApiUrl, { ...values, id: parseInt(channelId), models: modelsStr });
      } else {
        res = await API.post(baseApiUrl, { ...values, models: modelsStr });
      }
      const { success, message } = res.data;
      if (success) {
        if (channelId) {
          showSuccess(t('channel_edit.editSuccess'));
        } else {
          showSuccess(t('channel_edit.addSuccess'));
        }
        setSubmitting(false);
        setStatus({ success: true });
        onOk(true);
        return;
      } else {
        setStatus({ success: false });
        showError(message);
        setErrors({ submit: message });
      }
    } catch (error) {
      setStatus({ success: false });
      showError(error.message);
      setErrors({ submit: error.message });
      return;
    }
  };

  function initialModel(channelModel) {
    if (!channelModel) {
      return [];
    }

    // 如果 channelModel 是一个字符串
    if (typeof channelModel === 'string') {
      channelModel = channelModel.split(',');
    }
    let modelList = channelModel.map((model) => {
      const modelOption = modelOptions.find((option) => option.id === model);
      if (modelOption) {
        return modelOption;
      }
      return { id: model, group: t('channel_edit.customModelTip') };
    });
    return modelList;
  }

  const loadChannel = async () => {
    try {
      let baseApiUrl = `/api/channel/${channelId}`;

      if (isTag) {
        baseApiUrl = '/api/channel_tag/' + encodeURIComponent(channelId);
      }

      let res = await API.get(baseApiUrl);
      const { success, message, data } = res.data;
      if (success) {
        if (data.models === '') {
          data.models = [];
        } else {
          data.models = initialModel(data.models);
        }
        if (data.group === '') {
          data.groups = [];
        } else {
          data.groups = data.group.split(',');
        }

        data.model_mapping =
          data.model_mapping !== ''
            ? Object.entries(JSON.parse(data.model_mapping)).map(([key, value], index) => ({
                index,
                key,
                value
              }))
            : [];
        // if (data.model_headers) {
        data.model_headers =
          data.model_headers !== ''
            ? Object.entries(JSON.parse(data.model_headers)).map(([key, value], index) => ({
                index,
                key,
                value
              }))
            : [];
        // }

        // Format the custom_parameter JSON for better readability if it's not empty
        if (data.custom_parameter !== '') {
          try {
            // Parse and then stringify with indentation for formatting
            const parsedJson = JSON.parse(data.custom_parameter);
            data.custom_parameter = JSON.stringify(parsedJson, null, 2);
          } catch (error) {
            // If parsing fails, keep the original string
            console.log('Error parsing custom_parameter JSON:', error);
          }
        } else {
          data.custom_parameter = '';
        }

        data.base_url = data.base_url ?? '';
        data.is_edit = true;
        if (data.plugin === null) {
          data.plugin = {};
        }
        // 解析 custom_parameter 顶层 upstream 作为 UI 展示的上游选项（minimax/openai）。Gemini 渠道不使用第三方聚合上游。
        if ((data.type === 27 || data.type === 1) && data.custom_parameter) {
          try {
            const obj = JSON.parse(data.custom_parameter);
            if (obj && typeof obj === 'object') {
              if (data.type === 27) {
                if (obj.video && typeof obj.video === 'object' && obj.video.upstream) {
                  data.minimax_upstream = obj.video.upstream;
                } else if (obj.upstream) {
                  data.minimax_upstream = obj.upstream;
                }
              } else if (data.type === 1) {
                if (obj.upstream) {
                  data.openai_upstream = obj.upstream;
                }
              }
            }
          } catch (e) {
            // ignore parse error, 交给原有校验
          }
        }
        if (data.type === 27 && !data.minimax_upstream) data.minimax_upstream = 'official';
        if (data.type === 1 && !data.openai_upstream) data.openai_upstream = 'official';

        // Gemini：根据已保存的 plugin.gemini_video.vendor 或 base_url 推断 gemini_upstream，便于回显
        if (data.type === 25) {
          let vendor = 'google';
          try {
            if (data.plugin && typeof data.plugin === 'object') {
              const gv = data.plugin.gemini_video && data.plugin.gemini_video.vendor;
              const gv2 = data.plugin.gemini && data.plugin.gemini.video && data.plugin.gemini.video.vendor;
              if (typeof gv === 'string' && gv.trim() !== '') vendor = gv.trim();
              else if (typeof gv2 === 'string' && gv2.trim() !== '') vendor = gv2.trim();
            }
          } catch (_) {}
          if (!vendor || vendor === '') {
            const base = (data.base_url || '').toLowerCase();
            if (base.includes('sora2.pub') || base.includes('sutui') || base.includes('st-ai')) vendor = 'sutui';
            else if (base.includes('apimart')) vendor = 'apimart';
            else if (base.includes('ezlinkai')) vendor = 'ezlinkai';
            else vendor = 'google';
          }
          data.gemini_upstream = vendor;
        }

        initChannel(data.type);
        setInitialInput(data);

        if (!isTag && data.tag) {
          setHasTag(true);
        }
      } else {
        showError(message);
      }
    } catch (error) {
      return;
    }
  };

  useEffect(() => {
    if (open) {
      setBatchAdd(isTag);
      if (channelId) {
        loadChannel().then();
      } else {
        setHasTag(false);
        // 新建时默认选择 OpenAI；用户切换到 Gemini 时将自动填充 apimart 与 Veo 模型
        initChannel(1);
        setInitialInput({ ...defaultConfig.input, is_edit: false });
      }
    }

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [channelId, open]);

  return (
    <Dialog open={open} onClose={onCancel} fullWidth maxWidth={'md'}>
      <DialogTitle sx={{ margin: '0px', fontWeight: 700, lineHeight: '1.55556', padding: '24px', fontSize: '1.125rem' }}>
        {channelId ? t('common.edit') : t('common.create')}
      </DialogTitle>
      <Divider />
      <DialogContent>
        <Formik initialValues={initialInput} enableReinitialize validationSchema={getValidationSchema(t)} onSubmit={submit}>
          {({ errors, handleBlur, handleChange, handleSubmit, isSubmitting, touched, values, setFieldValue }) => {
            // 保存当前Formik状态，以便在模型选择器中使用
            const openModelSelector = () => {
              setTempFormikValues({ ...values });
              setTempSetFieldValue(() => setFieldValue); // 保存函数引用
              setModelSelectorOpen(true);
            };

            return (
              <form noValidate onSubmit={handleSubmit}>
                {!isTag && (
                  <FormControl fullWidth error={Boolean(touched.type && errors.type)} sx={{ ...theme.typography.otherInput }}>
                    <Autocomplete
                      id="channel-type-autocomplete"
                      options={providerOptions}
                      value={providerOptions.find((o) => o.value === values.type) || null}
                      getOptionLabel={(o) => (o?.label ? o.label : '')}
                      isOptionEqualToValue={(o, v) => o.value === v.value}
                      onChange={(_, option) => {
                        if (!option) return;
                        setFieldValue('type', option.value);
                        handleTypeChange(setFieldValue, option.value, values);
                        // 当选择 Gemini 渠道时：默认 vendor=google，默认 BaseURL 与模型
                        if (option.value === 25) {
                          // plugin 结构可能尚未初始化，做保护性设置
                          const plugin = values.plugin ? { ...values.plugin } : {};
                          plugin.gemini_video = plugin.gemini_video || {};
                          if (!plugin.gemini_video.vendor) plugin.gemini_video.vendor = 'google';
                          setFieldValue('plugin', plugin);
                          // 同步基础区上游下拉
                          if (!values.gemini_upstream) setFieldValue('gemini_upstream', 'google');
                          if (!values.base_url) setFieldValue('base_url', '');
                          // 若未选择模型，自动填充 Veo 官方模型
                          const hasModels = Array.isArray(values.models) ? values.models.length > 0 : !!values.models;
                          if (!hasModels) {
                            const defaults = basicModels(25);
                            if (defaults.length > 0) setFieldValue('models', defaults);
                            setFieldValue('test_model', 'veo-3.1-generate-preview');
                          }
                        }
                      }}
                      filterOptions={(opts, state) => {
                        const q = (state.inputValue || '').trim().toLowerCase();
                        if (!q) return opts;
                        return opts.filter((o) => {
                          const hay = (o.label + ' ' + (o.keywords || []).join(' ')).toLowerCase();
                          return hay.includes(q);
                        });
                      }}
                      renderInput={(params) => (
                        <TextField
                          {...params}
                          label={customizeT(inputLabel.type)}
                          error={Boolean(touched.type && errors.type)}
                          helperText={touched.type && errors.type ? errors.type : customizeT(inputPrompt.type)}
                          onBlur={handleBlur}
                        />
                      )}
                      autoHighlight
                      disableClearable
                      openOnFocus
                      disabled={hasTag}
                    />
                  </FormControl>
                )}

                {/* Gemini 上游供应商（简化版：只需一个下拉框） */}
                {!isTag && values.type === 25 && (
                  <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
                    <InputLabel htmlFor="gemini-upstream-label">上游供应商</InputLabel>
                    <Select
                      id="gemini-upstream-label"
                      name="gemini_upstream"
                      label="上游供应商"
                      value={values.gemini_upstream || 'google'}
                      onChange={(e) => {
                        const next = e.target.value;
                        setFieldValue('gemini_upstream', next);
                        
                        // 同步到 plugin 配置
                        const plugin = values.plugin ? { ...values.plugin } : {};
                        plugin.gemini_video = plugin.gemini_video || {};
                        plugin.gemini_video.vendor = next;
                        setFieldValue('plugin', plugin);
                        
                        // 根据供应商自动设置 base_url
                        if (next === 'google') {
                          setFieldValue('base_url', '');
                        } else if (next === 'sutui') {
                          setFieldValue('base_url', 'https://api.sora2.pub');
                        } else if (next === 'apimart') {
                          setFieldValue('base_url', 'https://api.apimart.ai');
                        } else if (next === 'ezlinkai') {
                          setFieldValue('base_url', 'https://api.ezlinkai.com');
                        }
                      }}
                      MenuProps={{ PaperProps: { style: { maxHeight: 200 } } }}
                    >
                      {GEMINI_UPSTREAM_OPTIONS.map((opt) => (
                        <MenuItem key={opt.value} value={opt.value}>
                          {opt.label}
                        </MenuItem>
                      ))}
                    </Select>
                    <FormHelperText>选择 Gemini 视频生成的上游供应商</FormHelperText>
                  </FormControl>
                )}

                {/* MiniMax 专用：上游供应商选择（顶层 upstream） */}
                {!isTag && values.type === 27 && (
                  <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
                    <InputLabel id="minimax-upstream-label">上游供应商</InputLabel>
                    <Select
                      labelId="minimax-upstream-label"
                      id="minimax-upstream"
                      name="minimax_upstream"
                      value={values.minimax_upstream || ''}
                      label="上游供应商"
                      onChange={(e) => {
                        const next = e.target.value;
                        setFieldValue('minimax_upstream', next);
                        if (next === 'ppinfra') {
                          setFieldValue('base_url', 'https://api.ppinfra.com');
                        } else if (next === 'polloi') {
                          setFieldValue('base_url', 'https://pollo.ai/api/platform');
                        } else if (next === 'official' || next === '') {
                          setFieldValue('base_url', 'https://api.minimaxi.com');
                        }
                        // 同步 custom_parameter 顶层 upstream，方便后端/视频/语音客户端解析
                        try {
                          const obj = values.custom_parameter ? JSON.parse(values.custom_parameter) : {};
                          if (next && next !== 'official') {
                            obj.upstream = next;
                          } else {
                            if (obj && typeof obj === 'object' && 'upstream' in obj) delete obj.upstream;
                          }
                          // 音频能力：为 PPInfra 写入 audio.upstream=ppinfra（后端也支持顶层 upstream；此处仅增强可读性与显式性）
                          if (!obj.audio || typeof obj.audio !== 'object') obj.audio = {};
                          if (next === 'ppinfra') {
                            obj.audio.upstream = 'ppinfra';
                          } else if (obj.audio && 'upstream' in obj.audio) {
                            delete obj.audio.upstream;
                          }
                          // 简化视频配置：选中聚合上游时，自动注入 video 段的合理默认值（用户可继续编辑）
                          if (!obj.video || typeof obj.video !== 'object') obj.video = {};
                          if (next === 'ppinfra') {
                            obj.video.base_url = obj.video.base_url || 'https://api.ppinfra.com';
                            obj.video.submit_path_template = obj.video.submit_path_template || '/v3/async/%s';
                            obj.video.query_path_template = obj.video.query_path_template || '/v3/async/task-result';
                            obj.video.auth_header = obj.video.auth_header || 'Authorization';
                            obj.video.auth_scheme = obj.video.auth_scheme || 'Bearer';
                          }
                          if (next === 'polloi') {
                            obj.video.base_url = obj.video.base_url || 'https://pollo.ai/api/platform';
                            obj.video.submit_path_template = obj.video.submit_path_template || '/generation/minimax/%s';
                            obj.video.query_path_template = obj.video.query_path_template || '/generation/%s/status';
                            obj.video.auth_header = obj.video.auth_header || 'x-api-key';
                            obj.video.auth_scheme = obj.video.auth_scheme || 'none';
                          }
                          setFieldValue('custom_parameter', JSON.stringify(obj, null, 2));
                        } catch (_) {
                          // 保守处理：不破坏用户自定义 JSON
                        }
                      }}
                    >
                      {MINIMAX_UPSTREAM_OPTIONS.map((opt) => (
                        <MenuItem key={opt.value} value={opt.value}>
                          {opt.label}
                        </MenuItem>
                      ))}
                    </Select>
                    <FormHelperText id="helper-tex-minimax-upstream-label">
                      仅影响 MiniMax 视频/文本/语音的上游请求入口，默认官方直连；高级按能力覆盖可在下方“额外参数”JSON中配置。
                    </FormHelperText>
                  </FormControl>
                )}

                {/* Vidu 专用：上游供应商选择（官方 / Pollo.ai） */}
                {!isTag && values.type === 57 && (
                  <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
                    <InputLabel id="vidu-upstream-label">上游供应商</InputLabel>
                    <Select
                      labelId="vidu-upstream-label"
                      id="vidu-upstream"
                      name="vidu_upstream"
                      value={values.vidu_upstream || ''}
                      label="上游供应商"
                      onChange={(e) => {
                        const next = e.target.value;
                        setFieldValue('vidu_upstream', next);

                        const group = typeConfig[57]?.modelGroup || 'Vidu';
                        const officialModels = ['viduq2-pro','viduq2-turbo','viduq1','viduq1-classic','vidu2.0','vidu1.5'].map(id => ({ id, group }));
                        const polloModels = ['viduq2-pro','viduq2-turbo','viduq1'].map(id => ({ id, group }));

                        if (next === 'pollo') {
                          // 切换到 Pollo 上游
                          setFieldValue('base_url', 'https://pollo.ai/api/platform');
                          // 写入 custom_parameter 顶层 upstream=pollo 便于后端识别
                          try {
                            const obj = values.custom_parameter ? JSON.parse(values.custom_parameter) : {};
                            obj.upstream = 'pollo';
                            setFieldValue('custom_parameter', JSON.stringify(obj, null, 2));
                          } catch (_) {}
                          // 自动填充 Pollo 支持的基础模型
                          setFieldValue('models', polloModels);
                        } else {
                          // 官方直连
                          setFieldValue('base_url', 'https://api.vidu.cn');
                          try {
                            const obj = values.custom_parameter ? JSON.parse(values.custom_parameter) : {};
                            if (obj && typeof obj === 'object' && 'upstream' in obj) delete obj.upstream;
                            setFieldValue('custom_parameter', JSON.stringify(obj, null, 2));
                          } catch (_) {}
                          // 自动填充官方基础模型
                          setFieldValue('models', officialModels);
                        }
                      }}
                    >
                      <MenuItem value="">默认（官方）</MenuItem>
                      <MenuItem value="official">官方（直连）</MenuItem>
                      <MenuItem value="pollo">Pollo.ai</MenuItem>
                    </Select>
                    <FormHelperText id="helper-tex-vidu-upstream-label">
                      默认直连官方；选择 Pollo.ai 将自动切换认证与路径，模型仅需基础名（系统会自动映射计费组合）。
                    </FormHelperText>
                  </FormControl>
                )}

                {/* OpenAI 专用：上游供应商选择（顶层 upstream） */}
                {!isTag && values.type === 1 && (
                  <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
                    <InputLabel id="openai-upstream-label">上游供应商</InputLabel>
                    <Select
                      labelId="openai-upstream-label"
                      id="openai-upstream"
                      name="openai_upstream"
                      value={values.openai_upstream || ''}
                      label="上游供应商"
                      onChange={(e) => {
                        const next = e.target.value;
                        setFieldValue('openai_upstream', next);
                        if (next === 'openrouter') {
                          setFieldValue('base_url', 'https://openrouter.ai/api');
                        } else if (next === 'mountsea') {
                          // MountSea 需使用其 API 域名，默认给出可识别占位，建议按实际端点覆盖
                          setFieldValue('base_url', 'https://api.mountsea.ai');
                        } else if (next === 'sutui') {
                          // Sutui 速推：使用官方提供的域名
                          setFieldValue('base_url', 'https://api.sora2.pub');
                        } else if (next === 'apimart') {
                          setFieldValue('base_url', 'https://api.apimart.ai');
                        } else if (next === 'official' || next === '') {
                          setFieldValue('base_url', 'https://api.openai.com');
                        }
                        // 将选择同步到 custom_parameter 顶层 upstream（OpenAIProvider 将按此切换基础域名/视频上游识别）
                        try {
                          const obj = values.custom_parameter ? JSON.parse(values.custom_parameter) : {};
                          if (next && next !== 'official') {
                            obj.upstream = next;
                          } else {
                            if (obj && typeof obj === 'object' && 'upstream' in obj) delete obj.upstream;
                          }
                          setFieldValue('custom_parameter', JSON.stringify(obj, null, 2));
                        } catch (_) {
                          // 不打断提交流程，交由用户自行修正 JSON
                        }

                        // Sora 视频上游：当选择 Sutui 时，显式写入插件字段以便后端识别；其他选项则清空
                        if (next === 'sutui') {
                          setFieldValue('plugin.sora.vendor', 'sutui');
                          // 对外仅暴露 OpenAI 官方模型：sora-2 / sora-2-pro
                          try {
                            const openaiVideoModels = ['sora-2', 'sora-2-pro'];
                            const sutuiVeoModels = ['veo3', 'veo3.1', 'veo3-pro', 'veo3.1-pro', 'veo3.1-components'];
                            const current = Array.isArray(values.models) ? values.models : [];
                            if (current.length === 0) {
                              const prefill = [...openaiVideoModels, ...sutuiVeoModels].map((id) => ({ id, group: 'OpenAI Video' }));
                              setFieldValue('models', prefill);
                              if (!values.test_model) setFieldValue('test_model', 'sora-2');
                            }
                          } catch (_) {}
                        } else if (next === 'apimart') {
                          setFieldValue('plugin.sora.vendor', 'apimart');
                          try {
                            const apimartVideoModels = ['sora-2', 'sora-2-pro'];
                            const current = Array.isArray(values.models) ? values.models : [];
                            if (current.length === 0) {
                              const prefill = apimartVideoModels.map((id) => ({ id, group: 'OpenAI Video' }));
                              setFieldValue('models', prefill);
                              if (!values.test_model) setFieldValue('test_model', 'sora-2');
                            }
                          } catch (_) {}
                        } else {
                          // 清理避免误判
                          setFieldValue('plugin.sora.vendor', '');
                        }
                      }}
                    >
                      {OPENAI_UPSTREAM_OPTIONS.map((opt) => (
                        <MenuItem key={opt.value} value={opt.value}>
                          {opt.label}
                        </MenuItem>
                      ))}
                    </Select>
                    <FormHelperText id="helper-tex-openai-upstream-label">
                      可选官方 / OpenRouter / MountSea / Sutui / Apimart。选择第三方后将自动切换基础地址与 custom_parameter.upstream；如使用 Sutui/Apimart，请确认 base_url 为实际端点。
                    </FormHelperText>
                  </FormControl>
                )}

                <FormControl fullWidth error={Boolean(touched.tag && errors.tag)} sx={{ ...theme.typography.otherInput }}>
                  <InputLabel htmlFor="channel-tag-label">{customizeT(inputLabel.tag)}</InputLabel>
                  <OutlinedInput
                    id="channel-tag-label"
                    label={customizeT(inputLabel.tag)}
                    type="text"
                    value={values.tag}
                    name="tag"
                    onBlur={handleBlur}
                    onChange={handleChange}
                    inputProps={{}}
                    aria-describedby="helper-text-channel-tag-label"
                  />
                  {touched.tag && errors.tag ? (
                    <FormHelperText error id="helper-tex-channel-tag-label">
                      {errors.tag}
                    </FormHelperText>
                  ) : (
                    <FormHelperText id="helper-tex-channel-tag-label"> {customizeT(inputPrompt.tag)} </FormHelperText>
                  )}
                </FormControl>

                {!isTag && (
                  <FormControl fullWidth error={Boolean(touched.name && errors.name)} sx={{ ...theme.typography.otherInput }}>
                    <InputLabel htmlFor="channel-name-label">{customizeT(inputLabel.name)}</InputLabel>
                    <OutlinedInput
                      id="channel-name-label"
                      label={customizeT(inputLabel.name)}
                      type="text"
                      value={values.name}
                      name="name"
                      onBlur={handleBlur}
                      onChange={handleChange}
                      inputProps={{ autoComplete: 'name' }}
                      aria-describedby="helper-text-channel-name-label"
                    />
                    {touched.name && errors.name ? (
                      <FormHelperText error id="helper-tex-channel-name-label">
                        {errors.name}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-name-label"> {customizeT(inputPrompt.name)} </FormHelperText>
                    )}
                  </FormControl>
                )}
                {inputPrompt.base_url && (
                  <FormControl fullWidth error={Boolean(touched.base_url && errors.base_url)} sx={{ ...theme.typography.otherInput }}>
                    <InputLabel htmlFor="channel-base_url-label">{customizeT(inputLabel.base_url)}</InputLabel>
                    <OutlinedInput
                      id="channel-base_url-label"
                      label={customizeT(inputLabel.base_url)}
                      type="text"
                      value={values.base_url}
                      name="base_url"
                      onBlur={handleBlur}
                      onChange={handleChange}
                      inputProps={{}}
                      aria-describedby="helper-text-channel-base_url-label"
                    />

                    {touched.base_url && errors.base_url ? (
                      <FormHelperText error id="helper-tex-channel-base_url-label">
                        {errors.base_url}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-base_url-label"> {customizeT(inputPrompt.base_url)} </FormHelperText>
                    )}
                  </FormControl>
                )}

                {inputPrompt.other && (
                  <FormControl fullWidth error={Boolean(touched.other && errors.other)} sx={{ ...theme.typography.otherInput }}>
                    <InputLabel htmlFor="channel-other-label">{customizeT(inputLabel.other)}</InputLabel>
                    <OutlinedInput
                      id="channel-other-label"
                      label={customizeT(inputLabel.other)}
                      type="text"
                      value={values.other}
                      name="other"
                      disabled={hasTag}
                      onBlur={handleBlur}
                      onChange={handleChange}
                      inputProps={{}}
                      aria-describedby="helper-text-channel-other-label"
                    />
                    {touched.other && errors.other ? (
                      <FormHelperText error id="helper-tex-channel-other-label">
                        {errors.other}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-other-label"> {customizeT(inputPrompt.other)} </FormHelperText>
                    )}
                  </FormControl>
                )}

                <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
                  <Autocomplete
                    multiple
                    id="channel-groups-label"
                    options={groupOptions}
                    value={values.groups}
                    disabled={hasTag}
                    onChange={(e, value) => {
                      const event = {
                        target: {
                          name: 'groups',
                          value: value
                        }
                      };
                      handleChange(event);
                    }}
                    onBlur={handleBlur}
                    filterSelectedOptions
                    renderInput={(params) => (
                      <TextField {...params} name="groups" error={Boolean(errors.groups)} label={customizeT(inputLabel.groups)} />
                    )}
                    aria-describedby="helper-text-channel-groups-label"
                  />
                  {errors.groups ? (
                    <FormHelperText error id="helper-tex-channel-groups-label">
                      {errors.groups}
                    </FormHelperText>
                  ) : (
                    <FormHelperText id="helper-tex-channel-groups-label"> {customizeT(inputPrompt.groups)} </FormHelperText>
                  )}
                </FormControl>

                <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
                  <Box sx={{ position: 'relative' }}>
                    <Autocomplete
                      multiple
                      freeSolo
                      disableCloseOnSelect
                      id="channel-models-label"
                      disabled={hasTag}
                      options={modelOptions}
                      value={values.models}
                      inputValue={inputValue}
                      onInputChange={(event, newInputValue) => {
                        if (newInputValue.includes(',')) {
                          const modelsList = newInputValue
                            .split(',')
                            .map((item) => ({
                              id: item.trim(),
                              group: t('channel_edit.customModelTip')
                            }))
                            .filter((item) => item.id);

                          const updatedModels = [...new Set([...values.models, ...modelsList])];
                          const event = {
                            target: {
                              name: 'models',
                              value: updatedModels
                            }
                          };
                          handleChange(event);
                          setInputValue('');
                        } else {
                          setInputValue(newInputValue);
                        }
                      }}
                      onChange={(e, value) => {
                        const event = {
                          target: {
                            name: 'models',
                            value: value.map((item) =>
                              typeof item === 'string' ? { id: item, group: t('channel_edit.customModelTip') } : item
                            )
                          }
                        };
                        handleChange(event);
                      }}
                      renderInput={(params) => (
                        <TextField
                          {...params}
                          name="models"
                          error={Boolean(errors.models)}
                          label={customizeT(inputLabel.models)}
                          InputProps={{
                            ...params.InputProps
                          }}
                        />
                      )}
                      groupBy={(option) => option.group}
                      getOptionLabel={(option) => {
                        if (typeof option === 'string') {
                          return option;
                        }
                        if (option.inputValue) {
                          return option.inputValue;
                        }
                        return option.id;
                      }}
                      filterOptions={(options, params) => {
                        const filtered = filter(options, params);
                        const { inputValue } = params;
                        const isExisting = options.some((option) => inputValue === option.id);
                        if (inputValue !== '' && !isExisting) {
                          filtered.push({
                            id: inputValue,
                            group: t('channel_edit.customModelTip')
                          });
                        }
                        return filtered;
                      }}
                      renderOption={(props, option, { selected }) => (
                        <li {...props}>
                          <Checkbox icon={icon} checkedIcon={checkedIcon} style={{ marginRight: 8 }} checked={selected} />
                          {option.id}
                        </li>
                      )}
                      renderTags={(value, getTagProps) =>
                        value.map((option, index) => {
                          const tagProps = getTagProps({ index });
                          const { key, ...chipProps } = tagProps || {};
                          return (
                            <Chip
                              key={key || option.id || index}
                              label={option.id}
                              {...chipProps}
                              onClick={() => copy(option.id)}
                              sx={{
                                maxWidth: '100%',
                                height: 'auto',
                                margin: '3px',
                                '& .MuiChip-label': {
                                  whiteSpace: 'normal',
                                  wordBreak: 'break-word',
                                  padding: '6px 8px',
                                  lineHeight: 1.4,
                                  fontWeight: 400
                                },
                                '& .MuiChip-deleteIcon': {
                                  margin: '0 5px 0 -6px'
                                }
                              }}
                            />
                          );
                        })
                      }
                      sx={{
                        '& .MuiAutocomplete-tag': {
                          margin: '2px'
                        },
                        '& .MuiAutocomplete-inputRoot': {
                          flexWrap: 'wrap'
                        }
                      }}
                    />
                  </Box>
                  {errors.models ? (
                    <FormHelperText error id="helper-tex-channel-models-label">
                      {errors.models}
                    </FormHelperText>
                  ) : (
                    <FormHelperText id="helper-tex-channel-models-label"> {customizeT(inputPrompt.models)} </FormHelperText>
                  )}
                </FormControl>
                <Container
                  sx={{
                    textAlign: 'right'
                  }}
                >
                  <ButtonGroup variant="outlined" aria-label="small outlined primary button group">
                    <Button
                      size="small"
                      onClick={() => {
                        const modelString = values.models.map((model) => model.id).join(',');
                        copy(modelString);
                      }}
                    >
                      {isMobile ? <Icon icon="mdi:content-copy" /> : t('channel_edit.copyModels')}
                    </Button>
                    <Button
                      disabled={hasTag}
                      size="small"
                      onClick={() => {
                        setFieldValue('models', basicModels(values.type));
                      }}
                    >
                      {isMobile ? <Icon icon="mdi:playlist-plus" /> : t('channel_edit.inputChannelModel')}
                    </Button>
                    {/* <Button
                      disabled={hasTag}
                      size="small"
                      onClick={() => {
                        setFieldValue('models', modelOptions);
                      }}
                    >
                      {t('channel_edit.inputAllModel')}
                    </Button> */}
                    {inputLabel.provider_models_list && (
                      <Tooltip title={customizeT(inputPrompt.provider_models_list)} placement="top">
                        <Button
                          disabled={hasTag}
                          size="small"
                          onClick={openModelSelector}
                          startIcon={!isMobile && <Icon icon="mdi:cloud-download" />}
                        >
                          {isMobile ? <Icon icon="mdi:cloud-download" /> : customizeT(inputLabel.provider_models_list)}
                        </Button>
                      </Tooltip>
                    )}
                  </ButtonGroup>
                </Container>
                <FormControl fullWidth error={Boolean(touched.key && errors.key)} sx={{ ...theme.typography.otherInput }}>
                  {!batchAdd ? (
                    <>
                      <InputLabel htmlFor="channel-key-label">{customizeT(inputLabel.key)}</InputLabel>
                      <OutlinedInput
                        id="channel-key-label"
                        label={customizeT(inputLabel.key)}
                        type="text"
                        value={values.key}
                        name="key"
                        onBlur={handleBlur}
                        onChange={handleChange}
                        inputProps={{}}
                        aria-describedby="helper-text-channel-key-label"
                      />
                    </>
                  ) : (
                    <TextField
                      multiline
                      id="channel-key-label"
                      label={customizeT(inputLabel.key)}
                      value={values.key}
                      name="key"
                      onBlur={handleBlur}
                      onChange={handleChange}
                      aria-describedby="helper-text-channel-key-label"
                      minRows={5}
                      placeholder={customizeT(inputPrompt.key) + t('channel_edit.batchKeytip')}
                    />
                  )}

                  {touched.key && errors.key ? (
                    <FormHelperText error id="helper-tex-channel-key-label">
                      {errors.key}
                    </FormHelperText>
                  ) : (
                    <FormHelperText id="helper-tex-channel-key-label" component="div">
                      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                        <span>{customizeT(inputPrompt.key)}</span>
                        {channelId === 0 && (
                          <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                            <Switch 
                              size="small"
                              checked={Boolean(batchAdd)} 
                              onChange={(e) => setBatchAdd(e.target.checked)} 
                            />
                            <Typography variant="body2" component="span">{t('channel_edit.batchAdd')}</Typography>
                          </Box>
                        )}
                      </Box>
                    </FormHelperText>
                  )}
                </FormControl>

                {inputPrompt.model_mapping && (
                  <FormControl
                    fullWidth
                    error={Boolean(touched.model_mapping && errors.model_mapping)}
                    sx={{ ...theme.typography.otherInput }}
                  >
                    <MapInput
                      mapValue={values.model_mapping}
                      onChange={(newValue) => {
                        setFieldValue('model_mapping', newValue);
                      }}
                      disabled={hasTag}
                      error={Boolean(touched.model_mapping && errors.model_mapping)}
                      label={{
                        keyName: customizeT(inputLabel.model_mapping),
                        valueName: customizeT(inputPrompt.model_mapping),
                        name: customizeT(inputLabel.model_mapping)
                      }}
                    />
                    {touched.model_mapping && errors.model_mapping ? (
                      <FormHelperText error id="helper-tex-channel-model_mapping-label">
                        {errors.model_mapping}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-model_mapping-label">{customizeT(inputPrompt.model_mapping)}</FormHelperText>
                    )}
                  </FormControl>
                )}
                {inputPrompt.model_headers && (
                  <FormControl
                    fullWidth
                    error={Boolean(touched.model_headers && errors.model_headers)}
                    sx={{ ...theme.typography.otherInput }}
                  >
                    <MapInput
                      mapValue={values.model_headers}
                      onChange={(newValue) => {
                        setFieldValue('model_headers', newValue);
                      }}
                      disabled={hasTag}
                      error={Boolean(touched.model_headers && errors.model_headers)}
                      label={{
                        keyName: customizeT(inputLabel.model_headers),
                        valueName: customizeT(inputPrompt.model_headers),
                        name: customizeT(inputLabel.model_headers)
                      }}
                    />
                    {touched.model_headers && errors.model_headers ? (
                      <FormHelperText error id="helper-tex-channel-model_headers-label">
                        {errors.model_headers}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-model_headers-label">{customizeT(inputPrompt.model_headers)}</FormHelperText>
                    )}
                  </FormControl>
                )}
                {inputPrompt.custom_parameter && (
                  <FormControl
                    fullWidth
                    error={Boolean(touched.custom_parameter && errors.custom_parameter)}
                    sx={{ ...theme.typography.otherInput }}
                  >
                    <TextField
                      id="channel-custom_parameter-label"
                      label={customizeT(inputLabel.custom_parameter)}
                      multiline={Boolean(values.custom_parameter || parameterFocused)}
                      rows={values.custom_parameter || parameterFocused ? 8 : 1}
                      value={values.custom_parameter}
                      name="custom_parameter"
                      disabled={hasTag}
                      error={Boolean(touched.custom_parameter && errors.custom_parameter)}
                      onChange={handleChange}
                      inputRef={parameterInputRef}
                      onBlur={(e) => {
                        handleBlur(e);
                        setParameterFocused(false);
                      }}
                      onFocus={() => {
                        setParameterFocused(true);
                        // 使用setTimeout确保状态更新后重新聚焦
                        setTimeout(() => {
                          if (parameterInputRef.current) {
                            parameterInputRef.current.focus();
                          }
                        }, 0);
                      }}
                      placeholder={
                        parameterFocused
                          ? '{\n  "temperature": 0.7,\n  "top_p": 0.9,\n  "nested_param": {\n      "key": "value"\n  }\n}'
                          : ''
                      }
                    />
                    {touched.custom_parameter && errors.custom_parameter ? (
                      <FormHelperText error id="helper-tex-channel-custom_parameter-label">
                        {errors.custom_parameter}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-custom_parameter-label">
                        {customizeT(inputPrompt.custom_parameter)}
                      </FormHelperText>
                    )}
                  </FormControl>
                )}
                {inputPrompt.disabled_stream && (
                  <FormControl
                    fullWidth
                    error={Boolean(touched.disabled_stream && errors.disabled_stream)}
                    sx={{ ...theme.typography.otherInput }}
                  >
                    <ListInput
                      listValue={values.disabled_stream}
                      onChange={(newValue) => {
                        setFieldValue('disabled_stream', newValue);
                      }}
                      disabled={hasTag}
                      error={Boolean(touched.disabled_stream && errors.disabled_stream)}
                      label={{
                        name: customizeT(inputLabel.disabled_stream),
                        itemName: customizeT(inputPrompt.disabled_stream)
                      }}
                    />
                  </FormControl>
                )}

                <FormControl fullWidth error={Boolean(touched.proxy && errors.proxy)} sx={{ ...theme.typography.otherInput }}>
                  <InputLabel htmlFor="channel-proxy-label">{customizeT(inputLabel.proxy)}</InputLabel>
                  <OutlinedInput
                    id="channel-proxy-label"
                    label={customizeT(inputLabel.proxy)}
                    disabled={hasTag}
                    type="text"
                    value={values.proxy}
                    name="proxy"
                    onBlur={handleBlur}
                    onChange={handleChange}
                    inputProps={{}}
                    aria-describedby="helper-text-channel-proxy-label"
                  />
                  {touched.proxy && errors.proxy ? (
                    <FormHelperText error id="helper-tex-channel-proxy-label">
                      {errors.proxy}
                    </FormHelperText>
                  ) : (
                    <FormHelperText id="helper-tex-channel-proxy-label"> {customizeT(inputPrompt.proxy)} </FormHelperText>
                  )}
                </FormControl>
                {inputPrompt.test_model && (
                  <FormControl fullWidth error={Boolean(touched.test_model && errors.test_model)} sx={{ ...theme.typography.otherInput }}>
                    <InputLabel htmlFor="channel-test_model-label">{customizeT(inputLabel.test_model)}</InputLabel>
                    <OutlinedInput
                      id="channel-test_model-label"
                      label={customizeT(inputLabel.test_model)}
                      type="text"
                      disabled={hasTag}
                      value={values.test_model}
                      name="test_model"
                      onBlur={handleBlur}
                      onChange={handleChange}
                      inputProps={{}}
                      aria-describedby="helper-text-channel-test_model-label"
                    />
                    {touched.test_model && errors.test_model ? (
                      <FormHelperText error id="helper-tex-channel-test_model-label">
                        {errors.test_model}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-test_model-label"> {customizeT(inputPrompt.test_model)} </FormHelperText>
                    )}
                  </FormControl>
                )}
                {inputPrompt.only_chat && (
                  <FormControl fullWidth>
                    <FormControlLabel
                      control={
                        <Switch
                          disabled={hasTag}
                          checked={Boolean(values.only_chat)}
                          onChange={(event) => {
                            setFieldValue('only_chat', event.target.checked);
                          }}
                        />
                      }
                      label={customizeT(inputLabel.only_chat)}
                    />
                    <FormHelperText id="helper-tex-only_chat_model-label"> {customizeT(inputPrompt.only_chat)} </FormHelperText>
                  </FormControl>
                )}
                {inputPrompt.pre_cost && (
                  <FormControl fullWidth error={Boolean(touched.pre_cost && errors.pre_cost)} sx={{ ...theme.typography.otherInput }}>
                    <InputLabel htmlFor="channel-pre_cost-label">{customizeT(inputLabel.pre_cost)}</InputLabel>
                    <Select
                      id="channel-pre_cost-label"
                      label={customizeT(inputLabel.pre_cost)}
                      value={values.pre_cost}
                      name="pre_cost"
                      onBlur={handleBlur}
                      onChange={handleChange}
                      disabled={hasTag}
                      MenuProps={{
                        PaperProps: {
                          style: {
                            maxHeight: 200
                          }
                        }
                      }}
                    >
                      {PreCostType.map((option) => {
                        return (
                          <MenuItem key={option.value} value={option.value}>
                            {option.label}
                          </MenuItem>
                        );
                      })}
                    </Select>
                    {touched.pre_cost && errors.pre_cost ? (
                      <FormHelperText error id="helper-tex-channel-pre_cost-label">
                        {errors.pre_cost}
                      </FormHelperText>
                    ) : (
                      <FormHelperText id="helper-tex-channel-pre_cost-label"> {customizeT(inputPrompt.pre_cost)} </FormHelperText>
                    )}
                  </FormControl>
                )}
                {inputPrompt.compatible_response && (
                  <FormControl fullWidth>
                    <FormControlLabel
                      control={
                        <Switch
                          disabled={hasTag}
                          checked={Boolean(values.compatible_response)}
                          onChange={(event) => {
                            setFieldValue('compatible_response', event.target.checked);
                          }}
                        />
                      }
                      label={customizeT(inputLabel.compatible_response)}
                    />
                    <FormHelperText id="helper-tex-compatible_response-label">{customizeT(inputPrompt.compatible_response)}</FormHelperText>
                  </FormControl>
                )}

                {pluginList[values.type] &&
                  Object.keys(pluginList[values.type]).map((pluginId) => {
                    const plugin = pluginList[values.type][pluginId];
                    return (
                        <Box key={pluginId}
                          sx={{
                            border: '1px solid #e0e0e0',
                            borderRadius: 2,
                            marginTop: 2,
                            marginBottom: 2,
                            overflow: 'hidden'
                          }}
                        >
                          <Box
                            sx={{
                              display: 'flex',
                              justifyContent: 'space-between',
                              alignItems: 'center',
                              padding: 2
                            }}
                          >
                            <Box sx={{ flex: 1 }}>
                              <Typography variant="h3">{customizeT(plugin.name)}</Typography>
                              <Typography variant="caption" component="span">{customizeT(plugin.description)}</Typography>
                            </Box>
                            <Button
                              onClick={() => setExpanded(!expanded)}
                              endIcon={
                                expanded ? (
                                  <Icon icon="solar:alt-arrow-up-line-duotone" />
                                ) : (
                                  <Icon icon="solar:alt-arrow-down-line-duotone" />
                                )
                              }
                              sx={{ textTransform: 'none', marginLeft: 2 }}
                            >
                              {expanded ? t('channel_edit.collapse') : t('channel_edit.expand')}
                            </Button>
                          </Box>

                          <Collapse in={expanded}>
                            <Box sx={{ padding: 2, marginTop: -3 }}>
                              {Object.keys(plugin.params).map((paramId) => {
                                const param = plugin.params[paramId];
                                const name = `plugin.${pluginId}.${paramId}`;
                                // 不再在插件面板重复渲染 Gemini vendor，下拉已提升到基础信息区（避免两个入口导致混乱）
                                if (pluginId === 'gemini_video' && paramId === 'vendor') {
                                  return null;
                                }
                                return param.type === 'bool' ? (
                                  <FormControl key={name} fullWidth sx={{ ...theme.typography.otherInput }}>
                                    <FormControlLabel
                                      key={name}
                                      required
                                      control={
                                        <Switch
                                          key={name}
                                          name={name}
                                          disabled={hasTag}
                                          checked={values.plugin?.[pluginId]?.[paramId] || false}
                                          onChange={(event) => {
                                            setFieldValue(name, event.target.checked);
                                          }}
                                        />
                                      }
                                      label={t('channel_edit.isEnable')}
                                    />
                                    <FormHelperText id="helper-tex-channel-key-label"> {customizeT(param.description)} </FormHelperText>
                                  </FormControl>
                                ) : (
                                  <FormControl key={name} fullWidth sx={{ ...theme.typography.otherInput }}>
                                    <TextField
                                      multiline
                                      key={name}
                                      name={name}
                                      disabled={hasTag}
                                      value={values.plugin?.[pluginId]?.[paramId] || ''}
                                      label={customizeT(param.name)}
                                      placeholder={customizeT(param.description)}
                                      onChange={handleChange}
                                    />
                                    <FormHelperText id="helper-tex-channel-key-label"> {customizeT(param.description)} </FormHelperText>
                                  </FormControl>
                                );
                              })}
                            </Box>
                          </Collapse>
                        </Box>
                    );
                  })}
                <DialogActions>
                  <Button onClick={onCancel}>{t('common.cancel')}</Button>
                  <Button disableElevation disabled={isSubmitting} type="submit" variant="contained" color="primary">
                    {t('common.submit')}
                  </Button>
                </DialogActions>
              </form>
            );
          }}
        </Formik>

        {/* 模型选择器弹窗 */}
        <ModelSelectorModal
          open={modelSelectorOpen}
          onClose={() => setModelSelectorOpen(false)}
          onConfirm={(selectedModels, mappings, overwriteModels, overwriteMappings) => {
            // 处理普通模型选择
            handleModelSelectorConfirm(selectedModels, overwriteModels);

            // 处理映射关系
            if (mappings && mappings.length > 0) {
              if (overwriteMappings) {
                // 覆盖映射模式：清空现有映射，使用新的
                tempSetFieldValue('model_mapping', mappings);
              } else {
                // 追加映射模式：
                const existingMappings = tempFormikValues?.model_mapping || [];
                const existingKeys = new Set(existingMappings.map((item) => item.key));
                const newMappings = mappings.filter((item) => !existingKeys.has(item.key));
                const mergedMappings = [...existingMappings, ...newMappings].map((item, index) => ({
                  ...item,
                  index
                }));
                tempSetFieldValue('model_mapping', mergedMappings);
              }
            }
          }}
          channelValues={tempFormikValues}
          prices={prices}
        />
      </DialogContent>
    </Dialog>
  );
};

export default EditModal;

EditModal.propTypes = {
  open: PropTypes.bool,
  channelId: PropTypes.oneOfType([PropTypes.number, PropTypes.string]),
  onCancel: PropTypes.func,
  onOk: PropTypes.func,
  groupOptions: PropTypes.array,
  isTag: PropTypes.bool,
  modelOptions: PropTypes.array,
  prices: PropTypes.array
};
