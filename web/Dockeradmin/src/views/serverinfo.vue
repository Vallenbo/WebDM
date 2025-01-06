<template>
    <div>
        <el-card>
            <el-row>
                <el-col :span="12">
                    <div ref="cpuChart" style="height: 300px;"></div>
                </el-col>
                <el-col :span="12">
                    <div ref="memoryChart" style="height: 300px;"></div>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <div ref="networkChart" style="height: 300px;"></div>
                </el-col>
                <el-col :span="12">
                    <div ref="diskChart" style="height: 300px;"></div>
                </el-col>
            </el-row>
        </el-card>
    </div>
</template>

<script>
import * as echarts from 'echarts';
import {ElRow, ElCol, ElCard} from 'element-plus';

export default {
    name: 'PerformanceDashboard',

    components: {
        ElRow,
        ElCol,
        ElCard
    },
    mounted() {
        this.renderCharts();
        this.init()
    },
    methods: {
        init() {
            let ws = new WebSocket("ws://localhost:5173/api/monitor");
            ws.onmessage = (msg) => {
                // console.log(msg)
                let axisData = new Date().toLocaleTimeString().replace(/^\D*/, '');
                let info = JSON.parse(msg.data)
                let categories = this.categories
                let cpuData = this.cpuData;
                let memData = this.memData;
                let diskRData = this.diskRData;
                let diskWData = this.diskWData;
                let netTData = this.netTData;
                let netRData = this.netRData;
                if (categories.length > 10) {
                    categories.shift();
                }
                if (cpuData.length > 10) {
                    cpuData.shift();
                }
                if (memData.length > 10) {
                    memData.shift();
                }
                if (diskRData.length > 10) {
                    diskRData.shift();
                }
                if (diskWData.length > 10) {
                    diskWData.shift();
                }
                if (netTData.length > 10) {
                    netTData.shift();
                }
                if (netRData.length > 10) {
                    netRData.shift();
                }


                categories.push(axisData)
                cpuData.push(info.CPUUsage)
                memData.push(info.CPUUsage)
                diskRData.push(this.bytesToMB(info.DiskRead))
                diskWData.push(this.bytesToMB(info.DiskWrite))
                netTData.push(this.bytesToMB(info.NetworkTX))
                netRData.push(this.bytesToMB(info.NetworkRX))


                this.cpuChart.setOption({
                    xAxis: [
                        {
                            data: categories
                        }
                    ],
                    series: [
                        {
                            data: cpuData
                        }]

                })


                this.memoryChart.setOption({
                    xAxis: [
                        {
                            data: categories
                        }
                    ],
                    series: [
                        {
                            data: memData
                        }]

                })
                this.networkChart.setOption({
                    xAxis: [
                        {
                            data: categories
                        }
                    ],
                    series: [{
                        name: 'Upload',
                        type: 'line',
                        data: netTData,
                        itemStyle: {
                            color: '#67C23A'
                        }
                    }, {
                        name: 'Download',
                        type: 'line',
                        data: netRData,
                        itemStyle: {
                            color: '#E6A23C'
                        }
                    }]

                })

                this.diskChart.setOption({
                    xAxis: [
                        {
                            data: categories
                        }
                    ],
                    series: [{
                        name: 'Read',
                        type: 'line',
                        data: diskRData,
                        itemStyle: {
                            color: '#909399'
                        }
                    }, {
                        name: 'Write',
                        type: 'line',
                        data: diskWData,
                        itemStyle: {
                            color: '#409EFF'
                        }
                    }]

                })

                // CPUUsage
                //     :
                //     1.1067708333333335
                // DiskRead
                //     :
                //     2991299584
                // DiskWrite
                //     :
                //     7607015424
                // Hostname
                //     :
                //     "LNB"
                // KernelArch
                //     :
                //     "x86_64"
                // MemoryUsage
                //     :
                //     83
                // NetworkRX
                //     :
                //     45731965
                // NetworkTX
                //     :
                //     10436544
                // OS
                //     :
                //     "windows"
                // Platform
                //     :
                //     "Microsoft Windows 11 Pro for Workstations"
                // PlatformVer
                //     :
                //     "10.0.22621 Build 22621"
                // Uptime
                //     :
                //     22698
            }
        },
        bytesToMB(bytes) {
            return (bytes / (1024 * 1024)).toFixed(2);
        },
        renderCharts() {
            // CPU chart
            const cpuChart = echarts.init(this.$refs.cpuChart);
            let categories = [];
            let cpuData = [];
            let memData = [];
            let diskRData = [];
            let diskWData = [];

            let netTData = [];
            let netRData = [];
            cpuChart.setOption({
                title: {
                    text: 'CPU Usage',
                    left: 'center'
                },
                tooltip: {
                    trigger: 'axis'
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: categories
                },
                yAxis: {
                    type: 'value',
                    axisLabel: {
                        formatter: '{value}%'
                    }
                },
                series: [{
                    name: 'CPU',
                    type: 'line',
                    data: cpuData,
                    itemStyle: {
                        color: '#F56C6C'
                    },
                    areaStyle: {}

                }]
            });

            // Memory chart
            const memoryChart = echarts.init(this.$refs.memoryChart);
            memoryChart.setOption({
                title: {
                    text: 'Memory Usage',
                    left: 'center'
                },
                tooltip: {
                    trigger: 'axis'
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: categories
                },
                yAxis: {
                    type: 'value',
                    axisLabel: {
                        formatter: '{value}%'
                    }
                },
                series: [{
                    name: 'Memory',
                    type: 'line',
                    data: memData,
                    itemStyle: {
                        color: '#409EFF'
                    }
                }]
            });

            // Network chart
            const networkChart = echarts.init(this.$refs.networkChart);
            networkChart.setOption({
                grid: {
                    // height: 600
                },
                title: {
                    text: 'Network IO',
                    left: 'center'
                },
                tooltip: {
                    trigger: 'axis'
                },
                legend: {
                    left: 10,
                    top: 10,
                    data: ['Upload', 'Download']
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: categories
                },
                yAxis: {
                    type: 'value',
                    axisLabel: {
                        formatter: '{value}KB/s'
                    }
                },
                series: [{
                    name: 'Upload',
                    type: 'line',
                    data: netTData,
                    itemStyle: {
                        color: '#67C23A'
                    }
                }, {
                    name: 'Download',
                    type: 'line',
                    data: netRData,
                    itemStyle: {
                        color: '#E6A23C'
                    }
                }]
            });

            // Disk chart
            const diskChart = echarts.init(this.$refs.diskChart);
            diskChart.setOption({
                title: {
                    text: 'Disk IO',
                    left: 'center'
                },
                tooltip: {

                    trigger: 'axis'
                },
                legend: {
                    left: 10,
                    top: 10,
                    data: ['Read', 'Write']
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: categories
                },
                yAxis: {
                    type: 'value',
                    axisLabel: {
                        formatter: '{value}KB/S'
                    }
                },
                series: [{
                    name: 'Read',
                    type: 'line',
                    data: diskRData,
                    itemStyle: {
                        color: '#909399'
                    }
                }, {
                    name: 'Write',
                    type: 'line',
                    data: diskWData,
                    itemStyle: {
                        color: '#409EFF'
                    }
                }]
            });

            this.cpuChart = cpuChart;
            this.memoryChart = memoryChart;
            this.networkChart = networkChart
            this.diskChart = diskChart
            this.categories = categories
            this.cpuData = cpuData;
            this.memData = memData;
            this.diskRData = diskRData;
            this.diskWData = diskWData;
            this.netTData = netTData;
            this.netRData = netRData;


        }
    }
};
</script>