import axios from 'axios';
import { Message,Loading } from 'element-ui';
import router from './router'

// 开始加载动画
let loading;
function startLoading() {
    loading = Loading.service({
        lock: true,
        text: "DollarKiller努力加载中",
        background:'rgba(0,0,0,0,7)'
    });
}

// 结束加载动画
function endLoading() {
    loading.close();
}

// 请求拦截
axios.interceptors.request.use(config => {
    // 加载动画
    startLoading();

    // 如果存在token 设置到请求头中
    if (localStorage.DKToken) {
         // 设置统一的请求header
        config.headers.token = localStorage.DKToken
    }

    return config;
},error => {
    return Promise.reject(error)
});

// 相应拦截
axios.interceptors.response.use( response => {
    // 结束加载动画
    endLoading();
    return response;
},error => {
    // 错误提醒
    endLoading();

    // 判断token是否过期
    const {status} = error.response;
    if (status == 401) {
        Message.error("token失效，请重新登录！")
        // 清除token
        localStorage.removeItem("DKToken")
        // 跳转到登录页面
        router.push('/login')
    }

    // Message.error(error.response.data);
    return Promise.reject(error);
});
export default axios;