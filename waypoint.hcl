project = "example-go"
app "example-go" {

  // New runner profile targeting
  runner {
    profile = "k8s-prod-profile"
  }

  build {
    use "pack" {}
  }
  deploy {
    use "kubernetes" {}
  }
}
