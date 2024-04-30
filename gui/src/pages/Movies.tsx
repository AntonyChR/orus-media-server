import { FC, useContext } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import TitleList from '../components/TitleList';
import ResetDatabase from '../components/ResetDatabase';
interface MoviesProps {}
const Movies: FC<MoviesProps> = () => {
    const { movies } = useContext(DataContext);
    return (
        <>
            <h1>Movies</h1>
            {movies.length > 0 ? (
                <TitleList titles={movies} />
            ) : (
                <ResetDatabase />
            )}
        </>
    );
};

export default Movies;
