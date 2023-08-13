"use strict";
import {toResume} from '/test/vc.js'
Object.defineProperty(exports, "__esModule", { value: true });
function restful(url, arg='') {
    if(arg){
        return "".concat(url, "/").concat(arg);
    }else{
        return url
    }
}
function translate(path) {
    return getApp().host + path;
}
App({
    host: "https://uapply.cloud",
    apis: {
        'user_login': '/inter/login', // 需要 code
        'inter_gerorgdep': "/inter/gerorgdep",
        'inter_turn': "/inter/turn",
        'inter_getresume': "/inter/getresume", // 需要轮次
        'inter_evaluate': "/inter/evaluate", // 需要轮次
        'inter_vieweva': "/inter/vieweva", // 需要轮次
        'inter_get-data': "/inter/get-data", // dep_id
    },
    // 用户信息
    user: {},
    // 简易封装toast函数
    _toast: function (title, icon) {
        if (icon === void 0) { icon = "none"; }
        wx.showToast({
            title: title,
            icon: icon,
            duration: 2000
        });
    },
    checkFormFull: function (formData, keys) {
        if (keys) {
            for (var _i = 0, keys_1 = keys; _i < keys_1.length; _i++) {
                var key = keys_1[_i];
                if (!formData[key]) {
                    return false;
                }
            }
            return true;
        }
        else {
            for (var d in formData) {
                if (!formData[d]) {
                    return false;
                }
            }
            return true;
        }
    },
    apiurl: function (key, arg) {
        if (arg === void 0) { arg = ''; }
        return restful(translate(this.apis[key]), arg);
    },
    err_toast: function (err) {
        this._toast(err.errMsg, 'error');
    },
    request(obj){
        /**
         * {
         *     url: {
         *         key: string,
         *         arg: any,
         *     }
         *     method: method = GET,
         *     accessToken: boolean = false,
         *     data: Object : {},
         *     fail_back: boolean: false
         * }
         */
        function failback(obj){
            if(obj.fail_back){
                setTimeout(() => {
                    wx.navigateBack()
                }, 3000)
            }
        }
        const app = this;
        if (obj.data) {
            obj.method = 'POST'
        }
        wx.request({
            url: app.apiurl(obj.url.key, obj.url.arg ? obj.url.arg : ""),
            method: obj.method ? obj.method : "GET",
            header: {
                Authorization: obj.accessToken ? app.user.accessToken : undefined,
            },
            data: {
                ...obj.data ? obj.data : {}
            },
            success(res){
                if (res.statusCode !== 200) {
                    wx.hideLoading().then(
                        wx.showToast({
                            title: `服务器出错 ${res.statusCode}`,
                            icon: 'error',
                            success(){
                                failback(obj)
                            }
                        })
                    )
                } else if (res.data.code){
                    wx.hideLoading().then(
                        wx.showToast({
                            title: res.data.msg,
                            icon: 'error',
                            success(){
                                failback(obj)
                            }
                        })
                    )
                } else {
                    obj.success(res)
                }
            },
            fail(err){
                app.err_toast(err)
                failback(obj)
            }
        })
    },
    onLaunch(){
        // toResume()
    }
});