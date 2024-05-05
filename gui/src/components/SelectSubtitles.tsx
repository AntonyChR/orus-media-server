import { FC, useContext } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import { t } from 'i18next';

interface SelectSubtitlesProps {
    videoId: number;
    className?: string;
}

const SelectSubtitles: FC<SelectSubtitlesProps> = ({ className, videoId }) => {
    const { subtitles, assignVideoIdToSubtitles } = useContext(DataContext);

    const handleChange: React.ChangeEventHandler<HTMLSelectElement> = (e) => {
        const subId = Number(e.target.value);
        assignVideoIdToSubtitles(videoId, subId);
    };

    return (
        <div className={className}>
            {subtitles.length === 0 ? (
                <p className=''>{t('No subtitles available')}</p>
            ) : (
                <select
                    className='rounded-md ml-3 bg-gray-700 text-white'
                    onChange={handleChange}
                    //value={selectedSubId}
                    defaultValue={0}
                    id='subtitle'
                >
                    <option value={0} disabled>
                        {t('-- subtitles --')}
                    </option>
                    {subtitles.map((subtitle) => {
                        return (
                            <option key={subtitle.Name} value={subtitle.ID}>
                                {subtitle.Name}
                            </option>
                        );
                    })}
                </select>
            )}
        </div>
    );
};

export default SelectSubtitles;
