variable "TAG" {
  default = "latest"
}

variable "REPOSITORY" {
  default = "docker.io"
}

group "default" {
  targets = []
}

target "verify-docker-context" {
  dockerfile = "dev-support/containers/alpine/verify-docker-context.dockerfile"
  target     = "verify-docker-context"
  tags       = ["${REPOSITORY}/library/verify-docker-context:${TAG}"]
}
