# Gate - Minecraft Proxy

Gate is a lightweight Minecraft proxy designed to forward connections to mc.hypixel.net:25565.

## Features

- Lightweight and efficient
- Simple configuration
- Easy deployment on Render

## Deployment

This project is configured for deployment on [Render](https://render.com/).

1. Fork this repository
2. Create a new Web Service on Render
3. Connect it to your forked repository
4. Use the provided `render.yaml` for automatic configuration

## Local Development

To run the proxy locally:

```bash
go run src/main.go
```

The proxy will listen on port 25565 and forward connections to mc.hypixel.net:25565.

## Configuration

The proxy is configured in `src/main.go` to forward connections to mc.hypixel.net:25565. You can modify this file to change the target server or listen port.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.