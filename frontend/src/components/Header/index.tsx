import React from 'react';
import memberAvatar from '@assets/member.png';
import './index.less';

const Header: React.FC = () => {
  return (
    <header className="header">
      <div className="header-actions">
        <div className="header-download-btn">
          <span className="header-download-text">下载App</span>
        </div>
        <div className="header-vip-btn">
          <span className="header-vip-text">会员中心</span>
        </div>
        <div className="header-avatar">
          <img src={memberAvatar} className="header-avatar-img" alt="用户头像" />
        </div>
      </div>
    </header>
  );
};

export default Header;
