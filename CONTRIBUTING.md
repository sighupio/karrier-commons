# Contribution

Read the following guide about how to contribute to this project and get familiar with.

Project structure based on:
[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Make

This project contains an easy-to-use interface with a `Makefile`. It allows you to pass all the checks the pipeline
will pass on certain events. So it worth get familiar with all the targets. To know what targets are available in
the project just run:

```bash
$ make

 Choose a command to run in fip-commons:

  init            Init the project. GITHUB_PROJECT=demo make init
  drone-init      Init the drone-project. GITHUB_PROJECT=demo GITHUB_TOKEN=123token321 DRONE_TOKEN=tokenhere REGISTRY=registry.sighup.io REGISTRY_USER=robotuser REGISTRY_PASSWORD=thepassword make drone-init
  lint            Run the policeman over the repository
  test            Run unit testing
  license         Check license headers are in-place in all files in the project
  clean-%         Clean the container image resulting from another target. make build clean-build

```

### init

The `init` target is required to be executed just one time after creating the project from the template.
Once the project is initialized you'll receive a message like:

```bash
$ GITHUB_PROJECT=a make init
Project already initialized with name demo-check
```

You can read more about this target in
[the template GitHub repository](https://github.com/sighupio/fip-commons)

### drone-init

The `drone-init` target is required to be executed just one time after creating the project from the template.
Once the project is initialized you'll receive a message like:

```bash
$ GITHUB_PROJECT=demo GITHUB_TOKEN=123token321 DRONE_TOKEN=tokenhere REGISTRY=registry.sighup.io REGISTRY_USER=robotuser REGISTRY_PASSWORD=thepassword make init
Project already initialized with name demo-check
```

It automates the configuration of the drone project with a single command, it requires a privilege token to run it.
If you don't have it, read more about how to set up it manually
[in the docs](https://github.com/sighupio/fip-commons)

### lint

This `make` target makes easy check if the (company-wide) linter pass or not. This is a basic requirement in order to
build and release a project at SIGHUP.io

The `lint` target uses the
[`policeman` *(with actually it is a wrapper around GitHub super-linter)*](https://github.com/sighupio/fury-kubernetes-policeman)
to check if everything is fine.

This project contains [a couple of additional linter rules](.rules) that has been tested in a real environment.

### test

The `test` target executes the command defined in the `tester` [`Dockerfile`](build/builder/Dockerfile)
target in order to unit test the application.

Again, if you need to add a different command to test the application, you can do it by modifying the right
chunk of the [`Dockerfile`](build/builder/Dockerfile):

```Dockerfile
.
..
...
RUN addlicense -c "SIGHUP s.r.l" -v -l bsd --check .

FROM golang:1.16 as tester

RUN mkdir /app
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY pkg pkg
COPY internal internal

RUN go test -v ./... -cover

FROM golang:1.16 as builder

RUN mkdir /app
...
..
.
```

Currently it executes `go test ./...` but feel free to modify with the right command.

### license

This make target checks for the license header in all the files. If you add new files that doesn't contain the
required header, you'll need to add them manually by:

```bash
go get -u github.com/google/addlicense
addlicense -c "SIGHUP s.r.l" -v -l bsd .
```

Then, you'll be able to re-check the `license` again:

```bash
make license
```

### clean-%v

The `clean-%v` target has been designed to remove the local built image resulting from the different targets in the
[`Makefile`](Makefile).

The main reason to implement this target is to save disk space.

Example usages:

- `make build clean-build`
- `make lint clean-lint`
- `make license clean-license`
- `make test clean-test`

## Pipeline

This project pipeline executes the `make` targets mentioned above.

It executes on every action to the Git Repository:

- `license`. It checks for the license headers.
- `lint`. It checks for common linting rules.
- `test`. It runs the tests of the project.

And finally, if the repo got a new `tag`:

- `release`: Creates a new GitHub release when a new `tag` is pushed to the git repo.
It uses the `docs/releases/${DRONE_TAG}.md`. It has to be present before creating the tag.


## Releases

This repository has been designed to follow trunk-based development. Then to create a new
release, you need to create a file describing the release with the name as the name of the
release.

You can see an example here: `docs/releases/v0.1.0.md`.

Then, in order to release it:

```bash
git tag v0.1.0
git push --tags
```
