if [ -f .env ]; then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

DB_HOST=${DB_HOST}
DB_PORT=${DB_PORT:-5432}

echo "Waiting for database at $DB_HOST:$DB_PORT..."
while ! nc -z $DB_HOST $DB_PORT; do
  sleep 1
done
echo "Database is ready!"

exec "$@"