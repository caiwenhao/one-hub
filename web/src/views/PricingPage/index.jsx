import { Box } from '@mui/material';

// 仅保留详细价格表
import UnifiedPricingTable from '../ModelPrice/components/UnifiedPricingTable';

// ----------------------------------------------------------------------
export default function PricingPage() {
  return (
    <Box sx={{ minHeight: '100vh' }}>
      {/* 统一的详细价格表格 */}
      <UnifiedPricingTable />
    </Box>
  );
}
