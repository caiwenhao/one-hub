// 搜索输入（防抖、清除按钮、占位语气，保留关键字）
import { useEffect, useMemo, useState } from 'react';
import { InputAdornment, IconButton, TextField } from '@mui/material';
import CloseIcon from '@mui/icons-material/Close';

export default function SearchInput({ value: external, onChange, delay = 300, placeholder = '搜索关键词' }) {
  const [value, setValue] = useState(external || '');
  useEffect(() => setValue(external || ''), [external]);

  const debounced = useMemo(() => {
    let t;
    return (v) => {
      clearTimeout(t);
      t = setTimeout(() => onChange?.(v), delay);
    };
  }, [onChange, delay]);

  const onClear = () => {
    setValue('');
    onChange?.('');
  };

  return (
    <TextField
      placeholder={placeholder}
      value={value}
      onChange={(e) => {
        const v = e.target.value;
        setValue(v);
        debounced(v);
      }}
      InputProps={{
        endAdornment:
          value && value.length > 0 ? (
            <InputAdornment position="end">
              <IconButton size="small" onClick={onClear} aria-label="清除关键词">
                <CloseIcon fontSize="small" />
              </IconButton>
            </InputAdornment>
          ) : null
      }}
      fullWidth
    />
  );
}
