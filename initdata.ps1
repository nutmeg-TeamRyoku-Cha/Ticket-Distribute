# APIが起動している状態のPowerShellで
# sh init.shとかしてください

# 建物
$body = @{
    building_name = "講義棟2階"
    latitude = 35.697403
    longitude = 139.701216
} | ConvertTo-Json
$headers = @{ "Content-Type" = "application/json; charset=utf-8" }
Invoke-RestMethod -Uri "http://localhost:8080/buildings" -Method POST -Headers $headers -Body $body

$body = @{
    building_name = "キッチンカー"
    latitude = 35.697403
    longitude = 139.701216
} | ConvertTo-Json
$headers = @{ "Content-Type" = "application/json; charset=utf-8" }
Invoke-RestMethod -Uri "http://localhost:8080/buildings" -Method POST -Headers $headers -Body $body

# 企画
$body = @{
    project_name      = "お化け屋敷"
    building_id       = 1
    requires_ticket   = $true
    remaining_tickets = 300
    start_time        = "2025-10-26T09:00:00Z"
    end_time          = "2025-10-26T18:00:00Z"
} | ConvertTo-Json
Invoke-RestMethod -Uri http://localhost:8080/projects -Method Post -Body $body -ContentType "application/json"

$body = @{
    project_name      = "スマブラ大会"
    building_id       = 1
    requires_ticket   = $true
    remaining_tickets = 150
    start_time        = "2025-10-27T09:00:00Z"
    end_time          = "2025-10-27T16:00:00Z"
} | ConvertTo-Json
Invoke-RestMethod -Uri http://localhost:8080/projects -Method Post -Body $body -ContentType "application/json"

$body = @{
    project_name      = "すし"
    building_id       = 2
    requires_ticket   = $true
    remaining_tickets = 1000
    start_time        = "2025-10-26T09:00:00Z"
    end_time          = "2025-10-26T18:00:00Z"
} | ConvertTo-Json
Invoke-RestMethod -Uri http://localhost:8080/projects -Method Post -Body $body -ContentType "application/json"

$body = @{
    project_name      = "ラーメン"
    building_id       = 2
    requires_ticket   = $true
    remaining_tickets = 1000
    start_time        = "2025-10-27T09:00:00Z"
    end_time          = "2025-10-27T16:00:00Z"
} | ConvertTo-Json
Invoke-RestMethod -Uri http://localhost:8080/projects -Method Post -Body $body -ContentType "application/json"

$body = @{
    project_name      = "ワークショップ"
    building_id       = 1
    requires_ticket   = $false
    remaining_tickets = 1000
    start_time        = "2025-10-26T09:00:00Z"
    end_time          = "2025-10-27T16:00:00Z"
} | ConvertTo-Json
Invoke-RestMethod -Uri http://localhost:8080/projects -Method Post -Body $body -ContentType "application/json"

$body = @{
    project_name      = "縁日"
    building_id       = 2
    requires_ticket   = $false
    remaining_tickets = 1000
    start_time        = "2025-10-27T09:00:00Z"
    end_time          = "2025-10-27T16:00:00Z"
} | ConvertTo-Json
Invoke-RestMethod -Uri http://localhost:8080/projects -Method Post -Body $body -ContentType "application/json"