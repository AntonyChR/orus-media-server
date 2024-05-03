import { isDevMode } from '../lib/isDevMode';

const API_URI = isDevMode
    ? import.meta.env.VITE_API_URI
    : `${window.location.origin}/api`;

export const ENDPOINTS = {
    media: {
        allTitles: API_URI + '/media/titles/all',
        videoFileInfo: API_URI + '/media/video',
        videoStream: API_URI + '/media/stream',
        allSubtitles: API_URI + '/media/all-subtitles',
        videoSubtitles: API_URI + '/media/video-subtitles',
        subtitlesServer: API_URI + '/media/subtitles',
        videoWithNoInfo: API_URI + '/media/no-info',
    },
    config: {
        resetDb: API_URI + '/manage/reset',
        setApiKey: API_URI + '/manage/api-key',
        serverLogs: API_URI + '/manage/events'
    },
};
