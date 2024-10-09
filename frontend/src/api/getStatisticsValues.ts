/* eslint-disable react-hooks/rules-of-hooks */
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

type ObjectType = {
  Name: string;
  Total: number;
  Percent: number;
};

const FnGetter = async () => {
  return await axios.get<ObjectType[]>('http://127.0.0.1:8080/api/record');
};

export const getStatisticsValues = () => {
  return useQuery({
    queryKey: ['dashboard'],
    queryFn: FnGetter,
    select: (data) => data.data,
  });
};
