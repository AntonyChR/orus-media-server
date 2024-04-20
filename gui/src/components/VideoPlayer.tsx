import { FC } from 'react';
interface VideoPlayerProps {
    src: string;
    className?: string;
    poster?: string;
}

const VideoPlayer: FC<VideoPlayerProps> = ({ src, className, poster }) => {
    return (
        <>
            <video className={className} src={src} controls poster={poster} />
            {/* <track src="/subt/sub.vtt" label="Spanish" kind="subtitles" srcLang="es" default/>  */}
        </>
    );
};

export default VideoPlayer;
