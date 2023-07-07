token=$(cat tgtoken.key);
dsn="host=localhost port=5432 user=user password=password dbname=wb_helper sslmode=disable";

export TELEGRAM_API_TOKEN=${token}
export DSN_DB=${dsn}

echo TELEGRAM_API_TOKEN=${TELEGRAM_API_TOKEN}
echo DB_DSN=${DSN_DB}