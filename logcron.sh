SH_FILE="logcron.sh"
DIR_PATH=$0
DIR_PATH=${DIR_PATH%"$SH_FILE"}

EXEC_FILE="google-cloud-ddns"
EXEC_FILEPATH="$DIR_PATH$EXEC_FILE"

LOGS_FILE="logs.txt"
LOGS_FILEPATH="$DIR_PATH$LOGS_FILE"

OUTPUT=$($EXEC_FILEPATH)

if [ "${#OUTPUT}" -ne 0 ]
then
	echo "$(date): ${OUTPUT}" >> $LOGS_FILEPATH
fi
