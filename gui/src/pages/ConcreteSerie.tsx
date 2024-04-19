import { useContext, useEffect, useState } from 'react';
import { NavLink, useParams } from 'react-router-dom';
import { DataContext } from '../providers/dataProvider/context';
import {
    getChapters,
    getVideoChapterSrc,
} from '../data_fetching/data_fetching';
import VideoPlayer from '../components/VideoPlayer';
import { FileInfo } from '../types/FileInfo';

const ConcreteSeries = () => {
    const [videoSrc, setVideoSrc] = useState('');
    const [chapters, setChapters] = useState<FileInfo[] | null>(null);

    const { seriesId, chapterId } = useParams();
    const { series } = useContext(DataContext);
    let titleInfo = null;

    for (let i = 0; i < series.length; i++) {
        if (series[i].ID == Number(seriesId)) {
            titleInfo = series[i];
            break;
        }
    }

    useEffect(() => {
        if (chapterId) {
            setVideoSrc(getVideoChapterSrc(chapterId));
        }
    }, [chapterId]);

    useEffect(() => {
        if (titleInfo) {
            getChapters(titleInfo.ID).then((resp) => setChapters(resp));
        }
    }, [titleInfo]);
    return (
        <>
            {titleInfo && (
                <div className='grid grid-cols-6 h-full'>
                    <div className='border border-white max-h-full overflow-auto'>
                        <h1 className='text-white'>{titleInfo.Title}</h1>
                        <ul className='text-white'>
                            {chapters &&
                                chapters.map((c) => (
                                    <li key={c.Name} className='bg-slate-950'>
                                        <NavLink
                                            className={({ isActive }) =>
                                                `${
                                                    isActive
                                                        ? 'bg-red-700'
                                                        : 'bg-gray-950'
                                                } hover:bg-red-700`
                                            }
                                            to={`/series/${titleInfo.ID}/${c.ID}`}
                                        >
                                            capitulo {c.Episode}, season:{' '}
                                            {c.Season}
                                        </NavLink>
                                    </li>
                                ))}
                        </ul>
                    </div>
                    <div className='col-span-5 flex justify-center'>
                        <VideoPlayer
                            className='max-h-[80vh] w-full'
                            src={videoSrc}
                            poster={titleInfo.Poster}
                        />
                    </div>
                </div>
            )}
        </>
    );
};

export default ConcreteSeries;
