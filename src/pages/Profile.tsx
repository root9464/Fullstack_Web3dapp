import { useQuery } from '@tanstack/react-query';
import { useTonAddress } from '@tonconnect/ui-react';
import axios from 'axios';
import { FC } from 'react';
import TonLogo from '../assets/ton.png';

type Body = {
  id: number;
  user_wallet_address: string;
  deposite_date: string;
  receiving_date: string;
  amount: number;
  rewards: number;
};

type Data = {
  id: number;
  user_wallet_address: string;
  hash: string;
  body_id: number;
  Body: Body;
};

export const Profile = () => {
  const address = useTonAddress();

  const { data } = useQuery({
    queryKey: ['profile'],
    queryFn: async () => axios.get<Data[]>(`https://localhost/api/transaction/${address}`),
    select: (data) => data.data,
    enabled: !!address,
  });

  console.log(data);

  return (
    <div className='w-full h-max mt-9 px-6 '>
      <div className='w-full h-max max-h-[700px] overflow-y-scroll flex flex-col gap-y-4'>
        {data?.map((item) => (
          <TransactionBLock
            key={item.id}
            amount={item.Body.amount}
            rewards={item.Body.rewards}
            depositDate={item.Body.deposite_date}
            receivingDate={item.Body.receiving_date}
          />
        ))}
      </div>
    </div>
  );
};

type TransactionBLockProps = {
  amount: number;
  rewards: number;
  depositDate: string;
  receivingDate: string;
};

const TransactionBLock: FC<TransactionBLockProps> = ({ amount, rewards, depositDate, receivingDate }) => (
  <div className='w-full h-max max-h-[700px] px-5 py-3 grid grid-cols-2 grid-rows-2 gap-5 bg-uiLowGray rounded-3xl'>
    <div className='col-span-1 row-span-1 flex flex-row gap-5'>
      <img src={TonLogo} alt='' className='w-6 h-6' />
      <h2>TON</h2>
    </div>

    <div className='col-span-1 row-span-1 flex flex-col gap-2 font-semibold text-sm'>
      <p>Amount: {amount} TON </p>
      <p>Rewards: {rewards} TON</p>
    </div>

    <div className='col-span-1 row-span-1 flex flex-col gap-2'>
      <p>Deposit Date: {depositDate}</p>
      <p>Receiving Date: {receivingDate}</p>
    </div>
  </div>
);
