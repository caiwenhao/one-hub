import PropTypes from 'prop-types';
import { Typography } from '@mui/material';
import { Link, useLocation } from 'react-router-dom';
import { useTheme } from '@mui/material/styles';

export default function NavLinkItem({ to, label, matchStart = false }) {
  const { pathname } = useLocation();
  const theme = useTheme();
  const isActive = matchStart ? pathname.startsWith(to) : pathname === to;
  const activeColor = theme.palette.primary.main;

  return (
    <Typography
      component={Link}
      to={to}
      sx={{
        color: isActive ? activeColor : '#718096',
        fontWeight: isActive ? 600 : 500,
        textDecoration: 'none',
        fontSize: '1rem',
        cursor: 'pointer',
        position: 'relative',
        transition: 'all 0.3s ease',
        '&:hover': { color: activeColor },
        '&::after': {
          content: '""',
          position: 'absolute',
          left: 0,
          right: 0,
          bottom: 0,
          height: '2px',
          background: activeColor,
          transform: isActive ? 'scaleX(1)' : 'scaleX(0)',
          transformOrigin: 'center',
          transition: 'transform 0.3s ease'
        },
        '&:hover::after': { transform: 'scaleX(1)' }
      }}
    >
      {label}
    </Typography>
  );
}

NavLinkItem.propTypes = {
  to: PropTypes.string.isRequired,
  label: PropTypes.node.isRequired,
  matchStart: PropTypes.bool
};

