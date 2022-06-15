wrk.method = "PUT"
wrk.headers["Content-Type"] = "application/json"

counter = 0
request = function()
    counter = counter + 1
    -- wrk.format(method, path, headers, body)
    return wrk.format(nil, nil, nil,  '{"task":"attemp #' .. counter .. '"}')
end