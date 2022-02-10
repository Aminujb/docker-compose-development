### Daily Standup App Docker Environment Setup
1. Ensure that [docker](https://www.docker.com/) is installed.
1. Ensure that [docker-compose](https://docs.docker.com/compose/) is also installed
1. Run `docker-compose up` from the project directory. This will build the configured images and create the configured containers then start up the containers.

To rebuild images run `docker-compose up --build`

Once all containers are running successfully
Open [http://localhost:3051](http://localhost:3051) to view the frontend in your browser.
