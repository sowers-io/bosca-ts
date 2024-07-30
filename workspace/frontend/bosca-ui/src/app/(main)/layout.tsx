import type {Metadata} from "next"
import "@/app/globals.css"

export const metadata: Metadata = {
  title: "Bosca",
  description: "Bosca: The AI Content Management System",
  manifest: '/site.webmanifest',
}

export default function RootLayout({children}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
    <body>
    {children}
    </body>
    </html>
  )
}
