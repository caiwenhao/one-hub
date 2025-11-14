import { useState, useEffect, useMemo } from 'react';
import { TextField, Autocomplete } from '@mui/material';
import { API } from 'utils/api';
import { showError } from 'utils/common';

// 防抖函数
const debounce = (func, wait) => {
  let timeout;
  return (...args) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => func(...args), wait);
  };
};

const UserSelector = ({ value, onChange, disabled = false }) => {
  const [userOptions, setUserOptions] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchUsers = async (keyword = '') => {
    setLoading(true);
    try {
      const res = await API.get('/api/user/', {
        params: {
          page: 1,
          size: 20,
          keyword: keyword.trim()
        }
      });
      const { success, message, data } = res.data || {};
      if (!success) {
        showError(message || '加载用户列表失败');
        setUserOptions([]);
        return;
      }
      setUserOptions(data?.data || []);
    } catch (e) {
      console.error(e);
      showError(e.message || '加载用户列表失败');
      setUserOptions([]);
    } finally {
      setLoading(false);
    }
  };

  // 防抖搜索
  const debouncedFetchUsers = useMemo(
    () => debounce((keyword) => fetchUsers(keyword), 300),
    []
  );

  useEffect(() => {
    // 初始加载一批用户
    fetchUsers('');
  }, []);

  return (
    <Autocomplete
      size="small"
      sx={{ minWidth: 260 }}
      options={userOptions}
      value={value}
      loading={loading}
      disabled={disabled}
      getOptionLabel={(option) =>
        option ? `${option.id} - ${option.username}${option.display_name ? `（${option.display_name}）` : ''}` : ''
      }
      isOptionEqualToValue={(opt, val) => opt.id === val.id}
      onChange={(_, newValue) => onChange(newValue)}
      onInputChange={(_, newInputValue) => {
        debouncedFetchUsers(newInputValue);
      }}
      renderInput={(params) => (
        <TextField
          {...params}
          label="用户（可搜索）"
          placeholder="输入用户名/ID 搜索"
        />
      )}
    />
  );
};

export default UserSelector;
