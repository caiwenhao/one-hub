import PropTypes from 'prop-types';
import { Box, Stack, Typography } from '@mui/material';
import { alpha, useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';

// 标准化的空状态展示组件，用于列表/内容缺省提示
const sizeConfig = {
  small: {
    iconSize: 32,
    spacing: 1.5,
    padding: 2
  },
  medium: {
    iconSize: 48,
    spacing: 2,
    padding: 3
  },
  large: {
    iconSize: 64,
    spacing: 2.5,
    padding: 4
  }
};

const defaultIcon = <Icon icon="solar:folder-with-files-linear" />;

export default function EmptyState({
  icon = defaultIcon,
  title = '暂无内容',
  description,
  action,
  size = 'medium',
  sx
}) {
  const theme = useTheme();
  const preset = sizeConfig[size] || sizeConfig.medium;

  return (
    <Stack
      spacing={preset.spacing}
      alignItems="center"
      justifyContent="center"
      textAlign="center"
      sx={{
        px: preset.padding,
        py: preset.padding * 1.5,
        color: theme.palette.text.secondary,
        backgroundColor: alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.08 : 0.04),
        borderRadius: 2,
        border: `1px dashed ${alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.2 : 0.12)}`,
        ...sx
      }}
    >
      <Box
        sx={{
          width: preset.iconSize,
          height: preset.iconSize,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: theme.palette.primary.main,
          '& svg': {
            width: preset.iconSize,
            height: preset.iconSize
          }
        }}
      >
        {icon}
      </Box>

      <Stack spacing={0.75} alignItems="center" justifyContent="center">
        <Typography variant="subtitle1" color="text.primary" sx={{ fontWeight: 600 }}>
          {title}
        </Typography>
        {description && (
          <Typography variant="body2" color="text.secondary" sx={{ maxWidth: 360 }}>
            {description}
          </Typography>
        )}
      </Stack>

      {action && <Box sx={{ mt: preset.spacing }}>{action}</Box>}
    </Stack>
  );
}

EmptyState.propTypes = {
  icon: PropTypes.node,
  title: PropTypes.node,
  description: PropTypes.node,
  action: PropTypes.node,
  size: PropTypes.oneOf(['small', 'medium', 'large']),
  sx: PropTypes.object
};
