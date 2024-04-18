//const API_URI = new URL(window.location.href).host
const API_URI = "http://localhost:3002"

export const ENDPOINTS = {
    media :{
        allTitles: API_URI+'/api/info/titles/all',
        videoFileInfo: API_URI+'/api/info/files/title',
        videoSrc: API_URI+'/api/info/video'
    },
    config :{}
}