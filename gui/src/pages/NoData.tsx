import { Link } from 'react-router-dom';

const NoData = () => {
    return (
        <div>
            <p className='text-white'>
                no information available, check{' '}
                <Link to={'/config'} className='text-red-500 underline'>
                    configuration
                </Link>
            </p>
        </div>
    );
};

export default NoData;
