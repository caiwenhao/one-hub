// material-ui
import { useSelector } from 'react-redux';
import { useTheme } from '@mui/material/styles';

// ==============================|| LOGO ||============================== //

const Logo = () => {
  const siteInfo = useSelector((state) => state.siteInfo);

  if (siteInfo.isLoading) {
    return null; // 数据加载未完成时不显示 logo
  }

  // 统一使用 web/public/logo.png
  const logoSrc = '/logo.png';
  return <img src={logoSrc} alt={siteInfo.system_name || 'Logo'} height="50" />;
};

export default Logo;
