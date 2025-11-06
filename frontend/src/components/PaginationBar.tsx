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
        <div className="bg-gh-bg-secondary rounded-lg p-4 px-6 shadow-lg flex flex-col gap-4 flex-shrink-0">
            <div className="flex justify-center items-center gap-4 flex-wrap">
                <button
                    className="bg-gh-success text-white border-none px-4 py-2 rounded-md cursor-pointer text-sm font-medium transition-colors min-w-[100px] hover:bg-gh-success-hover disabled:bg-gray-500 disabled:cursor-not-allowed disabled:opacity-50"
                    onClick={onFirstPage}
                    disabled={!hasPreviousPage}
                >
                    ⇤ First
                </button>
                <button
                    className="bg-gh-success text-white border-none px-4 py-2 rounded-md cursor-pointer text-sm font-medium transition-colors min-w-[100px] hover:bg-gh-success-hover disabled:bg-gray-500 disabled:cursor-not-allowed disabled:opacity-50"
                    onClick={onPreviousPage}
                    disabled={!hasPreviousPage}
                >
                    ← Previous
                </button>
                <button
                    className="bg-gh-success text-white border-none px-4 py-2 rounded-md cursor-pointer text-sm font-medium transition-colors min-w-[100px] hover:bg-gh-success-hover disabled:bg-gray-500 disabled:cursor-not-allowed disabled:opacity-50"
                    onClick={onNextPage}
                    disabled={!hasNextPage}
                >
                    Next →
                </button>
            </div>

            <div className="flex items-center gap-2 justify-center">
                <label htmlFor="page-size" className="text-gh-text-primary text-sm font-medium">Page size:</label>
                <input
                    id="page-size"
                    type="number"
                    min="1"
                    max="1000"
                    value={pageSize}
                    onChange={handlePageSizeChange}
                    className="bg-gh-bg-primary border border-gh-border rounded-md text-gh-text-secondary px-2 py-2 w-20 text-sm transition-colors focus:outline-none focus:border-blue-500 hover:border-gray-600"
                />
            </div>
        </div>
    );
}
