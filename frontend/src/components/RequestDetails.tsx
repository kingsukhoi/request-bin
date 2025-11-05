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
    <div className="request-details">
      <div className="details-header">
        <h2>Request Details</h2>
        <button className="close-button" onClick={onClose}>
          ×
        </button>
      </div>
      <div className="details-content">
        <div className="detail-section">
          <h3>Basic Info</h3>
          <p><strong>Method:</strong> {request.method}</p>
          <p><strong>Path:</strong> {request.path}</p>
          <p><strong>Timestamp:</strong> {request.timestamp.toLocaleString()}</p>
          <p><strong>IP Address:</strong> {request.sourceIp}</p>
        </div>

        <div className="detail-section">
          <h3>Headers</h3>
          <table className="headers-table">
            <thead>
            <tr>
              <th>Name</th>
              <th>Value</th>
            </tr>
            </thead>
            <tbody>
            {headers.sort().map((header, index) => (
              <tr key={index}>
                <td>{header.name}</td>
                <td>{header.value}</td>
              </tr>
            ))}
            </tbody>
          </table>
        </div>

        <div className="detail-section">
          <h3>Query Parameters</h3>
          <table className="headers-table">
            <thead>
            <tr>
              <th>Name</th>
              <th>Value</th>
            </tr>
            </thead>
            <tbody>
            {queryParams.length > 0 ? (
              queryParams.sort().map((param, index) => (
                <tr key={index}>
                  <td>{param.name}</td>
                  <td>{param.value}</td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan={2} style={{textAlign: 'center', fontStyle: 'italic', color: '#8b949e'}}>
                  No query parameters
                </td>
              </tr>
            )}
            </tbody>
          </table>
        </div>

        {request.content && (
          <div className="detail-section">
            <div className="section-header-with-button">
              <h3>Body</h3>
              <button
                className="copy-button"
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
                  />
                );
              } catch {
                // Not JSON, use SyntaxHighlighter for other formats
                return (
                  <div style={{overflow: 'hidden', maxWidth: '100%'}}>
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
