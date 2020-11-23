val while = f(whileTrue, doFunc) ->
   if (whileTrue()) ->
       doFunc()
       while(whileTrue, doFunc)
   end

   else ->
       true
   end
end

val i = 0

val fn = f() ->
    val i = i+1
end

val isTrue = f() ->
    i < 2
end

print(string(while(isTrue, fn)))