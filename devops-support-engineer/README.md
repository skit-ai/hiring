# DevOps Support Engineer Role

In this document we keep few questions for assignment rounds.

### 1. Docker
Docker is a tool to create, deploy, and run applications by using containers. It allows a developer to package up an application with all of the parts it needs, such as libraries and other dependencies thus allowing him/her to `Build once Run Anywhere`.

You have to write a [Dockerfile](https://docs.docker.com/engine/reference/builder/) with **Ubuntu 16.04** as a base image. The image should have following packages present:
- telnet
- curl
- ffmpeg

Finally, when running a container from the docker image, it should launch with `bash`.

##### Deliverables:
- A Dockerfile.
- A README file with instructions on how to build and run the image.

#### Learning Resources
- https://docs.docker.com/
- https://training.play-with-docker.com/beginner-linux/


### 2. Auditing Hardware
We work with companies that require our solution to be deployed on premise. For that we receive hardware, and auditing that hardware becomes important part of our deployment process.

Write a bash script to audit the following hardware spec in [RHEL](https://www.redhat.com/en/technologies/linux-platforms/enterprise-linux). The script should print out following specifications and also audit report should highlight specifications if they are not matching expected specifications

  - Server Uptime
  - Last Server Reboot Timestamp
  - Server Local Time Zone [Expected IST, Highlight to NON-IST ]
  - Last 10 installed packges with dates
  - OS version [Expected RHEL family, Highlight for different os]
  - Kernel version
  - CPU - Virtual cores
  - CPU - Clock speed
  - CPU - Architecture [Expected x86-64 , Highlight for x86-64]
  - Disk - Mounted/Unmounted volumes, type, storage
  - Private and Public IP
  - Private and Public DNS or Hostname
  - Networking - Bandwidth
  - Networking - OS Firewall (Allowed Ports & Protocols)
  - Networking - Network Firewall (Allowed Ports & Protocols)
  - CPU - Utilization [Expected Less than 60 %, Highlight CPU consumption]
  - RAM - Utilization [Expected Less than 60 %,Highlight RAM consumption]
  - Storage [Expected Less than 60 %, other wise Highlight Storage consumption]
  - Highlight when current User Password Exipring

##### Deliverables:
- A Shell Script.
- A README file with instructions how to run.

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

