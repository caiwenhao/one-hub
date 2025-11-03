export default function loadingButton(theme) {
  return {
    MuiLoadingButton: {
      defaultProps: {
        loadingPosition: 'start'
      },
      styleOverrides: {
        root: {
          borderRadius: 'var(--radii-md)',
          padding: 'var(--density-button-py) var(--density-button-px)'
        },
        loadingIndicator: {
          color: 'var(--color-text-inverse)'
        },
        containedPrimary: {
          '&.MuiLoadingButton-loading': {
            background: 'var(--color-brand-primary)',
            opacity: 0.9
          }
        }
      }
    }
  };
}
