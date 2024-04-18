import { FC, ReactNode } from 'react';
import Navbar from '../components/Navbar';

interface layoutProps {
    children: ReactNode;
}

const Layout: FC<layoutProps> = ({ children }) => {
    return (
        <div className='w-full h-screen bg-gray-950 overflow-auto'>
            <Navbar />
            <main className='pt-16 px-7'>{children}</main>
        </div>
    );
};

export default Layout;
