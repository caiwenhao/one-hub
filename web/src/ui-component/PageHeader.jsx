import PropTypes from 'prop-types';
import { Box, Stack, Typography } from '@mui/material';

// 统一的后台页面页头：标题、副标题、描述与操作按钮集中管理
export default function PageHeader({ title, subtitle, description, actions, sx }) {
  const actionItems = Array.isArray(actions) ? actions.filter(Boolean) : actions ? [actions] : [];

  return (
    <Stack
      direction={{ xs: 'column', md: 'row' }}
      alignItems={{ xs: 'flex-start', md: 'center' }}
      justifyContent="space-between"
      spacing={{ xs: 1.5, md: 2 }}
      sx={{ mb: { xs: 2, md: 3 }, ...sx }}
    >
      <Stack spacing={0.75} flex={1} minWidth={0}>
        <Typography variant="h2" noWrap={false} sx={{ wordBreak: 'break-word' }}>
          {title}
        </Typography>
        {subtitle && (
          <Typography variant="subtitle1" color="text.secondary">
            {subtitle}
          </Typography>
        )}
        {description && (
          <Typography variant="body2" color="text.secondary">
            {description}
          </Typography>
        )}
      </Stack>

      {actionItems.length > 0 && (
        <Stack
          direction="row"
          spacing={1}
          alignItems="center"
          justifyContent="flex-end"
          flexWrap="wrap"
        >
          {actionItems.map((action, index) => (
            <Box key={index} sx={{ display: 'inline-flex' }}>
              {action}
            </Box>
          ))}
        </Stack>
      )}
    </Stack>
  );
}

PageHeader.propTypes = {
  title: PropTypes.node.isRequired,
  subtitle: PropTypes.node,
  description: PropTypes.node,
  actions: PropTypes.oneOfType([PropTypes.node, PropTypes.arrayOf(PropTypes.node)]),
  sx: PropTypes.object
};
