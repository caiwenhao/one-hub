// material-ui
import { styled } from '@mui/material/styles';
import { useSelector } from 'react-redux';
import { useNavigate, useLocation } from 'react-router';
import { useEffect, useContext } from 'react';
import { UserContext } from 'contexts/UserContext';

// ==============================|| AUTHENTICATION 1 WRAPPER ||============================== //

const AuthStyle = styled('div')(({ theme }) => ({
  backgroundColor: theme.palette.background.default
}));

// eslint-disable-next-line
const AuthWrapper = ({ children }) => {
  const account = useSelector((state) => state.account);
  const { isUserLoaded } = useContext(UserContext);
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (isUserLoaded && account.user) {
      // 如果用户已登录，检查是否有重定向参数
      const searchParams = new URLSearchParams(location.search);
      const redirectUrl = searchParams.get('redirect') || '/panel';
      navigate(redirectUrl);
    }
  }, [account, navigate, isUserLoaded, location.search]);

  // 在用户信息加载完成前显示加载状态
  if (!isUserLoaded) {
    return <AuthStyle>加载中...</AuthStyle>;
  }

  return <AuthStyle> {children} </AuthStyle>;
};

export default AuthWrapper;
