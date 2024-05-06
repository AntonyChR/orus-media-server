import { FC, useEffect, useState } from 'react';
import { ENDPOINTS } from '../data_fetching/endpoints';
import { t } from 'i18next';
interface ServerLogsProps {
    className: string;
}

interface ServerEvent {
    id: string;
    content: string;
}

const ServerLogs: FC<ServerLogsProps> = ({ className }) => {
    const [connected, setServerConnected] = useState<boolean>(false);
    const [serverEvents, setServerEvents] = useState<ServerEvent[]>([]);

    useEffect(() => {
        const eventSourceRef = new EventSource(ENDPOINTS.config.serverLogs);
        eventSourceRef.onopen = () => {
            setServerConnected(true);
        };
        eventSourceRef.onerror = () => {
            setServerConnected(false);
        };
        eventSourceRef.onmessage = (event) => {
            const newEvent: ServerEvent = JSON.parse(event.data);
            setServerEvents((prevEvents) => [newEvent,...prevEvents]);
        };
        return () => {
            eventSourceRef.close();
            setServerConnected(false);
            setServerEvents([]);
        };
    }, []);

    return (
        <div className={`${className} pr-4`}>
            <h2 className='text-white text-md flex items-center mb-3'>
                {t("Server logs")}: {t(connected ? 'Connected' : 'Disconnected')}
                <div
                    className='w-[15px] h-[15px] border inline-block rounded-full mx-2'
                    style={{
                        backgroundColor: connected ? 'greenyellow' : 'red',
                    }}
                >
                    {' '}
                </div>
            </h2>
            <div className='h-[35vh] overflow-y-auto px-2 py-3 focus:scroll-auto bg-gray-800'>
                    {serverEvents.map((event) => (
                        <li className='text-white hover:bg-gray-800 animate-flash' key={event.id}>
                            {event.content}
                        </li>
                    ))}
            </div>
        </div>
    );
};

export default ServerLogs;
