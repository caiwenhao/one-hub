import { styled } from '@mui/material/styles';
import { Container } from '@mui/material';

const AdminContainer = styled(Container)(({ theme }) => ({
  paddingLeft: '0px !important',
  paddingRight: '0px !important',
  paddingBottom: `${theme.spacing(4)} !important`,
  marginBottom: `${theme.spacing(4)} !important`,
  width: '100%',
  // 流体布局：使用 clamp 在 1280–1440+ 之间平滑过渡
  '&.MuiContainer-root': {
    maxWidth: 'clamp(1200px, 90vw, 1440px)'
  }
}));

export default AdminContainer;
