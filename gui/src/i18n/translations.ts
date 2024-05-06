import { Resource } from 'i18next';

export const resources:Resource = {
    en: {
      translation: {
        // Navbar
        'Movies': 'Movies',
        'Series': 'Series',
        'Config': 'Settings',

        // Config
        'Save': 'Save',
        'Files with no info':'Media files without information',
        'Here videos with no info':'Here you can see the list of videos that have no information',
        'Info provider':'Info provider',
        'Select info provider':'Select the info provider you want to use to get information about the videos',
        'Get api key': 'You can get an API key from:',
        'After adding api key': 'After adding the API key and there is no information in the database, click on the "Reset database"',
        'Invalid api key':'Invalid api key',
        'Server logs':'Server logs',
        'Connected':'Connected',
        'Disconnected':'Disconnected',
        'Reset database':'Reset database',

        'Directory structure':'Directory structure',
        'The directory structure should be like:':'The directory structure containing the media files must be as follows:',

        'Name format':'The name format of the media files must be as follows:',
        'Movie name format':'Name of a movie file: "movie name (year).ext"',
        'Serie name format':'Series folder name format: "serie name (year)/sXeY.ext"',
        'Example directory':'Example directory',

        // Error page
        'No info available':'No information available, check',

        // floating info card
        'Duration': 'Duration',
        'Director': 'Director',
        'classification': 'Classification',

        // alert messages
        'Are you sure?':'Are you sure?',

        //subtitles
        'No subtitles available': 'No subtitles available',
        '-- subtitles --': '-- subtitles --'
      }
    },
    es: {
      translation: {
        // Navbar
        'Movies': 'Películas',
        'Series': 'Series',
        'Config': 'Configuración',

        // Config
        'Save': 'Guardar',
        'Files with no info':'Archivos multimedia sin información',
        'Here videos with no info':'Aquí puedes ver la lista de videos que no tienen información',
        'Info provider':'Proveedor de información',
        'Select info provider':'Selecciona el proveedor de información que deseas utilizar para obtener información sobre los videos',
        'Get api key': 'Puedes obtener una api-key en:',
        'After adding api key': 'Si después de agregar la api-key no hay información en la base de datos, haz clic en "Restablecer base de datos"',
        'Invalid api key':'Api-key inválida',
        'Server logs':'Eventos del servidor',
        'Connected':'Conectado',
        'Disconnected':'Desconectado',
        'Reset database':'Restablecer base de datos',

        'Directory structure':'Estructura de directorios',
        'The directory structure should be like:':'La estructura de directorios que contiene los archivos multimedia debe ser la siguiente:',
        'Name format':'El formato de nombre de los archivos multimedia debe ser el siguiente:',
        'Movie name format':'Nombre de una película: "nombre de la pelicula (año).ext"',
        'Serie name format':'Nombre de la carpeta de una serie: "nombre de la serie (año)/sXeY.ext"',
        'Example directory':'Directorio de ejemplo',

        // Error page
        'No info available':'No hay información disponible, verifique la',

        // floating info card
        'Duration': 'Duración',
        'Director': 'Director',
        'Classification': 'Clasificación',

        // alert messages
        'Are you sure?':'¿Estás seguro?',

        //subtitles
        'No subtitles available': 'No hay subtítulos disponibles',
        '-- subtitles --': '-- subtítulos --'
      }
    }
  };