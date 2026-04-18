import React from 'react';
import Sidebar from '../Sidebar';
// import Header from '../Header';
import './index.less';

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className="layout">
      <Sidebar />
      <div className="layout-right">
        {/* <Header /> */}
        <main className="layout-main">
          {children}
        </main>
      </div>
    </div>
  );
};

export default Layout;
