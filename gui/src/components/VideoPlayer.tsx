import { FC } from 'react';
interface VideoPlayerProps {
    src: string
    className: string
}

const VideoPlayer:FC<VideoPlayerProps> = ({src, className}) => {
    return (
        <video className={className} src={src} controls/>
    );
};

export default VideoPlayer;