import AppRouter from './AppRouter';
import AlertProvider from './components/Alert/AlertProvider';
import Layout from './layouts/layout';
import { DataProvider } from './providers/dataProvider/provider';

function App() {
    return (
        <Layout>
            <AlertProvider>
                <DataProvider>
                    <AppRouter />
                </DataProvider>
            </AlertProvider>
        </Layout>
    );
}

export default App;
