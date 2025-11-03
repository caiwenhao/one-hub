export default function paperCard(theme) {
  return {
    MuiPaper: {
      defaultProps: { elevation: 0 },
      styleOverrides: {
        root: {
          backgroundImage: 'none',
          borderRadius: 'var(--radii-lg)',
          background: theme.paper,
          border: `1px solid ${theme.divider}`,
          boxShadow: theme.mode === 'dark' ? 'none' : 'var(--shadow-md)'
        },
        rounded: { borderRadius: `${theme?.customization?.borderRadius || 12}px` },
        elevation1: { boxShadow: theme.mode === 'dark' ? 'none' : 'var(--shadow-sm)' },
        elevation2: { boxShadow: theme.mode === 'dark' ? 'none' : 'var(--shadow-lg)' }
      }
    },
    MuiCard: {
      styleOverrides: {
        root: {
          borderRadius: 'var(--radii-lg)',
          padding: 0,
          boxShadow: theme.mode === 'dark' ? 'none' : 'var(--shadow-md)',
          transition: 'box-shadow var(--motion-duration-normal) var(--motion-easing-standard)',
          overflow: 'hidden',
          '&:hover': {
            boxShadow: theme.mode === 'dark' ? 'none' : 'var(--shadow-lg)'
          },
          '& .MuiTableContainer-root': { borderRadius: 0 }
        }
      }
    }
  };
}
