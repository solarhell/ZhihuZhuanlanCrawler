workflow "Golang workflow" {
  on = "push"
  resolves = ["Test"]
}
 
action "Test" {
  uses = "./.github/actions/golang"
  args = "test"
}
