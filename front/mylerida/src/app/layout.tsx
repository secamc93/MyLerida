"use client";

import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import styles from "./page.module.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable}`}>
        <div className={styles.page}>
          <header className={styles.header}>
            <h1>Bienvenido a MyLerida</h1>
          </header>
          <main className={styles.main}>{children}</main>
          <footer className={styles.footer}>
            <p>© 2023 MyLerida. Todos los derechos reservados.</p>
          </footer>
        </div>
      </body>
    </html>
  );
}
