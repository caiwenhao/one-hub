import React, { useEffect, useState } from 'react';
import ModernHomePage from './ModernHomePage';
import { Box } from '@mui/material';
import ContentViewer from 'ui-component/ContentViewer';

const Home = () => {
  const [homePageContentLoaded, setHomePageContentLoaded] = useState(false);
  const [homePageContent, setHomePageContent] = useState('');

  const displayHomePageContent = async () => {
    const cached = localStorage.getItem('home_page_content') || '';
    setHomePageContent(cached);
    setHomePageContentLoaded(true);
  };

  useEffect(() => {
    displayHomePageContent().then();
  }, []);

  return (
    <>
      {homePageContentLoaded && homePageContent === '' ? (
        <ModernHomePage />
      ) : (
        <Box>
          <ContentViewer
            content={homePageContent}
            loading={!homePageContentLoaded}
            errorMessage={''}
            containerStyle={{ minHeight: 'calc(100vh - 136px)' }}
            contentStyle={{ fontSize: 'larger' }}
          />
        </Box>
      )}
    </>
  );
};

export default Home;
