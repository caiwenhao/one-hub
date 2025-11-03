import React, { useEffect, useState, useCallback } from 'react';
import { Box, Container, Typography } from '@mui/material';
import MainCard from 'ui-component/cards/MainCard';
import { useTranslation } from 'react-i18next';
import ContentViewer from 'ui-component/ContentViewer';

const About = () => {
  const { t } = useTranslation();
  const [about, setAbout] = useState('');
  const [aboutLoaded, setAboutLoaded] = useState(false);

  const displayAbout = useCallback(async () => {
    const cached = localStorage.getItem('about') || '';
    setAbout(cached);
    setAboutLoaded(true);
  }, []);

  useEffect(() => {
    displayAbout();
  }, [displayAbout]);

  return (
    <>
      {aboutLoaded && about === '' ? (
        <Box>
          <Container sx={{ paddingTop: '40px' }}>
            <MainCard title={t('about.aboutTitle')}>
              <Typography variant="body2">
                {t('about.aboutDescription')} <br />
                {t('about.projectRepo')}
                <a href="https://github.com/MartialBE/one-hub">https://github.com/MartialBE/one-hub</a>
              </Typography>
            </MainCard>
          </Container>
        </Box>
      ) : (
        <Box>
          <ContentViewer
            content={about}
            loading={!aboutLoaded}
            errorMessage={''}
            containerStyle={{ minHeight: 'calc(100vh - 136px)' }}
            contentStyle={{ fontSize: 'larger' }}
          />
        </Box>
      )}
    </>
  );
};

export default About;
