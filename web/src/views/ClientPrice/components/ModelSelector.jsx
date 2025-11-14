import { TextField, Autocomplete } from '@mui/material';

const ModelSelector = ({ value, onChange, options = [], disabled = false, loading = false }) => {
  return (
    <Autocomplete
      size="small"
      sx={{ minWidth: 280 }}
      options={options}
      value={value || null}
      loading={loading}
      disabled={disabled}
      onChange={(_, newValue) => onChange(newValue || '')}
      getOptionLabel={(option) => option || ''}
      filterOptions={(options, { inputValue }) => {
        // 支持模糊搜索
        const searchTerm = inputValue.toLowerCase();
        return options.filter((option) => option.toLowerCase().includes(searchTerm));
      }}
      renderInput={(params) => (
        <TextField
          {...params}
          label="模型（支持搜索）"
          placeholder="输入模型名搜索"
        />
      )}
    />
  );
};

export default ModelSelector;
