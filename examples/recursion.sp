value factorial = f(n) ->
    value z = 1

    if (n > 1) ->
        update z = (n * factorial(n - 1))
    end

    z
end

print(string(factorial(5)))