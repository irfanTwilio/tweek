# hello

Build and Run Docker Commands:
docker build -f Dockerfile -t hello:latest .
docker run -it -p 80:8090 --env-file .env hello:latest

Push Docker Image to Docker Hub:
docker login
docker tag go-tweek-sms:latest irfantwilio/hello
docker push irfantwilio/hello

Run app from command line:
cd /hello dir
run ngrok http http://127.0.0.1:8090 (Leave the ngrok session running)
go to Twilio console set webhook to path forwarding/sms (POST)
NOTE: each time ngrok is quit the path forwarding needs to be updated in twilio console
run go run testfinance_copy.go
send sms to TWILIO_PHONE_NUMBER

Run app in Docker Container:
docker build -f Dockerfile -t hello:latest .
docker run -it -p 80:8090 hello:latest
POSTMAN:
Method: POST
URL: http://localhost:80/sms
Body: 'form-data' Key=Body & Value=StockTicker
CURL:
curl --location --request POST 'http://localhost:80/sms' \
--form 'Body="TWLO"'



Git commands:
git add <filename>
git status
git commit -m "message"
git push --set-upstream origin main
git checkout -b main (switches to the main branch)


K8S:

hello.yaml# tweek
