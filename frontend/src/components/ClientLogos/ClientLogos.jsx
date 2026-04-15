import './ClientLogos.css'
import logo1 from '../../assets/clients/logo_jmw.png'
import logo2 from '../../assets/clients/logo_yj.png'
import logo3 from '../../assets/clients/logo_yhnz.png'
import logo4 from '../../assets/clients/logo_ysy.png'

const logosRow1 = [logo1, logo2, logo3, logo4, logo1, logo2, logo3, logo4]
const logosRow2 = [logo2, logo3, logo4, logo1, logo2, logo3, logo4, logo1]
const logosRow3 = [logo3, logo4, logo1, logo2, logo3, logo4, logo1, logo2]

function ClientLogos() {
  return (
    <div className="client_logo-wrap">
      <div className="client_logo-list">
        <div className="client_logo-row row-1">
          {[...logosRow1, ...logosRow1].map((logo, index) => (
            <img key={index} src={logo} className="client_logo-item" alt="" />
          ))}
        </div>
        <div className="client_logo-row row-2">
          {[...logosRow2, ...logosRow2].map((logo, index) => (
            <img key={index} src={logo} className="client_logo-item" alt="" />
          ))}
        </div>
        <div className="client_logo-row row-3">
          {[...logosRow3, ...logosRow3].map((logo, index) => (
            <img key={index} src={logo} className="client_logo-item" alt="" />
          ))}
        </div>
      </div>
      <div className="client_logo-mask"></div>
    </div>
  )
}

export default ClientLogos
