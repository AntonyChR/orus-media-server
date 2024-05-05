import { t } from 'i18next';
import { Link } from 'react-router-dom';

const NoData = () => {
    return (
        <div>
            <p className='text-white'>
                {t("No info available")}{' '}
                <Link to={'/config'} className='text-red-500 underline'>
                    {t("Config")}
                </Link>
            </p>
        </div>
    );
};

export default NoData;
