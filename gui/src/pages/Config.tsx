import ConfigInfoProvider from '../components/ConfigInfoProvider';
import ServerLogs from '../components/ServerLogs';
import VideosWithNoInfo from '../components/VideosWithNoInfo';

const Config = () => {
    return (
        <div className='grid grid-cols-1 md:grid-cols-2 md:h-full grid-rows-3 md:grid-rows-2'>
            <ConfigInfoProvider className='col-span-1 row-span-1 p-3'/>
            <VideosWithNoInfo   className='col-span-1 row-span-1 row-start-3 md:row-start-1 p-3'/>
            <ServerLogs className='col-span-1 md:col-span-2 md:h-full row-span-1'/>
        </div>
    );
};

export default Config;
