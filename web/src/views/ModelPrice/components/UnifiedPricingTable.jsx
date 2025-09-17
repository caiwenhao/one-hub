import { useState, useEffect, useCallback } from 'react';
import { useTranslation } from 'react-i18next';
import { useSelector } from 'react-redux';
import {
  Card,
  Stack,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Box,
  InputBase,
  Paper,
  IconButton,
  useMediaQuery,
  Avatar,
  ButtonBase,
  Tooltip
} from '@mui/material';
import { Icon } from '@iconify/react';
import { API } from 'utils/api';
import { showError, ValueFormatter, copy } from 'utils/common';
import { useTheme } from '@mui/material/styles';
import Label from 'ui-component/Label';
import ToggleButtonGroup from 'ui-component/ToggleButton';
import { alpha } from '@mui/material/styles';
import { getUnitPreference, setUnitPreference, PAGE_KEYS, UNIT_OPTIONS } from 'utils/unitPreferences';

const UnifiedPricingTable = () => {
  const { t } = useTranslation();
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  const ownedby = useSelector((state) => state.siteInfo?.ownedby);

  const [rows, setRows] = useState([]);
  const [filteredRows, setFilteredRows] = useState([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [availableModels, setAvailableModels] = useState({});
  const [userGroupMap, setUserGroupMap] = useState({});
  const [selectedGroup, setSelectedGroup] = useState('');
  const [selectedOwnedBy, setSelectedOwnedBy] = useState('all');
  // 使用偏好存储的单位默认值（M单位）
  const [unit, setUnit] = useState(() => getUnitPreference(PAGE_KEYS.UNIFIED_PRICING));
  const [onlyShowAvailable, setOnlyShowAvailable] = useState(false);

  const fetchAvailableModels = useCallback(async () => {
    try {
      const res = await API.get('/api/available_model');
      const { success, message, data } = res.data;
      if (success) {
        setAvailableModels(data);
      } else {
        showError(message);
      }
    } catch (error) {
      console.error(error);
    }
  }, []);

  const fetchUserGroupMap = useCallback(async () => {
    try {
      const res = await API.get('/api/user_group_map');
      const { success, message, data } = res.data;
      if (success) {
        setUserGroupMap(data);
        setSelectedGroup(Object.keys(data)[0]);
      } else {
        showError(message);
      }
    } catch (error) {
      console.error(error);
    }
  }, []);

  useEffect(() => {
    fetchAvailableModels();
    fetchUserGroupMap();
  }, [fetchAvailableModels, fetchUserGroupMap]);

  useEffect(() => {
    if (!availableModels || !userGroupMap || !selectedGroup) return;

    const newRows = Object.entries(availableModels)
      .filter(([, model]) => selectedOwnedBy === 'all' || model.owned_by === selectedOwnedBy)
      .filter(([, model]) => !onlyShowAvailable || model.groups.includes(selectedGroup))
      .map(([modelName, model], index) => {
        const group = userGroupMap[selectedGroup];
        const price = model.groups.includes(selectedGroup)
          ? {
              input: group.ratio * model.price.input,
              output: group.ratio * model.price.output
            }
          : { input: t('modelpricePage.noneGroup'), output: t('modelpricePage.noneGroup') };

        const formatPrice = (value, type) => {
          if (typeof value === 'number') {
            let nowUnit = '';
            let isM = unit === 'M';
            if (type === 'times') {
              isM = false;
            }
            if (type === 'tokens') {
              nowUnit = `/ 1${unit}`;
            }
            return ValueFormatter(value, true, isM) + nowUnit;
          }
          return value;
        };

        return {
          id: index + 1,
          model: modelName,
          provider: model.owned_by,
          userGroup: model.groups,
          type: model.price.type,
          input: formatPrice(price.input, model.price.type),
          output: formatPrice(price.output, model.price.type),
          extraRatios: model.price?.extra_ratios
        };
      });

    setRows(newRows);
    setFilteredRows(newRows);
  }, [availableModels, userGroupMap, selectedGroup, selectedOwnedBy, t, unit, onlyShowAvailable]);

  useEffect(() => {
    const filtered = rows.filter((row) => row.model.toLowerCase().includes(searchQuery.toLowerCase()));
    setFilteredRows(filtered);
  }, [searchQuery, rows]);

  const handleOwnedByChange = (newValue) => {
    setSelectedOwnedBy(newValue);
  };

  const handleGroupChange = (groupKey) => {
    setSelectedGroup(groupKey);
  };

  const handleSearchChange = (event) => {
    setSearchQuery(event.target.value);
  };

  const handleUnitChange = (event, newUnit) => {
    if (newUnit !== null) {
      setUnit(newUnit);
      // 保存用户单位偏好
      setUnitPreference(newUnit, PAGE_KEYS.UNIFIED_PRICING);
    }
  };

  const toggleOnlyShowAvailable = () => {
    setOnlyShowAvailable((prev) => !prev);
  };

  const uniqueOwnedBy = ['all', ...new Set(Object.values(availableModels).map((model) => model.owned_by))];

  const getIconByName = (name) => {
    if (name === 'all') return null;
    const owner = ownedby?.find((item) => item.name === name);
    return owner?.icon;
  };

  const clearSearch = () => {
    setSearchQuery('');
  };

  return (
    <Box
      sx={{
        py: { xs: 8, md: 10 },
        px: { xs: 3, md: 6, lg: 12 },
        background: 'linear-gradient(145deg, #ffffff 0%, #f8fafc 100%)',
        ...(theme.palette.mode === 'dark' && {
          background: 'linear-gradient(145deg, #1e293b 0%, #334155 100%)'
        })
      }}
    >
      <Box sx={{ maxWidth: '1400px', mx: 'auto' }}>
        {/* 标题区域 */}
        <Box sx={{ textAlign: 'center', mb: 6 }}>
          <Typography
            variant="h2"
            sx={{
              fontSize: { xs: '2rem', sm: '2.5rem', md: '3rem' },
              fontWeight: 700,
              color: theme.palette.text.primary,
              mb: 3,
              textShadow: 'none'
            }}
          >
            {t('modelpricePage.detailedPricing')}
          </Typography>
          <Typography
            variant="h6"
            sx={{
              color: theme.palette.text.secondary,
              fontWeight: 300
            }}
          >
            {t('modelpricePage.pricingDescription')}
          </Typography>
        </Box>

        {/* 筛选和搜索区域 */}
        <Card
          elevation={0}
          sx={{
            p: 3,
            mb: 4,
            overflow: 'visible',
            backgroundColor: theme.palette.mode === 'dark' ? alpha(theme.palette.background.paper, 0.6) : theme.palette.background.paper,
            borderRadius: '16px',
            border: '1px solid',
            borderColor: theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.05)',
            boxShadow: '0 8px 32px rgba(0,0,0,0.08)'
          }}
        >
          {/* 搜索和单位选择 */}
          <Box
            sx={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'space-between',
              flexWrap: 'wrap',
              gap: 2,
              mb: 3
            }}
          >
            <Paper
              component="form"
              sx={{
                p: '2px 4px',
                display: 'flex',
                alignItems: 'center',
                width: isMobile ? '100%' : 300,
                borderRadius: '12px',
                border: 'none',
                boxShadow: theme.palette.mode === 'dark' ? '0 2px 8px rgba(0,0,0,0.2)' : '0 2px 8px rgba(0,0,0,0.05)',
                backgroundColor:
                  theme.palette.mode === 'dark' ? alpha(theme.palette.background.default, 0.6) : theme.palette.background.default
              }}
            >
              <IconButton sx={{ p: '8px' }} aria-label="search">
                <Icon icon="eva:search-fill" width={18} height={18} />
              </IconButton>
              <InputBase sx={{ ml: 1, flex: 1 }} placeholder={t('modelpricePage.search')} value={searchQuery} onChange={handleSearchChange} />
              {searchQuery && (
                <IconButton sx={{ p: '8px' }} aria-label="clear" onClick={clearSearch}>
                  <Icon icon="eva:close-fill" width={16} height={16} />
                </IconButton>
              )}
            </Paper>

            <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Unit:
              </Typography>
              <ToggleButtonGroup
                value={unit}
                onChange={handleUnitChange}
                options={UNIT_OPTIONS}
                aria-label="unit toggle"
                size="small"
                sx={{
                  '& .MuiToggleButtonGroup-grouped': {
                    borderRadius: '8px !important',
                    mx: 0.5,
                    border: 0,
                    boxShadow: theme.palette.mode === 'dark' ? '0 1px 4px rgba(0,0,0,0.2)' : '0 1px 4px rgba(0,0,0,0.05)',
                    '&.Mui-selected': {
                      boxShadow: `0 0 0 1px ${theme.palette.primary.main}`
                    }
                  }
                }}
              />
            </Box>
          </Box>

          {/* 模型提供商标签 */}
          <Box sx={{ mb: 3 }}>
            <Typography
              variant="subtitle1"
              sx={{
                mb: 1.5,
                fontWeight: 600,
                color: theme.palette.text.primary,
                display: 'flex',
                alignItems: 'center',
                gap: 1
              }}
            >
              <Icon icon="eva:globe-outline" width={18} height={18} />
              {t('modelpricePage.channelType')}
            </Typography>
            <Box
              sx={{
                display: 'flex',
                flexWrap: 'wrap',
                gap: 1
              }}
            >
              {uniqueOwnedBy.map((ownedBy, index) => {
                const isSelected = selectedOwnedBy === ownedBy;
                return (
                  <ButtonBase
                    key={index}
                    onClick={() => handleOwnedByChange(ownedBy)}
                    sx={{
                      borderRadius: '8px',
                      overflow: 'hidden',
                      position: 'relative',
                      transition: 'all 0.2s ease',
                      transform: isSelected ? 'translateY(-1px)' : 'none',
                      '&:hover': {
                        transform: 'translateY(-1px)'
                      }
                    }}
                  >
                    <Box
                      sx={{
                        display: 'flex',
                        alignItems: 'center',
                        gap: 0.75,
                        py: 0.75,
                        px: 1.5,
                        borderRadius: '8px',
                        backgroundColor: isSelected
                          ? alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.25 : 0.1)
                          : theme.palette.mode === 'dark'
                            ? alpha(theme.palette.background.default, 0.5)
                            : theme.palette.background.default,
                        border: `1px solid ${
                          isSelected ? theme.palette.primary.main : theme.palette.mode === 'dark' ? alpha('#fff', 0.08) : alpha('#000', 0.05)
                        }`,
                        boxShadow: isSelected ? `0 2px 8px ${alpha(theme.palette.primary.main, 0.2)}` : 'none'
                      }}
                    >
                      {ownedBy !== 'all' ? (
                        <Avatar
                          src={getIconByName(ownedBy)}
                          alt={ownedBy}
                          sx={{
                            width: 20,
                            height: 20,
                            backgroundColor: theme.palette.mode === 'dark' ? '#fff' : theme.palette.background.paper,
                            '.MuiAvatar-img': {
                              objectFit: 'contain',
                              padding: '2px'
                            }
                          }}
                        >
                          {ownedBy.charAt(0).toUpperCase()}
                        </Avatar>
                      ) : (
                        <Icon
                          icon="eva:grid-outline"
                          width={18}
                          height={18}
                          color={isSelected ? theme.palette.primary.main : theme.palette.text.secondary}
                        />
                      )}
                      <Typography
                        variant="body2"
                        sx={{
                          fontWeight: isSelected ? 600 : 500,
                          color: isSelected ? theme.palette.primary.main : theme.palette.text.primary,
                          fontSize: '0.8125rem'
                        }}
                      >
                        {ownedBy === 'all' ? t('modelpricePage.all') : ownedBy}
                      </Typography>
                    </Box>
                  </ButtonBase>
                );
              })}
            </Box>
          </Box>

          {/* 用户组标签 */}
          <Box sx={{ mb: 0 }}>
            <Box
              sx={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'space-between',
                mb: 1.5
              }}
            >
              <Typography
                variant="subtitle1"
                sx={{
                  fontWeight: 600,
                  color: theme.palette.text.primary,
                  display: 'flex',
                  alignItems: 'center',
                  gap: 1
                }}
              >
                <Icon icon="eva:people-outline" width={18} height={18} />
                {t('modelpricePage.group')}
              </Typography>

              <Tooltip title={onlyShowAvailable ? t('modelpricePage.showAll') : t('modelpricePage.onlyAvailable')} arrow>
                <ButtonBase
                  onClick={toggleOnlyShowAvailable}
                  sx={{
                    position: 'relative',
                    borderRadius: '20px',
                    overflow: 'hidden',
                    transition: 'all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1)',
                    '&:hover': {
                      transform: 'translateY(-1px)',
                      boxShadow: theme.palette.mode === 'dark' ? '0 3px 10px rgba(0,0,0,0.4)' : '0 3px 10px rgba(0,0,0,0.1)'
                    },
                    '&:active': {
                      transform: 'translateY(0px)'
                    }
                  }}
                >
                  <Box
                    sx={{
                      display: 'flex',
                      alignItems: 'center',
                      gap: 1,
                      py: 0.6,
                      px: 1.5,
                      background: onlyShowAvailable
                        ? theme.palette.mode === 'dark'
                          ? `linear-gradient(45deg, ${alpha(theme.palette.primary.main, 0.8)}, ${alpha(theme.palette.primary.dark, 0.9)})`
                          : `linear-gradient(45deg, ${theme.palette.primary.main}, ${theme.palette.primary.dark})`
                        : theme.palette.mode === 'dark'
                          ? alpha(theme.palette.background.paper, 0.6)
                          : alpha(theme.palette.background.paper, 1),
                      border: `1px solid ${
                        onlyShowAvailable
                          ? theme.palette.primary.main
                          : theme.palette.mode === 'dark'
                            ? alpha('#fff', 0.1)
                            : alpha('#000', 0.08)
                      }`,
                      borderRadius: '20px',
                      boxShadow: onlyShowAvailable
                        ? `0 2px 8px ${alpha(theme.palette.primary.main, 0.4)}`
                        : theme.palette.mode === 'dark'
                          ? '0 2px 6px rgba(0,0,0,0.2)'
                          : '0 2px 6px rgba(0,0,0,0.05)'
                    }}
                  >
                    <Box
                      sx={{
                        width: 20,
                        height: 20,
                        borderRadius: '50%',
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        backgroundColor: onlyShowAvailable
                          ? '#fff'
                          : theme.palette.mode === 'dark'
                            ? alpha(theme.palette.primary.main, 0.2)
                            : alpha(theme.palette.primary.main, 0.1),
                        transition: 'all 0.2s ease'
                      }}
                    >
                      <Icon
                        icon={onlyShowAvailable ? 'eva:checkmark-outline' : 'eva:funnel-outline'}
                        width={14}
                        height={14}
                        color={onlyShowAvailable ? theme.palette.primary.main : theme.palette.text.secondary}
                      />
                    </Box>
                    <Typography
                      variant="body2"
                      sx={{
                        fontWeight: 600,
                        color: onlyShowAvailable ? '#fff' : theme.palette.text.primary,
                        fontSize: '0.75rem',
                        letterSpacing: '0.01em',
                        textTransform: 'uppercase'
                      }}
                    >
                      {t('modelpricePage.onlyAvailable')}
                    </Typography>
                  </Box>
                </ButtonBase>
              </Tooltip>
            </Box>
            <Box
              sx={{
                display: 'flex',
                flexWrap: 'wrap',
                gap: 1
              }}
            >
              {Object.entries(userGroupMap).map(([key, group]) => {
                const isSelected = selectedGroup === key;
                return (
                  <Tooltip
                    key={key}
                    title={group.ratio > 0 ? `${t('modelpricePage.rate')}: x${group.ratio}` : t('modelpricePage.free')}
                    arrow
                  >
                    <ButtonBase
                      onClick={() => handleGroupChange(key)}
                      sx={{
                        position: 'relative',
                        borderRadius: '8px',
                        overflow: 'hidden',
                        transition: 'all 0.2s ease',
                        transform: isSelected ? 'translateY(-1px)' : 'none',
                        '&:hover': {
                          transform: 'translateY(-1px)'
                        }
                      }}
                    >
                      <Box
                        sx={{
                          display: 'flex',
                          alignItems: 'center',
                          gap: 1,
                          py: 0.75,
                          px: 1.5,
                          borderRadius: '8px',
                          backgroundColor: isSelected
                            ? alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.25 : 0.1)
                            : theme.palette.mode === 'dark'
                              ? alpha(theme.palette.background.default, 0.5)
                              : theme.palette.background.default,
                          border: `1px solid ${
                            isSelected
                              ? theme.palette.primary.main
                              : theme.palette.mode === 'dark'
                                ? alpha('#fff', 0.08)
                                : alpha('#000', 0.05)
                          }`,
                          boxShadow: isSelected ? `0 2px 8px ${alpha(theme.palette.primary.main, 0.2)}` : 'none'
                        }}
                      >
                        <Icon
                          icon={isSelected ? 'eva:checkmark-circle-2-fill' : 'eva:radio-button-off-outline'}
                          width={16}
                          height={16}
                          color={isSelected ? theme.palette.primary.main : theme.palette.text.secondary}
                        />
                        <Typography
                          variant="body2"
                          sx={{
                            fontWeight: isSelected ? 600 : 500,
                            color: isSelected ? theme.palette.primary.main : theme.palette.text.primary,
                            fontSize: '0.8125rem'
                          }}
                        >
                          {group.name}
                        </Typography>
                        {group.ratio > 0 ? (
                          <Box
                            sx={{
                              display: 'flex',
                              alignItems: 'center',
                              justifyContent: 'center',
                              minWidth: 24,
                              height: 16,
                              borderRadius: '4px',
                              backgroundColor:
                                group.ratio > 1
                                  ? alpha(theme.palette.warning.main, theme.palette.mode === 'dark' ? 0.3 : 0.2)
                                  : alpha(theme.palette.info.main, theme.palette.mode === 'dark' ? 0.3 : 0.2),
                              color: group.ratio > 1 ? theme.palette.warning.main : theme.palette.info.main,
                              fontSize: '0.6875rem',
                              fontWeight: 600,
                              px: 0.5
                            }}
                          >
                            x{group.ratio}
                          </Box>
                        ) : (
                          <Box
                            sx={{
                              display: 'flex',
                              alignItems: 'center',
                              justifyContent: 'center',
                              minWidth: 24,
                              height: 16,
                              borderRadius: '4px',
                              backgroundColor: alpha(theme.palette.success.main, theme.palette.mode === 'dark' ? 0.3 : 0.2),
                              color: theme.palette.success.main,
                              fontSize: '0.6875rem',
                              fontWeight: 600,
                              px: 0.5
                            }}
                          >
                            {t('modelpricePage.free')}
                          </Box>
                        )}
                      </Box>
                    </ButtonBase>
                  </Tooltip>
                );
              })}
            </Box>
          </Box>
        </Card>

        {/* 价格表格 */}
        <Card
          sx={{
            borderRadius: '16px',
            overflow: 'hidden',
            boxShadow: '0 8px 32px rgba(0,0,0,0.08)',
            border: '1px solid',
            borderColor: theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.05)',
            background: theme.palette.mode === 'dark'
              ? 'rgba(30, 41, 59, 0.8)'
              : 'rgba(255, 255, 255, 0.9)',
            backdropFilter: 'blur(10px)'
          }}
        >
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow
                  sx={{
                    background: `linear-gradient(135deg, ${theme.palette.primary.main}15 0%, ${theme.palette.secondary.main}15 100%)`,
                    '& .MuiTableCell-head': {
                      borderBottom: 'none'
                    }
                  }}
                >
                  <TableCell width="25%" sx={{ fontWeight: 700, py: 2, px: 3, fontSize: '0.95rem' }}>
                    {t('modelpricePage.model')}
                  </TableCell>
                  <TableCell width="15%" sx={{ fontWeight: 700, py: 2, px: 3, fontSize: '0.95rem' }}>
                    {t('modelpricePage.channelType')}
                  </TableCell>
                  <TableCell width="10%" sx={{ fontWeight: 700, py: 2, px: 3, fontSize: '0.95rem' }}>
                    {t('modelpricePage.type')}
                  </TableCell>
                  <TableCell width="17.5%" sx={{ fontWeight: 700, py: 2, px: 3, fontSize: '0.95rem' }}>
                    {t('modelpricePage.inputMultiplier')}
                  </TableCell>
                  <TableCell width="17.5%" sx={{ fontWeight: 700, py: 2, px: 3, fontSize: '0.95rem' }}>
                    {t('modelpricePage.outputMultiplier')}
                  </TableCell>
                  <TableCell width="15%" sx={{ fontWeight: 700, py: 2, px: 3, fontSize: '0.95rem' }}>
                    {t('modelpricePage.other')}
                  </TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {filteredRows.length > 0 ? (
                  filteredRows.map((row, index) => (
                    <TableRow
                      key={row.id}
                      sx={{
                        '&:hover': {
                          backgroundColor: theme.palette.mode === 'dark'
                            ? 'rgba(255,255,255,0.05)'
                            : 'rgba(0,0,0,0.02)',
                          transform: 'translateY(-1px)',
                          boxShadow: '0 4px 12px rgba(0,0,0,0.1)'
                        },
                        transition: 'all 0.2s ease',
                        borderBottom: index === filteredRows.length - 1 ? 'none' : undefined
                      }}
                    >
                      <TableCell sx={{ py: 2, px: 3 }}>
                        <Stack direction="row" justifyContent="flex-start" alignItems="center" spacing={1}>
                          <Typography variant="body2" sx={{ fontWeight: 500 }}>
                            {row.model}
                          </Typography>
                          <IconButton size="small" onClick={() => copy(row.model)}>
                            <Icon icon="eva:copy-outline" width={16} height={16} />
                          </IconButton>
                        </Stack>
                      </TableCell>
                      <TableCell sx={{ py: 2, px: 3 }}>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                          <Avatar
                            src={getIconByName(row.provider)}
                            alt={row.provider}
                            sx={{
                              width: 20,
                              height: 20,
                              backgroundColor: theme.palette.mode === 'dark' ? '#fff' : theme.palette.background.paper,
                              '.MuiAvatar-img': {
                                objectFit: 'contain',
                                padding: '2px'
                              }
                            }}
                          >
                            {row.provider?.charAt(0).toUpperCase()}
                          </Avatar>
                          <Typography variant="body2">{row.provider}</Typography>
                        </Box>
                      </TableCell>
                      <TableCell sx={{ py: 2, px: 3 }}>
                        {row.type === 'tokens' ? (
                          <Label
                            color="primary"
                            sx={{
                              borderRadius: '6px',
                              fontWeight: 500,
                              fontSize: '0.75rem',
                              py: 0.25,
                              px: 0.75
                            }}
                          >
                            {t('modelpricePage.tokens')}
                          </Label>
                        ) : (
                          <Label
                            color="secondary"
                            sx={{
                              borderRadius: '6px',
                              fontWeight: 500,
                              fontSize: '0.75rem',
                              py: 0.25,
                              px: 0.75
                            }}
                          >
                            {t('modelpricePage.times')}
                          </Label>
                        )}
                      </TableCell>
                      <TableCell sx={{ py: 2, px: 3 }}>
                        <Label
                          color="info"
                          variant="outlined"
                          sx={{
                            borderRadius: '6px',
                            fontWeight: 500,
                            fontSize: '0.75rem',
                            py: 0.25,
                            px: 0.75
                          }}
                        >
                          {row.input}
                        </Label>
                      </TableCell>
                      <TableCell sx={{ py: 2, px: 3 }}>
                        <Label
                          color="info"
                          variant="outlined"
                          sx={{
                            borderRadius: '6px',
                            fontWeight: 500,
                            fontSize: '0.75rem',
                            py: 0.25,
                            px: 0.75
                          }}
                        >
                          {row.output}
                        </Label>
                      </TableCell>
                      <TableCell sx={{ py: 2, px: 3 }}>{getOther(t, row.extraRatios)}</TableCell>
                    </TableRow>
                  ))
                ) : (
                  <TableRow>
                    <TableCell colSpan={6} align="center" sx={{ py: 6 }}>
                      <Stack spacing={2} alignItems="center">
                        <Icon icon="eva:search-outline" width={48} height={48} color={theme.palette.text.secondary} />
                        <Typography variant="h6" color="text.secondary">
                          {t('common.noData')}
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          尝试调整筛选条件或搜索关键词
                        </Typography>
                      </Stack>
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </TableContainer>
        </Card>
      </Box>
    </Box>
  );
};

// 辅助函数
function getOther(t, extraRatios) {
  if (!extraRatios) return '';

  return (
    <Stack direction="column" spacing={0.5}>
      {Object.entries(extraRatios).map(([key, value]) => (
        <Label
          key={key}
          color="primary"
          variant="outlined"
          sx={{
            borderRadius: '4px',
            fontSize: '0.75rem',
            py: 0.25,
            px: 0.75
          }}
        >
          {t(`modelpricePage.${key}`)}: {value}
        </Label>
      ))}
    </Stack>
  );
}

export default UnifiedPricingTable;
