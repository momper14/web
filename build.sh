./make_bin.sh

tar -cvf app.tar static/ templates/ web
docker build -t mmomper/web:latest .
rm app.tar