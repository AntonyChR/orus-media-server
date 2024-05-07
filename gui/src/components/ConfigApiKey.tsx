import { FC, useContext, useState } from 'react';
import ApiDb from '../data_fetching/data_fetching';
import Loading from './Loading';
import { t } from 'i18next';
import ResetDatabase from './ResetDatabase';
import { AlertContext } from './Alert/AlertContext';
interface ConfigApiKeyProps {
    className?: string;
}

const ConfigApiKey: FC<ConfigApiKeyProps> = ({ className }) => {
    const [loading, setLoading] = useState(false);

    const {showAlert} = useContext(AlertContext)

    const onSetApiKey = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setLoading(true);
        const apiKey = (event.currentTarget.children[1] as HTMLInputElement)
            .value;
        //makeRequest(apiKey);
        const resp = await ApiDb.setApiKey(apiKey)

        if(resp === null){
            showAlert({alertType:'success', message:t('Api key was added'), timeout:5000})
        }else{
            showAlert({alertType:'error', message:t('Invalid api key'), timeout:5000})
        }
        setLoading(false);
    };

    return (
        <div className={`${className}`}>
            <h2 className='text-white text-xl'>{t('Info provider')}</h2>
            <p className='text-white'>{t('Select info provider')}</p>
            <form onSubmit={onSetApiKey}>
                <select className='mr-2 my-2' id='selectInfoProvider'>
                    <option value='omdb'>OMDB</option>
                    <option disabled value='imdb'>
                        IMDB
                    </option>
                    <option disabled value='tmdb'>
                        TMDB
                    </option>
                </select>
                <input
                    className=''
                    type='text'
                    placeholder='api key'
                    required
                />
                <p className='text-white italic text-sm my-1'>
                    <span className='text-red-500'>* </span>
                    {t('Get api key')}{' '}
                    <a
                        className='text-blue-500'
                        target='_blank'
                        href='https://www.omdbapi.com/apikey.aspx'
                    >
                        https://www.omdbapi.com/apikey.aspx
                    </a>
                </p>
                <p className='text-white italic font-semibold text-sm my-1'>
                    <span className='text-red-500'>* </span>
                    {t('After adding api key')}
                </p>
                <div className='flex my-2'>
                    <button
                        className='text-white red-button'
                        type='submit'
                        disabled={loading}
                    >
                        {t('Save')}
                    </button>
                    {loading && <Loading />}
                    <ResetDatabase />
                </div>
            </form>
        </div>
    );
};

export default ConfigApiKey;
