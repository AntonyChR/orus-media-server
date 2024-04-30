import ConfigInfoProvider from '../components/ConfigInfoProvider';
import ResetDatabase from '../components/ResetDatabase';
import VideosWithNoInfo from '../components/VideosWithNoInfo';

const Config = () => {
    return (
        <div className='grid grid-cols-1 md:grid-cols-2'>
            <VideosWithNoInfo   className='col-span-1 p-3'/>
            <ConfigInfoProvider className='col-span-1 p-3'/>
            <ResetDatabase/>
        </div>
    );
};

export default Config;
