import { useState } from 'react'
import { Routes, Route, useLocation } from 'react-router-dom'
import Header from './components/Header/Header'
import Footer from './components/Footer/Footer'
import LoginModal from './components/LoginModal/LoginModal'
import RightBar from './components/RightBar/RightBar'
import Home from './pages/Home/Home'
// import LoginPage from './pages/LoginPage/LoginPage'

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [userInfo, setUserInfo] = useState(null)
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false)
  const location = useLocation()

  // 判断是否在登录页面
  const isLoginPage = false; //location.pathname === '/login'

  const handleLogin = (user) => {
    setIsLoggedIn(true)
    setUserInfo(user)
    // 可以在这里存储到 localStorage
    localStorage.setItem('userInfo', JSON.stringify(user))
  }

  const handleLogout = () => {
    setIsLoggedIn(false)
    setUserInfo(null)
    localStorage.removeItem('userInfo')
  }

  const openLoginModal = () => {
    setIsLoginModalOpen(true)
  }

  const closeLoginModal = () => {
    setIsLoginModalOpen(false)
  }

  return (
    <div className="main">
      {/* 在非登录页面显示 Header */}
      {!isLoginPage && (
        <Header 
          isLoggedIn={isLoggedIn} 
          userInfo={userInfo}
          onLoginClick={openLoginModal}
        />
      )}
      
      <Routes>
        <Route path="/" element={<Home />} />
        {/* <Route path="/product" element={<Product />} />
        <Route path="/solution" element={<Solution />} />
        <Route path="/price" element={<Price />} />
        <Route path="/login" element={<LoginPage onLogin={handleLogin} />} /> */}
      </Routes>
      
      {/* 在非登录页面显示 Footer */}
      {!isLoginPage && <Footer />}
      
      {/* 登录弹窗 */}
      <LoginModal 
        isOpen={isLoginModalOpen} 
        onClose={closeLoginModal}
        onLogin={handleLogin}
      />
      
      {/* 右侧悬浮栏 */}
      {!isLoginPage && <RightBar />}
    </div>
  )
}

export default App
