import type {Request} from "../api"

interface RequestsTableProps {
    requests: Request[]
    onSelectRequest: (request: Request) => void
    onRefresh: () => void
    lastRefreshed: Date | null
    isRefreshing: boolean
    selectedRequest: Request | null
}

export function RequestsTable({requests, onSelectRequest, onRefresh, lastRefreshed, isRefreshing, selectedRequest}: RequestsTableProps) {
    const getMethodColor = (method: string) => {
        const colors: Record<string, string> = {
            GET: '#61affe',
            POST: '#49cc90',
            PUT: '#fca130',
            DELETE: '#f93e3e',
            PATCH: '#50e3c2'
        }
        return colors[method] || '#999'
    }

    requests.sort((a, b) => new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime())


    return (
        <div className="bg-gh-bg-secondary rounded-lg p-6 shadow-lg flex-1 min-w-0 overflow-x-auto">
            <div className="flex justify-between items-center mb-6">
                <h2 className="m-0 text-gh-text-primary text-2xl">Incoming Requests</h2>
                <div className="flex items-center gap-4">
                    {lastRefreshed && (
                        <span className="text-gh-text-muted text-sm">
                            Last updated: {lastRefreshed.toLocaleTimeString()}
                        </span>
                    )}
                    <button
                        className={`bg-gh-success text-white border-none w-10 h-10 rounded-md cursor-pointer text-xl font-medium transition-colors flex items-center justify-center hover:bg-gh-success-hover disabled:bg-gray-500 disabled:cursor-not-allowed ${isRefreshing ? 'animate-spin' : ''}`}
                        onClick={onRefresh}
                        disabled={isRefreshing}
                    >
                        ↻
                    </button>
                </div>
            </div>
            <table className="w-full border-collapse bg-gh-bg-primary rounded-lg overflow-x-auto">
                <thead className="bg-gh-bg-tertiary">
                <tr>
                    <th className="p-4 text-left font-semibold text-gh-text-primary border-b-2 border-gh-border">Method</th>
                    <th className="p-4 text-left font-semibold text-gh-text-primary border-b-2 border-gh-border">Path</th>
                    <th className="p-4 text-left font-semibold text-gh-text-primary border-b-2 border-gh-border">Timestamp</th>
                    <th className="p-4 text-left font-semibold text-gh-text-primary border-b-2 border-gh-border">IP
                        Address
                    </th>
                </tr>
                </thead>
                <tbody>
                {requests.length === 0 ? (
                    <tr>
                        <td colSpan={4} className="text-center p-12 text-gh-text-muted italic">
                            No requests yet. Start sending requests to your bin!
                        </td>
                    </tr>
                ) : (
                    requests.map((request) => (
                        <tr
                            key={request.id}
                            onClick={() => onSelectRequest(request)}
                            className={`border-b border-gh-border transition-colors cursor-pointer hover:bg-gh-bg-secondary ${selectedRequest?.id === request.id ? 'bg-gh-bg-tertiary' : ''}`}
                            style={selectedRequest?.id === request.id ? {boxShadow: 'inset 3px 0 0 #3b82f6'} : undefined}
                        >
                            <td className="p-4 text-gh-text-secondary">
                  <span
                      className="inline-block px-3 py-1 rounded font-semibold text-sm text-white uppercase"
                      style={{backgroundColor: getMethodColor(request.method)}}
                  >
                    {request.method}
                  </span>
                            </td>
                            <td className="p-4 text-gh-text-secondary font-mono text-gh-accent">{request.path}</td>
                            <td className="p-4 text-gh-text-secondary">{request.timestamp.toLocaleString()}</td>
                            <td className="p-4 text-gh-text-secondary">{request.sourceIp}</td>
                        </tr>
                    ))
                )}
                </tbody>
            </table>
        </div>
    )
}
