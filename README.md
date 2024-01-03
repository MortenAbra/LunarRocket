# Lunar - Rocket 🚀

Welcome to the Lunar - Rocket case! This readme will show you two ways of launching the project.

## Quick Start 🏁

To start the project, simply execute the following command in your terminal. This will set up everything you need, excluding the lunar rocket launch file.

```bash
chmod +x run.sh
./run.sh
```
This script will run everything from fetching docker images, running the containers to starting the go program.

### Troubleshooting 🛠️
If the prior solution doesn't work, try the following.

### Docker Compose 🐳
If you're unable to run the run.sh file, try using Docker. First, start by setting up your environment using Docker Compose:
```bash
docker compose up -d
```
This command will set up all the necessary containers in detached mode, ensuring your environment is ready.

### Go Modules 📦
Next, download the Go modules required for the project:
```bash
go mod download
```
This step ensures that all the dependencies are correctly installed and ready for use.

### Run the Application 🌟
Finally, you can start the application:

```bash
go run main.go
```
This will launch the main application.


### FYI
When launching the rockets application which sends the post requests, use the following endpoint for this application:
`http://localhost:8080/v1/rockets/messages`
