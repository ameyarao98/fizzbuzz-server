import redis.asyncio as redis
from litestar import get
from litestar.exceptions import HTTPException
from litestar.status_codes import HTTP_400_BAD_REQUEST

from .fizzbuzz import generate_fizz_buzz
from .redis import generate_redis_key, get_highest_count, increase_counter

HEALTH_RESPONSE = "fizz buzz: python"


@get()
async def get_health() -> str:
    return HEALTH_RESPONSE


@get()
async def get_fizzbuzz(
    redis_client: redis.Redis, int1: int, int2: int, limit: int, str1: str, str2: str
) -> list[str]:
    if int1 <= 0:
        raise HTTPException(
            status_code=HTTP_400_BAD_REQUEST,
            detail="int1 must be a positive integer",
        )

    if int2 <= 0:
        raise HTTPException(
            status_code=HTTP_400_BAD_REQUEST,
            detail="int2 must be a positive integer",
        )

    if limit <= 0:
        raise HTTPException(
            status_code=HTTP_400_BAD_REQUEST,
            detail="limit must be a positive integer",
        )

    result = generate_fizz_buzz(int1, int2, limit, str1, str2)
    key = generate_redis_key(int1, int2, limit, str1, str2)
    await increase_counter(redis_client, key)
    return result


@get()
async def get_statistics(redis_client: redis.Redis) -> dict[str, float]:
    highest_count = await get_highest_count(redis_client)
    return highest_count
