# Tools API

[![Go Report Card](https://goreportcard.com/badge/github.com/jasonbirchall/tools-api-poc)](https://goreportcard.com/report/github.com/jasonbirchall/tools-api-poc)
[![GoDoc](https://godoc.org/github.com/luraproject/lura?status.svg)](https://godoc.org/github.com/jasonbirchall/tools-api-poc)

The intention of this repository is to provide a simple API for the use of tool creation and lookup. At this time I wanted to keep the API as a separate entity and not use Kubernetes as an extension. Although this might be the case in the future.

This API is currently in a very early stage of development.

## Development

There are a few nuances I'd like this repository to have.

### GitFlow model

To structure the code within Git branches, we will use the GitFlow model. This approach consists of the following branches:

- main: This branch corresponds to the current production code. You can't commit directly, except for hotfixes. Git tags can be used to tag all the commits in the master branch with a version number (for instance, for using the semantic versioning convention, https://semver.org/, which has three parts: major, minor, and patch, so a tag with version 1.2.3 has 1 as its major version, 2 as its minor version, and 3 as its patch version).

- preprod: This is a release branch and is a mirror of production. It can be used to test all the new features that are developed on the develop branch before they are merged to the master branch.

- develop: This is the development integration branch, which contains the latest integrated development code.

- feature/X: This is an individual feature branch that's being developed. Each new feature resides in its own branch, and they're generally created for the latest develop branch.

- hotfix/X: When you need to solve something in production code, you can use the hotfix branch and open a pull request for the master branch. This branch is based on the master branch.

### Commit messages

Please if you can use [conventional commits](https://conventionalcommits.org/) to make your commits and follow the [conventional commit guidelines](https://conventionalcommits.org/en/v1.0.0/guidelines.html).

### Versioning

Please use [semantic versioning](https://semver.org/) for versioning when releasing. This will make it easier to track changes and to make it easier to find the latest version.

### Local development

Run the following commands to run locally:

```bash
docker compose up -d

go run main.go
```