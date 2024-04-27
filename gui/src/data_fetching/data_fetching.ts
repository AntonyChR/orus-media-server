import { FileInfo } from '../types/FileInfo';
import { Subtitle } from '../types/Subtitle';
import { TitleInfo } from '../types/TitleInfo';
import { ENDPOINTS } from './endpoints';


export const getAllSubtitles = async (): Promise<Subtitle[] | null> => {
    const url =ENDPOINTS.media.allSubtitles;
    try {
        const resp = await fetch(url);
        const data: Subtitle[] = await resp.json();
        return data;
    } catch (error) {
        return null;
    }
}

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
    const url = `${ENDPOINTS.media.videoFileInfo}/${titleId}`;
    try {
        const resp = await fetch(url);
        const data: FileInfo[] = await resp.json();
        return `${ENDPOINTS.media.videoStream}/${data[0].ID}`;
    } catch (error) {
        return null;
    }
};
export const getVideoChapterSrc = (videoId: string): string => {
    return `${ENDPOINTS.media.videoStream}/${videoId}`;
};

export const getChapters = async (
    titleId: number
): Promise<FileInfo[] | null> => {
    const url = `${ENDPOINTS.media.videoFileInfo}/${titleId}`;
    try {
        const resp = await fetch(url);
        const data: FileInfo[] = await resp.json();
        return data;
    } catch (error) {
        return null;
    }
};

export const resetDatabase = async (): Promise<Error | null>=> {
    try{
        const resp = await fetch(ENDPOINTS.config.resetDb)
        if (resp.status != 200){
            return new Error(`Bad request: ${resp.status} - ${resp.statusText}`)
        }
        return null
    }catch(error){
        if (error instanceof Error){
            return error
        }
        return new Error(`Request error: ${error}`)
    }

}