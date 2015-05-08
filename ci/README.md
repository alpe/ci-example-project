## CoreOS unit
Used by `systemd` and `fleet`.

### Service
Set user that systemd picks the `.dockercfg` file in the users home folder.
~~~
User=core
~~~
### ClI commands
* Submit unit file
~~~bash
fleetctl submit ci-example@.service
fleetctl list-unit-files
~~~

* Start 1 instance
~~~bash
fleetctl start ci-example@1.service
# status
fleetctl statuts ci-example@1.service
~~~

* Stop the instance
~~~bash
fleetctl stop ci-example@1.service
~~~

* Remove the unit file
~~~bash
fleetctl destroy ci-example@.service
~~~