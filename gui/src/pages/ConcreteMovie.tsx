import { useContext, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { DataContext } from '../providers/dataProvider/context';
import VideoPlayer from '../components/VideoPlayer';
import { getMovieSrc } from '../data_fetching/data_fetching';

const ConcreteMovie = () => {
    const [videoSrc, setVideoSrc] = useState("")

    const { movieId } = useParams();
    const { movies } = useContext(DataContext);
    let titleInfo = null;

    for (let i = 0; i < movies.length; i++) {
        if (movies[i].ID == Number(movieId)) {
            titleInfo = movies[i];
            break;
        }
    }

    useEffect(()=>{
        if(!titleInfo)return;
        getMovieSrc(titleInfo?.ID).then((resp:string |null)=>{
            if (resp){
                //console.log("video id",resp)
                setVideoSrc(String(resp))
            }
        })
    },[titleInfo])

    return (
        <div>
            {titleInfo && (
                <div>
                    <h1 className='text-white'>{titleInfo.Title}</h1>
                    <div className='flex justify-center w h-full'>
                        <VideoPlayer className='w-[80vw] h-[80vh]' src={videoSrc} poster={titleInfo.Poster} />
                    </div>
                </div>
            )}
        </div>
    );
};

export default ConcreteMovie;
