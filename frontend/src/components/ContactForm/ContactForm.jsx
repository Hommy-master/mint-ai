import { useState } from 'react'
import './ContactForm.css'

function ContactForm() {
  const [formData, setFormData] = useState({
    name: '',
    phone: '',
    company: '',
  })
  const [isSubmitting, setIsSubmitting] = useState(false)

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setIsSubmitting(true)
    
    // 模拟提交
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    alert('提交成功！我们会尽快与您联系。')
    setFormData({ name: '', phone: '', company: '' })
    setIsSubmitting(false)
  }

  return (
    <>
      <section className="home_form_part">
        <div className="home_form">
          <h2 className="part_title">联系我们</h2>
          <p className="part_desc">留下您的信息，我们将为您提供专业的解决方案</p>
          
          <form className="info_form" onSubmit={handleSubmit}>
            <div className="info_form-top">
              <div className="form-item required">
                <input
                  type="text"
                  name="name"
                  className="form-input"
                  placeholder="请输入您的姓名"
                  value={formData.name}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="form-item required">
                <input
                  type="tel"
                  name="phone"
                  className="form-input"
                  placeholder="请输入您的手机号"
                  value={formData.phone}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="form-item">
                <input
                  type="text"
                  name="company"
                  className="form-input"
                  placeholder="请输入您的公司名称"
                  value={formData.company}
                  onChange={handleChange}
                />
              </div>
            </div>
            <button 
              type="submit" 
              className={`primary-btn ${isSubmitting ? 'loading' : ''}`}
              disabled={isSubmitting}
            >
              {isSubmitting && (
                <span className="loading_icon">
                  <svg viewBox="0 0 24 24" width="16" height="16">
                    <circle cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="2" fill="none" strokeDasharray="31.416" strokeDashoffset="10"/>
                  </svg>
                </span>
              )}
              立即咨询
            </button>
          </form>
        </div>
      </section>

      {/* 底部 CTA */}
      <section className="bottom-try-bar">
        <h2 className="bottom-try-bar-title">开启您的 AI 营销之旅</h2>
        <a 
          href="https://jinshuju.net/f/Tw9wZH?x_field_1=https%3A%2F%2Ficut.cn%2F"
          target="_blank"
          rel="noreferrer"
          className="primary-btn"
        >
          免费试用
        </a>
      </section>
    </>
  )
}

export default ContactForm
