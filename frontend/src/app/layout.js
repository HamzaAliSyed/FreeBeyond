import './globals.css'

export const metadata = {
  title: 'My App',
  description: 'My Next.js App',
}

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}