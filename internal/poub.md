#	Ways to replace placeholders in URL

```go
replacements := map[string]string{
	"$NAME": cliConf.Name,
	"$TAG":  cliConf.Tag,
	"$OS":   osType,
	"$ARCH": osArch,
}

url := cliConf.Url
for placeholder, value := range replacements {
	url = strings.ReplaceAll(url, placeholder, value)
}
```