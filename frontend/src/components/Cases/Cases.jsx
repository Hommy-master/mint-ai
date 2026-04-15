import { useState, useEffect, useRef, useCallback } from 'react'
import logo1 from '../../assets/clients/logo_jmw.png'
import logo2 from '../../assets/clients/logo_yj.png'
import logo3 from '../../assets/clients/logo_yhnz.png'
import logo4 from '../../assets/clients/logo_ysy.png'
import img1 from '../../assets/cases/img_jmw.png'
import img2 from '../../assets/cases/img_yj.png'
import img3 from '../../assets/cases/img_yhnz.png'
import img4 from '../../assets/cases/img_ysy.png'
import ClientLogos from '../ClientLogos/ClientLogos'
import './Cases.css'

const casesData = [
  {
    id: 1,
    logo: logo1,
    image: img1,
    title: '九牧王',
    desc: '极睿科技与全球销量领先的男裤专家九牧王正式达成合作。基于极睿科技AIGC在电商领域的先进落地应用能力，引流宝批量生成直播高光切片，以矩阵式AIGC产品为九牧王至少降低70%人力运营成本、提升50%内容引流效率，以优质内容实现高效引流获客，用好内容成就好GMV'
  },
  {
    id: 2,
    logo: logo2,
    image: img2,
    title: '赢家',
    desc: '国内知名女装品牌赢家与极睿引流宝正式达成合作。通过AI直播切片，引流宝为赢家生成近5000条优质商品卖点短视频，助力其直播业务流量快速崛起，成为抖音女装类目中头部直播间，实现直播月场观人数破千万，线上年销售额破亿。'
  },
  {
    id: 3,
    logo: logo3,
    image: img3,
    title: '金辉男装',
    desc: '快手男装商家"金辉男装"与极睿引流宝达成合作。通过行业领先的AIGC在电商直播切片领域的变革性应用，极睿引流宝助力"金辉男装"产出切片540条，日均短视频发布数量提升4倍，多条智能切片入选账户精选视频，其中单条精选视频曝光高达37.8万自然流量、单条最高点赞量为5184，账号一个月增长精准粉丝6万。极睿引流宝赋能"金辉男装"以优质短视频内容快速提升营销GMV！'
  },
  {
    id: 4,
    logo: logo4,
    image: img4,
    title: '妍三也',
    desc: '网红服饰女装品牌妍三也与极睿引流宝正式开展合作。仅花费一个月左右时间，通过对直播内容的智能化识别与自动化裁剪，"短视频智能生成工具-引流宝"为其生成优质商品卖点短视频2000多条，以优质高频的短视频内容赋能该客户快速吸粉、抢占市场流量先机，提升直播间商品转化。'
  }
]

function Cases() {
  const [currentIndex, setCurrentIndex] = useState(0)
  const intervalRef = useRef(null)

  const handlePrev = () => {
    setCurrentIndex((prev) => (prev === 0 ? casesData.length - 1 : prev - 1))
  }

  const handleNext = useCallback(() => {
    setCurrentIndex((prev) => (prev === casesData.length - 1 ? 0 : prev + 1))
  }, [])

  // 自动轮播
  const startAutoPlay = useCallback(() => {
    intervalRef.current = setInterval(() => {
      handleNext()
    }, 5000)
  }, [handleNext])

  const stopAutoPlay = useCallback(() => {
    if (intervalRef.current) {
      clearInterval(intervalRef.current)
      intervalRef.current = null
    }
  }, [])

  useEffect(() => {
    startAutoPlay()
    return () => stopAutoPlay()
  }, [startAutoPlay, stopAutoPlay])

  return (
    <section className="pro-customer_case section" id="case">
      <div className="container-wrap">
        <div className="part_title">客户案例</div>
      </div>
      
      {/* H5 案例列表 - 横向滚动 */}
      <div className="pro-customer_case-h5list">
        {casesData.map((item) => (
          <div key={item.id} className="pro-customer_case-h5item">
            <div className="pro-customer_case-h5item-top">
              <img src={item.image} alt="" />
            </div>
            <div className="pro-customer_case-h5item-bottom">
              <div className="pro-customer_case-h5item-t">{item.title}</div>
              <div className="pro-customer_case-h5item-d">{item.desc}</div>
            </div>
          </div>
        ))}
      </div>

      {/* PC 案例轮播 */}
      <div 
        className="pro-customer_case-list"
        onMouseEnter={stopAutoPlay}
        onMouseLeave={startAutoPlay}
      >
        <div className="pro-customer_case-list-prev" onClick={handlePrev}>
          <svg width="16" height="28" viewBox="0 0 16 28" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M14 26L2 14L14 2" stroke="#C9CDD4" strokeWidth="4" strokeLinecap="round" strokeLinejoin="round"></path>
          </svg>
        </div>
        
        <div className="pro-customer_case-slider">
          {casesData.map((item, index) => (
            <div 
              key={item.id} 
              className={`pro-customer_case-item ${index === currentIndex ? 'active' : ''}`}
              style={{ transform: `translate3d(${(index - currentIndex) * 100}%, 0, 0)` }}
            >
              <div className="pro-customer_case-item-wrap">
                <div className="pro-customer_case-item-left">
                  <img className="pro-customer_case-item-logo" src={item.logo} alt="" />
                  <div className="pro-customer_case-item-desc">{item.desc}</div>
                </div>
                <img className="pro-customer_case-item-img" src={item.image} alt="" />
              </div>
            </div>
          ))}
        </div>

        <div className="pro-customer_case-list-next" onClick={handleNext}>
          <svg width="16" height="28" viewBox="0 0 16 28" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M2 26L14 14L2 2" stroke="#C9CDD4" strokeWidth="4" strokeLinecap="round" strokeLinejoin="round"></path>
          </svg>
        </div>
      </div>
      
      <ClientLogos />
    </section>
  )
}

export default Cases
