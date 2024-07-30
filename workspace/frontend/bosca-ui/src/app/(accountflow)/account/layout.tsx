import type {Metadata} from "next"
import "../../globals.css"

export const metadata: Metadata = {
  title: "Bosca",
  description: "Bosca Account Management",
  manifest: '/site.webmanifest',
}

export default function FullPageLayout({children}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en" className="h-full bg-gray-50">
    <body className="h-full">
    {children}
    </body>
    </html>
  )
}
