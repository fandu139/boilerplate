# boilerplate
This boilerplate support Golang version 1.13.x or latest, For futher information send email: meongbego@gmail.com

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

---
## Getting Started
---
To use this boilerplate in a new microservice, you are required to change the path of the module used, cross-use the search feature in your code editor by changing from this code.
```
replace: github.com/sofyan48/boilerplate | to: github.com/orn-id/$new_service_name
```

### Environment setup
#### Development
See your docekr-compose.yml search this line orn_service_api_golang see environtment section now adding your new environment variabel

#### Production
change your .env.example to .env

### Insight and Understanding Boilerplate
**Engine** This boilerplate uses the gin framework as the basis of the engine i.e:
1. src/config
2. src/router
3. src/middleware

**Global Utils Or Package**
For the global package itself, there is at
1. src/utils

**Application** The application structure is in the ***internal/$version/***
1. *package* is a helper or local utility that is used only in that version if the function you created can be consumed by global make it in ***/src/utils/$package_name***
1. *presentation* is a local utility that is used only in that version if the function you created can be consumed by global make it in ***/src/utils/$package_name***
2. *routes* are the initiation between the router engine and the API that will be made. All routing from the API will be collected in this section
3. *ucase* is a place where you are creative in forming a microservice


### AWS Development
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
    files: comm-boilerplate-dev.json
```
### AWS Production
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