import { G_TrialUrl } from '../../utils/const'
import './BottomTryBar.css'

function BottomTryBar() {
  return (
    <div className="bottom-try-bar">
      <p className="bottom-try-bar-title">快速开启您的数字化增长之路</p>
      <a 
        className="primary-btn" 
        target="_blank" 
        href={G_TrialUrl}
        rel="noreferrer"
      >
        预约演示
      </a>
    </div>
  )
}

export default BottomTryBar
