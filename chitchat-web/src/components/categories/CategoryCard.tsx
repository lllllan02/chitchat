import Link from 'next/link';
import React from 'react';

interface CategoryCardProps {
  id: number;
  name: string;
  description: string;
  postCount: number;
  icon?: string;
}

const CategoryCard: React.FC<CategoryCardProps> = ({
  id,
  name,
  description,
  postCount,
  icon,
}) => {
  // 默认图标
  const defaultIcon = (
    <svg xmlns="http://www.w3.org/2000/svg" className="h-12 w-12 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
    </svg>
  );

  return (
    <Link href={`/categories/${id}`} className="block">
      <div className="bg-white rounded-lg shadow-md p-6 border border-gray-100 hover:shadow-lg transition-shadow duration-200 flex items-start">
        <div className="mr-4">
          {icon ? (
            <img src={icon} alt={name} className="h-12 w-12" />
          ) : (
            defaultIcon
          )}
        </div>
        <div className="flex-grow">
          <h3 className="text-xl font-bold text-gray-800">{name}</h3>
          <p className="mt-2 text-gray-600">{description}</p>
          <div className="mt-3 text-sm text-gray-500">
            帖子数量: {postCount}
          </div>
        </div>
      </div>
    </Link>
  );
};

export default CategoryCard; 