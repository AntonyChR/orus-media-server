import { NavLink } from 'react-router-dom';

// active class
const aC = (isActive: boolean) => (isActive ? 'text-red-500' : 'text-gray-400');

const Navbar = () => {
    return (
        <nav className='fixed w-full bg-gray-950 z-20'>
            <ul className='grid grid-cols-2'>
                <li className='col-span-1 justify-start flex gap-4 p-4'>
                    <NavLink
                        to={'/movies'}
                        className={({ isActive: v }) => `${aC(v)}`}
                    >
                        Movies
                    </NavLink>
                    <NavLink
                        to={'/series'}
                        className={({ isActive: v }) => `${aC(v)}`}
                    >
                        Series
                    </NavLink>
                </li>
                <li className={'col-span-1 justify-end flex p-4'}>
                    <NavLink
                        className={({ isActive: v }) => `${aC(v)}`}
                        to={'/config'}
                    >
                        Config
                    </NavLink>
                </li>
            </ul>
        </nav>
    );
};

export default Navbar;
