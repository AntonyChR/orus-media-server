import { FC, useContext } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import TitleList from '../components/TitleList';
import NoData from './NoData';
interface MoviesProps {}
const Movies: FC<MoviesProps> = () => {
    const { movies } = useContext(DataContext);
    return (
        <>
            {movies.length > 0 ? (
                <TitleList titles={movies} />
            ) : (
                <NoData />
            )}
        </>
    );
};

export default Movies;
