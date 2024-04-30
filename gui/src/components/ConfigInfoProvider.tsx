import { FC } from 'react';
interface ConfigInfoProviderProps {
    className?: string;
}

const ConfigInfoProvider: FC<ConfigInfoProviderProps> = ({ className }) => {
    return (
        <div className={`${className}`}>
            <h2 className='text-white text-xl'>Info provider</h2>
            <p className='text-white'>
                Select the info provider you want to use to get information
                about the videos
            </p>
            <form action='' className=''>
                <select className='mr-2 my-2' id='selectInfoProvider'>
                    <option value='omdb'>OMDB</option>
                    <option disabled value='imdb'>IMDB</option>
                    <option disabled value='tmdb'>TMDB</option>
                </select>
                <input className='' type='text' placeholder='api key' />
                <p className='text-white italic text-sm'>
                    <span className='text-red-500'>* </span>You can get an API key from{' '}
                    <a
                        className='text-blue-500'
                        target='_blank'
                        href='https://www.omdbapi.com/apikey.aspx'
                    >
                        https://www.omdbapi.com/apikey.aspx
                    </a>
                </p>
                <button

                    className='text-white block bg-red-700 p-1 rounded-md my-2 active:bg-red-400'
                    type='submit'
                >
                    Save
                </button>
            </form>
        </div>
    );
};

export default ConfigInfoProvider;
