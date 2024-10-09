import { getStatisticsValues } from '../api/getStatisticsValues';

export const Dashboard = () => {
  const { data, isLoading, isSuccess } = getStatisticsValues();
  return (
    <div className='w-full h-max mt-2 flex flex-col gap-2'>
      <h1 className='text-[27px] font-climate'>Dashboard</h1>

      <div className='w-full h-[110px] bg-uiBgGradient rounded-32 px-4 py-3'>
        <div className='w-full h-full px-2 py-[6px] bg-white bg-opacity-10 rounded-3xl flex flex-row justify-between items-center'>
          <div className='w-1/2 h-full px-5 py-2 bg-black flex justify-center items-start flex-col rounded-3xl'>
            <h3 className='text-gray-300 text-base'>Total Supply</h3>
            <p className='text-white text-lg'>{isLoading ? 'Loading...' : isSuccess && data ? data[0].Total : '0'}$</p>
          </div>

          <div className='w-1/2 h-full px-5 py-2 bg-transparent flex justify-center items-start flex-col rounded-3xl'>
            <h3 className='text-gray-300 text-base'>Total Earn</h3>
            <p className='text-white text-lg'>{isLoading ? 'Loading...' : isSuccess && data ? data[1].Total : '0'}$</p>
          </div>
        </div>
      </div>
    </div>
  );
};
