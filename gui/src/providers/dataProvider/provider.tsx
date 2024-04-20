import { FC, ReactNode, useEffect, useState } from 'react';
import { DataContext } from './context';
import { TitleInfo } from '../../types/TitleInfo';
import { getAllTitles } from '../../data_fetching/data_fetching';
import { getLastRoute, saveLastRoute } from '../../storage/localstorage';
import { useNavigate } from 'react-router-dom';

interface Props {
    children: ReactNode;
}

interface StateProps {
    movies: TitleInfo[];
    series: TitleInfo[];
}

export const DataProvider: FC<Props> = ({ children }) => {
    const navigate = useNavigate();

    const [titles, setTitles] = useState<StateProps>({
        movies: [],
        series: [],
    });

    const getTitles = async () => {
        const data = await getAllTitles();
        if (!data) {
            // show error alert
            return;
        }

        const movies: TitleInfo[] = [];
        const series: TitleInfo[] = [];
        data.forEach((title) => {
            if (title.Title.trim() != '') {
                if (title.Type == 'movie') {
                    movies.push(title);
                } else {
                    series.push(title);
                }
            }
        });

        setTitles({ movies, series });
    };
    useEffect(() => {
        getTitles();
    }, []);

    useEffect(() => {
        const pathName = window.location.pathname
        if (pathName != '/movies' && pathName != '/') {
            saveLastRoute(pathName);
        }
    }, [window.location.pathname]);

    useEffect(() => {
        const lastRoute = getLastRoute();
        if (!lastRoute) return;

        const now: any = new Date();
        if (now - lastRoute.timeStamp <= 15_000) {
            navigate(lastRoute.route);
        }
    }, []);

    return (
        <DataContext.Provider value={{ ...titles }}>
            {children}
        </DataContext.Provider>
    );
};
