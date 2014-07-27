lmdtfy - Let me deploy that for you
======

### Job
A job is usually scheduled as a result of a github push or pull request.
It is then scheduled to a docker host in the cluster.

    - Current Stage
    - Branch
    - Status
    - Started On
    - Finished On
    - Total Time


### Stages
Any failed task in a 'stage' will result in a failed job.


* Git Hook (Push or Pull Request)
    * Creates a new Job with link to code and branch


* Job
    * Pulls code
    * Schedules to a node in the cluster
    * Build is kicked off


* Build
    * Parses .yml file
    * Starts a container for the environment. (ruby, golang, node, etc...)
    * Runs custom build commands
    * Builds Project
        * Fail: Ends process and logs errors
        * Pass:
            * Push to local docker registry tagged with `env-rev/commit`
            * Start Testing

* Test
    * Runs all tests in a test env
        * Fail: Ends process and logs errors
        * Pass: Deploys to Staging or Live environment


* Deploy
    *
    *
