<template>
    <el-button text @click="outerVisible = true">
        open the outer Dialog
    </el-button>

    <el-dialog v-model="outerVisible" title="Outer Dialog">
        <template #default>
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
        </template>
    </el-dialog>
</template>

<script setup>
import {ref} from 'vue'

const outerVisible = ref(false)
export default {
    name: 'VolumeDetails',
    props: ['volumes'],
}
</script>

<style scoped>
.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>