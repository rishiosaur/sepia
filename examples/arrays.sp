value map = f(arr, fn) ->
    value iterateOnArr = f(arr, accumulator) ->
        if (len(arr) == 0) ->
            accumulator
        end

        else ->
            iterateOnArr(rest(arr), append(accumulator, fn(first(arr))))
        end
    end

    iterateOnArr(arr, [])
end

value arr = ["hello", "world"]

value addWorld = f(str) ->
    str + "world"
end

print(len(map(arr, addWorld)))