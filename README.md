# mattermost-cf-integrator [![Build Status](https://travis-ci.org/cloudfoundry-community/mattermost-cf-integrator.svg)](https://travis-ci.org/cloudfoundry-community/mattermost-cf-integrator)
Integrate [mattermost](http://www.mattermost.org/) into Cloud Foundry.

[Mattermost](http://www.mattermost.org/) is an open source, on-prem Slack-alternative.
It offers modern communication from behind your firewall, including messaging and file sharing across PCs and phones with archiving and instant search.

## Getting started

1. Download latest release called `mattermost-cf.zip` in [releases page][1]
2. Unzip it
3. go inside the previously unzipped folder and push the app (e.g: `cf push mattermost -b binary_buildpack`), it will fail cause you need a database service connected
4. Create a mysql or postgres service on Cloud Foundry (e.g: `cf cs p-mysql 100mb db-mattermost`)
5. Bind to the app (e.g: `cf bs mattermost db-mattermost`)
6. restage your app (e.g: `cf restage mattermost`) and you're done


## Add SMTP Server (**Recommended**)

If you don't have an smtp server configured in your `config/config.json` mattermost will run in preview mode.
To fully use mattermost you should add a smtp, you have two ways to do it:

1. Do it manually by editing the `config/config.json` following the doc: http://docs.mattermost.com/install/smtp-email-setup.html
2. (*Preferred*) Bind a smtp service on your app and integrator will do the rest, example with Pivotal Web Service and sendgrid:

 ```
 $ cf cs sendgrid free mysmtp
 $ cf bs mattermost mysmtp
 $ cf restage mattermost
 ```

## Add an Amazon S3 Bucket (**Recommended**)

Mattermost by default will store data send by users (images, files, video ...) on the local file system but on Cloud Foundry app shouldn't write on the filesystem.
You should set an s3 storage on your mattermost, like for SMTP you have two ways to do it:

1. Do it manually by editing the `config/config.json` following the doc: http://docs.mattermost.com/administration/config-settings.html?highlight=amazons3endpoint#file-settings
2. Bind a s3 service on your app and integrator will do the rest, example [riak-cs](https://github.com/cloudfoundry/cf-riak-cs-release) service:

 ```
 $ cf cs p-riakcs developer mys3
 $ cf bs mattermost mys3
 $ cf restage mattermost
 ```


[1]: https://github.com/cloudfoundry-community/mattermost-cf-integrator/releases