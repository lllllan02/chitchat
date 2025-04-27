import Link from 'next/link';
import React, { useState } from 'react';

const Header: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <header className="bg-white shadow-sm">
      <div className="container mx-auto px-4 py-3 flex justify-between items-center">
        <div className="flex items-center">
          <Link href="/" className="text-xl font-bold text-gray-800 hover:text-gray-700">
            ChitChat
          </Link>
        </div>

        {/* 导航链接 - 桌面端 */}
        <nav className="hidden md:flex space-x-6">
          <Link href="/" className="text-gray-600 hover:text-gray-900">
            首页
          </Link>
          <Link href="/posts" className="text-gray-600 hover:text-gray-900">
            讨论区
          </Link>
          <Link href="/categories" className="text-gray-600 hover:text-gray-900">
            分类
          </Link>
        </nav>

        {/* 用户区域 */}
        <div className="hidden md:flex items-center space-x-4">
          {isLoggedIn ? (
            <>
              <Link href="/notifications" className="text-gray-600 hover:text-gray-900">
                <span className="relative">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                  </svg>
                  <span className="absolute -top-1 -right-1 bg-red-500 text-white rounded-full w-4 h-4 text-xs flex items-center justify-center">3</span>
                </span>
              </Link>
              <div className="relative">
                <button 
                  className="flex items-center focus:outline-none"
                  onClick={() => setIsMenuOpen(!isMenuOpen)}
                >
                  <div className="w-8 h-8 rounded-full bg-gray-300 flex items-center justify-center text-gray-700">
                    U
                  </div>
                </button>
                {isMenuOpen && (
                  <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-10">
                    <Link href="/profile" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                      个人主页
                    </Link>
                    <Link href="/profile/settings" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                      设置
                    </Link>
                    <button 
                      className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                      onClick={() => setIsLoggedIn(false)}
                    >
                      退出登录
                    </button>
                  </div>
                )}
              </div>
            </>
          ) : (
            <>
              <Link href="/auth/login" className="text-gray-600 hover:text-gray-900">
                登录
              </Link>
              <Link href="/auth/register" className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md">
                注册
              </Link>
            </>
          )}
        </div>

        {/* 移动端菜单按钮 */}
        <div className="md:hidden">
          <button 
            className="text-gray-600 hover:text-gray-900 focus:outline-none"
            onClick={() => setIsMenuOpen(!isMenuOpen)}
          >
            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16m-7 6h7" />
            </svg>
          </button>
        </div>
      </div>

      {/* 移动端菜单 */}
      {isMenuOpen && (
        <div className="md:hidden bg-white py-2 px-4 shadow-inner">
          <nav className="flex flex-col space-y-2">
            <Link href="/" className="text-gray-600 hover:text-gray-900 py-2">
              首页
            </Link>
            <Link href="/posts" className="text-gray-600 hover:text-gray-900 py-2">
              讨论区
            </Link>
            <Link href="/categories" className="text-gray-600 hover:text-gray-900 py-2">
              分类
            </Link>
            <div className="border-t border-gray-200 my-2"></div>
            {isLoggedIn ? (
              <>
                <Link href="/profile" className="text-gray-600 hover:text-gray-900 py-2">
                  个人主页
                </Link>
                <Link href="/notifications" className="text-gray-600 hover:text-gray-900 py-2">
                  通知消息
                </Link>
                <Link href="/profile/settings" className="text-gray-600 hover:text-gray-900 py-2">
                  设置
                </Link>
                <button 
                  className="text-left text-gray-600 hover:text-gray-900 py-2"
                  onClick={() => setIsLoggedIn(false)}
                >
                  退出登录
                </button>
              </>
            ) : (
              <>
                <Link href="/auth/login" className="text-gray-600 hover:text-gray-900 py-2">
                  登录
                </Link>
                <Link href="/auth/register" className="text-gray-600 hover:text-gray-900 py-2">
                  注册
                </Link>
              </>
            )}
          </nav>
        </div>
      )}
    </header>
  );
};

export default Header; 