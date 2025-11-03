import { useMemo } from 'react';
// material-ui
import { Typography } from '@mui/material';

// project imports
import NavGroup from './NavGroup';
import menuItem from 'menu-items';
import { useIsAdmin } from 'utils/common';
import { useTranslation } from 'react-i18next';
import { useSelector } from 'react-redux';

// ==============================|| SIDEBAR MENU LIST ||============================== //
const MenuList = () => {
  const userIsAdmin = useIsAdmin();
  const { t } = useTranslation();
  const siteInfo = useSelector((state) => state.siteInfo);

  const groups = useMemo(() => {
    // 递归转换：本地化标题 + 权限过滤
    const formatNodes = (nodes = []) =>
      nodes.reduce((acc, node) => {
        // 非管理员隐藏管理员菜单；缺少发票权限时隐藏发票
        if ((node.isAdmin && !userIsAdmin) || (siteInfo?.UserInvoiceMonth === false && node.id === 'invoice')) {
          return acc;
        }

        const localizedNode = {
          ...node,
          title: t(node.id, { defaultValue: node.title })
        };

        if (node.children?.length) {
          const nextChildren = formatNodes(node.children);
          if (!nextChildren.length) {
            // 折叠菜单在子项被过滤完时直接剔除
            if (node.type === 'collapse') {
              return acc;
            }
            localizedNode.children = [];
          } else {
            localizedNode.children = nextChildren;
          }
        }

        acc.push(localizedNode);
        return acc;
      }, []);

    return formatNodes(menuItem.items).filter((group) => group.type === 'group' && group.children?.length);
  }, [siteInfo?.UserInvoiceMonth, t, userIsAdmin]);

  if (!groups.length) {
    return (
      <Typography variant="h6" color="error" align="center">
        {t('menu.error')}
      </Typography>
    );
  }

  return (
    <>
      {groups.map((group, index) => (
        <NavGroup key={group.id} item={group} isLast={index === groups.length - 1} />
      ))}
    </>
  );
};

export default MenuList;
