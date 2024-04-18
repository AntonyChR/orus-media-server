import { FC, useContext } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import TitleList from '../components/TitleList';
import ReloadDatabase from '../components/ReloadDatabase';
interface MoviesProps {}
const Movies: FC<MoviesProps> = () => {
    const { movies } = useContext(DataContext);
    return (
        <>
            <h1>Movies</h1>
            {movies.length > 0 ? (
                <TitleList titles={movies} />
            ) : (
                <ReloadDatabase />
            )}
        </>
    );
};

export default Movies;
