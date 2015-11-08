# mattermost-cf-integrator
Integrate [mattermost](http://www.mattermost.org/) into Cloud Foundry.
[Mattermost](http://www.mattermost.org/) s an open source, on-prem Slack-alternative.
It offers modern communication from behind your firewall, including messaging and file sharing across PCs and phones with archiving and instant search.

## Getting start

1. Download latest release called `mattermost-cf.zip` in [releases page][1]
2. Unzip it
3. go inside the previously unzipped folder and push the app (e.g: `cf push mattermost -b binary_buildpack`), it will fail cause you need a database service connected
4. Create a mysql or postgres service on Cloud Foundry (e.g: `cf cs p-mysql 100mb db-mattermost`)
5. Bind to the app (e.g: `cf bs mattermost db-mattermost`)
6. restage your app (e.g: `cf restage mattermost`) and you're done

[1]: https://github.com/ArthurHlt/mattermost-cf-integrator/releases
