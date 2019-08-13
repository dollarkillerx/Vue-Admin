<template>
    <div class="login">
        <section class="form_container">
            <div class="manage_tip">
                <span class="title">DollarKiller通用后台管理系统</span>
                <el-form :model="loginUser" :rules="rules" ref="loginForm" label-width="80px"
                         class="loginForm">
                    <el-form-item label="邮箱" prop="email">
                        <el-input type="email" v-model="loginUser.email" placeholder="请输入Email"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input type="password" v-model="loginUser.password" placeholder="请输入密码"></el-input>
                    </el-form-item>



                    <el-form-item>
                        <el-button type="primary" class="submit_btn" @click="submitForm('loginForm')">登录 </el-button>
                    </el-form-item>
                    <div class="tiparea">
                        <p>还没有账号？现在<router-link to='/register'>注册</router-link></p>
                    </div>
                </el-form>
            </div>
        </section>
    </div>
</template>

<script>
    import base64 from 'base-64/base64'
    export default {
        name: "login",
        data() {
            return {
                loginUser: {
                    email: '',
                    password: ''
                },
                rules: {
                    email: [
                        {
                            type: "email",
                            required: true,
                            message: "邮箱格式不正确",
                            trigger: "blur"
                        }
                    ],
                    password: [
                        {required: true, message: "密码不能为空", trigger: "blur"},
                        {min: 6, max: 30, message: "长度在 6 到 30 个字符", trigger: "blur"}
                    ],
                }
            }
        },
        components: {},
        methods:{
            submitForm(formName) {
                this.$refs[formName].validate(valid => {
                    if (valid) {
                        // alert("submit!")
                        // 发送注册请求
                        this.$axios.post("/api/user/login",this.loginUser)
                            .then(res => {
                                this.$message({
                                    message:"登录成功",
                                    type: "success"
                                });
                                // 存储 token
                                const {token} = res.data;
                                // console.log(token)

                                // 存储到 localStorage DollarKillerToken
                                localStorage.setItem("DKToken",token)

                                // 解析token
                                let a = token.split(".")
                                let  b = a[0]
                                let jwt_string = base64.decode(b);
                                let jwt = JSON.parse(jwt_string)

                                // 存储到vuex
                                // 存储数据
                                this.$store.dispatch("setIsAutnenticated", !this.isEmpty(jwt));
                                this.$store.dispatch("setUser", jwt);

                                this.$router.push('/index')
                            }).catch(err => {
                            this.$message({
                                // message:"密码或者用户名错误",
                                message:err,
                                type:"error"
                            })
                        });
                    } else {
                        console.log("error submit!!");
                        return false;
                    }
                });
            },
            isEmpty(value) {
                return (
                    value === undefined ||
                    value === null ||
                    (typeof value === "object" && Object.keys(value).length === 0) ||
                    (typeof value === "string" && value.trim().length === 0)
                );
            }

        }
    }
</script>

<style scoped>
    .login {
        position: relative;
        width: 100%;
        height: 100%;
        background: url(../assets/bg.png) no-repeat center center;
        background-size: 100% 100%;
    }
    .form_container {
        width: 370px;
        height: 210px;
        position: absolute;
        top: 20%;
        left: 34%;
        padding: 25px;
        border-radius: 5px;
        text-align: center;
    }
    .form_container .manage_tip .title {
        font-family: "Microsoft YaHei";
        font-weight: bold;
        font-size: 26px;
        color: #fff;
    }
    .loginForm {
        margin-top: 20px;
        background-color: #fff;
        padding: 20px 40px 20px 20px;
        border-radius: 5px;
        box-shadow: 0px 5px 10px #cccc;
    }

    .submit_btn {
        width: 100%;
    }
    .tiparea {
        text-align: right;
        font-size: 12px;
        color: #333;
    }
    .tiparea p a {
        color: #409eff;
    }
</style>