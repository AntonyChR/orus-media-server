import { FC, useState } from 'react';
import { TitleInfo } from '../types/TitleInfo';
import FloatingInfo from './FloatingInfo';

interface TitleCardProps {
    titleInfo: TitleInfo;
}

const TitleCard: FC<TitleCardProps> = ({ titleInfo }) => {
    const [isOpen, setIsOpen] = useState(false);
    return (
        <div
            className='relative w-[250px]'
            onMouseEnter={() => setIsOpen(true)}
            onMouseLeave={() => setIsOpen(false)}
        >
            <div>
                <div className='w-[250px] h-[350px] overflow-hidden rounded-md'>
                    <img
                        height={300}
                        src={titleInfo.Poster}
                        alt={titleInfo.Title}
                    />
                </div>
                <p className='text-center text-white'>{titleInfo.Title}</p>
            </div>
            {isOpen && <FloatingInfo titleInfo={titleInfo}/>}
        </div>
    );
};

export default TitleCard;
