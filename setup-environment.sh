docker run -e POSTGRES_PASSWORD=password -e POSTGRES_USER=user -e POSTGRES_DB=testdb -p 5432:5432 -d postgres
export DATABASE_URL="postgres://user:password@localhost:5432/testdb"
sleep 5
psql $DATABASE_URL -c "CREATE TABLE users ( name TEXT, lastname TEXT);"
export PORT=8080
