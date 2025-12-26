FROM python:3.12-slim-trixie
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

WORKDIR /app

COPY pyproject.toml uv.lock /app/
RUN uv sync --locked --no-dev

COPY server /app/server

CMD ["uv", "run", "server/server.py"]
