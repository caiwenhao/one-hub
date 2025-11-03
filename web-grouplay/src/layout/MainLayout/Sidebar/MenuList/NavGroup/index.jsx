import PropTypes from 'prop-types';

// material-ui
import { useTheme } from '@mui/material/styles';
import { Divider, List, ListSubheader, Typography, Stack, Box } from '@mui/material';
import { alpha } from '@mui/material/styles';

// project imports
import NavItem from '../NavItem';
import NavCollapse from '../NavCollapse';

// ==============================|| SIDEBAR MENU LIST GROUP ||============================== //

const NavGroup = ({ item, isLast }) => {
  const theme = useTheme();

  // menu list collapse & items
  const items = item.children?.map((menu) => {
    switch (menu.type) {
      case 'collapse':
        return <NavCollapse key={menu.id} menu={menu} level={1} />;
      case 'item':
        return <NavItem key={menu.id} item={menu} level={1} />;
      default:
        return (
          <Typography key={menu.id} variant="h6" color="error" align="center">
            Menu Items Error
          </Typography>
        );
    }
  });

  return (
    <>
      <List
        disablePadding
        subheader={
          item.title ? (
            <ListSubheader
              component="div"
              disableSticky
              sx={{
                px: 2,
                pt: 2,
                pb: 1.5,
                backgroundColor: 'transparent'
              }}
            >
              <Stack spacing={0.5} alignItems="flex-start">
                <Stack direction="row" alignItems="center" spacing={1}>
                  <Box
                    sx={{
                      width: 8,
                      height: 8,
                      borderRadius: '50%',
                      backgroundColor: alpha(theme.palette.primary.main, theme.palette.mode === 'dark' ? 0.4 : 0.25)
                    }}
                  />
                  <Typography
                    variant="overline"
                    sx={{
                      fontSize: '0.72rem',
                      letterSpacing: '0.08em',
                      color: theme.palette.text.secondary
                    }}
                  >
                    {item.title}
                  </Typography>
                </Stack>
                {item.caption && (
                  <Typography
                    component="span"
                    sx={{
                      fontSize: '0.72rem',
                      fontWeight: 400,
                      letterSpacing: 'normal',
                      textTransform: 'none',
                      color: theme.palette.text.disabled,
                      pl: 2.5
                    }}
                  >
                    {item.caption}
                  </Typography>
                )}
              </Stack>
            </ListSubheader>
          ) : undefined
        }
        sx={{
          pb: 1.5
        }}
      >
        {items}
      </List>

      {/* group divider，末组不渲染以避免多余间隔 */}
      {!isLast && <Divider sx={{ mx: 2, my: 0.5, borderColor: theme.palette.divider }} />}
    </>
  );
};

NavGroup.propTypes = {
  item: PropTypes.object,
  isLast: PropTypes.bool
};

export default NavGroup;
