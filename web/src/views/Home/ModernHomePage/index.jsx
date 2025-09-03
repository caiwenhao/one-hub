import React from 'react';
import { Box } from '@mui/material';
import HeroSection from './components/HeroSection';
import FeaturesSection from './components/FeaturesSection';
import StatsSection from './components/StatsSection';
import TechSection from './components/TechSection';
import ModelsSection from './components/ModelsSection';
import AudienceSection from './components/AudienceSection';
import CTASection from './components/CTASection';

const ModernHomePage = () => {
  return (
    <Box
      sx={{
        minHeight: '100vh',
        backgroundColor: '#ffffff',
        overflow: 'hidden',
        fontFamily: 'Inter, sans-serif',
        '& *': {
          scrollbarWidth: 'none',
          '&::-webkit-scrollbar': {
            display: 'none'
          }
        }
      }}
    >
      <HeroSection />
      <FeaturesSection />
      <StatsSection />
      <TechSection />
      <ModelsSection />
      <AudienceSection />
      <CTASection />
    </Box>
  );
};

export default ModernHomePage;
