import ConfigInfoProvider from '../components/ConfigInfoProvider';
import VideosWithNoInfo from '../components/VideosWithNoInfo';

const Config = () => {
    return (
        <div className='grid grid-cols-2'>
            <VideosWithNoInfo   className='col-span-1 p-3'/>
            <ConfigInfoProvider className='col-span-1 p-3'/>
        </div>
    );
};

export default Config;
