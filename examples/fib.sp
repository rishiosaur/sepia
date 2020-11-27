value fibonacci = f(n) ->
    value z = 0

    if (n == 1 || n == 2) ->
        update z = 1
    end
    else ->
        update z = fibonacci(n - 1) + fibonacci(n - 2)
    end
    z
end

print(string(fibonacci(32)))