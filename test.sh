# echo "=== Test Reflection API ==="
grpcurl -plaintext localhost:8080 describe

echo "=== Insert Test Data ==="

grpcurl -plaintext -d  '{ "title": "christmas eve bike class", "url": "https://library", "config": "HJ library\nDJ library\nHJ www.library.ca" }' localhost:8080 api.v1.EntryServices/Insert

echo "=== Test Retrieve Descriptions ==="

grpcurl -plaintext -d '{ "id": 1 }' localhost:8080 api.v1.EntryServices/Retrieve | grep -q 'christmas eve bike class'

echo "=== Test List ==="

grpcurl -plaintext localhost:8080 api.v1.EntryServices/List | jq '.activities | length' |  grep -q '1'

echo "Success"
