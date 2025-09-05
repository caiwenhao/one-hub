import PropTypes from 'prop-types';
import * as Yup from 'yup';
import { Formik } from 'formik';
import { useTheme } from '@mui/material/styles';
import { useCallback, useEffect, useState } from 'react';
import { convertPrice, convertPriceData, isValidUnitType } from './priceConverter';
import {
  Alert,
  Autocomplete,
  Button,
  Checkbox,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  FormControl,
  FormHelperText,
  InputAdornment,
  InputLabel,
  MenuItem,
  OutlinedInput,
  Paper,
  Select,
  Stack,
  TextField,
  Typography
} from '@mui/material';

import { showError, showSuccess, trims } from 'utils/common';
import { API } from 'utils/api';
import { createFilterOptions } from '@mui/material/Autocomplete';
import { priceType, ValueFormatter } from './util';
import CheckBoxOutlineBlankIcon from '@mui/icons-material/CheckBoxOutlineBlank';
import CheckBoxIcon from '@mui/icons-material/CheckBox';
import { useTranslation } from 'react-i18next';
import ToggleButtonGroup from 'ui-component/ToggleButton';
import Decimal from 'decimal.js';
import { ExtraRatiosSelector } from './ExtraRatiosSelector';

const icon = <CheckBoxOutlineBlankIcon fontSize="small" />;
const checkedIcon = <CheckBoxIcon fontSize="small" />;

const filter = createFilterOptions();

// 单一模式下的表单验证
const validateSingleMode = (t, values, rows) => {
  if (values.model === '') {
    return t('pricing_edit.requiredModelName');
  }

  if (values.type !== 'tokens' && values.type !== 'times') {
    return t('pricing_edit.typeCheck');
  }

  if (values.channel_type <= 0) {
    return t('pricing_edit.channelTypeErr2');
  }

  // 判断model是否是唯一值
  if (rows && rows.filter((r) => r.model === values.model && (values.isNew || r.id !== values.id)).length > 0) {
    return t('pricing_edit.modelNameRe');
  }

  if (values.input === '' || values.input < 0) {
    return t('pricing_edit.inputVal');
  }
  if (values.output === '' || values.output < 0) {
    return t('pricing_edit.outputVal');
  }
  return false;
};

// 多选模式下的表单验证
const getValidationSchema = (t) =>
  Yup.object().shape({
    is_edit: Yup.boolean(),
    type: Yup.string().oneOf(['tokens', 'times'], t('pricing_edit.typeErr')).required(t('pricing_edit.requiredType')),
    channel_type: Yup.number().min(1, t('pricing_edit.channelTypeErr')).required(t('pricing_edit.requiredChannelType')),
    input: Yup.number().required(t('pricing_edit.requiredInput')),
    output: Yup.number().required(t('pricing_edit.requiredOutput')),
    models: Yup.array().min(1, t('pricing_edit.requiredModels'))
  });

// 多选模式初始值
const multipleOriginInputs = {
  is_edit: false,
  type: 'tokens',
  channel_type: 1,
  input: 0,
  output: 0,
  locked: false,
  models: [],
  extra_ratios: {}
};

// 单一模式初始值
const singleOriginInputs = {
  model: '',
  type: 'tokens',
  channel_type: 1,
  input: 0,
  output: 0,
  locked: false,
  extra_ratios: {}
};

const EditModal = ({
  open,
  pricesItem,
  onCancel,
  onOk,
  ownedby,
  noPriceModel,
  singleMode = false,
  price = null,
  rows = [],
  onSaveSingle = null,
  unit = 'K'
}) => {
  const { t } = useTranslation();
  const theme = useTheme();
  const [inputs, setInputs] = useState(singleMode ? singleOriginInputs : multipleOriginInputs);
  const [selectModel, setSelectModel] = useState([]);
  const [errors, setErrors] = useState({});

  const [unitType, setUnitType] = useState('rate');
  const [localUnit, setLocalUnit] = useState(unit);
  const [previousUnitType, setPreviousUnitType] = useState('rate');
  const [previousLocalUnit, setPreviousLocalUnit] = useState(unit);
  const [isConverting, setIsConverting] = useState(false);
  const [pendingConversion, setPendingConversion] = useState(null);

  // 当外部unit变化时同步本地unit
  useEffect(() => {
    setLocalUnit(unit);
  }, [unit]);

  const calculateRate = useCallback(
    (price) => {
      if (price === '' || price == null) return 0;
      // 使用统一转换器，确保与USD/RMB切换逻辑一致
      return convertPrice(price, unitType, localUnit, 'rate', 'K');
    },
    [unitType, localUnit]
  );

  const unitTypeOptions = [
    { value: 'rate', label: t('modelpricePage.rate') },
    { value: 'USD', label: 'USD' },
    { value: 'RMB', label: 'RMB' }
  ];

  const lockedOptions = [
    { value: true, label: t('pricing_edit.locked') },
    { value: false, label: t('pricing_edit.unlocked') }
  ];

  const unitOptions = [
    { value: 'K', label: 'K' },
    { value: 'M', label: 'M' }
  ];

  const handleEndAdornment = useCallback(
    (value) => {
      let endAdornment = '';

      switch (unitType) {
        case 'rate':
          endAdornment = ValueFormatter(value);
          break;
        case 'USD':
        case 'RMB':
          endAdornment = value === 0 ? 'Free' : calculateRate(value) + ' Rate';
          break;
      }

      return endAdornment;
    },
    [unitType, calculateRate]
  );

  const handleStartAdornment = useCallback(() => {
    switch (unitType) {
      case 'rate':
        return 'Rate：';
      case 'USD':
        return `USD(${localUnit})：`;
      case 'RMB':
        return `RMB(${localUnit})：`;
    }
  }, [unitType, localUnit]);

  // 表单提交处理
  const submit = async (values, { setErrors, setStatus, setSubmitting }) => {
    setSubmitting(true);

    // 单一模式处理
    if (singleMode) {
      // 验证表单
      const validationError = validateSingleMode(t, values, rows);
      if (validationError) {
        setStatus({ success: false });
        setErrors({ submit: validationError });
        return;
      }

      try {
        if (onSaveSingle) {
          await onSaveSingle({
            ...values,
            input: calculateRate(values.input),
            output: calculateRate(values.output)
          });
        }
        setSubmitting(false);
        return;
      } catch (error) {
        setStatus({ success: false });
        setErrors({ submit: error.message });
        return;
      }
    }

    // 多选模式处理
    values.models = trims(values.models);
    try {
      const res = await API.post(`/api/prices/multiple`, {
        original_models: inputs.models,
        models: values.models,
        price: {
          model: 'batch',
          type: values.type,
          channel_type: values.channel_type,
          input: calculateRate(values.input),
          output: calculateRate(values.output),
          locked: values.locked,
          extra_ratios: values.extra_ratios
        }
      });
      const { success, message } = res.data;
      if (success) {
        showSuccess(t('common.saveSuccess'));
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
    onOk();
  };

  // 处理表单字段变化 (单模式用)
  const handleChange = (event) => {
    if (!singleMode) return; // 单一模式专用

    const { name, value, checked } = event.target;

    let finalValue;
    if (name === 'input' || name === 'output') {
      finalValue = value === '' ? 0 : Number(value);
    } else {
      finalValue = name === 'locked' ? checked : value;
    }

    setInputs((prev) => ({
      ...prev,
      [name]: finalValue
    }));

    if (errors[name]) {
      setErrors((prev) => ({ ...prev, [name]: null }));
    }
  };

  // 处理extra_ratios变化 (单模式用)
  const handleChangeExtraRatios = (newExtraRatios) => {
    if (!singleMode) return; // 单一模式专用

    setInputs((prev) => ({
      ...prev,
      extra_ratios: newExtraRatios
    }));
  };

  useEffect(() => {
    if (singleMode) {
      // 单一模式初始化表单
      if (price) {
        setInputs({
          ...price,
          extra_ratios: price.extra_ratios || {}
        });
      } else {
        setInputs(singleOriginInputs);
      }
      setErrors({});
    } else {
      // 多选模式初始化
      if (pricesItem) {
        setSelectModel(pricesItem.models.concat(noPriceModel));
        setInputs(pricesItem);
      } else {
        setSelectModel(noPriceModel);
        setInputs(multipleOriginInputs);
      }
    }
  }, [singleMode, price, pricesItem, noPriceModel]);

  useEffect(() => {
    if (open) {
      setUnitType('rate');
      setLocalUnit('K');
      setPreviousUnitType('rate');
      setPreviousLocalUnit('K');
    }
  }, [open]);

  /**
   * 监听单位变化：仅同步 previous*，不在这里做转换
   * 转换统一在 pendingConversion 的 effect 中处理，避免重复转换
   */
  useEffect(() => {
    const hasUnitChanged = unitType !== previousUnitType || localUnit !== previousLocalUnit;
    if (hasUnitChanged) {
      setPreviousUnitType(unitType);
      setPreviousLocalUnit(localUnit);
    }
  }, [unitType, localUnit, previousUnitType, previousLocalUnit]);

  // 处理待处理的转换
  useEffect(() => {
    if (pendingConversion && singleMode) {
      console.log('Processing pending conversion:', pendingConversion);
      setIsConverting(true);

      // 使用 setTimeout 来确保状态更新完成
      setTimeout(() => {
        setInputs(prevInputs => {
          const convertedData = convertPriceData(
            prevInputs,
            pendingConversion.fromUnitType,
            pendingConversion.fromLocalUnit,
            pendingConversion.toUnitType,
            pendingConversion.toLocalUnit
          );

          console.log('Conversion completed:', {
            from: `${pendingConversion.fromUnitType}(${pendingConversion.fromLocalUnit})`,
            to: `${pendingConversion.toUnitType}(${pendingConversion.toLocalUnit})`,
            original: { input: prevInputs.input, output: prevInputs.output },
            converted: { input: convertedData.input, output: convertedData.output }
          });

          return convertedData;
        });

        // 更新前一次的单位状态
        setPreviousUnitType(pendingConversion.toUnitType);
        setPreviousLocalUnit(pendingConversion.toLocalUnit);

        setIsConverting(false);
        setPendingConversion(null);
      }, 50);
    }
  }, [pendingConversion, singleMode]);

  // 渲染类型选择表单
  const renderTypeSelector = (formProps) => {
    const { errors = {}, touched = {}, handleBlur, handleChange: formikHandleChange, values = {} } = formProps || {};

    return (
      <FormControl
        fullWidth
        error={singleMode ? !!errors.type : Boolean(touched?.type && errors?.type)}
        sx={{ ...theme.typography.otherInput }}
      >
        <InputLabel htmlFor="type-label">{t('pricing_edit.type')}</InputLabel>
        <Select
          id="type-label"
          label={t('pricing_edit.type')}
          value={singleMode ? inputs.type : values?.type}
          name="type"
          onBlur={handleBlur}
          onChange={singleMode ? handleChange : formikHandleChange}
          MenuProps={{
            PaperProps: {
              style: {
                maxHeight: 200
              }
            }
          }}
        >
          {Object.values(priceType).map((option) => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </Select>
        {!singleMode && touched?.type && errors?.type && (
          <FormHelperText error id="helper-tex-type-label">
            {errors.type}
          </FormHelperText>
        )}
      </FormControl>
    );
  };

  // 渲染渠道类型选择表单
  const renderChannelTypeSelector = (formProps) => {
    const { errors = {}, touched = {}, handleBlur, handleChange: formikHandleChange, values = {} } = formProps || {};

    return (
      <FormControl
        fullWidth
        error={singleMode ? !!errors.channel_type : Boolean(touched?.channel_type && errors?.channel_type)}
        sx={{ ...theme.typography.otherInput }}
      >
        <InputLabel htmlFor="channel_type-label">{t('pricing_edit.channelType')}</InputLabel>
        <Select
          id="channel_type-label"
          label={t('pricing_edit.channelType')}
          value={singleMode ? inputs.channel_type : values?.channel_type}
          name="channel_type"
          onBlur={handleBlur}
          onChange={singleMode ? handleChange : formikHandleChange}
          MenuProps={{
            PaperProps: {
              style: {
                maxHeight: 200
              }
            }
          }}
        >
          {ownedby.map((option) => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </Select>
        {!singleMode && touched?.channel_type && errors?.channel_type && (
          <FormHelperText error id="helper-tex-channel_type-label">
            {errors.channel_type}
          </FormHelperText>
        )}
      </FormControl>
    );
  };

  // 处理单位变化的统一函数
  const handleUnitChange = useCallback((newUnitType, newLocalUnit) => {
    const actualUnitType = newUnitType || unitType;
    const actualLocalUnit = newLocalUnit || localUnit;

    console.log('Unit change requested:', {
      from: `${unitType}(${localUnit})`,
      to: `${actualUnitType}(${actualLocalUnit})`,
      singleMode
    });

    // 如果在单一模式下且单位确实发生了变化
    if (
      singleMode &&
      (actualUnitType !== unitType || actualLocalUnit !== localUnit) &&
      isValidUnitType(unitType, localUnit) &&
      isValidUnitType(actualUnitType, actualLocalUnit)
    ) {
      const onlyKMChangeUnderRate =
        actualUnitType === unitType && unitType === 'rate' && actualLocalUnit !== localUnit;

      // 在倍率(rate)模式下，K/M 切换不应改变数值，因此不触发转换
      if (!onlyKMChangeUnderRate) {
        setPendingConversion({
          fromUnitType: unitType,
          fromLocalUnit: localUnit,
          toUnitType: actualUnitType,
          toLocalUnit: actualLocalUnit
        });
      }
    }

    // 更新状态
    if (newUnitType) setUnitType(newUnitType);
    if (newLocalUnit) setLocalUnit(newLocalUnit);
  }, [unitType, localUnit, singleMode]);

  // 渲染单位类型切换按钮组
  const renderUnitTypeToggle = () => {
    // 处理货币单位切换
    const handleUnitTypeChange = (_, newUnitType) => {
      if (newUnitType && newUnitType !== unitType) {
        handleUnitChange(newUnitType, null);
      }
    };

    // 处理数量单位切换
    const handleLocalUnitChange = (_, newUnit) => {
      if (newUnit && newUnit !== localUnit) {
        handleUnitChange(null, newUnit);
      }
    };

    return (
      <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
        <Stack direction="row" spacing={2}>
          <ToggleButtonGroup
            value={unitType}
            onChange={handleUnitTypeChange}
            options={unitTypeOptions}
            aria-label="unit toggle"
          />

          <ToggleButtonGroup
            value={localUnit}
            onChange={handleLocalUnitChange}
            options={unitOptions}
            aria-label="unit toggle"
          />
        </Stack>
      </FormControl>
    );
  };

  // 渲染输入价格表单
  const renderInputField = (formProps) => {
    const { errors = {}, touched = {}, handleBlur, handleChange: formikHandleChange, values = {} } = formProps || {};

    // 根据模式选择正确的处理函数和值
    const value = singleMode ? (inputs.input ?? 0) : (values.input ?? 0);
    const onChange = singleMode ? handleChange : formikHandleChange;
    const errorState = singleMode ? !!errors.input : Boolean(touched?.input && errors?.input);

    return (
      <FormControl fullWidth error={errorState} sx={{ ...theme.typography.otherInput }}>
        <InputLabel htmlFor="channel-input-label">{t('modelpricePage.inputMultiplier')}</InputLabel>
        <OutlinedInput
          id="channel-input-label"
          label={t('modelpricePage.inputMultiplier')}
          type="number"
          value={value}
          name="input"
          startAdornment={<InputAdornment position="start">{handleStartAdornment()}</InputAdornment>}
          endAdornment={<InputAdornment position="end">{handleEndAdornment(value)}</InputAdornment>}
          onBlur={handleBlur}
          onChange={onChange}
          aria-describedby="helper-text-channel-input-label"
          disabled={isConverting}
          sx={{
            '& .MuiOutlinedInput-root': {
              backgroundColor: isConverting ? 'rgba(25, 118, 210, 0.04)' : 'transparent',
              transition: 'background-color 0.2s ease-in-out'
            }
          }}
        />

        {(singleMode && errors.input) || (!singleMode && touched?.input && errors?.input) ? (
          <FormHelperText error id="helper-tex-channel-input-label">
            {errors.input}
          </FormHelperText>
        ) : null}
      </FormControl>
    );
  };

  // 渲染输出价格表单
  const renderOutputField = (formProps) => {
    const { errors = {}, touched = {}, handleBlur, handleChange: formikHandleChange, values = {} } = formProps || {};

    // 根据模式选择正确的处理函数和值
    const value = singleMode ? (inputs.output ?? 0) : (values.output ?? 0);
    const onChange = singleMode ? handleChange : formikHandleChange;
    const errorState = singleMode ? !!errors.output : Boolean(touched?.output && errors?.output);

    return (
      <FormControl fullWidth error={errorState} sx={{ ...theme.typography.otherInput }}>
        <InputLabel htmlFor="channel-output-label">{t('modelpricePage.outputMultiplier')}</InputLabel>
        <OutlinedInput
          id="channel-output-label"
          label={t('modelpricePage.outputMultiplier')}
          type="number"
          value={value}
          name="output"
          startAdornment={<InputAdornment position="start">{handleStartAdornment()}</InputAdornment>}
          endAdornment={<InputAdornment position="end">{handleEndAdornment(value)}</InputAdornment>}
          onBlur={handleBlur}
          onChange={onChange}
          aria-describedby="helper-text-channel-output-label"
          disabled={isConverting}
          sx={{
            '& .MuiOutlinedInput-root': {
              backgroundColor: isConverting ? 'rgba(25, 118, 210, 0.04)' : 'transparent',
              transition: 'background-color 0.2s ease-in-out'
            }
          }}
        />

        {(singleMode && errors.output) || (!singleMode && touched?.output && errors?.output) ? (
          <FormHelperText error id="helper-tex-channel-output-label">
            {errors.output}
          </FormHelperText>
        ) : null}
      </FormControl>
    );
  };

  // 渲染锁定切换按钮组
  const renderLockedToggle = (formProps) => {
    const { handleChange: formikHandleChange, values = {} } = formProps || {};

    // 在单模式下，我们需要使用不同的处理方式
    const handleLockChange = (_, newLocked) => {
      if (singleMode) {
        // 在单模式下，直接更新inputs状态
        handleChange({
          target: {
            name: 'locked',
            checked: newLocked
          }
        });
      } else {
        // 在多模式下，使用Formik的handleChange
        formikHandleChange({
          target: {
            name: 'locked',
            value: newLocked
          }
        });
      }
    };

    return (
      <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
        <Stack direction="row" spacing={2}>
          <ToggleButtonGroup
            value={singleMode ? inputs.locked : values?.locked}
            onChange={handleLockChange}
            options={lockedOptions}
            aria-label="locked toggle"
          />
        </Stack>
      </FormControl>
    );
  };

  // 渲染额外比率选择器
  const renderExtraRatioSelector = (formProps) => {
    const { setFieldValue, values = {} } = formProps || {};

    if (singleMode) {
      return (
        <Paper variant="outlined" sx={{ p: 2, mt: 2 }}>
          <ExtraRatiosSelector value={inputs.extra_ratios} onChange={handleChangeExtraRatios} />
        </Paper>
      );
    }

    return (
      <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
        <ExtraRatiosSelector
          value={values.extra_ratios || {}}
          onChange={(newExtraRatios) => {
            setFieldValue('extra_ratios', newExtraRatios);
          }}
          handleStartAdornment={handleStartAdornment}
        />
      </FormControl>
    );
  };

  // 渲染模型选择器 (多模式特有)
  const renderModelSelector = (formProps) => {
    if (!formProps) return null;

    const { errors = {}, handleBlur, handleChange, values = {} } = formProps;

    return (
      <FormControl fullWidth sx={{ ...theme.typography.otherInput }}>
        <Autocomplete
          multiple
          freeSolo
          id="channel-models-label"
          options={selectModel || []}
          value={values.models || []}
          onChange={(_, value) => {
            const event = {
              target: {
                name: 'models',
                value: value
              }
            };
            handleChange(event);
          }}
          onBlur={handleBlur}
          disableCloseOnSelect
          getOptionLabel={(option) => option || ''}
          renderInput={(params) => <TextField {...params} name="models" error={Boolean(errors.models)} label={t('pricing_edit.model')} />}
          filterOptions={(options, params) => {
            const filtered = filter(options, params);
            const { inputValue } = params;
            const isExisting = options.some((option) => inputValue === option);
            if (inputValue !== '' && !isExisting) {
              filtered.push(inputValue);
            }
            return filtered;
          }}
          renderOption={(props, option, { selected }) => (
            <li {...props}>
              <Checkbox icon={icon} checkedIcon={checkedIcon} style={{ marginRight: 8 }} checked={selected} />
              {option}
            </li>
          )}
        />
        {errors.models ? (
          <FormHelperText error id="helper-tex-channel-models-label">
            {errors.models}
          </FormHelperText>
        ) : (
          <FormHelperText id="helper-tex-channel-models-label"> {t('pricing_edit.modelTip')} </FormHelperText>
        )}
      </FormControl>
    );
  };

  // 渲染操作按钮
  const renderActions = (formProps) => {
    const { isSubmitting = false } = formProps || {};

    return (
      <DialogActions>
        <Button onClick={onCancel}>{t('common.cancel')}</Button>
        {singleMode ? (
          <Button
            onClick={() =>
              submit(inputs, {
                setErrors,
                setStatus: () => {},
                setSubmitting: () => {}
              })
            }
            variant="contained"
            color="primary"
          >
            {t('common.submit')}
          </Button>
        ) : (
          <Button disableElevation disabled={isSubmitting} type="submit" variant="contained" color="primary">
            {t('common.submit')}
          </Button>
        )}
      </DialogActions>
    );
  };

  return (
    <Dialog open={open} onClose={onCancel} fullWidth maxWidth="md">
      <DialogTitle sx={{ margin: '0px', fontWeight: 700, lineHeight: '1.55556', padding: '24px', fontSize: '1.125rem' }}>
        {singleMode ? (price ? t('common.edit') : t('common.create')) : pricesItem ? t('common.edit') : t('common.create')}
      </DialogTitle>
      <Divider />
      <DialogContent>
        {singleMode ? (
          // 单一模式表单
          <Stack spacing={2} sx={{ mt: 1 }}>
            <TextField
              label={t('pricing_edit.name')}
              name="model"
              value={inputs.model}
              onChange={handleChange}
              fullWidth
              error={!!errors.model}
              helperText={errors.model}
            />

            {renderTypeSelector()}
            {renderChannelTypeSelector()}
            {renderUnitTypeToggle()}

            <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
              {renderInputField()}
              {renderOutputField()}
            </Stack>

            {renderLockedToggle()}
            <Alert severity="warning">{t('pricing_edit.lockedTip')}</Alert>

            {isConverting && (
              <Alert severity="info" sx={{ mt: 1 }}>
                正在根据新单位自动调整价格...
              </Alert>
            )}

            {renderExtraRatioSelector()}

            {errors.general && (
              <Typography color="error" variant="body2">
                {errors.general}
              </Typography>
            )}

            {renderActions()}
          </Stack>
        ) : (
          // 多选模式表单
          <Formik initialValues={inputs} enableReinitialize validationSchema={getValidationSchema(t)} onSubmit={submit}>
            {(formProps) => (
              <form noValidate onSubmit={formProps.handleSubmit}>
                {renderTypeSelector(formProps)}
                {renderChannelTypeSelector(formProps)}
                {renderUnitTypeToggle()}
                {renderInputField(formProps)}
                {renderOutputField(formProps)}
                {renderModelSelector(formProps)}
                {renderLockedToggle(formProps)}
                <Alert severity="warning">{t('pricing_edit.lockedTip')}</Alert>
                {renderExtraRatioSelector(formProps)}
                {renderActions(formProps)}
              </form>
            )}
          </Formik>
        )}
      </DialogContent>
    </Dialog>
  );
};

export default EditModal;

EditModal.propTypes = {
  open: PropTypes.bool,
  pricesItem: PropTypes.oneOfType([PropTypes.object, PropTypes.any]),
  onCancel: PropTypes.func,
  onOk: PropTypes.func,
  ownedby: PropTypes.array,
  noPriceModel: PropTypes.array,
  // 以下是单一模式专用
  singleMode: PropTypes.bool,
  price: PropTypes.object,
  rows: PropTypes.array,
  onSaveSingle: PropTypes.func,
  unit: PropTypes.string
};
