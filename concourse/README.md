## Concourse Pipeline

This is a [concourse](https://concourse.ci) pipeline which create a new release of mattermost-cf-integrator 
when a new mattermost version exists.

This is not building a mattermost-cf-integrator, by creating a new release this will trigger travis which will build and send to release the package.

You shouldn't using it in your own concourse. You can do so if you see that new releases are not created anymore.

