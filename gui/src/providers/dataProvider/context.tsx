import { createContext } from 'react';
import { TitleInfo } from '../../types/TitleInfo';
import { Subtitle } from '../../types/Subtitle';

interface ContextProps {
    series: TitleInfo[];
    movies: TitleInfo[];
    subtitles: Subtitle[]
    assignVideoIdToSubtitles: (videoId: number, subtitleId: number) => void;
}

export const DataContext = createContext({} as ContextProps);
