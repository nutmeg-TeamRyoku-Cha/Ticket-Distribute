# Proxmox用
#!/usr/bin/env bash
set -euo pipefail

# Nginx 経由（HTTPS）
BASE_URL="https://localhost/api"
CURL_OPTS=(-sS -k -H "Content-Type: application/json; charset=utf-8")

# BASE_URL="http://localhost:8080"
# CURL_OPTS=(-sS -H "Content-Type: application/json; charset=utf-8")

post() {
  local path="$1"; shift
  local body="$1"; shift || true
  echo "POST ${BASE_URL}${path}"
  echo "${body}" | curl "${CURL_OPTS[@]}" -X POST "${BASE_URL}${path}" -d @- || {
    echo "POST ${path} failed" >&2; exit 1;
  }
  echo
}

# 建物
post "/buildings" '{
  "building_name": "講義棟2階",
  "latitude": 35.697403,
  "longitude": 139.701216
}'

post "/buildings" '{
  "building_name": "キッチンカー",
  "latitude": 35.697403,
  "longitude": 139.701216
}'

# 企画
post "/projects" '{
  "project_name": "お化け屋敷",
  "building_id": 1,
  "requires_ticket": true,
  "remaining_tickets": 300,
  "start_time": "2025-10-26T09:00:00Z",
  "end_time":   "2025-10-26T18:00:00Z"
}'

post "/projects" '{
  "project_name": "スマブラ大会",
  "building_id": 1,
  "requires_ticket": true,
  "remaining_tickets": 150,
  "start_time": "2025-10-27T09:00:00Z",
  "end_time":   "2025-10-27T16:00:00Z"
}'

post "/projects" '{
  "project_name": "すし",
  "building_id": 2,
  "requires_ticket": true,
  "remaining_tickets": 1000,
  "start_time": "2025-10-26T09:00:00Z",
  "end_time":   "2025-10-26T18:00:00Z"
}'

post "/projects" '{
  "project_name": "ラーメン",
  "building_id": 2,
  "requires_ticket": true,
  "remaining_tickets": 1000,
  "start_time": "2025-10-27T09:00:00Z",
  "end_time":   "2025_
