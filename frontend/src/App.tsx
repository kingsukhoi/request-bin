import {useState, useEffect, useCallback, useRef} from 'react'
import './App.css'
import {RequestsTable} from './components/RequestsTable'
import {RequestDetails} from './components/RequestDetails'
import {GetRequests, type Request} from "./api"

function App() {

    const [requests, setRequests] = useState<Request[]>([])

    const [selectedRequest, setSelectedRequest] = useState<Request | null>(null)
    const [lastRefreshed, setLastRefreshed] = useState<Date | null>(null)
    const [isRefreshing, setIsRefreshing] = useState(false)
    const debounceRef = useRef<number | null>(null)

    const refreshRequests = useCallback(async () => {
        if (isRefreshing) return

        setIsRefreshing(true)
        try {
            const requests = await GetRequests()
            setRequests(requests)
            setLastRefreshed(new Date())
        } finally {
          setIsRefreshing(false)
        }
    }, [isRefreshing])

    const debouncedRefresh = useCallback(() => {
        if (debounceRef.current) {
            clearTimeout(debounceRef.current)
        }
        debounceRef.current = setTimeout(() => {
            refreshRequests()
        }, 300)
    }, [refreshRequests])

    useEffect(() => {
        refreshRequests()
    }, [])//leave that empty or it spams the server


    return (
        <div className="app-container">
            <header className="app-header">
                <h1>Request Bin</h1>
                <p>Monitor and inspect HTTP requests</p>
            </header>

            <main className="main-content">
                <RequestsTable
                    requests={requests}
                    onSelectRequest={setSelectedRequest}
                    onRefresh={debouncedRefresh}
                    lastRefreshed={lastRefreshed}
                    isRefreshing={isRefreshing}
                    selectedRequest={selectedRequest}
                />

                {selectedRequest && (
                    <RequestDetails
                        request={selectedRequest}
                        onClose={() => setSelectedRequest(null)}
                    />
                )}
            </main>
        </div>
    )
}

export default App
