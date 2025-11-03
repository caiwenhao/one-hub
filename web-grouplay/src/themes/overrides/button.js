export default function button(theme) {
  return {
    MuiButton: {
      styleOverrides: {
        root: {
          fontWeight: 600,
          // 使用 Design Tokens 导出的 CSS 变量
          borderRadius: 'var(--radii-md)',
          textTransform: 'none',
          boxShadow: 'none',
          minHeight: '36px',
          padding: 'var(--density-button-py) var(--density-button-px)',
          letterSpacing: '0.01em',
          transition: 'background-color var(--motion-duration-normal) var(--motion-easing-standard), box-shadow var(--motion-duration-normal) var(--motion-easing-standard)',
          '&.Mui-disabled': {
            color: theme.colors?.grey600
          },
          '&:hover': {
            boxShadow: 'none'
          }
        },
        containedPrimary: {
          background: 'var(--color-brand-primary)',
          color: 'var(--color-text-inverse)',
          '&:hover': {
            background: 'var(--color-brand-primary-dark)'
          }
        },
        outlinedPrimary: {
          borderColor: 'var(--color-brand-primary)',
          color: 'var(--color-brand-primary)',
          '&:hover': {
            backgroundColor: 'rgba(22, 119, 255, 0.06)'
          }
        },
        containedSecondary: {
          background: 'var(--color-brand-secondary)',
          color: 'var(--color-text-inverse)',
          '&:hover': { filter: 'brightness(0.9)' }
        },
        containedError: {
          background: 'var(--color-semantic-error)',
          color: 'var(--color-text-inverse)',
          '&:hover': { filter: 'brightness(0.92)' }
        },
        outlinedError: {
          borderColor: 'var(--color-semantic-error)',
          color: 'var(--color-semantic-error)',
          '&:hover': { backgroundColor: 'rgba(239, 68, 68, 0.06)' }
        },
        text: {
          '&:hover': {
            backgroundColor: 'rgba(0, 0, 0, 0.04)'
          }
        },
        sizeSmall: {
          padding: 'var(--density-button-py) var(--density-button-px)',
          fontSize: '0.8125rem',
          minHeight: '32px',
          borderRadius: 'var(--radii-sm)'
        },
        sizeLarge: {
          padding: 'calc(var(--density-button-py) + 4px) calc(var(--density-button-px) + 4px)',
          fontSize: '1rem',
          minHeight: '48px'
        }
      }
    }
  };
}
