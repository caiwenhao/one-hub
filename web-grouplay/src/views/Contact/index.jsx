import React from 'react';
import { Box } from '@mui/material';
import HeroSection from './components/HeroSection';
import ContactChannels from './components/ContactChannels';
import HelpResources from './components/HelpResources';

// ==============================|| CONTACT PAGE ||============================== //

const Contact = () => {
  return (
    <Box sx={{
      minHeight: '100vh',
      backgroundColor: '#ffffff',
      overflow: 'hidden'
    }}>
      {/* Hero区域 */}
      <HeroSection />

      {/* 联系方式卡片区域 */}
      <ContactChannels />

      {/* 帮助资源区域 */}
      <HelpResources />
    </Box>
  );
};

export default Contact;
