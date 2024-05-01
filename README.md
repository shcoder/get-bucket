# get-bucket



## Установка
`go install github.com/shcoder/get-bucket@v0.1.3`

## Примеры использованич
Получить бакет и шард для posting-service для компании 1234
`get-bucket -ps 1234`

Получить бакет и шард для posting-data-service для компании 1234
`get-bucket -pds 1234`

Получить бакет и шард для кастомной базы для компании 1234
`get-bucket -schema 12 -shard 12 535`

