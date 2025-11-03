import PropTypes from 'prop-types';
import { Box } from '@mui/material';

// 统一滚动容器，替代 react-perfect-scrollbar
export default function ScrollArea({ children, sx, disablePadding = false, ...rest }) {
  return (
    <Box
      sx={{
        height: '100%',
        overflowY: 'auto',
        WebkitOverflowScrolling: 'touch',
        px: disablePadding ? 0 : { xs: 1.5, md: 2 },
        pb: disablePadding ? 0 : 2,
        ...sx
      }}
      {...rest}
    >
      {children}
    </Box>
  );
}

ScrollArea.propTypes = {
  children: PropTypes.node,
  sx: PropTypes.object,
  disablePadding: PropTypes.bool
};
