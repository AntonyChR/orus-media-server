import { FC, ReactNode, useEffect, useState } from 'react';
import { DataContext } from './context';
import { TitleInfo } from '../../types/TitleInfo';
import { getAllTitles } from '../../data_fetching/data_fetching';

interface Props {
    children: ReactNode;
}

interface StateProps {
    movies: TitleInfo[];
    series: TitleInfo[];
}

export const DataProvider: FC<Props> = ({ children }) => {
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

    return (
        <DataContext.Provider value={{ ...titles }}>
            {children}
        </DataContext.Provider>
    );
};
