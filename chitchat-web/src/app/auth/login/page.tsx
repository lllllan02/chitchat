import LoginForm from '@/components/auth/LoginForm';
import Layout from '@/components/layout/Layout';
import Link from 'next/link';

export const metadata = {
  title: '登录 - ChitChat',
  description: '登录您的ChitChat账号'
};

export default function LoginPage() {
  return (
    <Layout>
      <div className="flex items-center justify-center min-h-screen py-12 px-4 sm:px-6 lg:px-8 bg-gray-50">
        <div className="w-full max-w-md">
          <div className="text-center mb-6">
            <Link href="/" className="inline-block">
              <h2 className="text-3xl font-extrabold text-gray-900">ChitChat</h2>
            </Link>
          </div>
          
          <LoginForm />
        </div>
      </div>
    </Layout>
  );
} 