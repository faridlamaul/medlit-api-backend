## Medlit API Backend

How to use the API:

-   Make sure you have installed docker and docker-compose on your local computer. If you don't, please refer to the [Docker installation guide](https://docs.docker.com/install/).
-   Run `docker-compose up` to start the API.

    ```bash
    docker-compose up -d
    ```

-   Open your browser and go to http://localhost:8080/api/v1/
-   You should see json that returns `Welcome to Medlit API Backend`
-   You can use the following endpoints to test the API:

    -   GET Method

        -   `/api/v1/`: returns the list of available endpoints.
        -   `/api/v1/medicine/get/all`: returns all medicines.
        -   `/api/v1/medicine/get/{id}`: returns a medicine with the given id.
        -   `/api/v1/medicine/search?generic_name={name of medicine}`: returns all medicines with the given name.

    -   POST Method
        -   `/api/v1/medicine/add`: adds a new medicine.
        -   `/api/v1/register`: registers a new user.
        -   `/api/v1/login`: logs in a user.
