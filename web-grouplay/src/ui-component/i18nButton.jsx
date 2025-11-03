import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { alpha, useTheme } from '@mui/material/styles';
import { Box, ButtonBase, Menu, MenuItem, Typography } from '@mui/material';
import i18nList from 'i18n/i18nList';
import useI18n from 'hooks/useI18n';
import Flags from 'country-flag-icons/react/3x2';

export default function I18nButton({ sx }) {
  const theme = useTheme();
  const i18n = useI18n();

  const [anchorEl, setAnchorEl] = useState(null);

  const handleMenuOpen = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleLanguageChange = (lng) => {
    i18n.changeLanguage(lng);
    handleMenuClose();
  };

  // 获取当前语言的国家代码
  const getCurrentCountryCode = () => {
    const currentLang = i18n.language || 'zh_CN';
    const langItem = i18nList.find((item) => item.lng === currentLang) || i18nList[0];
    return langItem.countryCode;
  };

  // 动态获取当前语言的国旗组件
  const CurrentFlag = Flags[getCurrentCountryCode()];

  return (
    <>
      <ButtonBase
        aria-label="切换语言"
        onClick={handleMenuOpen}
        sx={{
          width: 40,
          height: 40,
          borderRadius: '12px',
          border: `1px solid ${alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.24 : 0.14)}`,
          backgroundColor: theme.palette.mode === 'dark' ? alpha(theme.palette.background.paper, 0.7) : theme.palette.background.paper,
          display: 'inline-flex',
          alignItems: 'center',
          justifyContent: 'center',
          overflow: 'hidden',
          transition: 'all .2s ease',
          '&:hover': {
            boxShadow: theme.shadows[4],
            transform: 'translateY(-1px)'
          },
          ...sx
        }}
      >
        {CurrentFlag && (
          <Box
            sx={{
              width: '20px',
              height: '14px',
              overflow: 'hidden',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              borderRadius: '6px',
              boxShadow: `0 0 0 1px ${alpha(theme.palette.common.black, 0.08)}`
            }}
          >
            <CurrentFlag style={{ width: '100%', height: '100%', objectFit: 'cover', display: 'block' }} />
          </Box>
        )}
      </ButtonBase>
      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleMenuClose}
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'center'
        }}
        transformOrigin={{
          vertical: 'top',
          horizontal: 'center'
        }}
      >
        {i18nList.map((item) => {
          const FlagComponent = Flags[item.countryCode];
          return (
            <MenuItem
              key={item.lng}
              onClick={() => handleLanguageChange(item.lng)}
              sx={{
                display: 'flex',
                alignItems: 'center',
                gap: 1
              }}
            >
              {FlagComponent && (
                <Box
                  sx={{
                    width: '1.45rem',
                    height: '1.125rem',
                    overflow: 'hidden',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    position: 'relative'
                  }}
                >
                  <FlagComponent style={{ width: '100%', height: '100%', objectFit: 'cover', display: 'block' }} />
                </Box>
              )}
              <Typography variant="body1">{item.name}</Typography>
            </MenuItem>
          );
        })}
      </Menu>
    </>
  );
}

I18nButton.propTypes = {
  sx: PropTypes.object
};
