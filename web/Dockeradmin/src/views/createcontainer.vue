<template>
    <div class="container-form">
        <el-form :model="form" label-width="120px">
            <el-form-item label="容器名">
                <el-input v-model="form.containerName"></el-input>
            </el-form-item>
            <el-form-item label="镜像名">
                <el-input v-model="form.imageName"></el-input>
            </el-form-item>
            <el-form-item label="端口映射列表">
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="容器端口">
                            <el-input v-model="newPortMapping.containerPort"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="主机端口">
                            <el-input v-model="newPortMapping.hostPort"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-button size="small" type="primary" @click="addPortMapping">添加端口映射</el-button>
                <el-table :data="form.portMappings" style="margin-top: 10px;">
                    <el-table-column label="容器端口" prop="containerPort"></el-table-column>
                    <el-table-column label="主机端口" prop="hostPort"></el-table-column>
                    <el-table-column label="操作">
                        <template #default="scope">
                            <el-button size="small" type="danger" @click="deletePortMapping(scope.$index)">删除
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-form-item>
            <el-form-item label="环境变量列表">
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="键">
                            <el-input v-model="newEnvVariable.key"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="值">
                            <el-input v-model="newEnvVariable.value"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-button size="small" type="primary" @click="addEnvVariable">添加环境变量</el-button>
                <el-table :data="form.envVariables" style="margin-top: 10px;">
                    <el-table-column label="键" prop="key"></el-table-column>
                    <el-table-column label="值" prop="value"></el-table-column>
                    <el-table-column label="操作">
                        <template #default="scope">
                            <el-button size="small" type="danger" @click="deleteEnvVariable(scope.$index)">删除
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm">创建容器</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import {CreateContainer} from "@/api/serviceapi"

export default {
    data() {
        return {
            form: {
                containerName: '',
                imageName: '',
                portMappings: [],
                envVariables: [],
            },
            newPortMapping: {
                containerPort: '',
                hostPort: '',
            },
            newEnvVariable: {
                key: '',
                value: '',
            },
        };
    },
    methods: {
        addPortMapping() {
            this.form.portMappings.push({
                containerPort: this.newPortMapping.containerPort,
                hostPort: this.newPortMapping.hostPort,
            });
            this.newPortMapping.containerPort = '';
            this.newPortMapping.hostPort = '';
        },
        deletePortMapping(index) {
            this.form.portMappings.splice(index, 1);
        },
        addEnvVariable() {
            this.form.envVariables.push({
                key: this.newEnvVariable.key,
                value: this.newEnvVariable.value,
            });
            this.newEnvVariable.key = '';
            this.newEnvVariable.value = '';
        },
        deleteEnvVariable(index) {
            this.form.envVariables.splice(index, 1);
        },
        submitForm() {
            console.log('Submitting form:', this.form);
            let {imageName: image, containerName: name, envVariables: env, portMappings: portMap} = this.form;
            let newEnv = [];
            let newPorMap = [];

            for (let k in env) {
                newEnv.push({k})
            }
            CreateContainer({name, image, env, portMap}).then(data => {
                this.$router.push("/container")
            })
        },
    },
    created() {
        let image = localStorage.getItem("image");
        this.form.imageName = image
    }
};
</script>

<style>
.container-form {
    padding: 20px;
}
</style>