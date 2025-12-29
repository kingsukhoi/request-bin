import {useEffect} from 'react'
import {useNavigate} from '@tanstack/react-router'

export function Home() {
  const navigate = useNavigate()

  useEffect(() => {
    // Redirect to /viewRequests - the router's beforeLoad hook will handle auth checking
    void navigate({
      to: '/viewRequests',
      search: {
        request_id: undefined,
        nextToken: undefined
      }
    })
  }, [navigate])

  return (
    <div className="min-h-screen bg-gh-bg-primary">
      <div className="flex justify-center items-center min-h-screen">
        <p className="text-gh-text-primary">Loading...</p>
      </div>
    </div>
  )
}
