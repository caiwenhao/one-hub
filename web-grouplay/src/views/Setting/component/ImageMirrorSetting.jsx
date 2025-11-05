import { useEffect, useState } from 'react';
import SubCard from 'ui-component/cards/SubCard';
import { Stack, FormControl, InputLabel, OutlinedInput, Checkbox, Button, FormControlLabel, Alert, Grid } from '@mui/material';
import { API } from 'utils/api';
import { showError, showSuccess } from 'utils/common';
import { useTranslation } from 'react-i18next';

const ImageMirrorSetting = () => {
  const { t } = useTranslation();
  const [inputs, setInputs] = useState({
    NewAPIMirrorImageToStorage: '',
    NewAPIAllowedAssetHosts: '',
    S3Endpoint: '',
    S3CDNURL: '',
    S3BucketName: '',
    S3AccessKeyId: '',
    S3AccessKeySecret: '',
    S3ExpirationDays: 0
  });
  const [originInputs, setOriginInputs] = useState({});
  const [loading, setLoading] = useState(false);

  const getOptions = async () => {
    try {
      const res = await API.get('/api/option/');
      const { success, message, data } = res.data;
      if (success) {
        const newInputs = { ...inputs };
        data.forEach((item) => {
          switch (item.key) {
            case 'NewAPIMirrorImageToStorage':
            case 'NewAPIAllowedAssetHosts':
            case 'S3Endpoint':
            case 'S3CDNURL':
            case 'S3BucketName':
            case 'S3AccessKeyId':
            case 'S3AccessKeySecret':
            case 'S3ExpirationDays':
              newInputs[item.key] = item.value;
              break;
            default:
              break;
          }
        });
        setInputs(newInputs);
        setOriginInputs(newInputs);
      } else {
        showError(message);
      }
    } catch (e) {
      // ignore
    }
  };

  useEffect(() => {
    getOptions();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const updateOption = async (key, value) => {
    setLoading(true);
    if (key.endsWith('Enabled')) {
      value = inputs[key] === 'true' ? 'false' : 'true';
    }
    try {
      const res = await API.put('/api/option/', { key, value });
      const { success, message } = res.data;
      if (!success) showError(message);
    } catch (e) {
      showError(e.message || '请求失败');
    } finally {
      setLoading(false);
    }
  };

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setInputs((prev) => ({ ...prev, [name]: value }));
  };

  const handleToggle = async (event) => {
    const { name } = event.target;
    const next = inputs[name] === 'true' ? 'false' : 'true';
    setInputs((prev) => ({ ...prev, [name]: next }));
  };

  const handleSave = async () => {
    setLoading(true);
    try {
      const ops = [];
      const keys = [
        'NewAPIMirrorImageToStorage',
        'NewAPIAllowedAssetHosts',
        'S3Endpoint',
        'S3CDNURL',
        'S3BucketName',
        'S3AccessKeyId',
        'S3AccessKeySecret',
        'S3ExpirationDays'
      ];
      for (const k of keys) {
        if (originInputs[k] !== inputs[k]) {
          ops.push(updateOption(k, inputs[k]));
        }
      }
      await Promise.all(ops);
      await getOptions();
      showSuccess(t('common.saveSuccess', { defaultValue: '保存成功！' }));
    } catch (e) {
      showError(e.message || '保存失败');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Stack spacing={3}>
      <SubCard title={t('setting_index.imageMirrorSettings.mirrorTitle', { defaultValue: '镜像策略' })}>
        <Stack spacing={2}>
          <FormControlLabel
            sx={{ marginLeft: '0px' }}
            label={t('setting_index.operationSettings.otherSettings.imageMirror.enabled')}
            control={
              <Checkbox
                checked={inputs.NewAPIMirrorImageToStorage === 'true'}
                onChange={handleToggle}
                name="NewAPIMirrorImageToStorage"
              />
            }
          />
          <FormControl>
            <InputLabel htmlFor="NewAPIAllowedAssetHosts">
              {t('setting_index.operationSettings.otherSettings.imageMirror.allowedHostsLabel')}
            </InputLabel>
            <OutlinedInput
              id="NewAPIAllowedAssetHosts"
              name="NewAPIAllowedAssetHosts"
              value={inputs.NewAPIAllowedAssetHosts}
              onChange={handleInputChange}
              label={t('setting_index.operationSettings.otherSettings.imageMirror.allowedHostsLabel')}
              placeholder={t('setting_index.operationSettings.otherSettings.imageMirror.allowedHostsPlaceholder')}
              disabled={loading}
            />
          </FormControl>
        </Stack>
      </SubCard>

      <SubCard title={t('setting_index.imageMirrorSettings.storageTitle', { defaultValue: '对象存储（S3/R2）' })}>
        <Stack spacing={2}>
          <Alert severity="info">{t('setting_index.operationSettings.otherSettings.s3r2.title')}</Alert>
          <Grid container spacing={2}>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel htmlFor="S3Endpoint">{t('setting_index.operationSettings.otherSettings.s3r2.endpoint')}</InputLabel>
                <OutlinedInput id="S3Endpoint" name="S3Endpoint" value={inputs.S3Endpoint} onChange={handleInputChange} label={t('setting_index.operationSettings.otherSettings.s3r2.endpoint')} disabled={loading} />
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel htmlFor="S3CDNURL">{t('setting_index.operationSettings.otherSettings.s3r2.cdnurl')}</InputLabel>
                <OutlinedInput id="S3CDNURL" name="S3CDNURL" value={inputs.S3CDNURL} onChange={handleInputChange} label={t('setting_index.operationSettings.otherSettings.s3r2.cdnurl')} disabled={loading} />
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel htmlFor="S3BucketName">{t('setting_index.operationSettings.otherSettings.s3r2.bucketName')}</InputLabel>
                <OutlinedInput id="S3BucketName" name="S3BucketName" value={inputs.S3BucketName} onChange={handleInputChange} label={t('setting_index.operationSettings.otherSettings.s3r2.bucketName')} disabled={loading} />
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel htmlFor="S3ExpirationDays">{t('setting_index.operationSettings.otherSettings.s3r2.expirationDays')}</InputLabel>
                <OutlinedInput id="S3ExpirationDays" name="S3ExpirationDays" value={inputs.S3ExpirationDays} onChange={handleInputChange} label={t('setting_index.operationSettings.otherSettings.s3r2.expirationDays')} disabled={loading} />
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel htmlFor="S3AccessKeyId">{t('setting_index.operationSettings.otherSettings.s3r2.accessKeyId')}</InputLabel>
                <OutlinedInput id="S3AccessKeyId" name="S3AccessKeyId" value={inputs.S3AccessKeyId} onChange={handleInputChange} label={t('setting_index.operationSettings.otherSettings.s3r2.accessKeyId')} disabled={loading} />
              </FormControl>
            </Grid>
            <Grid item xs={12} md={6}>
              <FormControl fullWidth>
                <InputLabel htmlFor="S3AccessKeySecret">{t('setting_index.operationSettings.otherSettings.s3r2.accessKeySecret')}</InputLabel>
                <OutlinedInput id="S3AccessKeySecret" name="S3AccessKeySecret" value={inputs.S3AccessKeySecret} onChange={handleInputChange} label={t('setting_index.operationSettings.otherSettings.s3r2.accessKeySecret')} disabled={loading} />
              </FormControl>
            </Grid>
          </Grid>
          <Button variant="contained" onClick={handleSave} disabled={loading}>
            {t('setting_index.imageMirrorSettings.saveButton', { defaultValue: '保存镜像与存储设置' })}
          </Button>
        </Stack>
      </SubCard>
    </Stack>
  );
};

export default ImageMirrorSetting;

