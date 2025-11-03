// 基于 design/tokens 生成 MUI 主题期望的 colors 结构
import { tokens } from 'design/tokens';

const withAlpha = (hex, alpha) => {
  // 简易 rgba 转换（不校验输入），用于 divider/hover 背景
  const bigint = parseInt(hex.replace('#', ''), 16);
  const r = (bigint >> 16) & 255;
  const g = (bigint >> 8) & 255;
  const b = bigint & 255;
  return `rgba(${r}, ${g}, ${b}, ${alpha})`;
};

export function buildColorsFromTokens(mode = 'light') {
  const brand = tokens.color.brand;
  const grey = tokens.color?.grey || {
    50: '#FAFAFA',
    100: '#F5F5F5',
    200: '#EEEEEE',
    300: '#E0E0E0',
    500: '#9E9E9E',
    600: '#757575',
    700: '#616161',
    900: '#212121'
  };

  const primary = {
    light: '#E6F0FF',
    200: '#ADC6FF',
    main: brand.primary,
    dark: brand.primaryDark,
    800: '#1D39C4'
  };

  // 次要色：采用更稳重的中性色为默认 secondary
  const secondary = {
    light: '#F3F4F6',
    200: '#E5E7EB',
    main: '#1F2937',
    dark: '#111827',
    800: '#0F172A'
  };

  // 语义色直接映射 tokens
  const success = {
    light: '#E6F8F1',
    200: '#66D9AB',
    main: tokens.color.success,
    dark: '#0AA366'
  };
  const error = {
    light: '#FFEBEE',
    main: tokens.color.error,
    dark: '#CF1322'
  };
  const orange = {
    light: '#FFF3E0',
    main: '#FF9800',
    dark: '#E65100'
  };
  const warning = {
    light: '#FFF8E1',
    main: tokens.color.warning,
    dark: '#D48806'
  };

  // 背景/文本
  const bg = {
    paper: tokens.color.bg.card,
    default: tokens.color.bg.page
  };

  // 暗色模式下适度调整
  if (mode === 'dark') {
    bg.paper = '#1A1D23';
    bg.default = '#13151A';
  }

  return {
    // paper & background
    paper: bg.paper,
    // primary
    primaryLight: primary.light,
    primaryMain: primary.main,
    primaryDark: primary.dark,
    primary200: primary[200],
    primary800: primary[800],
    // secondary
    secondaryLight: secondary.light,
    secondaryMain: secondary.main,
    secondaryDark: secondary.dark,
    secondary200: secondary[200],
    secondary800: secondary[800],
    // success
    successLight: success.light,
    success200: success[200],
    successMain: success.main,
    successDark: success.dark,
    // error
    errorLight: error.light,
    errorMain: error.main,
    errorDark: error.dark,
    // orange
    orangeLight: orange.light,
    orangeMain: orange.main,
    orangeDark: orange.dark,
    // warning
    warningLight: warning.light,
    warningMain: warning.main,
    warningDark: warning.dark,
    // grey
    grey50: grey[50],
    grey100: grey[100],
    grey200: grey[200],
    grey300: grey[300],
    grey500: grey[500],
    grey600: grey[600],
    grey700: grey[700],
    grey900: grey[900],
    // dark theme variants
    darkBackground: '#101820',
    darkPaper: '#1E2A38',
    darkLevel1: '#131E29',
    darkLevel2: '#1A2736',
    darkTableHeader: '#253545',
    darkPrimaryLight: '#60B8FF',
    darkPrimaryMain: brand.primary,
    darkPrimaryDark: brand.primaryDark,
    darkPrimary200: '#91CAFF',
    darkPrimary800: '#004D85',
    darkSecondaryLight: '#B4A1FF',
    darkSecondaryMain: '#7B61FF',
    darkSecondaryDark: '#5A3FD6',
    darkSecondary200: '#D4C8FF',
    darkSecondary800: '#4526BF',
    darkDivider: withAlpha('#F0F7FF', 0.12),
    darkSelectedBack: withAlpha(brand.primary, 0.12),
    tableBackground: '#F8FAFC',
    tableBorderBottom: mode === 'dark' ? withAlpha('#FFFFFF', 0.08) : '#E9EDF5'
  };
}

export default buildColorsFromTokens;
