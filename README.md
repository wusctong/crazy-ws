# CraftSocket Proxy for Render

This project deploys a CraftSocket proxy service to Render that connects to a Minecraft server (hub.mcsrv.top:25565) and exposes a WebSocket endpoint.

## Deployment to Render

1. Fork this repository to your GitHub account
2. Create a new Web Service on Render
3. Connect it to your forked repository
4. Set the build command to use Docker
5. Deploy!

The service will automatically start and listen on the Render-provided `$PORT` environment variable.

## Configuration

The proxy is configured to:
- Connect to `hub.mcsrv.top:25565` (Minecraft server)
- Expose a WebSocket endpoint at `/boost`
- Listen on the port provided by Render (`$PORT`)

## Usage

Once deployed, you can connect to the WebSocket endpoint at:
`wss://<your-render-url>/boost`

Replace `<your-render-url>` with the actual URL provided by Render.