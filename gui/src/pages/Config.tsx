import ConfigApiKey from '../components/ConfigApiKey';
import FileStructureInfo from '../components/FileStructureInfo';
import ServerLogs from '../components/ServerLogs';
import VideosWithNoInfo from '../components/VideosWithNoInfo';

const Config = () => {
    return (
        <div className='grid grid-cols-1 md:grid-cols-2 h-full overflow-y-scroll [&>*]:px-3 [&>*]:my-4'>
            <FileStructureInfo/>
            <VideosWithNoInfo/>
            <ConfigApiKey/>
            <ServerLogs className='md:col-span-2' />
        </div>
    );
};

export default Config;
