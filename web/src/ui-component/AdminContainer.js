import { styled } from '@mui/material/styles';
import { Container } from '@mui/material';

const AdminContainer = styled(Container)(({ theme }) => ({
  paddingLeft: '0px !important',
  paddingRight: '0px !important',
  paddingBottom: `${theme.spacing(4)} !important`,
  marginBottom: `${theme.spacing(4)} !important`,
  width: '100%',
  [theme.breakpoints.up('lg')]: {
    maxWidth: '1320px !important'
  },
  [theme.breakpoints.up('xl')]: {
    maxWidth: '1440px !important'
  }
}));

export default AdminContainer;
