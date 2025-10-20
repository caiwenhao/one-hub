export default function paperCard(theme) {
  return {
    MuiPaper: {
      defaultProps: { elevation: 0 },
      styleOverrides: {
        root: {
          backgroundImage: 'none',
          borderRadius: '12px',
          background: theme.paper,
          border: `1px solid ${theme.divider}`,
          boxShadow: theme.mode === 'dark' ? 'none' : '0 2px 8px rgba(0,0,0,0.06)'
        },
        rounded: { borderRadius: `${theme?.customization?.borderRadius || 12}px` },
        elevation1: { boxShadow: theme.mode === 'dark' ? 'none' : '0 1px 3px rgba(0,0,0,0.08)' },
        elevation2: { boxShadow: theme.mode === 'dark' ? 'none' : '0 3px 12px rgba(0,0,0,0.10)' }
      }
    },
    MuiCard: {
      styleOverrides: {
        root: {
          borderRadius: '12px',
          padding: 0,
          boxShadow: theme.mode === 'dark' ? 'none' : '0 2px 8px rgba(0,0,0,0.06)',
          transition: 'box-shadow 0.2s ease',
          overflow: 'hidden',
          '&:hover': {
            boxShadow: theme.mode === 'dark' ? 'none' : '0 6px 16px rgba(0,0,0,0.10)'
          },
          '& .MuiTableContainer-root': { borderRadius: 0 }
        }
      }
    }
  };
}

