<template>
    <el-container>
        <!-- 顶部标题栏 -->
        <el-header height="70px"
                   style="background-color: #333; color: #fff; display: flex; justify-content: center; align-items: center;">
            <h1 style="margin: 0;">Docker_admin</h1>
        </el-header>
        <el-container>

            <!-- 左侧导航栏 -->
            <el-aside style="background-color: #f0f0f0;" width="200px">
                <el-menu :default-openeds="['1']" active-text-color="#fff" background-color="#f0f0f0" text-color="#333">
                    <el-submenu index="1">
                        <template slot="title">
                            <i class="el-icon-menu"></i>
                            <span>Navigation One</span>
                        </template>
                        <el-menu-item index="1" @click="selectedItem = 'page1'">
                            <router-link to="/image">镜像列表</router-link>
                        </el-menu-item>
                        <el-menu-item index="1" @click="selectedItem = 'page2'">
                            <router-link to="/container">容器列表</router-link>
                        </el-menu-item>
                        <el-menu-item index="1-1" @click="selectedItem = 'page7'">
                            <router-link to="/searchimage">镜像搜索</router-link>
                        </el-menu-item>
                        <el-menu-item index="1-1" @click="selectedItem = 'page6'">
                            <router-link to="/pullImage">镜像拉取</router-link>
                        </el-menu-item>
                        <el-menu-item index="1-3" @click="selectedItem = 'page3'">
                            <router-link to="/createimage">Dockerfile构建镜像</router-link>
                        </el-menu-item>
                        <el-menu-item index="1-3" @click="selectedItem = 'page4'">
                            <router-link to="/monitor">服务器性能监控</router-link>
                        </el-menu-item>
                        <el-menu-item index="1-3" @click="selectedItem = 'page5'">
                            <router-link to="/version">服务器Docker版本信息</router-link>
                        </el-menu-item>
                        <el-menu-item index="1-3" @click="selectedItem = 'page8'">
                            <router-link to="/test">test</router-link>
                        </el-menu-item>
                    </el-submenu>
                </el-menu>
            </el-aside>

            <!-- 右侧页面主体部分 -->
            <el-main style="padding: 20px;">
                <router-view/>
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
export default {
    data() {
        return {
            selectedItem: 'page1', // 默认选中第一个选项
        };
    },
};
</script>

<style scoped>
el-menu-item {
    border: 3px solid #eee;
    border-radius: 4px;
    padding: 10px;
}
</style>