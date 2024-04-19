interface LastRoute {
    timeStamp: any;
    route: string;
}

export const saveLastRoute = (route: string) => {
    const data: LastRoute = {
        timeStamp: new Date(),
        route: route,
    };
    localStorage.setItem('last-route', JSON.stringify(data));
};

export const getLastRoute = (): LastRoute | null => {
    const data = localStorage.getItem('last-route');
    if (!data) {
        return null;
    }
    const lastRoute = JSON.parse(data);
    return { ...lastRoute, timeStamp: new Date(lastRoute.timeStamp) };
};
