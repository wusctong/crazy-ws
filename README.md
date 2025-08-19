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

You can also customize the target server using environment variables:
```bash
PROXY_TARGET_HOST=your-server.com PROXY_TARGET_PORT=25565 go run src/main.go
```

## Docker

To build and run using Docker:

```bash
docker build -t gate-proxy .
docker run -p 25565:25565 gate-proxy
```

You can customize the target server using environment variables:
```bash
docker run -e PROXY_TARGET_HOST=your-server.com -e PROXY_TARGET_PORT=25565 -p 25565:25565 gate-proxy
```

## Configuration

The proxy can be configured using the following environment variables:
- `PROXY_TARGET_HOST` - The target Minecraft server host (default: mc.hypixel.net)
- `PROXY_TARGET_PORT` - The target Minecraft server port (default: 25565)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.