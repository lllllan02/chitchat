import Link from 'next/link';
import React from 'react';

interface PostCardProps {
  id: number;
  title: string;
  content: string;
  author: {
    id: number;
    username: string;
    avatar?: string;
  };
  createdAt: string;
  category: {
    id: number;
    name: string;
  };
  commentCount: number;
  viewCount: number;
  likeCount: number;
  isPinned?: boolean;
  isFeatured?: boolean;
}

const PostCard: React.FC<PostCardProps> = ({
  id,
  title,
  content,
  author,
  createdAt,
  category,
  commentCount,
  viewCount,
  likeCount,
  isPinned,
  isFeatured
}) => {
  // 格式化时间
  const formattedDate = new Date(createdAt).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  });

  // 截取内容预览
  const contentPreview = content.length > 120 ? `${content.substring(0, 120)}...` : content;

  return (
    <div className="bg-white rounded-lg shadow-md p-5 border border-gray-100 hover:shadow-lg transition-shadow duration-200">
      <div className="flex justify-between items-start">
        <div>
          {/* 帖子标题与链接 */}
          <Link href={`/posts/${id}`} className="text-xl font-bold text-gray-800 hover:text-blue-600 transition-colors duration-200">
            <h2 className="flex items-center">
              {isPinned && (
                <span className="mr-2 px-2 py-1 bg-red-100 text-red-600 text-xs rounded-full">
                  置顶
                </span>
              )}
              {isFeatured && (
                <span className="mr-2 px-2 py-1 bg-yellow-100 text-yellow-600 text-xs rounded-full">
                  精华
                </span>
              )}
              {title}
            </h2>
          </Link>
          
          {/* 分类标签 */}
          <div className="mt-2">
            <Link href={`/categories/${category.id}`}>
              <span className="inline-block bg-gray-100 text-gray-600 px-2 py-1 rounded text-sm hover:bg-gray-200 transition-colors duration-200">
                {category.name}
              </span>
            </Link>
          </div>
        </div>
        
        {/* 作者信息 */}
        <Link href={`/profile/${author.id}`} className="flex items-center text-sm text-gray-600 ml-4">
          <div className="w-8 h-8 rounded-full bg-gray-300 flex items-center justify-center overflow-hidden">
            {author.avatar ? (
              <img src={author.avatar} alt={author.username} className="w-full h-full object-cover" />
            ) : (
              <span className="text-gray-700">{author.username.charAt(0).toUpperCase()}</span>
            )}
          </div>
          <span className="ml-2">{author.username}</span>
        </Link>
      </div>
      
      {/* 帖子内容预览 */}
      <p className="mt-3 text-gray-600">{contentPreview}</p>
      
      {/* 帖子元数据信息 */}
      <div className="mt-4 flex items-center justify-between">
        <div className="text-gray-500 text-sm">
          {formattedDate}
        </div>
        
        <div className="flex items-center space-x-4 text-sm text-gray-500">
          <div className="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
            <span>{viewCount}</span>
          </div>
          
          <div className="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
            </svg>
            <span>{commentCount}</span>
          </div>
          
          <div className="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
            <span>{likeCount}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default PostCard; 