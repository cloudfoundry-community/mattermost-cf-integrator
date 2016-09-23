# mattermost-cf-integrator [![Build Status](https://travis-ci.org/cloudfoundry-community/mattermost-cf-integrator.svg)](https://travis-ci.org/cloudfoundry-community/mattermost-cf-integrator)
Integrate [mattermost](http://www.mattermost.org/) into Cloud Foundry.

**Why it's different from the [mattermost-boshrelease](https://github.com/cloudfoundry-community/mattermost-boshrelease) ?** the boshrelease made by stark&wayne is made for Cloud Foundry operators and only them can deploy mattermost in an IaaS (not Cloud Foundry in fact). Here it's made for final Cloud Foundry users who want a private mattermost running inside a Cloud Foundry like [Bluemix](https://console.ng.bluemix.net/) or [PWS](https://run.pivotal.io).

[Mattermost](http://www.mattermost.org/) is an open source, on-prem Slack-alternative.
It offers modern communication from behind your firewall, including messaging and file sharing across PCs and phones with archiving and instant search.

![mattermost-preview](http://www.mattermost.org/wp-content/uploads/2015/09/20160315_v210.png)

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

## Integrate with UAA (to authenticate users over Cloud Foundry)

It can be interesting to login mattermost with Cloud Foundry. Mattermost can use GitLab SSO to login users and this functionality can be use to login over Cloud Foundry.

To do this we will need to register a client in the UAA, here an example to register a client:

```
$ uaac client add mattermost --authorized_grant_types "authorization_code" --name "mattermost" --scope "openid" --authorities "openid" -s "mysupersecret"
```

You will also need to deploy another app on Cloud Foundry, this app will be responsible to translate user informations from UAA to GitLab.
This app is a reverse proxy on your UAA which only transform json receive from the UAA on https://login.url-of-my.api/userinfo to another json asked by mattermost.

Follow these steps to deploy this application:

1. clone https://github.com/ArthurHlt/uaa-proxifier.git : `$ git clone https://github.com/ArthurHlt/uaa-proxifier.git`
2. set in the value for `UAA_URL` in `uaa-proxifier/manifest.yml` by the url of your UAA
3. deploy it: `$ cf push`


Now we can set in the file `config/config.json` endpoints for UAA and credentials, example:

```
"GitLabSettings": {
    "Enable": true,
    "Secret": "mysupersecret",
    "Id": "mattermost",
    "Scope": "openid",
    "AuthEndpoint": "https://login.url-of-my.api/oauth/authorize",
    "TokenEndpoint": "https://login.url-of-my.api/oauth/authorize/oauth/token",
    "UserApiEndpoint": "https://uaa-to-gitlab.my.cf.domain.com/userinfo"
}
```

**Note** the value for `UserApiEndpoint` this is the url of you previously deployed app.

To make users only login and register with cloudfoundry set to `false` values `EnableSignUpWithEmail`, `EnableSignInWithEmail` in `config/config.json`

Now, you can (re)deploy your mattermost: `$ cf push mattermost`

### (**Optional**) Remove reference about GitLab in Web UI


#### Change the word `GitLab` by `Cloud Foundry`

**For version 2.X.X**:

 ```
 $ sed -i "" 's/GitLab/Cloud Foundry/g' web/static/i18n/*`
 ```

 **For version 3.X.X**:

  ```
  $ sed -i "" 's/GitLab/Cloud Foundry/g' webapp/dist/i18n/*
  $ sed -i "" 's/"GitLab/"Cloud Foundry/g' webapp/dist/*.js
  ```

#### Change the GitLab logo by the Cloud Foundry logo:

**For version 2.X.X**:

 ```
 $ wget https://rawgit.com/cloudfoundry-community/mattermost-cf-integrator/master/cloudfoundryLogo.png
 $ cp cloudfoundryLogo.png web/static/images/gitlabLogo.png
 ```

**For version 3.X.X**:

 ```
 $ wget https://rawgit.com/cloudfoundry-community/mattermost-cf-integrator/master/cloudfoundryLogo.png
 $ cp cloudfoundryLogo.png webapp/dist/files/bf61680806a56e50a7857eeeea863f01.png
 ```

Now, you can (re)deploy your mattermost: `$ cf push mattermost`


[1]: https://github.com/cloudfoundry-community/mattermost-cf-integrator/releases
