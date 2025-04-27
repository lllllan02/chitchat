import React, { useState } from 'react';
import Layout from '../../components/layout/Layout';
import Link from 'next/link';
import Image from 'next/image';

// 模拟分类数据
const MOCK_CATEGORIES = [
  { id: 1, name: '综合讨论' },
  { id: 2, name: '技术交流' },
  { id: 3, name: '生活分享' },
  { id: 4, name: '兴趣爱好' },
  { id: 5, name: '求助问答' },
];

// 模拟热门标签
const MOCK_TAGS = [
  '前端', '后端', '设计', 'React', 'Vue', 'Python', 'Java', 
  '职场', '学习', '读书', '旅行', '摄影', '美食', '健身'
];

// 模拟帖子数据
const MOCK_POSTS = [
  {
    id: 1,
    title: '初学React，有什么好的学习资源推荐吗？',
    content: '最近开始学习React，想找一些系统的学习资源，希望大家能推荐一些好的教程、视频或者书籍，感谢！',
    author: {
      id: 101,
      name: 'React爱好者',
      avatar: '/images/avatars/avatar1.jpg'
    },
    category: { id: 2, name: '技术交流' },
    tags: ['React', '前端', '学习'],
    createdAt: new Date('2023-07-15T08:30:00'),
    viewCount: 342,
    likeCount: 56,
    commentCount: 24,
    isTop: true,
    isHot: true
  },
  {
    id: 2,
    title: '分享我的摄影作品集，记录城市的日与夜',
    content: '业余时间喜欢拍摄城市风光，这是我近期拍摄的一组照片，记录了城市从清晨到黄昏再到夜晚的变化...',
    author: {
      id: 102,
      name: '光影追逐者',
      avatar: '/images/avatars/avatar2.jpg'
    },
    category: { id: 4, name: '兴趣爱好' },
    tags: ['摄影', '城市', '风光'],
    createdAt: new Date('2023-07-14T15:45:00'),
    viewCount: 271,
    likeCount: 89,
    commentCount: 17,
    isTop: false,
    isHot: true
  },
  {
    id: 3,
    title: '远程工作一年的心得体会',
    content: '去年开始尝试远程工作，经过一年的实践，有一些心得想和大家分享，包括时间管理、沟通技巧、工作与生活平衡等...',
    author: {
      id: 103,
      name: '远程工作者',
      avatar: '/images/avatars/avatar3.jpg'
    },
    category: { id: 3, name: '生活分享' },
    tags: ['远程工作', '职场', '时间管理'],
    createdAt: new Date('2023-07-13T10:20:00'),
    viewCount: 498,
    likeCount: 132,
    commentCount: 45,
    isTop: false,
    isHot: false
  },
  {
    id: 4,
    title: 'Python爬虫实战：如何高效抓取网页数据',
    content: '本文介绍了使用Python进行网页数据抓取的几种常用方法，包括requests+BeautifulSoup、Scrapy框架的使用，以及如何处理反爬虫机制...',
    author: {
      id: 104,
      name: 'Python大师',
      avatar: '/images/avatars/avatar4.jpg'
    },
    category: { id: 2, name: '技术交流' },
    tags: ['Python', '爬虫', '数据分析'],
    createdAt: new Date('2023-07-12T14:30:00'),
    viewCount: 623,
    likeCount: 98,
    commentCount: 31,
    isTop: false,
    isHot: false
  },
  {
    id: 5,
    title: '如何培养高效的阅读习惯？',
    content: '在信息爆炸的时代，高效阅读变得尤为重要。本文分享了我坚持多年的阅读方法和习惯，希望能对大家有所帮助...',
    author: {
      id: 105,
      name: '阅读爱好者',
      avatar: '/images/avatars/avatar5.jpg'
    },
    category: { id: 3, name: '生活分享' },
    tags: ['读书', '习惯养成', '自我提升'],
    createdAt: new Date('2023-07-11T09:15:00'),
    viewCount: 376,
    likeCount: 72,
    commentCount: 28,
    isTop: false,
    isHot: false
  }
];

// 日期格式化函数
const formatDate = (date: Date) => {
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  
  // 小于24小时，显示"x小时前"
  if (diff < 24 * 60 * 60 * 1000) {
    const hours = Math.floor(diff / (60 * 60 * 1000));
    return hours <= 0 ? '刚刚' : `${hours}小时前`;
  }
  
  // 小于7天，显示"x天前"
  if (diff < 7 * 24 * 60 * 60 * 1000) {
    const days = Math.floor(diff / (24 * 60 * 60 * 1000));
    return `${days}天前`;
  }
  
  // 超过7天，显示具体日期
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
};

const PostsListPage: React.FC = () => {
  const [selectedCategory, setSelectedCategory] = useState<number | null>(null);
  const [selectedTag, setSelectedTag] = useState<string | null>(null);
  const [sortBy, setSortBy] = useState('latest'); // 'latest', 'hot', 'commented'
  
  // 过滤和排序帖子
  const filteredPosts = MOCK_POSTS
    .filter(post => 
      (selectedCategory === null || post.category.id === selectedCategory) &&
      (selectedTag === null || post.tags.includes(selectedTag))
    )
    .sort((a, b) => {
      // 首先按置顶排序
      if (a.isTop && !b.isTop) return -1;
      if (!a.isTop && b.isTop) return 1;
      
      // 然后按选择的排序方式
      switch (sortBy) {
        case 'latest':
          return b.createdAt.getTime() - a.createdAt.getTime();
        case 'hot':
          return b.viewCount - a.viewCount;
        case 'commented':
          return b.commentCount - a.commentCount;
        default:
          return 0;
      }
    });
  
  // 处理分类选择
  const handleCategoryClick = (categoryId: number) => {
    setSelectedCategory(selectedCategory === categoryId ? null : categoryId);
  };
  
  // 处理标签选择
  const handleTagClick = (tag: string) => {
    setSelectedTag(selectedTag === tag ? null : tag);
  };

  return (
    <Layout>
      <div className="container mx-auto px-4 py-8">
        {/* 页面头部 */}
        <div className="flex flex-col md:flex-row md:justify-between md:items-center mb-8">
          <h1 className="text-2xl font-bold text-gray-900 mb-4 md:mb-0">讨论区</h1>
          <Link href="/posts/new" className="inline-flex items-center justify-center px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
            </svg>
            发布新帖子
          </Link>
        </div>
        
        <div className="flex flex-col lg:flex-row gap-6">
          {/* 侧边栏 */}
          <div className="lg:w-1/4">
            {/* 分类过滤 */}
            <div className="bg-white rounded-lg shadow-md p-4 mb-6">
              <h2 className="text-lg font-medium text-gray-900 mb-3">分类</h2>
              <div className="space-y-2">
                {MOCK_CATEGORIES.map(category => (
                  <button
                    key={category.id}
                    onClick={() => handleCategoryClick(category.id)}
                    className={`w-full text-left px-3 py-2 rounded-md transition-colors ${
                      selectedCategory === category.id 
                        ? 'bg-blue-100 text-blue-700'
                        : 'hover:bg-gray-100 text-gray-700'
                    }`}
                  >
                    {category.name}
                  </button>
                ))}
              </div>
            </div>
            
            {/* 热门标签 */}
            <div className="bg-white rounded-lg shadow-md p-4">
              <h2 className="text-lg font-medium text-gray-900 mb-3">热门标签</h2>
              <div className="flex flex-wrap gap-2">
                {MOCK_TAGS.map(tag => (
                  <button
                    key={tag}
                    onClick={() => handleTagClick(tag)}
                    className={`px-3 py-1 rounded-full text-sm ${
                      selectedTag === tag
                        ? 'bg-blue-100 text-blue-700'
                        : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
                    }`}
                  >
                    #{tag}
                  </button>
                ))}
              </div>
            </div>
          </div>
          
          {/* 主内容区 */}
          <div className="lg:w-3/4">
            {/* 排序选项 */}
            <div className="bg-white rounded-lg shadow-md p-4 mb-6">
              <div className="flex items-center justify-between">
                <div className="flex space-x-4">
                  <button
                    onClick={() => setSortBy('latest')}
                    className={`px-3 py-1 rounded-md ${
                      sortBy === 'latest' ? 'bg-blue-100 text-blue-700' : 'text-gray-700'
                    }`}
                  >
                    最新
                  </button>
                  <button
                    onClick={() => setSortBy('hot')}
                    className={`px-3 py-1 rounded-md ${
                      sortBy === 'hot' ? 'bg-blue-100 text-blue-700' : 'text-gray-700'
                    }`}
                  >
                    热门
                  </button>
                  <button
                    onClick={() => setSortBy('commented')}
                    className={`px-3 py-1 rounded-md ${
                      sortBy === 'commented' ? 'bg-blue-100 text-blue-700' : 'text-gray-700'
                    }`}
                  >
                    多评论
                  </button>
                </div>
                
                <div className="text-sm text-gray-500">
                  共 {filteredPosts.length} 个帖子
                  {selectedCategory !== null && (
                    <span>
                      ，分类：{MOCK_CATEGORIES.find(c => c.id === selectedCategory)?.name}
                    </span>
                  )}
                  {selectedTag && <span>，标签：#{selectedTag}</span>}
                </div>
              </div>
            </div>
            
            {/* 帖子列表 */}
            <div className="space-y-4">
              {filteredPosts.length > 0 ? (
                filteredPosts.map(post => (
                  <div key={post.id} className="bg-white rounded-lg shadow-md overflow-hidden">
                    <div className="p-4">
                      {/* 帖子标题 */}
                      <div className="flex items-start mb-2">
                        {post.isTop && (
                          <span className="inline-block bg-red-100 text-red-800 text-xs px-2 py-1 rounded-full mr-2">
                            置顶
                          </span>
                        )}
                        {post.isHot && (
                          <span className="inline-block bg-orange-100 text-orange-800 text-xs px-2 py-1 rounded-full mr-2">
                            热门
                          </span>
                        )}
                        <Link href={`/posts/${post.id}`} className="text-lg font-medium text-gray-900 hover:text-blue-600">
                          {post.title}
                        </Link>
                      </div>
                      
                      {/* 帖子内容预览 */}
                      <div className="text-gray-600 mb-4 line-clamp-2">
                        {post.content}
                      </div>
                      
                      {/* 帖子元信息 */}
                      <div className="flex flex-wrap items-center text-sm text-gray-500">
                        <div className="flex items-center mr-4">
                          <div className="w-6 h-6 rounded-full overflow-hidden mr-1">
                            <Image
                              src={post.author.avatar}
                              alt={post.author.name}
                              width={24}
                              height={24}
                              className="object-cover"
                            />
                          </div>
                          <span>{post.author.name}</span>
                        </div>
                        
                        <div className="mr-4">
                          <span className="inline-block bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full">
                            {post.category.name}
                          </span>
                        </div>
                        
                        <div className="mr-4">
                          <time dateTime={post.createdAt.toISOString()}>
                            {formatDate(post.createdAt)}
                          </time>
                        </div>
                        
                        <div className="flex items-center space-x-3 mt-2 sm:mt-0">
                          <span className="flex items-center">
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                            </svg>
                            {post.viewCount}
                          </span>
                          
                          <span className="flex items-center">
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5" />
                            </svg>
                            {post.likeCount}
                          </span>
                          
                          <span className="flex items-center">
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
                            </svg>
                            {post.commentCount}
                          </span>
                        </div>
                      </div>
                      
                      {/* 标签 */}
                      {post.tags.length > 0 && (
                        <div className="mt-3 flex flex-wrap gap-2">
                          {post.tags.map(tag => (
                            <button
                              key={tag}
                              onClick={() => handleTagClick(tag)}
                              className="text-xs text-gray-500 hover:text-blue-600"
                            >
                              #{tag}
                            </button>
                          ))}
                        </div>
                      )}
                    </div>
                  </div>
                ))
              ) : (
                <div className="bg-white rounded-lg shadow-md p-8 text-center">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-12 w-12 mx-auto text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <p className="mt-4 text-gray-500">暂无符合条件的帖子</p>
                  <button
                    onClick={() => {
                      setSelectedCategory(null);
                      setSelectedTag(null);
                    }}
                    className="mt-2 text-blue-600 hover:text-blue-800"
                  >
                    清除筛选条件
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default PostsListPage; 