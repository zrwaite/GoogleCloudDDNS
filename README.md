# Google Cloud DDNS

My server has a dynamic DNS, but I want to point all my domains at it's nginx instance.

This script runs with a cronjob on the server and updates the google cloud DNS records upon a change.

Add to crontab script (`crontab -e`): `* * * * * <your dir>/logcron.sh`

Todo:
-   Perhaps make it more abstract for people to fork.
