import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/Sidebar";
import Header from "@/components/Header";
import { AuthProvider } from "@/context/AuthContext";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Geekhack - Keyboard Enthusiasts",
  description: "A complete overhaul of a legacy forum, architected from the ground up as a modern, high-performance web application.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <link href="https://fonts.googleapis.com/icon?family=Material+Icons+Outlined" rel="stylesheet"/>
      </head>
      <body className={`${inter.className} antialiased`}>
        <AuthProvider>
          <div className="flex">
            <Sidebar />
            <div className="flex-1">
              <Header />
              <main className="p-4 md:p-6 lg:p-8">
                {children}
              </main>
            </div>
          </div>
        </AuthProvider>
      </body>
    </html>
  );
}
