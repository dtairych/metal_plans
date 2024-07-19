# Standard output for Equinix Metal Server Plans

Golang code to interact with the Equinix Metal API which crates a standard output for all of the available server plans

To use, simply export your API token for Equinix Metal as an environment variable:

```console
export EQUINX_API_TOKEN="<insert token here>"
```

And then execute the module:

```console
go run main.go
```

An example printed output from a successful collection of the plans.  These are stored as structs that can be filtered on:

```console

Server: a3.large.opt-m4a4.x86
Description: a3.large.opt-m4a4.x86
Legacy:  false
Deployment Options:  []
CPUs:
  2 x Intel Xeon Gold 6338 Processor 64-Core @ 2.00GHz
    Manufacturer: Intel
    Model: 6338
    Cores: %!d(json.Number=64)
    Speed: 2.00GHz
Memory:
  Total: 1024GB
Drives:
  2 x 240GB SSD boot
NICs:
  4 x 25Gbps
------------------------
Server: a3.large.opt-s4a1.x86
Description: a3.large.opt-s4a1.x86
Legacy:  false
Deployment Options:  []
CPUs:
  2 x Intel Xeon Gold 6338 Processor 64-Core @ 2.00GHz
    Manufacturer: Intel
    Model: 6338
    Cores: %!d(json.Number=64)
    Speed: 2.00GHz
Memory:
  Total: 1024GB
Drives:
  2 x 240GB SSD boot
  6 x 7.84TB SSD storage
NICs:
  4 x 25Gbps
------------------------
Server: a3.large.opt-s4a5n1.x86
Description: a3.large.opt-s4a5n1.x86
Legacy:  false
Deployment Options:  []
CPUs:
  2 x Intel Xeon Gold 6338 Processor 64-Core @ 2.00GHz
    Manufacturer: Intel
    Model: 6338
    Cores: %!d(json.Number=64)
    Speed: 2.00GHz
Memory:
  Total: 1024GB
Drives:
  2 x 240GB SSD boot
NICs:
  6 x 25Gbps
------------------------
Server: a3.large.opt-s5a1.x86
Description: a3.large.opt-s5a1.x86
Legacy:  false
Deployment Options:  []
CPUs:
  2 x Intel Xeon Gold 6338 Processor 64-Core @ 2.00GHz
    Manufacturer: Intel
    Model: 6338
    Cores: %!d(json.Number=64)
    Speed: 2.00GHz
Memory:
  Total: 1024GB
Drives:
  2 x 240GB SSD boot
  10 x 7.84TB NVMe storage
NICs:
  4 x 25Gbps
------------------------

...

```