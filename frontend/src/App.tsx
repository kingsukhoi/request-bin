import {useState, useEffect} from 'react'
import './App.css'
import {RequestsTable} from './components/RequestsTable'
import {RequestDetails} from './components/RequestDetails'
import {GetRequests, type Request} from "./api"

function App() {

    const [requests, setRequests] = useState<Request[]>([])

    const [selectedRequest, setSelectedRequest] = useState<Request | null>(null)

    useEffect(() => {
        GetRequests().then(setRequests)

        const interval = setInterval(() => {
            GetRequests().then(setRequests)
        }, 10000)

        return () => clearInterval(interval)
    }, [])


    return (
        <div className="app-container">
            <header className="app-header">
                <h1>Request Bin</h1>
                <p>Monitor and inspect HTTP requests in real-time</p>
            </header>

            <main className="main-content">
                <RequestsTable
                    requests={requests}
                    onSelectRequest={setSelectedRequest}
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
