import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

function LoginPage({ onLogin }) {
  const [isLogin, setIsLogin] = useState(true)
  const [formData, setFormData] = useState({
    phone: '',
    password: '',
    code: '',
  })
  const [isSubmitting, setIsSubmitting] = useState(false)
  const navigate = useNavigate()

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setIsSubmitting(true)
    
    // 模拟登录/注册请求
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // 模拟登录成功
    const userInfo = {
      name: formData.phone.slice(-4) + '用户',
      phone: formData.phone,
    }
    
    onLogin(userInfo)
    setIsSubmitting(false)
    navigate('/')
  }

  const toggleMode = () => {
    setIsLogin(!isLogin)
    setFormData({ phone: '', password: '', code: '' })
  }

  return (
    <div className="login-page">
      <div className="login-modal">
        <h2 className="login-modal-title">
          {isLogin ? '欢迎回来' : '创建账号'}
        </h2>
        <p className="login-modal-subtitle">
          {isLogin ? '登录以继续使用引流宝' : '注册开始使用引流宝'}
        </p>
        
        <form className="login-form" onSubmit={handleSubmit}>
          <div className="form-item">
            <input
              type="tel"
              name="phone"
              className="form-input"
              placeholder="请输入手机号"
              value={formData.phone}
              onChange={handleChange}
              required
            />
          </div>
          
          <div className="form-item">
            <input
              type="password"
              name="password"
              className="form-input"
              placeholder="请输入密码"
              value={formData.password}
              onChange={handleChange}
              required
            />
          </div>
          
          {!isLogin && (
            <div className="form-item">
              <input
                type="text"
                name="code"
                className="form-input"
                placeholder="请输入验证码"
                value={formData.code}
                onChange={handleChange}
                required
              />
            </div>
          )}
          
          <button 
            type="submit" 
            className="login-btn"
            disabled={isSubmitting}
          >
            {isSubmitting ? '请稍候...' : (isLogin ? '登录' : '注册')}
          </button>
        </form>
        
        <div className="login-divider">
          <span>或</span>
        </div>
        
        <div className="social-login">
          <button className="social-login-btn" title="微信登录">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="#07C160">
              <path d="M8.691 2.188C3.891 2.188 0 5.476 0 9.53c0 2.212 1.17 4.203 3.002 5.55a.59.59 0 0 1 .213.665l-.39 1.48c-.019.07-.048.141-.048.213 0 .163.13.295.29.295a.326.326 0 0 0 .167-.054l1.903-1.114a.864.864 0 0 1 .717-.098 10.16 10.16 0 0 0 2.837.403c.276 0 .543-.027.811-.05-.857-2.578.157-4.972 1.932-6.446 1.703-1.415 3.882-1.98 5.853-1.838-.576-3.583-4.196-6.348-8.596-6.348zM5.785 5.991c.642 0 1.162.529 1.162 1.18a1.17 1.17 0 0 1-1.162 1.178A1.17 1.17 0 0 1 4.623 7.17c0-.651.52-1.18 1.162-1.18zm5.813 0c.642 0 1.162.529 1.162 1.18a1.17 1.17 0 0 1-1.162 1.178 1.17 1.17 0 0 1-1.162-1.178c0-.651.52-1.18 1.162-1.18zm5.34 2.867c-1.797-.052-3.746.512-5.28 1.786-1.72 1.428-2.687 3.72-1.78 6.22.942 2.453 3.666 4.229 6.884 4.229.826 0 1.622-.12 2.361-.336a.722.722 0 0 1 .598.082l1.584.926a.272.272 0 0 0 .14.047c.134 0 .24-.111.24-.247 0-.06-.023-.12-.038-.177l-.327-1.233a.582.582 0 0 1-.023-.156.49.49 0 0 1 .201-.398C23.024 18.48 24 16.82 24 14.98c0-3.21-2.931-5.837-6.656-6.088V8.89c-.135-.01-.269-.03-.407-.03zm-2.53 3.274c.535 0 .969.44.969.982a.976.976 0 0 1-.969.983.976.976 0 0 1-.969-.983c0-.542.434-.982.97-.982zm4.844 0c.535 0 .969.44.969.982a.976.976 0 0 1-.969.983.976.976 0 0 1-.969-.983c0-.542.434-.982.969-.982z"/>
            </svg>
          </button>
        </div>
        
        <div className="login-modal-footer">
          {isLogin ? (
            <>
              还没有账号？ <a onClick={toggleMode}>立即注册</a>
            </>
          ) : (
            <>
              已有账号？ <a onClick={toggleMode}>立即登录</a>
            </>
          )}
        </div>
      </div>
    </div>
  )
}

export default LoginPage
