import {vc, uvc} from '/test/vc.js'

function restful(url, arg){
	return `${url}/${arg}`
}

function translate(path){
	return getApp().host + path
}

App({
	host: "https://uapply.cloud", // 远程主机
	apis: {
		'user_login': '/user/login',
		'user_check-resume' : "/user/check-resume",
		'user_resume' : "/user/resume",
		'user_org-enrolling': "/user/org-enrolling",
		'user_org-deps': "/user/org-deps",
		'user_register': "/user/register",
	},
	// 用户信息
	user: {}, 
	// 组织
	organization: [],
	// 简易封装toast函数
	_toast(title, icon="none"){
		wx.showToast({
		  	title,
			icon,
			duration: 2000
		})
	},
	checkFormFull(formData, keys){
		if(keys){
			for (const key of keys){
				if(!formData[key]){
					return false
				}
			}
			return true
		}else{
			for(const d in formData){
				if(!formData[d]){
					return false
				}
			}
			return true
		}
	},
	apiurl(key, arg=''){
		return restful(translate(this.apis[key]), arg)
	},
	err_toast(err){
		this._toast(err.errMsg, 'error')
	},
	// 转换path为api的url
	onLaunch(){
		// 获取上次登陆时间与用户信息
		this.get_user()
	},
	// 尝试直接从缓存中获取信息
	get_user: function () {
		// uvc()
		var last_storage_time = wx.getStorageSync('lst');
		var user = wx.getStorageSync('user');
		if (user) {
			if (Date.now() - last_storage_time < user.accessExpire) {
				this.user = user;
				this._toast('自动登陆成功', 'success');
			}
			else {
				this._toast('登陆过期，请重新登陆');
			}
		}
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