import { FC } from 'react';
import { TitleInfo } from '../types/TitleInfo';
import { Link } from 'react-router-dom';
import TitleCard from './TitleCard';

interface TitleListProps {
    titles: TitleInfo[];
}

const TitleList: FC<TitleListProps> = ({ titles }) => {
    return (
        <div>
            <ul className='grid sm:grid-cols-1 md:grid-cols-3 xl:grid-cols-4 justify-items-center gap-y-8'>
                {titles.map((title) => (
                    <li key={title.imdbID}>
                        <Link
                            to={`/${title.Type}/${title.ID}/${
                                title.Type == 'series' ? 0 : ''
                            }`}
                        >
                            <TitleCard titleInfo={title} />
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TitleList;
