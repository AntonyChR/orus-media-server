import AppRouter from './AppRouter';
import Layout from './layouts/layout';
import { DataProvider } from './providers/dataProvider/provider';

function App() {
    return (
        <Layout>
            <DataProvider>
                <AppRouter />
            </DataProvider>
        </Layout>
    );
}

export default App;
