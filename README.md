## Medlit API Backend

How to use the API:

-   Open your browser and go to endpoint https://medlit-app.et.r.appspot.com/api/medlit/
-   You should see json that returns

    ```
    {
        "error": "false",
    	"message": "Welcome to Medlit API",
    }
    ```

-   You can use the following endpoints to test the API:
-   Endpoint: https://medlit-app.et.r.appspot.com/api/medlit/

-   GET Method

    -   `/medicine/get/all`: returns all medicines.

        ```
        {
            "error": "false",
            "medicineList": [
                {
                    "id": 1,
                    "generic_name": "Amlodipine Besilate",
                    "photo_url": "https://storage.googleapis.com/medlit-bucket/amlodipine.jpg",
                    "purpose": "Treat high blood pressure, certain types of angina and coronary artery disease",
                    "side_effects": "Swelling of the hands, feet, ankles, or lower legs, headache, upset stomach, nausea, stomach pain, dizziness or lightheadedness, drowsiness, etc",
                    "contraindication": "Hypersensitivity",
                    "dosage": "5 to 10 mg once a day for Usual Adult Dose, 2.5 to 5 mg once a day for Usual Pediatic Dose",
                    "ingredients": "Amlodipine 5 mg",
                    "created_at": "2022-06-12T15:25:49.401342+07:00",
                    "updated_at": "2022-06-12T15:25:49.401342+07:00"
                }
            ],
            "message": "Medicine list"
        }
        ```

    -   `/medicine/get/{id}`: returns a medicine with the given id.
    -   `/medicine/search?generic_name={name of medicine}`: returns all medicines with the given name.

-   POST Method
    -   `/medicine/add`: adds a new medicine.
    -   `/register`: registers a new user.
    -   `/login`: logs in a user.
