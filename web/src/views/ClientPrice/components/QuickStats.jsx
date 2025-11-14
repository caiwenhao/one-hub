import { Box, Card, Typography, Stack } from '@mui/material';
import { Icon } from '@iconify/react';

const StatCard = ({ icon, label, value, color = 'primary' }) => (
  <Card
    sx={{
      p: 2,
      display: 'flex',
      alignItems: 'center',
      gap: 2,
      flex: 1,
      minWidth: 200
    }}
  >
    <Box
      sx={{
        width: 48,
        height: 48,
        borderRadius: 2,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        bgcolor: `${color}.lighter`,
        color: `${color}.main`
      }}
    >
      <Icon icon={icon} width={28} />
    </Box>
    <Box>
      <Typography variant="h4" fontWeight={700}>
        {value}
      </Typography>
      <Typography variant="body2" color="text.secondary">
        {label}
      </Typography>
    </Box>
  </Card>
);

const QuickStats = ({ configuredPrices }) => {
  if (!configuredPrices || configuredPrices.length === 0) {
    return null;
  }

  const totalModels = configuredPrices.length;
  const totalGroups = configuredPrices.reduce(
    (sum, item) => sum + item.groups.filter((g) => g.has_customer_price).length,
    0
  );
  const secondsBillingCount = configuredPrices.reduce(
    (sum, item) => sum + item.groups.filter((g) => g.has_customer_price && g.type === 'seconds').length,
    0
  );
  const timesBillingCount = configuredPrices.reduce(
    (sum, item) => sum + item.groups.filter((g) => g.has_customer_price && g.type === 'times').length,
    0
  );

  return (
    <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2} sx={{ mb: 3 }}>
      <StatCard
        icon="solar:cpu-bolt-bold-duotone"
        label="已配置模型"
        value={totalModels}
        color="primary"
      />
      <StatCard
        icon="solar:layers-bold-duotone"
        label="已配置分组"
        value={totalGroups}
        color="success"
      />
      {secondsBillingCount > 0 && (
        <StatCard
          icon="solar:clock-circle-bold-duotone"
          label="按秒计费"
          value={secondsBillingCount}
          color="warning"
        />
      )}
      {timesBillingCount > 0 && (
        <StatCard
          icon="solar:hashtag-circle-bold-duotone"
          label="按次计费"
          value={timesBillingCount}
          color="info"
        />
      )}
    </Stack>
  );
};

export default QuickStats;
