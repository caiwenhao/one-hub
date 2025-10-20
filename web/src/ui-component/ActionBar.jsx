import PropTypes from 'prop-types';
import { Box, Stack } from '@mui/material';
import { useTheme } from '@mui/material/styles';

// 统一的操作区：延续 CardActions 节奏，兼顾响应式换行
export default function ActionBar({ children, justify = 'flex-end', sx }) {
  const theme = useTheme();
  return (
    <Box
      sx={{
        px: { xs: 1.5, md: 2 },
        py: { xs: 1.5, md: 1.75 },
        borderTop: `1px solid ${theme.palette.divider}`,
        backgroundColor: theme.palette.background.paper,
        ...sx
      }}
    >
      <Stack
        direction="row"
        spacing={{ xs: 1, md: 1.5 }}
        rowGap={1}
        alignItems="center"
        justifyContent={{ xs: 'flex-start', md: justify }}
        flexWrap="wrap"
      >
        {children}
      </Stack>
    </Box>
  );
}

ActionBar.propTypes = {
  children: PropTypes.node,
  justify: PropTypes.oneOf(['flex-start', 'center', 'flex-end', 'space-between', 'space-around']),
  sx: PropTypes.object
};
