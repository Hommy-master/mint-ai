import { useState } from 'react'
import './RightBar.css'
import rightbarTel from '../../assets/common/rightbar_tel.svg'
import rightbarTry from '../../assets/common/rightbar_try.svg'
import rightbarWx from '../../assets/common/rightbar_wx.svg'
import serviceAvatar from '../../assets/common/service_avatar.png'
import { DEFAULT_CUSTOMER_SERVICE_PHONE, DefaultQrCodeUrl, G_TrialUrl } from '../../utils/const'

function RightBar() {
  const [showQrCode, setShowQrCode] = useState(false)

  return (
    <div className="rightBar">
      {/* H5 版本 */}
      <div className="rightBar-h5">
        <a className="rightBar-h5-item" href={`tel:${DEFAULT_CUSTOMER_SERVICE_PHONE}`}>
          <img src={rightbarTel} alt="电话" />
        </a>
        <a 
          className="rightBar-h5-item" 
          target="_blank" 
          href={G_TrialUrl}
          rel="noreferrer"
        >
          <img src={rightbarTry} alt="试用" />
        </a>
        <div className="rightBar-h5-item" onClick={() => setShowQrCode(!showQrCode)}>
          <img src={rightbarWx} alt="微信" />
        </div>
      </div>

      {/* PC 版本 */}
      <div className="rightBar-pc">
        <img src={serviceAvatar} className="rightBar-pc-avatar" alt="客服" />
        <div className="rightBar-pc-telLabel">
          <img src={rightbarTel} alt="" /> 电话咨询
        </div>
        <a className="rightBar-pc-tel" href={`tel:${DEFAULT_CUSTOMER_SERVICE_PHONE}`}>{DEFAULT_CUSTOMER_SERVICE_PHONE}</a>
        <a 
          className="rightBar-pc-tryBtn" 
          target="_blank" 
          href={G_TrialUrl}
          rel="noreferrer"
        >
          申请试用
        </a>
        <div className="rightBar-pc-dot"></div>
        <div className="rightBar-pc-wx" onMouseEnter={() => setShowQrCode(true)} onMouseLeave={() => setShowQrCode(false)}>
          <img src={rightbarWx} alt="" /> 官方微信
        </div>
      </div>

      {/* 二维码弹窗 */}
      <div className="rightBar-qrcode" style={{ display: showQrCode ? 'block' : 'none' }}>
        <img className="rightBar-qrcode-img" src={DefaultQrCodeUrl} alt="微信二维码" />
        <div className="rightBar-qrcode-label">微信扫一扫</div>
      </div>
    </div>
  )
}

export default RightBar
