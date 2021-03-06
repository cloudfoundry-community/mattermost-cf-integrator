#!/usr/bin/env bash

set -xe

CW=$(pwd)

send_slack() {
    if [ "x$slack_url" == "x" ]; then
        return 0
    fi
    message=$1
    insecure=""
    if [ "x$slack_insecure" != "x" ]; then
        insecure="--insecure"
    fi
    notifslack $insecure --url $slack_url -c $slack_channel -u $slack_username -i $slack_icon "$message"
}

actual_version=$(curl -s http://about.mattermost.com/download/ -L | pup 'p.release-text span text{}' | sed -n 2p | awk '{print $1}')
current_version=$(head -n 1 "$CW/mattermost-integrator-release/tag")
if [ "v$actual_version" == "$current_version" ]; then
    echoc "[yellow]You should don't care about this error."
    echoc "[green]The version $actual_version already exists."
    exit 1
fi
semver=$(echo "$actual_version" | sed 's/v//g')
echo "$semver" | grep -Eq "^[0-9]{1,2}\.[0-9]{1,2}\.[0-9]{1,2}$"
if [ $? -ne 0 ]; then
    message="[mattermost] Parsed version do not follow semantic versioning, check https://www.mattermost.org/download page and fix the parsing"
    echo "$message"
    send_slack "$message"
    exit 1
fi
send_slack "Mattermost-cf-integrator bump its version to $semver"
echo "$semver" > "$CW/release-info/tag_to_release"
echo "Mattermost $semver on Cloud Foundry" > "$CW/release-info/name_of_release"
echo "See changelog from mattermost: https://docs.mattermost.com/administration/changelog.html#release-v$semver" > "$CW/release-info/body"
exit 0