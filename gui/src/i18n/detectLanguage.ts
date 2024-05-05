export const detectLanguage = () => {
    const lang = navigator.language;
    if (lang.split("-")[0] === 'es') {
        return 'es';
    }
    return 'en';
}