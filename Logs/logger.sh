timestamp=$(date +%Y-%m-%d)
logfile="./Build/build_$timestamp.log"

logger() {
    local level="$1"
    local result=$2
    local emoji=$3
    local message="$4"
	local process="$5"
    local timestamp=$(date +%Y-%m-%d-%H-%M-%S)

    echo "{'timestamp': '$timestamp', 'level': '$level', 'message': '$message', 'process': '$process', 'exit code': '$code'}" >> $logfile
}