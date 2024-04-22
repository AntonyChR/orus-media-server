interface LastRoute {
    timeStamp: Date;
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
    const lastRoute:LastRoute = JSON.parse(data);
    return { route: lastRoute.route, timeStamp: new Date(lastRoute.timeStamp) };
};
