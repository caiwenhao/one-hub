export default function dialog(theme) {
  return {
    MuiDialog: {
      styleOverrides: {
        paper: {
          borderRadius: '12px',
          boxShadow: theme.mode === 'dark' ? '0 8px 24px rgba(0,0,0,0.5)' : '0 8px 24px rgba(0,0,0,0.12)',
          overflow: 'visible',
          '&.MuiPaper-rounded': { borderRadius: '12px' }
        },
        paperWidthXs: { maxWidth: '360px' },
        paperWidthSm: { maxWidth: '480px' },
        paperWidthMd: { maxWidth: '640px' },
        paperWidthLg: { maxWidth: '840px' },
        paperWidthXl: { maxWidth: '1040px' }
      }
    },
    MuiDialogTitle: {
      styleOverrides: {
        root: {
          position: 'sticky',
          top: 0,
          background: theme.paper,
          zIndex: 1,
          borderBottom: `1px solid ${theme.divider}`
        }
      }
    },
    MuiDialogActions: {
      styleOverrides: {
        root: {
          position: 'sticky',
          bottom: 0,
          background: theme.paper,
          zIndex: 1,
          borderTop: `1px solid ${theme.divider}`
        }
      }
    }
  };
}

