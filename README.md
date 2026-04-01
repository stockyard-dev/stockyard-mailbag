# Stockyard Mailbag

**Transactional email service — SMTP relay with dashboard, templates, delivery tracking**

Part of the [Stockyard](https://stockyard.dev) family of self-hosted developer tools.

## Quick Start

```bash
docker run -p 9230:9230 -v mailbag_data:/data ghcr.io/stockyard-dev/stockyard-mailbag
```

Or with docker-compose:

```bash
docker-compose up -d
```

Open `http://localhost:9230` in your browser.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `9230` | HTTP port |
| `DATA_DIR` | `./data` | SQLite database directory |
| `MAILBAG_LICENSE_KEY` | *(empty)* | Pro license key |

## Free vs Pro

| | Free | Pro |
|-|------|-----|
| Limits | 1 SMTP config, 500 sends/mo | Unlimited configs, unlimited sends |
| Price | Free | $4.99/mo |

Get a Pro license at [stockyard.dev/tools/](https://stockyard.dev/tools/).

## Category

Operations & Teams

## License

Apache 2.0
