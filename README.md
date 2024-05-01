## Prepare Development Environment

To set up your development environment for the Orus Media Server project, follow these steps:

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/AntonyChR/orus-media-server.git
    ```

2. Change into the project directory:

    ```bash
    cd orus-media-server
    ```

3. Install the required dependencies:

    ```bash
    make prepare
    ```

    This command will configure custom git hooks and set up the necessary dependencies for the project.

4. Start the development server:

    ```bash
    go run main.go
    cd gui && npm run dev
    ```

    This command will start the Orus Media Server in development mode.

That's it! You're now ready to start developing with the Orus Media Server project. Happy coding!