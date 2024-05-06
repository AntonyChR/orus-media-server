import { t } from 'i18next';
import ApiDb from '../data_fetching/data_fetching';
import { useContext, useState } from 'react';
import { AlertContext } from './Alert/AlertContext';
import Loading from './Loading';


const ResetDatabase = () => {
    const [loading, setLoading] = useState(false);
    const {showAlert} = useContext(AlertContext);
    const onClick = () => {

        if(!confirm(t('Are you sure?'))) return;
        setLoading(true);
        ApiDb.resetDatabase().then((resp)=>{
            if(resp === null){
                window.location.reload();
            }else{
                showAlert({alertType: 'error', message: `Error: ${ resp.message}`, timeout: 5000});
                setLoading(false);
            }
        })
        ;

    }

    return (
            <div className='flex'>
                <button className='red-button' disabled={loading} onClick={onClick}>
                    {t('Reset database')}
                </button>
                {loading && <Loading/>}
            </div>
    );
};

export default ResetDatabase;
