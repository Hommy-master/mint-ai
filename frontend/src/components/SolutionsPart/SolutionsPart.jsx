import { useState, useEffect, useRef, useCallback } from 'react'
import Video1 from '../../assets/videos/1.mp4'
import Video2 from '../../assets/videos/2.mp4'
import Video3 from '../../assets/videos/3.mp4'
import Video4 from '../../assets/videos/4.mp4'
import Video5 from '../../assets/videos/5.mp4'

import Img1 from '../../assets/images/1.png'
import Img2 from '../../assets/images/2.png'
import Img3 from '../../assets/images/3.png'
import Img4 from '../../assets/images/4.png'
import Img5 from '../../assets/images/5.png'
import phoneImg from '../../assets/common/phone.png'

import './SolutionsPart.css'

const solutions = [
  {
    id: 0,
    title: '账号吸粉拉新',
    desc: 'icut自动生成优质短视频！主要用于账号涨粉，在直播中引流，在直播间进行成交和转粉',
    details: ['产品：爆款、FAB款、福利款', '时长：5s-10s，完播率70%', '内容：卖点展示/搭配展示/上身效果', '剪辑：单剪，整段抓取，适当增加音乐和字幕'],
    video: Video1,
    poster: Img1
  },
  {
    id: 1,
    title: '短视频带货',
    desc: 'icut自动生成优质短视频！通过短视频挂车，在直播之外的时间产生持续成交',
    details: ['产品：爆款、主推款', '时长：15s-30s', '内容：商品卖点+使用场景', '剪辑：精剪，突出重点'],
    video: Video2,
    poster: Img2
  },
  {
    id: 2,
    title: '短视频种草',
    desc: 'icut自动生成优质短视频！多平台推荐和展示商品，让消费者对商品产生兴趣和购买欲望',
    details: ['产品：新品、潜力款', '时长：10s-20s', '内容：商品展示+卖点介绍', '剪辑：轻剪，保持原貌'],
    video: Video3,
    poster: Img3
  },
  {
    id: 3,
    title: '主图视频/微详情',
    desc: 'icut自动生成优质短视频！淘宝—逛逛、微详情视频；京东——逛、新品、种草视频',
    details: ['产品：全品类', '时长：5s-15s', '内容：商品360°展示', '剪辑：自动适配各平台'],
    video: Video4,
    poster: Img4
  },
  {
    id: 4,
    title: 'IP授权带货',
    desc: 'icut自动生成优质短视频！借助已有大V流量，获得授权后直接带货分佣',
    details: ['产品：授权商品', '时长：15s-60s', '内容：IP形象+商品展示', '剪辑：混剪，增加趣味性'],
    video: Video5,
    poster: Img5
  }
]

function SolutionsPart() {
  const [activeIndex, setActiveIndex] = useState(0)
  const [isHovered, setIsHovered] = useState(false)
  const intervalRef = useRef(null)

  // 自动轮播逻辑
  const startAutoPlay = useCallback(() => {
    intervalRef.current = setInterval(() => {
      setActiveIndex((prev) => (prev + 1) % solutions.length)
    }, 5000) // 5秒切换一次
  }, [])

  const stopAutoPlay = useCallback(() => {
    if (intervalRef.current) {
      clearInterval(intervalRef.current)
      intervalRef.current = null
    }
  }, [])

  useEffect(() => {
    // 组件挂载时启动自动轮播
    startAutoPlay()
    
    return () => {
      // 组件卸载时清除定时器
      stopAutoPlay()
    }
  }, [startAutoPlay, stopAutoPlay])

  // 处理鼠标悬浮
  const handleMouseEnter = (index) => {
    setIsHovered(true)
    stopAutoPlay()
    setActiveIndex(index)
  }

  const handleMouseLeave = () => {
    setIsHovered(false)
    startAutoPlay()
  }

  return (
    <section className="solutions_part icut">
      <div className="container-wrap">
        <div className="part_title">效果展示</div>
        <div className="part_desc">满足多类型内容制作需求，商品卖点一网打尽</div>
      </div>
      
      {/* H5 版本 */}
      <div className="sol_h5_list">
        {solutions.map((sol) => (
          <div className="sol_h5_item" key={sol.id}>
            <div className="sol_h5_item-top">
              <video 
                className={`sol_h5_item-video v_${sol.id}`}
                src={sol.video}
                preload="metadata"
                poster={sol.poster}
                playsInline
                muted
                autoPlay
                loop
              />
            </div>
            <div className="sol_h5_item-bottom">
              <div className="sol_h5_item-t">{sol.title}</div>
              <div className="sol_h5_item-d">{sol.desc}</div>
            </div>
          </div>
        ))}
      </div>

      {/* PC 版本 */}
      <div className="sol_pc_pannel" onMouseLeave={handleMouseLeave}>
        <div className="sol_pc_pannel-menu">
          {solutions.map((sol, index) => (
            <div 
              key={sol.id}
              className={`sol_pc_pannel-menuItem ${activeIndex === index ? 'active' : ''}`}
              onMouseEnter={() => handleMouseEnter(index)}
            >
              {sol.title}
            </div>
          ))}
        </div>
        <div className="sol_pc_pannel-content">
          <div className="sol_pc_pannel-text">
            <div className="sol_pc_pannel-title">{solutions[activeIndex].title}</div>
            {solutions[activeIndex].details.map((detail, idx) => (
              <div className="sol_pc_pannel-desc" key={idx}>{detail}</div>
            ))}
          </div>
        </div>
        <div className="sol_pc_pannel-video-box">
          <img className="sol_pc_pannel-phone-img" src={phoneImg} alt="" />
          <div className="sol_pc_pannel-video-wrap">
            {solutions.map((sol, index) => (
              <video 
                key={sol.id}
                className={`sol_pc_pannel-video v-${index} ${activeIndex === index ? 'v-selected' : ''}`}
                src={sol.video}
                preload="metadata"
                poster={sol.poster}
                playsInline
                muted
                autoPlay
                loop
              />
            ))}
          </div>
        </div>
      </div>
    </section>
  )
}

export default SolutionsPart
