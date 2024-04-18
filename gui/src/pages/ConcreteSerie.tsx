import { useContext } from 'react';
import { useParams } from 'react-router-dom';
import { DataContext } from '../providers/dataProvider/context';
import Banner from '../components/Banner';

const ConcreteSeries = () => {
    const { seriesId } = useParams();
    const { series } = useContext(DataContext);
    let titleInfo = null;

    for (let i = 0; i < series.length; i++) {
        if (series[i].ID == Number(seriesId)) {
            titleInfo = series[i];
            break;
        }
    }
    return <div>{titleInfo && <Banner titleInfo={titleInfo} />}</div>;
};

export default ConcreteSeries;
