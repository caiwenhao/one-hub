import PropTypes from 'prop-types';
import { Icon } from '@iconify/react';

import Stack from '@mui/material/Stack';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputAdornment from '@mui/material/InputAdornment';

import { useTheme } from '@mui/material/styles';

// ----------------------------------------------------------------------

export default function TableToolBar({ placeholder }) {
  const theme = useTheme();
  const grey500 = theme.palette.grey[500];

  return (
    <Stack direction="row" alignItems="center" spacing={1.5} sx={{ width: '100%' }}>
      <OutlinedInput
        id="keyword"
        name="keyword"
        sx={{
          flex: 1,
          minWidth: 0,
          backgroundColor: theme.palette.background.paper,
          borderRadius: `${theme.shape.borderRadius}px`,
          '& .MuiInputAdornment-root:hover': {
            '& .search-icon': {
              borderRadius: '50%',
              boxShadow: '0 0 8px rgba(0,0,0,0.1)',
              transition: 'all 0.3s ease'
            }
          }
        }}
        placeholder={placeholder}
        size="small"
        startAdornment={
          <InputAdornment position="start">
            <Icon icon="solar:minimalistic-magnifer-line-duotone" className="search-icon" width="20" height="20" color={grey500} />
          </InputAdornment>
        }
      />
    </Stack>
  );
}

TableToolBar.propTypes = {
  filterName: PropTypes.string,
  handleFilterName: PropTypes.func,
  placeholder: PropTypes.string
};
