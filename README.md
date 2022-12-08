# ImageAPI

This project is an image API service that provides basic CRUD operations to accociate images with useful tags and information.
The API is hosted on Railway as a docker container, and it is the direct backend of: https://pepe-api.vercel.app/


# Architecture

The service is an interpretation of hexargonal architecture with buissness logic that is fully independant of infrastructure and http server implementationl.

Infrastructure is built with mongodb as the main database for storing associated image data and user credentials. The images themself are stored at an CDN in the cloud for easy user access. Redis is used as the secondary database for caching related tasks such as session based authentication.

Server is implemented with the golang standart library with the addition of Gorillas Mux router. Custom built middleware such as Api key authentication and user ratelimiting is also present to ensure stability of service

# Documentation
