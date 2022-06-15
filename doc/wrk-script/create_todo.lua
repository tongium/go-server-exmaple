wrk.method = "PUT"
wrk.headers["Content-Type"] = "application/json"

counter = 0
request = function()
    counter = counter + 1
    return wrk.format(nil, nil, nil,  '{"task":"attemp #' .. counter .. '"}')
 end