<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>登录页面</title>
    <!-- 引入 ElementUI Plus 样式 -->
    <link rel="stylesheet" href="https://unpkg.com/element-plus/lib/theme-chalk/index.css">
  </head>
  <body>
    <div id="app">
      <!-- 创建表单 -->
      <el-form ref="loginForm" :model="loginForm" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="loginForm.password" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- 引入 ElementUI Plus 脚本 -->
    <script src="https://unpkg.com/vue@next"></script>
    <script src="https://unpkg.com/element-plus@next"></script>
    <script>
      const { reactive } = Vue;
      const { ElMessage } = elementPlus;
      const loginPage = {
        setup() {
          const loginForm = reactive({
            username: '',
            password: '',
          });
          // 提交表单
          const submitForm = () => {
            // 在这里进行登录操作
            // 如果登录成功，可以使用 ElMessage.success() 显示成功消息
            // 如果登录失败，可以使用 ElMessage.error() 显示失败消息
          };
          return {
            loginForm,
            submitForm,
          };
        },
      };
      // 创建 Vue 应用
      const app = Vue.createApp(loginPage);
      // 挂载应用
      app.mount('#app');
    </script>
  </body>
</html>