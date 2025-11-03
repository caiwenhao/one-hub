// material-ui
import React from 'react';
import { Box, Container, Typography, Grid } from '@mui/material';
import { useSelector } from 'react-redux';

// ==============================|| SITE-WIDE DESIGN FOOTER (方案B) ||============================== //

const DesignFooter = () => {
  const navigateTo = (href) => {
    if (!href) return;
    if (href.startsWith('http') || href.startsWith('mailto:')) {
      window.open(href, '_blank');
    } else {
      window.location.href = href;
    }
  };

  const gradient = 'linear-gradient(-45deg, #0EA5FF, #22D3EE, #8B5CF6)';

  return (
    <Box component="footer" sx={{ backgroundColor: '#ffffff', borderTop: '1px solid rgba(0,0,0,0.08)', py: { xs: 8, md: 12 } }}>
      <Container maxWidth="lg" sx={{ maxWidth: '1200px' }}>
        <Grid container spacing={{ xs: 4, md: 6 }}>
          {/* 品牌区域 */}
          <Grid item xs={12} md={4}>
            <Box>
              {/* Logo */}
              <Box sx={{ display: 'flex', alignItems: 'center', mb: 3, cursor: 'pointer' }} onClick={() => navigateTo('/') }>
                <img
                  src="/logo.png"
                  alt="Logo"
                  style={{
                    height: '48px',
                    width: 'auto'
                  }}
                />
              </Box>
              {/* 标语 */}
              <Typography variant="body1" sx={{ color: '#718096', lineHeight: 1.6, fontSize: '1.125rem', fontWeight: 300, mb: 4 }}>
                稳定，是AI应用的唯一标准
              </Typography>
              {/* 社交/链接区域已移除 */}
            </Box>
          </Grid>

          {/* 导航与支持 */}
          <Grid item xs={12} md={8}>
            <Grid container spacing={{ xs: 4, md: 6 }}>
              {[
                { title: '导航', links: [ { label: '热门模型', href: '/models' }, { label: '价格方案', href: '/price' }, { label: '开发者中心', href: '/developer' } ] },
                { title: '支持', links: [ { label: '联系我们', href: '/contact' } ] },
                { title: '联系', links: [ { label: 'sales@grouplay.cn', href: null }, { label: 'support@grouplay.cn', href: null } ] }
              ].map((section) => (
                <Grid key={section.title} item xs={12} sm={4}>
                  <Typography variant="h6" sx={{ color: '#1A202C', fontWeight: 'bold', mb: 3, fontSize: '1.125rem' }}>
                    {section.title}
                  </Typography>
                  <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                    {section.links.map((l) => (
                      <Typography
                        key={l.label}
                        onClick={l.href ? () => navigateTo(l.href) : undefined}
                        sx={{
                          color: '#718096',
                          fontSize: '1rem',
                          fontWeight: 300,
                          cursor: l.href ? 'pointer' : 'default',
                          transition: l.href ? 'all 0.3s ease' : 'none',
                          '&:hover': l.href ? { color: '#0EA5FF', transform: 'translateX(4px)' } : {}
                        }}
                      >
                        {l.label}
                      </Typography>
                    ))}
                  </Box>
                </Grid>
              ))}
            </Grid>
          </Grid>
        </Grid>

        {/* 底部版权信息 */}
        <Box sx={{ borderTop: '1px solid rgba(0, 0, 0, 0.08)', pt: 6, mt: 8, textAlign: 'center' }}>
          <Typography variant="body2" sx={{ color: '#718096', fontSize: '1rem', fontWeight: 300 }}>
            © 2025 grouplay AI. All rights reserved.
          </Typography>
          <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4, opacity: 0.6 }}>
            <Box sx={{ display: 'flex', gap: 1.5, alignItems: 'center' }}>
              {[...Array(3)].map((_, i) => (
                <Box key={i} sx={{ width: 6, height: 6, borderRadius: '50%', background: gradient }} />
              ))}
            </Box>
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

const Footer = () => {
  const siteInfo = useSelector((state) => state.siteInfo);
  // 如果后台设置了自定义 Footer HTML，仍然优先展示
  if (siteInfo?.footer_html) {
    return (
      <Container sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', py: 4, borderRadius: 0 }}>
        <Box sx={{ textAlign: 'center', width: '100%' }}>
          <div className="custom-footer" dangerouslySetInnerHTML={{ __html: siteInfo.footer_html }} />
        </Box>
      </Container>
    );
  }
  // 否则使用设计稿版 Footer（整站统一）
  return <DesignFooter />;
};

export default Footer;
