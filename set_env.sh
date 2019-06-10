if [ "$WEBPASS" == "" ]; then echo "WEBPASS nicht gesetzt!" 
else
    export COUCHDB_URL="http://admin:$WEBPASS@localhost:5984"
fi