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
          to: '/viewRequests',
          search: {
            request_id: undefined,
            nextToken: undefined
          }
        })
      } else {
        navigate({to: '/login', search: {redirect: undefined}})
      }
    }

    checkAndRedirect()
  }, [navigate])

  // Show a loading state while checking auth
  return (
      <div className="min-h-screen bg-gh-bg-primary">
          <div className="flex justify-center items-center h-screen">
              <p className="text-gh-text-primary">Loading...</p>
      </div>
    </div>
  )
}
