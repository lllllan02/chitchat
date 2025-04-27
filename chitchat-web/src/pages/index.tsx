import React from 'react';
import Layout from '../components/layout/Layout';
import PostCard from '../components/posts/PostCard';
import CategoryCard from '../components/categories/CategoryCard';
import Link from 'next/link';

// 模拟数据
const MOCK_CATEGORIES = [
  {
    id: 1,
    name: '综合讨论',
    description: '各种话题的综合讨论区',
    postCount: 120,
  },
  {
    id: 2,
    name: '技术交流',
    description: '编程、技术相关的交流讨论',
    postCount: 85,
  },
  {
    id: 3,
    name: '生活分享',
    description: '日常生活经验分享',
    postCount: 64,
  },
  {
    id: 4,
    name: '兴趣爱好',
    description: '分享你的兴趣爱好',
    postCount: 52,
  },
];

const MOCK_POSTS = [
  {
    id: 1,
    title: '如何开始学习编程？新手指南',
    content: '编程是当今最有价值的技能之一，本文将分享如何从零开始学习编程。首先，选择一门适合初学者的语言，如Python或JavaScript。然后，找到优质的学习资源，如官方文档、在线课程或教程。接下来，通过实践项目巩固所学知识，从简单的项目开始，逐步挑战更复杂的任务。最后，加入编程社区，与其他学习者交流经验，获取反馈和建议。持续学习和实践是成为优秀程序员的关键。',
    author: {
      id: 1,
      username: '编程爱好者',
    },
    createdAt: '2023-11-15T08:30:00Z',
    category: {
      id: 2,
      name: '技术交流',
    },
    commentCount: 24,
    viewCount: 358,
    likeCount: 42,
    isPinned: true,
  },
  {
    id: 2,
    title: '分享我的户外摄影经验和器材推荐',
    content: '作为一名户外摄影爱好者，我想分享一些在野外拍摄的经验和推荐一些实用的器材。首先，选择一个适合自己的相机非常重要，如果你是初学者，可以考虑入门级单反或微单。其次，配备合适的镜头，广角镜头适合风景摄影，长焦镜头适合野生动物摄影。此外，三脚架、滤镜和备用电池也是户外摄影的必备装备。最后，学习构图技巧和后期处理也能大大提升照片质量。',
    author: {
      id: 2,
      username: '摄影师小王',
    },
    createdAt: '2023-11-14T14:20:00Z',
    category: {
      id: 4,
      name: '兴趣爱好',
    },
    commentCount: 18,
    viewCount: 246,
    likeCount: 36,
    isFeatured: true,
  },
  {
    id: 3,
    title: '今日热议：人工智能的未来发展趋势',
    content: '随着技术的快速发展，人工智能正在改变我们生活的方方面面。从自动驾驶汽车到智能助手，AI技术已经深入我们的日常生活。未来几年，AI可能会在医疗、教育和金融等领域带来更多创新。然而，随着AI的普及，我们也面临着一系列挑战，如数据隐私、算法偏见和就业变化等问题。因此，在推动AI发展的同时，我们也需要建立相应的伦理和监管框架。',
    author: {
      id: 3,
      username: '科技前沿',
    },
    createdAt: '2023-11-13T09:45:00Z',
    category: {
      id: 1,
      name: '综合讨论',
    },
    commentCount: 32,
    viewCount: 412,
    likeCount: 56,
  },
];

const Home: React.FC = () => {
  return (
    <Layout>
      {/* 欢迎横幅 */}
      <div className="bg-gradient-to-r from-blue-500 to-indigo-600 text-white py-16">
        <div className="container mx-auto px-4">
          <h1 className="text-4xl font-bold mb-4">欢迎来到 ChitChat</h1>
          <p className="text-xl max-w-2xl">一个自由分享知识、交流经验的社区平台</p>
          <div className="mt-8 flex space-x-4">
            <Link href="/posts" className="bg-white text-blue-600 px-6 py-2 rounded-md font-semibold hover:bg-gray-100 transition-colors duration-200">
              浏览讨论区
            </Link>
            <Link href="/auth/register" className="bg-transparent border-2 border-white text-white px-6 py-2 rounded-md font-semibold hover:bg-white hover:text-blue-600 transition-colors duration-200">
              立即注册
            </Link>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-4 py-10">
        {/* 分类板块 */}
        <section>
          <div className="flex justify-between items-center mb-6">
            <h2 className="text-2xl font-bold text-gray-800">分类板块</h2>
            <Link href="/categories" className="text-blue-600 hover:text-blue-800 font-medium">
              查看全部 &rarr;
            </Link>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            {MOCK_CATEGORIES.map(category => (
              <CategoryCard
                key={category.id}
                id={category.id}
                name={category.name}
                description={category.description}
                postCount={category.postCount}
              />
            ))}
          </div>
        </section>

        {/* 推荐帖子 */}
        <section className="mt-16">
          <div className="flex justify-between items-center mb-6">
            <h2 className="text-2xl font-bold text-gray-800">热门讨论</h2>
            <Link href="/posts" className="text-blue-600 hover:text-blue-800 font-medium">
              查看全部 &rarr;
            </Link>
          </div>
          <div className="space-y-6">
            {MOCK_POSTS.map(post => (
              <PostCard
                key={post.id}
                id={post.id}
                title={post.title}
                content={post.content}
                author={post.author}
                createdAt={post.createdAt}
                category={post.category}
                commentCount={post.commentCount}
                viewCount={post.viewCount}
                likeCount={post.likeCount}
                isPinned={post.isPinned}
                isFeatured={post.isFeatured}
              />
            ))}
          </div>
        </section>

        {/* 社区统计 */}
        <section className="mt-16 bg-gray-50 rounded-lg p-8">
          <h2 className="text-2xl font-bold text-gray-800 mb-8 text-center">社区数据</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="bg-white p-6 rounded-lg shadow-sm text-center">
              <div className="text-3xl font-bold text-blue-600 mb-2">1,245</div>
              <div className="text-gray-600">注册用户</div>
            </div>
            <div className="bg-white p-6 rounded-lg shadow-sm text-center">
              <div className="text-3xl font-bold text-blue-600 mb-2">3,872</div>
              <div className="text-gray-600">讨论帖子</div>
            </div>
            <div className="bg-white p-6 rounded-lg shadow-sm text-center">
              <div className="text-3xl font-bold text-blue-600 mb-2">12,564</div>
              <div className="text-gray-600">评论回复</div>
            </div>
          </div>
        </section>
      </div>
    </Layout>
  );
};

export default Home; 