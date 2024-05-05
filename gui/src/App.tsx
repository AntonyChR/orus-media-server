import AppRouter from './AppRouter';
import AlertProvider from './components/Alert/AlertProvider';
import i18n from './i18n';
import { detectLanguage } from './i18n/detectLanguage';
import Layout from './layouts/layout';
import { DataProvider } from './providers/dataProvider/provider';

function App() {
    const lang = detectLanguage();
    i18n.changeLanguage(lang);
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
