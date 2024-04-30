import { FC, useEffect, useState } from 'react';
import ApiDb from '../data_fetching/data_fetching';
import { Video } from '../types/Video';
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
        <div className={`${className}   `}>
            <h2 className='text-white text-xl'>
                Video files with no information
            </h2>
            {videosWithNoInfo.length === 0 ? (
                <p className='text-white italic text-sm'>
                    Here you will see the videos that have no information
                </p>
            ) : (
                <>
                    <div className='text-white p-3'>
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
