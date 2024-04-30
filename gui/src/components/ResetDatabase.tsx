import ApiDb from '../data_fetching/data_fetching';
import { useWrapFetch } from '../hooks/useWrapFetch';

import spinner from '/spinner.svg';

const ResetDatabase = () => {
    const {loading, makeRequest } = useWrapFetch<Error | null>(
        ApiDb.resetDatabase
    );
    return (
        <div className='w-full flex flex-col'>
            <div className='flex justify-center'>
                <button className='red-button' disabled={loading} onClick={makeRequest}>
                    Reset database
                </button>
                {loading && <img className='inline mb-2 animate-spin' width={20} src={spinner} />}
            </div>
        </div>
    );
};

export default ResetDatabase;
