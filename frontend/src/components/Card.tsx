import { HTMLAttributes } from 'react';

export default function Card({ ...props }: HTMLAttributes<HTMLDivElement>) {
  return <div className="bg-white shadow-md rounded-lg p-6" {...props} />;
}
