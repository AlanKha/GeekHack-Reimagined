import "~/styles/globals.css";
import TopNav from "~/app/_components/TopNav";
import SideBar from "~/app/_components/SideBar";

import { ClerkProvider } from "@clerk/nextjs";
import { GeistSans } from "geist/font/sans";
import { type Metadata } from "next";

export const metadata: Metadata = {
  title: "GeekHack, Reimagined",
  description: "A rehaul of the original GeekHack website",
  icons: [
    {
      rel: "icon",
      url: "/favicon.svg",
      type: "image/svg+xml",
    },
  ],
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <ClerkProvider>
      <html lang="en" className={`${GeistSans.variable} bg-black text-white`}>
        <body>
          <TopNav />
          <div className="flex">
            <SideBar />
            <div className="flex min-h-screen flex-grow flex-col items-center p-40 bg-black text-secondary">
              {children}
            </div>
          </div>
        </body>
      </html>
    </ClerkProvider>
  );
}
