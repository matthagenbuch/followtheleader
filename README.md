# Follow the Leader

A simple spanner-based leader election example which copies the hard parts from hashicorp/vault:
https://github.com/hashicorp/vault/blob/v1.13.3/sdk/physical/physical.go
https://github.com/hashicorp/vault/blob/v1.13.3/physical/spanner/spanner.go


```
$ go run main.go
2023/07/17 10:06:51 configuring backend
2023/07/17 10:06:51 creating HA client
2023/07/17 10:06:52 configuration database projects/lightstep-dev/instances/development/databases/dev-matth table Vault haEnabled true haTable alertevaluator_elections_ha maxParallel 0
2023/07/17 10:06:52 creating client
2023/07/17 10:06:52 running for leadership
2023/07/17 10:06:58 elected as leader
```