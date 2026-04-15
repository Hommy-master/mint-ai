import { useState } from 'react'
import './Products.css'
import featureIcon1 from '../../assets/common/feature_icon_1.svg'
import featureIcon1Active from '../../assets/common/feature_icon_1_active.svg'
import featureIcon2 from '../../assets/common/feature_icon_2.svg'
import featureIcon2Active from '../../assets/common/feature_icon_2_active.svg'
import featureIcon3 from '../../assets/common/feature_icon_3.svg'
import featureIcon3Active from '../../assets/common/feature_icon_3_active.svg'

const productsData = [
  {
    id: 1,
    title: '切片抓取',
    desc: '智能识别连续数小时直播中，优质的商品讲解与演示片段，自动完成多个版本的内容切片。',
    icon: featureIcon1,
    iconActive: featureIcon1Active,
  },
  {
    id: 2,
    title: '在线剪辑',
    desc: '基于直播切片可在线完成混剪、增加与去除贴片、加字幕、加背景音乐等功能，实现内容高效美化。',
    icon: featureIcon2,
    iconActive: featureIcon2Active,
  },
  {
    id: 3,
    title: '一键发布',
    desc: '对接多个短视频内容平台，基于已完成的切片内容，可以一键发布完成高效率多平台内容覆盖。',
    icon: featureIcon3,
    iconActive: featureIcon3Active,
  },
]

function Products() {
  const [activeProduct, setActiveProduct] = useState(0)

  return (
    <section className="products_part">
      <div className="container">
        <div className="part_title">"傻瓜式"直播短视频切片工具</div>
        <div className="part_desc">引流宝自动化生成直播商品短视频，助力每一个直播间的流量增长</div>
        
        <div className="product_list">
          {productsData.map((product, index) => (
            <div
              key={product.id}
              className="product_item"
              onMouseEnter={() => setActiveProduct(index)}
            >
              <div className="product_wrap">
                <div className={`product `}>
                  <div className="product-top">
                    <img
                      className="product-icon"
                      src={product.icon}
                      alt={product.title}
                    />
                    <img
                      className="product-iconActive"
                      src={product.iconActive}
                      alt={product.title}
                    />
                    <div className="product-title">{product.title}</div>
                  </div>
                  <div className="product-desc">{product.desc}</div>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}

export default Products
