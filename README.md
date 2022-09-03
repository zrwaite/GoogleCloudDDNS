# Google Cloud DDNS

My server has a dynamic DNS, but I want to point all my domains at it's nginx instance.

This script runs with a cronjob on the server and updates the google cloud DNS records upon a change.

Todo:
- Update params.json file upon Access Token Change.
- Add script for easier cronjob implementation. 
- Perhaps make it more abstract for people to fork.
