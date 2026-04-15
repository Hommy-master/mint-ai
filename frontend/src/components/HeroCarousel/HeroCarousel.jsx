import { useState, useEffect } from 'react'
import './HeroCarousel.css'
import banner1 from '../../assets/hero/banner_1.png'
import banner2 from '../../assets/hero/banner_2.png'
import banner3 from '../../assets/hero/banner_3.png'
import { DEFAULT_FUNCTION_NAME, G_TrialUrl } from '../../utils/const'

const heroData = [
  {
    id: 1,
    title: DEFAULT_FUNCTION_NAME,
    desc: '自动去除直播贴片，智能发现服装讲解亮点，多版本切片覆盖多平台内容要求。引流宝让直播卖家快速实现更多平台客户的覆盖',
    image: banner1,
  },
  {
    id: 2,
    title: DEFAULT_FUNCTION_NAME,
    desc: '实现直播内容自动化卖点切片，突破抖音快手短视频电商素材瓶颈。引流宝助力每一个服装直播间实现优质内容获客',
    image: banner2,
  },
  {
    id: 3,
    title: DEFAULT_FUNCTION_NAME,
    desc: '一边直播一边发短视频，不再是大直播间的专属，不再依赖中控的调度。引流宝真正做到"傻瓜"式直播短视频切片',
    image: banner3,
  },
]

function HeroCarousel() {
  const [currentSlide, setCurrentSlide] = useState(0)

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentSlide((prev) => (prev + 1) % heroData.length)
    }, 5000)
    return () => clearInterval(timer)
  }, [])

  return (
    <section className="hero_carousel">
      <div className="hero_carousel-wrap">
        <div className="hero_carousel-slider slick-slider">
          <div className="slick-list">
            <div 
              className="slick-track"
              style={{ 
                transform: `translateX(-${currentSlide * 33.333}%)`,
                transition: 'transform 0.5s ease'
              }}
            >
              {heroData.map((item) => (
                <div 
                  key={item.id}
                  className="hero_item slick-slide"
                >
                  <img
                    className="hero-img"
                    alt={item.title}
                    src={item.image}
                  />
                  <div className="hero-content">
                    <div className="hero-title">{item.title}</div>
                    <div className="hero-desc">{item.desc}</div>
                    <div className="hero-btns">
                      <a
                        className="hero-btn"
                        target="_blank"
                        href={G_TrialUrl}
                        rel="noreferrer"
                      >
                        免费试用
                      </a>
                    </div>
                  </div>
                  <img
                    className="hero-pcimg"
                    src={item.image}
                    alt={item.title}
                  />
                </div>
              ))}
            </div>
          </div>

          {/* 指示器 */}
          <ul className="slick-dots">
            {heroData.map((_, index) => (
              <li 
                key={index}
                className={currentSlide === index ? 'slick-active' : ''}
                onClick={() => setCurrentSlide(index)}
              >
                <button>
                  {index + 1}
                </button>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </section>
  )
}

export default HeroCarousel
