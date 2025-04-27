import React, { useState } from 'react';
import Layout from '../../components/layout/Layout';
import Link from 'next/link';
import { useRouter } from 'next/router';

// 模拟分类数据
const MOCK_CATEGORIES = [
  { id: 1, name: '综合讨论' },
  { id: 2, name: '技术交流' },
  { id: 3, name: '生活分享' },
  { id: 4, name: '兴趣爱好' },
  { id: 5, name: '求助问答' },
];

const NewPostPage: React.FC = () => {
  const router = useRouter();
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [category, setCategory] = useState('');
  const [tags, setTags] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [previewMode, setPreviewMode] = useState(false);

  // 处理提交
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim() || !content.trim() || !category) {
      alert('请填写必填项（标题、内容、分类）');
      return;
    }

    setIsSubmitting(true);
    
    // 模拟提交，实际项目中这里会调用API
    setTimeout(() => {
      alert('发布成功！');
      router.push('/posts');
    }, 1000);
  };

  // 处理标签输入
  const handleTagInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTags(e.target.value);
  };

  // 分段显示预览内容
  const renderPreviewContent = (content: string) => {
    return content.split('\n\n').map((paragraph, index) => (
      <p key={index} className="mb-4">{paragraph}</p>
    ));
  };

  return (
    <Layout>
      <div className="container mx-auto px-4 py-8">
        {/* 页面标题 */}
        <div className="flex items-center justify-between mb-6">
          <h1 className="text-2xl font-bold text-gray-900">发布新帖子</h1>
          <Link href="/posts" className="text-blue-500 hover:text-blue-700 flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
            </svg>
            返回讨论区
          </Link>
        </div>

        {/* 编辑/预览切换 */}
        <div className="bg-white rounded-lg shadow-md overflow-hidden mb-6">
          <div className="border-b border-gray-200">
            <div className="flex">
              <button
                onClick={() => setPreviewMode(false)}
                className={`py-3 px-6 font-medium focus:outline-none ${
                  !previewMode 
                    ? 'text-blue-600 border-b-2 border-blue-600' 
                    : 'text-gray-500 hover:text-gray-700'
                }`}
              >
                编辑
              </button>
              <button
                onClick={() => setPreviewMode(true)}
                className={`py-3 px-6 font-medium focus:outline-none ${
                  previewMode 
                    ? 'text-blue-600 border-b-2 border-blue-600' 
                    : 'text-gray-500 hover:text-gray-700'
                }`}
              >
                预览
              </button>
            </div>
          </div>

          {!previewMode ? (
            // 编辑模式
            <form onSubmit={handleSubmit} className="p-6">
              {/* 标题 */}
              <div className="mb-6">
                <label htmlFor="title" className="block text-sm font-medium text-gray-700 mb-1">
                  标题 <span className="text-red-500">*</span>
                </label>
                <input
                  type="text"
                  id="title"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="请输入帖子标题（5-100字）"
                  maxLength={100}
                  required
                />
              </div>

              {/* 分类 */}
              <div className="mb-6">
                <label htmlFor="category" className="block text-sm font-medium text-gray-700 mb-1">
                  分类 <span className="text-red-500">*</span>
                </label>
                <select
                  id="category"
                  value={category}
                  onChange={(e) => setCategory(e.target.value)}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                >
                  <option value="">请选择分类</option>
                  {MOCK_CATEGORIES.map((cat) => (
                    <option key={cat.id} value={cat.id}>
                      {cat.name}
                    </option>
                  ))}
                </select>
              </div>

              {/* 标签 */}
              <div className="mb-6">
                <label htmlFor="tags" className="block text-sm font-medium text-gray-700 mb-1">
                  标签
                </label>
                <input
                  type="text"
                  id="tags"
                  value={tags}
                  onChange={handleTagInput}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="添加标签，多个标签用逗号分隔"
                />
                <p className="mt-1 text-sm text-gray-500">最多添加5个标签，每个标签最多10个字</p>
              </div>

              {/* 内容 */}
              <div className="mb-6">
                <label htmlFor="content" className="block text-sm font-medium text-gray-700 mb-1">
                  内容 <span className="text-red-500">*</span>
                </label>
                <textarea
                  id="content"
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="请输入帖子内容..."
                  rows={15}
                  required
                ></textarea>
              </div>

              {/* 提交按钮 */}
              <div className="flex justify-end">
                <button
                  type="submit"
                  className="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50"
                  disabled={isSubmitting}
                >
                  {isSubmitting ? '发布中...' : '发布帖子'}
                </button>
              </div>
            </form>
          ) : (
            // 预览模式
            <div className="p-6">
              {title ? (
                <div>
                  <h1 className="text-2xl font-bold text-gray-900 mb-4">{title}</h1>
                  
                  <div className="flex items-center space-x-2 mb-6">
                    {category && (
                      <span className="inline-block bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full">
                        {MOCK_CATEGORIES.find(cat => cat.id.toString() === category)?.name || ''}
                      </span>
                    )}
                    
                    {tags && 
                      tags.split(',').map((tag, index) => (
                        tag.trim() && (
                          <span key={index} className="inline-block bg-gray-100 text-gray-800 text-xs px-2 py-1 rounded-full">
                            #{tag.trim()}
                          </span>
                        )
                      ))
                    }
                  </div>
                  
                  <div className="prose max-w-none text-gray-800">
                    {content ? renderPreviewContent(content) : (
                      <p className="text-gray-400 italic">帖子内容为空</p>
                    )}
                  </div>
                </div>
              ) : (
                <div className="flex items-center justify-center h-64">
                  <p className="text-gray-400">请先添加标题和内容以预览帖子</p>
                </div>
              )}
              
              <div className="border-t border-gray-200 mt-6 pt-6 flex justify-end">
                <button
                  onClick={() => setPreviewMode(false)}
                  className="px-4 py-2 mr-2 border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
                >
                  返回编辑
                </button>
                
                <button
                  onClick={handleSubmit}
                  className="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50"
                  disabled={isSubmitting || !title.trim() || !content.trim() || !category}
                >
                  {isSubmitting ? '发布中...' : '发布帖子'}
                </button>
              </div>
            </div>
          )}
        </div>

        {/* 发帖指南 */}
        <div className="bg-white rounded-lg shadow-md overflow-hidden">
          <div className="p-6">
            <h2 className="text-lg font-medium text-gray-900 mb-4">发帖指南</h2>
            <ul className="space-y-2 text-sm text-gray-600">
              <li className="flex items-start">
                <svg className="h-5 w-5 text-blue-500 mr-2 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>请选择合适的分类，以便其他用户更容易找到您的帖子</span>
              </li>
              <li className="flex items-start">
                <svg className="h-5 w-5 text-blue-500 mr-2 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>标题应当简洁明了，能够概括帖子的主要内容</span>
              </li>
              <li className="flex items-start">
                <svg className="h-5 w-5 text-blue-500 mr-2 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>内容请尽量详细，以便其他用户能够理解您的问题或观点</span>
              </li>
              <li className="flex items-start">
                <svg className="h-5 w-5 text-blue-500 mr-2 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>请添加相关标签，这有助于其他用户发现您的帖子</span>
              </li>
              <li className="flex items-start">
                <svg className="h-5 w-5 text-blue-500 mr-2 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>请遵守社区规则，不发布违规、低质量或无意义的内容</span>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default NewPostPage; 