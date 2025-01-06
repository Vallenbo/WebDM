<template>
    <div class="volume-details">
        <h1>Volume Details</h1>
        <ul>
            <li><span>Id:</span> {{ volumes.Id }}</li>
            <li><span>Created:</span> {{ volumes.Created }}</li>
            <li><span>Path:</span> {{ volumes.Path }}</li>
            <li><span>Args:</span>
                <ul>
                    <li v-for="(arg, index) in volumes.Args" :key="index">{{ arg }}</li>
                </ul>
            </li>
            <li><span>Status:</span> {{ volumes.State.Status }}</li>
            <li><span>ExitCode:</span> {{ volumes.State.ExitCode }}</li>
            <li><span>Image:</span> {{ volumes.Image }}</li>
            <li><span>Name:</span> {{ volumes.Name }}</li>
            <li><span>RestartCount:</span> {{ volumes.RestartCount }}</li>
        </ul>
    </div>
</template>

<script>
import {ContainerInspect} from "@/api/serviceapi";

export default {
    name: 'VolumeDetails',
    data() {
        return {
            ContainerInspects: [{
                "volumes": {
                    "Id": "0d0262ebdace9c5a7e3793ff35ddd6113ad89bec395e5c9ac08061f5bfd7306e",
                    "Created": "2023-04-18T02:31:40.407846008Z",
                    "Path": "/docker-entrypoint.sh",
                    "Args": [
                        "nginx",
                        "-g",
                        "daemon off;"
                    ],
                    "State": {
                        "Status": "exited",
                        "Running": false,
                        "Paused": false,
                        "Restarting": false,
                        "OOMKilled": false,
                        "Dead": false,
                        "Pid": 0,
                        "ExitCode": 0,
                        "Error": "",
                        "StartedAt": "2023-04-18T02:31:40.731743246Z",
                        "FinishedAt": "2023-04-18T02:32:44.74595922Z"
                    },
                    "Image": "sha256:605c77e624ddb75e6110f997c58876baa13f8754486b461117934b24a9dc3a85",
                    "ResolvConfPath": "/var/lib/docker/containers/0d0262ebdace9c5a7e3793ff35ddd6113ad89bec395e5c9ac08061f5bfd7306e/resolv.conf",
                    "HostnamePath": "/var/lib/docker/containers/0d0262ebdace9c5a7e3793ff35ddd6113ad89bec395e5c9ac08061f5bfd7306e/hostname",
                    "HostsPath": "/var/lib/docker/containers/0d0262ebdace9c5a7e3793ff35ddd6113ad89bec395e5c9ac08061f5bfd7306e/hosts",
                    "LogPath": "/var/lib/docker/containers/0d0262ebdace9c5a7e3793ff35ddd6113ad89bec395e5c9ac08061f5bfd7306e/0d0262ebdace9c5a7e3793ff35ddd6113ad89bec395e5c9ac08061f5bfd7306e-json.log",
                    "Name": "/romantic_sammet",
                    "RestartCount": 0,
                    "Driver": "overlay2",
                    "Platform": "linux",
                    "MountLabel": "",
                    "ProcessLabel": "",
                    "AppArmorProfile": "",
                    "ExecIDs": null,
                    "HostConfig": {
                        "Binds": null,
                        "ContainerIDFile": "",
                        "LogConfig": {
                            "Type": "json-file",
                            "Config": {}
                        },
                        "NetworkMode": "default",
                        "PortBindings": {
                            "80/tcp": [
                                {
                                    "HostIp": "",
                                    "HostPort": ""
                                }
                            ]
                        },
                        "RestartPolicy": {
                            "Name": "no",
                            "MaximumRetryCount": 0
                        },
                        "AutoRemove": false,
                        "VolumeDriver": "",
                        "VolumesFrom": null,
                        "ConsoleSize": [
                            35,
                            158
                        ],
                        "CapAdd": null,
                        "CapDrop": null,
                        "CgroupnsMode": "private",
                        "Dns": [],
                        "DnsOptions": [],
                        "DnsSearch": [],
                        "ExtraHosts": null,
                        "GroupAdd": null,
                        "IpcMode": "private",
                        "Cgroup": "",
                        "Links": null,
                        "OomScoreAdj": 0,
                        "PidMode": "",
                        "Privileged": false,
                        "PublishAllPorts": false,
                        "ReadonlyRootfs": false,
                        "SecurityOpt": null,
                        "UTSMode": "",
                        "UsernsMode": "",
                        "ShmSize": 67108864,
                        "Runtime": "runc",
                        "Isolation": "",
                        "CpuShares": 0,
                        "Memory": 0,
                        "NanoCpus": 0,
                        "CgroupParent": "",
                        "BlkioWeight": 0,
                        "BlkioWeightDevice": [],
                        "BlkioDeviceReadBps": [],
                        "BlkioDeviceWriteBps": [],
                        "BlkioDeviceReadIOps": [],
                        "BlkioDeviceWriteIOps": [],
                        "CpuPeriod": 0,
                        "CpuQuota": 0,
                        "CpuRealtimePeriod": 0,
                        "CpuRealtimeRuntime": 0,
                        "CpusetCpus": "",
                        "CpusetMems": "",
                        "Devices": [],
                        "DeviceCgroupRules": null,
                        "DeviceRequests": null,
                        "MemoryReservation": 0,
                        "MemorySwap": 0,
                        "MemorySwappiness": null,
                        "OomKillDisable": null,
                        "PidsLimit": null,
                        "Ulimits": null,
                        "CpuCount": 0,
                        "CpuPercent": 0,
                        "IOMaximumIOps": 0,
                        "IOMaximumBandwidth": 0,
                        "MaskedPaths": [
                            "/proc/asound",
                            "/proc/acpi",
                            "/proc/kcore",
                            "/proc/keys",
                            "/proc/latency_stats",
                            "/proc/timer_list",
                            "/proc/timer_stats",
                            "/proc/sched_debug",
                            "/proc/scsi",
                            "/sys/firmware"
                        ],
                        "ReadonlyPaths": [
                            "/proc/bus",
                            "/proc/fs",
                            "/proc/irq",
                            "/proc/sys",
                            "/proc/sysrq-trigger"
                        ]
                    },
                    "GraphDriver": {
                        "Data": {
                            "LowerDir": "/var/lib/docker/overlay2/fe98f4fbd54af7d95fddf9fb0e232f060931651ec82aa2a7e5c0246a99bce4bf-init/diff:/var/lib/docker/overlay2/39eeec4a6356279605e500b2ee1cbd5b017ee46431d3c2a98d495aada0efcb23/diff:/var/lib/docker/overlay2/22b0604d215b15281e2b88d1b60b7d8cc4fc0108a7adef82c209a260ca90cee4/diff:/var/lib/docker/overlay2/6ba3177ef12d42f6e68bfbd13310f704d0ca5ae32ef92b58c8f0a3a34739a29f/diff:/var/lib/docker/overlay2/dddf1b5a35f171a948a4cf8e67b98c22889e29c83f4b256a33562881d2efeca4/diff:/var/lib/docker/overlay2/c6c1b756e24dac59e14775f42414e2a932789917e0320ed234abb93b453c63a4/diff:/var/lib/docker/overlay2/e53b9b5714ff33481f215c6e6a099362cff13fa10bb4256ad7c31d6fca7e879c/diff",
                            "MergedDir": "/var/lib/docker/overlay2/fe98f4fbd54af7d95fddf9fb0e232f060931651ec82aa2a7e5c0246a99bce4bf/merged",
                            "UpperDir": "/var/lib/docker/overlay2/fe98f4fbd54af7d95fddf9fb0e232f060931651ec82aa2a7e5c0246a99bce4bf/diff",
                            "WorkDir": "/var/lib/docker/overlay2/fe98f4fbd54af7d95fddf9fb0e232f060931651ec82aa2a7e5c0246a99bce4bf/work"
                        },
                        "Name": "overlay2"
                    },
                    "Mounts": [],
                    "Config": {
                        "Hostname": "0d0262ebdace",
                        "Domainname": "",
                        "User": "",
                        "AttachStdin": false,
                        "AttachStdout": false,
                        "AttachStderr": false,
                        "ExposedPorts": {
                            "80/tcp": {}
                        },
                        "Tty": false,
                        "OpenStdin": false,
                        "StdinOnce": false,
                        "Env": [
                            "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                            "NGINX_VERSION=1.21.5",
                            "NJS_VERSION=0.7.1",
                            "PKG_RELEASE=1~bullseye"
                        ],
                        "Cmd": [
                            "nginx",
                            "-g",
                            "daemon off;"
                        ],
                        "Image": "nginx",
                        "Volumes": null,
                        "WorkingDir": "",
                        "Entrypoint": [
                            "/docker-entrypoint.sh"
                        ],
                        "OnBuild": null,
                        "Labels": {
                            "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>"
                        },
                        "StopSignal": "SIGQUIT"
                    },
                    "NetworkSettings": {
                        "Bridge": "",
                        "SandboxID": "52351db568efd3046d11688959ba4f9bd73c886a68cfa3d9cc344aa7e670cc0f",
                        "HairpinMode": false,
                        "LinkLocalIPv6Address": "",
                        "LinkLocalIPv6PrefixLen": 0,
                        "Ports": {},
                        "SandboxKey": "/var/run/docker/netns/52351db568ef",
                        "SecondaryIPAddresses": null,
                        "SecondaryIPv6Addresses": null,
                        "EndpointID": "",
                        "Gateway": "",
                        "GlobalIPv6Address": "",
                        "GlobalIPv6PrefixLen": 0,
                        "IPAddress": "",
                        "IPPrefixLen": 0,
                        "IPv6Gateway": "",
                        "MacAddress": "",
                        "Networks": {
                            "bridge": {
                                "IPAMConfig": null,
                                "Links": null,
                                "Aliases": null,
                                "NetworkID": "573845aa3958663cd40d538c8683708283d68a57d1d5f5f6156761bc2d5515bc",
                                "EndpointID": "",
                                "Gateway": "",
                                "IPAddress": "",
                                "IPPrefixLen": 0,
                                "IPv6Gateway": "",
                                "GlobalIPv6Address": "",
                                "GlobalIPv6PrefixLen": 0,
                                "MacAddress": "",
                                "DriverOpts": null
                            }
                        }
                    }
                }
            }],
        }
    },
    methods: {
        getInspect() {
            ContainerInspect.then(data => {

            })
        }
    },
};
</script>

<style>
.volume-details ul {
    list-style-type: none;
}

.volume-details li span {
    font-weight: bold;
}
</style>