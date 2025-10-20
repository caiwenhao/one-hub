import PropTypes from 'prop-types';
import { Box } from '@mui/material';
import { alpha, useTheme } from '@mui/material/styles';

// 统一的筛选区包裹组件：
// - 背景、边框、圆角与间距与主题对齐（企业稳重）
// - 接收任意子元素以保持灵活性
export default function FilterBar({ children, sx }) {
  const theme = useTheme();
  const subtleBg = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.12 : 0.06);
  return (
    <Box
      sx={{
        position: 'relative',
        bgcolor: subtleBg,
        border: `1px solid ${alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.24 : 0.12)}`,
        borderRadius: `${theme.shape.borderRadius}px`,
        p: { xs: 1.5, md: 2 },
        mb: 2,
        '&::before': {
          content: '""',
          position: 'absolute',
          inset: 0,
          borderRadius: 'inherit',
          border: `1px solid ${alpha(theme.palette.common.white, theme.palette.mode === 'dark' ? 0.06 : 0.1)}`,
          pointerEvents: 'none',
          mixBlendMode: theme.palette.mode === 'dark' ? 'screen' : 'normal'
        },
        ...sx
      }}
    >
      {children}
    </Box>
  );
}

FilterBar.propTypes = {
  children: PropTypes.node,
  sx: PropTypes.object
};
