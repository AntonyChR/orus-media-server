import { FC, useContext, useEffect, useState } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import { Subtitle } from '../types/Subtitle';
import { ENDPOINTS } from '../data_fetching/endpoints';
interface VideoPlayerProps {
    src: string;
    videoId: number;
    className?: string;
    poster?: string;
}

const VideoPlayer: FC<VideoPlayerProps> = ({
    src,
    className,
    poster,
    videoId,
}) => {
    const [subtitles, setSubtitles] = useState<Subtitle[]>([]);
    const { subtitles: allSubtitles } = useContext(DataContext);
    useEffect(() => {
        const videoSubtitles = allSubtitles.filter(
            (sub) => sub.VideoId == videoId
        );
        setSubtitles(videoSubtitles);
    }, [src, videoId, allSubtitles]);
    return (
            <video className={className} src={src}  controls crossOrigin='anonymous' poster={poster}>
                {subtitles.map((sub,i) => {
                    return (
                        <track
                            default={i === 0 ? true : false}
                            key={sub.Name}
                            src={`${ENDPOINTS.media.subtitlesServer}/${sub.Name}`}
                            srcLang={sub.Lang}
                            label={`${sub.Lang}-${sub.Name}`}
                            kind='subtitles'
                        />
                    );
                })}
            </video>
    );
};

export default VideoPlayer;
