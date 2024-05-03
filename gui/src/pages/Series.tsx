import { FC, useContext } from 'react';
import { DataContext } from '../providers/dataProvider/context';
import TitleList from '../components/TitleList';
import NoData from './NoData';
interface SeriesProps {}

const Series: FC<SeriesProps> = () => {
    const { series } = useContext(DataContext);
    return (
        <div className=''>
            {series.length > 0 ? (
                <TitleList titles={series} />
            ) : (
                <NoData />
            )}
        </div>
    );
};

export default Series;
