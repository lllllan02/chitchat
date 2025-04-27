import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Layout from '../../components/layout/Layout';
import Link from 'next/link';
import Image from 'next/image';

// 模拟帖子数据
const MOCK_POST = {
  id: 1,
  title: '如何开始学习编程？新手指南',
  content: '编程是当今最有价值的技能之一，本文将分享如何从零开始学习编程。\n\n首先，选择一门适合初学者的语言，如Python或JavaScript。Python语法简洁清晰，有丰富的库和资源，适合各种应用场景；JavaScript则是网页开发的基础，学习后可以直接在浏览器中看到效果，非常直观。\n\n然后，找到优质的学习资源，如官方文档、在线课程或教程。推荐几个学习平台：Codecademy、freeCodeCamp、Coursera和edX等，它们提供免费或付费的结构化课程。\n\n接下来，通过实践项目巩固所学知识，从简单的项目开始，逐步挑战更复杂的任务。比如，可以先做一个简单的计算器、待办事项列表或个人博客。实践是学习编程最有效的方式，通过解决实际问题，你会更深入地理解编程概念。\n\n最后，加入编程社区，与其他学习者交流经验，获取反馈和建议。GitHub、Stack Overflow和Reddit上的编程社区都是很好的选择。参与开源项目也是提升编程能力的好方法。\n\n持续学习和实践是成为优秀程序员的关键。编程领域发展迅速，新的技术和工具不断涌现，保持好奇心和学习热情非常重要。',
  author: {
    id: 1,
    username: '编程爱好者',
    avatar: '/images/avatar-1.jpg'
  },
  createdAt: '2023-11-15T08:30:00Z',
  updatedAt: '2023-11-15T10:15:00Z',
  category: {
    id: 2,
    name: '技术交流',
  },
  tags: ['编程', '学习方法', '新手指南'],
  commentCount: 24,
  viewCount: 358,
  likeCount: 42,
  isPinned: true,
};

// 模拟评论数据
const MOCK_COMMENTS = [
  {
    id: 1,
    content: '非常感谢分享！我刚开始学习编程，这些建议对我很有帮助。请问有没有推荐的Python入门书籍？',
    author: {
      id: 10,
      username: '学习者小李',
      avatar: '/images/avatar-2.jpg'
    },
    createdAt: '2023-11-15T09:45:00Z',
    likeCount: 5,
    replies: [
      {
        id: 3,
        content: '《Python编程：从入门到实践》是个不错的选择，浅显易懂，有很多实践项目。',
        author: {
          id: 1,
          username: '编程爱好者',
          avatar: '/images/avatar-1.jpg'
        },
        createdAt: '2023-11-15T10:20:00Z',
        likeCount: 3,
      }
    ]
  },
  {
    id: 2,
    content: '关于加入编程社区的建议很棒！我在Stack Overflow和GitHub上学到了很多东西。对于初学者来说，不要害怕提问和分享自己的代码，社区里的人通常很乐意帮助新人。',
    author: {
      id: 12,
      username: '代码大师',
      avatar: '/images/avatar-3.jpg'
    },
    createdAt: '2023-11-15T11:30:00Z',
    likeCount: 8,
    replies: []
  },
  {
    id: 4,
    content: '我认为选择编程语言时，还应该考虑你的目标领域。如果想做Web开发，JavaScript是必学的；如果对数据科学感兴趣，Python是最佳选择；如果想做移动应用，可以考虑学习Swift（iOS）或Kotlin（Android）。',
    author: {
      id: 15,
      username: '技术专家',
      avatar: '/images/avatar-4.jpg'
    },
    createdAt: '2023-11-15T14:05:00Z',
    likeCount: 12,
    replies: [
      {
        id: 5,
        content: '完全同意！目标导向的学习效率更高。不过对于完全的新手，我还是建议先从Python或JavaScript入手，因为它们的学习曲线相对平缓，能够快速看到成果，增强学习动力。',
        author: {
          id: 1,
          username: '编程爱好者',
          avatar: '/images/avatar-1.jpg'
        },
        createdAt: '2023-11-15T14:30:00Z',
        likeCount: 7,
      }
    ]
  }
];

const PostDetailPage: React.FC = () => {
  const router = useRouter();
  const { id } = router.query;
  const [commentText, setCommentText] = useState('');
  const [liked, setLiked] = useState(false);
  const [likeCount, setLikeCount] = useState(MOCK_POST.likeCount);

  // 处理评论提交
  const handleCommentSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (commentText.trim()) {
      alert('评论功能将在后端实现后生效');
      setCommentText('');
    }
  };

  // 处理点赞
  const handleLike = () => {
    if (!liked) {
      setLikeCount(prev => prev + 1);
    } else {
      setLikeCount(prev => prev - 1);
    }
    setLiked(!liked);
  };

  // 格式化日期
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('zh-CN', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  // 分段显示帖子内容
  const renderPostContent = (content: string) => {
    return content.split('\n\n').map((paragraph, index) => (
      <p key={index} className="mb-4">{paragraph}</p>
    ));
  };

  return (
    <Layout>
      <div className="container mx-auto px-4 py-8">
        {/* 返回按钮 */}
        <div className="mb-4">
          <Link href="/posts" className="flex items-center text-blue-500 hover:text-blue-700">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
            </svg>
            返回讨论区
          </Link>
        </div>

        {/* 帖子标题和信息 */}
        <div className="bg-white rounded-lg shadow-md overflow-hidden mb-6">
          <div className="p-6">
            <div className="flex items-center justify-between mb-4">
              <span className="inline-block bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full">
                {MOCK_POST.category.name}
              </span>
              <div className="flex space-x-2">
                {MOCK_POST.tags.map((tag, index) => (
                  <span key={index} className="inline-block bg-gray-100 text-gray-800 text-xs px-2 py-1 rounded-full">
                    #{tag}
                  </span>
                ))}
              </div>
            </div>
            
            <h1 className="text-3xl font-bold text-gray-900 mb-4">{MOCK_POST.title}</h1>
            
            <div className="flex items-center mb-6">
              <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-200 mr-3">
                {MOCK_POST.author.avatar ? (
                  <Image
                    src={MOCK_POST.author.avatar}
                    alt={MOCK_POST.author.username}
                    width={40}
                    height={40}
                  />
                ) : (
                  <div className="w-full h-full flex items-center justify-center bg-blue-500 text-white text-xl font-bold">
                    {MOCK_POST.author.username.charAt(0).toUpperCase()}
                  </div>
                )}
              </div>
              <div>
                <p className="font-medium text-gray-900">{MOCK_POST.author.username}</p>
                <p className="text-sm text-gray-500">
                  发布于 {formatDate(MOCK_POST.createdAt)}
                  {MOCK_POST.updatedAt !== MOCK_POST.createdAt && 
                    ` · 编辑于 ${formatDate(MOCK_POST.updatedAt)}`}
                </p>
              </div>
            </div>
            
            <div className="prose max-w-none text-gray-800">
              {renderPostContent(MOCK_POST.content)}
            </div>
            
            <div className="flex items-center justify-between mt-8 pt-4 border-t border-gray-200">
              <div className="flex items-center space-x-6">
                <button 
                  onClick={handleLike}
                  className={`flex items-center ${liked ? 'text-red-500' : 'text-gray-500 hover:text-red-500'}`}
                >
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 mr-1" fill={liked ? "currentColor" : "none"} viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                  </svg>
                  <span>{likeCount}</span>
                </button>
                
                <div className="flex items-center text-gray-500">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                  <span>{MOCK_POST.viewCount}</span>
                </div>
                
                <div className="flex items-center text-gray-500">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                  </svg>
                  <span>{MOCK_POST.commentCount}</span>
                </div>
              </div>
              
              <div className="flex space-x-2">
                <button className="text-gray-500 hover:text-blue-500">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                  </svg>
                </button>
                <button className="text-gray-500 hover:text-blue-500">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        {/* 评论区 */}
        <div className="bg-white rounded-lg shadow-md overflow-hidden mb-6">
          <div className="p-6">
            <h2 className="text-xl font-bold text-gray-900 mb-6">评论 ({MOCK_COMMENTS.length})</h2>
            
            {/* 评论输入 */}
            <form onSubmit={handleCommentSubmit} className="mb-8">
              <div className="flex mb-2">
                <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-200 mr-3">
                  <div className="w-full h-full flex items-center justify-center bg-blue-500 text-white text-xl font-bold">
                    游
                  </div>
                </div>
                <textarea
                  className="flex-grow p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  rows={4}
                  placeholder="写下你的评论..."
                  value={commentText}
                  onChange={(e) => setCommentText(e.target.value)}
                ></textarea>
              </div>
              <div className="flex justify-end">
                <button
                  type="submit"
                  className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50"
                  disabled={!commentText.trim()}
                >
                  发表评论
                </button>
              </div>
            </form>
            
            {/* 评论列表 */}
            <div className="space-y-6">
              {MOCK_COMMENTS.map((comment) => (
                <div key={comment.id} className="border-b border-gray-200 pb-6 last:border-b-0 last:pb-0">
                  <div className="flex mb-3">
                    <div className="w-10 h-10 rounded-full overflow-hidden bg-gray-200 mr-3">
                      {comment.author.avatar ? (
                        <Image
                          src={comment.author.avatar}
                          alt={comment.author.username}
                          width={40}
                          height={40}
                        />
                      ) : (
                        <div className="w-full h-full flex items-center justify-center bg-blue-500 text-white text-xl font-bold">
                          {comment.author.username.charAt(0).toUpperCase()}
                        </div>
                      )}
                    </div>
                    <div className="flex-grow">
                      <div className="flex items-start justify-between">
                        <div>
                          <p className="font-medium text-gray-900">{comment.author.username}</p>
                          <p className="text-sm text-gray-500">{formatDate(comment.createdAt)}</p>
                        </div>
                        <button className="text-gray-400 hover:text-gray-600">
                          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z" />
                          </svg>
                        </button>
                      </div>
                      <div className="mt-2 text-gray-800">
                        <p>{comment.content}</p>
                      </div>
                      <div className="mt-3 flex items-center space-x-4">
                        <button className="text-gray-500 hover:text-blue-500 text-sm flex items-center">
                          <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5" />
                          </svg>
                          {comment.likeCount}
                        </button>
                        <button className="text-gray-500 hover:text-blue-500 text-sm flex items-center">
                          <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6" />
                          </svg>
                          回复
                        </button>
                      </div>
                    </div>
                  </div>
                  
                  {/* 嵌套回复 */}
                  {comment.replies.length > 0 && (
                    <div className="ml-12 mt-4 space-y-4">
                      {comment.replies.map((reply) => (
                        <div key={reply.id} className="flex">
                          <div className="w-8 h-8 rounded-full overflow-hidden bg-gray-200 mr-3">
                            {reply.author.avatar ? (
                              <Image
                                src={reply.author.avatar}
                                alt={reply.author.username}
                                width={32}
                                height={32}
                              />
                            ) : (
                              <div className="w-full h-full flex items-center justify-center bg-blue-500 text-white text-sm font-bold">
                                {reply.author.username.charAt(0).toUpperCase()}
                              </div>
                            )}
                          </div>
                          <div className="flex-grow">
                            <div className="flex items-start justify-between">
                              <div>
                                <p className="font-medium text-gray-900">{reply.author.username}</p>
                                <p className="text-xs text-gray-500">{formatDate(reply.createdAt)}</p>
                              </div>
                            </div>
                            <div className="mt-1 text-gray-800 text-sm">
                              <p>{reply.content}</p>
                            </div>
                            <div className="mt-2 flex items-center space-x-4">
                              <button className="text-gray-500 hover:text-blue-500 text-xs flex items-center">
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-3 w-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5" />
                                </svg>
                                {reply.likeCount}
                              </button>
                              <button className="text-gray-500 hover:text-blue-500 text-xs flex items-center">
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-3 w-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6" />
                                </svg>
                                回复
                              </button>
                            </div>
                          </div>
                        </div>
                      ))}
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* 相关推荐 */}
        <div className="bg-white rounded-lg shadow-md overflow-hidden">
          <div className="p-6">
            <h2 className="text-xl font-bold text-gray-900 mb-4">相关讨论</h2>
            <div className="space-y-4">
              <div className="border-b border-gray-200 pb-3">
                <Link href="#" className="block hover:text-blue-500">
                  <h3 className="font-medium text-gray-900 mb-1">编程学习资源汇总：从入门到精通</h3>
                  <p className="text-sm text-gray-500">23 评论 · 340 次浏览</p>
                </Link>
              </div>
              <div className="border-b border-gray-200 pb-3">
                <Link href="#" className="block hover:text-blue-500">
                  <h3 className="font-medium text-gray-900 mb-1">新手必看：常见的编程错误和解决方法</h3>
                  <p className="text-sm text-gray-500">18 评论 · 285 次浏览</p>
                </Link>
              </div>
              <div>
                <Link href="#" className="block hover:text-blue-500">
                  <h3 className="font-medium text-gray-900 mb-1">为什么学习编程要选择正确的方向？</h3>
                  <p className="text-sm text-gray-500">31 评论 · 412 次浏览</p>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default PostDetailPage; 