import PropTypes from 'prop-types';
import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

// material-ui
import { useTheme } from '@mui/material/styles';
import { Box, Stack, Typography } from '@mui/material';
import MuiBreadcrumbs from '@mui/material/Breadcrumbs';

// project imports
import config from 'config';

// assets
import { IconTallymark1 } from '@tabler/icons-react';
import AccountTreeTwoToneIcon from '@mui/icons-material/AccountTreeTwoTone';
import HomeIcon from '@mui/icons-material/Home';
import HomeTwoToneIcon from '@mui/icons-material/HomeTwoTone';

const linkSX = {
  display: 'flex',
  color: 'grey.900',
  textDecoration: 'none',
  alignContent: 'center',
  alignItems: 'center'
};

// ==============================|| BREADCRUMBS ||============================== //

const Breadcrumbs = ({
  card = true,
  divider = true,
  icon,
  icons,
  maxItems,
  navigation,
  rightAlign,
  separator,
  title,
  titleBottom,
  ...others
}) => {
  const theme = useTheme();

  const iconStyle = {
    marginRight: theme.spacing(0.75),
    marginTop: `-${theme.spacing(0.25)}`,
    width: '1rem',
    height: '1rem',
    color: theme.palette.secondary.main
  };

  const [main, setMain] = useState();
  const [item, setItem] = useState();

  // set active item state
  const getCollapse = (menu) => {
    if (menu.children) {
      menu.children.filter((collapse) => {
        if (collapse.type && collapse.type === 'collapse') {
          getCollapse(collapse);
        } else if (collapse.type && collapse.type === 'item') {
          if (document.location.pathname === config.basename + collapse.url) {
            setMain(menu);
            setItem(collapse);
          }
        }
        return false;
      });
    }
  };

  useEffect(() => {
    navigation?.items?.map((menu) => {
      if (menu.type && menu.type === 'group') {
        getCollapse(menu);
      }
      return false;
    });
  });

  // item separator
  const separatorIcon = separator ? (
    typeof separator === 'function' ? (
      <separator stroke={1.5} size="1rem" />
    ) : (
      separator
    )
  ) : (
    <IconTallymark1 stroke={1.5} size="1rem" />
  );

  let mainContent;
  let itemContent;
  let breadcrumbContent = <Typography />;
  let itemTitle = '';
  let CollapseIcon;
  let ItemIcon;

  // collapse item
  if (main && main.type === 'collapse') {
    CollapseIcon = main.icon ? main.icon : AccountTreeTwoToneIcon;
    mainContent = (
      <Typography component={Link} to="#" variant="subtitle1" sx={linkSX}>
        {icons && <CollapseIcon style={iconStyle} />}
        {main.title}
      </Typography>
    );
  }

  // items
  if (item && item.type === 'item') {
    itemTitle = item.title;

    ItemIcon = item.icon ? item.icon : AccountTreeTwoToneIcon;
    itemContent = (
      <Typography
        variant="subtitle1"
        sx={{
          display: 'flex',
          textDecoration: 'none',
          alignContent: 'center',
          alignItems: 'center',
          color: 'grey.500'
        }}
      >
        {icons && <ItemIcon style={iconStyle} />}
        {itemTitle}
      </Typography>
    );

    if (item.breadcrumbs !== false) {
      const useRowLayout = rightAlign && !titleBottom;
      const renderRightBreadcrumb = useRowLayout && title;

      const titleNode = (
        <Typography variant="h3" sx={{ fontWeight: 500 }}>
          {item.title}
        </Typography>
      );

      const breadcrumbsNode = (
        <MuiBreadcrumbs
          sx={{ '& .MuiBreadcrumbs-separator': { width: 16, ml: 1.25, mr: 1.25 } }}
          aria-label="breadcrumb"
          maxItems={maxItems || 8}
          separator={separatorIcon}
        >
          <Typography component={Link} to="/" color="inherit" variant="subtitle1" sx={linkSX}>
            {icons && <HomeTwoToneIcon sx={iconStyle} />}
            {icon && <HomeIcon sx={{ ...iconStyle, mr: 0 }} />}
            {!icon && 'Dashboard'}
          </Typography>
          {mainContent}
          {itemContent}
        </MuiBreadcrumbs>
      );

      const leftSection = [];
      if (title && !titleBottom) {
        leftSection.push(titleNode);
      }
      if (!renderRightBreadcrumb) {
        leftSection.push(breadcrumbsNode);
      }
      if (title && titleBottom) {
        leftSection.push(titleNode);
      }

      breadcrumbContent = (
        <Box
          sx={{
            mb: theme.spacing(card ? 3 : 2.5),
            ...(card
              ? {
                  px: { xs: 1.5, md: 2 },
                  py: { xs: 1.25, md: 1.5 },
                  borderRadius: `${theme.shape.borderRadius}px`,
                  border: `1px solid ${theme.palette.divider}`,
                  backgroundColor: theme.palette.background.paper
                }
              : {
                  pb: divider !== false ? theme.spacing(1.5) : 0,
                  borderBottom: divider !== false ? `1px solid ${theme.palette.divider}` : 'none'
                })
          }}
          {...others}
        >
          <Stack
            direction={useRowLayout ? { xs: 'column', md: 'row' } : 'column'}
            alignItems={useRowLayout ? { xs: 'flex-start', md: 'center' } : 'flex-start'}
            justifyContent={useRowLayout ? { xs: 'flex-start', md: 'space-between' } : 'flex-start'}
            spacing={useRowLayout ? { xs: 1, md: 1.5 } : 0.75}
          >
            <Stack spacing={0.5} sx={{ minWidth: 0 }}>
              {leftSection.map((node, index) => (
                <Box key={index} sx={{ minWidth: 0 }}>
                  {node}
                </Box>
              ))}
            </Stack>
            {renderRightBreadcrumb && <Box sx={{ minWidth: 0 }}>{breadcrumbsNode}</Box>}
          </Stack>
        </Box>
      );
    }
  }

  return breadcrumbContent;
};

Breadcrumbs.propTypes = {
  card: PropTypes.bool,
  divider: PropTypes.bool,
  icon: PropTypes.bool,
  icons: PropTypes.bool,
  maxItems: PropTypes.number,
  navigation: PropTypes.object,
  rightAlign: PropTypes.bool,
  separator: PropTypes.oneOfType([PropTypes.func, PropTypes.object]),
  title: PropTypes.bool,
  titleBottom: PropTypes.bool
};

export default Breadcrumbs;
