import { FC, useContext } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import TitleList from '../components/TitleList';
import ReloadDatabase from '../components/ReloadDatabase';
interface SeriesProps {}

const Series: FC<SeriesProps> = () => {
    const { series } = useContext(DataContext);
    return (
        <>
            <h1>series</h1>
            {series.length > 0 ? (
                <TitleList titles={series} />
            ) : (
                <ReloadDatabase />
            )}
        </>
    );
};

export default Series;
