data_input = input().split(" ")

def is_canada_win(a: int, b: int, n: int) -> bool:
    return (a + n >= b)