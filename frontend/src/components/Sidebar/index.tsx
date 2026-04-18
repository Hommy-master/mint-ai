import React from 'react';
import logo from '/logo.png';
import recordIcon from '@assets/record.png';
import './index.less';

const Sidebar: React.FC = () => {
  return (
    <aside className="sidebar">
      <div className="sidebar-logo">
        <img src={logo} alt="直播快剪" />
        <span>直播快剪</span>
      </div>
      <nav className="sidebar-menu-wrapper">
        <ul className="sidebar-menu" role="menubar">
          <li className="sidebar-menu-item" role="menuitem">
            <img src={recordIcon} className="sidebar-menu-icon" alt="录制" />
            <span className="sidebar-menu-text">我的直播</span>
          </li>
        </ul>
      </nav>
    </aside>
  );
};

export default Sidebar;
