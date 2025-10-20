export default function button(theme) {
  return {
    MuiButton: {
      styleOverrides: {
        root: {
          fontWeight: 600,
          borderRadius: '10px',
          textTransform: 'none',
          boxShadow: 'none',
          minHeight: '36px',
          padding: '6px 16px',
          letterSpacing: '0.01em',
          transition: 'background-color 0.2s ease, box-shadow 0.2s ease',
          '&.Mui-disabled': {
            color: theme.colors?.grey600
          },
          '&:hover': {
            boxShadow: 'none'
          }
        },
        containedPrimary: {
          background: theme.colors?.primaryMain,
          color: '#fff',
          '&:hover': {
            background: theme.colors?.primaryDark
          }
        },
        outlinedPrimary: {
          borderColor: theme.colors?.primaryMain,
          color: theme.colors?.primaryMain,
          '&:hover': {
            backgroundColor: 'rgba(14, 165, 255, 0.06)'
          }
        },
        text: {
          '&:hover': {
            backgroundColor: 'rgba(0, 0, 0, 0.04)'
          }
        },
        sizeSmall: {
          padding: '6px 16px',
          fontSize: '0.8125rem',
          minHeight: '32px',
          borderRadius: '8px'
        },
        sizeLarge: {
          padding: '10px 24px',
          fontSize: '1rem',
          minHeight: '48px'
        }
      }
    }
  };
}

