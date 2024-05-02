import { FC, useEffect, useState } from 'react';

import ErrorIcon from '/error.svg';
import InfoIcon from '/info.svg';
import SuccessIcon from '/success.svg';
import warningIcon from '/warning.svg';

export type AlertTypes = 'success' | 'error' | 'warning' | 'info';

const getValueProps = (
    alertType: AlertTypes
): {
    icon: string;
    bg: string;
    textColor: string;
} => {
    const icons = {
        success: SuccessIcon,
        info: InfoIcon,
        warning: warningIcon,
        error: ErrorIcon,
    };

    const textColor = {
        success: '#000000',
        info: '#000000',
        warning: '#000000',
        error: '#FFFFFF',
    };

    const bg = {
        error: '#E23636',
        success: '#82DD55',
        info: '#4A90E2',
        warning: '#EDB95E',
    };
    return {
        icon: icons[alertType],
        bg: bg[alertType],
        textColor: textColor[alertType],
    };
};

export interface AlertProps {
    alertType: AlertTypes;
    message: string;
    timeout?: number; // default 3000
    bottom?: number;
}

export const Alert: FC<AlertProps> = ({
    alertType,
    message,
    bottom,
    timeout = 3000,
}) => {
    const timeInSec = timeout / 1000;
    const [timer, setTimer] = useState(timeInSec);
    const [show, setShow] = useState(true);

    const { icon, bg, textColor } = getValueProps(alertType);

    useEffect(() => {
        const interval = setInterval(() => {
            setTimer((prev) => prev - 1);
        }, 1000);
        const interval2 = setTimeout(() => {
            setShow(false);
        }, timeout);
        return () => {
            clearInterval(interval);
            clearTimeout(interval2);
        };
    // eslint-disable-next-line
    }, []);
    return (
        <>
            {show && (
                <div
                    style={{ backgroundColor: bg, bottom: bottom, animationDuration: `${timeInSec}s`}}
                    className="flex rounded-md p-2 items-center z-50 animate-fadeOut"

                >
                    <img src={icon} alt={alertType} width={30} className='' />
                    <p style={{ color: textColor }} className='mx-2'>
                        {message} {timer}
                    </p>
                </div>
            )}
        </>
    );
};
