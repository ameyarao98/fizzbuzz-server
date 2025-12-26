def generate_fizz_buzz(
    int1: int, int2: int, limit: int, str1: str, str2: str
) -> list[str]:
    result: list[str] = []

    for i in range(1, limit + 1):
        s = ""
        if not i % int1:
            s += str1
        if not i % int2:
            s += str2
        result.append(s or str(i))

    return result
