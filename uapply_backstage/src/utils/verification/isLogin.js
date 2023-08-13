import cookies from "vue-cookies";
import router from "../../router/index"
import { message } from "../message/message"

// 使用此方法判断是否已经登录
// 主要判断方法为检查是否存在Token
export function isLogin(){
    const token = cookies.get("Token");

    if (token == null){
        // 说明没有登录，就算有登录也是出bug了，要重新登录
        // 退回到登录界面，并且通知没有登录需要重新登录
        message.info("请重新登录！");
        router.replace("/");
    }
}