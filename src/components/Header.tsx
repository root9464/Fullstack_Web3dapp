import { TonConnectButton } from '@tonconnect/ui-react';
import Logo from '../assets/logotest.png';
import '../index.css';

export const Header = () => (
  <header className='w-full h-8 mt-5 flex flex-row justify-between'>
    <img src={Logo} alt='' />
    <TonConnectButton />
  </header>
);
