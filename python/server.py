import os

import uvicorn
from litestar import Litestar

from internal.router import api_router

app = Litestar(route_handlers=[api_router])

if __name__ == "__main__":
    uvicorn.run(
        "server:app",
        host=os.getenv("HOST", ""),
        port=int(os.getenv("PORT", "0")),
    )
