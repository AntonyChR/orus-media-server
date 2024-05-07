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
    const [show, setShow] = useState(true);

    const { icon, bg, textColor } = getValueProps(alertType);

    useEffect(() => {
        const interval = setTimeout(() => {
            setShow(false);
        }, timeout);
        return () => {
            clearTimeout(interval);
        };
        // eslint-disable-next-line
    }, []);
    return (
        <>
            {show && (
                <div
                    style={{
                        backgroundColor: bg,
                        bottom: bottom,
                        animationDuration: `${timeInSec}s`,
                    }}
                    className='rounded-md px-2 pt-2 items-center z-50 animate-fadeOut my-2 overflow-hidden'
                >
                    <div className='flex'>
                        <img
                            src={icon}
                            alt={alertType}
                            width={30}
                            className=''
                        />
                        <p style={{ color: textColor }} className='mx-2'>
                            {message}
                        </p>
                    </div>
                    <div
                        style={{ animationDuration: `${timeInSec}s` }}
                        className='h-[2px] bg-white w-full animate-contract left-[-0.5rem] relative rounded-md'
                    ></div>
                </div>
            )}
        </>
    );
};
