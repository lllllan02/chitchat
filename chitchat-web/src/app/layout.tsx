import '../styles/globals.css';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'ChitChat',
  description: '一个自由分享知识、交流经验的社区平台',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="zh-CN">
      <body>
        {children}
      </body>
    </html>
  );
} 