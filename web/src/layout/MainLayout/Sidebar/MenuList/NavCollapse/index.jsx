import PropTypes from 'prop-types';
import { useEffect, useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { useLocation } from 'react-router';

// material-ui
import { alpha, useTheme } from '@mui/material/styles';
import { Collapse, List, ListItemButton, ListItemIcon, ListItemText, Typography } from '@mui/material';

// project imports
import NavItem from '../NavItem';

// assets
import FiberManualRecordIcon from '@mui/icons-material/FiberManualRecord';
import { IconChevronDown, IconChevronUp } from '@tabler/icons-react';

// ==============================|| SIDEBAR MENU LIST COLLAPSE ITEMS ||============================== //

const NavCollapse = ({ menu, level }) => {
  const theme = useTheme();
  const customization = useSelector((state) => state.customization);
  const { pathname } = useLocation();
  const [open, setOpen] = useState(false);
  const [selected, setSelected] = useState(null);

  const leftPaddingUnit = level === 1 ? 2 : 2 + (level - 1) * 1.5;
  const leftPadding = theme.spacing(leftPaddingUnit);
  const childListMargin = theme.spacing(Math.max(1.5, leftPaddingUnit - 0.5));
  const activeBg = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.28 : 0.12);
  const hoverBg = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.18 : 0.08);
  const connectorColor = alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.25 : 0.16);

  const hasActiveChild = useMemo(() => {
    const loop = (nodes = []) =>
      nodes.some((child) => {
        if (child.children?.length) {
          return loop(child.children);
        }
        return child.url === pathname;
      });
    return loop(menu.children);
  }, [menu.children, pathname]);

  useEffect(() => {
    if (hasActiveChild) {
      setOpen(true);
      setSelected(menu.id);
    } else {
      setSelected(null);
    }
  }, [hasActiveChild, menu.id]);

  const triggerResize = () => window.requestAnimationFrame(() => window.dispatchEvent(new Event('resize')));

  const handleClick = () => {
    setOpen((prev) => !prev);
    setSelected((prev) => (prev ? null : menu.id));
    triggerResize();
  };

  const iconSize = 20;
  const iconWrapperSize = 36;
  const IconComponent = menu.icon;
  const isSelected = selected === menu.id;

  const menuIcon = IconComponent ? (
    <IconComponent strokeWidth={1.5} size={iconSize} />
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

  const menus = menu.children?.map((item) => {
    switch (item.type) {
      case 'collapse':
        return <NavCollapse key={item.id} menu={item} level={level + 1} />;
      case 'item':
        return <NavItem key={item.id} item={item} level={level + 1} />;
      default:
        return (
          <Typography key={item.id} variant="h6" color="error" align="center">
            Menu Items Error
          </Typography>
        );
    }
  });

  return (
    <>
      <ListItemButton
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
        onClick={handleClick}
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
          {menuIcon}
        </ListItemIcon>
        <ListItemText
          primary={menu.title}
          primaryTypographyProps={{
            variant: isSelected ? 'subtitle1' : 'body2',
            fontWeight: isSelected ? 600 : 500,
            noWrap: true
          }}
          secondary={
            menu.caption && (
              <Typography variant="caption" sx={{ ...theme.typography.subMenuCaption }} display="block" gutterBottom>
                {menu.caption}
              </Typography>
            )
          }
        />
        {open ? (
          <IconChevronUp stroke={1.5} size="1rem" style={{ marginTop: 'auto', marginBottom: 'auto' }} />
        ) : (
          <IconChevronDown stroke={1.5} size="1rem" style={{ marginTop: 'auto', marginBottom: 'auto' }} />
        )}
      </ListItemButton>
      <Collapse in={open} timeout="auto" unmountOnExit onEntered={triggerResize} onExited={triggerResize}>
        <List
          component="div"
          disablePadding
          sx={{
            position: 'relative',
            ml: childListMargin,
            pl: theme.spacing(1.5),
            borderLeft: `1px solid ${connectorColor}`,
            mt: 0.5
          }}
        >
          {menus}
        </List>
      </Collapse>
    </>
  );
};

NavCollapse.propTypes = {
  menu: PropTypes.object,
  level: PropTypes.number
};

export default NavCollapse;
