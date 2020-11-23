# A standard factorial implementation.

value factorial = f(n) ->
    if (n < 0) ->
        -1
    end
    
    else ->
        if (n == 0) ->
            1
        end

        else ->
            (n * factorial(n - 1))
        end
    end
end

value n = 16

factorial(n)