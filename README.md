# ORN-MN-BOILERPLATE
This boilerplate support Golang version 1.13.x or latest, For futher information, send a email:

1. meongbego@gmail.com

## Foreword
---
### Install Golang
**Darwin**
```
wget https://dl.google.com/go/go1.14.3.darwin-amd64.pkg
```
double click to install

**Linux**
```
curl -fLo /usr/local https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
```
export path and write to .bashrc
```
export PATH=$PATH:/usr/local/go/bin
```

**Windows**
Download [Golang](https://dl.google.com/go/go1.14.3.windows-amd64.msi) double click binary to install

### How To Contribute
---

Please refer to each project's style and contribution guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.
 1. ***Fork*** the repo on GitHub
 2. ***Clone*** the project to your own machine
 3. ***Commit*** changes to your own branch
 4. ***Push*** your work back up to your fork
 5. Submit a ***Pull request*** so that we can review your changes


# SERVICE
---
## Getting Started
---
To use this boilerplate in a new microservice, you are required to change the path of the module used, cross-use the search feature in your code editor by changing from this code.
```
replace: github.com/orn/rl-ms-boilerplate-go | to: github.com/orn/$new_service_name
```

### Insight and Understanding Boilerplate
**Engine** This boilerplate uses the gin framework as the basis of the engine i.e:
1. src/config
2. src/router
3. src/middleware

**Global Utils Or Package**
For the global package itself, there is at
1. src/utils

**Application** The application structure is in the ***app/$version/***
1. *package* is a helper or local utility that is used only in that version if the function you created can be consumed by global make it in ***/src/utils/$package_name***
2. *routes* are the initiation between the router engine and the API that will be made. All routing from the API will be collected in this section
3. *api* is a place where you are creative in forming a microservice

The structure used in this boilerplate is the structure that is applied to the ***Clean Architercture from Uncle Bob*** model, to see more clearly the structure please read [here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

### Environment setup
---

```
user$ cp .env.example .env
```
see this value for example
```
####################################################################
# SERVICE CONFIGURATION
####################################################################
SERVER_ADDRESS=0.0.0.0
SERVER_PORT=3000
SERVER_TIMEZONE=Asia/Jakarta
SECRET_KEY=xXbxnad!!sadsa
APP_NAME=order

####################################################################
# DATABASE CONFIGURATION
####################################################################
######################## WRITE DATABASE #############################
DB_HOST=localhost
#DB_HOST=orn_order_api_mysql
DB_USER=root
DB_PASSWORD=orn-order_password
DB_PORT=3306
DB_NAME=orn-order
######################## READ DATABASE #############################
DB_USER_READ=orn-order_username
DB_PASSWORD_READ=orn-order_password
DB_HOST_READ=orn_order_api_mysql
DB_PORT_READ=3306
DB_NAME_READ=orn-order
```

### Local Development
---

After installing the Golang, the next step is to install all the needs to run the program, one of them is a database, if you want to take advantage of the needs that we have prepared, you can combine it with the docker as needed.

Examples to take advantage of the needs we have prepared *(can be seen in the list of services available in the compose docker)* according to your needs are as follows:
list of service container
```
user$ docker-compose config --service

orn_order_api_mysql
orn_ecatalog_api_golang
```
run service container
```
docker-compose start $service_container_in_docker_compose
```
After install all requirements now run application from application root directory
```
user$ go run src/main.go
```
#### Auto Refresh
To activate Live Reload install air

on macOS
```
user$ curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air
user$ chmod +x /usr/local/bin/air
```
on Linux
```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
chmod +x /usr/local/bin/air
```
on Windows
```
user$ curl -fLo ~/air.exe \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/windows/air.exe
```
see *.air.conf* setting watcher file for air config now then starting application from air
```
user$ air
```
or dont use default configuration
```
user$ air -c config.conf
```

### Dockerizing Development
---

To use the docker as a development model, make sure that the specifications of the computer you are using are able to run all the containers that are made, for more details, please read [here](https://docs.docker.com/datacenter/ucp/1.1/installation/system-requirements/).

To run dockerizing development, please execute this command:
```
user$ docker-compose up
```
All the needs regarding development are available, you just need to make changes to the code. Please wait for a few moments until all service containers are ready to be used.

*note: If there is a database connection problem you can simply change some code to run the server again*


### Production
---

Production section of this boilerplate provides AWS support spesificly for CodeBuild and CodePipeline service found in the buildspec folder. Call devops or sysops for your access or email [mailto](mailto:devops@orn.com).

To save the storage capacity of the Dockerfile each environment uses a 2 stage model. 

#### AWS Development
setting runtime version to golang 1.13 or latest and for environtment setup dont forget setup AWS System manager to storing environment value.
```
version: 0.2
phases:
  install:
    runtime-versions:
      golang: 1.13
  pre_build:
    commands:
      - COMMIT_HASH=$(echo $CODEBUILD_RESOLVED_SOURCE_VERSION | cut -c 1-7)
      - IMAGE_TAG=${COMMIT_HASH:=latest}
  build:
    commands:
      - docker build -t $REPOSITORY_URI:latest -f dockerfiles/dockerfile-comm-dev
      - docker tag $REPOSITORY_URI:latest $REPOSITORY_URI:$IMAGE_TAG
  post_build:
    commands:
      - docker push $REPOSITORY_URI:$IMAGE_TAG
artifacts:
    files: comm-ecatalog-dev.json
```
#### AWS Production
setting runtime version to golang 1.13 or latest and for environtment setup dont forget setup AWS System manager to storing environment value.
```
version: 0.2
phases:
  install:
    runtime-versions:
      golang: 1.13
  pre_build:
    commands:
      - COMMIT_HASH=$(echo $CODEBUILD_RESOLVED_SOURCE_VERSION | cut -c 1-7)
      - IMAGE_TAG=${COMMIT_HASH:=latest}
  build:
    commands:
      - docker build -t $REPOSITORY_URI:latest -f dockerfiles/dockerfile-comm-dev
      - docker tag $REPOSITORY_URI:latest $REPOSITORY_URI:$IMAGE_TAG
  post_build:
    commands:
      - docker push $REPOSITORY_URI:$IMAGE_TAG
artifacts:
    files: comm-svc-prd.json
```