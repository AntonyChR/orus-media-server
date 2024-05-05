import { Video } from '../types/Video';
import { Subtitle } from '../types/Subtitle';
import { TitleInfo } from '../types/TitleInfo';
import { ENDPOINTS } from './endpoints';


const getAllSubtitles = async (): Promise<Subtitle[] | null> => {
    const url =ENDPOINTS.media.allSubtitles;
    try {
        const resp = await fetch(url);
        const data: Subtitle[] = await resp.json();
        return data;
    } catch (error) {
        return null;
    }
}

const assignVideoIdToSubtitle = async(videoId: number, subtitleId: number): Promise<Error|null> => {

    const url = `${ENDPOINTS.media.videoSubtitles}/${subtitleId}/${videoId}`;

    try {
        await fetch(url, { method: 'POST' });
        return null;
    }catch(error){
        if (error instanceof Error){
            return error
        }
        return new Error(`Request error: ${error}`)
    }
}

const getAllTitles = async (): Promise<TitleInfo[] | null> => {
    try {
        const resp = await fetch(ENDPOINTS.media.allTitles);
        const data: TitleInfo[] = await resp.json();
        return data;
    } catch (error) {
        return null;
    }
};

export interface VideoInfo extends Video{
    stream: string;
}

const getVideoInfo = async (titleId: number): Promise<VideoInfo| null> => {
    const url = `${ENDPOINTS.media.videoFileInfo}/${titleId}`;
    try {
        const resp = await fetch(url);
        const data: Video[] = await resp.json();
        return {stream:`${ENDPOINTS.media.videoStream}/${data[0].ID}`, ...data[0]};
    } catch (error) {
        return null;
    }
};
const getVideoChapterSrc = (videoId: string): string => {
    return `${ENDPOINTS.media.videoStream}/${videoId}`;
};

const getChapters = async (
    titleId: number
): Promise<Video[] | null> => {
    const url = `${ENDPOINTS.media.videoFileInfo}/${titleId}`;
    try {
        const resp = await fetch(url);
        const data: Video[] = await resp.json();
        return data;
    } catch (error) {
        return null;
    }
};

const resetDatabase = async (): Promise<Error | null>=> {
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

const getVideoWithNoTitleInfo = async (): Promise<Video[] | null> => {
    try {
        const resp = await fetch(ENDPOINTS.media.videoWithNoInfo);
        const data: Video[] = await resp.json();
        return data;
    } catch (error) {
        return null;
    }

}

const setApiKey = async (apiKey?: string): Promise<Error | null> => {
    if (!apiKey){
        return new Error('No api key provided')
    }
    const url = new URL(ENDPOINTS.config.setApiKey);
    url.searchParams.append('apiKey', apiKey);
    try{
        const resp = await fetch(url.toString(), { method: 'POST'});
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

const ApiDb = {
    getAllSubtitles,
    assignVideoIdToSubtitle,
    getAllTitles,
    getVideoInfo,
    getVideoChapterSrc,
    getChapters,
    resetDatabase,
    getVideoWithNoTitleInfo,
    setApiKey
}

export default ApiDb;