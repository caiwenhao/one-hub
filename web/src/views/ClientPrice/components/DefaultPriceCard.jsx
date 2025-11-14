import { Card, Typography, Box } from '@mui/material';
import { convertPrice } from '../../Pricing/component/priceConverter';

const DefaultPriceCard = ({ defaultPrice }) => {
  if (!defaultPrice) return null;

  return (
    <Card sx={{ p: 2 }}>
      <Typography variant="subtitle1" gutterBottom>
        官方基础价格（仅供参考，来自全局价表）
      </Typography>
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 0.5 }}>
        <Typography variant="body2" color="text.secondary">
          类型：{defaultPrice.type || '-'}
        </Typography>
        {defaultPrice.type === 'tokens' && (
          <>
            <Typography variant="body2" color="text.secondary">
              输入：Rate(K) = {defaultPrice.input ?? 0}，约 RMB{' '}
              {convertPrice(defaultPrice.input ?? 0, 'rate', 'K', 'RMB', 'M')} / 百万 tokens
            </Typography>
            <Typography variant="body2" color="text.secondary">
              输出：Rate(K) = {defaultPrice.output ?? 0}，约 RMB{' '}
              {convertPrice(defaultPrice.output ?? 0, 'rate', 'K', 'RMB', 'M')} / 百万 tokens
            </Typography>
          </>
        )}
        {(defaultPrice.type === 'times' || defaultPrice.type === 'seconds') && (
          <Typography variant="body2" color="text.secondary">
            输出：Rate(K) = {defaultPrice.output ?? 0}，约 RMB{' '}
            {convertPrice(defaultPrice.output ?? 0, 'rate', 'K', 'RMB', 'K')}{' '}
            / {defaultPrice.type === 'times' ? '次' : '秒'}
          </Typography>
        )}
      </Box>
    </Card>
  );
};

export default DefaultPriceCard;
