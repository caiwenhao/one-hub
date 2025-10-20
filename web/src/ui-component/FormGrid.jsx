// 标准表单栅格：单列/双列与断点切换；标签顶部对齐
import { Grid } from '@mui/material';

export default function FormGrid({ children, columns = 1, spacing = 2 }) {
  const col = Math.max(1, Math.min(2, columns));
  if (col === 1)
    return (
      <Grid container spacing={spacing}>
        {children}
      </Grid>
    );
  return (
    <Grid container spacing={spacing}>
      {Array.isArray(children)
        ? children.map((child, idx) => (
            <Grid item xs={12} md={6} key={idx}>
              {child}
            </Grid>
          ))
        : children}
    </Grid>
  );
}
