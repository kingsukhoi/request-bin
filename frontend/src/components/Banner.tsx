interface BannerProps {
    title: string
    subtitle: string
}

export function Banner({ title, subtitle }: BannerProps) {
    return (
        <header className="app-header">
            <h1>{title}</h1>
            <p>{subtitle}</p>
        </header>
    )
}
