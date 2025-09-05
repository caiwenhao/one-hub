import PropTypes from 'prop-types';
import { useEffect, useState, useCallback, useContext } from 'react';
import { API } from 'utils/api';
import { getChatLinks, showError, replaceChatPlaceholders } from 'utils/common';
import { Typography, Tabs, Tab, Box, Card } from '@mui/material';
import SubCard from 'ui-component/cards/SubCard';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { UserContext } from 'contexts/UserContext';

function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`playground-tabpanel-${index}`}
      aria-labelledby={`playground-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.number.isRequired,
  value: PropTypes.number.isRequired
};

function a11yProps(index) {
  return {
    id: `playground-tab-${index}`,
    'aria-controls': `playground-tabpanel-${index}`
  };
}

const Playground = () => {
  const [value, setValue] = useState('');
  const [tabIndex, setTabIndex] = useState(0);
  const [isLoading, setIsLoading] = useState(true);
  const siteInfo = useSelector((state) => state.siteInfo);
  const account = useSelector((state) => state.account);
  const { isUserLoaded } = useContext(UserContext);
  const navigate = useNavigate();
  const chatLinks = getChatLinks(true);
  const [iframeSrc, setIframeSrc] = useState(null);

  const loadTokens = useCallback(async () => {
    setIsLoading(true);
    try {
      const res = await API.get(`/api/token/playground`);
      const { success, message, data } = res.data;
      if (success) {
        setValue(data);
      } else {
        showError(message);
      }
    } catch (error) {
      // 如果是401错误（未认证），跳转到登录页
      if (error.response?.status === 401) {
        navigate('/login?redirect=/playground');
        return;
      }
      showError(error.message || '获取令牌失败');
    }
    setIsLoading(false);
  }, [navigate]);

  const handleTabChange = useCallback(
    (event, newIndex) => {
      setTabIndex(newIndex);
      let server = '';
      if (siteInfo?.server_address) {
        server = siteInfo.server_address;
      } else {
        server = window.location.host;
      }
      server = encodeURIComponent(server);
      const key = 'sk-' + value;

      setIframeSrc(replaceChatPlaceholders(chatLinks[newIndex].url, key, server));
    },
    [siteInfo, value, chatLinks]
  );

  // 检查登录状态
  useEffect(() => {
    if (isUserLoaded && !account.user) {
      // 保存当前页面路径，登录后跳转回来
      navigate('/login?redirect=/playground');
      return;
    }
  }, [account, navigate, isUserLoaded]);

  useEffect(() => {
    // 只有在用户已登录的情况下才加载令牌
    if (isUserLoaded && account.user) {
      loadTokens().then(() => {
        if (value !== '') {
          handleTabChange(null, 0);
        }
      });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [loadTokens, value, isUserLoaded, account.user]);

  // 在用户信息加载完成前显示加载状态
  if (!isUserLoaded) {
    return (
      <SubCard title="Playground">
        <Typography align="center">加载中...</Typography>
      </SubCard>
    );
  }

  // 如果用户未登录，不显示内容（会被重定向到登录页）
  if (!account.user) {
    return null;
  }

  if (chatLinks.length === 0 || isLoading || value === '') {
    return (
      <SubCard title="Playground">
        <Typography align="center">{isLoading ? 'Loading...' : 'No playground available'}</Typography>
      </SubCard>
    );
  } else if (chatLinks.length === 1) {
    return (
      <iframe title="playground" src={iframeSrc} style={{ width: '100%', height: '85vh', border: 'none' }} />
    );
  } else {
    return (
      <Card>
        <Tabs variant="scrollable" value={tabIndex} onChange={handleTabChange} sx={{ borderRight: 1, borderColor: 'divider' }}>
          {chatLinks.map((link, index) => link.show && <Tab label={link.name} {...a11yProps(index)} key={index} />)}
        </Tabs>
        <Box>
          <iframe title="playground" src={iframeSrc} style={{ width: '100%', height: '85vh', border: 'none' }} />
        </Box>
      </Card>
    );
  }
};

export default Playground;
