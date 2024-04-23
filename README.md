# NTHU-Distributed-System

The repository includes microservices for the NTHU Distributed System course lab. The goal of this project is to **introduce a production, realworld microservices backend mono-repo architecture** for teaching purpose.

Before going through the following parts, make sure your Docker is running since we are generating/testing/building code inside a Docker container to prevent dependencies from conflicting/missing on your host machine.

## goal
In this assignment, implement a message queue by using Kafka to support communication between the video server (based on the previous lab) and the stream server (the new one).
* Learn to use Apache Kafka by Sarama, which is go client library for Apache Kafka

## Problem Description
In this project, be going to mock video transcoding by a new server, stream server. Thus, when the video uploading process is handled by the video server done successfully, need to send a relative message to Kafka. 
When the stream server consumes the message from Kafka, it knows the video is the original video without target resolution to transcode by a specific field in the message. Then, the stream server produces four messages with different resolutions [360, 480, 720, 1080] to Kafka.
Now, if the stream server consumes the message with target resolution, it mocks video transcoding by waiting for 3 seconds.

The detail progress is as below: 
1. In the video server, the video is uploaded successfully and writes the information into the database successfully.
2. The video server sends a message to Kafka server with the field “Scale” = 0.
3. The stream server consumes a message with “Scale” = 0, produces another four messages with the same video id, url but fills the “Scale” field with four resolutions. [360, 480, 720, 1080]
4. The stream server consumes a message with “Scale” not equal to 0, mocks the video transcoding process by sleeping for 3 seconds.
5. After the stream server done the video transcoding process, update database variants part. Since we just mock the transcoding progress, we just put the same URL into it.

What have to do in this project:
1. Understand the logic in the 3 Code Structure part.
2. Finish the code in modules/video/service/service.go, handle after the video is uploaded.
3. Finish the code in modules/video/stream/stream.go, handle after receiving a message from Kafka.
4. [HINT] can search all “Kafka TODO” comments in the code to complete this lab!
## Code Structure
* cmd/video/api.go
  * Where to start up all the services in the video server.
* cmd/video/stream.go
  * Where to start up all the services in the stream server.
* pkg/kafkakit/producer.go
  * Kafka producer startup, send messages operation.
* pkg/kafkakit/consumer.go
  * Kafka consumer group startup, consume messages operation.
* modules/video/pb/stream.pb
  * Message content definition.
* modules/video/pb/stream.pb.sarama.go
  * Kafka consumer group handler implementation
* modules/video/service/service.go
  * Most of all are as same as the previous lab.
  * need to implement the logic after a video is uploaded.
* modules/video/stream/stream.go
  * need to implement the logic after a message receives from Kafka 
* [NOTE] only need to modify where [Kafka TODO]
* [NOTE] All “*_test.go”  files and files in “mock” directory are for unit testing. They are not related to backend operations.
## Prerequisite
1. Docker
2. Docker compose
3. Fork this project to your own github repo, checkout to branch “Lab-Kafka” to complete this lab.
4. [NOTE] If you already have forked this project, you can just fetch the new branch “Lab-Kafka”. 
Reference of how to fetch: [link](https://stackoverflow.com/questions/7244321/how-do-i-update-or-sync-a-forked-repository-on-github)
## Test implementation
Unit testing 
1. Make sure your Docker is running.
2. Generate gRPC client and server interface from .proto file:  
```make dc.generate```    
Or if only want to generate codes in certain module:    
```make dc.{module}.generate```  
For example:  
```make dc.video.generate```
3. Test implementation. 
```make dc.test```  
Or if only want to test certain module:  
```make dc.{module}.test```  
For example:  
```make dc.video.test```  

Test like the example video in ppt
1. Make sure your Docker is running, and set the memory space to 10GB
2. Generate docker image
```make dc.image```
3. Start up Kafka server, it may take a while  
```docker compose up -d kafka```  
You can test whether kafka runs successfully, if you do not get an error after entering the below command, kafka is running. This command try to have Kafka list all topics  
```docker exec -it nthu-distributed-system-kafka-1 kafka-topics --bootstrap-server kafka:9092 --list```
4. Start up video-api, video-gateway server
```docker compose up -d video-api video-gateway```
5. Start up video-stream server  
```docker compose up -d video-stream```  
6. Send request to the video server  
```curl -X POST localhost:10080/v1/videos -F "file=@/{Your path}/NTHU-Distributed-System/modules/video/service/fixtures/big_buck_bunny_240p_1mb.mp4"```
7. Check database  
enter mongo CLI by below commands  
```docker exec -it nthu-distributed-system-mongo-1 bash```  
Then, enter below commands to get records in the database and the videos collection  
```mongo```  
```use nthu_distributed_system```  
```db.videos.find({})```  
8. Result
should see a record with “Variants” has 4 different values and map to 4 URLs. Like the image below.
![image](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/ccd600e4-a8ef-4d65-89eb-3ea7c8747d88)

## result
test:  
![image](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/ba104c6e-b96e-480f-9791-8363b60dc8bd)
![image](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/bceb1626-a322-4e2c-9382-1500bbe8a42c)
![image](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/99d047c2-4c3e-4f3d-9709-144abb1a641a)
![image](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/72eb53db-6fd0-4403-b3e5-f6f065dc9a43)

run:
![image](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/618c9a19-cacd-4c08-91ec-fcb6b2002820)


