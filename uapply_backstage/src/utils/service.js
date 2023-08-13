// 使用Naice-ui进行请求失败的消息通知
// 要使Message API脱离setup使用，得通过createDiscreteApi创建对应API才可以在JS中调用
// 调用需要使用到的模块
import { message } from './message/message';
import cookies from "vue-cookies";

/****   request.js   ****/
// 导入axios
import axios from 'axios'
// 使用element-ui Message做消息提醒
// import { Message } from 'element-ui';
//1. 创建新的axios实例，
const service = axios.create({
  // 公共接口--process.env.BASE_API是webpack中的全局环境变量，但是这里可以写死
  // baseURL: process.env.BASE_API,
  baseURL: 'https://uapply.cloud',
  // 超时时间 单位是ms，这里设置了3s的超时时间
  timeout: 3 * 1000
})
// 2.请求拦截器
service.interceptors.request.use(config => {
  //发请求前做的一些处理，数据转化，配置请求头，设置token,设置loading等，根据需求去添加
  config.data = JSON.stringify(config.data); //数据转化,也可以使用qs转换
  config.headers = {
    'Content-Type': 'application/json' //配置请求头
  }
  //如有需要：注意使用token的时候需要引入cookie方法或者用本地localStorage等方法，推荐js-cookie
  const token = cookies.get('Token');//这里取token之前，你肯定需要先拿到token,存一下
  if(token){
    // config.params = {'Authorization':token} //如果要求携带在参数中
    config.headers.Authorization= token; //如果要求携带在请求头中
  }
  return config
}, error => {
  Promise.reject(error)
})

// 3.响应拦截器
service.interceptors.response.use(response => {
  //接收到响应数据并成功后的一些共有的处理，关闭loading等

  return response
}, error => {
  /***** 接收到异常响应的处理开始 *****/
  if (error && error.response) {
    // 1.公共错误处理
    // 2.根据响应码具体处理
    let ERR = "";
    switch (error.response.status) {
      case 0:
        ERR = "您已离线，请检查网络设置";
        break;
      case 400:
        ERR = '错误请求'
        break;
      case 401:
        ERR = '未授权，请重新登录'
        break;
      case 403:
        ERR = '拒绝访问';
        break;
      case 404:
        ERR = '请求错误,未找到该资源'
        window.location.href = "/error_pages/error_404"
        break;
      case 405:
        ERR = '请求方法未允许'
        break;
      case 408:
        ERR = '请求超时'
        break;
      case 500:
        ERR = '服务器端出错'
        break;
      case 501:
        ERR = '网络未实现'
        break;
      case 502:
        ERR = '网络错误'
        break;
      case 503:
        ERR = '服务不可用'
        break;
      case 504:
        ERR = '网络超时'
        break;
      case 505:
        ERR = 'http版本不支持该请求'
        break;
      default:
        ERR = `连接错误${error.response.status}`
    }
    message.info(ERR);
  } else {
    // 超时处理
    if (JSON.stringify(error).includes('timeout')) {
      message.info('服务器响应超时，请刷新当前页')
    }
    message.info('连接服务器失败');
  }

  // Message.error(error.message)
  /***** 处理结束 *****/
  //如果不需要错误处理，以上的处理过程都可省略
  return Promise.resolve(error.response)
})
//4.导入文件
export default service
