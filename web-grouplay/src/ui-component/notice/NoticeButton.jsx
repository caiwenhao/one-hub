import PropTypes from 'prop-types';
import { ButtonBase } from '@mui/material';
import { alpha, useTheme } from '@mui/material/styles';

import { Icon } from '@iconify/react';

import { useNotice } from './NoticeContext';

export function NoticeButton({ sx }) {
  const theme = useTheme();
  const { openNotice } = useNotice();

  const surfaceColor = theme.palette.mode === 'dark' ? alpha(theme.palette.background.paper, 0.7) : theme.palette.background.paper;
  const borderColor = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.24 : 0.14);

  return (
    <ButtonBase
      aria-label="打开通知"
      onClick={openNotice}
      sx={{
        width: 40,
        height: 40,
        borderRadius: '12px',
        border: `1px solid ${borderColor}`,
        backgroundColor: surfaceColor,
        color: theme.palette.mode === 'dark' ? theme.palette.primary.light : theme.palette.primary.main,
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
      <Icon icon="lets-icons:message-duotone" width="20" />
    </ButtonBase>
  );
}

NoticeButton.propTypes = {
  sx: PropTypes.object
};
