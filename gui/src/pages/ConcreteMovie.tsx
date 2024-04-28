import { useContext, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { DataContext } from '../providers/dataProvider/context';
import VideoPlayer from '../components/VideoPlayer';
import SelectSubtitle from '../components/SelectSubtitle';
import ApiDb from '../data_fetching/data_fetching';

const ConcreteMovie = () => {
    const [videoSrc, setVideoSrc] = useState('');
    const { movieId } = useParams();
        const { movies, assignVideoIdToSubtitles } = useContext(DataContext);

    let titleInfo = null;

    for (let i = 0; i < movies.length; i++) {
        if (movies[i].ID == Number(movieId)) {
            titleInfo = movies[i];
            break;
        }
    }

    const onSelectSubtitle = (subtitleId: number) => {
        if (titleInfo) {
            assignVideoIdToSubtitles(titleInfo.ID, subtitleId);
        }
    }

    useEffect(() => {
        if (!titleInfo) return;
        ApiDb.getMovieSrc(titleInfo?.ID).then((resp: string | null) => {
            if (resp) {
                setVideoSrc(String(resp));
            }
        });
    }, [titleInfo]);

    useEffect(() => {
        return () => {
            setVideoSrc('');
        };
    }, []);

    return (
        <div>
            {titleInfo && (
                <div>
                    <div className='grid grid-cols-2 w-full'>
                        <h1 className='text-white col-span-1'>{titleInfo.Title}</h1>
                        <SelectSubtitle className='col-span-1 flex justify-end' onSelect={onSelectSubtitle}/>
                    </div>
                    <div className='flex justify-center w h-full'>
                        <VideoPlayer
                            className='w-[80vw] h-[80vh]'
                            src={videoSrc}
                            videoId={titleInfo.ID}
                            poster={titleInfo.Poster}
                        />
                    </div>
                </div>
            )}
        </div>
    );
};

export default ConcreteMovie;
