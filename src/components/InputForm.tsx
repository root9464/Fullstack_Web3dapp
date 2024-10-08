import { FC, useState } from 'react';
import TonLogo from '../assets/ton.png';
export const InputForm = () => {
  const [activeButton, setActiveButton] = useState(1);

  return (
    <div className='w-full h-max mt-4 bg-uiGrayGradient flex flex-col gap-3 rounded-32 px-5 py-4'>
      <h2 className='text-2xl font-climate'>Market Overview</h2>

      <div className='w-full h-8 px-2 py-1 flex flex-row gap-x-1 bg-uiLowGray rounded-32'>
        <ToggleButton label='1 day' isActive={activeButton === 1} onClick={() => setActiveButton(1)} />
        <ToggleButton label='7 day' isActive={activeButton === 7} onClick={() => setActiveButton(7)} />
        <ToggleButton label='30 day' isActive={activeButton === 30} onClick={() => setActiveButton(30)} />
      </div>

      <div className='w-full h-fit text-gray-400 flex flex-row justify-between text-xs'>
        <p>Assets</p>
        <p>Earnings for {activeButton} day</p>
      </div>

      <div className='w-full h-max px-4 py-5 flex flex-col gap-2 bg-uiLowGray rounded-32'>
        <HeaderForm activeButton={activeButton} />
        <input
          type='text'
          placeholder='Amount'
          className='w-[110px] bg-[#3F3F3F] rounded-[4px] border border-[#818181] outline-none px-2 py-1 text-sm'
        />
        <p className='text-gray-400 text-xs'>Min. 1 - Max. 8594</p>
      </div>
    </div>
  );
};

const HeaderForm: FC<{ activeButton: number }> = ({ activeButton }) => (
  <div className='w-full h-full flex flex-row justify-between'>
    <div className='w-max h-full flex flex-row items-center gap-x-3'>
      <img src={TonLogo} alt='' className='w-5' />
      <h2>TON</h2>
    </div>
    <div className='w-max h-full flex flex-row items-center gap-x-3'>
      <p>103.58К</p>
      <p>${activeButton === 1 ? 1 : activeButton === 7 ? 9 : 40}%</p>
    </div>
  </div>
);

type ToggleButtonProps = {
  label: string;
  isActive: boolean;
  onClick: () => void;
} & Partial<HTMLButtonElement>;

const ToggleButton: FC<ToggleButtonProps> = ({ label, isActive, onClick }) => (
  <button className={`w-1/3 h-full rounded-full ${isActive ? 'bg-black' : 'bg-[#3F3F3F]'}`} onClick={onClick}>
    {label}
  </button>
);
