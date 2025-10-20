export default function datagrid(theme) {
  return {
    MuiDataGrid: {
      styleOverrides: {
        root: {
          border: 'none',
          borderRadius: 'var(--radii-lg)',
          backgroundColor: theme.mode === 'dark' ? theme.paper : theme.colors?.paper,
          overflow: 'hidden',
          width: '100%',
          margin: 0,
          padding: 0,
          '& .MuiPaper-root': { borderRadius: 'var(--radii-lg)' },
          '&.MuiPaper-root': { borderRadius: 'var(--radii-lg)' },
          '& .MuiDataGrid-main': {
            width: '100%',
            margin: 0,
            padding: 0,
            '& .MuiDataGrid-columnHeaders': {
              borderBottom: `1px solid ${theme.tableBorderBottom}`,
              borderRadius: 'var(--radii-lg)',
              backgroundColor: theme.headBackgroundColor,
              minHeight: '48px',
              width: '100%',
              margin: 0
            },
            '& .MuiDataGrid-virtualScroller': {
              backgroundColor: theme.mode === 'dark' ? theme.paper : '#fff',
              width: '100%',
              margin: 0
            },
            '& .MuiDataGrid-columnHeadersInner': {
              width: '100%',
              margin: 0,
              padding: 0,
              '& .MuiDataGrid-columnHeader': {
                padding: 0,
                margin: 0,
                '& .MuiDataGrid-columnHeaderTitleContainer': { justifyContent: 'center', padding: '0 var(--density-table-cell-px)', margin: 0 },
                '&:first-of-type': { '& .MuiDataGrid-columnHeaderTitleContainer': { paddingLeft: '24px' } },
                '&:last-of-type': { '& .MuiDataGrid-columnHeaderTitleContainer': { paddingRight: '24px' } }
              }
            },
            '& .MuiDataGrid-cellContent': { justifyContent: 'center', width: '100%', display: 'flex' },
            '& .MuiDataGrid-toolbarContainer': {
              position: 'sticky',
              top: 0,
              zIndex: 2,
              backgroundColor: theme.headBackgroundColor,
              borderBottom: `1px solid ${theme.tableBorderBottom}`
            }
          },
          footerContainer: {
            borderTop: `1px solid ${theme.tableBorderBottom}`,
            minHeight: '56px',
            backgroundColor: theme.headBackgroundColor,
            width: '100%',
            margin: 0,
            padding: '0 24px',
            '& .MuiTablePagination-root': { overflow: 'visible', backgroundColor: 'transparent', color: theme.textDark, borderTop: 'none' },
            '& .MuiToolbar-root': {
              minHeight: '56px',
              padding: '0',
              '& > p:first-of-type': { fontSize: '0.875rem', color: theme.darkTextSecondary }
            }
          }
        },
        columnHeader: {
          padding: 'var(--density-table-head-py) var(--density-table-cell-px)',
          fontSize: '0.875rem',
          fontWeight: 600,
          color: theme.darkTextSecondary,
          height: '48px',
          textAlign: 'center',
          '&:focus': { outline: 'none' },
          '&:focus-within': { outline: 'none' }
        },
        columnHeaderTitle: { color: theme.darkTextSecondary, fontWeight: 600, fontSize: '0.875rem', textAlign: 'center' },
        columnSeparator: { color: theme.divider },
        cell: {
          fontSize: '0.875rem',
          padding: 'var(--density-table-cell-py) var(--density-table-cell-px)',
          borderBottom: `1px solid ${theme.tableBorderBottom}`,
          textAlign: 'center',
          '&:focus': { outline: 'none' },
          '&:focus-within': { outline: 'none' }
        },
        row: {
          transition: 'background-color 0.2s ease',
          '&:hover': { backgroundColor: theme.mode === 'dark' ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.02)' },
          '&.Mui-selected': {
            backgroundColor: theme.mode === 'dark' ? 'rgba(22, 119, 255, 0.18)' : 'rgba(22, 119, 255, 0.08)',
            '&:hover': { backgroundColor: theme.mode === 'dark' ? 'rgba(22, 119, 255, 0.24)' : 'rgba(22, 119, 255, 0.12)' }
          }
        },
        rowCount: { color: theme.darkTextSecondary },
        selectedRowCount: { color: theme.colors?.primaryMain, fontWeight: 500 },
        toolbarContainer: { backgroundColor: theme.mode === 'dark' ? theme.paper : theme.colors?.paper, padding: 'var(--density-table-cell-py) var(--density-table-cell-px)', '& .MuiButton-root': { marginRight: 'var(--spacing-sm)' } },
        panelHeader: { backgroundColor: theme.mode === 'dark' ? theme.paper : theme.colors?.paper, padding: '16px 20px', borderBottom: `1px solid ${theme.divider}` },
        panelContent: { padding: '16px 20px' }
      }
    }
  };
}

