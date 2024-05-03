import { FC, useEffect, useState } from 'react';
import { ENDPOINTS } from '../data_fetching/endpoints';
interface ServerLogsProps {
    className: string;
}

interface ServerEvent {
    id: string;
    content: string;
}

const ServerLogs: FC<ServerLogsProps> = ({ className }) => {
    const [connected, setServerConnected] = useState<boolean>(false);
    const [serverEvents, setServerEvents] = useState<
        { id: string; content: string }[]
    >([]);
    //const eventSourceRef = useRef<EventSource | null>(null);

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
        <div className={`${className}`}>
            <h2 className='text-white text-xl'>
                Server logs: {connected ? 'CONNECTED' : 'DISCONNECTED'}{' '}
                <div
                    className='w-[15px] h-[15px] border inline-block rounded-full'
                    style={{
                        backgroundColor: connected ? 'greenyellow' : 'red',
                    }}
                >
                    {' '}
                </div>
            </h2>
            <div className='border max-h-[35vh] overflow-y-auto px-2 py-3 focus:scroll-auto flex flex-col-reverse'>
                    {serverEvents.map((event, index) => (
                        <li className='text-white hover:bg-gray-800' key={event.id}>
                            {index}: {event.content}
                        </li>
                    ))}
            </div>
        </div>
    );
};

export default ServerLogs;
