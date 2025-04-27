import { NextResponse } from 'next/server';
import type { UserRegisterData } from '@/types/user';

// 这是一个模拟函数，实际项目中应该连接数据库
async function registerUser(userData: UserRegisterData) {
  // 模拟检查邮箱是否已注册
  if (userData.email === 'test@example.com') {
    throw new Error('邮箱已被注册');
  }

  // 模拟创建用户并返回用户信息（不含密码）
  return {
    id: Math.floor(Math.random() * 1000),
    username: userData.username,
    email: userData.email,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  };
}

export async function POST(request: Request) {
  try {
    const userData: UserRegisterData = await request.json();

    // 基本验证
    if (!userData.username || !userData.email || !userData.password) {
      return NextResponse.json(
        { error: '请提供用户名、邮箱和密码' },
        { status: 400 }
      );
    }

    // 邮箱格式验证
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(userData.email)) {
      return NextResponse.json(
        { error: '请提供有效的邮箱地址' },
        { status: 400 }
      );
    }

    // 密码长度验证
    if (userData.password.length < 6) {
      return NextResponse.json(
        { error: '密码至少需要6个字符' },
        { status: 400 }
      );
    }

    // 注册用户
    const user = await registerUser(userData);

    // 返回成功响应
    return NextResponse.json(
      { 
        message: '注册成功', 
        user 
      },
      { status: 201 }
    );
  } catch (error: any) {
    console.error('注册失败:', error);
    
    // 处理已知错误
    if (error.message === '邮箱已被注册') {
      return NextResponse.json(
        { error: '该邮箱已被注册' },
        { status: 409 }
      );
    }
    
    // 处理其他错误
    return NextResponse.json(
      { error: '注册失败，请稍后再试' },
      { status: 500 }
    );
  }
} 