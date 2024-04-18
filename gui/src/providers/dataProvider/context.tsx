import { createContext } from 'react';
import { TitleInfo } from '../../types/TitleInfo';

interface ContextProps {
    series: TitleInfo[];
    movies: TitleInfo[];
}

export const DataContext = createContext({} as ContextProps);
