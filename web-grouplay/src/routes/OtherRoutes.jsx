import { lazy } from 'react';
import { Box } from '@mui/material';

// project imports
import Loadable from 'ui-component/Loadable';
import MinimalLayout from 'layout/MinimalLayout';

// login option 3 routing
const AuthLogin = Loadable(lazy(() => import('views/Authentication/Auth/Login')));
const AuthRegister = Loadable(lazy(() => import('views/Authentication/Auth/Register')));
const GitHubOAuth = Loadable(lazy(() => import('views/Authentication/Auth/GitHubOAuth')));
const LarkOAuth = Loadable(lazy(() => import('views/Authentication/Auth/LarkOAuth')));
const OIDCOAuth = Loadable(lazy(() => import('views/Authentication/Auth/OIDCOAuth')));
const ForgetPassword = Loadable(lazy(() => import('views/Authentication/Auth/ForgetPassword')));
const ResetPassword = Loadable(lazy(() => import('views/Authentication/Auth/ResetPassword')));
const Home = Loadable(lazy(() => import('views/Home')));
const NotFoundView = Loadable(lazy(() => import('views/Error')));
const Jump = Loadable(lazy(() => import('views/Jump')));
const Playground = Loadable(lazy(() => import('views/Playground')));
const ModelPrice = Loadable(lazy(() => import('views/ModelPrice')));
const PricingPage = Loadable(lazy(() => import('views/PricingPage')));
const HotModels = Loadable(lazy(() => import('views/HotModels')));
const DeveloperCenter = Loadable(lazy(() => import('views/DeveloperCenter')));
const Contact = Loadable(lazy(() => import('views/Contact')));

const WithMargins = ({ children }) => (
  <Box
    sx={{
      maxWidth: '1200px',
      margin: '0 auto',
      padding: { xs: 0, sm: '0 24px' }
    }}
  >
    {children}
  </Box>
);

// ==============================|| AUTHENTICATION ROUTING ||============================== //

const OtherRoutes = {
  path: '/',
  element: <MinimalLayout />,
  children: [
    {
      path: '',
      element: <Home />
    },
    {
      path: '/login',
      element: <AuthLogin />
    },
    {
      path: '/register',
      element: <AuthRegister />
    },
    {
      path: '/reset',
      element: <ForgetPassword />
    },
    {
      path: '/user/reset',
      element: <ResetPassword />
    },
    {
      path: '/oauth/github',
      element: <GitHubOAuth />
    },
    {
      path: '/oauth/oidc',
      element: <OIDCOAuth />
    },
    {
      path: '/oauth/lark',
      element: <LarkOAuth />
    },
    {
      path: '/404',
      element: <NotFoundView />
    },
    {
      path: '/jump',
      element: <Jump />
    },
    {
      path: '/playground',
      element: <Playground />
    },
    {
      path: '/price',
      element: (
        <WithMargins>
          <PricingPage />
        </WithMargins>
      )
    },
    {
      path: '/models',
      element: (
        <WithMargins>
          <HotModels />
        </WithMargins>
      )
    },
    {
      path: '/developer',
      element: <DeveloperCenter />
    },
    {
      path: '/contact',
      element: <Contact />
    }
  ]
};

export default OtherRoutes;
