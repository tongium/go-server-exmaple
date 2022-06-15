counter = 0
request = function()
    counter = counter + 1
    path = wrk.path:gsub("%:id", counter)
    return wrk.format(nil, path)
end