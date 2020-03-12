rm -rf /tmp/tendermint-server
cd $GOPATH/src/github.com/tendermint/tendermint
make install
TMHOME="/tmp/tendermint-server" tendermint init
./tendermint-server -config "/tmp/tendermint-server/config/config.toml"

# Now open another tab in your terminal and try sending a transaction
curl -s 'localhost:26657/broadcast_tx_commit?tx="tendermint=rocks"'

# Now let’s check if the given key now exists and its value:
curl -s 'localhost:26657/abci_query?data="tendermint"'
# {
#  "jsonrpc": "2.0",
#  "id": "",
#  "result": {
#    "response": {
#      "log": "exists",
#      "key": "dGVuZGVybWludA==",
#      "value": "cm9ja3M="
#    }
#  }
# }
# “dGVuZGVybWludA==” and “cm9ja3M=” are the base64-encoding of the ASCII of “tendermint” and “rocks” accordingly.