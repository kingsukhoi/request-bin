interface BannerProps {
    title: string
    subtitle: string
}

export function Banner({ title, subtitle }: BannerProps) {
    return (
        <header className="bg-gradient-to-br from-blue-900 to-blue-500 md:p-8 p-6 text-center text-white shadow-lg">
            <h1 className="m-0 md:text-4xl text-3xl font-bold">{title}</h1>
            <p className="mt-2 mb-0 md:text-lg text-base opacity-90">{subtitle}</p>
        </header>
    )
}
