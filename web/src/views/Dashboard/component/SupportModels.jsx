import { useState, useEffect } from 'react';
import { API } from 'utils/api';
import { showError, copy } from 'utils/common';
import { Box, Card, Stack, alpha, Tooltip, IconButton, Typography } from '@mui/material';
import Label from 'ui-component/Label';
import { useTranslation } from 'react-i18next';
import { ExpandMore, ExpandLess } from '@mui/icons-material';
import { useSelector } from 'react-redux';
import IconWrapper from 'ui-component/IconWrapper';

const SupportModels = () => {
  const [modelList, setModelList] = useState([]);
  const [priceMap, setPriceMap] = useState({}); // model -> price_display
  const [expanded, setExpanded] = useState(false);
  const { t } = useTranslation();
  const ownedby = useSelector((state) => state.siteInfo?.ownedby);

  const fetchModels = async () => {
    try {
      const res = await API.get(`/api/available_model`);
      const { data, success } = res.data;
      if (!success) return;

      const modelGroup = Object.entries(data).reduce((acc, [modelId, modelInfo]) => {
        const { owned_by } = modelInfo;
        if (!acc[owned_by]) {
          acc[owned_by] = [];
        }
        acc[owned_by].push(modelId);
        return acc;
      }, {});

      // build price map for tooltip
      const map = {};
      Object.entries(data).forEach(([modelId, modelInfo]) => {
        if (modelInfo?.price_display) {
          map[modelId] = modelInfo.price_display;
        }
      });

      Object.values(modelGroup).forEach((models) => models.sort());

      const sortedModelGroup = Object.keys(modelGroup)
        .sort()
        .reduce((acc, key) => {
          acc[key] = modelGroup[key];
          return acc;
        }, {});

      setModelList(sortedModelGroup);
      setPriceMap(map);
    } catch (error) {
      showError(error.message);
    }
  };

  useEffect(() => {
    fetchModels();
  }, []);

  const getIconByName = (name) => {
    const owner = ownedby.find((item) => item.name === name);
    return owner?.icon;
  };

  return (
    <Card>
      <Box sx={{ p: 2 }}>
        <Box sx={{ position: 'relative' }}>
          <Stack
            direction="row"
            alignItems="center"
            spacing={2}
            sx={{
              mb: expanded ? 2 : 0,
              pr: 5
            }}
          >
            <Typography variant="subtitle1" sx={{ color: 'text.secondary', whiteSpace: 'nowrap' }}>
              {t('dashboard_index.model_price')}:
            </Typography>

            {!expanded && (
              <Box
                sx={{
                  flex: 1,
                  overflow: 'auto',
                  display: 'flex',
                  gap: 1,
                  '&::-webkit-scrollbar': { display: 'none' },
                  scrollbarWidth: 'none',
                  maskImage: 'linear-gradient(to right, black 90%, transparent 100%)',
                  WebkitMaskImage: 'linear-gradient(to right, black 90%, transparent 100%)'
                }}
              >
                {Object.entries(modelList)
                  .slice(0, 1)
                  .map(([provider, models]) => (
                    <Box
                      key={provider}
                      sx={{
                        display: 'flex',
                        gap: 1,
                        alignItems: 'center'
                      }}
                    >
                      <Typography
                        variant="subtitle2"
                        sx={{
                          color: 'text.secondary',
                          whiteSpace: 'nowrap',
                          fontWeight: 'bold'
                        }}
                      >
                        {provider}:
                      </Typography>
                      {models.map((model) => {
                        const p = priceMap[model];
                        let tip = p
                          ? p.type === 'times'
                            ? `单价：$${p.input_usd ?? p.input_rmb} / 次`
                            : `输入：$${p.input_usd ?? p.input_rmb} / 1k，输出：$${p.output_usd ?? p.output_rmb} / 1k`
                          : undefined;
                        // minimaxi 视频模型为“组合计费”，基础模型可能无单价，给出友好提示
                        if (!tip) {
                          const lower = model.toLowerCase();
                          const isMiniMaxVideoBase =
                            lower.includes('minimax-hailuo-02') ||
                            lower === 't2v-01' ||
                            lower === 't2v-01-director' ||
                            lower === 'i2v-01' ||
                            lower === 'i2v-01-live' ||
                            lower === 's2v-01' ||
                            lower.startsWith('minimax-');
                          if (isMiniMaxVideoBase) {
                            tip = '视频按分辨率/时长精确计费，提交时自动匹配组合价格';
                          }
                        }
                        const label = (
                          <Label
                            key={model}
                            variant="soft"
                            color="primary"
                            onClick={() => copy(model, t('dashboard_index.model_name'))}
                            sx={{
                              cursor: 'pointer',
                              whiteSpace: 'nowrap',
                              '&:hover': {
                                bgcolor: (theme) => alpha(theme.palette.primary.main, 0.16)
                              }
                            }}
                          >
                            {model}
                          </Label>
                        );
                        return tip ? (
                          <Tooltip key={model} title={tip} placement="top">
                            <span>{label}</span>
                          </Tooltip>
                        ) : (
                          label
                        );
                      })}
                    </Box>
                  ))}
              </Box>
            )}
          </Stack>

          <Box
            sx={{
              position: 'absolute',
              right: 0,
              top: -2,
              bgcolor: (theme) => theme.palette.background.paper,
              background: (theme) => `linear-gradient(to right, transparent, ${theme.palette.background.paper} 20%)`,
              pl: 1
            }}
          >
            <Tooltip>
              <IconButton
                size="small"
                onClick={() => setExpanded(!expanded)}
                sx={{
                  color: 'text.secondary',
                  '&:hover': {
                    color: 'text.primary'
                  }
                }}
              >
                {expanded ? <ExpandLess sx={{ width: 20 }} /> : <ExpandMore sx={{ width: 20 }} />}
              </IconButton>
            </Tooltip>
          </Box>
        </Box>

        {expanded && (
          <Stack spacing={2}>
            {Object.entries(modelList).map(([provider, models]) => (
              <Box key={provider}>
                <Typography
                  variant="subtitle2"
                  sx={{
                    color: 'text.secondary',
                    display: 'block',
                    mb: 1,
                    fontWeight: 'bold'
                  }}
                >
                  <Stack direction="row" alignItems="center" spacing={1}>
                    <IconWrapper url={getIconByName(provider)} />
                    <span>{provider}</span>
                  </Stack>
                </Typography>
                <Box
                  sx={{
                    display: 'flex',
                    flexWrap: 'wrap',
                    gap: 1,
                    pl: 1
                  }}
                >
                  {models.map((model) => {
                    const p = priceMap[model];
                    let tip = p
                      ? p.type === 'times'
                        ? `单价：$${p.input_usd ?? p.input_rmb} / 次`
                        : `输入：$${p.input_usd ?? p.input_rmb} / 1k，输出：$${p.output_usd ?? p.output_rmb} / 1k`
                      : undefined;
                    if (!tip) {
                      const lower = model.toLowerCase();
                      const isMiniMaxVideoBase =
                        lower.includes('minimax-hailuo-02') ||
                        lower === 't2v-01' ||
                        lower === 't2v-01-director' ||
                        lower === 'i2v-01' ||
                        lower === 'i2v-01-live' ||
                        lower === 's2v-01' ||
                        lower.startsWith('minimax-');
                      if (isMiniMaxVideoBase) {
                        tip = '视频按分辨率/时长精确计费，提交时自动匹配组合价格';
                      }
                    }
                    const label = (
                      <Label
                        key={model}
                        variant="soft"
                        color="primary"
                        onClick={() => copy(model, t('dashboard_index.model_name'))}
                        sx={{
                          cursor: 'pointer',
                          '&:hover': {
                            bgcolor: (theme) => alpha(theme.palette.primary.main, 0.16)
                          }
                        }}
                      >
                        {model}
                      </Label>
                    );
                    return tip ? (
                      <Tooltip key={model} title={tip} placement="top">
                        <span>{label}</span>
                      </Tooltip>
                    ) : (
                      label
                    );
                  })}
                </Box>
              </Box>
            ))}
          </Stack>
        )}
      </Box>
    </Card>
  );
};

export default SupportModels;
