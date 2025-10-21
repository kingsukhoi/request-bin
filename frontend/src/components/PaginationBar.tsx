interface PaginationBarProps {
    pageSize: number;
    onPageSizeChange: (size: number) => void;
    onFirstPage: () => void;
    onPreviousPage: () => void;
    onNextPage: () => void;
    hasPreviousPage: boolean;
    hasNextPage: boolean;
}

export function PaginationBar({
                                  pageSize,
                                  onPageSizeChange,
                                  onFirstPage,
                                  onPreviousPage,
                                  onNextPage,
                                  hasPreviousPage,
                                  hasNextPage
                              }: PaginationBarProps) {
    const handlePageSizeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = parseInt(e.target.value);
        if (value > 0 && value <= 1000) {
            onPageSizeChange(value);
        }
    };

    return (
        <div className="pagination-bar">
            <button
                className="pagination-button"
                onClick={onFirstPage}
                disabled={!hasPreviousPage}
            >
                ⇤ First
            </button>
            <button
                className="pagination-button"
                onClick={onPreviousPage}
                disabled={!hasPreviousPage}
            >
                ← Previous
            </button>

            <div className="page-size-control">
                <label htmlFor="page-size">Page size:</label>
                <input
                    id="page-size"
                    type="number"
                    min="1"
                    max="1000"
                    value={pageSize}
                    onChange={handlePageSizeChange}
                    className="page-size-input"
                />
            </div>

            <button
                className="pagination-button"
                onClick={onNextPage}
                disabled={!hasNextPage}
            >
                Next →
            </button>
        </div>
    );
}
