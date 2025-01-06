<template>
    <el-card>
        <el-form :model="form" label-width="100px">
            <el-form-item label="镜像名">
                <el-input v-model="form.imageName"></el-input>
            </el-form-item>
            <el-form-item label="标签">
                <el-input v-model="form.tag"></el-input>
            </el-form-item>
            <el-form-item label="模板名称">
                <el-select v-model="form.templateName" placeholder="请选择">
                    <el-option v-for="item in templates" :key="item.value" :label="item.label"
                               :value="item.value"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="命令列表">
                <el-row>
                    <el-col :span="20">
                        <el-input v-model="commandInput"></el-input>
                    </el-col>
                    <el-col :span="4">
                        <el-button type="primary" @click="addCommand">添加</el-button>
                    </el-col>
                </el-row>
                <el-row v-for="(command, index) in form.commands" :key="index">
                    <el-col :span="20">
                        <el-input v-model="form.commands[index]"></el-input>
                    </el-col>
                    <el-col :span="4">
                        <el-button type="danger" @click="deleteCommand(index)">删除</el-button>
                    </el-col>
                </el-row>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm">创建</el-button>
            </el-form-item>
        </el-form>
    </el-card>
</template>

<script>
import {dockerfilereturn, CreateImage} from "../api/serviceapi"

export default {
    data() {
        return {
            form: {
                imageName: '',
                tag: '',
                templateName: '',
                commands: []
            },
            templates: [],
            commandInput: ''
        }
    },
    methods: {
        addCommand() {
            this.form.commands.push(this.commandInput)
            this.commandInput = ''
        },
        deleteCommand(index) {
            this.form.commands.splice(index, 1)
        },
        async submitForm() {
            let {imageName: name, tag, templateName, commands: templateCommand} = this.form;
            // 提交表单逻辑
            await CreateImage({name, tag, templateName, templateCommand})
            this.$router.push("/image");
        }
    },
    async created() {
        let data = await dockerfilereturn();
        let tmps = [];

        for (let tmpl of data.images) {
            tmps.push({
                label: tmpl,
                value: tmpl
            })
        }
        this.templates = tmps;
    }
}
</script>