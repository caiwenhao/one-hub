import { ToggleButtonGroup, ToggleButton, Box } from '@mui/material';
import { Icon } from '@iconify/react';

const PriceInputMode = ({ value, onChange }) => {
  return (
    <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
      <ToggleButtonGroup
        size="small"
        value={value}
        exclusive
        onChange={(_, newValue) => {
          if (newValue !== null) {
            onChange(newValue);
          }
        }}
        aria-label="价格输入模式"
      >
        <ToggleButton value="discount" aria-label="折扣模式">
          <Icon icon="solar:percent-bold-duotone" width={18} style={{ marginRight: 4 }} />
          折扣
        </ToggleButton>
        <ToggleButton value="rmb" aria-label="百万 tokens 价格模式">
          <Icon icon="solar:wallet-money-bold-duotone" width={18} style={{ marginRight: 4 }} />
          百万 tokens 价格
        </ToggleButton>
      </ToggleButtonGroup>
    </Box>
  );
};

export default PriceInputMode;
