import {vc, uvc} from '/test/vc.js'

App({
	host: "https://uapply.cloud", // 远程主机
	apis: {
		'user_login': '/user/login',
		'user_check-resume' : "/user/check-resume",
		'user_resume' : "/user/resume",
		'user_org-enrolling': "/user/org-enrolling",
		'user_org-deps': "/user/org-deps",
		'user_register': "/user/register",
		'nice': "/user/mini-code"
	},
	// 版本code值
	code: 0,
	// 用户信息
	user: {}, 
	// 组织
	organization: [],
	// 简易封装toast函数
	translate(path){
		return this.host + path
	},
	restful(url, arg){
		return `${url}/${arg}`
	},
	_toast(title, icon="none"){
		wx.showToast({
		  	title,
			icon,
			duration: 2000
		})
	},
	checkFormFull(formData, keys){
		if (keys){
			for (const key of keys){
				if(!formData[key]){
					return false
				}
			}
		}else{
			for(const d in formData){
				if(!formData[d]){
					return false
				}
			}
		}
		return true
	},
	apiurl(key, arg=''){
		return this.restful(this.translate(this.apis[key]), arg)
	},
	err_toast(err){
		this._toast(err.errMsg, 'error')
	},
	// 转换path为api的url
	onLaunch(){
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
})