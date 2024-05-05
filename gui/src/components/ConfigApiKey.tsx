import { FC } from 'react';
import { useWrapFetch } from '../hooks/useWrapFetch';
import ApiDb from '../data_fetching/data_fetching';
import Loading from './Loading';
import { t } from 'i18next';
interface ConfigApiKeyProps {
    className?: string;
}

const ConfigApiKey: FC<ConfigApiKeyProps> = ({ className }) => {

    const {makeRequest,loading, data} = useWrapFetch<Error|null,string>(ApiDb.setApiKey);

    const onSetApiKey = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const apiKey = (event.currentTarget.children[1] as HTMLInputElement).value;
        makeRequest(apiKey);
    }

    return (
        <div className={`${className}`}>
            <h2 className='text-white text-xl'>{t("Info provider")}</h2>
            <p className='text-white'>
                {t("Select info provider")}
            </p>
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
                <input className='' type='text' placeholder='api key' />
                <p className='text-white italic text-md my-1'>
                    <span className='text-red-500'>* </span>{t("Get api key")}{' '}
                    <a
                        className='text-blue-500'
                        target='_blank'
                        href='https://www.omdbapi.com/apikey.aspx'
                    >
                        https://www.omdbapi.com/apikey.aspx
                    </a>
                </p>
                <p className='text-white italic text-md my-1'>
                    <span className='text-red-500'>* </span>{t("After adding api key")}
                </p>
                <button
                    className='text-white block red-button my-2'
                    type='submit'
                >
                    {t('Save')}
                </button>{loading && <Loading/>}
            </form>
            {data && <p className='text-red-500'>{t("Invalid api key")}</p>}
        </div>
    );
};

export default ConfigApiKey;
