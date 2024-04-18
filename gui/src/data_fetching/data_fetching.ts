import { FileInfo } from '../types/FileInfo';
import { TitleInfo } from '../types/TitleInfo';
import { ENDPOINTS } from './endpoints';


export const getAllTitles = async (): Promise<TitleInfo[] | null> => {
    try {
        const resp = await fetch(ENDPOINTS.media.allTitles);
        const data: TitleInfo[] = await resp.json();
        return data;
    } catch (error) {
        return null;
    }
};

export const getMovieSrc = async (titleId: number): Promise<string | null> => {
    const url = `${ENDPOINTS.media.videoFileInfo}/${titleId}`
    try {
        const resp = await fetch(url);
        const data: FileInfo[] = await resp.json();
        return `${ENDPOINTS.media.videoSrc}/${data[0].ID}`;
    } catch (error) {
        return null;
    }
};



