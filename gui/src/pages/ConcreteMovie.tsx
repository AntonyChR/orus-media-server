import { useContext, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { DataContext } from '../providers/dataProvider/context';
import VideoPlayer from '../components/VideoPlayer';
import SelectSubtitle from '../components/SelectSubtitle';
import ApiDb, { VideoInfo } from '../data_fetching/data_fetching';

const ConcreteMovie = () => {
    const [videoInfo, setVideoInfo] = useState<VideoInfo|null>(null);
    const { movieId } = useParams();
    const { movies } = useContext(DataContext);

    let titleInfo = null;

    for (let i = 0; i < movies.length; i++) {
        if (movies[i].ID == Number(movieId)) {
            titleInfo = movies[i];
            break;
        }
    }

    useEffect(() => {
        if (!titleInfo) return;
        ApiDb.getVideoInfo(titleInfo?.ID).then((resp: VideoInfo | null) => {
            if (resp) {
                setVideoInfo(resp);
            }
        });
    }, [titleInfo]);

    useEffect(() => {
        return () => {
            setVideoInfo(null);
        };
    }, []);

    return (
        <div>
            {titleInfo && videoInfo && (
                <div>
                    <div className='grid grid-cols-2 w-full'>
                        <h1 className='text-white col-span-1'>
                            {titleInfo.Title}
                        </h1>
                        <SelectSubtitle
                            videoId={videoInfo.ID}
                            className='col-span-1 flex justify-end'
                        />
                    </div>
                    <div className='flex justify-center w h-full'>
                        <VideoPlayer
                            className='w-[80vw] h-[80vh]'
                            src={videoInfo.stream}
                            videoId={videoInfo.ID}
                            poster={titleInfo.Poster}
                        />
                    </div>
                </div>
            )}
        </div>
    );
};

export default ConcreteMovie;
