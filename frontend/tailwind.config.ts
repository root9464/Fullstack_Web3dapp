import daisyui from 'daisyui';
import { type Config } from 'tailwindcss';

const config: Config = {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        uiMainBgColor: '#151515',
        uiPurple: '#6A00C6',
        uiLowGray: '#3F3F3F',
        // uiBgGradient: 'bg-gradient-to-r from-[#34343480] via-[#4E4E4E80] to-[#34343480]',
        // uiPurpleGradient: 'bg-gradient-to-l from-[#343434] via-[#8800ff] to-[#343434]',
      },

      fontFamily: {
        climate: ['"Climate Crisis"', 'sans-serif'],
      },

      backgroundImage: {
        uiBgGradient: 'linear-gradient(265.14deg, rgba(52, 52, 52, 0.5) -2.3%, rgba(136, 0, 255, 0.5) 32.55%, rgba(52, 52, 52, 0.5) 101.09%)',
        uiGrayGradient:
          'linear-gradient(224deg, rgba(52.24, 52.24, 52.24, 0.50) 0%, rgba(77.53, 77.53, 77.53, 0.50) 34%, rgba(52.24, 52.24, 52.24, 0.50) 100%)',
      },

      borderRadius: {
        '32': '32px',
      },
    },
  },
  plugins: [daisyui],
  daisyui: {
    base: false,
  },
};

export default config;
