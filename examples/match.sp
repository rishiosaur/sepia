value fibonacci = f(n) -> match (n) {
    0 -> 0,
    1 -> 1,
    2 -> 1,
    default -> fibonacci(n-1) + fibonacci(n-2),
} end

print(fibonacci(6))