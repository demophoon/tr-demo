project = "targetable-runners"
app "targetable-runners" {

  // New runner profile targeting
  runner {
    profile = "k8s-prod-profile"
  }

  build {
    use "docker" {}
    registry {
      use "docker" {
        image = "registry.services.demophoon.com/tr-demo"
        tag = "latest"
      }
    }
  }
  deploy {
    use "kubernetes" {}
  }
}
