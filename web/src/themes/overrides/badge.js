export default function badge(theme) {
  return {
    MuiBadge: {
      defaultProps: { max: 99 },
      styleOverrides: {
        badge: {
          minWidth: 18,
          height: 18,
          padding: '0 5px',
          borderRadius: 9,
          fontSize: '0.75rem',
          boxShadow: theme.mode === 'dark' ? '0 0 0 1px rgba(0,0,0,0.4)' : '0 0 0 1px rgba(255,255,255,0.8)'
        },
        colorPrimary: { backgroundColor: theme.colors?.primaryMain },
        colorSecondary: { backgroundColor: theme.colors?.secondaryMain },
        colorError: { backgroundColor: theme.colors?.errorMain }
      }
    }
  };
}
