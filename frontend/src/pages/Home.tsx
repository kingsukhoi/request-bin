import {useEffect} from 'react'
import {useNavigate} from '@tanstack/react-router'
import {checkAuth} from '../api'

export function Home() {
  const navigate = useNavigate()

  useEffect(() => {
    const checkAndRedirect = async () => {
      const isAuthenticated = await checkAuth()

      if (isAuthenticated) {
        navigate({
          to: '/viewRequests'
        })
      } else {
        navigate({to: '/login'})
      }
    }

    checkAndRedirect()
  }, [navigate])

  // Show a loading state while checking auth
  return (
    <div className="app-container">
      <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh'
      }}>
        <p>Loading...</p>
      </div>
    </div>
  )
}
