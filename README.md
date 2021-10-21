# Web application on Golang

## About

This application was created as part of the DevOps exam. Application is very simple and replies only with "Hello World 2" message. 

More interesting thing is a Github flow workflow created for deployment of this app.

## About Github flow

Github flow is the simple trunk-based workflow which is consists of one main branch and several feature branches. Every developer is workin in his own feature branch. Developer commiting in his own branch with new versions of feature with reasonable commit messages. Commit messages allows to check development history and roll back the changes if critical issue was found. 

After feature is complete developer creating a pull request for merging his branch into main. After the pull request is opened code reviewing and conversations are started. Developer is in process of updating and improvement of his code. If everyone is desided that code is complete, then pull request is applied and feature branch is deleted. 

Code is in CD stage now. It will be delivered to the test environment, then to the staging and production. Very simple development process. 

## My implementation of this process

My implementation is based on Github Actions and its workflows. Each workflow can be described by the .yml/.yaml file and can be triggered by many events: pull requests, pushes, merges, issues, etc. 

My workflow works like this:

* At first we need a running infrastructure. Infrustructure creation can be triggered by [Run Infrustrictire](https://github.com/RainbowGravity/golang-app/actions/workflows/run_infrastructure.yml) Github workflow. After process is complete we are ready to go.
* Then we need to create a pull request with the new code version. After this automated code tests will run. In Golang case: dockerfilelint, golangci-lint, Snyk test, Snyk code test. Results of the code checks will be posted as comment for the PR and automatically will be sent to my Telegram and Email. 
* If everything is OK and code is ready to merge, then pull request will be closed and feature branch will be merged to the main. 
* When code is pushed to the main, then code docker image building process is started. After successfully code building container will be pushed to the ECR golang-app-dev repository with a unique tag that is commit hash.
* Then Terraform will do its part: create a Dev ECS Service, tasks definition, target group and 9090 port listener. Now service can be accessed by the 9090 port of the ALB.
* Same is for the Blue and Green ECS Services. But each deployment of them requied a review and approval. e.g. It allows to change traffic between Blue and Green by [Redirect Traffic](https://github.com/RainbowGravity/golang-app/actions/workflows/redirect_traffic.yml) Github workflow.
* After process is complete [automated workflow](https://github.com/RainbowGravity/golang-app/actions/workflows/destroy_gruen.yml) of redirecting a traffic back to Blue and destroying the Green is started. New service is the Blue one now.

<p align=center><b>Workflow diagram</b></p>
<p align=center>

  <img width="800" height="357" src="https://user-images.githubusercontent.com/89798605/138294880-c891d5e3-a818-42e7-9956-96b3c401a712.png">

</p>
