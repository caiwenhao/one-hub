import { Grid } from '@mui/material';
import { gridSpacing } from 'store/constant';
import PageHeader from 'ui-component/PageHeader';
import { useTranslation } from 'react-i18next';

// Import components
import SystemLogs from './components/SystemLogs';

// Main SystemInfo Component
const SystemInfo = () => {
  const { t } = useTranslation();
  return (
    <Grid container spacing={gridSpacing}>
      <Grid item xs={12}>
        <PageHeader title={t('systemInfo')} subtitle="System Info" />
      </Grid>
      {/* 系统日志卡片 */}
      <Grid item xs={12}>
        <SystemLogs />
      </Grid>
    </Grid>
  );
};

export default SystemInfo;
