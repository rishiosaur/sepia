value fibonacci = f(n) -> 

    if (n == 0) ->
        return 0;
    end

    if (n == 1 or n == 2) ->
        return 1;
    end

    fibonacci(n - 1) + fibonacci(n - 2)
end

print(string(fibonacci(5)))