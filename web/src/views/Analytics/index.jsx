import { gridSpacing } from 'store/constant';
import { Grid } from '@mui/material';
import MainCard from 'ui-component/cards/MainCard';
import PageHeader from 'ui-component/PageHeader';
import { useTranslation } from 'react-i18next';
import Statistics from './component/Statistics';
import Overview from './component/Overview';

export default function MarketingData() {
  const { t } = useTranslation();

  return (
    <Grid container spacing={gridSpacing}>
      <Grid item xs={12}>
        <PageHeader title={t('analytics')} subtitle="Analytics" />
      </Grid>
      <Grid item xs={12}>
        <Statistics />
      </Grid>
      <Grid item xs={12}>
        <MainCard>
          <Overview />
        </MainCard>
      </Grid>
    </Grid>
  );
}
