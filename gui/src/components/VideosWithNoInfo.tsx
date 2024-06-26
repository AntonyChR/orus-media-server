import { FC, useEffect, useState } from 'react';
import ApiDb from '../data_fetching/data_fetching';
import { Video } from '../types/Video';
import { t } from 'i18next';
interface VideosWithNoInfoProps {
    className?: string;
}

const VideosWithNoInfo: FC<VideosWithNoInfoProps> = ({ className }) => {
    const [videosWithNoInfo, setVideos] = useState<Video[]>([]);
    useEffect(() => {
        ApiDb.getVideoWithNoTitleInfo().then((videos) => {
            if (videos) {
                setVideos(videos);
            }
        });
    }, []);
    return (
        <div className={`${className}`}>
            <h2 className='text-white text-xl mb-3'>
                {t('Files with no info')}
            </h2>
            {videosWithNoInfo.length === 0 ? (
                <p className='text-white italic text-sm'>
                    {t('Here videos with no info')}
                </p>
            ) : (
                <>
                    <div className='text-white bg-gray-800 p-3'>
                        <ul className=''>
                            {videosWithNoInfo.map((video) => {
                                return (
                                    <li
                                        key={video.ID}
                                        className='hover:bg-gray-700 hover:cursor-pointer'
                                    >
                                        {video.Name}
                                    </li>
                                );
                            })}
                        </ul>
                    </div>
                </>
            )}
        </div>
    );
};

export default VideosWithNoInfo;
