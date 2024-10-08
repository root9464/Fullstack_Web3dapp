import { Link } from '@tanstack/react-router';
import { DashboardIcon } from '../assets/tsx/Dashboard';
import { ProfileIcon } from '../assets/tsx/Profile';

export const Menu = () => (
  <div className='w-full h-[71px] absolute left-0 bottom-0 mt-5 bg-black flex flex-row justify-center items-center gap-x-5 py-3'>
    <div className='w-fit h-min flex flex-col justify-center items-center'>
      <DashboardIcon fill='white' className='w-8 h-8' />
      <Link to='/'>Dashboard</Link>
    </div>

    <div className='w-fit h-min flex flex-col justify-center items-center'>
      <ProfileIcon fill='white' className='w-8 h-8' />
      <Link to='/profile'>Profile</Link>
    </div>
  </div>
);
