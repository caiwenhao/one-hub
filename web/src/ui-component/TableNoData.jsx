import PropTypes from 'prop-types';
import { TableRow, TableCell } from '@mui/material';
import EmptyState from './EmptyState';

const TableNoData = ({ message = '暂无数据' }) => {
  return (
    <TableRow>
      <TableCell colSpan={1000}>
        <EmptyState title={message} description={null} sx={{ minHeight: 240 }} />
      </TableCell>
    </TableRow>
  );
};
export default TableNoData;

TableNoData.propTypes = {
  message: PropTypes.string
};
