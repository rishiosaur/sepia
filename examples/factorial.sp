# A standard factorial implementation.

val factorial = f(n) ->
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

val n = 16

factorial(n)