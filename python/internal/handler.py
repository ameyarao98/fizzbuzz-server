from litestar import get
from litestar.exceptions import HTTPException
from litestar.status_codes import HTTP_400_BAD_REQUEST

from .fizzbuzz import generate_fizz_buzz

HEALTH_RESPONSE = "fizz buzz: python"


@get()
async def get_health() -> str:
    return HEALTH_RESPONSE


@get()
async def get_fizzbuzz(
    int1: int, int2: int, limit: int, str1: str, str2: str
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
    return result
