import { useState } from 'react'
import { Link, useLocation } from 'react-router-dom'
import { DefaultLogoUrl } from '../../utils/const'

import './Header.css'

function Header({ isLoggedIn, userInfo, onLoginClick }) {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false)
  const location = useLocation()

  const toggleMobileMenu = () => {
    setMobileMenuOpen(!mobileMenuOpen)
  }

  const closeMobileMenu = () => {
    setMobileMenuOpen(false)
  }

  return (
    <>
      <header className="header">
        <div className="navbar_pc">
          <Link to="/" className="navbar_pc-logo" aria-label="Home">
            <img src={DefaultLogoUrl} alt="逗赛科技" /> 
            <span>逗赛科技</span>
          </Link>
        </div>
      </header>
    </>
  )
}

export default Header
