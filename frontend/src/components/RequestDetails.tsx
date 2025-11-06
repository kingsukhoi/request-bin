import type {Request} from '../api'
import type {Header} from "../api/getHeaders.ts";
import {GetHeaders} from "../api/getHeaders.ts";
import type {QueryParam} from "../api/getQueryParams.ts";
import {GetQueryParams} from "../api/getQueryParams.ts";
import {useEffect, useState} from "react";
import JsonView from '@uiw/react-json-view';
import {githubDarkTheme} from '@uiw/react-json-view/githubDark';
import SyntaxHighlighter from 'react-syntax-highlighter';
import {vs2015} from 'react-syntax-highlighter/dist/esm/styles/hljs';

interface RequestDetailsProps {
    request: Request
    onClose: () => void
}

export function RequestDetails({request, onClose}: RequestDetailsProps) {

    const [headers, setHeaders] = useState<Header[]>([])
    const [queryParams, setQueryParams] = useState<QueryParam[]>([])
    const [copied, setCopied] = useState(false)

    useEffect(() => {
        GetHeaders(request.id).then(setHeaders)
        GetQueryParams(request.id).then(setQueryParams)
    }, [request.id])

    const copyToClipboard = async () => {
        if (!request.content) return

        try {
            await navigator.clipboard.writeText(request.content)
            setCopied(true)
            setTimeout(() => setCopied(false), 2000)
        } catch (err) {
            console.error('Failed to copy:', err)
        }
    }


    return (
        <div
            className="bg-gh-bg-secondary md:rounded-lg p-6 shadow-lg md:sticky md:top-8 md:flex-1 md:max-h-[calc(100vh-4rem)] overflow-y-auto md:min-w-0 md:w-full fixed inset-0 z-50 md:relative md:z-auto">
            <div className="flex justify-between items-center mb-6 pb-4 border-b-2 border-gh-border gap-4">
                <h2 className="m-0 text-gh-text-primary text-2xl">Request Details</h2>
                <div className="flex items-center gap-4">
                    <span className="text-sm font-mono text-gh-text-muted">{request.id}</span>
                    <button
                        className="bg-gh-danger text-white border-none w-8 h-8 rounded-md cursor-pointer text-2xl leading-none transition-colors flex items-center justify-center hover:bg-gh-danger-hover flex-shrink-0"
                        onClick={onClose}>
                        ×
                    </button>
                </div>
            </div>
            <div className="grid gap-6">
                <div>
                    <h3 className="text-gh-text-primary mt-0 mb-3 text-lg">Basic Info</h3>
                    <p className="text-gh-text-secondary my-2 leading-relaxed"><strong>Method:</strong> {request.method}
                    </p>
                    <p className="text-gh-text-secondary my-2 leading-relaxed"><strong>Path:</strong> {request.path}</p>
                    <p className="text-gh-text-secondary my-2 leading-relaxed"><strong>Response
                        Code:</strong> {request.responseCode}</p>
                    <p className="text-gh-text-secondary my-2 leading-relaxed">
                        <strong>Timestamp:</strong> {request.timestamp.toLocaleString()}</p>
                    <p className="text-gh-text-secondary my-2 leading-relaxed"><strong>IP
                        Address:</strong> {request.sourceIp}</p>
                </div>

                <div>
                    <h3 className="text-gh-text-primary mt-0 mb-3 text-lg">Headers</h3>
                    <table className="w-full border-collapse bg-gh-bg-primary rounded-md overflow-hidden">
                        <thead>
                        <tr>
                            <th className="p-3 text-left font-semibold text-gh-text-primary bg-gh-bg-tertiary border-b-2 border-gh-border">Name</th>
                            <th className="p-3 text-left font-semibold text-gh-text-primary bg-gh-bg-tertiary border-b-2 border-gh-border">Value</th>
                        </tr>
                        </thead>
                        <tbody>
                        {headers.sort().map((header, index) => (
                            <tr key={index}>
                                <td className="p-3 text-gh-text-secondary border-b border-gh-border last:border-b-0">{header.name}</td>
                                <td className="p-3 text-gh-text-secondary border-b border-gh-border last:border-b-0">{header.value}</td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                </div>

                <div>
                    <h3 className="text-gh-text-primary mt-0 mb-3 text-lg">Query Parameters</h3>
                    <table className="w-full border-collapse bg-gh-bg-primary rounded-md overflow-hidden">
                        <thead>
                        <tr>
                            <th className="p-3 text-left font-semibold text-gh-text-primary bg-gh-bg-tertiary border-b-2 border-gh-border">Name</th>
                            <th className="p-3 text-left font-semibold text-gh-text-primary bg-gh-bg-tertiary border-b-2 border-gh-border">Value</th>
                        </tr>
                        </thead>
                        <tbody>
                        {queryParams.length > 0 ? (
                            queryParams.sort().map((param, index) => (
                                <tr key={index}>
                                    <td className="p-3 text-gh-text-secondary border-b border-gh-border last:border-b-0">{param.name}</td>
                                    <td className="p-3 text-gh-text-secondary border-b border-gh-border last:border-b-0">{param.value}</td>
                                </tr>
                            ))
                        ) : (
                            <tr>
                                <td colSpan={2} className="p-3 text-center italic text-gh-text-muted">
                                    No query parameters
                                </td>
                            </tr>
                        )}
                        </tbody>
                    </table>
                </div>

                {request.content && (
                    <div>
                        <div className="flex justify-between items-center mb-3">
                            <h3 className="text-gh-text-primary mt-0 mb-0 text-lg">Body</h3>
                            <button
                                className="bg-gh-success text-white border-none px-4 py-2 rounded-md cursor-pointer text-sm font-medium transition-colors flex items-center gap-1 hover:bg-gh-success-hover active:scale-95"
                                onClick={copyToClipboard}
                                title="Copy to clipboard"
                            >
                                {copied ? '✓ Copied' : '📋 Copy'}
                            </button>
                        </div>
                        {(() => {
                            try {
                                // Try to parse as JSON
                                const jsonData = JSON.parse(request.content);
                                return (
                                    <JsonView
                                        value={jsonData}
                                        style={githubDarkTheme}
                                        displayDataTypes={true}
                                        displayObjectSize={true}
                                        enableClipboard={false}
                                        shortenTextAfterLength={50}
                                    />
                                );
                            } catch {
                                // Not JSON, use SyntaxHighlighter for other formats
                                return (
                                    <div className="overflow-hidden max-w-full">
                                        <SyntaxHighlighter
                                            style={vs2015}
                                            wrapLongLines={true}
                                            showLineNumbers={true}
                                            showInlineLineNumbers={true}
                                            PreTag="div"
                                            customStyle={{
                                                background: '#0d1117',
                                                padding: '1rem',
                                                borderRadius: '6px',
                                                border: '1px solid #30363d',
                                                margin: 0,
                                                maxWidth: '100%',
                                                overflow: 'hidden'
                                            }}
                                            codeTagProps={{
                                                style: {
                                                    whiteSpace: 'pre-wrap',
                                                    wordBreak: 'normal',
                                                    overflowWrap: 'break-word'
                                                }
                                            }}
                                        >
                                            {request.content}
                                        </SyntaxHighlighter>
                                    </div>
                                );
                            }
                        })()}
                    </div>
                )}
            </div>
        </div>
    )
}
