import { useState } from 'react'
import './PropositionPart.css'
import p1IconActive from '../../assets/products/p_1_icon_active.png'
import p1Icon from '../../assets/products/p_1_icon.png'
import p2IconActive from '../../assets/products/p_2_icon_active.png'
import p2Icon from '../../assets/products/p_2_icon.png'
import p3IconActive from '../../assets/products/p_3_icon_active.png'
import p3Icon from '../../assets/products/p_3_icon.png'
import p4IconActive from '../../assets/products/p_4_icon_active.png'
import p4Icon from '../../assets/products/p_4_icon.png'
import p1 from '../../assets/products/p_1.png'
import p2 from '../../assets/products/p_2.png'
import p3 from '../../assets/products/p_3.png'
import p4 from '../../assets/products/p_4.png'


const icons = [
  { active: p1IconActive, inactive: p1Icon },
  { active: p2IconActive, inactive: p2Icon },
  { active: p3IconActive, inactive: p3Icon },
  { active: p4IconActive, inactive: p4Icon }
]

const contentList = [
  {
    img: p1,
    title: '以直播内容生成更多短视频',
    desc: '每一条视频都能带来平台自然流量，更多数量才能更多流量'
  },
  {
    img: p2,
    title: '智能识别商品讲解亮点',
    desc: '自动发现直播中的商品卖点，精准切片'
  },
  {
    img: p3,
    title: '多版本切片覆盖多平台',
    desc: '适配抖音、快手、淘宝等多平台内容要求'
  },
  {
    img: p4,
    title: '傻瓜式操作无需专业技能',
    desc: '简单几步即可完成专业级短视频制作'
  }
]

function PropositionPart() {
  const [activeIndex, setActiveIndex] = useState(0)

  return (
    <section className="proposition_part">
      <div className="container">
        <div className="part_title">使用引流宝，直播切片自动化生成，多、快、好、省！</div>
        <div className="part_desc">产品和服务让团队直播业务省心省力，助力品牌生成更多更好的切片短视频，斩获更好的 GMV</div>
        <div className="proposition_imgList">
          {icons.map((icon, index) => (
            <img 
              key={index}
              className={`proposition_img ${activeIndex === index ? 'active' : ''}`}
              src={index === activeIndex ? icon.active : icon.inactive} 
              alt=""
              onMouseEnter={() => setActiveIndex(index)}
            />
          ))}
        </div>
        <div className="proposition_content container">
          <img src={contentList[activeIndex].img} className="proposition_content-img" alt="" />
          <div className="proposition_content-right">
            <div className="proposition_content-svg">
              <svg width="24" height="16" viewBox="0 0 24 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M0 5L10 0V16H0V5Z" fill="#C9CDD4"></path>
                <path d="M14 5L24 0V16H14V5Z" fill="#C9CDD4"></path>
              </svg>
            </div>
            <div className="proposition_content-t">{contentList[activeIndex].title}</div>
            <div className="proposition_content-d">{contentList[activeIndex].desc}</div>
          </div>
        </div>
      </div>
    </section>
  )
}

export default PropositionPart
