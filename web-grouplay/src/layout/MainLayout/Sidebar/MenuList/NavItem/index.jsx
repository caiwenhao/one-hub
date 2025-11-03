import PropTypes from 'prop-types';
import { forwardRef, useEffect } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';

// material-ui
import { alpha, useTheme } from '@mui/material/styles';
import { Avatar, Chip, ListItemButton, ListItemIcon, ListItemText, Typography, useMediaQuery } from '@mui/material';

// project imports
import { MENU_OPEN, SET_MENU } from 'store/actions';

// assets
import FiberManualRecordIcon from '@mui/icons-material/FiberManualRecord';

// ==============================|| SIDEBAR MENU LIST ITEMS ||============================== //

const NavItem = ({ item, level }) => {
  const theme = useTheme();
  const dispatch = useDispatch();
  const { pathname } = useLocation();
  const customization = useSelector((state) => state.customization);
  const matchesSM = useMediaQuery(theme.breakpoints.down('lg'));
  const isSelected = customization.isOpen.includes(item.id);

  const leftPaddingUnit = level === 1 ? 2 : 2 + (level - 1) * 1.5;
  const leftPadding = theme.spacing(leftPaddingUnit);
  const activeBg = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.28 : 0.12);
  const hoverBg = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.18 : 0.08);

  const iconSize = 20;
  const iconWrapperSize = 36;
  const Icon = item.icon;
  const itemIcon = Icon ? (
    <Icon stroke={1.5} size={iconSize} />
  ) : (
    <FiberManualRecordIcon
      sx={{
        width: isSelected ? 10 : 8,
        height: isSelected ? 10 : 8,
        transition: theme.transitions.create(['transform', 'color'], {
          duration: theme.transitions.duration.shorter
        })
      }}
      fontSize={level > 0 ? 'inherit' : 'medium'}
    />
  );

  let itemTarget = '_self';
  if (item.target) {
    itemTarget = '_blank';
  }

  let listItemProps = {
    component: forwardRef((props, ref) => <Link ref={ref} {...props} to={item.url} target={itemTarget} />)
  };
  if (item?.external) {
    listItemProps = { component: 'a', href: item.url, target: itemTarget };
  }

  const itemHandler = (id) => {
    dispatch({ type: MENU_OPEN, id });
    if (matchesSM) dispatch({ type: SET_MENU, opened: false });
  };

  // 页面刷新时保持选中态
  useEffect(() => {
    const currentIndex = document.location.pathname
      .toString()
      .split('/')
      .findIndex((id) => id === item.id);
    if (currentIndex > -1) {
      dispatch({ type: MENU_OPEN, id: item.id });
    }
    // eslint-disable-next-line
  }, [pathname]);

  return (
    <ListItemButton
      {...listItemProps}
      disabled={item.disabled}
      sx={{
        position: 'relative',
        borderRadius: `${Math.max(customization.borderRadius, 10)}px`,
        mb: 0.5,
        alignItems: 'center',
        minHeight: level > 1 ? 40 : 44,
        px: 1.5,
        py: 0.75,
        pl: leftPadding,
        color: isSelected ? theme.palette.primary.main : theme.palette.text.secondary,
        backgroundColor: isSelected ? activeBg : 'transparent',
        transition: theme.transitions.create(['background-color', 'color'], {
          duration: theme.transitions.duration.shorter
        }),
        '&::before': {
          content: '""',
          position: 'absolute',
          left: theme.spacing(1),
          top: theme.spacing(0.75),
          bottom: theme.spacing(0.75),
          width: '3px',
          borderRadius: '8px',
          backgroundColor: isSelected ? theme.palette.primary.main : 'transparent',
          transform: isSelected ? 'scaleY(1)' : 'scaleY(0.2)',
          opacity: isSelected ? 1 : 0,
          transition: theme.transitions.create(['transform', 'opacity', 'background-color'], {
            duration: theme.transitions.duration.shorter
          })
        },
        '&:hover': {
          backgroundColor: hoverBg,
          color: theme.palette.primary.main,
          '& .MuiListItemIcon-root': {
            color: theme.palette.primary.main
          },
          '&::before': {
            transform: 'scaleY(1)',
            opacity: 1,
            backgroundColor: theme.palette.primary.main
          }
        },
        '&.Mui-selected': {
          backgroundColor: activeBg,
          color: theme.palette.primary.main,
          '& .MuiListItemIcon-root': {
            color: theme.palette.primary.main
          },
          '&:hover': {
            backgroundColor: activeBg
          }
        }
      }}
      selected={isSelected}
      onClick={() => itemHandler(item.id)}
    >
      <ListItemIcon
        sx={{
          my: 'auto',
          minWidth: iconWrapperSize,
          color: isSelected ? theme.palette.primary.main : theme.palette.text.secondary,
          transition: theme.transitions.create(['color'], {
            duration: theme.transitions.duration.shorter
          }),
          '& svg': {
            width: iconSize,
            height: iconSize
          }
        }}
      >
        {itemIcon}
      </ListItemIcon>
      <ListItemText
        primary={item.title}
        primaryTypographyProps={{
          variant: isSelected ? 'subtitle1' : 'body2',
          fontWeight: isSelected ? 600 : 500,
          noWrap: true
        }}
        secondary={
          item.caption && (
            <Typography variant="caption" sx={{ ...theme.typography.subMenuCaption }} display="block" gutterBottom>
              {item.caption}
            </Typography>
          )
        }
      />
      {item.chip && (
        <Chip
          color={item.chip.color}
          variant={item.chip.variant}
          size={item.chip.size}
          label={item.chip.label}
          avatar={item.chip.avatar && <Avatar>{item.chip.avatar}</Avatar>}
        />
      )}
    </ListItemButton>
  );
};

NavItem.propTypes = {
  item: PropTypes.object,
  level: PropTypes.number
};

export default NavItem;
