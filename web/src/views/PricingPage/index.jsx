import { Box } from '@mui/material';

// 导入营销页面组件
import FreeTrialSection from '../ModelPrice/components/FreeTrialSection';
import UnifiedPricingTable from '../ModelPrice/components/UnifiedPricingTable';

// ----------------------------------------------------------------------
export default function PricingPage() {
  return (
    <Box sx={{ minHeight: '100vh' }}>
      {/* 免费试用区域 */}
      <FreeTrialSection />

      {/* 统一的详细价格表格 */}
      <UnifiedPricingTable />
    </Box>
  );
}
