import { useContext, useEffect, useState } from 'react';
import { NavLink, useNavigate, useParams } from 'react-router-dom';
import { DataContext } from '../providers/dataProvider/context';
import VideoPlayer from '../components/VideoPlayer';
import { Video } from '../types/Video';
import ApiDb from '../data_fetching/data_fetching';

const ConcreteSeries = () => {
    const [videoSrc, setVideoSrc] = useState('');
    const [chapters, setChapters] = useState<Video[] | null>(null);

    const { seriesId, chapterId } = useParams();
    const { series } = useContext(DataContext);
    const navigate = useNavigate();
    let titleInfo = null;

    for (let i = 0; i < series.length; i++) {
        if (series[i].ID == Number(seriesId)) {
            titleInfo = series[i];
            break;
        }
    }

    useEffect(() => {
        if (chapterId) {
            setVideoSrc(ApiDb.getVideoChapterSrc(chapterId));
        }
    }, [chapterId]);

    useEffect(() => {
        if (titleInfo) {
            ApiDb.getChapters(titleInfo.ID).then((resp) => {
                if (!resp) return;
                const orderedChapters = resp.sort((a, b) => {
                    if (a.Season > b.Season) {
                        return 1;
                    } else if (a.Season < b.Season) {
                        return -1;
                    } else {
                        return a.Episode - b.Episode;
                    }
                });
                setChapters(orderedChapters);
            });
        }
    }, [titleInfo]);

    useEffect(() => {
        if (chapters && chapters.length>0 && chapterId == '0') {
            navigate(`/series/${titleInfo!.ID}/${chapters[0].ID}`);
        }
        // eslint-disable-next-line
    }, [chapters]);

    const onNextChapter = () => {
        if (!chapters) return;

        for (let i = 0; i < chapters?.length; i++) {
            if (chapters[i].ID == Number(chapterId)) {
                if (i < chapters.length - 1) {
                    navigate(`/series/${titleInfo!.ID}/${chapters[i + 1].ID}`);
                }
                break;
            }
        }
    };

    const onPrevChapter = () => {
        if (!chapters) return;
        for (let i = 0; i < chapters.length; i++) {
            if (chapters[i].ID == Number(chapterId)) {
                if (i > 0) {
                    navigate(`/series/${titleInfo!.ID}/${chapters[i - 1].ID}`);
                }
                break;
            }
        }
    };

    useEffect(() => {
        return () => {
            setChapters(null);
            setVideoSrc('');
        };
    }, []);
    return (
        <>
            {titleInfo && (
                <div className='grid grid-cols-6'>
                    <div className='text-white col-span-6 flex justify-between my-3'>
                        <h1 className='font-bold text-xl'>{titleInfo.Title}</h1>
                        <div className='justify-end flex gap-x-3'>
                            <button
                                className='red-button px-2'
                                onClick={onPrevChapter}
                            >
                                Prev
                            </button>
                            <button
                                className='red-button px-2'
                                onClick={onNextChapter}
                            >
                                Next
                            </button>
                        </div>
                    </div>



                    <div className='col-span-6 xl:col-span-5 flex justify-center flex-col'>
                        <VideoPlayer
                            className='h-[80vh]'
                            src={videoSrc}
                            poster={titleInfo.Poster}
                            videoId={Number(chapterId)}
                        />
                    </div>
                    <div className='max-h-[80vh] col-span-6 xl:col-span-1'>
                        <div className='text-white flex flex-col bg-gray-900 overflow-y-auto h-full'>
                            {chapters &&
                                chapters.map((c) => (
                                    <NavLink
                                        key={c.Name}
                                        className={({ isActive }) =>
                                            `${
                                                isActive
                                                    ? 'bg-red-700'
                                                    : 'bg-gray-900'
                                            } hover:bg-red-900`
                                        }
                                        to={`/series/${titleInfo.ID}/${c.ID}`}
                                    >
                                        Chapter {c.Episode} - Season: {c.Season}
                                    </NavLink>
                                ))}
                        </div>
                    </div>
                </div>
            )}
        </>
    );
};

export default ConcreteSeries;
