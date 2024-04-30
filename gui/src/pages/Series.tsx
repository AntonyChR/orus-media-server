import { FC, useContext } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import TitleList from '../components/TitleList';
import ResetDatabase from '../components/ResetDatabase';
interface SeriesProps {}

const Series: FC<SeriesProps> = () => {
    const { series } = useContext(DataContext);
    return (
        <>
            <h1>series</h1>
            {series.length > 0 ? (
                <TitleList titles={series} />
            ) : (
                <ResetDatabase />
            )}
        </>
    );
};

export default Series;
