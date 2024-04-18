import { FC } from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import Series from './pages/Series';
import Movies from './pages/Movies';
import ConcreteSeries from './pages/ConcreteSerie';
import Config from './pages/Config';
import ConcreteMovie from './pages/ConcreteMovie';

const AppRouter: FC = () => {
    return (
        <Routes>
            <Route
                path='/'
                element={<Navigate to='/movies' replace={true} />}
            />
            <Route path='/movies' element={<Movies />} />
            <Route path='/movie/:movieId' element={<ConcreteMovie />} />

            <Route path='/series' element={<Series />} />
            <Route path='/series/:seriesId' element={<ConcreteSeries />} />
            
            <Route path='/config' element={<Config />} />
        </Routes>
    );
};

export default AppRouter;
