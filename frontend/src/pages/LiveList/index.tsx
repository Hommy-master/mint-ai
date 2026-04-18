import React, { useState } from 'react';
import tipIcon from '@assets/tip-icon.png';
import noDataImage from '@assets/no-data.png';
import './index.less';

interface LiveRecord {
  id: string;
  title: string;
  duration: string;
  source: string;
  status: 'recording' | 'completed' | 'failed';
  addTime: string;
}

const LiveList: React.FC = () => {
  const [searchUrl, setSearchUrl] = useState('');
  const [records] = useState<LiveRecord[]>([]);
  const [currentPage] = useState(1);
  const [tipVisible, setTipVisible] = useState(false);

  const handleStartRecord = () => {
    if (!searchUrl.trim()) {
      alert('请输入直播链接');
      return;
    }
    console.log('开始录制:', searchUrl);
  };

  const getStatusText = (status: LiveRecord['status']) => {
    const statusMap = {
      recording: '录制中',
      completed: '已完成',
      failed: '失败',
    };
    return statusMap[status];
  };

  const getStatusClass = (status: LiveRecord['status']) => {
    const classMap = {
      recording: 'live-status-recording',
      completed: 'live-status-completed',
      failed: 'live-status-failed',
    };
    return classMap[status];
  };

  return (
    <div className="live">
      <div className="live-search">
        <input
          type="text"
          className="live-search-input"
          placeholder="输入要粘贴的直播链接"
          value={searchUrl}
          onChange={(e) => setSearchUrl(e.target.value)}
        />
        <div className="live-search-actions">
          <div className="live-record-btn" onClick={handleStartRecord}>
            <span className="live-record-text">一键成片</span>
          </div>
        </div>
        <div className="live-search-tip">
          请自觉遵守平台链接导入规范
          <img src={tipIcon} className="live-tip-icon" alt="提示" onClick={() => setTipVisible(true)} />
        </div>
      </div>

      {tipVisible && (
        <div className="live-modal-mask" onClick={() => setTipVisible(false)}>
          <div className="live-modal" onClick={(e) => e.stopPropagation()}>
            <div className="live-modal-title">温馨提示</div>
            <div className="live-modal-content">
              坚持创作高质量且充满人文关怀的原创内容，请勿搬运或发布侵权他人、违反国家法律法规、公序良俗的不良内容；因违反上述规定而产生的一切后果，均由用户自行承担。
            </div>
            <div className="live-modal-btn" onClick={() => setTipVisible(false)}>
              <span className="live-modal-btn-text">我知道了</span>
            </div>
          </div>
        </div>
      )}

      <div className="live-table-section">
        <div className="live-table-container">
          <table className="live-table">
            <thead>
              <tr>
                <th className="live-col-info">直播信息</th>
                <th className="live-col-duration">录制时长</th>
                <th className="live-col-source">直播来源</th>
                <th className="live-col-status">录制状态</th>
                <th className="live-col-time">添加时间</th>
                <th className="live-col-action">更多操作</th>
              </tr>
            </thead>
            <tbody>
              {records.length > 0 ? (
                records.map((record) => (
                  <tr key={record.id}>
                    <td>{record.title}</td>
                    <td>{record.duration}</td>
                    <td>{record.source}</td>
                    <td>
                      <span className={`live-status-tag ${getStatusClass(record.status)}`}>
                        {getStatusText(record.status)}
                      </span>
                    </td>
                    <td>{record.addTime}</td>
                    <td>
                      <button className="live-action-btn">查看</button>
                    </td>
                  </tr>
                ))
              ) : null}
            </tbody>
          </table>
          
          {records.length === 0 && (
            <div className="live-empty">
              <img src={noDataImage} alt="无直播记录" className="live-empty-img" />
              <div className="live-empty-text">暂无直播记录</div>
            </div>
          )}
        </div>
      </div>

      <div className="live-pagination">
        <button className="live-page-btn prev" disabled={currentPage === 1}>
          &lt;
        </button>
        <span className="live-page-number active">1</span>
        <button className="live-page-btn next" disabled>
          &gt;
        </button>
      </div>
    </div>
  );
};

export default LiveList;
