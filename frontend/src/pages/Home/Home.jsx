import HeroCarousel from '../../components/HeroCarousel/HeroCarousel'
import TrustSignals from '../../components/TrustSignals/TrustSignals'
import MyData from '../../components/MyData/MyData'
import PropositionPart from '../../components/PropositionPart/PropositionPart'
import Products from '../../components/Products/Products'
import SolutionsPart from '../../components/SolutionsPart/SolutionsPart'
import Cases from '../../components/Cases/Cases'
import BottomTryBar from '../../components/BottomTryBar/BottomTryBar'

function Home() {
  return (
    <main className="main-content">
      <HeroCarousel />
      <TrustSignals />
      <MyData />
      <PropositionPart />
      <Products />
      <SolutionsPart />
      <Cases />
      <BottomTryBar />
    </main>
  )
}

export default Home
