import { Dashboard } from '../components/Dashboard';
import { Faq } from '../components/Faq';
import { InputForm } from '../components/InputForm';

export default function App() {
  return (
    <div className='w-full h-min overflow-y-scroll pb-20 px-6'>
      <Dashboard />
      <InputForm />
      <Faq />
    </div>
  );
}
