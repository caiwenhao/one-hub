import PropTypes from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import { SET_THEME } from 'store/actions';
import { alpha, useTheme } from '@mui/material/styles';
import { ButtonBase } from '@mui/material';
import { Icon } from '@iconify/react';

export default function ThemeButton({ sx }) {
  const dispatch = useDispatch();
  const defaultTheme = useSelector((state) => state.customization.theme);
  const theme = useTheme();

  const isLight = defaultTheme === 'light';
  const handler = () => {
    const nextTheme = isLight ? 'dark' : 'light';
    dispatch({ type: SET_THEME, theme: nextTheme });
    localStorage.setItem('theme', nextTheme);
  };

  return (
    <ButtonBase
      aria-label={isLight ? '切换至深色主题' : '切换至浅色主题'}
      onClick={handler}
      sx={{
        width: 40,
        height: 40,
        borderRadius: '12px',
        border: `1px solid ${alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.32 : 0.18)}`,
        backgroundColor: isLight ? alpha(theme.palette.primary.main, 0.08) : alpha(theme.palette.primary.main, 0.16),
        color: isLight ? theme.palette.primary.main : theme.palette.primary.light,
        display: 'inline-flex',
        alignItems: 'center',
        justifyContent: 'center',
        transition: 'all .2s ease',
        '&:hover': {
          boxShadow: theme.shadows[4],
          transform: 'translateY(-1px)'
        },
        ...sx
      }}
    >
      <Icon icon={isLight ? 'solar:sun-2-bold-duotone' : 'solar:moon-bold-duotone'} width="20" />
    </ButtonBase>
  );
}

ThemeButton.propTypes = {
  sx: PropTypes.object
};
