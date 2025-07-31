package helpers

import (
	"bytes"
	"compress/gzip"
	"io"
)

func isGzipEncoded(data []byte) bool {
	// Check if data starts with gzip magic number (0x1f, 0x8b)
	return len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b
}

func decompressGzip(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return data, err
	}
	defer reader.Close()

	return io.ReadAll(reader)
}

func CheckAndDecompressGzip(data []byte) ([]byte, error) {
	if isGzipEncoded(data) {
		return decompressGzip(data)
	}
	// Return original data if not gzip encoded
	return data, nil
}
