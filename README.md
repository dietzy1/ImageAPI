# ImageAPI

This project is an image API service that provides basic CRUD operations to accociate images with useful tags and information.
The API is hosted on Railway as a docker container, and it is the direct backend of: https://pepe-api.vercel.app/


# Architecture

The service is an interpretation of hexargonal architecture with buissness logic that is fully independant of infrastructure and http server implementationl.

Infrastructure is built with mongodb as the main database for storing associated image data and user credentials. The images themself are stored at an CDN in the cloud for easy user access. Redis is used as the secondary database for caching related tasks such as session based authentication.

Server is implemented with the golang standart library with the addition of Gorillas Mux router. Custom built middleware such as Api key authentication and user ratelimiting is also present to ensure stability of service

Inner domain logic contains extraction of regular image data but also image hashing in order to prevent multiple entries of the same image.

# Documentation
All GET endpoints require an API-key as a query parameter.
If none is provided the request will be denied.

Quantity is one of the optional query parameters. And
determines the number of images to be returned. If the
parameter is not specified then the default value of 10 is
used.

 /api/v0/image/random/
Returns a single pepe image object that is randomly selected
from the image database.

  /api/v0/image/:uuid
Returns a single pepe image object that is selected based on
UUID provided in the request.

/api/v0/images/:tags
Returns an array of pepe image objects that are selected
                    based on the tags provided in the request. Quantity can be
                    used as a query parameter.
                    
                    /api/v0/images/random/
 
  Returns an array of pepe image objects that are randomly
                    selected from the image database. Quantity can be used as a
                    query parameter.
 
