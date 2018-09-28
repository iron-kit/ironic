package template

var (
	DockerFNC = `FROM alpine:3.2
ADD {{.Alias}}-{{.Type}} /{{.Alias}}-{{.Type}}
ENTRYPOINT [ "/{{.Alias}}-{{.Type}}" ]
`

	DockerSRV = `FROM alpine:3.2
RUN mkdir /app
WORKDIR /app
ADD {{.Alias}}-{{.Type}} /app/{{.Alias}}-{{.Type}}
ENTRYPOINT [ "/app/{{.Alias}}-{{.Type}}" ]
CMD ["./{{.Alias}}-{{.Type}}"]
`

	DockerWEB = `FROM alpine:3.2
ADD html /html
ADD {{.Alias}}-{{.Type}} /{{.Alias}}-{{.Type}}
WORKDIR /
ENTRYPOINT [ "/{{.Alias}}-{{.Type}}" ]
`
)
