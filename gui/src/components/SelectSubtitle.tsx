import { FC, useContext, useState } from 'react';
import { DataContext } from '../providers/dataProvider/context';
interface SelectSubtitleProps {
    onSelect: (subtitleId: number) => void;
}

const SelectSubtitle: FC<SelectSubtitleProps> = ({onSelect}) => {
    const { subtitles } = useContext(DataContext);
    const [selectedSubtitle, setSelectedSubtitle] = useState<number>(0);
    const handleChange:React.ChangeEventHandler<HTMLSelectElement> = (e) =>{
        setSelectedSubtitle(Number(e.target.value))
        onSelect(Number(e.target.value))
    }

    return (
        <>
            {subtitles.length === 0 ? (
                <p>No subtitles available</p>
            ) : (
                <select onChange={handleChange} value={selectedSubtitle}>
                    {subtitles.map((subtitle) => {
                        return (
                            <option key={subtitle.Name} value={subtitle.ID}>
                                {subtitle.Name}
                            </option>
                        );
                    })}
                </select>
            )}
        </>
    );
};

export default SelectSubtitle;
