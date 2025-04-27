"use client";

import { useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import type { UserLoginData } from '@/types/user';
import { useAuth } from '@/context/AuthContext';

export default function LoginForm() {
  const router = useRouter();
  const { login } = useAuth();
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const [errors, setErrors] = useState({
    email: '',
    password: '',
    server: ''
  });

  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
    
    // 清除对应字段的错误信息
    if (errors[name as keyof typeof errors]) {
      setErrors(prev => ({
        ...prev,
        [name]: ''
      }));
    }
    
    // 如果有服务器错误，也一并清除
    if (errors.server) {
      setErrors(prev => ({
        ...prev,
        server: ''
      }));
    }
  };

  const validateForm = () => {
    let valid = true;
    const newErrors = {
      email: '',
      password: '',
      server: ''
    };

    // 验证邮箱
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!formData.email.trim()) {
      newErrors.email = '请输入电子邮箱';
      valid = false;
    } else if (!emailRegex.test(formData.email)) {
      newErrors.email = '请输入有效的电子邮箱';
      valid = false;
    }

    // 验证密码
    if (!formData.password) {
      newErrors.password = '请输入密码';
      valid = false;
    }

    setErrors(newErrors);
    return valid;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validateForm()) {
      return;
    }

    setIsSubmitting(true);
    setErrors(prev => ({ ...prev, server: '' }));

    try {
      // 使用AuthContext的login方法
      await login(formData.email, formData.password);
      
      // 登录成功，跳转到首页
      router.push('/');
    } catch (error: any) {
      console.error('登录失败:', error);
      setErrors(prev => ({
        ...prev,
        server: error.message || '登录失败，请检查您的邮箱和密码'
      }));
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="w-full max-w-md p-8 space-y-8 bg-white rounded-lg shadow-md">
      <div className="text-center">
        <h1 className="text-2xl font-bold">登录</h1>
        <p className="mt-2 text-sm text-gray-600">
          欢迎回来，请登录您的账号
        </p>
      </div>

      <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
        {/* 服务器错误提示 */}
        {errors.server && (
          <div className="p-3 text-sm text-red-700 bg-red-100 rounded-md">
            {errors.server}
          </div>
        )}

        <div>
          <label htmlFor="email" className="block text-sm font-medium text-gray-700">
            电子邮箱
          </label>
          <input
            id="email"
            name="email"
            type="email"
            required
            value={formData.email}
            onChange={handleChange}
            className="block w-full px-3 py-2 mt-1 placeholder-gray-400 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="请输入电子邮箱"
          />
          {errors.email && (
            <p className="mt-1 text-xs text-red-500">{errors.email}</p>
          )}
        </div>

        <div>
          <div className="flex items-center justify-between">
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              密码
            </label>
            <Link href="/auth/forgot-password" className="text-xs text-blue-600 hover:text-blue-500">
              忘记密码?
            </Link>
          </div>
          <input
            id="password"
            name="password"
            type="password"
            required
            value={formData.password}
            onChange={handleChange}
            className="block w-full px-3 py-2 mt-1 placeholder-gray-400 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="请输入密码"
          />
          {errors.password && (
            <p className="mt-1 text-xs text-red-500">{errors.password}</p>
          )}
        </div>

        <div className="flex items-center">
          <input
            id="remember-me"
            name="remember-me"
            type="checkbox"
            className="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
          />
          <label htmlFor="remember-me" className="block ml-2 text-sm text-gray-700">
            记住我
          </label>
        </div>

        <div>
          <button
            type="submit"
            disabled={isSubmitting}
            className="flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {isSubmitting ? '登录中...' : '登录'}
          </button>
        </div>

        <div className="text-sm text-center">
          <p className="text-gray-600">
            还没有账号？
            <Link href="/auth/register" className="font-medium text-blue-600 hover:text-blue-500">
              注册
            </Link>
          </p>
        </div>
      </form>
    </div>
  );
} 