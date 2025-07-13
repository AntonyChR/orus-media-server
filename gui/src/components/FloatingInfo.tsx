import { FC } from 'react';
import { TitleInfo } from '../types/TitleInfo';
import star from '/star.svg';
import { t } from 'i18next';

interface FloatingInfoProps {
    titleInfo: TitleInfo;
}

const FloatingInfo: FC<FloatingInfoProps> = ({ titleInfo }) => {
    return (
        <div
            className='text-white absolute top-2 right-[-20%] w-[300px] h-max z-10 p-3 space-y-2 rounded-md bg-gray-950 border border-gray-900 animate-fadeIn'
        >
            <p>{titleInfo.Plot}</p>
            <p>
                <span className='font-bold'>{t("Duration")}</span>: {titleInfo.Runtime}
            </p>
            <p>
                <span className='font-bold'>{t("Director")}</span>:{' '}
                {titleInfo.Director}
            </p>
            <p>
                <span className='font-bold'>{t("Classification")}</span>:{' '}
                {titleInfo.Rated}
            </p>
            <p>
                <span className='font-bold'>IMDb: </span>
                {titleInfo.imdbRating}/10{' '}
                <img className='inline mb-2' width={20} src={star} />
            </p>
        </div>
    );
};

export default FloatingInfo;
