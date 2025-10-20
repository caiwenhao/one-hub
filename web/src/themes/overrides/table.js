export default function table(theme) {
  return {
    MuiTableContainer: {
      styleOverrides: {
        root: {
          overflowX: 'auto',
          overflowY: 'auto',
          borderRadius: 'var(--radii-lg)',
          boxShadow: 'none',
          ...(theme.breakpoints && {
            [theme.breakpoints.down('sm')]: { maxWidth: '100%', whiteSpace: 'nowrap' }
          })
        }
      }
    },
    MuiTable: {
      styleOverrides: { root: { borderCollapse: 'separate', borderSpacing: 0, width: '100%', margin: 0, padding: 0 } }
    },
    MuiTableHead: {
      styleOverrides: {
        root: {
          backgroundColor: theme.headBackgroundColor,
          borderBottom: `1px solid ${theme.tableBorderBottom}`,
          width: '100%',
          margin: 0,
          '& tr': {
            width: '100%',
            '& th:first-of-type': { borderTopLeftRadius: 0 },
            '& th:last-of-type': { borderTopRightRadius: 0 }
          }
        }
      }
    },
    MuiTableCell: {
      styleOverrides: {
        root: {
          borderBottom: `1px solid ${theme.tableBorderBottom}`,
          fontSize: '0.875rem',
          padding: 'var(--density-table-cell-py) var(--density-table-cell-px)',
          textAlign: 'center',
          '&:first-of-type': { paddingLeft: '12px' },
          '&:last-of-type': { paddingRight: '12px' }
        },
        head: {
          fontSize: '0.875rem',
          fontWeight: 600,
          color: theme.darkTextSecondary,
          borderBottom: 'none',
          whiteSpace: 'nowrap',
          padding: 'var(--density-table-head-py) var(--density-table-cell-px)',
          textAlign: 'center'
        },
        body: { color: theme.textDark }
      }
    },
    MuiTableRow: {
      styleOverrides: {
        root: {
          transition: 'background-color 0.2s ease',
          '&:hover': { backgroundColor: theme.mode === 'dark' ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.02)' },
          '&:last-child td': { borderBottom: 0 },
          '&.MuiTableRow-root.MuiTableRow-hover:hover': {
            backgroundColor: theme.mode === 'dark' ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.02)'
          },
          '& .MuiCollapse-root': { '&:hover': { backgroundColor: 'transparent' } },
          '& .MuiCollapse-root .MuiTableRow-root': { '&:hover': { backgroundColor: 'transparent' } }
        }
      }
    },
    MuiTablePagination: {
      styleOverrides: {
        root: {
          color: theme.textDark,
          borderTop: `1px solid ${theme.tableBorderBottom}`,
          overflow: 'auto',
          backgroundColor: theme.headBackgroundColor,
          minHeight: '56px',
          width: '100%',
          margin: 0,
          padding: '0 24px',
          '& .MuiToolbar-root': {
            minHeight: '56px',
            padding: '0',
            ...(theme.breakpoints && {
              [theme.breakpoints.down('sm')]: { flexWrap: 'wrap', justifyContent: 'center', padding: '8px 0' }
            }),
            '& > p:first-of-type': { fontSize: '0.875rem', color: theme.darkTextSecondary }
          }
        },
        select: { paddingTop: '6px', paddingBottom: '6px', paddingLeft: '12px', paddingRight: '28px', borderRadius: '8px' }
      }
    }
  };
}

