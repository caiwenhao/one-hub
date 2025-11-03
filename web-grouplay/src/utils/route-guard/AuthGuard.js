import { useSelector } from 'react-redux';
import { useEffect, useContext } from 'react';
import { UserContext } from 'contexts/UserContext';
import { useNavigate } from 'react-router-dom';

const AuthGuard = ({ children }) => {
  const account = useSelector((state) => state.account);
  const { isUserLoaded } = useContext(UserContext);
  const navigate = useNavigate();

  // 开发期基线截图/演示绕过（仅 DEV 且带 ?baseline=1 或 ?demo=1）
  const params = new URLSearchParams(window.location.search);
  const bypass = import.meta.env.DEV && (params.get('baseline') === '1' || params.get('demo') === '1');

  useEffect(() => {
    if (!bypass && isUserLoaded && !account.user) {
      navigate('/login');
      return;
    }
  }, [account, navigate, isUserLoaded, bypass]);

  // 在用户信息加载完成前不渲染子组件（bypass 时直接放行）
  if (!bypass && !isUserLoaded) {
    return null;
  }

  return children;
};

export default AuthGuard;
