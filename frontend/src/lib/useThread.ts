'use client';

import { useEffect, useState } from 'react';
import { useParams } from 'next/navigation';
import api from './api';
import { Thread } from '@/types/thread';

export const useThread = () => {
  const params = useParams();
  const { id } = params;
  const [thread, setThread] = useState<Thread | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (id) {
      api
        .get(`/threads/${id}`)
        .then((response) => {
          setThread(response.data);
        })
        .catch((error) => {
          console.error('Failed to fetch thread:', error);
        })
        .finally(() => {
          setLoading(false);
        });
    }
  }, [id]);

  return { thread, loading };
};
