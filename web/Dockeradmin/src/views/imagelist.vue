<template>
    <div class="image-list">
        <el-table :data="images" style="width: 100%">
            <el-table-column label="镜像ID" prop="name"></el-table-column>
            <el-table-column label="标签" prop="tag"></el-table-column>
            <el-table-column label="大小(MB)" prop="size"></el-table-column>
            <el-table-column label="创建时间" prop="createdTime">
                <template #default="scope">
                    {{ formatDate(scope.row.createdTime * 1000) }}
                </template>
            </el-table-column>

            <el-table-column label="操作">
                <template #default="scope">
                    <el-button size="small" type="primary" @click="startContainer(scope.$index)">创建容器</el-button>
                    <el-button size="small" type="danger" @click="deleteImage(scope.$index)">删除镜像</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import {GetAllImages, StopremoveContainer} from "../api/serviceapi"
import {DeleteImage} from "../api/serviceapi"
import {ElMessage} from "element-plus";

export default {
    data() {
        return {
            images: [
                {
                    "name": "sha256:0b345b61f6808139ae02ea4178d6b85d53671551d7791a73e769ba171e763374",
                    "tag": "n3:latest",
                    "digest": "",
                    "size": 231268856,
                    "createdTime": 1631730005,
                    "author": "",
                    "architecture": "",
                    "os": ""
                },
            ],
        };
    },
    methods: {
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
            // console.log(`Starting container for image ${this.images[index].name}:${this.images[index].tag}`);
            // 在这里实现启动容器的逻辑
            // console.log(index)
            let image = this.images[index];
            localStorage.setItem("image", image.tag)
            this.$router.push("/image/container/1")
        },
        async deleteImage(index) {
            // 在这里实现删除镜像的逻辑
            let image = this.images[index];
            await DeleteImage(image.name)
            ElMessage({
                message: 'Congrats, this is a success message.',
                type: 'success',
            })
            GetAllImages().then(data => {
                this.images = data.images;
            })
        },
    },

    created() {
        GetAllImages().then(data => {
            this.images = data.images;
        })
    }
};
</script>

<style>
.image-list {
    padding: 20px;
}
</style>