import { FC, ReactNode } from 'react';
import Navbar from '../components/Navbar';

interface layoutProps {
    children: ReactNode;
}

const Layout: FC<layoutProps> = ({ children }) => {
    return (
        <div className='w-screen h-screen bg-gray-950 overflow-auto'>
            <Navbar />
            <main className='pt-16 px-7 h-full md:overflow-hidden'>{children}</main>
        </div>
    );
};

export default Layout;
