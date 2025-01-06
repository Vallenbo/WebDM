<template>
    <div class="image-list">
        <el-table :data="Containers" style="width: 100%">
            <el-table-column label="容器id" prop="id"></el-table-column>
            <el-table-column label="容器名" prop="name">
                <template #default="scope">
                    <a href="javascript:;" @click="getContainerInfo(scope.row.id)">{{ scope.row.name }}</a>
                </template>
            </el-table-column>
            <el-table-column label="使用镜像" prop="image"></el-table-column>
            <el-table-column label="端口映射">
                <template #default="scope">
                    <span v-for="port in scope.row.ports"
                          :key="port.PublicPort">({{ port.IP }}:{{ port.PrivatePort }}) ->{{
                        port.PublicPort
                        }}/{{ port.Type }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态" prop="status"></el-table-column>
            <el-table-column label="创建时间" prop="createdTime">
                <template #default="scope">
                    {{ formatDate(scope.row.createdTime * 1000) }}
                </template>
            </el-table-column>
            <el-table-column label="COMMEND" prop="command"></el-table-column>

            <el-table-column label="操作">
                <template #default="scope">
                    <el-button size="small" type="primary" @click="startContainer(scope.$index)">启动容器</el-button>
                    <el-button size="small" type="info" @click="stopContainer(scope.$index)">停止容器</el-button>
                    <el-button size="small" type="danger" @click="deleteContainer(scope.$index)">删除容器</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>

    <el-dialog v-model="dialogVisible" width="50%">
        <div slot="title">Docker 容器信息</div>

        <div class="modal-body">
            <table class="table table-striped">
                <tr>
                    <td><strong>ID:</strong></td>
                    <td>{{ container.Id }}</td>
                </tr>
                <tr>
                    <td><strong>状态:</strong></td>
                    <td>{{ container.State.Status }}</td>
                </tr>
                <tr>
                    <td><strong>创建时间:</strong></td>
                    <td>{{ container.Created }}</td>
                </tr>
                <tr>
                    <td><strong>镜像:</strong></td>
                    <td>{{ container.Config.Image }}</td>
                </tr>
                <tr>
                    <td><strong>绑定的端口:</strong></td>
                    <td>
                        <ul>
                            <li v-for="(port, index) in container.NetworkSettings.Ports" :key="index">
                                <span v-if="port.length === 0">{{ index }}</span>
                                <ul>
                                    <li v-for="(binding, i) in port" :key="i">{{ binding.HostIp }}:{{
                                        binding.HostPort
                                        }}
                                    </li>
                                </ul>
                            </li>
                        </ul>
                    </td>
                </tr>
            </table>
        </div>
        <div slot="footer" class="dialog-footer">
            <el-button @click="dialogVisible = false">关闭</el-button>
        </div>

    </el-dialog>
</template>

<script>
import {ListContainers, RunContainer, StopContainer, StopremoveContainer, containerInfo} from "@/api/serviceapi";
import {ElMessage} from 'element-plus';

export default {
    data() {
        return {
            dialogVisible: false,
            container: {},
            Containers: [
                {
                    "id": "ab86172c962996af2c5b32e26636526b3094e1f4a3499325cc5a4c444e540979",
                    "name": "romantic_mahavira",
                    "image": "nginx",
                    "createdTime": "",
                    "status": "",
                    "ip": "",
                    "publicPort": "",
                    "privatePort": "",
                    "labels": "",
                    "ports": [],
                },
            ],
        };
    },
    filters: {},
    methods: {
        getContainerInfo(id) {
            containerInfo(id).then(c => {
                console.log(c);
                this.dialogVisible = true;
                this.container = c.volumes;
            })
        },
        formatDate(timestamp) {
            const date = new Date(timestamp);
            const year = date.getFullYear();
            const month = ('0' + (date.getMonth() + 1)).slice(-2);
            const day = ('0' + date.getDate()).slice(-2);
            const hours = ('0' + date.getHours()).slice(-2);
            const minutes = ('0' + date.getMinutes()).slice(-2);
            const seconds = ('0' + date.getSeconds()).slice(-2);
            return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
        },
        startContainer(index) {
            RunContainer(this.Containers[index].id).then(data => {
                this.ListContainers();
                ElMessage({
                    message: 'Congrats, this is a success message.',
                    type: 'success',
                });
            });
        },
        stopContainer(index) {
            StopContainer(this.Containers[index].id).then(data => {
                this.ListContainers();
                ElMessage({
                    message: 'Congrats, this is a success message.',
                    type: 'success',
                });
            });
        },
        deleteContainer(index) {
            StopremoveContainer(this.Containers[index].id).then(data => {
                this.ListContainers();
                ElMessage({
                    message: 'Congrats, this is a success message.',
                    type: 'success',
                });
            });
        },
        ListContainers() {
            ListContainers().then(data => {
                this.Containers = data.containers;
                this.Containers.forEach(container => {
                    container.ports = [container.ports];
                    container.ports.forEach(port => {
                        port.publicPort = port.PublicPort || "-";
                        port.privatePort = port.PrivatePort || "-";
                        port.type = port.Type || "-";
                    });
                });
            });
        },
    },
    created() {
        this.ListContainers();
    }
};
</script>

<style>
.image-list {
    padding: 20px;
}
</style>