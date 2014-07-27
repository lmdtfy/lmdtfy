lmdtfy - Let me deploy that for you
======


## Stages

* Git Hook (Push)
    * Kicks off a build.

* Build
    * Parses .yml file
    * Starts a container for the environment. (ruby, golang, node, etc...)
    * Runs custom build commands.
    * Builds Project
        * Fail: Ends process and logs errors.
        * Pass: Will then start Testing.

* Test
    * Runs all tests in a test env.
        * Fail: Ends process and logs errors.
        * Pass: Deploys to Staging or Live environment.

* Deploy
    * 
