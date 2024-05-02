import { FC, ReactNode, useContext, useEffect, useState } from 'react';
import { DataContext } from './context';
import { TitleInfo } from '../../types/TitleInfo';
import { getLastRoute, saveLastRoute } from '../../storage/localstorage';
import { useNavigate } from 'react-router-dom';
import { Subtitle } from '../../types/Subtitle';
import ApiDb from '../../data_fetching/data_fetching';
import { AlertContext } from '../../components/Alert/AlertContext';

interface Props {
    children: ReactNode;
}

interface StateProps {
    movies: TitleInfo[];
    series: TitleInfo[];
    subtitles: Subtitle[];
}

export const DataProvider: FC<Props> = ({ children }) => {
    const navigate = useNavigate();

    const [titles, setTitles] = useState<StateProps>({
        movies: [],
        series: [],
        subtitles: [],
    });

    const { showAlert } = useContext(AlertContext);

    const getTitles = async () => {
        const data = await ApiDb.getAllTitles();
        if (!data) {
            showAlert({
                message: 'Error fetching titles',
                alertType: 'error',
                timeout: 5000,
            });
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

        let subtitles: Subtitle[] = [];

        const subtitlesResp = await ApiDb.getAllSubtitles();
        if (subtitlesResp) {
            subtitles = subtitlesResp;
        }
        setTitles({ movies, series, subtitles });
    };

    const assignVideoIdToSubtitles = async (
        videoId: number,
        subtitleId: number
    ) => {
        const err = await ApiDb.assignVideoIdToSubtitle(videoId, subtitleId);
        if (err) {
            // show error alert
            return;
        }
        const subtitles = titles.subtitles.map((sub) => {
            if (sub.ID === subtitleId) {
                sub.VideoId = videoId;
            }
            return sub;
        });
        setTitles({ ...titles, subtitles });
    };

    useEffect(() => {
        getTitles();
        // eslint-disable-next-line
    }, []);

    useEffect(() => {
        const pathName = window.location.pathname;
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
        <DataContext.Provider value={{ ...titles, assignVideoIdToSubtitles }}>
            {children}
        </DataContext.Provider>
    );
};
