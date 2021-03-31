# Contributing to the Airplane Go SDK

## Releases

Releases are managed by git tags that follow semver. To create a new release:

```sh
export AIRPLANE_TAG=v0.0.1 && \
  git tag ${AIRPLANE_TAG} && \
  git push origin ${AIRPLANE_TAG}
```
