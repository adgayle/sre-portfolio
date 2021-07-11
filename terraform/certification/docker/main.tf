resource "docker_image" "nginx" {
  name = "nginx:stable"
}

resource "docker_container" "web" {
  image   = docker_image.nginx.latest
  name    = "web"
  restart = "on-failure"
  rm      = false

  ports {
    internal = 80
    external = 8080
  }
}
