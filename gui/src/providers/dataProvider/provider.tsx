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
    // eslint-disable-next-line
    }, [window.location.pathname]);

    useEffect(() => {
        const lastRoute = getLastRoute();
        if (!lastRoute) return;

        const now = new Date();
        if (Number(now) - Number(lastRoute.timeStamp) <= 15_000) {
            navigate(lastRoute.route);
        }
    // eslint-disable-next-line
    }, []);

    return (
        <DataContext.Provider value={{ ...titles }}>
            {children}
        </DataContext.Provider>
    );
};
