import {type FormEvent, useState} from 'react'
import {useNavigate, useSearch} from '@tanstack/react-router'
import {login} from '../api'
import {Banner} from '../components/Banner'

export function Login() {
  const navigate = useNavigate()
  const search = useSearch({from: '/login'})
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setError(null)
    setIsLoading(true)

    try {
      await login({username, password})
      // On success, navigate to the redirect URL or default to viewRequests
      if (search.redirect) {
        // Parse the redirect URL to extract pathname and search params
        const url = new URL(search.redirect, window.location.origin)
        const searchParams: Record<string, string> = {}
        url.searchParams.forEach((value, key) => {
          searchParams[key] = value
        })

        navigate({
          to: url.pathname,
          search: searchParams
        })
      } else {
        navigate({
          to: '/viewRequests',
          search: {
            request_id: "",
            nextToken: "",
          }
        })
      }
    } catch (err) {
      // Handle login error
      if (err instanceof Error) {
        setError(err.message)
      } else {
        setError('Login failed. Please check your credentials.')
      }
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="app-container">
      <Banner title="Request Bin" subtitle="Login to continue"/>

      <main className="login-content">
        <div className="login-card">
          <h2>Login</h2>

          <form onSubmit={handleSubmit} className="login-form">
            <div className="form-group">
              <label htmlFor="username">Username</label>
              <input
                id="username"
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
                disabled={isLoading}
                autoComplete="username"
              />
            </div>

            <div className="form-group">
              <label htmlFor="password">Password</label>
              <input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                disabled={isLoading}
                autoComplete="current-password"
              />
            </div>

            {error && (
              <div className="error-message">
                {error}
              </div>
            )}

            <button
              type="submit"
              className="login-button"
              disabled={isLoading}
            >
              {isLoading ? 'Logging in...' : 'Login'}
            </button>
          </form>
        </div>
      </main>
    </div>
  )
}
