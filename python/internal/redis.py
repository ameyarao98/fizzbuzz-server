import os
from collections.abc import AsyncGenerator

import redis.asyncio as redis

REQUESTS_KEY = "requests"


async def init_redis() -> AsyncGenerator[redis.Redis, None]:
    redis_addr = os.getenv("REDIS_DSN", "")
    redis_db = os.getenv("REDIS_DB")
    if redis_db is None or not redis_db.isdigit():
        raise ValueError(f"Invalid REDIS_DB value: {redis_db}")

    redis_db = int(redis_db)

    redis_client = redis.Redis(
        host=redis_addr.split(":")[0],
        port=int(redis_addr.split(":")[1]),
        db=int(redis_db),
        decode_responses=True,
    )
    try:
        yield redis_client
    finally:
        await redis_client.aclose()


def generate_redis_key(int1: int, int2: int, limit: int, str1: str, str2: str) -> str:
    return f"{int1}:{int2}:{limit}:{str1}:{str2}"


async def increase_counter(redis_client: redis.Redis, key: str) -> None:
    await redis_client.zincrby(REQUESTS_KEY, 1, key)
    return None


async def get_highest_count(redis_client: redis.Redis) -> dict[str, float]:
    result: list[tuple[str, float]] = await redis_client.zrevrange(
        REQUESTS_KEY, 0, 0, withscores=True
    )
    if not result:
        return {}

    member, score = result[0]
    return {member: score}
