import { t } from 'i18next';
import ApiDb from '../data_fetching/data_fetching';
import { useWrapFetch } from '../hooks/useWrapFetch';
import Loading from './Loading';


const ResetDatabase = () => {
    const {loading, makeRequest } = useWrapFetch<Error | null>(
        ApiDb.resetDatabase
    );
    return (
            <div className='flex'>
                <button className='red-button' disabled={loading} onClick={makeRequest}>
                    {t('Reset database')}
                </button>
                {loading && <Loading/>}
            </div>
    );
};

export default ResetDatabase;
