import './MyData.css'

const dataItems = [
  { title: '400万+', desc: '累计完成视频切片' },
  { title: '1万+', desc: '累计覆盖直播间次数' },
  { title: '2000+', desc: '累计赋能品牌' },
  { title: '20+', desc: '单场直播可自动生成切片数' }
]

function MyData() {
  return (
    <section className="my_data">
      <div className="my_data-list">
        {dataItems.map((item, index) => (
          <div className="my_data-item" key={index}>
            <div className="my_data-svg">
              <svg width="24" height="16" viewBox="0 0 24 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M0 5L10 0V16H0V5Z" fill="white" fillOpacity="0.6"></path>
                <path d="M14 5L24 0V16H14V5Z" fill="white" fillOpacity="0.6"></path>
              </svg>
            </div>
            <div className="my_data-title">{item.title}</div>
            <div className="my_data-desc">{item.desc}</div>
          </div>
        ))}
      </div>
    </section>
  )
}

export default MyData
