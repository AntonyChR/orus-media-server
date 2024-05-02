import { FC, ReactNode, useState } from 'react';
import { v4 as uuidv4 } from 'uuid';

import { AlertContext } from './AlertContext';
import { Alert, AlertProps } from './Alert';
interface AlertProviderProps {
    children: ReactNode;
}

interface ExtendedAlertProps extends AlertProps {
    id: string;
}

const AlertProvider: FC<AlertProviderProps> = ({ children }) => {
    const [alerts, setAlerts] = useState<ExtendedAlertProps[]>([]);

    const showAlert = (props: AlertProps) => {
        const timeout = props.timeout || 3000;
        const bottom = (alerts.length + 1) * 60;
        const id = uuidv4();
        setAlerts((prevAlerts) => [
            ...prevAlerts,
            { ...props, bottom, id, timeout: timeout },
        ]);

        setTimeout(() => {
            setAlerts((prevAlerts) =>{
                if (prevAlerts.length === 1) {
                    return [];
                }
                return prevAlerts.filter((alert) => alert.id !== id)
            }
            );
        }, timeout);


    };

    return (
        <AlertContext.Provider value={{ showAlert }}>
            {children}
            <div className='fixed bottom-2'>
                {alerts.map((alert, index) => (
                    <Alert key={alert.id} {...alert} bottom={index * 60} />
                ))}
            </div>
        </AlertContext.Provider>
    );
};
export default AlertProvider;
