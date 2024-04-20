import { isDevMode } from '../lib/isDevMode';

const API_URI = isDevMode
    ? import.meta.env.VITE_API_URI
    : window.location.origin;

export const ENDPOINTS = {
    media: {
        allTitles: API_URI + '/api/media/titles/all',
        videoFileInfo: API_URI + '/api/media/files',
        videoSrc: API_URI + '/api/media/video',
    },
    config: {},
};
