import { FC, useState } from 'react';
import { TitleInfo } from '../types/TitleInfo';

interface TitleCardProps {
    titleInfo: TitleInfo;
}

const TitleCard: FC<TitleCardProps> = ({ titleInfo }) => {
    const [isOpen, setIsOpen] = useState(false);
    return (
        <div
            className='relative'
            onMouseEnter={() => setIsOpen(true)}
            onMouseLeave={() => setIsOpen(false)}
        >
            <div>
                <div className='w-[250px] h-[350px] overflow-hidden'>
                    <img
                        height={300}
                        src={titleInfo.Poster}
                        alt={titleInfo.Title}
                    />
                </div>
                <p className='text-center text-white'>{titleInfo.Title}</p>
            </div>
            {isOpen && (
                <div className='text-white absolute top-9 right-[-20%] w-[200px] h-[350px] bg-gray-800 z-10 bg-opacity-90'>
                    <h3>
                        {titleInfo.Title} ({titleInfo.Year})
                    </h3>
                    <p>{titleInfo.Plot}</p>
                    <p>classification: {titleInfo.Rated}</p>
                    <p>IMDb: {titleInfo.imdbRating}/10</p>
                    <p>Duration: {titleInfo.Runtime}</p>
                    <p>Director: {titleInfo.Director}</p>
                </div>
            )}
        </div>
    );
};

export default TitleCard;
