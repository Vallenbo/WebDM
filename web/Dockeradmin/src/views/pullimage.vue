<template>
    <el-container>
        <el-header>
            <h1>拉取Docker镜像</h1>
        </el-header>
        <el-main>
            <el-row>
                <el-col :span="24">
                    <el-form ref="form" :model="form" label-position="top">
                        <el-form-item label="镜像名">
                            <el-input v-model="form.imageName" placeholder="请输入镜像名"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click="pullImage">拉取镜像</el-button>
                        </el-form-item>
                    </el-form>
                </el-col>

            </el-row>
            <el-row>
                <el-col :span="24">
                    <el-card>
                        <div slot="header" class="clearfix">
                            <span>拉取信息</span>
                            <el-button icon="el-icon-close" style="float: right; padding: 3px 0" type="text"
                                       @click="clearLog"></el-button>
                        </div>
                        <el-scrollbar v-loading="loading" style="height: 500px">
                            <el-timeline>
                                <el-timeline-item v-for="(log, index) in logs" :key="index"
                                                  :timestamp="new Date().toLocaleString()">
                                    <p>{{ log }}</p>
                                </el-timeline-item>
                            </el-timeline>
                        </el-scrollbar>
                    </el-card>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</template>

<script>
import {pullImage} from "@/api/serviceapi"

export default {
    data() {
        return {
            form: {
                imageName: ''
            },
            logs: [],
            loading: false,
        }
    },
    methods: {
        async pullImage() {
            const imageName = this.form.imageName.trim()
            // if (!imageName) {
            //     this.$message.error('请输入镜像名')
            //     return
            // }
            // this.logs.push({
            //     time: new Date(),
            //     message: `开始拉取镜像 ${imageName} ...`
            // })
            this.loading = true;
            try {
                // 使用 axios 或其他库发送拉取镜像的命令并获取输出信息
                let resp = await pullImage(imageName);
                let text = await resp.text();
                this.logs = text.split("\n");
                this.loading = false
            } catch (err) {
                // this.logs.push({
                //     time: new Date(),
                //     message: `拉取镜像 ${imageName} 失败：${err.message}`
                // })
            }
        },
        clearLog() {
            this.logs = []
        }
    }
}
</script>