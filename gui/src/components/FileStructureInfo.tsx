import { FC } from 'react';
import { t } from 'i18next';

import folder from '/folder.svg';
import video from '/video.svg';
import text from '/text.svg';


interface FileStructureInfoProps {
    className?: string;
}

const FileStructureInfo: FC<FileStructureInfoProps> = ({ className }) => {
    return (
        <div className={`${className} text-white`}>
            <h2 className='text-xl'>{t("Directory structure")}</h2>
            <p>{t('The directory structure should be like:')}</p>
            <ul className='my-4 bg-gray-800 p-2'>
                <li><img className='inline mr-1' width={20} src={folder}/>media/</li>
                <li><img className='ml-8 inline mr-1' width={20} src={video}/>Godzilla(2014).mov</li>
                <li><img className='ml-8 inline mr-1' width={20} src={video}/>The matrix(1999).mp4</li>
                <li><img className='ml-8 inline mr-1' width={20} src={folder}/>The flash(2013)/</li>
                <li className='ml-14'><img className='inline mr-1' width={20} src={video}/>s1e1.mp4</li>
                <li className='ml-14'><img className='inline mr-1' width={20} src={video}/>s1e2.mp4</li>
                <li className='ml-14'><img className='inline mr-1' width={20} src={video}/>s1e3.mp4</li>
                <li><img className='inline mr-1' width={20} src={folder}/>subtitles/</li>
                <li className='ml-8'><img className='inline mr-1' width={20} src={text}/>Godzilla(2014).en.vtt</li>
            </ul>
            <p className='text-sm'><span className='text-red-500'>* </span>{t('Movie name format')}</p>
            <p className='text-sm'><span className='text-red-500'>* </span>{t('Serie name format')}</p>
        </div>
    );
};

export default FileStructureInfo;
