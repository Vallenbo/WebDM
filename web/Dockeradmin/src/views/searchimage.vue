<template>
    <el-container>
        <el-header>
            <h1>搜索 Docker 镜像</h1>
        </el-header>
        <el-main>
            <el-row :gutter="20">
                <el-col :span="12">
                    <el-input placeholder="输入镜像名" v-model="imageName"></el-input>
                </el-col>
                <el-col :span="12">
                    <el-button type="primary" @click="searchImage">搜索</el-button>
                </el-col>
            </el-row>
            <el-divider></el-divider>
            <el-row :gutter="20">
                <el-col :span="24">
                    <h2>搜索结果</h2>
                    <el-table :data="searchResult" style="width: 100%">
                        <el-table-column prop="name" label="名称"></el-table-column>
                        <el-table-column prop="star_count" label="推荐星数"></el-table-column>
                        <el-table-column prop="is_official" label="官方镜像">
                            <template #scope>
                                <el-tag v-if="scope.row.is_official" type="success">是</el-tag>
                                <el-tag v-else type="info">否</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="is_automated" label="自制镜像">
                            <template #scope>
                                <el-tag v-if="scope.row.is_automated" type="success">是</el-tag>
                                <el-tag v-else type="info">否</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="description" label="描述"></el-table-column>
                    </el-table>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</template>


<script>
import {pullImage, SearchImage} from "@/api/serviceapi"

export default {
    data() {
        return {
            imageName: '',
            searchResult: [
                {
                    "name": "nginx",
                    "star_count": 18377,
                    "is_official": true,
                    "is_automated": false,
                    "description": "Official build of Nginx."
                }
            ],
            columns: [
                {
                    label: '镜像名称',
                    prop: 'name',
                },
                {
                    label: '推荐数',
                    prop: 'star_count',
                },
                {
                    label: '是否官方',
                    prop: 'is_official',
                },
                {
                    label: '是否自制',
                    prop: 'is_automated',
                },
                {
                    label: '镜像描述',
                    prop: 'description',
                },
            ],
        };
    },
    methods: {
        async searchImage() {
            // 在此处编写搜索 Docker 镜像的逻辑
            // 将搜索结果存储在 this.searchResult 中，格式类似于以下示例数据
            let resp = await SearchImage(this.imageName);
            this.searchResult = resp.message
        },
    },
};
</script>