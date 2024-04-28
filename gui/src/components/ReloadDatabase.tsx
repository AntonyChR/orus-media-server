import ApiDb from '../data_fetching/data_fetching';
import { useWrapFetch } from '../hooks/useWrapFetch';

import spinner from '/spinner.svg';

const ReloadDatabase = () => {
    const {loading, makeRequest } = useWrapFetch<Error | null>(
        ApiDb.resetDatabase
    );
    return (
        <div className='w-full flex flex-col'>
            <p className='text-white text-center'>
                The database appears to be empty.
            </p>
            <div className='flex justify-center'>
                <button className='bg-red-800 p-2 rounded-md m-3' disabled={loading} onClick={makeRequest}>
                    Reset database to get information about existing files 
                </button>
                {!loading && <img className='inline mb-2 animate-spin' width={20} src={spinner} />}
            </div>
        </div>
    );
};

export default ReloadDatabase;
