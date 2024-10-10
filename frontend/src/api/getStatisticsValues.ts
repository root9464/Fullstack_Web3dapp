/* eslint-disable react-hooks/rules-of-hooks */
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

type ObjectType = {
  name: string;
  total: number | number[];
  percent: number;
};

const FnGetter = async () => {
  return await axios.get<ObjectType[]>(`https://anytapton.ru/go/api/record`);
};

export const getStatisticsValues = () => {
  return useQuery({
    queryKey: ['dashboard'],
    queryFn: FnGetter,
    select: (data) => data.data,
  });
};
