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
      <div className="min-h-screen bg-gh-bg-primary">
      <Banner title="Request Bin" subtitle="Login to continue"/>

          <main className="max-w-md mx-auto md:p-8 p-4 flex justify-center items-center min-h-[calc(100vh-200px)]">
              <div className="bg-gh-bg-secondary rounded-lg md:p-8 p-6 shadow-lg w-full">
                  <h2 className="text-gh-text-primary mt-0 mb-6 text-center text-3xl">Login</h2>

                  <form onSubmit={handleSubmit} className="flex flex-col gap-5">
                      <div className="flex flex-col gap-2">
                          <label htmlFor="username"
                                 className="text-gh-text-primary text-sm font-medium">Username</label>
              <input
                id="username"
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
                disabled={isLoading}
                autoComplete="username"
                className="bg-gh-bg-primary border border-gh-border rounded-md text-gh-text-secondary px-3 py-3 text-base transition-colors focus:outline-none focus:border-blue-500 hover:border-gray-600 disabled:opacity-60 disabled:cursor-not-allowed"
              />
            </div>

                      <div className="flex flex-col gap-2">
                          <label htmlFor="password"
                                 className="text-gh-text-primary text-sm font-medium">Password</label>
              <input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                disabled={isLoading}
                autoComplete="current-password"
                className="bg-gh-bg-primary border border-gh-border rounded-md text-gh-text-secondary px-3 py-3 text-base transition-colors focus:outline-none focus:border-blue-500 hover:border-gray-600 disabled:opacity-60 disabled:cursor-not-allowed"
              />
            </div>

            {error && (
                <div className="bg-gh-danger text-white px-3 py-3 rounded-md text-sm text-center">
                {error}
              </div>
            )}

            <button
              type="submit"
              className="bg-gh-success text-white border-none px-3 py-3 rounded-md cursor-pointer text-base font-semibold transition-colors mt-2 hover:bg-gh-success-hover disabled:bg-gray-500 disabled:cursor-not-allowed disabled:opacity-70"
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
