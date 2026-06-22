import type { Metadata } from "next";
import "./globals.css";
import "highlight.js/styles/github-dark.css";

export const metadata: Metadata = {
	title: "go-emoji-shortcode",
	description:
		"Dependency-free Go library for mapping emoji shortcodes like :smile: to Unicode emoji characters, with lookup and suggestion APIs.",
};

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<html lang="en">
			<body>
				<header className="site-header">
					<a href="/" className="brand">
						go-emoji-shortcode
					</a>
					<nav>
						<a href="https://github.com/floatpane/go-emoji-shortcode">GitHub</a>
					</nav>
				</header>
				<main>{children}</main>
			</body>
		</html>
	);
}

export const viewport = { width: "device-width", initialScale: 1 };
