import { useQuery } from '@tanstack/react-query';
import { useTonAddress } from '@tonconnect/ui-react';
import axios from 'axios';
import { FC } from 'react';
import TonLogo from '../assets/ton.png';

type TransactionsUser = {
  id: number;
  user_wallet_address: string;
  deposite_date: string;
  receiving_date: string;
  amount: string;
  rewards: string;
};
export const Profile = () => {
  const address = useTonAddress();

  const { data } = useQuery({
    queryKey: ['profile'],
    queryFn: async () => axios.get<TransactionsUser[]>(`https://earnton.ru/go/api/transaction/${address}`),
    select: (data) => data.data,
  });

  console.log(data);

  return (
    <div className='w-full h-screen mt-9 px-6 '>
      <div className='w-full h-max max-h-[700px] overflow-y-scroll flex flex-col gap-y-4'>
        {data?.map((item) => (
          <TransactionBLock
            key={item.id}
            id={item.id}
            amount={item.amount}
            rewards={item.rewards}
            depositDate={item.deposite_date}
            receivingDate={item.receiving_date}
          />
        ))}
      </div>
    </div>
  );
};

type TransactionBLockProps = {
  id: number;
  amount: string;
  rewards: string;
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

    <div className='col-span-1 row-span-1 flex flex-col gap-2 w-max'>
      <p>Deposit Date: {depositDate}</p>
      <p>Receiving Date: {receivingDate}</p>
    </div>
  </div>
);
