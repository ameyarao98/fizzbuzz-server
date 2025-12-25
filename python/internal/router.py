from litestar import Router

from .handler import get_fizzbuzz, get_health, get_statistics

api_router = Router(
    path="",
    route_handlers=(
        Router(
            path="/health",
            route_handlers=(get_health,),
        ),
        Router(
            path="/fizzbuzz",
            route_handlers=(get_fizzbuzz,),
        ),
        Router(
            path="/statistics",
            route_handlers=(get_statistics,),
        ),
    ),
)
