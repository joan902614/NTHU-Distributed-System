# NTHU-Distributed-System

The repository includes microservices for the NTHU Distributed System course lab. The goal of this project is to **introduce a production, realworld microservices backend mono-repo architecture** for teaching purpose.

Before going through the following parts, make sure your Docker is running since we are generating/testing/building code inside a Docker container to prevent dependencies from conflicting/missing on your host machine.

## goal
In this assignment, you are asked to monitor the system with metrics data, and learn the following skills:
* Get familiar with OpenTelemetry and Prometheus.
* Learn how to use metrics to monitor a system.
* Know how to update telemetry data with gRPC interceptor.

## Test implementation
1. Make sure  Docker is running.
2. Test implementation in the otelkit package.
```make dc.pkg.test```
3. Verify implementation of the interceptor with the following steps.
  1. build your image.
  ```make dc.image```
  2. start comment-server, comment-gateway and prometheus server.
  ```docker-compose up -d comment-api comment-gateway prometheus```
  3. Make some requests to the comment server.
  For example:
  ```curl -X GET "http://localhost:10081"```
  ```curl -X POST "http://localhost:10081/v1/comments"```
  5. open browser and type localhost:9090 to check if Prometheus collects the metrics data successfully

## test
![1](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/277900fc-7b70-433b-b904-379bdffa0842)
![2](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/e788496a-39aa-4797-b000-4e5638351938)
![3](https://github.com/joan902614/NTHU-Distributed-System/assets/132533584/27bd81dd-d04b-435d-9d9a-7867e80e863a)

## result

