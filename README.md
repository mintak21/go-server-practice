# GoGin Practice With Clean Architecture API

# API Descriptions

**A Simple API**

Frameworks used in this project below:

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/jinzhu/gorm)
- postgres

# Entity Relation
`users` (1)<------>(many) `comments`

# Directory Structure

### `cmd`
Main Application for this project.
Includes small main function that invokes the codes form `/pkg`.

### `pkg`
Main Architecture codes are inculeded.
`pkg` directories are based on Clean Architecture.

- `domain` : `Entities` Layer
- `usecase` : `Use cases` Layer
- `interface` : `Interface Adapters` Layer
- `infra` : `Frameworks And Drivers` Layer

### `build`
- `build/package` : Container Configurations = Dockerfile
- `build/ci : CI tools Configurations. In this Project, using `CircleCI`.

### `deployments`
Container orchestration deployment configurations such as `docker-compose`, `k8s`.

# Reference Pages

- [Qiita]
	- [駆け出しエンジニアのGo入門](https://qiita.com/_nas/items/1dae3a19be2256890649)
	- [Goで簡単TwitterAPIを作った](https://qiita.com/kohama66/items/735e70e6e3215942b02f)
	- [GormとGinによるWebApp](https://qiita.com/Anharu/items/ce644c521a4d52fafb7e)
- [Github]
	- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

