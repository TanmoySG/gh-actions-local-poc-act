# gh-actions-local-poc-act
Trying out https://github.com/nektos/act to run GitHub Actions Locally 🚀

## How to Run?

Install act using `brew`
```
brew install act
```
Go to a repo with `.github/workflows` , create a workflow. This repo is a good starting point!

Run `act` in the directory with the workflow files.

```
act
``` 

A prompt asks to select the resource/size of the image to run act on. Micro should be fine.

```
? Please choose the default image you want to use with act:

  - Large size image: +20GB Docker image, includes almost all tools used on GitHub Actions (IMPORTANT: currently only ubuntu-18.04 platform is available)
  - Medium size image: ~500MB, includes only necessary tools to bootstrap actions and aims to be compatible with all actions
  - Micro size image: <200MB, contains only NodeJS required to bootstrap actions, doesn't work with all actions
```

The server now be should run on docker. **Docker must be Running.**

To see the the logs run
```
act -v
```

To See the Graph of the Workflow
```
act -g
```

See more here - https://github.com/nektos/act#example-commands

### Known Issues

- While setting up go and dependancies, there were some ca-caertificate error and if any package requires C , GCC Compiler also needs to be configured, hence the apt-get's 
