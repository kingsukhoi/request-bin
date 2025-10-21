import {useCallback, useEffect, useState} from 'react'
import './App.css'
import {RequestsTable} from './components/RequestsTable'
import {RequestDetails} from './components/RequestDetails'
import {PaginationBar} from './components/PaginationBar'
import {GetRequests, type Request} from "./api"

function App() {

    const [requests, setRequests] = useState<Request[]>([])

    const [selectedRequest, setSelectedRequest] = useState<Request | null>(null)
    const [lastRefreshed, setLastRefreshed] = useState<Date | null>(null)
    const [isRefreshing, setIsRefreshing] = useState(false)

    const [pageSize, setPageSize] = useState(20)
    const [currentToken, setCurrentToken] = useState<string | undefined>(undefined)
    const [previousTokens, setPreviousTokens] = useState<string[]>([])

    const refreshRequests = useCallback(async (token?: string) => {
        if (isRefreshing) return

        setIsRefreshing(true)
        try {
            const requestsData = await GetRequests({
                limit: pageSize,
                nextToken: token
            })
            setRequests(requestsData)
            setCurrentToken(token)
            setLastRefreshed(new Date())
        } finally {
          setIsRefreshing(false)
        }
    }, [isRefreshing, pageSize])

    useEffect(() => {
        refreshRequests()
    }, [])//leave that empty or it spams the server

    useEffect(() => {
        // Reset to first page when page size changes
        setPreviousTokens([])
        refreshRequests()
    }, [pageSize])

    const handleNextPage = () => {
        // Use the ID of the last request as the next token
        if (requests.length > 0) {
            const lastRequestId = requests[requests.length - 1].id
            // Store current token so we can go back
            setPreviousTokens([...previousTokens, currentToken])
            refreshRequests(lastRequestId)
        }
    }

    const handleFirstPage = () => {
        setPreviousTokens([])
        refreshRequests(undefined)
    }

    const handlePreviousPage = () => {
        if (previousTokens.length > 0) {
            const newPreviousTokens = [...previousTokens]
            const previousToken = newPreviousTokens.pop()
            setPreviousTokens(newPreviousTokens)
            refreshRequests(previousToken)
        } else {
            // Go back to first page
            refreshRequests(undefined)
        }
    }

    const handlePageSizeChange = (size: number) => {
        setPageSize(size)
    }


    return (
        <div className="app-container">
            <header className="app-header">
                <h1>Request Bin</h1>
                <p>Monitor and inspect HTTP requests</p>
            </header>

            <main className="main-content">
                <div className="requests-section">
                    <RequestsTable
                        requests={requests}
                        onSelectRequest={setSelectedRequest}
                        onRefresh={() => refreshRequests()}
                        lastRefreshed={lastRefreshed}
                        isRefreshing={isRefreshing}
                        selectedRequest={selectedRequest}
                    />

                    <PaginationBar
                        pageSize={pageSize}
                        onPageSizeChange={handlePageSizeChange}
                        onFirstPage={handleFirstPage}
                        onPreviousPage={handlePreviousPage}
                        onNextPage={handleNextPage}
                        hasPreviousPage={previousTokens.length > 0}
                        hasNextPage={requests.length >= pageSize}
                    />
                </div>

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
