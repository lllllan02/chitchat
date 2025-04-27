import RegisterForm from '@/components/auth/RegisterForm';
import Layout from '@/components/layout/Layout';
import Link from 'next/link';

export const metadata = {
  title: '注册 - ChitChat',
  description: '创建一个新的ChitChat账号，加入我们的社区'
};

export default function RegisterPage() {
  return (
    <Layout>
      <div className="flex items-center justify-center min-h-screen py-12 px-4 sm:px-6 lg:px-8 bg-gray-50">
        <div className="w-full max-w-md">
          <div className="text-center mb-6">
            <Link href="/" className="inline-block">
              <h2 className="text-3xl font-extrabold text-gray-900">ChitChat</h2>
            </Link>
          </div>
          
          <RegisterForm />
          
          <div className="mt-6 text-center text-sm">
            <p className="text-gray-600">
              注册即表示您同意我们的
              <Link href="/terms" className="mx-1 font-medium text-blue-600 hover:text-blue-500">
                服务条款
              </Link>
              和
              <Link href="/privacy" className="ml-1 font-medium text-blue-600 hover:text-blue-500">
                隐私政策
              </Link>
            </p>
          </div>
        </div>
      </div>
    </Layout>
  );
} 