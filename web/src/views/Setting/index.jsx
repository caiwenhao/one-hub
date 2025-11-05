import { useState, useEffect, useMemo } from 'react';
import PropTypes from 'prop-types';
import { Tabs, Tab, Box, Card, CardHeader, CardContent, Divider, Stack } from '@mui/material';
import { IconWorldCog, IconCpu, IconServerCog } from '@tabler/icons-react';
import OperationSetting from './component/OperationSetting';
import ImageMirrorSetting from './component/ImageMirrorSetting';
import SystemSetting from './component/SystemSetting';
import OtherSetting from './component/OtherSetting';
import { useLocation, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import PageHeader from 'ui-component/PageHeader';

function CustomTabPanel({ children, value, index }) {
  return (
    <Box
      role="tabpanel"
      hidden={value !== index}
      id={`setting-tabpanel-${index}`}
      aria-labelledby={`setting-tab-${index}`}
      sx={{ width: '100%' }}
    >
      {value === index && <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>{children}</Box>}
    </Box>
  );
}

CustomTabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.number.isRequired,
  value: PropTypes.number.isRequired
};

function a11yProps(index) {
  return {
    id: `setting-tab-${index}`,
    'aria-controls': `setting-tabpanel-${index}`
  };
}

const Setting = () => {
  const { t } = useTranslation();
  const location = useLocation();
  const navigate = useNavigate();
  const hash = location.hash.replace('#', '');
  const tabMap = useMemo(
    () => ({
      operation: 0,
      system: 1,
      mirror: 2,
      other: 3
    }),
    []
  );
  const [value, setValue] = useState(tabMap[hash] || 0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
    const hashArray = Object.keys(tabMap);
    navigate(`#${hashArray[newValue]}`);
  };

  useEffect(() => {
    const handleHashChange = () => {
      const hash = location.hash.replace('#', '');
      setValue(tabMap[hash] || 0);
    };
    window.addEventListener('hashchange', handleHashChange);
    return () => {
      window.removeEventListener('hashchange', handleHashChange);
    };
  }, [location, tabMap]);

  const headerTitle = t('setting_index.title', { defaultValue: '系统设置' });
  const headerDescription = t('setting_index.description', {
    defaultValue: '管理站点运行参数、系统行为与扩展配置。'
  });

  return (
    <Stack spacing={3} sx={{ width: '100%' }}>
      <PageHeader title={headerTitle} description={headerDescription} />

      <Card elevation={0} sx={{ borderRadius: 2, border: (theme) => `1px solid ${theme.palette.divider}` }}>
        <CardHeader
          title={
            <Tabs
              value={value}
              onChange={handleChange}
              variant="scrollable"
              scrollButtons="auto"
              sx={{
                '& .MuiTab-root': {
                  minHeight: 44
                }
              }}
            >
              <Tab
                label={t('setting_index.operationSettings.title')}
                {...a11yProps(0)}
                icon={<IconCpu size={18} />}
                iconPosition="start"
              />
              <Tab
                label={t('setting_index.systemSettings.title')}
                {...a11yProps(1)}
                icon={<IconServerCog size={18} />}
                iconPosition="start"
              />
              <Tab
                label={t('setting_index.imageMirrorSettings.title')}
                {...a11yProps(2)}
                icon={<IconWorldCog size={18} />}
                iconPosition="start"
              />
              <Tab
                label={t('setting_index.otherSettings.title')}
                {...a11yProps(3)}
                icon={<IconWorldCog size={18} />}
                iconPosition="start"
              />
            </Tabs>
          }
          sx={{
            px: { xs: 2, md: 3 },
            pt: { xs: 1.5, md: 2 },
            pb: 0,
            '& .MuiCardHeader-title': {
              width: '100%'
            }
          }}
        />

        <Divider sx={{ mt: { xs: 1, md: 1.5 } }} />

        <CardContent sx={{ px: { xs: 2, md: 3 }, py: { xs: 2, md: 3 } }}>
          <CustomTabPanel value={value} index={0}>
            <OperationSetting />
          </CustomTabPanel>
          <CustomTabPanel value={value} index={1}>
            <SystemSetting />
          </CustomTabPanel>
          <CustomTabPanel value={value} index={2}>
            <ImageMirrorSetting />
          </CustomTabPanel>
          <CustomTabPanel value={value} index={3}>
            <OtherSetting />
          </CustomTabPanel>
        </CardContent>
      </Card>
    </Stack>
  );
};

export default Setting;
