import { FC, useContext, useState } from 'react';
import { DataContext } from '../providers/dataProvider/context';
interface SelectSubtitleProps {
    onSelect: (subtitleId: number) => void;
    className?: string;
}

const SelectSubtitle: FC<SelectSubtitleProps> = ({ onSelect, className }) => {
    const { subtitles } = useContext(DataContext);
    const [selectedSubtitle, setSelectedSubtitle] = useState<number>(0);
    const handleChange: React.ChangeEventHandler<HTMLSelectElement> = (e) => {
        setSelectedSubtitle(Number(e.target.value));
        onSelect(Number(e.target.value));
    };

    return (
        <div className={className}>
            {subtitles.length === 0 ? (
                <p className=''>No subtitles available</p>
            ) : (
                <>
                    <label htmlFor='subtitle' className='text-white'>
                        Available subtitles:
                    </label>
                    <select
                    className='rounded-md ml-3 bg-gray-700 text-white'
                        onChange={handleChange}
                        value={selectedSubtitle}
                        id='subtitle'
                    >
                        {subtitles.map((subtitle) => {
                            return (
                                <option key={subtitle.Name} value={subtitle.ID}>
                                    {subtitle.Name}
                                </option>
                            );
                        })}
                    </select>
                </>
            )}
        </div>
    );
};

export default SelectSubtitle;
