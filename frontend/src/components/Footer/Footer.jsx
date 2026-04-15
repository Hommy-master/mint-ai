
import { DEFAULT_CUSTOMER_SERVICE_PHONE, DEFAULT_CUSTOMER_SERVICE_PHONE2, DEFAULT_EMAIL, DEFAULT_EMAIL2, DefaultQrCodeUrl } from '../../utils/const'
import './Footer.css'

function Footer() {
  return (
    <footer className="footer">
      <div className="footer-wrap">
        {/* 快速入口 */}
        {/* <div className="footer-item">
          <div className="footer-item-title">快速入口</div>
          <div className="footer-item-content">
            <div className="footer-item-list">
              <div className="footer-item-link">
                <a target="_blank" rel="noreferrer" href="https://www.infimind.com/">北京极睿科技有限责任公司</a>
              </div>
              <div className="footer-item-link">
                <a target="_blank" rel="noreferrer" href="http://www.icut.cn/">引流宝 短视频智能生成平台</a>
              </div>
              <div className="footer-item-link">
                <a target="_blank" rel="noreferrer" href="http://www.ipim.cn/">IPim 商品数字化管理中台</a>
              </div>
            </div>
          </div>
        </div> */}

        {/* 服务与支持 */}
        <div className="footer-item">
          <div className="footer-item-title">服务与支持</div>
          <div className="footer-item-content">
            <div className="footer-item-list">
              <div className="footer-item-link">
                <span>售后服务：</span>
                <a target="_blank" rel="noreferrer" href={`mailto:${DEFAULT_EMAIL}`}>{DEFAULT_EMAIL}</a>
              </div>
              <div className="footer-item-link">
                <span>商务合作：</span>
                <a target="_blank" rel="noreferrer" href={`mailto:${DEFAULT_EMAIL2}`}>{DEFAULT_EMAIL2}</a>
              </div>
              <div className="footer-item-link">
                <span>工作时间：9:00-18:00（工作日）</span>
              </div>
            </div>
          </div>
        </div>

        {/* 联系我们 */}
        <div className="footer-item">
          <div className="footer-item-title">联系我们</div>
          <div className="footer-item-content">
            <div className="footer-item-list">
              <div className="footer-item-link">
                <span>客服电话1：</span>
                <a target="_blank" rel="noreferrer" href={`tel:${DEFAULT_CUSTOMER_SERVICE_PHONE}`}>{DEFAULT_CUSTOMER_SERVICE_PHONE}</a>
              </div>
              <div className="footer-item-link">
                <span>客服电话2：</span>
                <a target="_blank" rel="noreferrer" href={`tel:${DEFAULT_CUSTOMER_SERVICE_PHONE2}`}>{DEFAULT_CUSTOMER_SERVICE_PHONE2}</a>
              </div>
              <div className="footer-item-qrcode">
                <img className="footer-item-qrcodeImg" alt="" src={DefaultQrCodeUrl} />
                <div className="footer-item-qrcodeText">微信扫码关注</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* 备案信息 */}
      {/* <div className="footer-wrap">
        <a href="https://beian.miit.gov.cn/" className="footer-beian">京ICP备17058102号</a>
      </div> */}
    </footer>
  )
}

export default Footer
