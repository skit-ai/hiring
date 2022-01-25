# DevOps Engineer Role

In this document we keep few questions for assignment rounds.

### 1. Docker
Docker is a tool to create, deploy, and run applications by using containers. It allows a developer to package up an application with all of the parts it needs, such as libraries and other dependencies thus allowing him/her to `Build once Run Anywhere`.

`assignment/main.go` is an http server written in Golang. 

You have to write a [Dockerfile](https://docs.docker.com/engine/reference/builder/) for the same. 


Dockerfile should have:
- minimal base image requirement to run the go application
- multi-stage build
- an executable to run by default when starting the container

Deliverables:
- A Dockerfile
- A README file with instructions on how to build and run the image

#### Learning Resources
- https://docs.docker.com/
- https://training.play-with-docker.com/beginner-linux/

--- 

### 2. Deploy above application using Kubernetes

[Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/) is a portable, extensible, open-source platform for managing containerized applications. 

- Write a [kubernetes deployment](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/) object having two replicas (pods) of the above application.
- Load balance the pods using a [service](https://kubernetes.io/docs/concepts/services-networking/service/). Expose service at port 31000.

**Tip**: Push image to [dockerhub](http://dockerhub.com/) for your convenience.

Extra points if you can write a [helm chart](https://helm.sh/docs/) for the same (`values.yaml` should contain variables such as image name and version).

Deliverables:
- A file containing relevant kubernetes object
- A README file with instructions to run application on a kubernetes cluster
- For helm, submit a complete helm chart along with instructions to run on a kubernetes cluster.

##### Learning Resources
- [Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)

---
### 3. Managing Disk Space
We provide call center automation as one of our products. Essentially it processes user's audio and return a correct audio response. For managing all these audios that are generated in the due process we use [NFS](https://en.wikipedia.org/wiki/Network_File_System). With the call volumes we have, it essential that disk space of NFS should be managed properly.

Write a bash script to clean out all the `.wav ` audios data in `/data/audios/`folder that are older than `40hrs` by default (If no argrument passed) and if you pass first arguments to script with relative time like `1hrs` or `100hrs` then script should clean out audio data. Also generate a log file with the name`deleted-files-<date>-<month>-<year>.log` which should contain:
  * name of audio file
  * time of creation of audio file ([ISO Format](https://en.wikipedia.org/wiki/ISO_8601))
  * time of deletion of audio file ([ISO Format](https://en.wikipedia.org/wiki/ISO_8601))

For example:

A sample script should like `sh managing-disk-space.sh 10hrs`

A sample log file `deleted-files-12-02-2020.log` will contain:
```shell
...
audio_0010023.wav 10-02-2020T08:37:16+05:30 12-02-2020T10:18:40+05:30
audio_0010024.wav 10-02-2020T08:38:16+05:30 12-02-2020T10:18:42+05:30
audio_0010025.wav 10-02-2020T08:38:50+05:30 12-02-2020T10:18:45+05:30
...
```

##### Deliverables:
- A Shell Script.
- A Log file.
- A README file with instructions how to run.
