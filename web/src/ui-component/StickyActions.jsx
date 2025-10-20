import PropTypes from 'prop-types';
import { Box, Stack } from '@mui/material';

// 吸顶/吸底操作栏（随滚动始终可见）
export default function StickyActions({ children, position = 'bottom', sx }) {
  const isBottom = position === 'bottom';
  return (
    <Box
      sx={{
        position: 'sticky',
        [isBottom ? 'bottom' : 'top']: 0,
        zIndex: (theme) => theme.zIndex.appBar,
        bgcolor: 'background.paper',
        borderTop: isBottom ? (theme) => `1px solid ${theme.palette.divider}` : 'none',
        borderBottom: !isBottom ? (theme) => `1px solid ${theme.palette.divider}` : 'none',
        px: { xs: 1.5, md: 2 },
        py: 1,
        boxShadow: (theme) => (theme.palette.mode === 'dark' ? 'none' : '0 -4px 16px rgba(0,0,0,0.04)'),
        ...sx
      }}
    >
      <Stack direction="row" justifyContent="flex-end" alignItems="center" spacing={1}>
        {children}
      </Stack>
    </Box>
  );
}

StickyActions.propTypes = {
  children: PropTypes.node,
  position: PropTypes.oneOf(['top', 'bottom']),
  sx: PropTypes.object
};

