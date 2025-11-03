import React from 'react';
import { Box } from '@mui/material';
import FeaturedModels from './components/FeaturedModels';
import CategorySection from './components/CategorySection';

const HotModelsPage = () => {
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
      <FeaturedModels />
      <CategorySection />
    </Box>
  );
};

export default HotModelsPage;
