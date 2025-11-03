import { createRoot } from 'react-dom/client';

// third party
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';

// project imports
import App from 'App';
import { store } from 'store';

// style + assets
import './styles/tokens.css';
import './styles/density.css';
import './styles/table-helpers.css';
import 'assets/scss/style.scss';
import './styles/motion.css';
import './styles/i18n.css';
import { applyTokens } from 'design/applyTokens';
import { applyDensity } from 'design/density';
import config from './config';
import reportWebVitals from 'reportWebVitals';
// ==============================|| REACT DOM RENDER  ||============================== //

const container = document.getElementById('root');
// 注入设计令牌 CSS 变量（企业稳重风格）
applyTokens();
// 应用密度（本地持久化，默认 standard）
applyDensity();
const root = createRoot(container); // createRoot(container!) if you use TypeScript
root.render(
  <Provider store={store}>
    <BrowserRouter basename={config.basename}>
      <App />
    </BrowserRouter>
  </Provider>
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
reportWebVitals();
