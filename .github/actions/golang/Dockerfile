FROM golang:latest

LABEL "name"="Golang workflow" \
    "com.github.actions.icon"="code" \
    "com.github.actions.color"="green-dark" \
    "com.github.actions.name"="golang　workflow" \
    "com.github.actions.description"="This is an Action to run go and go test commands."

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
