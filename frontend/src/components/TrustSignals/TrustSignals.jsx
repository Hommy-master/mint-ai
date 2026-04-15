import './TrustSignals.css'
import signal1 from '../../assets/common/signal_1.png'
import signal2 from '../../assets/common/signal_2.png'
import signal3 from '../../assets/common/signal_3.png'

const signals = [
  {
    icon: signal1,
    title: '头部客户见证',
    desc: '各行业龙头企业推广技术采购订单过亿'
  },
  {
    icon: signal2,
    title: '顶级资本认同',
    desc: '累计获顶尖资本投资近N亿'
  },
  {
    icon: signal3,
    title: '斩获国赛冠军',
    desc: '加推推广技术斩获国赛冠军'
  }
]

function TrustSignals() {
  return (
    <div className="trust_signals">
      {signals.map((signal, index) => (
        <div className="t_s_item" key={index}>
          <img className="t_s_icon" alt={signal.title} src={signal.icon} />
          <div className="t_s_content">
            <div className="t_s_title">{signal.title}</div>
            <div className="t_s_desc">{signal.desc}</div>
          </div>
        </div>
      ))}
    </div>
  )
}

export default TrustSignals
